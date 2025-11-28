module app/server

go 1.25.3

require (
	app/server/graph v0.0.0-00010101000000-000000000000
	app/server/graph/model v0.0.0-00010101000000-000000000000
	app/sqlmanager v0.0.0-00010101000000-000000000000
	github.com/99designs/gqlgen v0.17.81
	github.com/go-mysql-org/go-mysql v1.13.0
	github.com/go-sql-driver/mysql v1.9.3
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.3
	github.com/graphql-go/graphql v0.8.1
	github.com/graphql-go/handler v0.2.4
	github.com/rs/cors v1.11.1
	github.com/vektah/gqlparser/v2 v2.5.30
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/pingcap/errors v0.11.5-0.20250318082626-8f80e5cb09ec // indirect
	github.com/pingcap/failpoint v0.0.0-20240528011301-b51a646c7c86 // indirect
	github.com/pingcap/log v1.1.1-0.20241212030209-7e3ff8601a2a // indirect
	github.com/pingcap/tidb/pkg/parser v0.0.0-20250421232622-526b2c79173d // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)

replace app/server/graph => ./graph

replace app/server/graph/mode => ./graph/model

replace app/server/graph/model => ./graph/model

replace app/sqlmanager => ../.././SQL-Manager
