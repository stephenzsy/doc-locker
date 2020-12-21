import { InteractionType } from "@azure/msal-browser";
import {
  AuthenticatedTemplate,
  UnauthenticatedTemplate,
  useMsalAuthentication
} from "@azure/msal-react";
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
  return <div />;
}
