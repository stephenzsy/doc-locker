import { Auth, CognitoUser } from "@aws-amplify/auth";
import { Button, TextField, Typography } from "@material-ui/core";
import React from "react";
import { useAppContext } from "../AppContext";
import { AuthContext } from "../auth/AuthContext";

function AwsSettingsForm(): JSX.Element | null {
  const { aws } = React.useContext(AuthContext);
  const [currentUser, setCurrentUser] = React.useState<CognitoUser | undefined>(
    undefined
  );
  const [
    cognitoIdentityPoolIdPending,
    setCognitoIdentityPoolIdPending,
  ] = React.useState<string>("");
  const [
    cognitoIdentityPoolRegionPending,
    setCognitoIdentityPoolRegionPending,
  ] = React.useState<string>("");
  const [configPathPending, setConfigPathPending] = React.useState<string>("");
  React.useEffect(() => {
    if (aws) {
      (async () => {
        const user = (await Auth.currentAuthenticatedUser()) as CognitoUser;
        if (user) {
          setCurrentUser(user);
        }
        const attributes = await Auth.userAttributes(user);
        for (const attr of attributes) {
          switch (attr.Name) {
            case aws.cognitoUserPoolAttributesMapping.cognitoIdentityPoolId:
              setCognitoIdentityPoolIdPending(attr.Value);
              break;
            case aws.cognitoUserPoolAttributesMapping.cognitoIdentityPoolRegion:
              setCognitoIdentityPoolRegionPending(attr.Value);
              break;
            case aws.cognitoUserPoolAttributesMapping.configPath:
              setConfigPathPending(attr.Value);
              break;
          }
        }
        console.log(attributes);
      })();
    }
  }, [aws, setCognitoIdentityPoolIdPending]);

  if (!aws) {
    return null;
  }

  return (
    <form noValidate>
      <div>
        <div>
          <TextField
            fullWidth
            label="Cognito Identity Pool ID"
            value={cognitoIdentityPoolIdPending}
            onChange={(ev) => {
              setCognitoIdentityPoolIdPending(ev.target.value);
            }}
          />
        </div>
        <div>
          <TextField
            fullWidth
            label="Cognito Identity Pool Region"
            value={cognitoIdentityPoolRegionPending}
            onChange={(ev) => {
              setCognitoIdentityPoolRegionPending(ev.target.value);
            }}
          />
        </div>
        <div>
          <TextField
            fullWidth
            label="Config File S3 ARN"
            value={configPathPending}
            onChange={(ev) => {
              setConfigPathPending(ev.target.value);
            }}
          />
        </div>
        <div>
          <Button
            variant="contained"
            onClick={() => {
              if (currentUser) {
                Auth.updateUserAttributes(currentUser, {
                  [aws.cognitoUserPoolAttributesMapping
                    .cognitoIdentityPoolId]: cognitoIdentityPoolIdPending.trim(),
                  [aws.cognitoUserPoolAttributesMapping
                    .cognitoIdentityPoolRegion]: cognitoIdentityPoolRegionPending.trim(),
                  [aws.cognitoUserPoolAttributesMapping
                    .configPath]: configPathPending.trim(),
                });
              }
            }}
          >
            Save
          </Button>
        </div>
      </div>
    </form>
  );
}

export function Settings(): JSX.Element {
  const { endpoint, setEndpoint } = useAppContext();
  const [endpointPending, setEndpointPending] = React.useState<string>(
    endpoint || ""
  );

  return (
    <div>
      <div>
        <Typography variant="h6">Endpoint</Typography>
        <form noValidate>
          <div>
            <TextField
              label="Endpoint"
              value={endpointPending}
              onChange={(ev) => {
                setEndpointPending(ev.target.value);
              }}
            />
            <Button
              variant="contained"
              onClick={() => {
                if (endpointPending) {
                  setEndpoint(endpointPending);
                }
              }}
            >
              Save
            </Button>
          </div>
        </form>
      </div>
      <div>
        <Typography variant="h6">AWS Settings</Typography>
        <AwsSettingsForm />
      </div>
    </div>
  );
}
