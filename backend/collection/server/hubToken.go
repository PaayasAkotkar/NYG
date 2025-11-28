package server

import (
	"sync"
)

type Hub struct {
	key           string                          // room-id
	trigger       chan map[string]map[string]bool // key : send item key: true|false
	done          chan bool
	dictionaryKey string
	listKey       string
	bookKey       string
	mu            sync.Mutex
	wg            sync.WaitGroup
	once          sync.Once
}

func NewHub() *Hub {
	return &Hub{
		trigger: make(chan map[string]map[string]bool),
		done:    make(chan bool),
	}
}

type Monitor struct {
	going chan map[string]map[string]bool
	done  chan bool
}

var (
	monitor     Monitor
	viewMonitor = make(map[string]map[string]bool)
)
