import * as jspb from 'google-protobuf'



export class FileMetadata extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FileMetadata.AsObject;
  static toObject(includeInstance: boolean, msg: FileMetadata): FileMetadata.AsObject;
  static serializeBinaryToWriter(message: FileMetadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FileMetadata;
  static deserializeBinaryFromReader(message: FileMetadata, reader: jspb.BinaryReader): FileMetadata;
}

export namespace FileMetadata {
  export type AsObject = {
  }
}

export class AwsS3FileDestination extends jspb.Message {
  getName(): string;
  setName(value: string): AwsS3FileDestination;

  getRegion(): string;
  setRegion(value: string): AwsS3FileDestination;

  getBucketname(): string;
  setBucketname(value: string): AwsS3FileDestination;

  getPrefix(): string;
  setPrefix(value: string): AwsS3FileDestination;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AwsS3FileDestination.AsObject;
  static toObject(includeInstance: boolean, msg: AwsS3FileDestination): AwsS3FileDestination.AsObject;
  static serializeBinaryToWriter(message: AwsS3FileDestination, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AwsS3FileDestination;
  static deserializeBinaryFromReader(message: AwsS3FileDestination, reader: jspb.BinaryReader): AwsS3FileDestination;
}

export namespace AwsS3FileDestination {
  export type AsObject = {
    name: string,
    region: string,
    bucketname: string,
    prefix: string,
  }
}

export class FileDestination extends jspb.Message {
  getAws(): AwsS3FileDestination | undefined;
  setAws(value?: AwsS3FileDestination): FileDestination;
  hasAws(): boolean;
  clearAws(): FileDestination;

  getDestinationCase(): FileDestination.DestinationCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FileDestination.AsObject;
  static toObject(includeInstance: boolean, msg: FileDestination): FileDestination.AsObject;
  static serializeBinaryToWriter(message: FileDestination, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FileDestination;
  static deserializeBinaryFromReader(message: FileDestination, reader: jspb.BinaryReader): FileDestination;
}

export namespace FileDestination {
  export type AsObject = {
    aws?: AwsS3FileDestination.AsObject,
  }

  export enum DestinationCase { 
    DESTINATION_NOT_SET = 0,
    AWS = 1,
  }
}

export class UploadFileRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UploadFileRequest;

  getChunkoffset(): number;
  setChunkoffset(value: number): UploadFileRequest;

  getChunk(): Uint8Array | string;
  getChunk_asU8(): Uint8Array;
  getChunk_asB64(): string;
  setChunk(value: Uint8Array | string): UploadFileRequest;

  getMetadata(): FileMetadata | undefined;
  setMetadata(value?: FileMetadata): UploadFileRequest;
  hasMetadata(): boolean;
  clearMetadata(): UploadFileRequest;

  getDataCase(): UploadFileRequest.DataCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UploadFileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UploadFileRequest): UploadFileRequest.AsObject;
  static serializeBinaryToWriter(message: UploadFileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UploadFileRequest;
  static deserializeBinaryFromReader(message: UploadFileRequest, reader: jspb.BinaryReader): UploadFileRequest;
}

export namespace UploadFileRequest {
  export type AsObject = {
    id: string,
    chunkoffset: number,
    chunk: Uint8Array | string,
    metadata?: FileMetadata.AsObject,
  }

  export enum DataCase { 
    DATA_NOT_SET = 0,
    CHUNK = 3,
    METADATA = 4,
  }
}

export class UploadFileResponse extends jspb.Message {
  getId(): string;
  setId(value: string): UploadFileResponse;

  getDestinationname(): string;
  setDestinationname(value: string): UploadFileResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UploadFileResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UploadFileResponse): UploadFileResponse.AsObject;
  static serializeBinaryToWriter(message: UploadFileResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UploadFileResponse;
  static deserializeBinaryFromReader(message: UploadFileResponse, reader: jspb.BinaryReader): UploadFileResponse;
}

export namespace UploadFileResponse {
  export type AsObject = {
    id: string,
    destinationname: string,
  }
}

export class BeginUploadFileRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BeginUploadFileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BeginUploadFileRequest): BeginUploadFileRequest.AsObject;
  static serializeBinaryToWriter(message: BeginUploadFileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BeginUploadFileRequest;
  static deserializeBinaryFromReader(message: BeginUploadFileRequest, reader: jspb.BinaryReader): BeginUploadFileRequest;
}

export namespace BeginUploadFileRequest {
  export type AsObject = {
  }
}

export class BeginUploadFileResponse extends jspb.Message {
  getId(): string;
  setId(value: string): BeginUploadFileResponse;

  getDestinationsList(): Array<FileDestination>;
  setDestinationsList(value: Array<FileDestination>): BeginUploadFileResponse;
  clearDestinationsList(): BeginUploadFileResponse;
  addDestinations(value?: FileDestination, index?: number): FileDestination;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BeginUploadFileResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BeginUploadFileResponse): BeginUploadFileResponse.AsObject;
  static serializeBinaryToWriter(message: BeginUploadFileResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BeginUploadFileResponse;
  static deserializeBinaryFromReader(message: BeginUploadFileResponse, reader: jspb.BinaryReader): BeginUploadFileResponse;
}

export namespace BeginUploadFileResponse {
  export type AsObject = {
    id: string,
    destinationsList: Array<FileDestination.AsObject>,
  }
}

