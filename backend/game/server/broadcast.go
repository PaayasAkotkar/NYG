package server

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

// BroadcastToRoom  broadcast to all the clients that are in the room
func BroadcastToRoom(h *Hub, roomID string, token string) {
	h.raw.RLock()
	defer h.raw.RUnlock()
	log.Println("boradcast to room: ", roomID)
	if _, ok := h.rooms[roomID]; ok {
		for conn, c := range h.rooms[roomID] {
			copyConn := conn
			copyClient := c
			h.wg.Add(1)
			go func(conn *websocket.Conn, c *roomClient) {

				defer h.wg.Done()

				c.mu.Lock()
				defer c.mu.Unlock()
				if c.isClosing {
					log.Println("closing")
					return
				}
				log.Println("broadcasting to room: ", roomID, "the token: ", string(token))
				err := conn.WriteMessage(websocket.TextMessage, []byte(token))
				if err != nil {
					log.Println(err)
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					conn.Close()
					h.unregister <- conn
				}
				log.Println("done broadcasting")
			}(copyConn, copyClient)
		}
	} else {
		log.Println("room not found")
	}

}

// BroadcastAll @RETURN broadcast to all the clients; doesnt matter if they are in room or not

func BroadcastAll(h *Hub, token string) {
	for connection, c := range h.clients {
		go func(conn *websocket.Conn, c *roomClient) {
			c.mu.Lock()
			defer c.mu.Unlock()
			if c.isClosing {
				return
			}
			err := conn.WriteMessage(websocket.TextMessage, []byte(token))
			if err != nil {
				// log.Println("err == nil broadcast")
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				conn.Close()
				h.unregister <- conn
			}
		}(connection, c)
	}
}

//	GameRoomBroadcast returns  broadcast to specific client that is in the playing hall
//
// note: not in the game hall
// _sleep: to broadcast to the player that has not won the point
func GameRoomBroadcast(h *Hub, roomID string, token string, to string, _sleep bool) {
	// @TODO: make sure that the client is in the room
	// if not do not broadcast to that client
	h.raw.Lock()
	defer h.raw.Unlock()
	log.Println("game room broadcast: ", roomID)

	for connection, c := range h.clients {
		copyConn := connection
		copyClient := c

		h.wg.Add(1)
		go func(conn *websocket.Conn, c *roomClient) {

			c.mu.Lock()
			defer c.mu.Unlock()
			defer h.wg.Done()

			if c.isClosing {
				return
			}

			ID := h.clients[conn].clientID
			log.Println("broadcast to: ", ID, "token: ", token)
			// if _sleep {
			// 	if ID != to {
			// 		err := conn.WriteMessage(websocket.TextMessage, []byte(token))
			// 		if err != nil {
			// 			// log.Println("err == nil broadcast")
			// 			conn.WriteMessage(websocket.CloseMessage, []byte{})
			// 			conn.Close()
			// 			h.unregister <- conn
			// 		}
			// 	}
			// } else

			if ID == to && token != "" {
				log.Println("token: ", token)

				err := conn.WriteMessage(websocket.TextMessage, []byte(token))
				if err != nil {
					// log.Println("err == nil broadcast")
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					conn.Close()
					h.unregister <- conn
				}
			} else {
				log.Println("empty token")
				log.Println("token: ", token)
			}

		}(copyConn, copyClient)
	}

}

func GameBroadcastToRoom(h *Hub, roomID string, token string, to string) {

	h.raw.RLock()
	defer h.raw.RUnlock()
	log.Println("game to room board: ", roomID)

	if _, ok := h.rooms[roomID]; ok {
		for conn, c := range h.rooms[roomID] {
			copyConn := conn
			copyClient := c
			h.wg.Add(1)
			go func(conn *websocket.Conn, c *roomClient) {
				c.mu.Lock()
				defer c.mu.Unlock()
				defer h.wg.Done()
				if c.isClosing {
					// log.Println("closing")
					return
				}
				ID := h.clients[conn].clientID

				if ID == to && token != "" {
					err := conn.WriteMessage(websocket.TextMessage, []byte(token))
					if err != nil {
						// log.Println("error")
						conn.WriteMessage(websocket.CloseMessage, []byte{})
						conn.Close()
						h.unregister <- conn
					}
					log.Println("game token:", string(token), "token ID: ", ID)
					log.Println("done game broadcasting")
				}
			}(copyConn, copyClient)
		}
	} else {
		log.Println("game room not found")
	}

}
