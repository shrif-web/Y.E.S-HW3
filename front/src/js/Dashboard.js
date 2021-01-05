import React from "react";
import { useQuery, useMutation } from "@apollo/client";
import gql from "graphql-tag";

const GET_USER_POSTS = gql`
  query Posts($username: String!) {
    postsOfUser(username: $username) {
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

//   const { data, loading, error } = useQuery(GET_USER_POSTS, {
//       variables: {
//           username: 
//       }
//   });



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
