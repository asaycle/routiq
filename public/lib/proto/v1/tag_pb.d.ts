import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"


export class Tag extends jspb.Message {
  getId(): string;
  setId(value: string): Tag;

  getName(): string;
  setName(value: string): Tag;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Tag.AsObject;
  static toObject(includeInstance: boolean, msg: Tag): Tag.AsObject;
  static serializeBinaryToWriter(message: Tag, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Tag;
  static deserializeBinaryFromReader(message: Tag, reader: jspb.BinaryReader): Tag;
}

export namespace Tag {
  export type AsObject = {
    id: string,
    name: string,
  }
}

export class CreateTagRequest extends jspb.Message {
  getTag(): Tag | undefined;
  setTag(value?: Tag): CreateTagRequest;
  hasTag(): boolean;
  clearTag(): CreateTagRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTagRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTagRequest): CreateTagRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTagRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTagRequest;
  static deserializeBinaryFromReader(message: CreateTagRequest, reader: jspb.BinaryReader): CreateTagRequest;
}

export namespace CreateTagRequest {
  export type AsObject = {
    tag?: Tag.AsObject,
  }
}

export class ListTagsRequest extends jspb.Message {
  getFilter(): string;
  setFilter(value: string): ListTagsRequest;

  getPageSize(): number;
  setPageSize(value: number): ListTagsRequest;

  getPageToken(): string;
  setPageToken(value: string): ListTagsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTagsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListTagsRequest): ListTagsRequest.AsObject;
  static serializeBinaryToWriter(message: ListTagsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTagsRequest;
  static deserializeBinaryFromReader(message: ListTagsRequest, reader: jspb.BinaryReader): ListTagsRequest;
}

export namespace ListTagsRequest {
  export type AsObject = {
    filter: string,
    pageSize: number,
    pageToken: string,
  }
}

export class ListTagsResponse extends jspb.Message {
  getTagsList(): Array<Tag>;
  setTagsList(value: Array<Tag>): ListTagsResponse;
  clearTagsList(): ListTagsResponse;
  addTags(value?: Tag, index?: number): Tag;

  getNextPageToken(): string;
  setNextPageToken(value: string): ListTagsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTagsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListTagsResponse): ListTagsResponse.AsObject;
  static serializeBinaryToWriter(message: ListTagsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTagsResponse;
  static deserializeBinaryFromReader(message: ListTagsResponse, reader: jspb.BinaryReader): ListTagsResponse;
}

export namespace ListTagsResponse {
  export type AsObject = {
    tagsList: Array<Tag.AsObject>,
    nextPageToken: string,
  }
}

