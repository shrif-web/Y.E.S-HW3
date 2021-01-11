import React from "react";
import { useQuery } from "@apollo/client";
import { Grid, Card, Icon, Header } from "semantic-ui-react";
import gql from "graphql-tag";

const GET_POSTS_QUERY = gql`
  {
    timeline(start: 0, amount: 1000) {
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

const PostCell = ({ title, content, created_by, isMobile }) => {
  return (
    <Grid.Column fluid="true" textAlign="left">
      <Card fluid>
        <Card.Content header={title} />
        <Card.Content description={content} />
        <Card.Content extra>
          <Icon name="user" />
          created by <b>{created_by.name}</b>
        </Card.Content>
      </Card>
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

  return (
    <div style={{ top: "50px", position: "absolute", width: "100%" }}>
      <Header as='h2' icon textAlign='center'>
      <Icon name='users' circular />
      <Header.Content>Timeline</Header.Content>
    </Header>
      <Grid columns={3} stackable style={{ margin: 20 }}>
        {!loading &&
          data.timeline.map((post, i) => {
            return (
              <PostCell
                key={i}
                title={post.title}
                content={post.content}
                created_by={post.created_by}
                isMobile={props.isMobile}
              />
            );
          })}
      </Grid>
    </div>
  );
};

export default PostGrid;
