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
  Form,
  Modal,
  List,
  Label
} from "semantic-ui-react";
import { tupleExpression } from "@babel/types";

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
  mutation UpdatePost(
    $targetID: String!
    $newTitle: String!
    $newContent: String!
  ) {
    updatePost(
      targetID: $targetID
      input: { title: $newTitle, content: $newContent }
    ) {
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

const AddPostModal = () => {
  const [state, setState] = useState({
    addingPost: false,
    newTitle: "",
    newContent: ""
  });

  const [createPost] = useMutation(CREATE_POST_MUTATION, {
    onCompleted({ createPost }) {
      console.log("created post succesfully :D", createPost);
    }
  });

  return (
    <div>
      <Modal open={state.addingPost} dimmer="blurring">
        <Modal.Header>Create a new post :)</Modal.Header>
        <Modal.Content>
          <List>
            <List.Item>
              <Header>Title :</Header>
              <Input
                placeholder="Please write the title of your post."
                onChange={e => {
                  setState({ ...state, newTitle: e.target.value });
                }}
              />
            </List.Item>
            <List.Item>
              <Header>Content :</Header>
              <Form>
                <TextArea
                  placeholder="Please write the content of your post."
                  onChange={e => {
                    setState({ ...state, newContent: e.target.value });
                  }}
                />
              </Form>
            </List.Item>
          </List>
        </Modal.Content>
        <Modal.Actions>
          <Button onClick={() => setState({ addingPost: false })}>
            Cancel
          </Button>
          <Button
            onClick={() => {
              createPost({
                variables: {
                  title: state.newTitle,
                  content: state.newContent
                }
              });
            }}
            positive
          >
            Add
          </Button>
        </Modal.Actions>
      </Modal>
      <Button
        positive
        onClick={() => {
          setState({ addingPost: true });
        }}
      >
        Add Post
      </Button>
    </div>
  );
};

const PostCell = ({ post }) => {
  const [state, setState] = useState({
    editingActive: false,
    newTitle: post.title,
    newContent: post.content
  });

  const [updatePost] = useMutation(UPDATE_POST_MUTATION, {
    onCompleted: ({ updatePost }) => {
      handleHide();
    }
  });

  const [deletePost] = useMutation(DELETE_POST_MUTATION, {
    onCompleted({ message }) {
      // triggerState();
    }
  });

  const handleShow = () => setState({ editingActive: true });
  const handleHide = () => setState({ editingActive: false });

  return (
    <div>
      <Card
        key={post.id}
        id={post.id}
        header={post.title}
        meta={post.created_by.name}
        description={post.content}
        fluid
      />
      <Button.Group>
        <Button
          icon="edit"
          onClick={() => {
            handleShow();
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
      <Modal open={state.editingActive} dimmer="blurring">
        <Modal.Header>Updating your post ...</Modal.Header>
        <Modal.Content>
          <List>
            <List.Item>
              <Header>New Title :</Header>
              <Input
                placeholder="New Title"
                defaultValue={post.title}
                onChange={e => {
                  setState({ ...state, newTitle: e.target.value });
                }}
              />
            </List.Item>
            <List.Item>
              <Header>New Content :</Header>
              <Form>
                <TextArea
                  placeholder="New Content"
                  defaultValue={post.content}
                  onChange={e => {
                    setState({ ...state, newContent: e.target.value });
                  }}
                />
              </Form>
            </List.Item>
          </List>
        </Modal.Content>
        <Modal.Actions>
          <Button onClick={() => handleHide()}>Cancel</Button>
          <Button
            onClick={() => {
              updatePost({
                variables: {
                  targetID: post.id,
                  newTitle: state.newTitle,
                  newContent: state.newContent
                }
              });
            }}
            positive
          >
            Update
          </Button>
        </Modal.Actions>
      </Modal>
    </div>
  );
};

const UserPosts = props => {
  const [state, setState] = useState({
    trigger: false
  });

  const { data, loading, error } = useQuery(GET_USER_POSTS);

  function triggerState() {
    console.log("here trigerringgggggg //////////");
    setState({ trigger: !state.trigger });
  }

  return (
    <div>
      <AddPostModal />
      <Grid columns={2} stackable>
        {!loading &&
          data.user.posts.map(post => {
            return (
              <Grid.Column>
                <PostCell post={post} />
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
