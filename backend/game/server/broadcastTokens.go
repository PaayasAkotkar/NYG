package server

import (
	"encoding/json"
	"log"
)

type reqGameRoomBroadcast struct {
	roomname string
	token    string
	to       string
	_sleep   bool
}
type BMiscelleanousTokens struct {
	RoomMode   string         `json:"roomMode"`
	Nicknames  []string       `json:"nicknames"`
	CheatSheet map[string]int `json:"cheatSheet"`
}

func BoardcastSession(h *Hub, roomname string, clash, toss, challenge, dictionary, game bool) {
	log.Println("in brodacast session")
	var s BSession
	s.Clash = clash
	s.Toss = toss
	s.Dictionary = dictionary
	s.Challenge = challenge
	s.Game = game
	c, _ := json.Marshal(s)
	token := SessionKey + string(c)
	h.broadcast <- BroadcastReq{Token: token, RoomID: roomname}
}
