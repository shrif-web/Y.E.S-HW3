import React from "react";
import { useQuery, useMutation } from "@apollo/client";
import { Grid, Card, Button } from "semantic-ui-react";
import gql from "graphql-tag";

const GET_POSTS_QUERY = gql`
  {
    posts(start: 0, amount: 100) {
      id
      title
      content
      created_at
      created_by {
        name
      }
    }
  }
`;

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

const PostCell = ({ username, title, description }) => {
  return (
    <Grid.Column>
      <Card header={username} meta={title} description={description} fluid />
    </Grid.Column>
  );
};

const PostGrid = props => {
  const { data, loading, error } = useQuery(GET_USER_QUERY);
  console.log("data:", data);
  console.log("error:", error);
  console.log("loading:", loading);

  const [createPost] = useMutation(CREATE_POST_MUTATION, {
    variables: {
      title: "TITLE",
      content: "CONTENT"
    }
  });

  return (
    <div>
      <Grid columns={3} stackable>
        {!loading && data.users.map((user, i) => {
          return (
            <PostCell
              key={i}
              username={user.name}
              title={user.email}
              description=":))))))))"
            />
          );
        })}
      </Grid>
    </div>
  );
};

export default PostGrid;
