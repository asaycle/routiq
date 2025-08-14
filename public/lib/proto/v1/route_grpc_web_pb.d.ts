import * as grpcWeb from 'grpc-web';

import * as v1_route_pb from '../v1/route_pb'; // proto import: "v1/route.proto"


export class RouteServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  createRoute(
    request: v1_route_pb.CreateRouteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_route_pb.Route) => void
  ): grpcWeb.ClientReadableStream<v1_route_pb.Route>;

  listRoutes(
    request: v1_route_pb.ListRoutesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_route_pb.ListRoutesResponse) => void
  ): grpcWeb.ClientReadableStream<v1_route_pb.ListRoutesResponse>;

  getRoute(
    request: v1_route_pb.GetRouteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_route_pb.Route) => void
  ): grpcWeb.ClientReadableStream<v1_route_pb.Route>;

}

export class RouteServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  createRoute(
    request: v1_route_pb.CreateRouteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_route_pb.Route>;

  listRoutes(
    request: v1_route_pb.ListRoutesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_route_pb.ListRoutesResponse>;

  getRoute(
    request: v1_route_pb.GetRouteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_route_pb.Route>;

}

