import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Button, CircularProgress, Typography, Paper, Box } from '@mui/material';
import { MapContainer, GeoJSON, TileLayer, useMap } from 'react-leaflet';
import L from 'leaflet';
import { Route } from '../../../../lib/proto/v1/route_pb';

type RouteMapProps = {
    route: Route
}

const RouteMap: React.FC<RouteMapProps> = ({ route }) => {
    const parsedJSON = JSON.parse(route.getGeoJson());
    const bounds = L.geoJSON(parsedJSON).getBounds();
    const center = bounds.getCenter();
    return (
        <Box mt={4}>
            <MapContainer boxZoom={true} center={center} zoom={10} style={{ height: "50vh", width: "100%" }}>
                <CustomTileLayer />
                <GeoJSONLayer geoJson={parsedJSON} />
            </MapContainer>
        </Box>
    );
};

const GeoJSONLayer: React.FC<{ geoJson: any }> = ({ geoJson }) => {
    const map = useMap();
    const [isBoundsSet, setBoundsSet] = useState(false);

    useEffect(() => {
        if (geoJson && !isBoundsSet) {
            const geoJsonLayer = L.geoJSON(geoJson);
            const bounds = geoJsonLayer.getBounds();

            if (bounds.isValid()) {
                map.fitBounds(bounds, { padding: [50, 50], maxZoom: 15 });
                setBoundsSet(true);
            }
        }
    }, [geoJson, map, isBoundsSet]);

    return <GeoJSON data={geoJson} />;
};

const CustomTileLayer = () => (
    <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://osm.org/copyright">OpenStreetMap</a> contributors'
    />
);


export default RouteMap;