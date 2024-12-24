package maps

import (
	"fmt"
	"strings"

	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/xerrors"
)

func GenerateGoogleMapsURL(feature *geojson.Feature) (string, error) {
	if feature.Geometry == nil {
		return "", xerrors.New("geometry is nil")
	}
	geometry := feature.Geometry
	// GeoJSONがLineStringか確認
	if geometry.Type != "LineString" {
		return "", fmt.Errorf("unsupported geometry type: %s", feature.Type)
	}

	// GeoJSONの座標を取得
	coordinates := geometry.LineString
	if len(coordinates) < 2 {
		return "", fmt.Errorf("at least two coordinates are required")
	}

	// 出発地点と目的地を設定
	origin := fmt.Sprintf("%f,%f", coordinates[0][1], coordinates[0][0])
	destination := fmt.Sprintf("%f,%f", coordinates[len(coordinates)-1][1], coordinates[len(coordinates)-1][0])

	// 経由地点を設定
	var waypoints []string
	for _, coord := range coordinates[1 : len(coordinates)-1] {
		waypoints = append(waypoints, fmt.Sprintf("%f,%f", coord[1], coord[0]))
	}
	waypointsParam := strings.Join(waypoints, "|")

	// Google Maps URLを生成
	url := fmt.Sprintf("https://www.google.com/maps/dir/?api=1&origin=%s&destination=%s&travelmode=driving", origin, destination)
	if len(waypoints) > 0 {
		url += fmt.Sprintf("&waypoints=%s", waypointsParam)
	}

	return url, nil
}
