package server

import (
	"sync"

	"github.com/tmaxmax/go-sse"
)

// BroadcastReq note: token must be struct nothing else
type BroadcastReq struct {
	Token any
	ID    string
	Conn  *sse.Server
}
type Hub struct {
	Broadcast chan BroadcastReq
	mu        sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Broadcast: make(chan BroadcastReq),
	}
}
