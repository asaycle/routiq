import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"
import * as v1_tag_pb from '../v1/tag_pb'; // proto import: "v1/tag.proto"


export class TagCount extends jspb.Message {
  getTag(): v1_tag_pb.Tag | undefined;
  setTag(value?: v1_tag_pb.Tag): TagCount;
  hasTag(): boolean;
  clearTag(): TagCount;

  getCount(): number;
  setCount(value: number): TagCount;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TagCount.AsObject;
  static toObject(includeInstance: boolean, msg: TagCount): TagCount.AsObject;
  static serializeBinaryToWriter(message: TagCount, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TagCount;
  static deserializeBinaryFromReader(message: TagCount, reader: jspb.BinaryReader): TagCount;
}

export namespace TagCount {
  export type AsObject = {
    tag?: v1_tag_pb.Tag.AsObject,
    count: number,
  }
}

export class Route extends jspb.Message {
  getName(): string;
  setName(value: string): Route;

  getDisplayName(): string;
  setDisplayName(value: string): Route;

  getDescription(): string;
  setDescription(value: string): Route;

  getDistance(): number;
  setDistance(value: number): Route;

  getGeoJson(): string;
  setGeoJson(value: string): Route;

  getGoogleMapUrl(): string;
  setGoogleMapUrl(value: string): Route;

  getTagCountsList(): Array<TagCount>;
  setTagCountsList(value: Array<TagCount>): Route;
  clearTagCountsList(): Route;
  addTagCounts(value?: TagCount, index?: number): TagCount;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Route;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Route;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Route.AsObject;
  static toObject(includeInstance: boolean, msg: Route): Route.AsObject;
  static serializeBinaryToWriter(message: Route, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Route;
  static deserializeBinaryFromReader(message: Route, reader: jspb.BinaryReader): Route;
}

export namespace Route {
  export type AsObject = {
    name: string,
    displayName: string,
    description: string,
    distance: number,
    geoJson: string,
    googleMapUrl: string,
    tagCountsList: Array<TagCount.AsObject>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class CreateRouteRequest extends jspb.Message {
  getRoute(): Route | undefined;
  setRoute(value?: Route): CreateRouteRequest;
  hasRoute(): boolean;
  clearRoute(): CreateRouteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRouteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRouteRequest): CreateRouteRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRouteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRouteRequest;
  static deserializeBinaryFromReader(message: CreateRouteRequest, reader: jspb.BinaryReader): CreateRouteRequest;
}

export namespace CreateRouteRequest {
  export type AsObject = {
    route?: Route.AsObject,
  }
}

export class ListRoutesRequest extends jspb.Message {
  getPageSize(): number;
  setPageSize(value: number): ListRoutesRequest;

  getPageToken(): string;
  setPageToken(value: string): ListRoutesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoutesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoutesRequest): ListRoutesRequest.AsObject;
  static serializeBinaryToWriter(message: ListRoutesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoutesRequest;
  static deserializeBinaryFromReader(message: ListRoutesRequest, reader: jspb.BinaryReader): ListRoutesRequest;
}

export namespace ListRoutesRequest {
  export type AsObject = {
    pageSize: number,
    pageToken: string,
  }
}

export class ListRoutesResponse extends jspb.Message {
  getRoutesList(): Array<Route>;
  setRoutesList(value: Array<Route>): ListRoutesResponse;
  clearRoutesList(): ListRoutesResponse;
  addRoutes(value?: Route, index?: number): Route;

  getNextPageToken(): string;
  setNextPageToken(value: string): ListRoutesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoutesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoutesResponse): ListRoutesResponse.AsObject;
  static serializeBinaryToWriter(message: ListRoutesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoutesResponse;
  static deserializeBinaryFromReader(message: ListRoutesResponse, reader: jspb.BinaryReader): ListRoutesResponse;
}

export namespace ListRoutesResponse {
  export type AsObject = {
    routesList: Array<Route.AsObject>,
    nextPageToken: string,
  }
}

export class GetRouteRequest extends jspb.Message {
  getId(): string;
  setId(value: string): GetRouteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRouteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRouteRequest): GetRouteRequest.AsObject;
  static serializeBinaryToWriter(message: GetRouteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRouteRequest;
  static deserializeBinaryFromReader(message: GetRouteRequest, reader: jspb.BinaryReader): GetRouteRequest;
}

export namespace GetRouteRequest {
  export type AsObject = {
    id: string,
  }
}

