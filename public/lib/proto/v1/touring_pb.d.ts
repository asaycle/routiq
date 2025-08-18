import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"
import * as google_type_date_pb from '../google/type/date_pb'; // proto import: "google/type/date.proto"
import * as v1_route_pb from '../v1/route_pb'; // proto import: "v1/route.proto"
import * as v1_tag_pb from '../v1/tag_pb'; // proto import: "v1/tag.proto"


export class Touring extends jspb.Message {
  getName(): string;
  setName(value: string): Touring;

  getRoute(): string;
  setRoute(value: string): Touring;

  getTagsList(): Array<string>;
  setTagsList(value: Array<string>): Touring;
  clearTagsList(): Touring;
  addTags(value: string, index?: number): Touring;

  getDate(): google_type_date_pb.Date | undefined;
  setDate(value?: google_type_date_pb.Date): Touring;
  hasDate(): boolean;
  clearDate(): Touring;

  getDisplayName(): string;
  setDisplayName(value: string): Touring;

  getNote(): string;
  setNote(value: string): Touring;

  getScore(): number;
  setScore(value: number): Touring;

  getUserId(): string;
  setUserId(value: string): Touring;

  getCreateTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreateTime(value?: google_protobuf_timestamp_pb.Timestamp): Touring;
  hasCreateTime(): boolean;
  clearCreateTime(): Touring;

  getUpdateTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdateTime(value?: google_protobuf_timestamp_pb.Timestamp): Touring;
  hasUpdateTime(): boolean;
  clearUpdateTime(): Touring;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Touring.AsObject;
  static toObject(includeInstance: boolean, msg: Touring): Touring.AsObject;
  static serializeBinaryToWriter(message: Touring, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Touring;
  static deserializeBinaryFromReader(message: Touring, reader: jspb.BinaryReader): Touring;
}

export namespace Touring {
  export type AsObject = {
    name: string,
    route: string,
    tagsList: Array<string>,
    date?: google_type_date_pb.Date.AsObject,
    displayName: string,
    note: string,
    score: number,
    userId: string,
    createTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updateTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class CreateTouringRequest extends jspb.Message {
  getParent(): string;
  setParent(value: string): CreateTouringRequest;

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
    parent: string,
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

  getView(): TouringView;
  setView(value: TouringView): ListTouringsRequest;

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
    view: TouringView,
  }
}

export class ListTouringsResponse extends jspb.Message {
  getTouringsList(): Array<Touring>;
  setTouringsList(value: Array<Touring>): ListTouringsResponse;
  clearTouringsList(): ListTouringsResponse;
  addTourings(value?: Touring, index?: number): Touring;

  getIncludedTagsMap(): jspb.Map<string, v1_tag_pb.Tag>;
  clearIncludedTagsMap(): ListTouringsResponse;

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
    includedTagsMap: Array<[string, v1_tag_pb.Tag.AsObject]>,
    nextPageToken: string,
  }
}

export class GetTouringRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetTouringRequest;

  getView(): TouringView;
  setView(value: TouringView): GetTouringRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTouringRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTouringRequest): GetTouringRequest.AsObject;
  static serializeBinaryToWriter(message: GetTouringRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTouringRequest;
  static deserializeBinaryFromReader(message: GetTouringRequest, reader: jspb.BinaryReader): GetTouringRequest;
}

export namespace GetTouringRequest {
  export type AsObject = {
    name: string,
    view: TouringView,
  }
}

export class GetTouringResponse extends jspb.Message {
  getTouring(): Touring | undefined;
  setTouring(value?: Touring): GetTouringResponse;
  hasTouring(): boolean;
  clearTouring(): GetTouringResponse;

  getIncludedTagsMap(): jspb.Map<string, v1_tag_pb.Tag>;
  clearIncludedTagsMap(): GetTouringResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTouringResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTouringResponse): GetTouringResponse.AsObject;
  static serializeBinaryToWriter(message: GetTouringResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTouringResponse;
  static deserializeBinaryFromReader(message: GetTouringResponse, reader: jspb.BinaryReader): GetTouringResponse;
}

export namespace GetTouringResponse {
  export type AsObject = {
    touring?: Touring.AsObject,
    includedTagsMap: Array<[string, v1_tag_pb.Tag.AsObject]>,
  }
}

export enum TouringView { 
  TOURING_VIEW_UNSPECIFIED = 0,
  TOURING_VIEW_BASIC = 1,
  TOURING_VIEW_WITH_TAGS = 2,
}
