package server

import (
	"app/server/graph"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"

	"github.com/gorilla/websocket"
	"github.com/vektah/gqlparser/v2/ast"
)

type NYGResolver struct{}

type Services struct{ *NYGResolver }
type QServices struct{ *NYGResolver }

func (r *NYGResolver) Subscription() graph.SubscriptionResolver { return &Services{r} }
func (r *NYGResolver) Query() graph.QueryResolver               { return &QServices{r} }

func Messenger() {

	port := "6060"
	// imp for query
	c := cors.New(cors.Options{
		AllowedOrigins:   allowOrgs,
		AllowCredentials: true,
	})
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &NYGResolver{}}))

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
	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.Options{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/default", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/settings-changes", srv)
	http.Handle("/login", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			panic(err)
		}
	}()
	wg.Wait()
}

var wg sync.WaitGroup
var mu sync.Mutex

func Server() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		Watch()
	}()

	go func() {
		defer wg.Done()
		Messenger()
	}()

	wg.Wait()

}
