package server

import (
	"encoding/json"
	"log"
)

func LocifySumUp(h *Hub, id string, roomname string, challengeToken string, fromBet, fromTimeup bool, bets []string) {
	log.Println("in sum tokens")
	myProfile := getLocifyProfile[id][roomname]
	myTeam := myProfile.MyTeam
	myOpponentID := myProfile.Against
	oppoProfile := getLocifyProfile[myOpponentID][roomname]
	_fromOpp := myProfile.OppoSetChallenge
	token := "ChallengeGuess: " + challengeToken
	fromOpp := "ChallengeGuess: " + getLocifyProfile[id][roomname].OppoSetChallenge

	log.Println("view powers oppo: ", myProfile.OppoPowerUp)
	log.Println("view powers my: ", oppoProfile.OppoPowerUp)

	// if the last call from the bet session
	if fromBet {
		BroadcastBets(h, myOpponentID, myTeam, roomname, bets)

	} else if myProfile.OppoBetDone {
		// if my oppo has used the bet than send the bet to me
		pr := getLocifyProfile[myOpponentID][roomname]
		_t := pr.MyTeam
		log.Println("bets: ", pr.OppoBets)
		BroadcastBets(h, id, _t, roomname, pr.OppoBets)

	}

	for _, _id := range saveShuffle[roomname][myProfile.OppoTeamname] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, roomname: roomname, _sleep: false, to: _id}
	}

	for _, _id := range saveShuffle[roomname][myTeam] {
		if !myProfile.OppoPowerUp[_BetKey] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: fromOpp, roomname: roomname, _sleep: false, to: _id}
		}
	}

	h.broadcast <- BroadcastReq{Token: RemoveChallenge + challengeToken, RoomID: roomname}

	// sending powers to opponent
	// its important to reset power of the id from which the power is being spotted than actual id
	switch true {
	case myProfile.OppoPowerUp[_NexusKey]:
		dictionary := getLocifyProfile[id][roomname].SetDictionary
		book := getLocifyProfile[id][roomname].Book
		PNexusMechanism(h, roomname, myOpponentID, book, dictionary, challengeToken, oppoProfile.MyGameProfile.NexusLevel)
		LocifyResetPowerUp(id, roomname, _NexusKey)
		BoardcastRemovePower(h, myOpponentID, roomname, _NexusKey)

	case myProfile.OppoPowerUp[_CovertKey]:
		PCovertMechanism(h, roomname, id) // id must be this cause we coverting the player's id
		LocifyResetPowerUp(id, roomname, _CovertKey)
		BoardcastRemovePower(h, myOpponentID, roomname, _CovertKey)

	case myProfile.OppoPowerUp[_FreezeKey]:
		PFreezeMechanism(h, roomname, myOpponentID, myTeam, false, false, oppoProfile.MyGameProfile.FreezeLevel)
		LocifyResetPowerUp(id, roomname, _FreezeKey)
		BoardcastRemovePower(h, myOpponentID, roomname, _FreezeKey)

	}

	switch true {
	case oppoProfile.OppoPowerUp[_NexusKey]:
		dictionary := getLocifyProfile[id][roomname].SetDictionary
		book := getLocifyProfile[id][roomname].Book
		PNexusMechanism(h, roomname, id, book, dictionary, _fromOpp, myProfile.MyGameProfile.NexusLevel)
		LocifyResetPowerUp(myOpponentID, roomname, _NexusKey)
		BoardcastRemovePower(h, id, roomname, _NexusKey)

	case oppoProfile.OppoPowerUp[_CovertKey]:
		PCovertMechanism(h, roomname, myOpponentID)
		LocifyResetPowerUp(myOpponentID, roomname, _CovertKey)
		BoardcastRemovePower(h, id, roomname, _CovertKey)

	case oppoProfile.OppoPowerUp[_FreezeKey]:
		PFreezeMechanism(h, roomname, id, myTeam, false, false, myProfile.MyGameProfile.FreezeLevel)
		LocifyResetPowerUp(myOpponentID, roomname, _CovertKey)
		BoardcastRemovePower(h, id, roomname, _FreezeKey)
	}

	SingleSave(false, myOpponentID,
		roomname, _StringSentinel_,
		_StringSentinel_, challengeToken,
		_StringSentinel_, nil,
		false, false, true, false)

	h.broadcast <- BroadcastReq{Token: Unblock, RoomID: roomname}
	h.broadcast <- BroadcastReq{Token: waiting, RoomID: roomname}
	BoardcastSession(h, roomname, false, false, false, false, true)

	match1 := SumupGametokens(roomname, myProfile.NickNamesViaID[id], myProfile.NickNamesViaID[myOpponentID], getLocifyProfile[id][roomname].MyGameProfile.Pic, getLocifyProfile[myOpponentID][roomname].MyGameProfile.Pic)
	_c, _ := json.Marshal(match1)
	_match1 := "NYGMatchup: " + string(_c)
	for _, _id := range saveShuffle[roomname][myTeam] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: _match1, to: _id, roomname: roomname,
			_sleep: false,
		}
	}
	LocifyHandleGameTime(h, roomname, myProfile, id, fromTimeup)
}

