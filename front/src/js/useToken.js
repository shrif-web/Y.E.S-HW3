import React, { useState } from "react";
import constants from "../constants.js";

function useToken() {
  const getToken = () => {
    const tokenString = localStorage.getItem(constants.AUTH_TOKEN);
    // console.log("??????", tokenString)
    return tokenString
    // const userToken = JSON.parse(tokenString);
    // return userToken;
  };

  // console.log("in use token:)))))=============")

  const [token, setToken] = useState(getToken());

  const saveToken = userToken => {
    if (userToken == undefined) {
      localStorage.removeItem(constants.AUTH_TOKEN);
    } else {
      localStorage.setItem(constants.AUTH_TOKEN, JSON.stringify(userToken));
      setToken(userToken);
    }
  };

  return {
    setToken: saveToken,
    token
  };
}

export default useToken;
