import { AppBar, Box, Tab, Tabs } from "@material-ui/core";
import React from "react";
import {
  BrowserRouter,
  Route,
  Switch,
  useHistory,
  useLocation,
} from "react-router-dom";
import "./App.css";
import { Host } from "./host/host";

function App(): JSX.Element {
  return (
    <BrowserRouter>
      <div>
        <AppBarTabs />
        <Switch>
          <Route path="/host">
            <Box>
              <Host />
            </Box>
          </Route>
          <Route path="/login">
            <Box>
              <Host />
            </Box>
          </Route>
          <Route path="/">Hello World!</Route>
        </Switch>
      </div>
    </BrowserRouter>
  );
}

function AppBarTabs(): JSX.Element {
  const location = useLocation();
  const history = useHistory();
  const onChange = React.useCallback(
    (event, value) => {
      history.push(value);
    },
    [history]
  );

  return (
    <AppBar position="static">
      <Tabs value={location.pathname} onChange={onChange}>
        <Tab label="Home" value="/" />
        <Tab label="Host" value="/host" />
        <Tab label="Login" value="/login" />
      </Tabs>
    </AppBar>
  );
}

export default App;
