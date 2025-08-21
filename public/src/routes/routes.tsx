import React from 'react';
import { Routes, Route } from 'react-router-dom';

import AnnouncementList from '../pages/Announcement/AnnouncementsList';
import LoginPage from '../components/page/LoginPage';
import OAuthCallbackPage from '../components/page/OAuthCallbackPage';
import TagList from '../components/tag/TagList';
import RouteCreate from '../components/route/RouteCreate';
import RouteList from '../components/route/RouteList';
import RouteShow from '../components/route/RouteShow';
import TouringCreate from '../components/touring/TouringCreate';
import HomePage from '../pages/HomePage';

const AppRoutes: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/announcements" element={<AnnouncementList />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/oauth/callback" element={<OAuthCallbackPage />} />
      <Route path="/tags" element={<TagList />} />
      <Route path="/routes" element={<RouteList />} />
      <Route path="/routes/new" element={<RouteCreate />} />
      <Route path="/routes/:id" element={<RouteShow />} />
      <Route path="/routes/:id/tourings/new" element={<TouringCreate />} />
    </Routes>
  );
}

export default AppRoutes;