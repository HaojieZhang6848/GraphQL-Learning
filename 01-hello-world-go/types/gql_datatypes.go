package types

import (
	"01-hello-world-go/handler"

	"github.com/graphql-go/graphql"
)

// GqlAuthorType 定义GraphQL的作者类型
var GqlAuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// GqlBookType 定义GraphQL的书籍类型
var GqlBookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.String},
		"title":  &graphql.Field{Type: graphql.String},
		"author": &graphql.Field{Type: graphql.String},
	},
})

// GqlEmployeeType 定义GraphQL的员工类型，这里带有子查询（salary字段）
var GqlEmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.String},
		"name":       &graphql.Field{Type: graphql.String},
		"age":        &graphql.Field{Type: graphql.Int},
		"department": &graphql.Field{Type: graphql.String},
		"gender":     &graphql.Field{Type: graphql.String},
		"salary": &graphql.Field{
			Type: graphql.Float,
			Args: graphql.FieldConfigArgument{
				"city":    &graphql.ArgumentConfig{Type: graphql.String},
				"taxRate": &graphql.ArgumentConfig{Type: graphql.Float},
			},
			Resolve: handler.ResolveSalary,
		},
	},
})

// GqlAccountType 定义GraphQL的账户类型
var GqlAccountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Account",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.String},
		"owner":  &graphql.Field{Type: graphql.String},
		"amount": &graphql.Field{Type: graphql.Float},
		"status": &graphql.Field{Type: graphql.String},
	},
})

// GQLCreateAccountInputType 定义创建账户的输入类型
var GQLCreateAccountInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateAccountInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"owner":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"amount": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
	},
})

// GQLUpdateAccountInputType 定义更新账户的输入类型
var GQLUpdateAccountInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateAccountInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"owner":  &graphql.InputObjectFieldConfig{Type: graphql.String},
		"amount": &graphql.InputObjectFieldConfig{Type: graphql.Float},
		"status": &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})
