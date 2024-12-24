import * as jspb from 'google-protobuf'



export class ListPrefsRequest extends jspb.Message {
  getPageSize(): number;
  setPageSize(value: number): ListPrefsRequest;

  getPageToken(): string;
  setPageToken(value: string): ListPrefsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPrefsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPrefsRequest): ListPrefsRequest.AsObject;
  static serializeBinaryToWriter(message: ListPrefsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPrefsRequest;
  static deserializeBinaryFromReader(message: ListPrefsRequest, reader: jspb.BinaryReader): ListPrefsRequest;
}

export namespace ListPrefsRequest {
  export type AsObject = {
    pageSize: number,
    pageToken: string,
  }
}

export class ListPrefsResponse extends jspb.Message {
  getNextPageToken(): string;
  setNextPageToken(value: string): ListPrefsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPrefsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPrefsResponse): ListPrefsResponse.AsObject;
  static serializeBinaryToWriter(message: ListPrefsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPrefsResponse;
  static deserializeBinaryFromReader(message: ListPrefsResponse, reader: jspb.BinaryReader): ListPrefsResponse;
}

export namespace ListPrefsResponse {
  export type AsObject = {
    nextPageToken: string,
  }
}

