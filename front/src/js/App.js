import logo from "../logo.svg";
import "../styles/App.css";
import React, { useState } from "react";
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
    refreshToken {
      __typename
      ... on Token {
        token
      }
      ... on Exception {
        message
      }
    }
  }
`;

function App(props) {
  const { token, setToken } = useToken();

  const [state, setState] = useState({
    refreshed: false
  });

  const [intervalID, setIntervalID] = useState();

  const [refreshToken] = useMutation(REFRESH_TOKEN_MUTATION, {
    onCompleted: ({ refreshToken }) => {
      console.log("()()()() refreshed token", refreshToken);
      setToken(refreshToken.token);
    }
  });

  // function setTokenInterval() {
  //   const interval = setInterval(() => {
  //     refreshToken()
  //   }, 360000)
  //   return interval
  // }

  if (!token) {
    console.log("********* there is NO token!!!! ***********");
  } else {
    console.log("+++++++++ there IS token!!!! +++++++++");
    // const intr = setInterval(function() {
    //   refreshToken()
    // }, 120000)
    // setIntervalID(intr)
  }

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
        intervalID={intervalID}
      />
      {token && <Redirect exact from="/login" to="/dashboard" />}
      {token && <Redirect exact from="/register" to="/dashboard" />}
      {!token && <Redirect exact from="/dashboard" to="/" />}
      {!token && (
        <Route exact path="/">
          <MainPage isMobile={isMobile} />
        </Route>
      )}
      {!token && (
        <Route exact path="/register">
          <Register setToken={setToken} setIntervalID={setIntervalID} />
        </Route>
      )}
      {!token && (
        <Route exact path="/login">
          <Login setToken={setToken} setIntervalID={setIntervalID} />
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
      <Redirect exact from="/*" to="/" />
    </div>
  );
}

export default App;
