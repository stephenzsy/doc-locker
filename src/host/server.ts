import process from 'process';
import grpc, { sendUnaryData, ServerUnaryCall } from 'grpc';

import { HostServiceService, IHostServiceService } from '../generated/host_grpc_pb';
import { HostStatusRequest, HostStatusResponse } from '../generated/host_pb';

export class HostServiceServer implements IHostServiceService {
    private readonly server: grpc.Server;

    constructor(host: string, port: number) {
        this.server = new grpc.Server();

        // Add a descriptor with gRPC metadata that says what message types should be received
        // and returned at various URLS, followed by a map of handlers that match with RPC names
        // that implement the functionality at those endpoints. This class puts those handlers
        // into public instance fields below, so we can just pass `this` as the implementation.
        this.server.addService(HostServiceService, this);

        // Tell the server where it should listen for network connections.
        this.server.bind(`${host}:${port}`, grpc.ServerCredentials.createInsecure());
    }

    public status(call: ServerUnaryCall<HostStatusRequest>, callback: sendUnaryData<HostStatusResponse>): void {
        const response = new HostStatusResponse();
        const statusJson = {
            'node-version': process.version,
        };
        response.setStatusJson(JSON.stringify(statusJson));
        // First parameter is error, second is response message
        callback(null, response);
    }
}
