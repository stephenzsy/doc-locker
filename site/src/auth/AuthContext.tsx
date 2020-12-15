import { MsalProvider } from "@azure/msal-react";
import React from "react";

export interface IAuthContext {
  provider: string;
}

export const AuthContext = React.createContext<IAuthContext>({
  provider: "",
});

export function AuthContextProvider(props: {
  children?: React.ReactNode;
}): JSX.Element {
  return (
    <MsalProvider instance={undefined as any}>{props.children}</MsalProvider>
  );
}
