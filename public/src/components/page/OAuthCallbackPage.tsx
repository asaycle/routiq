import React, { useEffect } from 'react';
import { useNavigate } from "react-router-dom";
import { exchangeOAuthCode } from '../../api/authClient';
import { useRole } from '../../contexts/RoleContext';

const OAuthCallbackPage: React.FC = () => {
    const navigate = useNavigate();
    const { setRole } = useRole();
    useEffect(() => {
        const exchange = async () => {
            const params = new URLSearchParams(window.location.search);
            const code = params.get("code");
            if (!code) {
                console.error("err")
                return
            }

            try {
                const resp = await exchangeOAuthCode(code);
                if (!resp) {
                    return
                }
                localStorage.setItem("access_token", resp.getAccessToken());
                localStorage.setItem("refresh_token", resp.getRefreshToken());
                setRole(resp.getRole())
                navigate(resp.getRedirectUrl()!)
            } catch (error) {
                console.error("failed exchange oauth code")
            }
        };
        exchange();
    }, []);

    return (
        <p>processing authentication...</p>
    );
};

export default OAuthCallbackPage;
