import logo from "../logo.svg";
import "../styles/App.css";
import React from "react";
import { Button, Input } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
import Main from "./main.js";
import MainPage from "./MainPage.js";
import Login from "./Login.js";
import Register from "./Register.js";
import { Redirect, Route, Switch } from "react-router-dom";
import Header from "./Header.js";
import useToken from "./useToken";
import userToken from "./useToken";

function App() {
  const { token, setToken } = userToken();

  return (
    <div className="App">
      <Header />
      <Route exact path="/" component={MainPage} />
      <Route exact path="/register" component={Register} />
      <Route exact path="/login">
        <Login setToken={setToken} />
      </Route>
    </div>
  );
}

export default App;
