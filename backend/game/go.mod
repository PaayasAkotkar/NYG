module nyg/app

go 1.25.3

replace nyg/dataset => ./dataset

replace nyg/server => ./server

require nyg/server v0.0.0-00010101000000-000000000000

require nyg/list v0.0.0-00010101000000-000000000000 // indirect

require (
	app/nygpostprotoc v0.0.0-00010101000000-000000000000 // indirect
	app/nygprotoc v0.0.0-00010101000000-000000000000 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/fasthttp/websocket v1.5.8 // indirect
	github.com/gofiber/contrib/websocket v1.3.4 // indirect
	github.com/gofiber/fiber/v2 v2.52.6 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/savsgio/gotils v0.0.0-20240303185622-093b76447511 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.58.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/grpc v1.76.0 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
	nyg/dataset v0.0.0-00010101000000-000000000000 // indirect
	nyg/deck v0.0.0-00010101000000-000000000000 // indirect
	nyg/dictionary v0.0.0-00010101000000-000000000000 // indirect
	nyg/profiles v0.0.0-00010101000000-000000000000 // indirect
	nyg/validate v0.0.0-00010101000000-000000000000 // indirect
	resty.dev/v3 v3.0.0-beta.3 // indirect
)

replace nyg/books => ./books

replace nyg/dictionary => ./dictionary

replace nyg/list => ./lists

replace nyg/validate => ./validate

replace nyg/deck => ./decks

replace nyg/profiles => ./profiles

replace app/nygprotoc => ./../protos/result-protoc

replace app/nygpostprotoc => ./../protos/post-protoc
