package server

import (
	"encoding/json"
	"log"
)

// InitBet id must be the used power
func InitBet(h *Hub, id string, teamname string, roomname string, book, dictionary string, clash bool) []string {
	log.Println("in bet init")
	_teamname, cups := BetCups(book, dictionary)
	log.Println("bet cups: ", cups)
	_token := PlayBet{}
	_token.DoBetting = true
	_token.Challenge = false
	_token.Bets = []string{cups.FirstCup, cups.SecondCup, cups.ThirdCup}
	_c, _ := json.Marshal(_token)
	token := string(_c)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashBet: " + token,
		roomname: roomname, _sleep: false, to: id}

	log.Println("sending bet")
	if clash {
		CommonSave(true, roomname, teamname, _token.Bets, false, false, false, false)
		// if the last call is made by none bet
		SingleSave(true, getClashProfile[id][roomname].Against, roomname, teamname, _StringSentinel_, _teamname, _StringSentinel_,
			nil, false, false, false, false)
	} else {
		// if the last call is made by none bet
		SingleSave(false, getLocifyProfile[id][roomname].Against, roomname,
			_StringSentinel_, _StringSentinel_, _teamname, _StringSentinel_,
			nil, false, false, false, false)
	}
	return _token.Bets
}
