import React, { useState } from "react";
import { Form, Input, Button, Grid, Segment, Message } from "semantic-ui-react";
import gql from "graphql-tag";
import { useMutation } from "@apollo/client";
import { useHistory } from "react-router-dom";
import constants from "../constants.js";

const REGISTER_MUTATION = gql`
  mutation CreateUser($username: String!, $email: String!, $password: String!) {
    createUser(
      target: { username: $username, password: $password, email: $email }
    ) {
      __typename
      ... on User {
        name
      }
      ... on Exception {
        message
      }
    }
  }
`;

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

const RegisterForm = props => {
  const [state, setState] = useState({
    username: "",
    email: "",
    password: "",
    confirmPass: "",
    error: ""
  });

  const history = useHistory();

  const [registerUser] = useMutation(REGISTER_MUTATION, {
    variables: {
      username: state.username,
      email: state.email,
      password: state.password
    },
    onCompleted: ({ createUser }) => {
      console.log("createUser:", createUser);
      if (createUser.__typename == "User") {
        login();
        setState({ ...state, error: "" });
      } else {
        //Handle Errors
        console.log("+_+_+_+_+_+_+_+", createUser.message);
        switch (createUser.__typename) {
          case "DuplicateUsernameException":
            setState({ ...state, error: constants.DUPLICATE_USERNAME_ERROR });
            break;
          case "InternalServerException":
            setState({ ...state, error: constants.INTERNAL_SERVER_EXCEPTION });
            break;
        }
      }
    }
  });

  const [login] = useMutation(LOGIN_MUTATION, {
    variables: {
      username: state.username,
      password: state.password
    },
    onCompleted: ({ login }) => {
      if (login.__typename == "Token") {
        props.setToken(login.token);
        history.push("/dashbaord");
      } else {
        // Todo : ERROR!
      }
    }
  });

  function handleRegister() {
    if (
      state.username !== "" &&
      state.password !== "" &&
      state.email !== "" &&
      state.confirmPass !== ""
    ) {
      if (state.password === state.confirmPass) {
        registerUser();
      } else {
        setState({ ...state, error: constants.PASSWORDS_DIFFER });
      }
    } else {
      setState({ error: constants.EMPTY_FIELDS });
    }
  }

  return (
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
        />
        <Form.Input
          icon="mail"
          iconPosition="left"
          placeholder="Enter your email"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              email: e.target.value
            });
          }}
        />
        <Form.Input
          icon="lock"
          iconPosition="left"
          type="password"
          placeholder="Choose a password ..."
          control={Input}
          onChange={e => {
            setState({
              ...state,
              password: e.target.value
            });
          }}
        />
        <Form.Input
          icon="lock"
          iconPosition="left"
          type="password"
          placeholder="Repeat your password ..."
          control={Input}
          onChange={e => {
            setState({
              ...state,
              confirmPass: e.target.value
            });
          }}
        />
        <Form.Button
          fluid
          color="blue"
          content="Register"
          control={Button}
          onClick={() => {
            handleRegister();
          }}
        />
      </Segment>
      <Message>
        Already have an account? <a href="/login">Login</a>
      </Message>
      {state.error !== "" && <Message negative>{state.error}</Message>}
    </Form>
  );
};

class Register extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div style={{ top: "50px", position: "absolute", width: "100%" }}>
        <Grid
          centered
          verticalAlign="middle"
          textAlign="center"
          style={{ height: "100vh" }}
        >
          <Grid.Row>
            <Grid.Column
              style={{ maxWidth: 450, marginRight: 20, marginLeft: 20 }}
            >
              <RegisterForm
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

export default Register;
