import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { listRoutes } from '../../api/routeClient';
import { Box, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Button } from '@mui/material';
import { Route } from '../../../lib/proto/v1/route_pb';

const RouteList: React.FC = () => {
  const [routes, setRoutes] = useState<Route[]>([]);

  useEffect(() => {
    const fetchRoutes = async () => {
      try {
        setRoutes(await listRoutes());
      } catch (error) {
        console.error('Error fetching tags:', error);
      }
    };

    fetchRoutes();
  }, []);


  return (
    <div>
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
        <h1>Route List</h1>
        <Button
          variant="contained"
          color="primary"
          component={Link}
          to="/routes/new" // RouteCreate コンポーネントへのリンク
        >
          新しいルートを作成
        </Button>
      </Box>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {routes.map((route, index) => (
              <TableRow key={index}>
                <TableCell>
                  <Link to={`/routes/${route.getId()}`}>
                    {route.getDisplayName()}
                  </Link>
                  </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
};

export default RouteList;
