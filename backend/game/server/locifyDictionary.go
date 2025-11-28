package server

import (
	"log"
)

// LocifyDictionarySession returns
// cases:
// highest lives
// id must be of the winner
// note: in Locify one team sets the dictionary while other sets the challenge
// for Locify MaxPlayerCount =3
func LocifyDictionarySession(h *Hub, roomname string, id string, dictionaryToken string) {

	log.Println("dictionary disucssion")
	myProfile := getLocifyProfile[id][roomname]
	myTeam := getLocifyProfile[id][roomname].MyTeam
	_isBet := myProfile.OppoPowerUp[_BetKey]
	b := getLocifyRoomSettings[roomname].Book
	token := SendList(b, dictionaryToken)
	var src = myProfile.RoomSettings
	isOppoSpectating := false
	if src.Reverse {
		isOppoSpectating = true
		// proceed only if both of them done with the challennge
		for _, _id := range saveShuffle[roomname][myProfile.OppoTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
		}
		for _, _id := range saveShuffle[roomname][myTeam] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}

		}

	} else {
		for _, _id := range saveShuffle[roomname][myTeam] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
		}
		for _, _id := range saveShuffle[roomname][myProfile.OppoTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}
		}
	}

	if _isBet {
		bets := InitBet(h, myProfile.Against, myTeam, roomname, b, dictionaryToken, false)
		LocifyResetPowerUp(id, roomname, _BetKey)

		MasterSave(false, _IntSentinel, _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			roomname, _StringSentinel_, _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			false, false, _IntSentinel, _IntSentinel,
			_StringSentinel_, dictionaryToken, bets, nil)

	} else if x := getLocifyProfile[myProfile.Against][roomname].OppoPowerUp[_BetKey]; x {
		bets := InitBet(h, id, myTeam, roomname, b, dictionaryToken, false)
		MasterSave(false, _IntSentinel, _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			roomname, _StringSentinel_, _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			false, false, _IntSentinel, _IntSentinel,
			_StringSentinel_, dictionaryToken, bets, nil)
		LocifyResetPowerUp(myProfile.Against, roomname, _BetKey)

	} else {
		MasterSave(false, _IntSentinel, _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			roomname, _StringSentinel_, _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			false, false, _IntSentinel, _IntSentinel,
			_StringSentinel_, dictionaryToken, nil, nil)
	}
	var ux = GSpectate(myProfile.OppoTeamname, myProfile.NickNamesViaID[myProfile.Against])

	if isOppoSpectating {

		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: specate_ + ux, roomname: roomname, to: myProfile.Against, _sleep: false}

	} else {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: specate_ + ux, roomname: roomname, to: id, _sleep: false}

	}
	LocifyDictionaryDoneMessage(h, roomname, dictionaryToken)

	h.broadcast <- BroadcastReq{Token: waiting, RoomID: roomname}
	h.broadcast <- BroadcastReq{Token: RemoveDictionary + dictionaryToken, RoomID: roomname}
	h.broadcast <- BroadcastReq{Token: ItemsURL + token, RoomID: roomname}

	SingleSave(false, myProfile.Against,
		roomname, _StringSentinel_,
		_StringSentinel_, _StringSentinel_,
		_StringSentinel_, nil, false, true, false, false)

	BoardcastSession(h, roomname, false, false, true,
		false, false)
	log.Println("dictionary token: ", dictionaryToken)

}
