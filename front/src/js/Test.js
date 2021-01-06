import React, { useState } from "react";
import App from "./App.js";
import { Button } from "semantic-ui-react";

class Test extends React.Component {
  state = {
    trigger: false
  };

  refresh() {
    this.setState({ trigger: !this.state.trigger });
  }

  render() {
    console.log("rendreing Test component..............");
    return (
      <div>
        {/* <Button onClick={this.refresh.bind(this)}>Refresh</Button>
        <Button
          onClick={() => {
            localStorage.setItem("test", "salam bahar:)");
            this.refresh();
          }}
        >
          Set In Localstorage
        </Button>
        <Button
          onClick={() => {
            console.log("?? => ", localStorage.getItem("test"));
          }}
        >
          Check Localstorage
        </Button>
        <Button
          onClick={() => {
            localStorage.removeItem("test");
            this.refresh();
          }}
        >
          Remove Test from Localstorage
        </Button>
        {localStorage.getItem("test") && <div>Has test :D</div>} */}
        <App refresh={this.refresh.bind(this)} />
      </div>
    );
  }
}

export default Test;
