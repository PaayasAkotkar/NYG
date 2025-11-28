package server

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Book struct {
	Entertainment []string `json:"entertainment"`
	Sports        []string `json:"sports"`
}

func InitBooks() Book {
	var b Book
	b.Entertainment = []string{"MOVIES", "MUSIC", "SHOWS"}
	b.Sports = []string{"BASKETBALL", "CRICKET", "FOOTBALL"}
	return b
}
func Books() *handler.Handler {
	var b Book
	b.Entertainment = []string{"MOVIES", "MUSIC", "SHOWS", "BOOKS"}
	b.Sports = []string{"BASEBALL", "BASKETBALL", "CRICKET", "FOOTBALL", "HOCKEY", "RUGBY"}
	_sportBook := graphql.NewObject(graphql.ObjectConfig{
		Name: "Sports", // type name
		Fields: graphql.Fields{
			"book": {
				Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return b.Sports, nil
				},
			},
		},
	})
	_entertainment := graphql.NewObject(graphql.ObjectConfig{
		Name: "Entertainment",
		Fields: graphql.Fields{
			"book": {
				Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return b.Entertainment, nil
				},
			},
		},
	})
	_root := graphql.NewObject(graphql.ObjectConfig{
		Name: "updatedBooks",
		Fields: graphql.Fields{
			"Entertainment": {
				Type: _entertainment,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return b.Entertainment, nil
				},
			},
			"Sports": {
				Type: _sportBook,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return b.Sports, nil
				},
			},
		},
	})
	schemaConfig := graphql.SchemaConfig{Query: _root}
	schama, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Println(err)
	}
	h := handler.New(&handler.Config{
		GraphiQL:   true,
		Pretty:     true,
		Schema:     &schama,
		Playground: true,
	})
	return h
}
