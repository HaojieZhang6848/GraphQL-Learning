package handler

import (
	"01-hello-world-go/db"

	"github.com/graphql-go/graphql"
)

var ResolveEmployees = func(p graphql.ResolveParams) (interface{}, error) {
	department := p.Args["department"].(string)
	return db.Employees[department], nil
}

var ResolveSalary = func(p graphql.ResolveParams) (interface{}, error) {
	city := p.Args["city"].(string)
	taxRate := p.Args["taxRate"].(float64)
	highSalaryCities := []string{"Beijing", "Shanghai", "Guangzhou", "Shenzhen"}
	lowSalaryCities := []string{"Chengdu", "Hangzhou", "Wuhan", "Nanjing"}

	baseSalary := 0.0
	if contains(highSalaryCities, city) {
		baseSalary = 20000.0
	} else if contains(lowSalaryCities, city) {
		baseSalary = 10000.0
	} else {
		baseSalary = 15000.0
	}
	return baseSalary * (1 - taxRate), nil
}

func contains[T comparable](arr []T, target T) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
