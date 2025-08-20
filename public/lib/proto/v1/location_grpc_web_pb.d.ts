import * as grpcWeb from 'grpc-web';

import * as v1_location_pb from '../v1/location_pb'; // proto import: "v1/location.proto"


export class LocationServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listPrefs(
    request: v1_location_pb.ListPrefsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_location_pb.ListPrefsResponse) => void
  ): grpcWeb.ClientReadableStream<v1_location_pb.ListPrefsResponse>;

  getPref(
    request: v1_location_pb.GetPrefRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_location_pb.Pref) => void
  ): grpcWeb.ClientReadableStream<v1_location_pb.Pref>;

  listCities(
    request: v1_location_pb.ListCitiesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_location_pb.ListCitiesResponse) => void
  ): grpcWeb.ClientReadableStream<v1_location_pb.ListCitiesResponse>;

  getCity(
    request: v1_location_pb.GetCityRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_location_pb.City) => void
  ): grpcWeb.ClientReadableStream<v1_location_pb.City>;

}

export class LocationServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listPrefs(
    request: v1_location_pb.ListPrefsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_location_pb.ListPrefsResponse>;

  getPref(
    request: v1_location_pb.GetPrefRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_location_pb.Pref>;

  listCities(
    request: v1_location_pb.ListCitiesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_location_pb.ListCitiesResponse>;

  getCity(
    request: v1_location_pb.GetCityRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_location_pb.City>;

}

