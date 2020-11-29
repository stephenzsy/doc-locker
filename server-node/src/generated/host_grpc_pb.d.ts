import { handleUnaryCall, ServiceDefinition } from "@grpc/grpc-js";
import { HostStatusRequest, HostStatusResponse } from './host_pb';

export interface IHostServiceService {
    status: handleUnaryCall<HostStatusRequest, HostStatusResponse>;
}

export declare const HostServiceService: ServiceDefinition<IHostServiceService>;
