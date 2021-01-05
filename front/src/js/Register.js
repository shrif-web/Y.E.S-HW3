import React, {useState} from "react";
import { Form, Input, Button, Grid } from "semantic-ui-react";
import gql from "graphql-tag";
import { useMutation } from "@apollo/client";
import {useHistory} from 'react-router-dom'

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
    username: '',
    email: '',
    password: '',
    confirmPass: ''
  })

  const history = useHistory()

  const [registerUser] = useMutation(REGISTER_MUTATION, {
    variables: {
      username: state.username,
      email: state.email,
      password: state.password
    },
    onCompleted: ({createUser}) => {
      if (createUser.__typename == "User") {
        login()
      } else {
        // Todo : ERROR!
        alert(createUser.__typename)
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
          })
        }}
      />
      <Form.Field
        label="Email"
        placeholder="Enter your email ..."
        control={Input}
        onChange={e => {
          setState({
            ...state,
            email: e.target.value
          })
        }}
      />
      <Form.Field
        label="Password"
        placeholder="Choose a password ..."
        control={Input}
        onChange={e => {
          setState({
            ...state,
            password: e.target.value
          })
        }}
      />
      <Form.Field
        label="Confirm Password"
        placeholder="Repeat your password ..."
        control={Input}
        onChange={e => {
          setState({
            ...state,
            confirmPass: e.target.value
          })
        }}
      />
      <Form.Button
        content="Register"
        control={Button}
        onClick={() => {
          registerUser();
          // Handle error if passwords are not the same
          // login()
        }}
      />
    </Form>
  );
};

class Register extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Grid centered verticalAlign="middle">
        <Grid.Row columns={3}>
          <Grid.Column>
            <RegisterForm setToken={this.props.setToken} />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    );
  }
}

export default Register;
