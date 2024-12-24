import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"
import * as google_type_date_pb from '../google/type/date_pb'; // proto import: "google/type/date.proto"
import * as v1_route_pb from '../v1/route_pb'; // proto import: "v1/route.proto"
import * as v1_tag_pb from '../v1/tag_pb'; // proto import: "v1/tag.proto"


export class Riding extends jspb.Message {
  getId(): string;
  setId(value: string): Riding;

  getRouteId(): string;
  setRouteId(value: string): Riding;

  getTagsList(): Array<v1_tag_pb.Tag>;
  setTagsList(value: Array<v1_tag_pb.Tag>): Riding;
  clearTagsList(): Riding;
  addTags(value?: v1_tag_pb.Tag, index?: number): v1_tag_pb.Tag;

  getDate(): google_type_date_pb.Date | undefined;
  setDate(value?: google_type_date_pb.Date): Riding;
  hasDate(): boolean;
  clearDate(): Riding;

  getTitle(): string;
  setTitle(value: string): Riding;

  getNote(): string;
  setNote(value: string): Riding;

  getScore(): number;
  setScore(value: number): Riding;

  getUserId(): string;
  setUserId(value: string): Riding;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Riding;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Riding;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Riding.AsObject;
  static toObject(includeInstance: boolean, msg: Riding): Riding.AsObject;
  static serializeBinaryToWriter(message: Riding, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Riding;
  static deserializeBinaryFromReader(message: Riding, reader: jspb.BinaryReader): Riding;
}

export namespace Riding {
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

export class CreateRidingRequest extends jspb.Message {
  getRiding(): Riding | undefined;
  setRiding(value?: Riding): CreateRidingRequest;
  hasRiding(): boolean;
  clearRiding(): CreateRidingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRidingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRidingRequest): CreateRidingRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRidingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRidingRequest;
  static deserializeBinaryFromReader(message: CreateRidingRequest, reader: jspb.BinaryReader): CreateRidingRequest;
}

export namespace CreateRidingRequest {
  export type AsObject = {
    riding?: Riding.AsObject,
  }
}

export class ListRidingsRequest extends jspb.Message {
  getFilter(): string;
  setFilter(value: string): ListRidingsRequest;

  getPageSize(): number;
  setPageSize(value: number): ListRidingsRequest;

  getPageToken(): string;
  setPageToken(value: string): ListRidingsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRidingsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRidingsRequest): ListRidingsRequest.AsObject;
  static serializeBinaryToWriter(message: ListRidingsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRidingsRequest;
  static deserializeBinaryFromReader(message: ListRidingsRequest, reader: jspb.BinaryReader): ListRidingsRequest;
}

export namespace ListRidingsRequest {
  export type AsObject = {
    filter: string,
    pageSize: number,
    pageToken: string,
  }
}

export class ListRidingsResponse extends jspb.Message {
  getRidingsList(): Array<Riding>;
  setRidingsList(value: Array<Riding>): ListRidingsResponse;
  clearRidingsList(): ListRidingsResponse;
  addRidings(value?: Riding, index?: number): Riding;

  getNextPageToken(): string;
  setNextPageToken(value: string): ListRidingsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRidingsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRidingsResponse): ListRidingsResponse.AsObject;
  static serializeBinaryToWriter(message: ListRidingsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRidingsResponse;
  static deserializeBinaryFromReader(message: ListRidingsResponse, reader: jspb.BinaryReader): ListRidingsResponse;
}

export namespace ListRidingsResponse {
  export type AsObject = {
    ridingsList: Array<Riding.AsObject>,
    nextPageToken: string,
  }
}

export class GetRidingRequest extends jspb.Message {
  getId(): string;
  setId(value: string): GetRidingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRidingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRidingRequest): GetRidingRequest.AsObject;
  static serializeBinaryToWriter(message: GetRidingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRidingRequest;
  static deserializeBinaryFromReader(message: GetRidingRequest, reader: jspb.BinaryReader): GetRidingRequest;
}

export namespace GetRidingRequest {
  export type AsObject = {
    id: string,
  }
}

