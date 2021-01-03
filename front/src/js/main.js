import logo from "../logo.svg";
import "../styles/App.css";
import React, { useState } from "react";
import { Button, Input } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";

import { useQuery, gql } from "@apollo/client";

const CREATE_USER_QUERY = gql`
  {
    users(start:1, amount:3) {
      name
      email
      posts {
        amount
      }
    }
  }
`;

const Main = props => {
  const [mainState, setMainState] = useState({
    inputText: ""
  });

  function onButtonClick() {
    console.log(mainState.inputText)
  }

  const test = useQuery(CREATE_USER_QUERY)
  console.log("data:", test)

  return (
    <div className="App">
      <div className="createUser">
        <Button 
        onClick={onButtonClick}
        // onClick={props.onCreateUserClick}
        >Create User</Button>
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
      </div>
    </div>
  );
};

export default Main;
