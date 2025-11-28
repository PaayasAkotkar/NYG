package server

import (
	"log"
)

func LocifyChallengeSession(h *Hub, id string, roomname string, challengeToken string, fromTimeup bool) {
	log.Println("Locify challenge disucssion")

	myTeam, myOpponentID := getLocifyProfile[id][roomname].MyTeam, getLocifyProfile[id][roomname].Against
	myProfile := getLocifyProfile[id][roomname]
	sessionDone := getLocifyProfile[id][roomname].OppoChallengeDone || getLocifyProfile[id][roomname].OppoBetDone
	var ux = GSpectate(myProfile.OppoTeamname, myProfile.NickNamesViaID[myProfile.Against])

	log.Println("set dictionary: ", getLocifyProfile[id][roomname].SetDictionary)
	log.Println("oppo bet: ", myProfile.OppoBetDone)
	switch true {
	case sessionDone:
		log.Println("seesionDone!!!!")
		bets := getLocifyProfile[id][roomname]
		LocifySumUp(h, id, roomname, challengeToken, false, fromTimeup, bets.OppoBets)

	default:
		log.Println("default!!!!")
		LocifyChallengeAlertMessage(h, roomname, id)
		for _, _id := range saveShuffle[roomname][getLocifyProfile[id][roomname].OppoTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: specate_ + ux, roomname: roomname, to: id, _sleep: false}
		}
		h.broadcast <- BroadcastReq{Token: RemoveChallenge + challengeToken,
			RoomID: roomname}

		for _, _id := range saveShuffle[roomname][myTeam] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}
		}

		// if the oppo had used bet
		if getLocifyProfile[id][roomname].OppoPowerUp[_BetKey] {
			dictionary := getLocifyProfile[id][roomname].SetDictionary
			b := getLocifyRoomSettings[roomname].Book

			bets := InitBet(h, myOpponentID, myTeam, roomname, b, dictionary, false)
			log.Println("fetched bets: ", bets)
			MasterSave(false, _IntSentinel, _StringSentinel_, _StringSentinel_,
				_StringSentinel_, roomname, myTeam,
				_StringSentinel_, _StringSentinel_,
				myOpponentID, false, false, _IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, bets, nil)

			LocifyResetPowerUp(id, roomname, _BetKey)
		}

		go func() {
			SingleSave(false, myOpponentID, roomname,
				_StringSentinel_, _StringSentinel_,
				challengeToken, _StringSentinel_, nil, false, false, true, false)
		}()

	}

}
