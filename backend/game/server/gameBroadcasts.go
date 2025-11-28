package server

import (
	"encoding/json"
	"log"
)

type IWait struct {
	Msg     string `json:"msg"`
	Waiting bool   `json:"waiting"`
}
type ISessionUpdate struct {
	Toss       bool `json:"toss"`
	Dictionary bool `json:"dictionary"`
	Challenge  bool `json:"challenge"`
	Game       bool `json:"game"`
}

func BWait(h *Hub, msg, roomname, teamname string, waiting_ bool) {
	var w IWait
	w.Waiting = waiting_
	w.Msg = msg

	c, err := json.Marshal(w)
	token := "NYGwait: " + string(c)
	if err != nil {
		log.Println(err)
		return
	}
	for _, _id := range saveShuffle[roomname][teamname] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: token, roomname: roomname, to: _id,
			_sleep: false,
		}
	}
}
