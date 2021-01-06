import React from "react";
import { useQuery, useMutation } from "@apollo/client";
import { Grid, Card, Button } from "semantic-ui-react";
import gql from "graphql-tag";
import constants from "../constants";
import { async } from "q";

const GET_POSTS_QUERY = gql`
  {
    posts(start: 0, amount: 1000) {
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

const PostCell = ({ title, content, created_by }) => {
  return (
    <Grid.Column>
      <Card header={title} meta={created_by.name} description={content} fluid />
    </Grid.Column>
  );
};

const PostGrid = props => {

  const { data, loading, error } = useQuery(GET_POSTS_QUERY);
  if (error) {
    return (
      <div>
        Error!!!!! data:
        {data}
      </div>
    );
  }
  console.log("data:", data);
  console.log("error:", error);
  console.log("loading:", loading);

  return (
    <div>
      <Grid columns={3} stackable>
        {!loading &&
          data.posts.map((post, i) => {
            return (
              <PostCell
                key={i}
                title={post.title}
                content={post.content}
                created_by={post.created_by}
              />
            );
          })}
      </Grid>
    </div>
  );
};

export default PostGrid;
