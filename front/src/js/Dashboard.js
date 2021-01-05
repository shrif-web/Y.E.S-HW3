import React from "react";
import { useQuery, useMutation } from "@apollo/client";
import gql from "graphql-tag";
import constants from "../constants";

const GET_USER_POSTS = gql`
  query Posts($username: String!) {
    postsOfUser(userName: $username) {
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

const UserPosts = props => {

  const { data, loading, error } = useQuery(GET_USER_POSTS, {
      variables: {
          username: localStorage.getItem(constants.AUTH_TOKEN)
      }
  });

  console.log("in dashboard")
  console.log("data: ", data)
  console.log("error:", error)
  console.log("loading:", loading)



  return <div>posts!</div>;
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
