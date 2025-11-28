package server

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
)

func SignedUp(h *Hub, conn *websocket.Conn, l Lobby, token []byte) {
	err := json.Unmarshal(token, &l)
	if err != nil {
		log.Println("err in lobby")
		log.Println(err)
		return
	}
	switch true {
	case l.Clash:
		log.Println("clash")

		ClashIdsStoreRoom(h, conn)
		done, roomname := CreateClashRoom(h, conn, l)

		log.Println("create: ", done)

		if done {
			ClashTags(h, conn, roomname)
		}
	case l.JoinRoom:

		log.Println("room joinings")

		start, roomname := RoomJoin(l, h, conn)
		// Monitor(h, l, token)
		onevone := l.RoomCapacity == 2
		BoardcastProfile(h, l.ID, roomname, onevone)

		l.CreateRoom = false
		l.JoinRoom = false

		if start {

			// imp to append the id cause the count is always +1 during validation
			x := cRoom[roomname]

			if len(x) != getLocifyRoomSettings[roomname].Capacity {
				x = append(x, l.ID)
			}

			LocifyTags(h, l, roomname, x)

		} else {
			log.Println("team has not created yet waiting for players")
		}
		fmt.Println("signup done")

	case l.CreateRoom:

		log.Println("creating room")

		RoomCreate(l, h, conn)
		onevone := l.RoomCapacity == 2
		BoardcastProfile(h, l.ID, l.RoomName, onevone)

	default:
		log.Println("none found")
	}
}
