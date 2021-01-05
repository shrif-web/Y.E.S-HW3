import React from "react";
import { useQuery, useMutation } from "@apollo/client";
import gql from "graphql-tag";
import constants from "../constants";
import { Card, Grid, Button } from "semantic-ui-react";

const GET_USER_POSTS = gql`
  query getUser {
    user {
      name
      posts {
        title
        created_by {
          name
          posts {
            title
          }
        }
      }
    }
  }
`;

const UserPosts = props => {
  const { data, loading, error } = useQuery(GET_USER_POSTS);

  //   console.log("in dashboard");
  //   console.log("data: ", data);
  //   console.log("error:", error);
  //   console.log("loading:", loading);

  return (
    <div>
      <Button.Group>
        <Button icon="plus" onClick={() => {
            console.log("clicked in +")
        }} />
        <Button icon="minus" onClick={() => {
            console.log("clicked on -")
        }} />
      </Button.Group>
      <Grid columns={2} stackable>
        {!loading &&
          data.user.posts.map(post => {
            return (
              <Grid.Column>
                <Card
                  header={post.title}
                  meta={post.created_by.name}
                  description={post.content}
                  fluid
                />
              </Grid.Column>
            );
          })}
      </Grid>
    </div>
  );
};

class Dashboard extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <UserPosts />
      </div>
    );
  }
}

export default Dashboard;
