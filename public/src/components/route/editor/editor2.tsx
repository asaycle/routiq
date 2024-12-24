import React, { useEffect, useRef, useState } from 'react';
import { MapContainer, TileLayer, useMap } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet-draw/dist/leaflet.draw.css';
import L, { Map as LeafletMap } from 'leaflet';
import { Box, Tab } from '@mui/material';
import { TabContext, TabList, TabPanel } from '@mui/lab';
import MonacoEditor from 'react-monaco-editor';
import * as monaco from 'monaco-editor/esm/vs/editor/editor.api';
import 'leaflet-draw';
import { Feature, LineString } from 'geojson';
import 'leaflet/dist/leaflet.css'; // Leaflet の基本スタイル
import 'leaflet-draw/dist/leaflet.draw.css'; // Leaflet.Draw のスタイル

interface GeoJsonEditorProps {
    onGeoJsonChange: (geoJson: string) => void;
}

const GeoJsonEditor: React.FC<GeoJsonEditorProps> = ({ onGeoJsonChange }) => {
    const mapRef = useRef<L.Map | null>(null);
    const drawnItems = useRef(new L.FeatureGroup()).current;
    const [geoJson, setGeoJson] = useState<string>(() => {
        try {
            return JSON.stringify(
                {
                    type: 'Feature',
                    geometry: {
                        type: 'LineString',
                        coordinates: [
                            [135.9, 35.0],
                            [135.8, 35.0],
                        ],
                    },
                    properties: {}
                },
                null,
                2
            );
        } catch (error) {
            console.error('Invalid initial GeoJSON data:', error);
            return '{}';
        }
    });
    const [activeTab, setActiveTab] = useState("1"); // タブの状態管理

    const MapWithDraw = () => {
        const map = useMap();

        useEffect(() => {
            onGeoJsonChange(geoJson);

            mapRef.current = map;

            const drawControl = new L.Control.Draw({
                edit: {
                    featureGroup: drawnItems,
                },
                draw: {
                    polyline: {
                        shapeOptions: {
                            color: 'blue',
                            weight: 5,
                        },
                    },
                },
            });

            map.addLayer(drawnItems);
            map.addControl(drawControl);

            map.on(L.Draw.Event.CREATED, (e: any) => {
                const layer = e.layer;
                drawnItems.addLayer(layer);

                const updatedGeoJson = drawnItems.toGeoJSON();
                setGeoJson(JSON.stringify(updatedGeoJson, null, 2));
                updateEditor()
            });

            map.on(L.Draw.Event.EDITED, () => {
                updateEditor()
            });

            map.on(L.Draw.Event.DELETED, () => {
                updateEditor()
            });

            const updateEditor = () => {
                const featureCollection = drawnItems.toGeoJSON() as GeoJSON.FeatureCollection;
                const features = featureCollection.features.map((feature) => ({
                    type: "Feature",
                    geometry: feature.geometry,
                    properties: feature.properties || {},
                }));
                setGeoJson(JSON.stringify(features[0], null, 2));
            };

        }, [map]);

        return null;
    };

    const handleEditorChange = (value: string) => {
        setGeoJson(value);

        try {
            const parsedGeoJson = JSON.parse(value);
            drawnItems.clearLayers();
            L.geoJSON(parsedGeoJson).addTo(drawnItems);

            const bounds = drawnItems.getBounds();
            if (bounds.isValid()) {
                mapRef.current?.fitBounds(bounds);
            }
        } catch (error) {
            console.error('Invalid GeoJSON:', error);
        }
    };

    const handleTabChange = (event: React.SyntheticEvent, newValue: string) => {
        setActiveTab(newValue);
    };


    useEffect(() => {
        console.log("aaaaaa")
        const map = mapRef.current;
        if (!map) {
            console.log("nil")
            return;
        }
        map.addLayer(drawnItems);
        const drawControl = new L.Control.Draw({
            edit: {
                featureGroup: drawnItems,
            },
            draw: {
                polyline: {
                    shapeOptions: {
                        color: 'blue', // ラインの色
                        weight: 5,     // ラインの太さ
                        opacity: 0.8,  // ラインの透明度
                    },
                },
                polygon: false,
                rectangle: false,
                circle: false,
                marker: false,
            },
        });
        map.addControl(drawControl);
        // 描画イベントのリスナーを追加
        map.on(L.Draw.Event.CREATED, (e: any) => {
            const layer = e.layer;
            drawnItems.addLayer(layer);
            console.log('Drawn layer:', layer.toGeoJSON());
            setGeoJson(JSON.stringify(layer.toGeoJSON(), null, 2));
        });

        map.on(L.Draw.Event.EDITED, (e: any) => {
            const layers = e.layers;
            layers.eachLayer((layer: any) => {
                console.log('Edited layer:', layer.toGeoJSON());
            });
        });

        map.on(L.Draw.Event.DELETED, (e: any) => {
            const layers = e.layers;
            layers.eachLayer((layer: any) => {
                console.log('Deleted layer:', layer.toGeoJSON());
            });
        });
    }, [mapRef, drawnItems]);

    return (
        <Box sx={{ width: '100%', typography: 'body1' }}>
            <TabContext value={activeTab}>
                <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                    <TabList onChange={handleTabChange} centered>
                        <Tab label="Map" value="1" />
                        <Tab label="Editor" value="2" />
                    </TabList>
                    <TabPanel value="1">
                        <div>
                            <MapContainer
                                center={[35, 135.9]}
                                zoom={8}
                                style={{ height: '500px', width: '100%' }}
                            >
                                <TileLayer
                                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                    attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
                                />
                                <MapWithDraw />
                            </MapContainer >
                        </div>
                    </TabPanel>
                    <TabPanel value="2">
                        <div style={{ flex: 1, height: '500px' }}>
                            <MonacoEditor
                                language="json"
                                value={geoJson}
                                theme="vs-light"
                                options={{
                                    automaticLayout: true,
                                    scrollBeyondLastLine: false,
                                    minimap: { enabled: false },
                                    formatOnPaste: true,
                                    formatOnType: true,
                                }}
                                onChange={handleEditorChange}
                            />
                        </div >
                    </TabPanel>
                </Box>
            </TabContext>
        </Box >
    );
};

export default GeoJsonEditor;
