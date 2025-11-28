package server

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomUnregister(t *testing.T) {

	type roomClient struct {
		mu        sync.Mutex
		isClosing bool
		roomID    string
		clientID  string
	}
	type ClientIDUnregistration struct {
		Roomname string
		ID       string
	}

	roomname := "x"
	// consider conn as websocket.conn
	conn := "cin"
	clientIDToRemove := "kk"

	removePattern := ClientIDUnregistration{}
	removePattern.Roomname = roomname
	removePattern.ID = clientIDToRemove

	client := map[string]map[string]*roomClient{}
	client[roomname] = map[string]*roomClient{
		conn: {
			clientID: clientIDToRemove,
		},
	}
	if _, ok := client[roomname]; ok {
		client[roomname] = map[string]*roomClient{}
	}
	if _, ok := client[roomname][conn]; !ok {

		client[roomname] = map[string]*roomClient{
			conn: {
				clientID: "aasa",
			},
		}
	}

	if _, ok := client[removePattern.Roomname]; ok {
		for conn, r := range client[removePattern.Roomname] {
			if removePattern.ID == r.clientID {
				delete(client, conn)
			}
		}
	}

	assert.Equal(t, "aasa", client[roomname][conn].clientID)
	assert.NotEqual(t, clientIDToRemove, client[roomname][conn].clientID)

}
