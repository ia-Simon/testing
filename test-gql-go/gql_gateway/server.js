const { ApolloServer } = require('apollo-server');
const { ApolloGateway, IntrospectAndCompose } = require("@apollo/gateway");

const subgraphs = process.env.SUBGRAPH_LIST
  .split(",")
  .map(subgraphPair => ({ 
    name: subgraphPair.split("@")[0],
    url: subgraphPair.split("@")[1],
  }));
console.log(`Listening to subgraphs: ${JSON.stringify(subgraphs, null, 2)}`);

const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({ subgraphs }),
});

const server = new ApolloServer({ gateway });

server.listen({ port: process.env.PORT }).then(({ url }) => {
  console.log(`Gateway listening at ${url}`);
});