import React from 'react';
import { BrowserRouter, Link, Route, Switch } from 'react-router-dom';
import './App.css';
import { Host } from './host/host';

function App() {
  return (
    <BrowserRouter>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/host">Host</Link>
            </li>
          </ul>
        </nav>
        <Switch>
          <Route path="/host">
            <Host />
          </Route>
          <Route path="/">
            Hello World!
          </Route>
        </Switch>
      </div>
    </BrowserRouter>
  );
}

export default App;
