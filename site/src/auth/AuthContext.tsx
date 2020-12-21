import Amplify from "@aws-amplify/core";
import { Empty } from "google-protobuf/google/protobuf/empty_pb";
import React from "react";
import { useAppContext } from "../AppContext";
import { ConfigurationsServicePromiseClient } from "../generated/configurations_grpc_web_pb";
import awsConfig from "../aws-exports";

export interface IAwsSiteConfigurations {
  cognitoRegion: string;
  cognitoUserPoolId: string;
  cognitoUserPoolWebClientId: string;
  cognitoUserPoolAttributesMapping: {
    configPath: string;
    cognitoIdentityPoolId: string;
    cognitoIdentityPoolRegion: string;
  };
}

export interface ISiteConfigurations {
  aws: IAwsSiteConfigurations;
}

export interface IAuthContext {
  aws: IAwsSiteConfigurations | undefined;
}

export const AuthContext = React.createContext<IAuthContext>({
  aws: undefined,
});

export function AuthContextProvider(props: {
  children?: React.ReactNode;
}): JSX.Element {
  const { endpoint } = useAppContext();
  const [siteConfigs, setSiteConfigs] = React.useState<
    ISiteConfigurations | undefined
  >(undefined);
  React.useEffect(() => {
    if (!endpoint) {
      return;
    }
    const client = new ConfigurationsServicePromiseClient(endpoint);
    (async () => {
      const response = await client.siteConfigurations(new Empty());
      const configs = JSON.parse(
        response.getSiteconfigurationsjson()
      ) as ISiteConfigurations;
      console.log(configs);
      Amplify.configure(awsConfig);
      setSiteConfigs(configs);
    })();
  }, [endpoint]);

  return (
    <AuthContext.Provider value={{ aws: siteConfigs?.aws }}>
      {props.children}
    </AuthContext.Provider>
  );
}
