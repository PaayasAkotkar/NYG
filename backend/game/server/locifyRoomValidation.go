package server

import (
	"encoding/json"
	"log"
	"nyg/profiles"
	"strconv"

	"github.com/gofiber/contrib/websocket"
)

type IRoomJoin struct {
	Name    string             `json:"name"`
	Profile profiles.IProfile_ `json:"profile"` // can be the profile that the room has created or joined
}

func RoomJoin(l Lobby, h *Hub, conn *websocket.Conn) (bool, string) {
	log.Println("joining to room")

	// check the if the room exists
	// exists, isFriendRoom, isPrivateRoom := false, getLocifyRoomSettings[l.RoomName].Friend, getLocifyRoomSettings[l.RoomName].Private
	exists := false
	startGame := getLocifyRoomSettings[l.To].Capacity == len(h.rooms[l.To])
	roomname := l.To
	id := l.ID

	// get the current client count of the room
	clientCount := len(h.rooms[roomname])

	// validation
	isFull := l.RoomCapacity <= clientCount

	// for friend room there are two ways
	// two players are joined and one will be joined via code
	_friendRoom, _publicRoom, _privateRoom := false, false, false

	for _, det := range getLocifyRoomSettings {

		if l.Code == det.Code || l.To == det.Name {
			log.Println("public: ", det.Public)
			roomname = det.Name
			_friendRoom = det.Friend
			_publicRoom = det.Public
			_privateRoom = det.Private
			clientCount = len(h.rooms[roomname])
			isFull = det.Capacity <= clientCount
			exists = true
			startGame = det.Capacity == clientCount // counting current client
			break
		}
	}

	switch true {

	case exists && _publicRoom:

		if !isFull {
			log.Println("public room joining")
			_l := RegisterRoom(h, conn, l.To, l.ID, false)

			h.broadcast <- BroadcastReq{Token: "ActiveConns: " + strconv.Itoa(_l)}
			BoardcastRoomJoin(h, id, roomname)
			startGame = getLocifyRoomSettings[l.To].Capacity == _l

		} else {
			log.Println("room full")
			h.broadcast <- BroadcastReq{Token: l.To + "RoomCap: room-full"}
		}

	case exists && _friendRoom:

		log.Println("friend room")

		switch true {
		// if the friend has not joined yet
		case l.Code == "":
			log.Println("joining without code")

			cap := getLocifyRoomSettings[roomname].Capacity - clientCount

			// keep one of the space for the friend
			if cap == 1 && !getLocifyRoomSettings[l.To].FriendJoin {
				log.Println("friend room full")

				// let the client join
			} else {

				_l := RegisterRoom(h, conn, roomname, l.ID, false)
				startGame = getLocifyRoomSettings[roomname].Capacity == _l

				// broadcast to all the clients
				h.broadcast <- BroadcastReq{Token: "ActiveConns: " + strconv.Itoa(_l), RoomID: l.To}
				BoardcastRoomJoin(h, id, roomname)
			}

		// friend joined via code
		case !getLocifyRoomSettings[roomname].FriendJoin:
			// todo: for invalid code
			log.Println("joining via code")

			_l := RegisterRoom(h, conn, roomname, l.ID, false)
			startGame = getLocifyRoomSettings[roomname].Capacity == _l

			var f FriendTrack
			f.ID = l.ID
			f.Join = true

			_join := map[string]FriendTrack{}

			_join[roomname] = f
			friendJoin <- _join

			BoardcastRoomJoin(h, id, roomname)

		}

	case exists && _privateRoom:
		if !isFull {

			log.Println("private room joining")

			if roomname != "" {

				_l := RegisterRoom(h, conn, roomname, l.ID, false)
				startGame = getLocifyRoomSettings[roomname].Capacity == _l
				BoardcastRoomJoin(h, id, roomname)
				h.broadcast <- BroadcastReq{Token: "ActiveConns: " + strconv.Itoa(_l), RoomID: l.To}

			} else {
				log.Println("empty room name")
			}

		} else {
			log.Println("room full")
			h.broadcast <- BroadcastReq{Token: roomname + "RoomCap: room-full", RoomID: "private-room"}
		}

	default:
		log.Println("cant confirm for joining")
		log.Println("game session is going on")
	}

	return startGame, roomname
}

