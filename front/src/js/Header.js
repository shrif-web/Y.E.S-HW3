import React, { useState } from "react";
import { Input, Menu, Button } from "semantic-ui-react";
import { Redirect, Route, Switch, useHistory } from "react-router-dom";
import constants from "../constants";
import { useMutation, gql } from "@apollo/client";

// const REFRESH_TOKEN_MUTATION = gql`
//   mutation RefreshToken {
//     refreshToken {
//       __typename
//       ... on Token {
//         token
//       }
//       ... on Exception {
//         message
//       }

//     }
//   }
// `;

const LoggedInHeader = props => {

  // const [refreshToken] = useMutation(REFRESH_TOKEN_MUTATION, {
  //   onCompleted: ({ refreshToken }) => {
  //     console.log("refreshed token", refreshToken.token);
  //     props.setToken(refreshToken.token)
  //   }
  // });

  return (
    <Menu fixed="top" inverted style={{ borderRadius: "0px" }}>
      {/* <Menu.Item name="Refresh Token" onClick={() => {
        refreshToken()
      }} /> */}
      {props.isMobile && (
        <Menu.Item
          icon="bars"
          onClick={() => {
            props.setSidebarIsOpen(!props.sidebarIsOpen);
          }}
        />
      )}
      <Menu.Item
        name="Logout"
        active={props.state.activeItem === "Logout"}
        onClick={props.handleItemClick}
      />
      <Menu.Menu position="right">
        <Menu.Item
          name="Dashboard"
          active={props.state.activeItem === "Dashboard"}
          onClick={props.handleItemClick}
        />
      </Menu.Menu>
    </Menu>
  );
};

const MainHeader = props => {
  return (
    <Menu
      fixed="top"
      inverted
      style={{ borderRadius: "0px", marginBottom: 20 }}
    >
      <Menu.Item
        name="Login"
        active={props.state.activeItem === "Login"}
        onClick={props.handleItemClick}
      />
      <Menu.Item
        name="Register"
        active={props.state.activeItem === "Register"}
        onClick={props.handleItemClick}
      />
      <Menu.Menu position="right">
        <Menu.Item
          name="Homepage"
          active={props.state.activeItem === "Homepage"}
          onClick={props.handleItemClick}
        />
      </Menu.Menu>
    </Menu>
  );
};

const Header = props => {
  const [state, setState] = useState({
    activeItem: "",
    loggedIn: false
  });

  const history = useHistory();
  const auth_token = localStorage.getItem(constants.AUTH_TOKEN);


  function handleItemClick(e, { name }) {
    setState({ activeItem: name });
    switch (name) {
      case "Login":
        console.log("login");
        props.refresh();
        // interval = props.setTokenInterval()
        history.push("/login");
        break;
      case "Register":
        props.refresh();
        console.log("Register");
        history.push("/register");
        break;
      case "Homepage":
        console.log("Homepage");
        history.push("/");
        break;
      case "Logout":
        console.log("Logout");
        localStorage.removeItem(constants.AUTH_TOKEN);
        clearInterval(props.intervalID)
        history.push("/");
        window.location.reload(false);
        break;
      case "Dashboard":
        console.log("Dashboard");
        history.push("/dashboard");
    }
  }

  return (
    <div>
      {!auth_token ? (
        <MainHeader
          handleItemClick={handleItemClick}
          state={state}
          setState={setState}
        />
      ) : (
        <LoggedInHeader
          handleItemClick={handleItemClick}
          state={state}
          setState={setState}
          setToken={props.setToken}
          isMobile={props.isMobile}
          setSidebarIsOpen={props.setSidebarIsOpen}
          sidebarIsOpen={props.sidebarIsOpen}
        />
      )}
    </div>
  );
};

export default Header;
