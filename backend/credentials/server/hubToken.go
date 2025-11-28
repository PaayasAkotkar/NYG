package server

import "sync"

type Hub struct {
	Coin     chan map[Coin]bool
	Spur     chan map[Spur]bool
	Upgarade chan map[Upgarade]bool
	mu       sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Coin:     make(chan map[Coin]bool),
		Spur:     make(chan map[Spur]bool),
		Upgarade: make(chan map[Upgarade]bool),
	}
}
