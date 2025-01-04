package handler

import (
	"01-hello-world-go/models"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/graphql-go/graphql"
)

var ResolveAuthors = func(p graphql.ResolveParams) (interface{}, error) {
	number := rand.Intn(10)
	authors := make([]interface{}, number)
	for i := 0; i < number; i++ {
		authors[i] = models.Author{
			ID:   gofakeit.UUID(),
			Name: gofakeit.Name(),
		}
	}
	return authors, nil
}
		