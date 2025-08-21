import { AuthServiceClient } from '../../lib/proto/v1/AuthServiceClientPb';
import {
  ExchangeOAuthCodeRequest,
  ExchangeOAuthCodeResponse,
  VerifyTokenRequest,
  VerifyTokenResponse,
} from '../../lib/proto/v1/auth_pb';

// gRPC-Webクライアントを初期化
const client = new AuthServiceClient(process.env.REACT_APP_API_BASE!);

export const exchangeOAuthCode = async (
  code: string
): Promise<ExchangeOAuthCodeResponse | null> => {
  const request = new ExchangeOAuthCodeRequest();
  request.setCode(code);

  return new Promise((resolve, reject) => {
    client.exchangeOAuthCode(request, {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        resolve(response || null);
      }
    });
  });
};

export const verifyToken = async (accessToken: string): Promise<{ role: string; userId: string; }> => {
  return new Promise((resolve, reject) => {
    const request = new VerifyTokenRequest();
    request.setAccessToken(accessToken);

    client.verifyToken(request, {}, (err, response) => {
      if (err) {
        reject(new Error(err.message));
        return;
      }

      const role = response?.getRole();
      const userId = response?.getUserId();

      if (role && userId) {
        resolve({ role, userId });
      } else {
        reject(new Error('Invalid response from server'));
      }
    });
  });
};