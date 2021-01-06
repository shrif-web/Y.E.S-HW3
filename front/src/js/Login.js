import React, { useState } from "react";
import { Form, Input, Button, Grid, Segment, Message } from "semantic-ui-react";
import gql from "graphql-tag";
import { useQuery, useMutation } from "@apollo/client";
import constants from "../constants.js";
import { useHistory } from "react-router-dom";
import "../styles/Login.css";

const LOGIN_MUTATION = gql`
  mutation LoginUser($username: String!, $password: String!) {
    login(input: { username: $username, password: $password }) {
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

const LoginForm = props => {
  const [state, setState] = useState({
    username: " ",
    password: " "
  });

  const history = useHistory();

  const [login] = useMutation(LOGIN_MUTATION, {
    variables: {
      username: state.username,
      password: state.password
    },
    onCompleted: ({ login }) => {
      console.log("loginnnnnnnn:", login);
      if (login.__typename == "Token") {
        console.log("haaaaaaaaa?");
        // localStorage.setItem(constants.AUTH_TOKEN, login.token);
        props.setToken(login.token);
        history.push("/");
      } else {
        // Todo : ERROR!
        console.log("error in login!", login);
        alert(login.__typename);
      }
    }
  });

  function handleLogin() {
    if (state.username && state.password) {
      login();
    }
  }

  console.log("login rendereddddddddddd+++++++++++++++++++++++++++");

  return (
    <div>
      <Form>
        <Segment>
          <Form.Input
            icon="user"
            iconPosition="left"
            placeholder="Enter your username"
            control={Input}
            onChange={e => {
              setState({
                ...state,
                username: e.target.value
              });
            }}
            error={
              !state.username && {
                content: "Please enter a username!",
                pointing: "below"
              }
            }
          />
          <Form.Input
            icon="lock"
            iconPosition="left"
            placeholder="Enter your password"
            control={Input}
            type="password"
            onChange={e => {
              setState({
                ...state,
                password: e.target.value
              });
            }}
            error={
              !state.password && {
                content: "Please enter a password please!",
                pointing: "below"
              }
            }
          />
          <Form.Button
            fluid
            color="blue"
            content="Login"
            control={Button}
            onClick={() => {
              handleLogin();
            }}
          />
        </Segment>
      </Form>
      <Message>
        New to us? <a href="/register">Sign Up</a>
      </Message>
    </div>
  );
};

class Login extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Grid
        centered
        style={{ height: "100vh" }}
        verticalAlign="middle"
        textAlign="center"
      >
        <Grid.Row>
          <Grid.Column
            style={{
              maxWidth: 450,
              marginRight: 20,
              marginLeft: 20,
              marginTop: -100
            }}
          >
            <LoginForm setToken={this.props.setToken} />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    );
  }
}

export default Login;
