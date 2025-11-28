package server

import "log"

type PlayBet struct {
	DoBetting bool     `json:"doBet"`
	Bets      []string `json:"bets"`
	Challenge bool     `json:"challenge"`
}

// ClashDictionarySession returns
// cases:
// highest lives
// id must be of the winner
// note: in clash one team sets the dictionary while other sets the challenge
// for clash MaxPlayerCount =3
func ClashDictionarySession(h *Hub, roomname string, id string, dictionaryToken string, MaxPlayerCount int) {

	log.Println("dictionary disucssion")
	myProfile := getClashProfile[id][roomname]
	myTeam := getClashProfile[id][roomname].MyTeam
	// myOpponentID := getClashProfile[id][roomname].Against
	Count := int(1)
	// because the default count is 1
	sessionDone := (getClashProfile[id][roomname].Count + Count) == int(MaxPlayerCount)
	_isBet := getClashProfile[id][roomname].OppoPowerUp[_BetKey]
	b := myProfile.Book
	token := SendList(b, dictionaryToken)

	// proceed only if both of them done with the challennge
	if sessionDone {
		for _, _id := range saveShuffle[roomname][myTeam] {
			if id != _id {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
			} else {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RemoveDictionary + dictionaryToken, roomname: roomname, _sleep: false, to: _id}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: ItemsURL + token, roomname: roomname, _sleep: false, to: _id}

		}

		if _isBet {
			b := getClashProfile[id][roomname].Book
			InitBet(h, myProfile.Against, myTeam, roomname, b, dictionaryToken, true)
			ClashResetPowerUp(id, roomname, _BetKey)
			BoardcastRemovePower(h, myProfile.Against, roomname, _BetKey)

		} else if x := getClashProfile[myProfile.Against][roomname]; x.OppoPowerUp[_BetKey] {
			b := getClashProfile[id][roomname].Book
			ClashResetPowerUp(id, roomname, _BetKey)
			InitBet(h, myProfile.Against, myTeam, roomname, b, dictionaryToken, true)
			BoardcastRemovePower(h, id, roomname, _BetKey)

		}
		// h.broadcast <- BroadcastReq{Token: _challengeDiscussion, RoomID: roomname}
		// h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomID: roomname}

		// h.broadcast <- BroadcastReq{Token: _DictionarySet, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: waiting, RoomID: roomname}

		CommonSave(true, roomname, myTeam, nil, false, true, false, false)
		OneTimeSave(true, roomname, myTeam, dictionaryToken)

		BoardcastSession(h, roomname, false, false, true, false, false)

		re := map[string]bool{}
		re[roomname] = true
		clashResetCount <- re
	} else {

		for _, _id := range saveShuffle[roomname][myTeam] {
			if id != _id {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
			} else {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: DictionaryEvent + dictionaryToken, roomname: roomname, _sleep: false, to: _id}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _waiting, roomname: roomname, _sleep: false, to: _id}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _ClashWaitMessgae, roomname: roomname, _sleep: false, to: _id}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RemoveDictionary + dictionaryToken, roomname: roomname, _sleep: false, to: _id}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: ItemsURL + token, roomname: roomname, _sleep: false, to: _id}
		}

		if _isBet {
			b := getClashProfile[id][roomname].Book
			InitBet(h, myProfile.Against, myTeam, roomname, b, dictionaryToken, true)
			ClashResetPowerUp(id, roomname, _BetKey)
			BoardcastRemovePower(h, myProfile.Against, roomname, _BetKey)

		} else if x := getClashProfile[myProfile.Against][roomname]; x.OppoPowerUp[_BetKey] {
			b := getClashProfile[id][roomname].Book
			ClashResetPowerUp(id, roomname, _BetKey)
			InitBet(h, myProfile.Against, myTeam, roomname, b, dictionaryToken, true)
			BoardcastRemovePower(h, id, roomname, _BetKey)
		}

		MasterSave(true, int(Count), _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			roomname, myTeam, _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			false, false, _IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, nil, nil)
		CommonSave(true, roomname, myTeam, nil, false, true, false, false)
		OneTimeSave(true, roomname, myTeam, dictionaryToken)
	}

	log.Println("dictionary token: ", dictionaryToken)

}
