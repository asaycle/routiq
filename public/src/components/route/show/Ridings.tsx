import React, { useEffect, useState } from 'react';
import {
    Avatar,
    Box,
    CircularProgress,
    Paper,
    Rating,
    Typography,
} from '@mui/material';
import { listRouteRidings } from '../../../api/ridingClient';
import { Route } from '../../../../lib/proto/v1/route_pb';
import { Riding } from '../../../../lib/proto/v1/riding_pb';
import { format } from 'date-fns';

type RouteRidingsProps = {
    route: Route
}

const RouteRidings: React.FC<RouteRidingsProps> = ({ route }) => {
    const [loading, setLoading] = useState<boolean>(true);
    const [ridings, setRidings] = useState<Riding[] | null>(null);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchRoute = async () => {
            try {
                const resp = await listRouteRidings(route.getId() || '');
                setRidings(resp);
                setLoading(false);
            } catch (err) {
                setError('Failed to load route data.');
                setLoading(false);
            }
        };

        fetchRoute();
    }, [route]);

    if (loading) {
        return <CircularProgress />;
    }

    if (error) {
        return <Typography color="error">{error}</Typography>;
    }

    if (!ridings) {
        return <Typography color="error">Route not found.</Typography>;
    }

    return (
        <Paper elevation={3} style={{ padding: '20px', marginBottom: '20px' }}>
            <Typography variant="h6" gutterBottom>
                ライディング記録
            </Typography>
            <Box mt={2}>
                {ridings.map((riding) => (
                    <Box
                        key={riding.getId()}
                        display="flex"
                        flexDirection="row"
                        borderBottom="1px solid #e0e0e0"
                        paddingY="10px"
                    >
                        {/* ユーザーアイコン */}
                        <Avatar alt={riding.getUserId()} style={{ marginRight: '10px' }}>
                            {riding.getUserId().charAt(0)} {/* ユーザーIDの頭文字 */}
                        </Avatar>

                        {/* ツーリング情報 */}
                        <Box flex={1}>
                            <Box display="flex" justifyContent="space-between">
                                <Typography variant="body1" fontWeight="bold">
                                    {riding.getTitle() || 'タイトルなし'}
                                </Typography>
                                <Typography variant="body2" color="textSecondary">
                                </Typography>
                            </Box>
                            <Rating value={riding.getScore()} readOnly size="small" />
                            <Typography variant="body2" mt={1}>
                                {riding.getNote() || 'コメントなし'}
                            </Typography>
                        </Box>
                    </Box>
                ))}
            </Box>
        </Paper>
    )
};

export default RouteRidings;