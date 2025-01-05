package types

import (
	"01-hello-world-go/handler"

	"github.com/graphql-go/graphql"
)

// Name为Mutation的GraphQL对象类型是所有GraphQL变更的入口点，变更内可以嵌套子变更
var GQLMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createAccount": &graphql.Field{
			Type: GqlAccountType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{Type: GQLCreateAccountInputType},
			},
			Resolve: handler.ResolveCreateAccount,
		},
		"updateAccount": &graphql.Field{
			Type: GqlAccountType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{Type: GQLUpdateAccountInputType},
			},
			Resolve: handler.ResolveUpdateAccount,
		},
	},
})