import React from "react";
import { Form, Input, Button, Grid } from "semantic-ui-react";

const LoginForm = props => {
  return (
    <Form>
      <Form.Field
        label="Username"
        placeholder="Enter your username ..."
        control={Input}
      />
      <Form.Field
        label="Password"
        placeholder="Enter your password ..."
        control={Input}
      />
      <Form.Field content="Login" control={Button} />
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
