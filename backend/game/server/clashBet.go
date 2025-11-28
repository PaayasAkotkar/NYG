package server

import (
	"encoding/json"
	"log"
)

// ClashBetSession implements only if the dictionary is set
// case from dictionary setter-> than opponent either to the challenge session or if used bet than send bets
// case from non dictionary setter-> count+1 and opponent either to the challenge session or if used bet than send bets
// count will always be send
// maxCount will be same of the challenge session
func ClashBetSession(h *Hub, id string, roomname string, betToken string, maxCount int, fromTimeup bool) {
	log.Println("in clash bet session")

	myProfile := getClashProfile[id][roomname]
	proceed := myProfile.Count == int(maxCount)
	_token := PlayBet{}
	_token.DoBetting = true
	_token.Bets = myProfile.OppoBets
	_c, _ := json.Marshal(_token)
	token := string(_c)

	// if the opponent has used bet too
	if myProfile.OppoPowerUp[_BetKey] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: "ClashBet: " + token, to: myProfile.Against, roomname: roomname, _sleep: false,
		}
	} else {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: Unblock, roomname: roomname,
			to: myProfile.Against, _sleep: false,
		}
	}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _waiting, to: id, roomname: roomname, _sleep: false}

	if proceed {
		SumupTokensGame(h, id, roomname, "", true, fromTimeup, _token.Bets)
		CommonSave(true, roomname, myProfile.MyTeam, _token.Bets, false, false, true, false)
		SingleSave(true, myProfile.Against, roomname, myProfile.MyTeam,
			_StringSentinel_, _StringSentinel_, betToken, nil, false, false, false, true)
	}
	MasterSave(true, 1, _StringSentinel_, _StringSentinel_, _StringSentinel_, roomname, myProfile.MyTeam,
		_StringSentinel_, _StringSentinel_, _StringSentinel_, false,
		false, _IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, nil, nil)
	CommonSave(true, roomname, myProfile.MyTeam, _token.Bets, false, false, true, false)
	SingleSave(true, myProfile.Against, roomname, myProfile.MyTeam,
		_StringSentinel_, _StringSentinel_, betToken, nil, false, false, false, true)

}
