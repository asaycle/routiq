package main

import (
	"fmt"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"title": func(s string) string {
		if s == "" {
			return ""
		}
		return strings.ToUpper(s[:1]) + s[1:]
	},
	"seg": func(s string) string {
		if strings.HasPrefix(s, "{") {
			return fmt.Sprintf("_seg{_segVar, %q}", strings.Trim(s, "{}"))
		}
		return fmt.Sprintf("_seg{_segLit, %q}", s)
	},
	"fmtArgs": func(segs []string) string {
		// "users/{user}" → "user string"
		vars := []string{}
		seen := map[string]bool{}
		for _, s := range segs {
			if strings.HasPrefix(s, "{") {
				v := strings.Trim(s, "{}")
				if !seen[v] {
					vars = append(vars, fmt.Sprintf("%s string", v))
					seen[v] = true
				}
			}
		}
		return strings.Join(vars, ", ")
	},
	"fmtJoin": func(segs []string) string {
		parts := make([]string, 0, len(segs))
		for _, s := range segs {
			if strings.HasPrefix(s, "{") {
				v := strings.Trim(s, "{}")
				parts = append(parts, v)
			} else {
				parts = append(parts, fmt.Sprintf("%q", s))
			}
		}
		return strings.Join(parts, " + \"/\" + ")
	},
	"join": strings.Join,
}

// ランタイム本文（ヘッダは main.go で付与）
var runtimeBodyTmpl = template.Must(
	template.New("runtimeBody").
		Option("missingkey=error").
		Parse(`import (
	"errors"
	"strings"
)

type _segKind uint8
const (
	_segLit _segKind = iota
	_segVar
)
type _seg struct{ k _segKind; s string }
type _pattern []_seg

func _match(segs []string, idx int, p _pattern, vars map[string]string) (bool, int, map[string]string) {
	if len(p) == 0 { return true, idx, vars }
	if idx >= len(segs) { return false, idx, vars }
	h, t := p[0], p[1:]
	switch h.k {
	case _segLit:
		if segs[idx] != h.s { return false, idx, vars }
		return _match(segs, idx+1, t, vars)
	case _segVar:
		v := segs[idx]
		if v == "" { return false, idx, vars }
		if vars == nil { vars = make(map[string]string) }
		vars[h.s] = v
		return _match(segs, idx+1, t, vars)
	default:
		return false, idx, vars
	}
}

func _parseByVariants(name string, variants [][]_pattern) (map[string]string, int, error) {
	segs := strings.Split(name, "/")
	for i, ps := range variants {
		for _, p := range ps {
			if ok, consumed, vars := _match(segs, 0, p, nil); ok && consumed == len(segs) {
				return vars, i, nil
			}
		}
	}
	return nil, -1, errors.New("invalid resource name")
}
`))

// 各リソース本文（ヘッダは main.go で付与）
var fileBodyTmpl = template.Must(
	template.New("fileBody").
		Option("missingkey=error").
		Funcs(funcMap).
		Parse(`{{range .Targets}}
{{- $t := . -}}
// {{$t.TypeName}} patterns
type {{$t.VariantType}} int
const (
	{{- range $i, $v := .Variants }}
	{{$t.TypeName}}NameVariant{{$i}} {{$t.VariantType}} = {{$i}} // {{join $v.Segs "/"}}
	{{- end }}
)

var _{{$t.TypeName}}Variants = [][]_pattern{
	{{- range .Variants }}
	{ _pattern{ {{- range $i, $s := .Segs}}{{if $i}}, {{end}}{{seg $s}}{{- end}} } },
	{{- end }}
}

type {{$t.NameType}} struct{
	Variant {{$t.VariantType}}
	{{- range .Accessors}}
	{{.Field}} string   {{/* 非公開フィールド */}}
	{{- end}}
}

func Parse{{$t.TypeName}}Name(name string) ({{$t.NameType}}, error) {
	vars, which, err := _parseByVariants(name, _{{$t.TypeName}}Variants)
	if err != nil { return {{$t.NameType}}{}, err }
	out := {{$t.NameType}}{ Variant: {{$t.VariantType}}(which) }
	{{- range .Accessors}}
	out.{{.Field}} = vars["{{.Field}}"]
	{{- end}}
	return out, nil
}

{{range .FormatFuncs}}
func {{.Func}}({{fmtArgs .Segs}}) string {
	return {{fmtJoin .Segs}}
}
{{end}}

{{range .Accessors}}
func (n {{$t.NameType}}) Get{{.Field | title}}() string { return n.{{.Field}} }
{{end}}

{{end}}
`))
