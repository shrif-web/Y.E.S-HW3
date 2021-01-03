import React from "react";
import ReactDOM from "react-dom";
import "./styles/index.css";
import App from "./js/App";
import reportWebVitals from "./reportWebVitals";
import * as serviceWorker from "./serviceWorker";

import { ApolloClient } from "apollo-client";
import { InMemoryCache } from "apollo-cache-inmemory";
import { HttpLink } from "apollo-link-http";
import { ApolloProvider } from "@apollo/client";

const cache = new InMemoryCache();
const link = new HttpLink({
  uri: "http://localhost:8080"
});

const apolloClient = new ApolloClient({
  cache,
  link
});

ReactDOM.render(
  // <React.StrictMode>
    <ApolloProvider client={apolloClient}>
      <App />
    </ApolloProvider>,
  // </React.StrictMode>,
  document.getElementById("root")
);

serviceWorker.unregister();

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
