import logo from "../logo.svg";
import "../styles/App.css";
import React, { useState } from "react";
import { Button, Input, Header } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";

import { useQuery, useMutation } from "@apollo/client";
import gql from "graphql-tag";

const GET_USER_QUERY = gql`
  {
    users(start: 0, amount: 100) {
      name
      email
      posts {
        title
      }
    }
  }
`;

const CREATE_USER_MUTATION = gql`
  mutation createUser($username: String!, $password: String!, $email: String!) {
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

const Main = props => {
  const [mainState, setMainState] = useState({
    inputText: ""
  });

  function onButtonClick() {
    console.log(mainState.inputText);
  }

  const { data } = useQuery(GET_USER_QUERY);
  console.log("data:", data);

  const [createUser] = useMutation(CREATE_USER_MUTATION, {
    variables: {
      username: "44",
      password: "55",
      email: "66@66.com"
    }
  });

  return (
    <div className="App">
      <div className="createUser">
        <Button
          onClick={() => {
            createUser({

            });
          }}
          // onClick={props.onCreateUserClick}
        >
          Create User
        </Button>
        <Input
          onChange={e => {
            setMainState({
              ...mainState,
              inputText: e.target.value
            });
          }}
          //   onChange={props.onUsernameChange.bind(this)}
          placeholder="Enter name ..."
        />
        <Header as="h1">{JSON.stringify(data)}</Header>
      </div>
    </div>
  );
};

export default Main;
