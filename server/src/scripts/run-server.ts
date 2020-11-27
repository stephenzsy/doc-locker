import { Server, ServerCredentials, UntypedServiceImplementation } from '@grpc/grpc-js';
import { HostServiceService } from '../generated/host_grpc_pb';
import { HostServiceServer } from '../host/server';
import { Configurations } from '../lib/common/configuration';

function main() {
    const serverConfig = Configurations.getServerSetupConfiguration().serverListenerConfiguration;
    const server = new Server();
    server.addService(HostServiceService, new HostServiceServer() as unknown as UntypedServiceImplementation);
    const serviceAddress: string = `${serverConfig.address}:${serverConfig.port}`;
    server.bindAsync(
        serviceAddress,
        ServerCredentials.createInsecure(), () => {
            server.start();
            console.log("started server at: " + serviceAddress)
        });
}

main();
