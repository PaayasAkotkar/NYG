package server

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

func StoreLocifyRoomSettings(l Lobby, gameType, code string, public bool) {
	fmt.Println("in store locify room")
	var lr LocifyRoomSettings
	lr.Name = l.RoomName
	lr.RoomOwnerID = l.ID
	lr.Category = l.Category
	lr.Starter = l.Starter
	lr.Book = l.Book
	lr.Capacity = l.RoomCapacity
	lr.Field = l.Field
	lr.SetupToss = l.SetToss
	lr.Reverse = l.Reverse
	lr.TwoVTwo = l.RoomCapacity == 4
	lr.Mode = gameType
	lr.Private = l.PrivateRoom
	lr.Public = public
	lr.Friend = l.Friend
	lr.DecisionTime = l.DecisionTime
	lr.GameTime = l.GameTime
	lr.FriendJoin = false
	lr.Code = code
	if code == _StringSentinel_ {
		x := uuid.NewString()[:10]
		lr.Code = x
	}
	powers := map[string]bool{}
	powers[_RewindKey] = l.RewindPower
	powers[_CovertKey] = l.CovertPower
	powers[_NexusKey] = l.NexusPower
	powers[_BetKey] = l.BetPower
	powers[_DrawKey] = l.DrawPower
	powers[_TagKey] = l.TagPower
	powers[_FreezeKey] = l.FreezePower
	lr.Powers = powers
	go func() {
		log.Println("storing book: ", lr.Book)
		store := map[string]LocifyRoomSettings{}
		store[l.RoomName] = lr
		storeLocifyRoomSettings <- store
		log.Println("storing: ", store)
	}()
}

type TBoardRoomCreate struct {
	Book            string          `json:"book"`
	Mode            string          `json:"mode"`
	Category        string          `json:"category"`
	Code            string          `json:"code"`
	Roomname        string          `json:"roomname"`
	GameTime        int             `json:"gameTime"`
	DecisionTime    int             `json:"decisionTime"`
	Cap             int             `json:"cap"`
	Powers          map[string]bool `json:"powers"`
	IsEntertainment bool            `json:"isEntertainment"`
	Friend          bool            `json:"friend"`
	Private         bool            `json:"private"`
	Reverse         bool            `json:"reverse"`
	Starter         bool            `json:"starter"`
	SetupToss       bool            `json:"setupToss"`
	Client          bool            `json:"client"`
	Session         bool            `json:"session"`
}

func BoardcastRoomCreate(h *Hub, roomname, roomBook,
	roomMode, category string, gameTime, decisionTime int,
	powers map[string]bool, cap int,
	friend bool, private bool, code string, setupToss, reverse, starter bool) {
	fmt.Println("boardcast creating")
	var t TBoardRoomCreate
	t.Book = roomBook
	t.Mode = roomMode
	t.Category = category
	t.GameTime = gameTime
	t.DecisionTime = decisionTime
	t.Powers = powers
	t.Cap = cap
	t.Roomname = roomname
	t.Friend = friend
	t.Private = private
	t.Reverse = reverse
	t.Starter = starter
	t.SetupToss = setupToss
	t.Session = true
	t.Client = false

	if code != _StringSentinel_ {
		t.Code = code
	} else {
		t.Code = "NYG"
	}

	if strings.ToLower(category) == "entertainment" {
		t.IsEntertainment = true
	} else {
		t.IsEntertainment = false
	}

	c, _ := json.Marshal(t)
	token := _RoomSettingKey + string(c)
	h.broadcast <- BroadcastReq{Token: token, RoomID: roomname}
}

// BoardcastRoomJoin only send friend but not code
func BoardcastRoomJoin(h *Hub, id, roomname string) {
	var r = getLocifyRoomSettings[roomname]
	var t TBoardRoomCreate
	t.Book = r.Book
	t.Mode = r.Mode
	t.Category = r.Category
	t.GameTime = r.GameTime
	t.DecisionTime = r.DecisionTime
	t.Powers = r.Powers
	t.Cap = r.Capacity
	t.Roomname = roomname
	t.Friend = r.Friend // if the room is of friend
	t.Private = r.Private
	t.SetupToss = getLocifyRoomSettings[roomname].SetupToss
	t.Reverse = getLocifyRoomSettings[roomname].Reverse
	t.Starter = getLocifyRoomSettings[roomname].Starter
	t.Code = "NYG"
	t.Client = true
	t.Session = true

	if strings.ToLower(r.Category) == "entertainment" {
		t.IsEntertainment = true
	} else {
		t.IsEntertainment = false
	}
	c, _ := json.Marshal(t)
	token := _RoomSettingKey + string(c)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		token: token, roomname: roomname, to: id, _sleep: false,
	}
}
