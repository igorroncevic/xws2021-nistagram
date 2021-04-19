import './App.css';
import * as React from "react";
import {BrowserRouter, Switch, Route} from 'react-router-dom';
import IndexPage from "./pages/IndexPage";

export default class App extends React.Component {
  constructor () {
    super();
    this.state = {
      userRole : "",
      username : "",
      Id : ""
    }
  };
  render() {
    document.title = "Nistagram"
    return (
        <BrowserRouter>
          <Switch>
            <Route exact path="/"  render={(props) => <IndexPage {...props} /> } />

          </Switch>
        </BrowserRouter>
    );
  }
}

