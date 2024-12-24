import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Button, CircularProgress, Typography, Paper, Box } from '@mui/material';
import { MapContainer, GeoJSON, useMap } from 'react-leaflet';
import L from 'leaflet';
import { Route } from '../../../../lib/proto/v1/route_pb';

type RouteMapProps = {
    route: Route
}

const RouteMap: React.FC<RouteMapProps> = ({ route }) => {

    return (
        <Box mt={4}>
            <MapContainer center={[35, 136]} zoom={10} style={{ height: "50vh", width: "100%" }}>
                {/* TileLayer for map tiles */}
                <CustomTileLayer />
                {/* GeoJSON Layer */}
                <GeoJSON data={JSON.parse(route.getGeoJson())} />
            </MapContainer>
        </Box>
    );
};

// カスタムTileLayerコンポーネント
const CustomTileLayer = () => {
    const map = useMap();

    React.useEffect(() => {
        // タイルレイヤーを追加
        const tileLayer = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        });

        tileLayer.addTo(map); // マップに追加

        return () => {
            map.removeLayer(tileLayer); // クリーンアップ
        };
    }, [map]);

    return null; // JSXを返さない
};


export default RouteMap;