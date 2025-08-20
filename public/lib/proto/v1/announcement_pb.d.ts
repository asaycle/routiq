import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"


export class Announcement extends jspb.Message {
  getName(): string;
  setName(value: string): Announcement;

  getTitle(): string;
  setTitle(value: string): Announcement;

  getBody(): string;
  setBody(value: string): Announcement;

  getCreateTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreateTime(value?: google_protobuf_timestamp_pb.Timestamp): Announcement;
  hasCreateTime(): boolean;
  clearCreateTime(): Announcement;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Announcement.AsObject;
  static toObject(includeInstance: boolean, msg: Announcement): Announcement.AsObject;
  static serializeBinaryToWriter(message: Announcement, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Announcement;
  static deserializeBinaryFromReader(message: Announcement, reader: jspb.BinaryReader): Announcement;
}

export namespace Announcement {
  export type AsObject = {
    name: string,
    title: string,
    body: string,
    createTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class ListAnnouncementsRequest extends jspb.Message {
  getPageSize(): number;
  setPageSize(value: number): ListAnnouncementsRequest;

  getPageToken(): string;
  setPageToken(value: string): ListAnnouncementsRequest;

  getFilter(): string;
  setFilter(value: string): ListAnnouncementsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAnnouncementsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListAnnouncementsRequest): ListAnnouncementsRequest.AsObject;
  static serializeBinaryToWriter(message: ListAnnouncementsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAnnouncementsRequest;
  static deserializeBinaryFromReader(message: ListAnnouncementsRequest, reader: jspb.BinaryReader): ListAnnouncementsRequest;
}

export namespace ListAnnouncementsRequest {
  export type AsObject = {
    pageSize: number,
    pageToken: string,
    filter: string,
  }
}

export class ListAnnouncementsResponse extends jspb.Message {
  getAnnouncementsList(): Array<Announcement>;
  setAnnouncementsList(value: Array<Announcement>): ListAnnouncementsResponse;
  clearAnnouncementsList(): ListAnnouncementsResponse;
  addAnnouncements(value?: Announcement, index?: number): Announcement;

  getNextPageToken(): string;
  setNextPageToken(value: string): ListAnnouncementsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAnnouncementsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListAnnouncementsResponse): ListAnnouncementsResponse.AsObject;
  static serializeBinaryToWriter(message: ListAnnouncementsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAnnouncementsResponse;
  static deserializeBinaryFromReader(message: ListAnnouncementsResponse, reader: jspb.BinaryReader): ListAnnouncementsResponse;
}

export namespace ListAnnouncementsResponse {
  export type AsObject = {
    announcementsList: Array<Announcement.AsObject>,
    nextPageToken: string,
  }
}

export class GetAnnouncementRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetAnnouncementRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAnnouncementRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAnnouncementRequest): GetAnnouncementRequest.AsObject;
  static serializeBinaryToWriter(message: GetAnnouncementRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAnnouncementRequest;
  static deserializeBinaryFromReader(message: GetAnnouncementRequest, reader: jspb.BinaryReader): GetAnnouncementRequest;
}

export namespace GetAnnouncementRequest {
  export type AsObject = {
    name: string,
  }
}

