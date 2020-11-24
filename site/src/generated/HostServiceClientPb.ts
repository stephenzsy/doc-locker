/**
 * @fileoverview gRPC-Web generated client stub for doclocker.host
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as host_pb from './host_pb';


export class HostServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoStatus = new grpcWeb.AbstractClientBase.MethodInfo(
    host_pb.HostStatusResponse,
    (request: host_pb.HostStatusRequest) => {
      return request.serializeBinary();
    },
    host_pb.HostStatusResponse.deserializeBinary
  );

  status(
    request: host_pb.HostStatusRequest,
    metadata: grpcWeb.Metadata | null): Promise<host_pb.HostStatusResponse>;

  status(
    request: host_pb.HostStatusRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: host_pb.HostStatusResponse) => void): grpcWeb.ClientReadableStream<host_pb.HostStatusResponse>;

  status(
    request: host_pb.HostStatusRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: host_pb.HostStatusResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/doclocker.host.HostService/Status',
        request,
        metadata || {},
        this.methodInfoStatus,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/doclocker.host.HostService/Status',
    request,
    metadata || {},
    this.methodInfoStatus);
  }

}

