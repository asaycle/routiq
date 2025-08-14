import * as grpcWeb from 'grpc-web';

import * as v1_user_pb from '../v1/user_pb'; // proto import: "v1/user.proto"


export class UserServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getUser(
    request: v1_user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_user_pb.User) => void
  ): grpcWeb.ClientReadableStream<v1_user_pb.User>;

  listUsers(
    request: v1_user_pb.ListUsersRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_user_pb.ListUsersResponse) => void
  ): grpcWeb.ClientReadableStream<v1_user_pb.ListUsersResponse>;

}

export class UserServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getUser(
    request: v1_user_pb.GetUserRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_user_pb.User>;

  listUsers(
    request: v1_user_pb.ListUsersRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_user_pb.ListUsersResponse>;

}

