package handler

import (
	"github.com/graphql-go/graphql"
)

var fibonacci = func(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

var ResolveFibonacci = func(p graphql.ResolveParams) (interface{}, error) {
	n := p.Args["n"].(int)
	return fibonacci(n), nil
}
