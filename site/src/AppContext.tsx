import React from "react";

export interface IAppContext {
  endpoint: string;
  setEndpoint: (value: string) => void;
}

const AppContext = React.createContext<IAppContext>({
  endpoint: "",
  setEndpoint: () => undefined,
});

const localStorageKey = "doc-locker-current-endpoint";

export function AppContextProvider(props: {
  children?: React.ReactNode;
}): JSX.Element {
  const [endpoint, setEndpointState] = React.useState<string>(() => {
    return localStorage.getItem(localStorageKey) || "";
  });
  const setEndpoint = React.useCallback<(value: string) => void>(
    (value) => {
      localStorage.setItem(localStorageKey, value);
      setEndpointState(value);
    },
    [setEndpointState]
  );
  const contextValue = React.useMemo(
    (): IAppContext => ({
      endpoint,
      setEndpoint,
    }),
    [endpoint, setEndpoint]
  );
  return (
    <AppContext.Provider value={contextValue}>
      {props.children}
    </AppContext.Provider>
  );
}

export function useAppContext(): IAppContext {
  return React.useContext(AppContext);
}
