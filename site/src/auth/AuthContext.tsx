import React from "react";

export interface IAuthContext {
  provider: string;
}

export const AuthContext = React.createContext<IAuthContext>({
  provider: "",
});
