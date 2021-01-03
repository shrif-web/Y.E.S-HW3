import React from "react";
import { Form, Input, Button, Grid } from "semantic-ui-react";
import gql from "graphql-tag";

const REGISTER_MUTATION = gql`
{
    mutation register($username:String!, $email: String!, $password: String!) {
        
    }
}
`;

const RegisterForm = props => {
  return (
    <Form>
      <Form.Field
        label="Username"
        placeholder="Enter your username ..."
        control={Input}
      />
      <Form.Field
        label="Email"
        placeholder="Enter your email ..."
        control={Input}
      />
      <Form.Field
        label="Password"
        placeholder="Choose a password ..."
        control={Input}
      />
      <Form.Field
        label="Confirm Password"
        placeholder="Repeat your password ..."
        control={Input}
      />
      <Form.Field content="Register" control={Button} />
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
            <RegisterForm />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    );
  }
}

export default Register;
