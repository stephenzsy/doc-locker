import * as grpcWeb from 'grpc-web';

import * as file_pb from './file_pb';


export class FileServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  beginUploadFile(
    request: file_pb.BeginUploadFileRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: file_pb.BeginUploadFileResponse) => void
  ): grpcWeb.ClientReadableStream<file_pb.BeginUploadFileResponse>;

}

export class FileServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  beginUploadFile(
    request: file_pb.BeginUploadFileRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<file_pb.BeginUploadFileResponse>;

}

