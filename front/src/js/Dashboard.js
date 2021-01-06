import React, { useState } from "react";
import { useQuery, useMutation } from "@apollo/client";
import gql from "graphql-tag";
import constants from "../constants";
import {
  Card,
  Grid,
  Button,
  Dimmer,
  Segment,
  Image,
  Header,
  Input,
  TextArea,
  Form
} from "semantic-ui-react";
import { from } from "zen-observable";

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

const UPDATE_POST_MUTATION = gql`
  mutation UpdatePost($targetID: String!, $newTitle: String!, $newContent: String!) {
    updatePost(targetID: $targetID, input: {title: $newTitle, content: $newContent}) {
      __typename
    ... on OperationSuccessfull{
      message
    }
    ... on Exception{
      message
    }
    }
  }
`;

const PostCell = ({ post }) => {
  const [state, setState] = useState({
    editingActive: false,
    newTitle: post.title,
    newContent: post.content
  });

  const [updatePost] = useMutation(UPDATE_POST_MUTATION, {
    onCompleted: ({updatePost}) => {
      handleHide()
    }
  });

  const handleShow = () => setState({ editingActive: true });
  const handleHide = () => setState({ editingActive: false });

  return (
    <div>
      <Dimmer.Dimmable as={Segment} dimmed={state.editingActive}>
        <Card
          key={post.id}
          id={post.id}
          header={post.title}
          meta={post.created_by.name}
          description={post.content}
          fluid
          onClick={() => {
            handleShow();
          }}
        />

        <Dimmer
          active={state.editingActive}
          onClickOutside={handleHide}
        >
          <Form inverted>
            <Form.Group>
              <Form.Input
                label="New Title"
                placeholder="First name"
                defaultValue={post.title}
                onChange={e => {
                  setState({ ...state, newTitle: e.target.value });
                }}
              />
              <Form.TextArea
                label="New Content"
                placeholder="Last name"
                defaultValue={post.content}
                onChange={e => {
                  setState({ ...state, newContent: e.target.value });
                }}
              />
              <Form.Button
                onClick={() => {
                  updatePost({
                    variables: {
                      targetID: post.id,
                      newTitle: state.newTitle,
                      newContent: state.newContent
                    }
                  });
                }}
              >
                Update
              </Form.Button>
              <Form.Button onClick={handleHide}>Cancel</Form.Button>
            </Form.Group>
          </Form>
        </Dimmer>
      </Dimmer.Dimmable>
    </div>
  );
};

const UserPosts = props => {
  const [state, setState] = useState({
    trigger: false
  });

  const { data, loading, error } = useQuery(GET_USER_POSTS);

  const [createPost] = useMutation(CREATE_POST_MUTATION, {
    variables: {
      title: "title of the test post :)",
      content: "content of the test post !"
    },
    onCompleted({ data }) {
      triggerState();
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
      {/* <EditPostContent /> */}
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
                <PostCell post={post} />
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
