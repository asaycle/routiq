package location

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/asaycle/routiq/pkg/domain/entity"
	"github.com/asaycle/routiq/pkg/domain/repository"
	"github.com/asaycle/routiq/pkg/infrastructure/db"
	"github.com/asaycle/routiq/pkg/lib/config"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"golang.org/x/xerrors"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import master data (prefs, cities)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Please specify a subcommand (e.g., grpc, gateway)")
		ctx := cmd.Context()
		cfg := config.FromContext(ctx)
		return importLocation(ctx, cfg)
	},
}

func importLocation(ctx context.Context, cfg *config.Config) error {
	xr, err := download(ctx, cfg.App.LocationConfig.SourceURL)
	if err != nil {
		return xerrors.Errorf("failed to download file: %w", err)
	}
	defer xr.Close()

	// 2) Excelを開く（ファイルに落とさずReaderで）
	xf, err := excelize.OpenReader(xr)
	if err != nil {
		return fmt.Errorf("open xlsx: %w", err)
	}
	defer xf.Close()

	sheet := xf.GetSheetName(0)
	slog.Info("Found sheet", "name", sheet)
	hm := headerMap{codeCol: 0, prefCol: 1, cityCol: 2}
	records, err := readRecords(xf, sheet, hm, 3)
	if err != nil {
		return err
	}
	slog.Info("Read records from Excel", "count", len(records))
	if len(records) == 0 {
		return xerrors.New("no rows parsed from Excel")
	}
	prefs := make(map[string]*entity.Pref) // key: pref name
	cities := make(map[string]*entity.City)
	for _, record := range records {
		// pref
		if record.CityName == "" {
			pref := entity.NewPref(record.Code, record.PrefName)
			prefs[pref.DisplayName] = pref
			continue
		}
		// city
		parentPref := prefs[record.PrefName]
		city := entity.NewCity(record.Code, parentPref.ID, record.CityName)
		slog.Info("Record", "code", record.Code, "pref", record.PrefName, "city", record.CityName)
		cities[city.ID] = city
	}

	pgdb, err := db.NewPgDBFromCfg(&cfg.DB)
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	txManager := repository.NewTransactionManager(pgdb)
	repo := repository.NewLocationRepository()

	for _, pref := range prefs {
		slog.Info("Inserting Pref", "pref", pref)
		err := txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
			if err := repo.CreatePref(ctx, tx, pref); err != nil {
				return xerrors.Errorf("Error CreatePref: %w", err)
			}
			return nil
		})
		if err != nil {
			return xerrors.Errorf("Failed Transaction for Pref: %w", err)
		}
	}

	for _, city := range cities {
		slog.Info("Inserting City", "city", city)
		err := txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
			if err := repo.CreateCity(ctx, tx, city); err != nil {
				return xerrors.Errorf("Error CreateCity: %w", err)
			}
			return nil
		})
		if err != nil {
			return xerrors.Errorf("Failed Transaction for Pref: %w", err)
		}
	}

	return nil
}

func download(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	// 呼び出し側でCloseする
	return resp.Body, nil
}

type rowRecord struct {
	Code     string // 5桁 JIS X 0402
	PrefName string
	CityName string
}

func readRecords(xf *excelize.File, sheet string, hm headerMap, startRow int) ([]rowRecord, error) {
	rows, err := xf.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	out := make([]rowRecord, 0, max(0, len(rows)-startRow))
	for i := startRow; i < len(rows); i++ {
		r := rows[i]
		get := func(idx int) string {
			if idx >= 0 && idx < len(r) {
				return strings.TrimSpace(r[idx])
			}
			return ""
		}
		code := normalizeCode5(get(hm.codeCol))
		if len(code) != 5 {
			continue
		}
		rec := rowRecord{
			Code:     code,
			PrefName: strings.TrimSpace(get(hm.prefCol)),
			CityName: strings.TrimSpace(get(hm.cityCol)),
		}
		if rec.PrefName == "" && rec.CityName == "" {
			continue
		}
		out = append(out, rec)
	}
	return out, nil
}

func normalizeCode5(s string) string {
	s = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(s, "　", ""), " ", ""))
	if i := strings.Index(s, "."); i > 0 { // "25201.0"対策
		s = s[:i]
	}
	if len(s) == 4 {
		return "0" + s
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type headerMap struct {
	codeCol     int
	prefCol     int
	cityCol     int
	gunCol      *int
	seireiKuCol *int
}

func detectHeaders(xf *excelize.File, sheet string) (headerMap, int, error) {
	rows, err := xf.GetRows(sheet)
	if err != nil {
		return headerMap{}, 0, err
	}
	for i := 0; i < min(5, len(rows)); i++ {
		h := rows[i]
		hm := headerMap{codeCol: 0, prefCol: 1, cityCol: 2}
		for colIdx, v := range h {
			n := normalizeHeader(v)
			switch {
			case matchAny(n, "団体コード", "jisコード", "jis code", "code"):
				hm.codeCol = colIdx
			case matchAny(n, "都道府県名", "都道府県"):
				hm.prefCol = colIdx
			case matchAny(n, "市区町村名", "市区町村", "自治体名", "団体名"):
				hm.cityCol = colIdx
			case matchAny(n, "郡名", "郡"):
				c := colIdx
				hm.gunCol = &c
			case matchAny(n, "政令市区名", "区名", "政令区", "政令指定市区名"):
				c := colIdx
				hm.seireiKuCol = &c
			}
		}
		if hm.codeCol >= 0 && hm.prefCol >= 0 && hm.cityCol >= 0 {
			return hm, i + 1, nil
		}
	}
	return headerMap{}, 0, errors.New("header row not detected")
}

func matchAny(n string, cands ...string) bool {
	n = strings.ToLower(n)
	for _, c := range cands {
		c = strings.ToLower(c)
		if n == c || strings.Contains(n, c) {
			return true
		}
	}
	return false
}

func findTargetSheet(xf *excelize.File) (string, error) {
	return xf.GetSheetName(0), nil
}

func headerHas(h []string, keys []string) bool {
	norm := make([]string, len(h))
	for i, x := range h {
		norm[i] = normalizeHeader(x)
	}
	contains := func(needle string) bool {
		needle = normalizeHeader(needle)
		for _, col := range norm {
			if col == needle || strings.Contains(col, needle) {
				return true
			}
		}
		return false
	}
	for _, k := range keys {
		if !contains(k) {
			return false
		}
	}
	return true
}

var spaceRe = regexp.MustCompile(`\s+`)

func normalizeHeader(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "　", " ")
	s = spaceRe.ReplaceAllString(s, " ")
	return strings.ToLower(s)
}
