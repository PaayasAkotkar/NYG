package server

import (
	"sync"
)


type RoomClient struct {
	mu    sync.Mutex
	rooms chan map[string]string
	done  chan bool
}

type TStoreShuffle struct {
	store chan map[string]map[string][]string // room-name, team-name and team's id
	mu    sync.Mutex
	done  chan bool
}

var (
	ClashIdsDetails = make(chan map[string][]string)

	// stores the name of the room and client id's of the associated room
	cRommstore = RoomClient{rooms: make(chan map[string]string)}
	cRoom      = make(map[string][]string)

	storeShuffle = TStoreShuffle{store: make(chan map[string]map[string][]string), done: make(chan bool)}
	saveShuffle  = make(map[string]map[string][]string) // room-name + team name and players

)
