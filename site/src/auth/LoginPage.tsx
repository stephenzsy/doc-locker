import { AmplifyAuthenticator } from "@aws-amplify/ui-react";
import { InteractionType } from "@azure/msal-browser";
import {
  AuthenticatedTemplate,
  UnauthenticatedTemplate,
  useMsalAuthentication,
} from "@azure/msal-react";
import { Box, Paper, Tab, Tabs } from "@material-ui/core";
import React from "react";

function MsalLogin(): JSX.Element {
  useMsalAuthentication(InteractionType.Popup);
  return (
    <>
      <p>Anyone can see this paragraph.</p>
      <AuthenticatedTemplate>
        <p>At least one account is signed in!</p>
      </AuthenticatedTemplate>
      <UnauthenticatedTemplate>
        <p>No users are signed in!</p>
      </UnauthenticatedTemplate>
    </>
  );
}

export function LoginPage(): JSX.Element {
  const [selectedTab, setSelectedTab] = React.useState(0);
  return (
    <Box>
      <Paper>
        <Tabs
          value={selectedTab}
          indicatorColor="primary"
          textColor="primary"
          onChange={(ev, value) => {
            setSelectedTab(value);
          }}
          aria-label="disabled tabs example"
        >
          <Tab label="Azure" />
          <Tab label="AWS" />
        </Tabs>
        {selectedTab === 0 && <MsalLogin />}
        {selectedTab === 1 && <AmplifyAuthenticator />}
      </Paper>
    </Box>
  );
}
