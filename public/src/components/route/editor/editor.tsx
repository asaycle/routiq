import React, { useRef, useEffect } from "react";
import { MapContainer, TileLayer } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import "leaflet-draw/dist/leaflet.draw.css";
import L, { Map as LeafletMap } from "leaflet";
import "leaflet-draw";

interface RouteEditorProps {
  onSave: (geoJson: GeoJSON.FeatureCollection) => void;
}

const RouteEditor: React.FC<RouteEditorProps> = ({ onSave }) => {
  const mapRef = useRef<LeafletMap | null>(null);
  const drawnItems = useRef(new L.FeatureGroup()).current;

  useEffect(() => {
    const map = mapRef.current;
    if (!map) return;

    // 描画用のレイヤーを追加
    map.addLayer(drawnItems);

    // Draw Control を初期化
    const drawControl = new L.Control.Draw({
      edit: {
        featureGroup: drawnItems,
      },
      draw: {
        polyline: {}, // Polyline のデフォルト設定
        polygon: false,
        circle: false,
        rectangle: false,
        marker: false,
      },
    });

    map.addControl(drawControl);

    // イベントハンドラ設定
    map.on(L.Draw.Event.CREATED, (e: any) => {
      const layer = e.layer;
      drawnItems.addLayer(layer);
    });

    return () => {
      // イベントのクリーンアップ
      map.off();
    };
  }, [onSave, drawnItems]);

  const handleSave = () => {
    const geoJson = drawnItems.toGeoJSON() as GeoJSON.FeatureCollection;
    onSave(geoJson);
  };

  return (
    <div>
      <MapContainer
        center={[35.6895, 139.6917]}
        zoom={13}
        style={{ height: "400px", width: "100%" }}
        ref={(mapInstance) => {
          if (mapInstance && !mapRef.current) {
            mapRef.current = mapInstance;
          }
        }}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        />
      </MapContainer>
      <button onClick={handleSave}>ルートを保存</button>
    </div>
  );
};

export default RouteEditor;
