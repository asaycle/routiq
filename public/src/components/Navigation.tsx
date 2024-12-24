// src/components/Navigation.tsx
import React from 'react';
import { AppBar, Toolbar, Typography, Button, IconButton } from '@mui/material';
import { Link } from 'react-router-dom';
import { useRole } from '../contexts/RoleContext';

const Navigation: React.FC = () => {
    const { role, logout } = useRole();
    return (
  <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          Motify
        </Typography>
        <Button color="inherit" component={Link} to="/">
          Home
        </Button>
        <Button color="inherit" component={Link} to="/tags">
          Tags
        </Button>
        <Button color="inherit" component={Link} to="/routes">
          Routes
        </Button>
        {role === 'guest' ? (
          <Button color="inherit" component={Link} to="/login">
            Login
          </Button>
        ) : (
          <>
            <Button color="inherit" onClick={logout}>
              Logout
            </Button>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Navigation;
