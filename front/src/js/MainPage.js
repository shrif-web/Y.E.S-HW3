import React from "react";
import { Grid, Card, Icon } from "semantic-ui-react";
import "../styles/MainPage.css";
import Header from "./Header.js";

const posts = [
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "KK",
    title: "Back End Development",
    description: "Bye! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  },
  {
    username: "Bahar",
    title: "Front End Development",
    description: "Hi! nice to meet you"
  }
];

const PostCell = ({ username, title, description }) => {
  return (
    <Grid.Column>
      <Card header={username} meta={title} description={description} fluid />
    </Grid.Column>
  );
};

class MainPage extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <Header />
        <Grid columns={3}>
          {posts.map(post => {
            return (
              <PostCell
                username={post.username}
                title={post.title}
                description={post.description}
              />
            );
          })}
        </Grid>
      </div>
    );
  }
}

export default MainPage;
