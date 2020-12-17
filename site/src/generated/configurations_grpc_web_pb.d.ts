import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as configurations_pb from './configurations_pb';


export class ConfigurationsServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  siteConfigurations(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: configurations_pb.SiteConfigurationsResponse) => void
  ): grpcWeb.ClientReadableStream<configurations_pb.SiteConfigurationsResponse>;

  userProfile(
    request: configurations_pb.UserProfileRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: configurations_pb.UserProfileResponse) => void
  ): grpcWeb.ClientReadableStream<configurations_pb.UserProfileResponse>;

}

export class ConfigurationsServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  siteConfigurations(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata
  ): Promise<configurations_pb.SiteConfigurationsResponse>;

  userProfile(
    request: configurations_pb.UserProfileRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<configurations_pb.UserProfileResponse>;

}

