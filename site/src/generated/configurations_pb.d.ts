import * as jspb from 'google-protobuf'



export class SiteConfigurationsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SiteConfigurationsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SiteConfigurationsRequest): SiteConfigurationsRequest.AsObject;
  static serializeBinaryToWriter(message: SiteConfigurationsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SiteConfigurationsRequest;
  static deserializeBinaryFromReader(message: SiteConfigurationsRequest, reader: jspb.BinaryReader): SiteConfigurationsRequest;
}

export namespace SiteConfigurationsRequest {
  export type AsObject = {
  }
}

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

