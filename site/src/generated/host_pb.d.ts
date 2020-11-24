import * as jspb from 'google-protobuf'



export class HostStatusRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HostStatusRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HostStatusRequest): HostStatusRequest.AsObject;
  static serializeBinaryToWriter(message: HostStatusRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HostStatusRequest;
  static deserializeBinaryFromReader(message: HostStatusRequest, reader: jspb.BinaryReader): HostStatusRequest;
}

export namespace HostStatusRequest {
  export type AsObject = {
  }
}

export class HostStatusResponse extends jspb.Message {
  getStatusjson(): string;
  setStatusjson(value: string): HostStatusResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HostStatusResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HostStatusResponse): HostStatusResponse.AsObject;
  static serializeBinaryToWriter(message: HostStatusResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HostStatusResponse;
  static deserializeBinaryFromReader(message: HostStatusResponse, reader: jspb.BinaryReader): HostStatusResponse;
}

export namespace HostStatusResponse {
  export type AsObject = {
    statusjson: string,
  }
}

