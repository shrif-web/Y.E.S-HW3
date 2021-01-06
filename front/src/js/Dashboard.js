import React, { useState } from "react";
import { useQuery, useMutation } from "@apollo/client";
import gql from "graphql-tag";
import constants from "../constants";
import { Card, Grid, Button } from "semantic-ui-react";

const GET_USER_POSTS = gql`
  # query getUser
  {
    user {
      name
      posts {
        id
        title
        content
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

const CREATE_POST_MUTATION = gql`
  mutation CreatePost($title: String!, $content: String!) {
    createPost(input: { title: $title, content: $content }) {
      __typename
      ... on Post {
        id
        title
      }
      ... on Exception {
        message
      }
    }
  }
`;

const DELETE_POST_MUTATION = gql`
  mutation DeletePost($targetID: String!) {
    deletePost(targetID: $targetID) {
      __typename
      ... on OperationSuccessfull {
        message
      }
      ... on Exception {
        message
      }
    }
  }
`;

const UserPosts = props => {
  const [state, setState] = useState({
    trigger: false
  });
  console.log("in UserPostsssssss:", props);
  const { data, loading, error } = useQuery(GET_USER_POSTS);

  const [createPost] = useMutation(CREATE_POST_MUTATION, {
    variables: {
      title: "title of the test post :)",
      content: "content of the test post !"
    },
    onCompleted({ data }) {
      triggerState();
      // console.log("data?");
    }
  });

  const [deletePost] = useMutation(DELETE_POST_MUTATION, {
    onCompleted({ message }) {
      triggerState();
    }
  });

  console.log("in dashboard");
  console.log("data: ", data);
  console.log("error:", error);
  console.log("loading:", loading);

  function addPost() {}

  function triggerState() {
    console.log("here trigerringgggggg //////////");
    setState({ trigger: !state.trigger });
  }

  return (
    <div>
      <Button
        icon="plus"
        onClick={() => {
          createPost();
        }}
      />
      <Grid columns={2} stackable>
        {!loading &&
          data.user.posts.map(post => {
            return (
              <Grid.Column>
                <Card
                  key={post.id}
                  id={post.id}
                  header={post.title}
                  meta={post.created_by.name}
                  description={post.content}
                  fluid
                  onClick={() => {
                    console.log("++_+_+_+_++_", post);
                  }}
                />
                <Button.Group>
                  <Button
                    icon="edit"
                    onClick={() => {
                      console.log("clicked in edit");
                    }}
                  />
                  <Button
                    icon="minus"
                    onClick={() => {
                      console.log("clicked on -");
                      deletePost({ variables: { targetID: post.id } });
                    }}
                  />
                </Button.Group>
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
