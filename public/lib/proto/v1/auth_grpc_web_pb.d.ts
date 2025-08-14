import * as grpcWeb from 'grpc-web';

import * as v1_auth_pb from '../v1/auth_pb'; // proto import: "v1/auth.proto"


export class AuthServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  exchangeOAuthCode(
    request: v1_auth_pb.ExchangeOAuthCodeRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_auth_pb.ExchangeOAuthCodeResponse) => void
  ): grpcWeb.ClientReadableStream<v1_auth_pb.ExchangeOAuthCodeResponse>;

  refreshToken(
    request: v1_auth_pb.RefreshTokenRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_auth_pb.RefreshTokenResponse) => void
  ): grpcWeb.ClientReadableStream<v1_auth_pb.RefreshTokenResponse>;

  verifyToken(
    request: v1_auth_pb.VerifyTokenRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_auth_pb.VerifyTokenResponse) => void
  ): grpcWeb.ClientReadableStream<v1_auth_pb.VerifyTokenResponse>;

}

export class AuthServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  exchangeOAuthCode(
    request: v1_auth_pb.ExchangeOAuthCodeRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_auth_pb.ExchangeOAuthCodeResponse>;

  refreshToken(
    request: v1_auth_pb.RefreshTokenRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_auth_pb.RefreshTokenResponse>;

  verifyToken(
    request: v1_auth_pb.VerifyTokenRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_auth_pb.VerifyTokenResponse>;

}

