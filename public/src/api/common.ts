import * as grpcWeb from "grpc-web";

/**
 * Metadata（HTTPヘッダー）を生成する関数
 * @returns grpcWeb.Metadata - gRPCリクエスト用のメタデータ
 */
export const createMetadata = (): grpcWeb.Metadata => {
  const token = localStorage.getItem("access_token");
  const metadata: grpcWeb.Metadata = {};

  if (token) {
    metadata.Authorization = `Bearer ${token}`;
  }

  return metadata;
};

export const handleGrpcError = (err: grpcWeb.RpcError): void => {
  if (err.code === grpcWeb.StatusCode.UNAUTHENTICATED) {
    // Unauthenticatedエラーの場合、ログインページにリダイレクト
    console.warn("Unauthenticated: Redirecting to login page.");
    window.location.href = "/login";
  } else {
    console.error("gRPC Error:", err.message);
  }
};