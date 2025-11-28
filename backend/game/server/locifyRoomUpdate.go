package server

import (
	"encoding/json"
)

type FriendTrack struct {
	ID   string
	Join bool
}

var (
	// room-name ->...
	friendJoin = make(chan map[string]FriendTrack)
)

func UpdateLocifyRoomSettings(h *Hub, roomname, friendID, book string, dTime, gTime int,
	setupToss, reverse, starter bool, nexus, covert, bet, freeze, draw, tag, rewind bool) {
	powers := map[string]bool{}

	if nexus {
		powers[_NexusKey] = nexus

	}
	if covert {
		powers[_CovertKey] = covert

	}
	if bet {
		powers[_BetKey] = bet

	}
	if freeze {
		powers[_FreezeKey] = freeze

	}
	if draw {
		powers[_DrawKey] = draw

	}
	if tag {
		powers[_TagKey] = tag

	}
	if rewind {
		powers[_RewindKey] = rewind
	}

	if !covert && !rewind && !freeze && !draw && !tag && !bet {
		powers = nil
	}

	var r RoomSettingsSave
	r.FriendID = friendID
	r.Book = getLocifyRoomSettings[roomname].Book
	r.DTime = getLocifyRoomSettings[roomname].DecisionTime
	r.GTime = getLocifyRoomSettings[roomname].GameTime

	r.SetupToss = setupToss
	r.Reverse = reverse
	r.Starter = starter

	r.Powers = powers

	var t TBoardRoomCreate

	t.Book = r.Book
	t.Mode = getLocifyRoomSettings[roomname].Mode
	t.Category = getLocifyRoomSettings[roomname].Category
	t.GameTime = r.GTime
	t.DecisionTime = r.DTime
	t.Powers = r.Powers
	t.Cap = getLocifyRoomSettings[roomname].Capacity
	t.Reverse = reverse
	t.Starter = starter
	t.Powers = powers
	t.Friend = getLocifyRoomSettings[roomname].Friend
	t.Client = false // becuase the req came from the owner
	t.IsEntertainment = getLocifyRoomSettings[roomname].Category == "entertainment"
	if getLocifyRoomSettings[roomname].Private || getLocifyRoomSettings[roomname].Friend {
		t.Code = getLocifyRoomSettings[roomname].Code
	}
	t.Session = true
	t.Roomname = roomname
	t.Category = getLocifyRoomSettings[roomname].Category

	c, _ := json.Marshal(t)
	token := _RoomSettingKey + string(c)
	h.broadcast <- BroadcastReq{Token: token, RoomID: roomname}

	store := map[string]RoomSettingsSave{roomname: r}
	saveRoomSettings <- store
}
