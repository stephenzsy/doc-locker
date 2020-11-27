import { sendUnaryData, ServerUnaryCall } from '@grpc/grpc-js';
import process from 'process';
import { IHostServiceService } from '../generated/host_grpc_pb';
import { HostStatusRequest, HostStatusResponse } from '../generated/host_pb';


export class HostServiceServer implements IHostServiceService {

    public status(
        call: ServerUnaryCall<HostStatusRequest, HostStatusResponse>,
        callback: sendUnaryData<HostStatusResponse>): void {
        const response = new HostStatusResponse();
        const statusJson = {
            'node-version': process.version,
        };
        response.setStatusjson(JSON.stringify(statusJson));
        // First parameter is error, second is response message
        callback(null, response);
    }
}