func LocifyHandleGameTime(h *Hub, roomname string, myProfile LocifyFixtures, id string, fromTimeup bool) {

	tim := "10"
	if fromTimeup {

		tim = "8"
		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: id, _sleep: false, token: gameTime + tim}
		x := getLocifyProfile[myProfile.Against][roomname].MyPenalties[myProfile.CurrentRound] == 1
		if x {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: myProfile.Against, _sleep: false, token: gameTime + tim}

		}
	} else {
		for _, _id := range saveShuffle[roomname][myProfile.MyTeam] {
			if getLocifyProfile[_id][roomname].MyPenalties[myProfile.CurrentRound] == 1 {

				tim = "8"
				h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: _id, _sleep: false, token: gameTime + tim}
			} else {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: _id, _sleep: false, token: gameTime + tim}
			}
		}
		for _, _id := range saveShuffle[roomname][myProfile.OppoTeamname] {
			if getLocifyProfile[_id][roomname].MyPenalties[myProfile.CurrentRound] == 1 {

				tim = "8"
				h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: _id, _sleep: false, token: gameTime + tim}
			} else {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: _id, _sleep: false, token: gameTime + tim}
			}
		}
	}

}

type IRemovePower struct {
	Power map[string]bool `json:"power"`
}

// BoardcastRemovePower oppoID to reset covert cases
func BoardcastRemovePower(h *Hub, id, roomname, power string) {
	log.Println("in BoardcastRemovePower")
	var p IRemovePower
	p.Power = make(map[string]bool)
	p.Power[power] = false
	c, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("removing power: ", power)

	ptoken := "NYGremovePower: " + string(c)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: ptoken, roomname: roomname, _sleep: false, to: id}
}

// BoardcastRefresh resets all the powers signal to default
func BoardcastRefresh(h *Hub, roomname string) {
	// freeze
	var t ParcelFreezePackage
	t.CanUnfreeze = false
	t.FreezeUse = false
	t.Unfreeze = true
	t.FreezeTime = 0
	x, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	tx := "NYGFreeze: " + string(x)
	h.broadcast <- BroadcastReq{
		RoomID: roomname,
		Token:  tx,
	}

	// end

	// nexus
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: _NexusUse,
	}
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: nexusWord,
	}
	// end

	// draw
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: _DrawUse,
	}
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: _drawOffer,
	}
	// end

	// rewind
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: _RewindUse,
	}
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: _backClock,
	}
	// end

	// tag
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: _TagUse,
	}
	// end

	// covert
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: _CovertUse,
	}
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: _underTest,
	}
	// end

	// bet
	b := PlayBet{DoBetting: false, Bets: nil, Challenge: false}
	b.DoBetting = false
	b.Bets = nil
	b.Challenge = false
	c, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}
	token := "ClashBet: " + string(c)
	h.broadcast <- BroadcastReq{
		RoomID: roomname, Token: token,
	}
	// end
}
