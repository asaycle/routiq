import React, { useState } from 'react';
import { TextField, Button, Box, Typography, Paper } from '@mui/material';
import { createRoute } from '../../api/routeClient';
import ErrorBoundary from '../error_boundary';
import GeoJsonEditor from './editor/editor2';

const RouteCreate: React.FC = () => {
    const [displayName, setDisplayName] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [geometry, setGeometry] = useState<string | null>(null);
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
    const [geoJson, setGeoJson] = useState<string>("");

    const handleSave = async () => {
        if (!displayName.trim()) {
            alert('ルート名を入力してください');
            return;
        }

        if (!geometry) {
            alert('ルートを作成してください');
            return;
        }

        setIsSubmitting(true);

        try {
            console.log("geojson", geoJson)
            console.log("geometry", geometry)
            const newRoute = await createRoute(displayName, description, geometry);
            alert('ルートが作成されました');
            console.log('Created Route:', newRoute);
        } catch (error) {
            console.error('Error creating route:', error);
            alert('ルートの作成中にエラーが発生しました');
        } finally {
            setIsSubmitting(false);
        }
    };

    const handleGeoJsonChange = (newGeoJson: string) => {
        setGeoJson(newGeoJson);
        setGeometry(geoJson)
    };

    return (
        <Box
            sx={{
                maxWidth: '800px',
                margin: 'auto',
                padding: 2,
                display: 'flex',
                flexDirection: 'column',
                gap: 3,
            }}
        >
            <Typography variant="h4" component="h1" gutterBottom>
                新しいルートを作成
            </Typography>

            <Paper elevation={3} sx={{ padding: 3 }}>
                <Box component="form" display="flex" flexDirection="column" gap={2}>
                    <TextField
                        label="ルート名"
                        variant="outlined"
                        value={displayName}
                        onChange={(e) => setDisplayName(e.target.value)}
                        fullWidth
                        required
                    />
                    <TextField
                        label="説明"
                        variant="outlined"
                        multiline
                        rows={4}
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        fullWidth
                    />
                </Box>
                <Box component="section">
                    地図上でルートを作成
                    <GeoJsonEditor onGeoJsonChange={handleGeoJsonChange} />
                </Box>
            </Paper>

            <Button
                variant="contained"
                color="primary"
                onClick={handleSave}
                disabled={isSubmitting}
                sx={{ alignSelf: 'flex-end' }}
            >
                {isSubmitting ? '送信中...' : 'ルートを保存'}
            </Button>
        </Box>
    );
};

export default RouteCreate;
