import React from "react";
import ReactDOM from "react-dom";
import "./styles/index.css";
import App from "./js/App";
import reportWebVitals from "./reportWebVitals";
import * as serviceWorker from "./serviceWorker";
import { BrowserRouter } from "react-router-dom";
import constants from './constants.js'
import {setContext} from '@apollo/client/link/context'

import { ApolloClient } from "apollo-client";
import {
  InMemoryCache,
  IntrospectionFragmentMatcher
} from "apollo-cache-inmemory";
import { HttpLink } from "apollo-link-http";
import { ApolloProvider } from "@apollo/client";

import introspectionQueryResultData from "../src/fragmentTypes.json";

const fragmentMatcher = new IntrospectionFragmentMatcher({
  introspectionQueryResultData
});

const cache = new InMemoryCache({ fragmentMatcher });
const link = new HttpLink({
  uri: "http://localhost:8080/query"
});

const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem(constants.AUTH_TOKEN);
  return {
    headers: {
      ...headers,
      Authorization: token ? `Bearer ${token}` : ''
    }
  };
});

const apolloClient = new ApolloClient({
  cache,
  link: authLink.concat(link)
});

ReactDOM.render(
  // <React.StrictMode>
  <BrowserRouter>
    <ApolloProvider client={apolloClient}>
      <App />
    </ApolloProvider>
  </BrowserRouter>,
  document.getElementById("root")

  // </React.StrictMode>,
);

serviceWorker.unregister();

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
