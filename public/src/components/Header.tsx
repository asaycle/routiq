// src/components/Header.tsx
import React from 'react';
import { useRole } from '../contexts/RoleContext';
import { roles } from '../constants/roles';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import AccountCircle from '@mui/icons-material/AccountCircle';

const Header: React.FC = () => {
  const { role, loginAsUser, loginAsAdmin, logout } = useRole();

  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" sx={{ flexGrow: 1 }}>
          My Application
        </Typography>
        {role === roles.GUEST && (
          <>
            <Button color="inherit" onClick={loginAsUser}>
              ログイン (ユーザー)
            </Button>
            <Button color="inherit" onClick={loginAsAdmin}>
              ログイン (管理者)
            </Button>
          </>
        )}
        {role === roles.USER && (
          <>
            <IconButton color="inherit">
              <AccountCircle />
            </IconButton>
            <Button color="inherit" onClick={logout}>
              ログアウト
            </Button>
          </>
        )}
        {role === roles.ADMIN && (
          <>
            <IconButton color="inherit">
              <AccountCircle />
            </IconButton>
            <Button color="inherit" onClick={logout}>
              ログアウト
            </Button>
            <Button color="inherit">管理画面</Button>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Header;
