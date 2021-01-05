import React, { useState } from "react";
import { Input, Menu, Button } from "semantic-ui-react";
import { Redirect, Route, Switch, useHistory } from "react-router-dom";
import constants from "../constants";

const LoggedInHeader = props => {
  return (
    <Menu fixed={props.fixed} inverted style={{ borderRadius: "0px" }}>
      <Menu.Item
        name="Logout"
        active={props.state.activeItem === "Logout"}
        onClick={props.handleItemClick}
      />
      <Menu.Menu position="right">
        <Menu.Item
          name="Info"
          active={props.state.activeItem === "Info"}
          onClick={props.handleItemClick}
        />
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
    <Menu fixed={props.fixed} inverted style={{ borderRadius: "0px" }}>
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
        history.push("/login");
        break;
      case "Register":
        console.log("Register");
        history.push("/register");
        break;
      case "Homepage":
        console.log("Homepage");
        history.push("/");
        break;
      case "Logout":
        console.log("Logout");
        localStorage.removeItem(constants.AUTH_TOKEN)
        // props.setToken('')
        history.push("/");
        break;
      case "Info":
        console.log("Info");
        break;
      case "Dashboard":
        console.log("Dashboard");
        history.push("/");
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
        />
      )}
      {/* <Button
        onClick={() => {
          localStorage.removeItem(constants.AUTH_TOKEN);
          setState({ ...state, loggedIn: false });
        }}
      >
        Remove token (test button)
      </Button> */}
      {/* <Button
        onClick={() => {
          localStorage.setItem(constants.AUTH_TOKEN, "test1");
          setState({ ...state, loggedIn: true });
        }}
      >
        Add token
      </Button> */}
    </div>
  );
};

export default Header;
