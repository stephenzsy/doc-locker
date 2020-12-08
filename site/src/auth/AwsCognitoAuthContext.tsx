import React from "react";
import { AuthContext, IAuthContext } from "./AuthContext";

export type IAwsCognitoAuthContext = IAuthContext;
const provider = "aws-cognito";

export function AwsCognitoAuthContextProvider(props: {
  children?: React.ReactNode;
}): JSX.Element {
  return (
    <AuthContext.Provider value={{ provider }}>
      {props.children}
    </AuthContext.Provider>
  );
}
