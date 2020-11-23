// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var host_pb = require('./host_pb.js');

function serialize_HostStatusRequest(arg) {
  if (!(arg instanceof host_pb.HostStatusRequest)) {
    throw new Error('Expected argument of type HostStatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_HostStatusRequest(buffer_arg) {
  return host_pb.HostStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_HostStatusResponse(arg) {
  if (!(arg instanceof host_pb.HostStatusResponse)) {
    throw new Error('Expected argument of type HostStatusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_HostStatusResponse(buffer_arg) {
  return host_pb.HostStatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var HostServiceService = exports.HostServiceService = {
  status: {
    path: '/HostService/Status',
    requestStream: false,
    responseStream: false,
    requestType: host_pb.HostStatusRequest,
    responseType: host_pb.HostStatusResponse,
    requestSerialize: serialize_HostStatusRequest,
    requestDeserialize: deserialize_HostStatusRequest,
    responseSerialize: serialize_HostStatusResponse,
    responseDeserialize: deserialize_HostStatusResponse,
  },
};

exports.HostServiceClient = grpc.makeGenericClientConstructor(HostServiceService);
