import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as file_pb from './file_pb';


export class SiteConfigurationsResponse extends jspb.Message {
  getSiteconfigurationsjson(): string;
  setSiteconfigurationsjson(value: string): SiteConfigurationsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SiteConfigurationsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SiteConfigurationsResponse): SiteConfigurationsResponse.AsObject;
  static serializeBinaryToWriter(message: SiteConfigurationsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SiteConfigurationsResponse;
  static deserializeBinaryFromReader(message: SiteConfigurationsResponse, reader: jspb.BinaryReader): SiteConfigurationsResponse;
}

export namespace SiteConfigurationsResponse {
  export type AsObject = {
    siteconfigurationsjson: string,
  }
}

export class UserProfileRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): UserProfileRequest;

  getConfigpath(): string;
  setConfigpath(value: string): UserProfileRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserProfileRequest): UserProfileRequest.AsObject;
  static serializeBinaryToWriter(message: UserProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserProfileRequest;
  static deserializeBinaryFromReader(message: UserProfileRequest, reader: jspb.BinaryReader): UserProfileRequest;
}

export namespace UserProfileRequest {
  export type AsObject = {
    provider: string,
    configpath: string,
  }
}

export class UserProfileResponse extends jspb.Message {
  getLastupdated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastupdated(value?: google_protobuf_timestamp_pb.Timestamp): UserProfileResponse;
  hasLastupdated(): boolean;
  clearLastupdated(): UserProfileResponse;

  getFiledestinationsList(): Array<file_pb.FileDestination>;
  setFiledestinationsList(value: Array<file_pb.FileDestination>): UserProfileResponse;
  clearFiledestinationsList(): UserProfileResponse;
  addFiledestinations(value?: file_pb.FileDestination, index?: number): file_pb.FileDestination;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserProfileResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserProfileResponse): UserProfileResponse.AsObject;
  static serializeBinaryToWriter(message: UserProfileResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserProfileResponse;
  static deserializeBinaryFromReader(message: UserProfileResponse, reader: jspb.BinaryReader): UserProfileResponse;
}

export namespace UserProfileResponse {
  export type AsObject = {
    lastupdated?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    filedestinationsList: Array<file_pb.FileDestination.AsObject>,
  }
}

