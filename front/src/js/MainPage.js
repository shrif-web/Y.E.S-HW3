import React from "react";
import { Grid, Card, Icon, Message, Segment } from "semantic-ui-react";
import "../styles/MainPage.css";
import Header from "./Header.js";
import { Redirect, Route, Switch } from "react-router-dom";
import PostGrid from "./PostGrid";

class MainPage extends React.Component {
  constructor(props) {
    super(props);

    this.updateSize = this.updateSize.bind(this);
    // window.addEventListener("resize", this.updateSize);

    this.state = {
      isMobile: false
    };
  }

  updateSize() {
    if (window.innerWidth < 600 && !this.state.isMobile) {
      this.setState({ isMobile: true });
    }

    if (window.innerWidth > 600 && this.state.isMobile) {
      this.setState({ isMobile: false });
    }
  }

  render() {
    return (
      <div>
        <PostGrid />
      </div>
    );
  }
}

export default MainPage;