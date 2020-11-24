import React from 'react';
import { HostServiceClient } from '../generated/HostServiceClientPb';
import { HostStatusRequest } from '../generated/host_pb';

export function Host(): JSX.Element {
    const [json, setJson] = React.useState<string | undefined>('Waiting');

    React.useEffect(() => {
        (async () => {
            const client = new HostServiceClient('localhost:10000');
            const response = await client.status(new HostStatusRequest(), {});
            setJson(response.getStatusjson());
        })();
    }, []);

    return (
        <div>
            <h1>Host Info</h1>
            {json}
        </div>
    );
}
