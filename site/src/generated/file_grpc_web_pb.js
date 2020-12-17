/**
 * @fileoverview gRPC-Web generated client stub for doclocker.file
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.doclocker = {};
proto.doclocker.file = require('./file_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.doclocker.file.FileServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.doclocker.file.FileServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.doclocker.file.BeginUploadFileRequest,
 *   !proto.doclocker.file.BeginUploadFileResponse>}
 */
const methodDescriptor_FileService_BeginUploadFile = new grpc.web.MethodDescriptor(
  '/doclocker.file.FileService/BeginUploadFile',
  grpc.web.MethodType.UNARY,
  proto.doclocker.file.BeginUploadFileRequest,
  proto.doclocker.file.BeginUploadFileResponse,
  /**
   * @param {!proto.doclocker.file.BeginUploadFileRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.doclocker.file.BeginUploadFileResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.doclocker.file.BeginUploadFileRequest,
 *   !proto.doclocker.file.BeginUploadFileResponse>}
 */
const methodInfo_FileService_BeginUploadFile = new grpc.web.AbstractClientBase.MethodInfo(
  proto.doclocker.file.BeginUploadFileResponse,
  /**
   * @param {!proto.doclocker.file.BeginUploadFileRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.doclocker.file.BeginUploadFileResponse.deserializeBinary
);


/**
 * @param {!proto.doclocker.file.BeginUploadFileRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.doclocker.file.BeginUploadFileResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.doclocker.file.BeginUploadFileResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.doclocker.file.FileServiceClient.prototype.beginUploadFile =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/doclocker.file.FileService/BeginUploadFile',
      request,
      metadata || {},
      methodDescriptor_FileService_BeginUploadFile,
      callback);
};


/**
 * @param {!proto.doclocker.file.BeginUploadFileRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.doclocker.file.BeginUploadFileResponse>}
 *     Promise that resolves to the response
 */
proto.doclocker.file.FileServicePromiseClient.prototype.beginUploadFile =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/doclocker.file.FileService/BeginUploadFile',
      request,
      metadata || {},
      methodDescriptor_FileService_BeginUploadFile);
};


module.exports = proto.doclocker.file;

