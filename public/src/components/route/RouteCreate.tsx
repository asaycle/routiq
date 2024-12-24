import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { TextField, Button, Box, Typography, Paper } from '@mui/material';
import { createRoute } from '../../api/routeClient';
import GeoJsonEditor from './editor/editor';

const RouteCreate: React.FC = () => {
    const navigate = useNavigate();
    const [displayName, setDisplayName] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
    const [geoJson, setGeoJson] = useState<string>("{}");

    const handleSave = async () => {
        if (!displayName.trim()) {
            alert('ルート名を入力してください');
            return;
        }

        if (!geoJson) {
            alert('ルートを作成してください');
            return;
        }

        setIsSubmitting(true);

        try {
            console.log("geojson", geoJson)
            const newRoute = await createRoute(displayName, description, geoJson);
            alert('ルートが作成されました');
            console.log('Created Route:', newRoute);
            navigate(`/routes/${newRoute?.getId()}`);
        } catch (error) {
            console.error('Error creating route:', error);
            alert('ルートの作成中にエラーが発生しました');
        } finally {
            setIsSubmitting(false);
        }
    };

    const handleGeoJsonChange = (newGeoJson: string) => {
        setGeoJson(newGeoJson);
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
                    <GeoJsonEditor geoJson={geoJson} onGeoJsonChange={handleGeoJsonChange} />
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
