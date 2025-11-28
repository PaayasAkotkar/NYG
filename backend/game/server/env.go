package server

import "flag"

var (
	namespaces = []string{"/ws/:id", "/ws/:id/:room"}
	addr       = flag.String("addr", "localhost:6070", "grpc address")
	allowOrgs  = []string{
		"*",
		"http://localhost:4200",
		"http://172.26.128.1:4200",
		"http://192.168.1.7:4200",
	}
)
