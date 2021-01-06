import React, { useState } from "react";
import { Form, Input, Button, Grid } from "semantic-ui-react";
import gql from "graphql-tag";
import { useQuery, useMutation } from "@apollo/client";
import constants from "../constants.js";
import { useHistory } from "react-router-dom";

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
      console.log("loginnnnnnnn:", login)
      if (login.__typename == "Token") {
        console.log("haaaaaaaaa?")
        // localStorage.setItem(constants.AUTH_TOKEN, login.token);
        props.setToken(login.token);
        history.push("/");
      } else {
        // Todo : ERROR!
        console.log("error in login!", login)
        alert(login.__typename)
      }
    }
  });

  function handleLogin() {
    if (state.username && state.password) {
      login();
    }
  }

  return (
    <Form>
      <Form.Field
        label="Username"
        placeholder="Enter your username ..."
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
      <Form.Field
        label="Password"
        placeholder="Enter your password ..."
        control={Input}
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
        content="Login"
        control={Button}
        onClick={() => {
          handleLogin();
        }}
      />
    </Form>
  );
};

class Login extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Grid centered verticalAlign="middle">
        <Grid.Row columns={3}>
          <Grid.Column>
            <LoginForm setToken={this.props.setToken} />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    );
  }
}

export default Login;
