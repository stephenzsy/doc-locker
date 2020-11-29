// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var host_pb = require('./host_pb.js');

function serialize_doclocker_host_HostStatusRequest(arg) {
  if (!(arg instanceof host_pb.HostStatusRequest)) {
    throw new Error('Expected argument of type doclocker.host.HostStatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_doclocker_host_HostStatusRequest(buffer_arg) {
  return host_pb.HostStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_doclocker_host_HostStatusResponse(arg) {
  if (!(arg instanceof host_pb.HostStatusResponse)) {
    throw new Error('Expected argument of type doclocker.host.HostStatusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_doclocker_host_HostStatusResponse(buffer_arg) {
  return host_pb.HostStatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var HostServiceService = exports.HostServiceService = {
  status: {
    path: '/doclocker.host.HostService/Status',
    requestStream: false,
    responseStream: false,
    requestType: host_pb.HostStatusRequest,
    responseType: host_pb.HostStatusResponse,
    requestSerialize: serialize_doclocker_host_HostStatusRequest,
    requestDeserialize: deserialize_doclocker_host_HostStatusRequest,
    responseSerialize: serialize_doclocker_host_HostStatusResponse,
    responseDeserialize: deserialize_doclocker_host_HostStatusResponse,
  },
};

exports.HostServiceClient = grpc.makeGenericClientConstructor(HostServiceService);
