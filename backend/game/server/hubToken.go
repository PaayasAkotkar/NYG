package server

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type roomClient struct {
	mu        sync.Mutex
	isClosing bool
	roomID    string
	clientID  string
	conn      *websocket.Conn
}

type BroadcastReq struct {
	RoomID string
	Token  string
}

type ClientIDUnregistration struct {
	Roomname string
	ID       string
}

type Hub struct {
	clients               map[*websocket.Conn]*roomClient
	register              chan *websocket.Conn
	unregister            chan *websocket.Conn        // cleaner way if have conn
	unregisterID          chan ClientIDUnregistration // complex way if dont have conn and must be the id
	unregisterSession     chan *websocket.Conn
	unregisterRoomSession chan map[*websocket.Conn]SignOut
	broadcast             chan BroadcastReq // broadcasts to all the clients in the room
	Cbroadcast            chan BroadcastReq // broadcasts to the specific client that is in the room
	gameRoomBroadcast     chan reqGameRoomBroadcast
	rooms                 map[string]map[*websocket.Conn]*roomClient
	raw                   sync.RWMutex
	wg                    sync.WaitGroup
}

func NewHub() *Hub {
	return &Hub{
		clients:               make(map[*websocket.Conn]*roomClient),
		register:              make(chan *websocket.Conn),
		unregister:            make(chan *websocket.Conn),
		unregisterID:          make(chan ClientIDUnregistration),
		unregisterSession:     make(chan *websocket.Conn),
		unregisterRoomSession: make(chan map[*websocket.Conn]SignOut),
		broadcast:             make(chan BroadcastReq, 4),
		rooms:                 make(map[string]map[*websocket.Conn]*roomClient),
		raw:                   sync.RWMutex{},
		Cbroadcast:            make(chan BroadcastReq),
		gameRoomBroadcast:     make(chan reqGameRoomBroadcast, 4),
		wg:                    sync.WaitGroup{},
	}
}
