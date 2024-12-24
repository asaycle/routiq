import React, { useEffect, useRef } from 'react';
import { MapContainer, TileLayer, useMap } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet-draw/dist/leaflet.draw.css';
import L from 'leaflet';
import 'leaflet-draw';
import { Box, Tab } from '@mui/material';
import { TabContext, TabList, TabPanel } from '@mui/lab';
import MonacoEditor from 'react-monaco-editor';

interface GeoJsonEditorProps {
    geoJson: string;
    onGeoJsonChange: (geoJson: string) => void;
}

const GeoJsonEditor: React.FC<GeoJsonEditorProps> = ({ geoJson, onGeoJsonChange }) => {
    const [activeTab, setActiveTab] = React.useState<string>("1");
    const handleTabChange = (event: React.SyntheticEvent, newValue: string) => {
        setActiveTab(newValue);
    };

    return (
        <Box sx={{ width: '100%', typography: 'body1' }}>
            <TabContext value={activeTab}>
                <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                    <TabList onChange={handleTabChange} centered>
                        <Tab label="Map" value="1" />
                        <Tab label="Editor" value="2" />
                    </TabList>
                    <TabPanel value="1">
                        <MapContainer
                            center={[35, 135.9]}
                            zoom={8}
                            style={{ height: '500px', width: '100%' }}
                        >
                            <TileLayer
                                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
                            />
                            <DrawControl geoJson={geoJson} onGeoJsonChange={onGeoJsonChange}/>
                        </MapContainer >
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
                                onChange={onGeoJsonChange}
                            />
                        </div >
                    </TabPanel>
                </Box>
            </TabContext>
        </Box >
    );
};

interface DrawControlProps {
    geoJson: string;
    onGeoJsonChange: (geoJson: string) => void;
}

const DrawControl: React.FC<DrawControlProps> = ({ geoJson, onGeoJsonChange }) => {
    const map = useMap();
    const geoJsonLayerRef = useRef<L.GeoJSON | null>(null);

    useEffect(() => {
        if (geoJsonLayerRef.current) {
            map.removeLayer(geoJsonLayerRef.current);
        }
        try {
            const parsedGeoJson = JSON.parse(geoJson);
            geoJsonLayerRef.current = L.geoJSON(parsedGeoJson).addTo(map);
        } catch (error) {
            console.error('Invalid GeoJSON:', error);
        }
        const featureGroup = new L.FeatureGroup().addTo(map);
        const drawControl = new L.Control.Draw({
            draw: {
                polygon: false,
                rectangle: false,
                circle: false,
                marker: false,
                circlemarker: false,
            },
            edit: {
                featureGroup
            },
        });
        map.addControl(drawControl);

        map.on(L.Draw.Event.CREATED, (e: L.LeafletEvent) => {
            const event= e as L.DrawEvents.Created;
            featureGroup.addLayer(event.layer)
            onGeoJsonChange(JSON.stringify(featureGroup.toGeoJSON(), null, 2));
        });

        return () => {
            map.removeControl(drawControl);
        };
    }, [map]);

    return null;
};

export default GeoJsonEditor;
