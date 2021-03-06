import React, { useState } from "react";
import { useQuery, useMutation } from "@apollo/client";
import gql from "graphql-tag";
import _ from "lodash";
import constants from "../constants.js";
import {
  Card,
  Grid,
  Button,
  Segment,
  Header,
  Input,
  TextArea,
  Form,
  Modal,
  List,
  Sidebar,
  Menu,
  Icon
} from "semantic-ui-react";
import { useHistory } from "react-router-dom";

const GET_USER_POSTS = gql`
  {
    user {
      name
      email
      posts {
        id
        title
        content
        created_by {
          name
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
        content
        created_by {
          name
        }
      }
      ... on Exception {
        message
      }
    }
  }
`;

const DELETE_POST_MUTATION = gql`
  mutation DeletePost($targetID: String!, $username: String!) {
    deletePost(targetID: $targetID, username: $username) {
      __typename
      ... on Post {
        id
        title
        content
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
    $username: String
  ) {
    updatePost(
      targetID: $targetID
      input: { title: $newTitle, content: $newContent }
      username: $username
    ) {
      __typename
      ... on Post {
        id
        content
        title
      }
      ... on Exception {
        message
      }
    }
  }
`;

const AddPostModal = ({ addingPost, setAddingPost }) => {
  const [state, setState] = useState({
    // addingPost: false,
    newTitle: "",
    newContent: ""
  });

  const history = useHistory();

  function logout() {
    localStorage.removeItem(constants.AUTH_TOKEN);
    history.push("/");
    window.location.reload(false);
  }

  const [createPost] = useMutation(CREATE_POST_MUTATION, {
    onError(err) {
      if (err.message.includes("401")) {
        logout();
      }
    },
    update(cache, { data: { createPost } }) {
      const data = cache.readQuery({
        query: GET_USER_POSTS
      });

      const localData = _.cloneDeep(data);

      localData.user.posts = [...localData.user.posts, createPost];

      cache.writeQuery({
        query: GET_USER_POSTS,
        data: {
          ...localData
        }
      });
    },
    onCompleted({ createPost }) {
      console.log("created post succesfully :D", createPost);
      setAddingPost(false);
    }
  });

  return (
    <div>
      <Modal open={addingPost} dimmer="blurring" fluid="true">
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
          <Button onClick={() => setAddingPost(false)}>Cancel</Button>
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
    </div>
  );
};

const PostCell = ({ post, rerenderComponent }) => {
  const [state, setState] = useState({
    editingActive: false,
    newTitle: post.title,
    newContent: post.content
  });

  const history = useHistory()

  function logout() {
    localStorage.removeItem(constants.AUTH_TOKEN);
    history.push("/");
    window.location.reload(false);
  }

  const [updatePost] = useMutation(UPDATE_POST_MUTATION, {
    onError(err) {
      if (err.message.includes("401")) {
        logout();
      }
    },
    update(cache, { data: { updatePost } }) {

      const data = cache.readQuery({
        query: GET_USER_POSTS
      });

      const localData = _.cloneDeep(data);
      localData.user.posts = localData.user.posts.map(post => {
        return post.id === updatePost.id ? updatePost : post;
      });

      cache.writeQuery({
        query: GET_USER_POSTS,
        data: {
          ...localData
        }
      });
    },
    onCompleted: ({ updatePost }) => {
      handleHide();
    }
  });

  const [deletePost] = useMutation(DELETE_POST_MUTATION, {
    onError(err) {
      if (err.message.includes("401")) {
        logout();
      }
    },
    update(cache, { data: { deletePost } }) {

      const data = cache.readQuery({
        query: GET_USER_POSTS
      });

      const localData = _.cloneDeep(data);
      localData.user.posts = localData.user.posts.filter(function(post) {
        return post.id !== deletePost.id;
      });

      cache.writeQuery({
        query: GET_USER_POSTS,
        data: {
          ...localData
        }
      });
    },
    onCompleted({ message }) {
      // triggerState();
    }
  });

  const handleShow = () => setState({ ...state, editingActive: true });
  const handleHide = () => setState({ ...state, editingActive: false });

  return (
    <div>
      <Segment>
        <Card fluid color="blue">
          <Card.Content header={post.title} />
          <Card.Content description={post.content} />
        </Card>

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
              deletePost({
                variables: { targetID: post.id, username: post.created_by.name }
              });
            }}
          />
        </Button.Group>
      </Segment>

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
                  newContent: state.newContent,
                  username: post.created_by.name
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
    trigger: false,
    addingPost: false
  });

  const { data, loading, error } = useQuery(GET_USER_POSTS);

  function setAddingPost(value) {
    setState({ addingPost: value });
  }

  return (
    <div>
      {!loading && (
        <Sidebar
          as={Menu}
          animation="overlay"
          icon="labeled"
          direction="left"
          vertical
          visible={props.isMobile ? props.sidebarIsOpen : true}
          width="thin"
          style={{ width: 250, top: 40 }}
        >
          <Menu.Item as="a">
            <Icon name="user" />
            {data.user.name}
          </Menu.Item>
          <Menu.Item as="a">
            <Icon name="mail" />
            {data.user.email}
          </Menu.Item>
          <Menu.Item as="a">
            <Icon name="question circle" />
            You have {data.user.posts.length} post
            {data.user.posts.length == 1 ? "" : "s"}!
          </Menu.Item>
          <Menu.Item as="a">
            <Button
              positive
              onClick={() => {
                setState({ addingPost: true });
              }}
            >
              Add New Post
            </Button>
          </Menu.Item>
        </Sidebar>
      )}
      <AddPostModal
        addingPost={state.addingPost}
        setAddingPost={setAddingPost}
      />
      <Grid.Row>
        <Grid
          columns={2}
          stackable
          style={{
            position: "absolute",
            left: props.isMobile ? 0 : 250,
            right: 0,
            margin: 30
          }}
        >
          {!loading &&
            data.user.posts.map(post => {
              return (
                <Grid.Column textAlign="left" key={post.id}>
                  <PostCell
                    post={post}
                    rerenderComponent={props.rerenderComponent}
                  />
                </Grid.Column>
              );
            })}
        </Grid>
      </Grid.Row>
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
        <UserPosts
          isMobile={this.props.isMobile}
          sidebarIsOpen={this.props.sidebarIsOpen}
        />
      </div>
    );
  }
}

export default Dashboard;
