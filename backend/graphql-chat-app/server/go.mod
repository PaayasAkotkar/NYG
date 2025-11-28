module app/server

go 1.25.3

require (
	app/server/graph v0.0.0-00010101000000-000000000000
	app/server/graph/model v0.0.0-00010101000000-000000000000
	app/sqlmanager v0.0.0-00010101000000-000000000000
	github.com/99designs/gqlgen v0.17.81
	github.com/go-sql-driver/mysql v1.9.3
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.0
	github.com/rs/cors v1.11.1
	github.com/vektah/gqlparser/v2 v2.5.31
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
)

replace app/server/graph/model => ./graph/model

replace app/server/graph => ./graph

replace app/sqlmanager => ./../../SQL-Manager
