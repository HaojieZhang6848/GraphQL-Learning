const express = require('express');
const { buildSchema } = require('graphql');
const { graphqlHTTP } = require('express-graphql');

const schema = buildSchema(`
    type Student {
        id: Int!
        name: String!
        age: Int!
        email: String!
    }

    type Query {
        hello: String
        students: [Student]
    }
`);

const root = {
    hello: () => {
        return 'Hello world!';
    },
    students: () => {
        return [
            { id: 1, name: 'John Doe', age: 23, email: 'zz@163.com' },
            { id: 2, name: 'Jane Doe', age: 22, email: 'rr@163.co,' },
        ];
    }
};

const app = express();
app.use('/graphql', graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true,
}));

app.listen(4000, () => {
    console.log('Running a GraphQL API server at http://localhost:4000/graphql');
});