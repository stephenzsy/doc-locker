import { AmplifyAuthenticator } from "@aws-amplify/ui-react";
import { Box } from "@material-ui/core";
import React from "react";

export function LoginPage(): JSX.Element {
  return (
    <Box>
      <AmplifyAuthenticator />
    </Box>
  );
}
