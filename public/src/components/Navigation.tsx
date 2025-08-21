// src/components/Navigation.tsx
import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@mui/material';
import { Link } from 'react-router-dom';
import { useRole } from '../contexts/RoleContext';

const Navigation: React.FC = () => {
  const { role, logout } = useRole();
  return (
    <AppBar position="static">
      <Toolbar>
        <Typography
          variant="h6"
          component={Link}
          to="/"
          sx={{
            flexGrow: 1,
            textDecoration: 'none',
            color: 'inherit',
          }}
        >
          Routiq
        </Typography>
        <Button color="inherit" component={Link} to="/tags">
          タグから探す
        </Button>
        <Button color="inherit" component={Link} to="/routes">
          ルート一覧
        </Button>
        {role === 'guest' ? (
          <Button color="inherit" component={Link} to="/login">
            ログイン
          </Button>
        ) : (
          <Button color="inherit" onClick={logout}>
            ログアウト
          </Button>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Navigation;
