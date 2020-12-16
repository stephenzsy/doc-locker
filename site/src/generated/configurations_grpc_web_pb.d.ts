import * as grpcWeb from 'grpc-web';

import * as configurations_pb from './configurations_pb';


export class ConfigurationsServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  siteConfigurations(
    request: configurations_pb.SiteConfigurationsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: configurations_pb.SiteConfigurationsResponse) => void
  ): grpcWeb.ClientReadableStream<configurations_pb.SiteConfigurationsResponse>;

}

export class ConfigurationsServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  siteConfigurations(
    request: configurations_pb.SiteConfigurationsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<configurations_pb.SiteConfigurationsResponse>;

}

