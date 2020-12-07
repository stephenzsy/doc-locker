import React from "react";
import { useAppContext } from "../AppContext";
import { HostServiceClient } from "../generated/HostServiceClientPb";
import { HostStatusRequest } from "../generated/host_pb";

export function Host(): JSX.Element {
  const [json, setJson] = React.useState<string | undefined>("Waiting");
  const { endpoint } = useAppContext();

  React.useEffect(() => {
    (async () => {
      const client = new HostServiceClient(endpoint);
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
