import * as grpcWeb from 'grpc-web';

import * as v1_touring_pb from '../v1/touring_pb'; // proto import: "v1/touring.proto"


export class TouringServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  createTouring(
    request: v1_touring_pb.CreateTouringRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_touring_pb.Touring) => void
  ): grpcWeb.ClientReadableStream<v1_touring_pb.Touring>;

  listTourings(
    request: v1_touring_pb.ListTouringsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_touring_pb.ListTouringsResponse) => void
  ): grpcWeb.ClientReadableStream<v1_touring_pb.ListTouringsResponse>;

  getTouring(
    request: v1_touring_pb.GetTouringRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_touring_pb.Touring) => void
  ): grpcWeb.ClientReadableStream<v1_touring_pb.Touring>;

}

export class TouringServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  createTouring(
    request: v1_touring_pb.CreateTouringRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_touring_pb.Touring>;

  listTourings(
    request: v1_touring_pb.ListTouringsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_touring_pb.ListTouringsResponse>;

  getTouring(
    request: v1_touring_pb.GetTouringRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_touring_pb.Touring>;

}

