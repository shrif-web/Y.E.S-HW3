import React from "react";
import "../styles/MainPage.css";
import PostGrid from "./PostGrid";

class MainPage extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <PostGrid isMobile={this.props.isMobile}/>
      </div>
    );
  }
}

export default MainPage;
