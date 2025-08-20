import * as jspb from 'google-protobuf'

import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb'; // proto import: "google/protobuf/field_mask.proto"


export class Pref extends jspb.Message {
  getName(): string;
  setName(value: string): Pref;

  getDisplayName(): string;
  setDisplayName(value: string): Pref;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pref.AsObject;
  static toObject(includeInstance: boolean, msg: Pref): Pref.AsObject;
  static serializeBinaryToWriter(message: Pref, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pref;
  static deserializeBinaryFromReader(message: Pref, reader: jspb.BinaryReader): Pref;
}

export namespace Pref {
  export type AsObject = {
    name: string,
    displayName: string,
  }
}

export class City extends jspb.Message {
  getName(): string;
  setName(value: string): City;

  getDisplayName(): string;
  setDisplayName(value: string): City;

  getPref(): string;
  setPref(value: string): City;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): City.AsObject;
  static toObject(includeInstance: boolean, msg: City): City.AsObject;
  static serializeBinaryToWriter(message: City, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): City;
  static deserializeBinaryFromReader(message: City, reader: jspb.BinaryReader): City;
}

export namespace City {
  export type AsObject = {
    name: string,
    displayName: string,
    pref: string,
  }
}

export class Location extends jspb.Message {
  getName(): string;
  setName(value: string): Location;

  getUserLabel(): string;
  setUserLabel(value: string): Location;

  getPlaceId(): string;
  setPlaceId(value: string): Location;

  getLat(): number;
  setLat(value: number): Location;

  getLng(): number;
  setLng(value: number): Location;

  getPref(): string;
  setPref(value: string): Location;

  getCity(): string;
  setCity(value: string): Location;

  getUserLabelNorm(): string;
  setUserLabelNorm(value: string): Location;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Location.AsObject;
  static toObject(includeInstance: boolean, msg: Location): Location.AsObject;
  static serializeBinaryToWriter(message: Location, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Location;
  static deserializeBinaryFromReader(message: Location, reader: jspb.BinaryReader): Location;
}

export namespace Location {
  export type AsObject = {
    name: string,
    userLabel: string,
    placeId: string,
    lat: number,
    lng: number,
    pref: string,
    city: string,
    userLabelNorm: string,
  }
}

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

export class GetPrefRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetPrefRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPrefRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPrefRequest): GetPrefRequest.AsObject;
  static serializeBinaryToWriter(message: GetPrefRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPrefRequest;
  static deserializeBinaryFromReader(message: GetPrefRequest, reader: jspb.BinaryReader): GetPrefRequest;
}

export namespace GetPrefRequest {
  export type AsObject = {
    name: string,
  }
}

export class ListCitiesRequest extends jspb.Message {
  getFilter(): string;
  setFilter(value: string): ListCitiesRequest;

  getPageSize(): number;
  setPageSize(value: number): ListCitiesRequest;

  getPageToken(): string;
  setPageToken(value: string): ListCitiesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCitiesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListCitiesRequest): ListCitiesRequest.AsObject;
  static serializeBinaryToWriter(message: ListCitiesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCitiesRequest;
  static deserializeBinaryFromReader(message: ListCitiesRequest, reader: jspb.BinaryReader): ListCitiesRequest;
}

export namespace ListCitiesRequest {
  export type AsObject = {
    filter: string,
    pageSize: number,
    pageToken: string,
  }
}

export class ListCitiesResponse extends jspb.Message {
  getCitiesList(): Array<City>;
  setCitiesList(value: Array<City>): ListCitiesResponse;
  clearCitiesList(): ListCitiesResponse;
  addCities(value?: City, index?: number): City;

  getNextPageToken(): string;
  setNextPageToken(value: string): ListCitiesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCitiesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCitiesResponse): ListCitiesResponse.AsObject;
  static serializeBinaryToWriter(message: ListCitiesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCitiesResponse;
  static deserializeBinaryFromReader(message: ListCitiesResponse, reader: jspb.BinaryReader): ListCitiesResponse;
}

export namespace ListCitiesResponse {
  export type AsObject = {
    citiesList: Array<City.AsObject>,
    nextPageToken: string,
  }
}

export class GetCityRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetCityRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCityRequest): GetCityRequest.AsObject;
  static serializeBinaryToWriter(message: GetCityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCityRequest;
  static deserializeBinaryFromReader(message: GetCityRequest, reader: jspb.BinaryReader): GetCityRequest;
}

export namespace GetCityRequest {
  export type AsObject = {
    name: string,
  }
}

