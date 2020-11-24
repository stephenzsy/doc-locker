import { handleUnaryCall, MethodDefinition, ServiceDefinition } from "grpc";

import { MethodDefinition, ServiceDefinition } from 'grpc';
import { HostStatusRequest, HostStatusResponse } from './host_pb';

export interface IHostServiceService {
    status: handleUnaryCall<HostStatusRequest, HostStatusResponse>;
}

export declare const HostServiceService: ServiceDefinition<IHostServiceService>;
