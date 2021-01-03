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
        # posts {
        #   id
        # }
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
      <>salam</>
      <Button onClick={() => {
        createPost()
      }}>Create Post</Button>
      {/* <Grid columns={3} stackable>
        {posts.map((post, i) => {
          return (
            // <PostCell
            //   key={i}
            //   username={post.username}
            //   title={post.title}
            //   description={post.description}
            // />
          );
        })}
      </Grid> */}
      {/* {this.state.isMobile && <Header fixed="bottom" />} */}
    </div>
  );
};

export default PostGrid;
