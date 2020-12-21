import { Layout, Menu } from "antd";
import React from "react";
import {
  BrowserRouter,
  Route,
  Switch,
  useHistory,
  useLocation,
} from "react-router-dom";
import "./App.css";
import { AppContextProvider } from "./AppContext";
import { AuthContextProvider } from "./auth/AuthContext";
import { Host } from "./host/host";
import { Settings as SettingsPage } from "./settings/Settings";

function HeaderMenu(): JSX.Element {
  const location = useLocation();
  const history = useHistory();
  const onChange = React.useCallback(
    (value: string) => {
      history.push(value);
    },
    [history]
  );

  return (
    <Menu
      theme="dark"
      mode="horizontal"
      selectedKeys={[location.pathname]}
      onSelect={(ev) => {
        onChange(ev.key.toString());
      }}
    >
      <Menu.Item key="/">Home</Menu.Item>
      <Menu.Item key="/host">Host</Menu.Item>
      <Menu.Item key="/settings">Settings</Menu.Item>
    </Menu>
  );
}

function App(): JSX.Element {
  return (
    <AppContextProvider>
      <AuthContextProvider>
        <BrowserRouter>
          <Layout className="layout">
            <Layout.Header>
              <HeaderMenu />
            </Layout.Header>
            <Layout.Content>
              <Switch>
                <Route path="/host">
                  <Host />
                </Route>
                <Route path="/settings">
                  <SettingsPage />
                </Route>
                <Route path="/">Hello World!</Route>
              </Switch>
            </Layout.Content>
          </Layout>
        </BrowserRouter>
      </AuthContextProvider>
    </AppContextProvider>
  );
}

export default App;