func RoomCreate(l Lobby, h *Hub, conn *websocket.Conn) {
	// todo: the only difference between private room and public room is that
	// we wont be broadasting the name of the room else everything is same
	log.Println("making room")

	// check if the room is created
	exists := false
	roomMode := "" // either 2v2 or 1v1
	_publicRoom, _privateRoom, _friendRoom := !l.PrivateRoom && !l.Friend, l.PrivateRoom, l.Friend

	if l.RoomCapacity == 2 {
		roomMode = "1v1"
	} else {
		roomMode = "2v2"
	}
	log.Println("book: ", l.Book)

	switch true {
	case l.CreateRoom && getLocifyRoomSettings != nil:
		for room := range getLocifyRoomSettings {
			if room == l.RoomName {
				exists = true
			}
		}

		// if doesnt exists than create
		if !exists {

			switch true {
			case _friendRoom:
				log.Println("friend room")

				var code = GenCode() // random code generated for a friend to join the group

				h.wg.Go(func() {
					RegisterRoom(h, conn, l.RoomName, l.ID, false)
				})
				go func() {

					powers := map[string]bool{}
					powers[_RewindKey] = l.RewindPower
					powers[_CovertKey] = l.CovertPower
					powers[_NexusKey] = l.NexusPower
					powers[_BetKey] = l.BetPower
					powers[_DrawKey] = l.DrawPower
					powers[_TagKey] = l.TagPower
					powers[_FreezeKey] = l.FreezePower

					BoardcastRoomCreate(h, l.RoomName, l.Book,
						roomMode, l.Category, l.GameTime,
						l.DecisionTime, powers, l.RoomCapacity, true,
						false, code, l.SetToss, l.Reverse, l.Starter)

					StoreLocifyRoomSettings(l, roomMode, code, false)

					if l.RoomCapacity != 2 {
						r := []RoomList{}
						var t RoomList
						t.Book = l.Book
						t.Category = l.Category
						t.RoomName = l.RoomName
						t.Type = roomMode
						r = append(r, t)
						c, err := json.Marshal(&r)
						if err != nil {
							log.Println(err)
							return
						}
						token := "RoomLists: " + string(c)
						log.Println("room lists: ", r)
						BroadcastAll(h, token)
					}
				}()

			// testing done working perfectly
			// the only difference is that we are sending a friend code to join
			case _publicRoom:
				log.Println("public room")
				h.wg.Go(func() {
					RegisterRoom(h, conn, l.RoomName, l.ID, false)
				})

				go func() {

					powers := map[string]bool{}
					powers[_RewindKey] = l.RewindPower
					powers[_CovertKey] = l.CovertPower
					powers[_NexusKey] = l.NexusPower
					powers[_BetKey] = l.BetPower
					powers[_DrawKey] = l.DrawPower
					powers[_TagKey] = l.TagPower
					powers[_FreezeKey] = l.FreezePower

					StoreLocifyRoomSettings(l, roomMode, _StringSentinel_, true)
					BoardcastRoomCreate(h, l.RoomName, l.Book,
						roomMode, l.Category, l.GameTime,
						l.DecisionTime, powers, l.RoomCapacity, false, false, _StringSentinel_, l.SetToss, l.Reverse, l.Starter)
					r := []RoomList{}
					var t RoomList
					t.Book = l.Book
					t.Category = l.Category
					t.RoomName = l.RoomName
					t.Type = roomMode
					r = append(r, t)
					c, err := json.Marshal(&r)
					if err != nil {
						log.Println(err)
						return
					}
					token := "RoomLists: " + string(c)
					BroadcastAll(h, token)
				}()

			// it doesnt matter whether they have also
			// picked the friend to true or false
			// only diff between !friend room and private room
			// is i am just broadcasting the room code rather than saving owner
			case _privateRoom:
				log.Println("private room")
				h.wg.Go(func() {
					RegisterRoom(h, conn, l.RoomName, l.ID, false)
				})
				go func() {
					var code = GenCode() // random code generated for a friend to join the group

					powers := map[string]bool{}
					powers[_RewindKey] = l.RewindPower
					powers[_CovertKey] = l.CovertPower
					powers[_NexusKey] = l.NexusPower
					powers[_BetKey] = l.BetPower
					powers[_DrawKey] = l.DrawPower
					powers[_TagKey] = l.TagPower
					powers[_FreezeKey] = l.FreezePower

					StoreLocifyRoomSettings(l, roomMode, code, false)

					BoardcastRoomCreate(h, l.RoomName, l.Book, roomMode, l.Category, l.GameTime, l.DecisionTime,
						powers, l.RoomCapacity, false, true, code, l.SetToss, l.Reverse, l.Starter)

				}()

			default:
				log.Println("not able to decide to make a room for what")
			}

		} else {
			log.Println("room exists please create another room or wait")
		}

	default:
		log.Println("wrong request")

	}

	log.Println("room registeration doneXXX")
}

func BoardcastProfile(h *Hub, id, roomname string, onevone bool) {
	credits := profiles.FetchCredits([]string{id}, false, onevone)
	n := profiles.Fetch([]string{id})
	var roomJoin IRoomJoin
	for _, cred := range credits.PlayerCredits {
		roomJoin.Name = n.NickNames[id]
		roomJoin.Profile.Gamesplayed = cred.Profile.Gamesplayed
		roomJoin.Profile.Rating = cred.Profile.Rating
		roomJoin.Profile.Tier = cred.Profile.Tier
		roomJoin.Profile.Points = cred.Profile.Points
	}
	t, err := json.Marshal(roomJoin)
	if err != nil {
		log.Println(err)
		return
	}
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: "NYGplayerJoin: " + string(t)}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, token: "NYGmyNickname: " + n.NickNames[id], to: id, _sleep: false}
}
