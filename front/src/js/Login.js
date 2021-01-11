import React, { useState } from "react";
import { Form, Input, Button, Grid, Segment, Message } from "semantic-ui-react";
import gql from "graphql-tag";
import { useMutation } from "@apollo/client";
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

const LoginForm = props => {
  const [state, setState] = useState({
    username: " ",
    password: " ",
    error: ""
  });

  const [refreshToken] = useMutation(REFRESH_TOKEN_MUTATION, {
    onCompleted: ({ refreshToken }) => {
      console.log("()()()() refreshed token", refreshToken);
      props.setToken(refreshToken.token);
    }
  });

  const history = useHistory();

  const [login] = useMutation(LOGIN_MUTATION, {
    variables: {
      username: state.username,
      password: state.password
    },
    onCompleted: ({ login }) => {
      if (login.__typename == "Token") {
        props.setToken(login.token);
        history.push("/dashboard");
      } else {
        switch (login.__typename) {
          case "UserPassMissMatchException":
            setState({ ...state, error: constants.USER_PASS_MISMATCH });
            break;
          case "InternalServerException":
            setState({ ...state, error: constants.INTERNAL_SERVER_EXCEPTION });
            break;
        }
      }
    }
  });

  function handleLogin() {
    if (state.username && state.password) {
      console.log("handliing login?????????");
      login();
      setState({ ...state, error: "" });
    }
  }

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
      {state.error !== "" && <Message negative>{state.error}</Message>}
    </div>
  );
};

class Login extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div style={{ top: "50px", position: "absolute", width: "100%" }}>
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
                marginLeft: 20
              }}
            >
              <LoginForm
                setToken={this.props.setToken}
                setIntervalID={this.props.setIntervalID}
              />
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </div>
    );
  }
}

export default Login;
