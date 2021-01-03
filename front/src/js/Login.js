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
    username: "",
    password: ""
  });

  const history = useHistory();

  const [login] = useMutation(LOGIN_MUTATION, {
    variables: {
      username: state.username,
      password: state.password
    },
    onCompleted: ({ login }) => {
      console.log("herererererere", login);
      if (login.__typename == "Token") {
        localStorage.setItem(constants.AUTH_TOKEN, login.token);
        // history.push("/");
      } else {
        // Todo : ERROR!
      }
    }
  });

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
      />
      <Form.Button
        content="Login"
        control={Button}
        onClick={() => {
          login();
        }}
      />
      <Form.Button
        content="Test Button (Token?)"
        control={Button}
        onClick={() => {
          console.log(localStorage.getItem(constants.AUTH_TOKEN));
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
            <LoginForm />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    );
  }
}

export default Login;
