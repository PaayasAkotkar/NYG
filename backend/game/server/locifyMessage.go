// Package server you can send the id of the current requst
package server

import (
	"encoding/json"
	"fmt"
	"log"
)

func LocifyTossMessage(h *Hub, roomname, id string) {
	myProfile := getLocifyProfile[id][roomname]
	msg := "player is about to toss!!!!"
	var t AlertMessage

	t.Alert = true
	t.Message = msg
	c, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname, to: id, _sleep: false, token: Message + string(c),
	}

	msg = "you can now toss the coin!!!!"

	t.Alert = true
	t.Message = msg
	c, err = json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname, to: myProfile.Against, _sleep: false, token: Message + string(c),
	}

}

func LocifyTossDoneMessage(h *Hub, roomname, losserID, winnerID, winnerTeam string) {
	msg := fmt.Sprintf("you lost the toss %s team will set the dictionary!!!!", winnerTeam)
	var t AlertMessage

	t.Alert = true
	t.Message = msg
	c, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname, to: losserID, _sleep: false, token: Message + string(c),
	}

	msg = "you can now set the dictionary!!!!"

	t.Alert = true
	t.Message = msg
	c, err = json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname, to: winnerID, _sleep: false, token: Message + string(c),
	}
}

func LocifyDictionaryDoneMessage(h *Hub, roomname, dictionaryToken string) {
	msg := fmt.Sprintf("round style %s", dictionaryToken)
	var t AlertMessage
	t.Alert = true
	t.Message = msg
	c, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: Message + string(c)}

}

func LocifyChallengeAlertMessage(h *Hub, roomname, id string) {
	myProfile := getLocifyProfile[id][roomname]
	myOpponentID := myProfile.Against
	msg := "Hang in there!!!"
	var t AlertMessage
	t.Alert = true
	t.Message = msg
	c, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname, to: id, _sleep: false, token: Message + string(c),
	}

	msg = "you can now set the challenge"
	t.Alert = true
	t.Message = msg
	c, err = json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname, to: myProfile.Against, _sleep: false, token: Message + string(c),
	}

	if getLocifyRoomSettings[roomname].Capacity == 4 {
		msg = "your teammate is setting the challenge"
		t.Alert = true
		t.Message = msg
		c, err = json.Marshal(t)
		if err != nil {
			log.Println(err)
			return
		}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			roomname: roomname, to: getLocifyProfile[myOpponentID][roomname].MyPatnerID, _sleep: false, token: Message + string(c),
		}
		msg = "soon the players will be on the playground"
		t.Alert = true
		t.Message = msg
		c, err = json.Marshal(t)
		if err != nil {
			log.Println(err)
			return
		}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			roomname: roomname, to: myProfile.MyPatnerID, _sleep: false, token: Message + string(c),
		}

	}
}

func LocifyRoundOverMessage(h *Hub, winnerTeam, roomname string) {
	msg := fmt.Sprintf("team %s won this round", winnerTeam)
	var t AlertMessage
	t.Alert = true
	t.Message = msg
	c, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.broadcast <- BroadcastReq{Token: Message + string(c), RoomID: roomname}
}
