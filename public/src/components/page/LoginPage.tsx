import React from 'react';
import { Button, Container, Typography } from '@mui/material';
import { CONFIG } from '../../config';

const LoginPage: React.FC = () => {
    console.log(CONFIG)
    return (
        <Container maxWidth="sm" style={{ textAlign: 'center', marginTop: '100px' }}>
            <Typography variant="h4" gutterBottom>
                Motifyにログイン
            </Typography>
            <Typography variant="body1" gutterBottom>
                Googleアカウントで簡単にログインできます。
            </Typography>
            <LoginButton />
        </Container>
    );
};

const LoginButton: React.FC = () => {
    //const authUrl = `https://accounts.google.com/o/oauth2/auth?response_type=code&client_id=${CONFIG.GOOGLE_CLIENT_ID}&redirect_uri=${CONFIG.REDIRECT_URI}&scope=openid email profile`;
    const authUrl = `https://accounts.google.com/o/oauth2/auth?response_type=code&client_id=${CONFIG.GOOGLE_CLIENT_ID}&redirect_uri=${CONFIG.REDIRECT_URI}&scope=openid profile`;

    const handleLogin = () => {
        console.log(authUrl)
        window.location.href = authUrl;
    };

    return (
        <Button variant="contained" color="primary" onClick={handleLogin}>
            Googleでログイン
        </Button>
    );
};

export default LoginPage;
