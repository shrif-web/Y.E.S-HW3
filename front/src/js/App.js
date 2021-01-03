import logo from "../logo.svg";
import "../styles/App.css";
import React from "react";
import { Button, Input } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
import Main from './main.js'
import MainPage from './MainPage.js'


class App extends React.Component {
  constructor(props) {
    super(props);

    // this.onCreateUserClick = this.onCreateUserClick.bind(this);
  }

  onCreateUserClick() {
    // console.log("salam", this.inputName.value);
  }

  onUsernameChange(e) {
    // this.inputName.value = e.target.value;
  }

  render() {

    return (
      <div className="App">
        {/* <Main 
          onCreateUserClick={this.onCreateUserClick}
          onUsernameChange={this.onUsernameChange}
        /> */}
        <MainPage />
      </div>
    );
  }
}

export default App;
