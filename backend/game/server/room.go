package server

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Room(namespace string, app *fiber.App, h *Hub) {
	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		// if c.Path() != "/ws" {
		// 	return c.Next()
		// }
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Get(namespace, websocket.New(func(c *websocket.Conn) {

		h.register <- c

		log.Println("registered successful")
		defer func() {
			h.unregister <- c // Remove the connection from your active list
			c.Close()
			// c.SetCloseHandler(func(code int, text string) error {
			// 	log.Println("WebSocket closed with code:", code, "reason:", text)

			// })
		}()
		// assign the value of the room parameter from the namespace: done

		id := c.Params("id")

		for {
			tokenType, token, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}
				return
			}

			if tokenType == websocket.TextMessage {
				log.Println("recieved message: ", string(token))

				var l Lobby
				var GR GameRoom
				var LO SignOut

				json.Unmarshal(token, &l)
				json.Unmarshal(token, &GR)
				json.Unmarshal(token, &LO)

				if LO.Logout {
					h.unregister <- c
				}

				switch true {

				case l.Clash && l.CreateRoom:
					SignedUp(h, c, l, token)

				case l.CreateRoom && !l.RoomCreated || l.JoinRoom && !l.RoomCreated:

					// sign up for lobby
					SignedUp(h, c, l, token)

				case l.RoomCreated && l.ChangeSettings:
					log.Println("changing the room settings")
					UpdateLocifyRoomSettings(h, l.RoomName, _StringSentinel_, l.Book, l.DecisionTime,
						l.GameTime, l.SetToss, l.Reverse, l.Starter, l.NexusPower, l.CovertPower, l.BetPower, l.FreezePower, l.DrawPower, l.TagPower, l.RewindPower)

				case GR.Session && !GR.Clash:

					log.Println("game room session")

					LocifyGame(h, c, GR, id)

				case GR.Session && GR.Clash:
					ClashGame(h, c, GR, id)

				default:
					for room, det := range getLocifyRoomSettings {
						r := []RoomList{}
						var t RoomList
						if det.Public {
							t.Book = det.Book
							t.Category = det.Category
							t.RoomName = room
							t.Type = det.Mode
							r = append(r, t)
							c, err := json.Marshal(&r)
							if err != nil {
								log.Println(err)
								return
							}
							h.broadcast <- BroadcastReq{Token: "RoomLists: " + string(c), RoomID: "private-room"}

						}
					}
				}
				log.Println("recieved token: ", string(token))

				// h.broadcast <- BroadcastReq{RoomID: RoomID, Token: string(token)}
			} else {

				log.Println("invalid message type")
			}
		}
	},
		websocket.Config{Origins: allowOrgs})) // very very important: prevents the site from fetching the data
}
