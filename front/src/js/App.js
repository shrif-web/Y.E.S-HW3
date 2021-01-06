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
import Dashboard from "./Dashboard.js";
import { useMutation, gql } from "@apollo/client";

const REFRESH_TOKEN_MUTATION = gql`
  mutation RefreshToken {
    refreshToken
  }
`;

function App(props) {
  const { token, setToken } = useToken();

  console.log("------------------------ token in App compoennt:", token);

  return (
    <div className="App">
      <Header setToken={setToken} refresh={props.refresh} />
      <Route exact path="/" component={MainPage} />
      {!token && (
        <Route exact path="/register">
          <Register setToken={setToken} />
        </Route>
      )}
      {!token && (
        <Route exact path="/login">
          <Login setToken={setToken} />
        </Route>
      )}
      {token && (
        <Route exact path="/dashboard">
          <Dashboard setToken={setToken} />
        </Route>
      )}
    </div>
  );
}

export default App;
