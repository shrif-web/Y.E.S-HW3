import logo from "../logo.svg";
import "../styles/App.css";
import React, { useState } from "react";
import { Button, Input } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
import MainPage from "./MainPage.js";
import Login from "./Login.js";
import Register from "./Register.js";
import { Redirect, Route, Switch } from "react-router-dom";
import Header from "./Header.js";
import useToken from "./useToken";
import Dashboard from "./Dashboard.js";
import { useMutation, gql } from "@apollo/client";
import { useMediaQuery } from "react-responsive";

const REFRESH_TOKEN_MUTATION = gql`
  mutation RefreshToken {
    refreshToken
  }
`;

function App(props) {
  const { token, setToken } = useToken();

  const isMobile = useMediaQuery({
    query: "(max-device-width: 570px)"
  });

  const [sidebarIsOpen, setSidebarIsOpen] = useState(false);

  return (
    <div className="App">
      <Header
        setToken={setToken}
        refresh={props.refresh}
        isMobile={isMobile}
        sidebarIsOpen={sidebarIsOpen}
        setSidebarIsOpen={setSidebarIsOpen}
      />
      <Route exact path="/" >
        <MainPage isMobile={isMobile} />
      </Route>
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
          <Dashboard
            setToken={setToken}
            sidebarIsOpen={sidebarIsOpen}
            isMobile={isMobile}
            // setSidebarIsOpen={setSidebarIsOpen}
          />
        </Route>
      )}
    </div>
  );
}

export default App;
