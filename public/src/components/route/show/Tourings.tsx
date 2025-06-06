import React, { useEffect, useState } from 'react';
import {
    Avatar,
    Box,
    CircularProgress,
    Paper,
    Rating,
    Typography,
} from '@mui/material';
import { listRouteTourings } from '../../../api/touringClient';
import { Route } from '../../../../lib/proto/v1/route_pb';
import { Touring } from '../../../../lib/proto/v1/touring_pb';
import { format } from 'date-fns';

type RouteTouringsProps = {
    route: Route
}

const RouteTourings: React.FC<RouteTouringsProps> = ({ route }) => {
    const [loading, setLoading] = useState<boolean>(true);
    const [tourings, setTourings] = useState<Touring[] | null>(null);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchRoute = async () => {
            try {
                const resp = await listRouteTourings(route.getId() || '');
                setTourings(resp);
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

    if (!tourings) {
        return <Typography color="error">Route not found.</Typography>;
    }

    return (
        <Paper elevation={3} style={{ padding: '20px', marginBottom: '20px' }}>
            <Typography variant="h6" gutterBottom>
                ライディング記録
            </Typography>
            <Box mt={2}>
                {tourings.map((touring) => (
                    <Box
                        key={touring.getId()}
                        display="flex"
                        flexDirection="row"
                        borderBottom="1px solid #e0e0e0"
                        paddingY="10px"
                    >
                        {/* ユーザーアイコン */}
                        <Avatar alt={touring.getUserId()} style={{ marginRight: '10px' }}>
                            {touring.getUserId().charAt(0)} {/* ユーザーIDの頭文字 */}
                        </Avatar>

                        {/* ツーリング情報 */}
                        <Box flex={1}>
                            <Box display="flex" justifyContent="space-between">
                                <Typography variant="body1" fontWeight="bold">
                                    {touring.getTitle() || 'タイトルなし'}
                                </Typography>
                                <Typography variant="body2" color="textSecondary">
                                </Typography>
                            </Box>
                            <Rating value={touring.getScore()} readOnly size="small" />
                            <Typography variant="body2" mt={1}>
                                {touring.getNote() || 'コメントなし'}
                            </Typography>
                        </Box>
                    </Box>
                ))}
            </Box>
        </Paper>
    )
};

export default RouteTourings;