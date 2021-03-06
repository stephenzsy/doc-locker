/**
 * @fileoverview gRPC-Web generated client stub for doclocker.configurations
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var file_pb = require('./file_pb.js')
const proto = {};
proto.doclocker = {};
proto.doclocker.configurations = require('./configurations_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.doclocker.configurations.ConfigurationsServiceClient =
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
proto.doclocker.configurations.ConfigurationsServicePromiseClient =
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
 *   !proto.google.protobuf.Empty,
 *   !proto.doclocker.configurations.SiteConfigurationsResponse>}
 */
const methodDescriptor_ConfigurationsService_SiteConfigurations = new grpc.web.MethodDescriptor(
  '/doclocker.configurations.ConfigurationsService/SiteConfigurations',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.doclocker.configurations.SiteConfigurationsResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.doclocker.configurations.SiteConfigurationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.doclocker.configurations.SiteConfigurationsResponse>}
 */
const methodInfo_ConfigurationsService_SiteConfigurations = new grpc.web.AbstractClientBase.MethodInfo(
  proto.doclocker.configurations.SiteConfigurationsResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.doclocker.configurations.SiteConfigurationsResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.doclocker.configurations.SiteConfigurationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.doclocker.configurations.SiteConfigurationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.doclocker.configurations.ConfigurationsServiceClient.prototype.siteConfigurations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/doclocker.configurations.ConfigurationsService/SiteConfigurations',
      request,
      metadata || {},
      methodDescriptor_ConfigurationsService_SiteConfigurations,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.doclocker.configurations.SiteConfigurationsResponse>}
 *     Promise that resolves to the response
 */
proto.doclocker.configurations.ConfigurationsServicePromiseClient.prototype.siteConfigurations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/doclocker.configurations.ConfigurationsService/SiteConfigurations',
      request,
      metadata || {},
      methodDescriptor_ConfigurationsService_SiteConfigurations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.doclocker.configurations.UserProfileRequest,
 *   !proto.doclocker.configurations.UserProfileResponse>}
 */
const methodDescriptor_ConfigurationsService_UserProfile = new grpc.web.MethodDescriptor(
  '/doclocker.configurations.ConfigurationsService/UserProfile',
  grpc.web.MethodType.UNARY,
  proto.doclocker.configurations.UserProfileRequest,
  proto.doclocker.configurations.UserProfileResponse,
  /**
   * @param {!proto.doclocker.configurations.UserProfileRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.doclocker.configurations.UserProfileResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.doclocker.configurations.UserProfileRequest,
 *   !proto.doclocker.configurations.UserProfileResponse>}
 */
const methodInfo_ConfigurationsService_UserProfile = new grpc.web.AbstractClientBase.MethodInfo(
  proto.doclocker.configurations.UserProfileResponse,
  /**
   * @param {!proto.doclocker.configurations.UserProfileRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.doclocker.configurations.UserProfileResponse.deserializeBinary
);


/**
 * @param {!proto.doclocker.configurations.UserProfileRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.doclocker.configurations.UserProfileResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.doclocker.configurations.UserProfileResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.doclocker.configurations.ConfigurationsServiceClient.prototype.userProfile =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/doclocker.configurations.ConfigurationsService/UserProfile',
      request,
      metadata || {},
      methodDescriptor_ConfigurationsService_UserProfile,
      callback);
};


/**
 * @param {!proto.doclocker.configurations.UserProfileRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.doclocker.configurations.UserProfileResponse>}
 *     Promise that resolves to the response
 */
proto.doclocker.configurations.ConfigurationsServicePromiseClient.prototype.userProfile =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/doclocker.configurations.ConfigurationsService/UserProfile',
      request,
      metadata || {},
      methodDescriptor_ConfigurationsService_UserProfile);
};


module.exports = proto.doclocker.configurations;

