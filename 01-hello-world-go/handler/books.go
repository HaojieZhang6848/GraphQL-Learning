package handler

import (
	"01-hello-world-go/models"
	"math/rand"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/graphql-go/graphql"
)

var ResolveBooks = func(p graphql.ResolveParams) (interface{}, error) {
	number := rand.Intn(10)
	books := make([]interface{}, number)
	for i := 0; i < number; i++ {
		books[i] = models.Book{
			ID:     gofakeit.UUID(),
			Title:  gofakeit.Sentence(3),
			Author: gofakeit.Name(),
		}
	}
	return books, nil
}
