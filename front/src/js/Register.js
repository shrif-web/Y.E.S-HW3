import React, { useState } from "react";
import { Form, Input, Button, Grid, Segment } from "semantic-ui-react";
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

const RegisterForm = props => {
  const [state, setState] = useState({
    username: "",
    email: "",
    password: "",
    confirmPass: ""
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
      } else {
        //Handle Errors
        switch (createUser.__typename) {
          case "DuplicateUsernameException":
            alert(constants.DUPLICATE_USERNAME_ERROR);
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
        history.push("/");
      } else {
        // Todo : ERROR!
      }
    }
  });

  function handleRegister() {
    if (state.password === state.confirmPass) {
      registerUser();
    } else {
      alert(constants.PASSWORDS_DIFFER);
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
    </Form>
  );
};

class Register extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Grid
        centered
        verticalAlign="middle"
        textAlign="center"
        style={{ height: "100vh" }}
      >
        <Grid.Row>
          <Grid.Column style={{ maxWidth: 450, marginRight: 20, marginLeft: 20, marginTop: -100 }}>
            <RegisterForm setToken={this.props.setToken} />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    );
  }
}

export default Register;
