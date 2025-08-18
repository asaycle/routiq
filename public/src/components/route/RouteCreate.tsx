import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { TextField, Button, Box, Typography, Paper } from '@mui/material';
import { createRoute } from '../../api/routeClient';
import GeoJsonEditor from './editor/editor';
import { useForm, Controller } from "react-hook-form";

const RouteCreate: React.FC = () => {
    const { control, handleSubmit, setError, formState: { errors } } = useForm();
    const navigate = useNavigate();
    const [displayName, setDisplayName] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
    const [geoJson, setGeoJson] = useState<string>("{}");

    const onSubmit = async (data: any) => {
        try {
            console.log(data, geoJson)
            const route = await createRoute(data.displayName, data.description, geoJson);
            console.log(route?.getName())
            navigate(`${route?.getName()}`);
        } catch (error: any) {
        }
    };

    const handleGeoJsonChange = (newGeoJson: string) => {
        setGeoJson(newGeoJson);
    };

    return (
        <Box
            component="form"
            onSubmit={handleSubmit(onSubmit)}
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

            <Controller
                name="displayName"
                control={control}
                defaultValue=""
                rules={{ required: "ルート名は必須です。" }}
                render={({ field }) => (
                    <TextField
                        {...field}
                        label="ルート名"
                        fullWidth
                        margin="normal"
                        error={!!errors.display_name}
                        helperText={errors.display_name ? (errors.display_name.message as string) : ""}
                    />
                )}
            />

            <Controller
                name="description"
                control={control}
                defaultValue=""
                rules={{ required: "説明は必須です。" }}
                render={({ field }) => (
                    <TextField
                        {...field}
                        label="説明"
                        fullWidth
                        margin="normal"
                        error={!!errors.description}
                        helperText={errors.description ? (errors.description.message as string) : ""}
                    />
                )}
            />

            <Box component="section">
                地図上でルートを作成
                <GeoJsonEditor geoJson={geoJson} onGeoJsonChange={handleGeoJsonChange} />
            </Box>

            <Button
                variant="contained"
                color="primary"
                type="submit"
                sx={{ alignSelf: 'flex-end' }}
            >
                {isSubmitting ? '送信中...' : 'ルートを保存'}
            </Button>
        </Box>
    );
};

export default RouteCreate;
