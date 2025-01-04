package types

import (
	"01-hello-world-go/handler"
	"github.com/graphql-go/graphql"
)

// Name为Query的GraphQL对象类型是所有GraphQL查询的入口点，查询内可以嵌套子查询（见Employee的例子）
// 查询的入口点上定义了authors, books, employees, fibonacci四个查询字段
var GQLQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"authors": &graphql.Field{
			Type:    graphql.NewList(GqlAuthorType),
			Resolve: handler.ResolveAuthors,
		},
		"books": &graphql.Field{
			Type:    graphql.NewList(GqlBookType),
			Resolve: handler.ResolveBooks,
		},
		"employees": &graphql.Field{
			Type: graphql.NewList(GqlEmployeeType),
			Args: graphql.FieldConfigArgument{
				"department": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: handler.ResolveEmployees,
		},
		"fibonacci": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"n": &graphql.ArgumentConfig{Type: graphql.Int}, // 定义参数n, 类型为Int
			},
			Resolve: handler.ResolveFibonacci,
		},
	},
})
