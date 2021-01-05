import React from "react";
import { useQuery, useMutation } from "@apollo/client";
import gql from "graphql-tag";
import constants from "../constants";

const GET_USER_POSTS = gql`
  query Posts {
    user {
        name
    #   posts {
    #     id
    #     title
    #     content
    #     created_at
    #     created_by {
    #       name
    #     }
    #   }
    }
  }
`;

const UserPosts = props => {
  console.log(
    "===========",
    localStorage.getItem(constants.AUTH_TOKEN).slice(1, -1)
  );
  const { data, loading, error } = useQuery(GET_USER_POSTS);

  console.log("in dashboard");
  console.log("data: ", data);
  console.log("error:", error);
  console.log("loading:", loading);

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
