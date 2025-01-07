import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getRoute } from '../../api/routeClient';
import RouteMap from './show/Map';
import RouteRidings from './show/Ridings';
import { Route } from '../../../lib/proto/v1/route_pb';
import { Link as RouterLink } from 'react-router-dom';
import { Button, Chip, Grid, CircularProgress, Link, Typography, Paper, Box, Stack } from '@mui/material';
import Breadcrumbs from '@mui/material/Breadcrumbs';

const RouteShow: React.FC = () => {
  const { id } = useParams<{ id: string }>(); // React Routerのパラメータを取得
  console.log('Route ID:', id);
  const [route, setRoute] = useState<Route | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchRoute = async () => {
      try {
        const fetchedRoute = await getRoute(id || '');
        setRoute(fetchedRoute);
        setLoading(false);
      } catch (err) {
        setError('Failed to load route data.');
        setLoading(false);
      }
    };

    fetchRoute();
  }, [id]);

  if (loading) {
    return <CircularProgress />;
  }

  if (error) {
    return <Typography color="error">{error}</Typography>;
  }

  if (!route) {
    return <Typography color="error">Route not found.</Typography>;
  }

  console.log(JSON.parse(route.getFeature()))

  return (
    <Box style={{ padding: '20px' }}>
      {/* パンくずリスト */}
      <Box mb={2}>
        <Breadcrumbs aria-label="breadcrumb">
          <Link underline="hover" color="inherit" href="/">
            ホーム
          </Link>
          <Link underline="hover" color="inherit" href="/routes">
            ルート一覧
          </Link>
          <Typography color="textPrimary">{route.getDisplayName()}</Typography>
        </Breadcrumbs>
      </Box>

      {/* メインコンテンツ */}
      <Grid container spacing={2}>
        {/* 左側: 地図ギャラリー */}
        <Grid item xs={12} md={7}>
          <Paper elevation={3} style={{ padding: '20px' }}>
            <Typography variant="h4" gutterBottom>
              {route.getDisplayName()}
            </Typography>
            <Box mt={2}>
              <RouteMap route={route} />
            </Box>
            <Typography variant="h5" gutterBottom>
              ルートの詳細
            </Typography>
            <Stack direction="row" spacing={1} flexWrap="wrap">
              {route.getTagCountsList().map((tag) => (
                <Chip
                  key={tag.getTag()?.getId()}
                  label={tag.getTag()?.getName()}
                  color="primary"
                  variant="outlined"
                />
              ))}
            </Stack>
            <Typography variant="body1" color="textSecondary" gutterBottom>
              {route.getDescription()}
            </Typography>
          </Paper>
        </Grid>

        {/* 右側: 商品情報/ツーリング概要 */}
        <Grid item xs={12} md={5}>
          <Paper elevation={3} style={{ padding: '20px' }}>
            <Typography variant="body1" color="textSecondary" gutterBottom>
              {route.getDescription()}
            </Typography>
            <Box mt={2}>
              <Button
                variant="contained"
                component={RouterLink}
                to={`/routes/${route.getId()}/ridings/new`}
                color="primary"
                fullWidth
              >
                ツーリング記録を作成
              </Button>
            </Box>
            <Box mt={2}>
              <Button
                variant="contained"
                color="primary"
                href={route.getGoogleMapUrl()}
                target="_blank"
                rel="noopener noreferrer"
                fullWidth
              >
                地図を見る
              </Button>
            </Box>
            <Box mt={2}>
              <Button
                variant="outlined"
                color="secondary"
                onClick={() => alert('お気に入りに追加しました')}
                fullWidth
              >
                お気に入りに追加
              </Button>
            </Box>
            <Box mt={2}>
              <Button
                variant="outlined"
                color="secondary"
                onClick={() => alert('共有しました')}
                fullWidth
              >
                共有
              </Button>
            </Box>
          </Paper>
        </Grid>
      </Grid>

      {/* 下部セクション: ツーリング記録 */}
      <RouteRidings route={route} />
    </Box>
  );
};

export default RouteShow;
