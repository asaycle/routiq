import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"


export class ExchangeOAuthCodeRequest extends jspb.Message {
  getCode(): string;
  setCode(value: string): ExchangeOAuthCodeRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExchangeOAuthCodeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ExchangeOAuthCodeRequest): ExchangeOAuthCodeRequest.AsObject;
  static serializeBinaryToWriter(message: ExchangeOAuthCodeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExchangeOAuthCodeRequest;
  static deserializeBinaryFromReader(message: ExchangeOAuthCodeRequest, reader: jspb.BinaryReader): ExchangeOAuthCodeRequest;
}

export namespace ExchangeOAuthCodeRequest {
  export type AsObject = {
    code: string,
  }
}

export class ExchangeOAuthCodeResponse extends jspb.Message {
  getRedirectUrl(): string;
  setRedirectUrl(value: string): ExchangeOAuthCodeResponse;

  getAccessToken(): string;
  setAccessToken(value: string): ExchangeOAuthCodeResponse;

  getRefreshToken(): string;
  setRefreshToken(value: string): ExchangeOAuthCodeResponse;

  getRole(): string;
  setRole(value: string): ExchangeOAuthCodeResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExchangeOAuthCodeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ExchangeOAuthCodeResponse): ExchangeOAuthCodeResponse.AsObject;
  static serializeBinaryToWriter(message: ExchangeOAuthCodeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExchangeOAuthCodeResponse;
  static deserializeBinaryFromReader(message: ExchangeOAuthCodeResponse, reader: jspb.BinaryReader): ExchangeOAuthCodeResponse;
}

export namespace ExchangeOAuthCodeResponse {
  export type AsObject = {
    redirectUrl: string,
    accessToken: string,
    refreshToken: string,
    role: string,
  }
}

export class RefreshTokenRequest extends jspb.Message {
  getRefreshToken(): string;
  setRefreshToken(value: string): RefreshTokenRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshTokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshTokenRequest): RefreshTokenRequest.AsObject;
  static serializeBinaryToWriter(message: RefreshTokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshTokenRequest;
  static deserializeBinaryFromReader(message: RefreshTokenRequest, reader: jspb.BinaryReader): RefreshTokenRequest;
}

export namespace RefreshTokenRequest {
  export type AsObject = {
    refreshToken: string,
  }
}

export class RefreshTokenResponse extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): RefreshTokenResponse;

  getRefreshToken(): string;
  setRefreshToken(value: string): RefreshTokenResponse;

  getExpiresIn(): number;
  setExpiresIn(value: number): RefreshTokenResponse;

  getTokenType(): string;
  setTokenType(value: string): RefreshTokenResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshTokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshTokenResponse): RefreshTokenResponse.AsObject;
  static serializeBinaryToWriter(message: RefreshTokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshTokenResponse;
  static deserializeBinaryFromReader(message: RefreshTokenResponse, reader: jspb.BinaryReader): RefreshTokenResponse;
}

export namespace RefreshTokenResponse {
  export type AsObject = {
    accessToken: string,
    refreshToken: string,
    expiresIn: number,
    tokenType: string,
  }
}

export class VerifyTokenRequest extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): VerifyTokenRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VerifyTokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: VerifyTokenRequest): VerifyTokenRequest.AsObject;
  static serializeBinaryToWriter(message: VerifyTokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VerifyTokenRequest;
  static deserializeBinaryFromReader(message: VerifyTokenRequest, reader: jspb.BinaryReader): VerifyTokenRequest;
}

export namespace VerifyTokenRequest {
  export type AsObject = {
    accessToken: string,
  }
}

export class VerifyTokenResponse extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): VerifyTokenResponse;

  getRole(): string;
  setRole(value: string): VerifyTokenResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VerifyTokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: VerifyTokenResponse): VerifyTokenResponse.AsObject;
  static serializeBinaryToWriter(message: VerifyTokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VerifyTokenResponse;
  static deserializeBinaryFromReader(message: VerifyTokenResponse, reader: jspb.BinaryReader): VerifyTokenResponse;
}

export namespace VerifyTokenResponse {
  export type AsObject = {
    userId: string,
    role: string,
  }
}

