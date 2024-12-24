import { grpc } from "@improbable-eng/grpc-web";
import { RouteServiceClient } from '../../lib/proto/v1/RouteServiceClientPb';
import { CreateRouteRequest, GetRouteRequest, ListRoutesRequest, Route } from '../../lib/proto/v1/route_pb';
import { createMetadata, handleGrpcError } from "./common";

// gRPC-Webクライアントを初期化
const client = new RouteServiceClient('http://localhost:8080', null, {
  withCredentials: true
});

/**
 * CreateRoute APIを呼び出す関数
 */
export const createRoute = async (
  displayName: string,
  description: string,
  geometry: string // GeoJSON形式の文字列
): Promise<Route | null> => {
  const route = new Route();
  route.setDisplayName(displayName);
  route.setDescription(description);
  route.setFeature(geometry);

  console.log(route)

  const request = new CreateRouteRequest();
  request.setRoute(route);

  return new Promise((resolve, reject) => {
    client.createRoute(request, createMetadata(), (err, response) => {
      if (err) {
        handleGrpcError(err)
        reject(err);
      } else {
        resolve(response || null);
      }
    });
  });
};

/**
 * GetRoute APIを呼び出す関数
 */
export const getRoute = async (id: string): Promise<Route | null> => {
  console.log("getRoute: ", id)
  const request = new GetRouteRequest();
  request.setId(id);
  console.log("getRoute 2: ", request.getId())

  return new Promise((resolve, reject) => {
    client.getRoute(request, createMetadata(), (err, response) => {
      if (err) {
        handleGrpcError(err)
        reject(err);
      } else {
        resolve(response || null);
      }
    });
  });
};

/**
 * ListTags APIを呼び出す関数
 */
export const listRoutes = async (): Promise<Route[]> => {
  const request = new ListRoutesRequest();
  return new Promise((resolve, reject) => {
    client.listRoutes(request, createMetadata(), (err, response) => {
      if (err) {
        handleGrpcError(err)
        reject(err);
      } else {
        resolve(response?.getRoutesList() || []);
      }
    });
  });
};
