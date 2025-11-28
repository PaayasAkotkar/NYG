package server

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
)

// RegisterRoom registers the room and returns the current len of the room
func RegisterRoom(h *Hub, conn *websocket.Conn, RoomID string, ID string, isClash bool) int {
	h.raw.Lock()
	defer h.raw.Unlock()
	roomclient, ok := h.clients[conn]
	if roomclient.isClosing {
		return -1
	}
	if !ok {
		return -1
	}
	roomclient.roomID = RoomID
	if _, ok := h.rooms[RoomID]; !ok {
		h.rooms[RoomID] = make(map[*websocket.Conn]*roomClient)
	}
	h.rooms[RoomID][conn] = roomclient
	h.clients[conn].clientID = conn.Params("id")

	if !isClash {
		go func() {
			// send the host ID that has created the room
			// Host = h.clients[conn].clientID
			Conns := strconv.Itoa(len(h.rooms[RoomID]))

			// stores the room with the client ID
			store := make(map[string]string)
			store[RoomID] = ID
			cRommstore.rooms <- store

			h.broadcast <- BroadcastReq{RoomID: RoomID, Token: "ActiveConns: " + Conns}

		}()
	} else {
		log.Println("created for clash")
		go func() {
			h.broadcast <- BroadcastReq{RoomID: RoomID, Token: "Clash: true"}
		}()
	}
	fmt.Println("len: ", len(h.rooms[RoomID]))
	// joinDone := "JoinDone: true"
	// BroadcastToClient(h, joinDone)
	// h.broadcast <- BroadcastReq{Token: "JoinDone: true"} // to send the player to the party chat
	return len(h.rooms[RoomID])
}

func UnregisterRoom(h *Hub, conn *websocket.Conn) {
	log.Println("unergistered process begins")
	client, ok := h.clients[conn]
	if !ok {
		return
	}

	RoomID := client.roomID
	h.raw.Lock()
	if room, ok := h.rooms[RoomID]; ok {
		log.Println("deleting room and its connection")
		// will remove the current connection from the room
		delete(room, conn) // removes the connection from the room
		if len(room) == 0 {
			log.Println("deleting rooms and rooms ID")
			var lg LocfiyGameInfo
			lg.Session = false
			_c, err := json.Marshal(lg)
			if err != nil {
				log.Println(err)
				return
			}
			token := "LocifyGame: " + string(_c)

			h.broadcast <- BroadcastReq{RoomID: RoomID, Token: token}

			// remove room if and only if there are no members left in the room
			UnregisterStoreRoom(h, conn, RoomID)

			delete(h.rooms, RoomID) // empties the room

		} else {
			// send the active conns
			// if there are still members left
			Conns := strconv.Itoa(len(h.rooms[RoomID]))

			log.Println("active: ", Conns)
			h.broadcast <- BroadcastReq{RoomID: RoomID, Token: "ActiveConns: " + Conns}
		}

		log.Println("unregisterd successfuly")
		log.Println("rooms: ", h.rooms)
	} else {
		log.Println("room not ok")
	}
	h.raw.Unlock()
	delete(h.clients, conn) // removes the connection
	// conn.Close()
	log.Println("unregistered successful")
}

func UnregisterRoomSession(h *Hub, conn *websocket.Conn, roomname, id string) {

	log.Println("unergistered process begins")
	h.raw.Lock()
	defer h.raw.Unlock()
	if _, ok := h.clients[conn]; !ok {
		return
	}
	room, exists := h.rooms[roomname]
	if !exists {
		delete(h.clients, conn)
		return
	}
	info, inroom := room[conn]
	if !inroom {
		delete(h.clients, conn)
		return
	}
	owner := getLocifyRoomSettings[roomname].RoomOwnerID
	cur := info.clientID
	if owner == cur {
		for c := range room {
			delete(h.clients, c)
		}
		delete(h.rooms, roomname)
		return
	}

	delete(room, conn)
	delete(h.clients, conn)

	Conns := strconv.Itoa(len(h.rooms[roomname]))
	var lg LocfiyGameInfo
	lg.Session = false
	_c, err := json.Marshal(lg)
	if err != nil {
		log.Println(err)
		return
	}
	token := "LocifyGame: " + string(_c)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, roomname: roomname, to: id, _sleep: false}
	log.Println("active: ", Conns)
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}

	log.Println("unregistered successful")
}

func UnregisterClientFromRoom(h *Hub, roomID string, clientID string) {
	h.raw.Lock()
	defer h.raw.Unlock()
	log.Println("for roomID: ", roomID, "clientID: ", clientID)
	log.Println("unergistering client from room")
	if room, ok := h.rooms[roomID]; ok {
		for conn, d := range room {
			if d.clientID == clientID {
				delete(room, conn)
				log.Println("removeing clinet wth id: ", clientID)
			}
		}
	} else {
		log.Println("room not ok")
	}
}

func UnregisterStoreRoom(h *Hub, conn *websocket.Conn, RoomID string) {
	log.Println("deleting created rooms")
	delete(getLocifyRoomSettings, RoomID)

	for id := range getLocifyProfile {
		delete(getLocifyProfile[id], RoomID)
	}
	for id := range getClashProfile {
		delete(getClashProfile[id], RoomID)
	}

	cRoom[RoomID] = nil
	delete(cRoom, RoomID)

	if _, ok := saveShuffle[RoomID]; ok {
		delete(saveShuffle, RoomID)
		delete(saveShuffle, RoomID)
	}

}
