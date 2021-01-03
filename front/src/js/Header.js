import React, { useState } from "react";
import { Input, Menu } from "semantic-ui-react";
import { Redirect, Route, Switch, useHistory } from "react-router-dom";

const Header = props => {
  const [state, setState] = useState({
    activeItem: ""
  });

  const history = useHistory();

  function handleItemClick(e, { name }) {
    setState({ activeItem: name });
    switch (name) {
      case "Login":
        console.log("login");
        history.push("/login");
        // return <Redirect to="/login" />
        break;
      case "Register":
        console.log("Register");
        break;
      case "Homepage":
        console.log("Homepage");
        // return <Redirect to="/mainpage" />
        history.push("/");
    }
  }

  return (
    <Menu fixed={props.fixed} inverted>
      <Menu.Item
        name="Login"
        active={state.activeItem === "Login"}
        onClick={handleItemClick}
      />
      <Menu.Item
        name="Register"
        active={state.activeItem === "Register"}
        onClick={handleItemClick}
      />
      <Menu.Menu position="right">
        <Menu.Item
          name="Homepage"
          active={state.activeItem === "Homepage"}
          onClick={handleItemClick}
        />
      </Menu.Menu>
    </Menu>
  );
};

export default Header;

// export default class Header extends React.Component {
//   state = { activeItem: "Homepage" };

//   handleItemClick = (e, { name }) => {
//     this.setState({ activeItem: name });
//     switch (name) {
//       case "Login":
//         console.log("login");
//         history.push("/login");
//         // return <Redirect to="/login" />
//         break;
//       case "Register":
//         console.log("Register");
//         break;
//       case "Homepage":
//         console.log("Homepage");
//         // return <Redirect to="/mainpage" />
//         history.push(".mainpage");
//     }
//   };

//   render() {
//     const { activeItem } = this.state;

//     return (
//       <Menu fixed={this.props.fixed} inverted>
//         <Menu.Item
//           name="Login"
//           active={activeItem === "Login"}
//           onClick={this.handleItemClick}
//         />
//         <Menu.Item
//           name="Register"
//           active={activeItem === "Register"}
//           onClick={this.handleItemClick}
//         />
//         <Menu.Menu position="right">
//           <Menu.Item
//             name="Homepage"
//             active={activeItem === "Homepage"}
//             onClick={this.handleItemClick}
//           />
//         </Menu.Menu>
//       </Menu>
//     );
//   }
// }
