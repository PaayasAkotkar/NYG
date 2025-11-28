package server

import (
	"app/server/graph"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

func Server() {
	port := "1299"

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &NYGServices{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.GRAPHQL{})
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			WriteBufferSize: 1024,
			ReadBufferSize:  1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/nyg-chat-init", c.Handler(srv))
	http.Handle("/nyg-chat-latest-msg", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println(err)
		return
	}

}
