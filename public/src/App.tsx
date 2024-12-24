import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { AppBar, Toolbar, Typography, Button, Container } from '@mui/material';
import './App.css';
import 'leaflet/dist/leaflet.css';
import { RoleProvider } from './contexts/RoleContext';
import Header from './components/Header';
import Navigation from './components/Navigation';
import AppRoutes from './routes/routes'

function App() {
  return (
    <RoleProvider>
      <Router basename="/">
      <Navigation />
      <Container>
        <AppRoutes />
      </Container>
      </Router>
    </RoleProvider>
  );
}

export default App;
