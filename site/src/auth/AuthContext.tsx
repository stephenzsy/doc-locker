import Amplify from "@aws-amplify/core";
import React from "react";
import { useAppContext } from "../AppContext";
import { ConfigurationsServicePromiseClient } from "../generated/configurations_grpc_web_pb";
import { SiteConfigurationsRequest } from "../generated/configurations_pb";

export interface IAwsSiteConfigurations {
  cognitoIdenityPoolId: string;
  cognitoRegion: string;
  cognitoUserPoolId: string;
  cognitoUserPoolWebClientId: string;
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
      const response = await client.siteConfigurations(
        new SiteConfigurationsRequest()
      );
      const configs = JSON.parse(
        response.getSiteconfigurationsjson()
      ) as ISiteConfigurations;
      console.log(configs);
      setSiteConfigs(configs);
      Amplify.configure({
        Auth: {
          // REQUIRED - Amazon Cognito Region
          identityPoolId: configs.aws.cognitoIdenityPoolId,
          region: configs.aws.cognitoRegion,
          userPoolId: configs.aws.cognitoUserPoolId,
          userPoolWebClientId: configs.aws.cognitoUserPoolWebClientId,
        },
      });
    })();
  }, [endpoint]);

  return (
    <AuthContext.Provider value={{ aws: siteConfigs?.aws }}>
      {props.children}
    </AuthContext.Provider>
  );
}
