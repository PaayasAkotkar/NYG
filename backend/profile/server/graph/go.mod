module app/server/graph

go 1.25.3

require (
	app/server/graph/model v0.0.0-00010101000000-000000000000
	github.com/99designs/gqlgen v0.17.81
	github.com/vektah/gqlparser/v2 v2.5.30
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
)

replace app/profile/graph/model => ./model

replace app/sqlmanager => .././../../SQL-Manager

replace app/server/graph/model => ./model
