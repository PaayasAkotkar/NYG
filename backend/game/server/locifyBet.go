package server

import (
	"encoding/json"
	"log"
)

func LocifyBetSession(h *Hub, id string, roomname string, betToken string, fromTimeUp bool) {
	log.Println("in Locify bet session")

	myProfile := getLocifyProfile[id][roomname]
	proceed := getLocifyProfile[id][roomname].OppoChallengeDone
	_token := PlayBet{}
	_token.DoBetting = true
	_token.Bets = myProfile.OppoBets
	_c, _ := json.Marshal(_token)
	token := string(_c)

	// if the opponent has used bet too
	if myProfile.OppoPowerUp[_BetKey] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: "LocifyBet: " + token, to: myProfile.Against, roomname: roomname, _sleep: false,
		}
	} else {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: Unblock, roomname: roomname,
			to: myProfile.Against, _sleep: false,
		}
	}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _waiting, to: id, roomname: roomname, _sleep: false}

	MasterSave(false, 1, _StringSentinel_, _StringSentinel_, _StringSentinel_, roomname, myProfile.MyTeam,
		_StringSentinel_, _StringSentinel_, _StringSentinel_, false,
		false, _IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, nil, nil)

	SingleSave(false, myProfile.Against, roomname, myProfile.MyTeam,
		_StringSentinel_, _StringSentinel_, betToken, nil, false, false, false, true)

	if proceed {
		log.Println("from bet to  challenge")
		LocifySumUp(h, id, roomname, "", true, fromTimeUp, _token.Bets)
	}

}
