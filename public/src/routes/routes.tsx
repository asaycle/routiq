import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { AppBar, Toolbar, Typography, Button, Container } from '@mui/material';

import LoginPage from '../components/page/LoginPage';
import OAuthCallbackPage from '../components/page/OAuthCallbackPage';
import TagList from '../components/tag/TagList';
import RouteCreate from '../components/route/RouteCreate';
import RouteList from '../components/route/RouteList';
import RouteShow from '../components/route/RouteShow';
import RidingCreate from '../components/riding/RidingCreate';

const AppRoutes: React.FC = () => {
    return (
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/oauth/callback" element={<OAuthCallbackPage />} />
          <Route path="/tags" element={<TagList />} />
          <Route path="/routes" element={<RouteList />} />
          <Route path="/routes/new" element={<RouteCreate />} />
          <Route path="/routes/:id" element={<RouteShow />} />
          <Route path="/routes/:id/ridings/new" element={<RidingCreate />} />
        </Routes>
    );
}

export default AppRoutes;