import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"
import * as google_type_date_pb from '../google/type/date_pb'; // proto import: "google/type/date.proto"
import * as v1_route_pb from '../v1/route_pb'; // proto import: "v1/route.proto"
import * as v1_tag_pb from '../v1/tag_pb'; // proto import: "v1/tag.proto"


export class Touring extends jspb.Message {
  getId(): string;
  setId(value: string): Touring;

  getRouteId(): string;
  setRouteId(value: string): Touring;

  getTagsList(): Array<v1_tag_pb.Tag>;
  setTagsList(value: Array<v1_tag_pb.Tag>): Touring;
  clearTagsList(): Touring;
  addTags(value?: v1_tag_pb.Tag, index?: number): v1_tag_pb.Tag;

  getDate(): google_type_date_pb.Date | undefined;
  setDate(value?: google_type_date_pb.Date): Touring;
  hasDate(): boolean;
  clearDate(): Touring;

  getTitle(): string;
  setTitle(value: string): Touring;

  getNote(): string;
  setNote(value: string): Touring;

  getScore(): number;
  setScore(value: number): Touring;

  getUserId(): string;
  setUserId(value: string): Touring;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Touring;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Touring;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Touring.AsObject;
  static toObject(includeInstance: boolean, msg: Touring): Touring.AsObject;
  static serializeBinaryToWriter(message: Touring, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Touring;
  static deserializeBinaryFromReader(message: Touring, reader: jspb.BinaryReader): Touring;
}

export namespace Touring {
  export type AsObject = {
    id: string,
    routeId: string,
    tagsList: Array<v1_tag_pb.Tag.AsObject>,
    date?: google_type_date_pb.Date.AsObject,
    title: string,
    note: string,
    score: number,
    userId: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class CreateTouringRequest extends jspb.Message {
  getTouring(): Touring | undefined;
  setTouring(value?: Touring): CreateTouringRequest;
  hasTouring(): boolean;
  clearTouring(): CreateTouringRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTouringRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTouringRequest): CreateTouringRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTouringRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTouringRequest;
  static deserializeBinaryFromReader(message: CreateTouringRequest, reader: jspb.BinaryReader): CreateTouringRequest;
}

export namespace CreateTouringRequest {
  export type AsObject = {
    touring?: Touring.AsObject,
  }
}

export class ListTouringsRequest extends jspb.Message {
  getFilter(): string;
  setFilter(value: string): ListTouringsRequest;

  getPageSize(): number;
  setPageSize(value: number): ListTouringsRequest;

  getPageToken(): string;
  setPageToken(value: string): ListTouringsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTouringsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListTouringsRequest): ListTouringsRequest.AsObject;
  static serializeBinaryToWriter(message: ListTouringsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTouringsRequest;
  static deserializeBinaryFromReader(message: ListTouringsRequest, reader: jspb.BinaryReader): ListTouringsRequest;
}

export namespace ListTouringsRequest {
  export type AsObject = {
    filter: string,
    pageSize: number,
    pageToken: string,
  }
}

export class ListTouringsResponse extends jspb.Message {
  getTouringsList(): Array<Touring>;
  setTouringsList(value: Array<Touring>): ListTouringsResponse;
  clearTouringsList(): ListTouringsResponse;
  addTourings(value?: Touring, index?: number): Touring;

  getNextPageToken(): string;
  setNextPageToken(value: string): ListTouringsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTouringsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListTouringsResponse): ListTouringsResponse.AsObject;
  static serializeBinaryToWriter(message: ListTouringsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTouringsResponse;
  static deserializeBinaryFromReader(message: ListTouringsResponse, reader: jspb.BinaryReader): ListTouringsResponse;
}

export namespace ListTouringsResponse {
  export type AsObject = {
    touringsList: Array<Touring.AsObject>,
    nextPageToken: string,
  }
}

export class GetTouringRequest extends jspb.Message {
  getId(): string;
  setId(value: string): GetTouringRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTouringRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTouringRequest): GetTouringRequest.AsObject;
  static serializeBinaryToWriter(message: GetTouringRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTouringRequest;
  static deserializeBinaryFromReader(message: GetTouringRequest, reader: jspb.BinaryReader): GetTouringRequest;
}

export namespace GetTouringRequest {
  export type AsObject = {
    id: string,
  }
}

