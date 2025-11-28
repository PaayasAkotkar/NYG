package server

import (
	"encoding/json"
	"log"
)

type MatchUp struct {
	Img  string `json:"img"`
	Name string `json:"name"`
}
type GameMiscelleanous struct {
	Matchup map[string]MatchUp `json:"matchup"`
}

// SumupGametokens returns the struct the redPlayer and bluePlayer is not the id it can be user name
func SumupGametokens(roomname, redPlayer, bluePlayer, redImg, blueImg string) GameMiscelleanous {
	var matchups GameMiscelleanous
	matchups.Matchup = map[string]MatchUp{}
	var j = matchups.Matchup
	var mat MatchUp
	mat.Name = redPlayer
	mat.Img = redImg
	j["RED"] = mat
	mat.Img = blueImg
	mat.Name = bluePlayer
	j["BLUE"] = mat
	matchups.Matchup = j
	return matchups
}

// SumupTokensGame fromBet if the method is called from the bet
func SumupTokensGame(h *Hub, id string, roomname string, challengeToken string, fromBet, fromTimeUp bool, bets []string) {
	log.Println("in sum tokens")
	myProfile := getClashProfile[id][roomname]
	myTeam := myProfile.MyTeam
	myOpponentID := myProfile.Against
	oppoProfile := getClashProfile[myOpponentID][roomname]

	_fromOpp := myProfile.OppoSetChallenge
	token := "ChallengeGuess: " + challengeToken
	fromOpp := "ChallengeGuess: " + getClashProfile[id][roomname].OppoSetChallenge

	log.Println("bet ids: ", myProfile.BetIDs, "my id: ", id)
	log.Println("from oppo: ", fromOpp)
	log.Println("opp prof: ", getClashProfile[myOpponentID][roomname])
	log.Println("my prof: ", myProfile)
	log.Println("challenge token: ", challengeToken)
	log.Println("my team: ", myTeam, " next team: ", myProfile.NextTeamname)

	// if the last call from the bet session
	if fromBet {
		BroadcastBets(h, myOpponentID, myTeam, roomname, bets)

	} else {
		for _, _id := range myProfile.BetIDs {
			pr := getClashProfile[_id][roomname]
			_t := pr.MyTeam
			log.Println("bets: ", pr.OppoBets)
			BroadcastBets(h, _id, _t, roomname, pr.OppoBets)
		}
	}

	for _, _id := range saveShuffle[roomname][myTeam] {
		if _id != id {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, roomname: roomname, _sleep: false, to: _id}
		} else if !myProfile.OppoPowerUp[_BetKey] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: fromOpp, roomname: roomname, _sleep: false, to: _id}
		}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RemoveChallenge + challengeToken, roomname: roomname, _sleep: false, to: _id}
	}

	// sending powers to opponent
	if getClashProfile[id][roomname].OppoPowerUp[_FreezeKey] && getClashProfile[myOpponentID][roomname].OppoPowerUp[_FreezeKey] {
		log.Println("both")
		PFreezeMechanism(h, roomname, myOpponentID, myTeam, true, true, myProfile.MyGameProfile.FreezeLevel)
		go func() {
			ClashResetPowerUp(myOpponentID, roomname, _FreezeKey)
		}()
		go func() {
			ClashResetPowerUp(id, roomname, _FreezeKey)
		}()
		BoardcastRemovePower(h, myOpponentID, roomname, _FreezeKey)
		BoardcastRemovePower(h, id, roomname, _FreezeKey)

	} else {
		switch true {
		case getClashProfile[id][roomname].OppoPowerUp[_NexusKey]:
			dictionary := getClashProfile[id][roomname].SetDictionary
			book := getClashProfile[id][roomname].Book
			PNexusMechanism(h, roomname, myOpponentID, book, dictionary, challengeToken, oppoProfile.MyGameProfile.NexusLevel)
			ClashResetPowerUp(id, roomname, _NexusKey)
			BoardcastRemovePower(h, myOpponentID, roomname, _NexusKey)

		case getClashProfile[id][roomname].OppoPowerUp[_CovertKey]:
			PCovertMechanism(h, roomname, id) // id must be this cause we coverting the player's id
			ClashResetPowerUp(id, roomname, _CovertKey)
			BoardcastRemovePower(h, myOpponentID, roomname, _CovertKey)

		case getClashProfile[id][roomname].OppoPowerUp[_FreezeKey]:
			PFreezeMechanism(h, roomname, myOpponentID, myTeam, true, false, oppoProfile.MyGameProfile.FreezeLevel)
			ClashResetPowerUp(id, roomname, _FreezeKey)
			BoardcastRemovePower(h, myOpponentID, roomname, _FreezeKey)

		}

		switch true {
		case getClashProfile[myOpponentID][roomname].OppoPowerUp[_NexusKey]:
			dictionary := getClashProfile[id][roomname].SetDictionary
			book := getClashProfile[id][roomname].Book
			PNexusMechanism(h, roomname, id, book, dictionary, _fromOpp, myProfile.MyGameProfile.NexusLevel)
			ClashResetPowerUp(myOpponentID, roomname, _NexusKey)
			BoardcastRemovePower(h, id, roomname, _NexusKey)

		case getClashProfile[myOpponentID][roomname].OppoPowerUp[_CovertKey]:
			PCovertMechanism(h, roomname, myOpponentID)
			ClashResetPowerUp(myOpponentID, roomname, _CovertKey)
			BoardcastRemovePower(h, id, roomname, _CovertKey)

		case getClashProfile[myOpponentID][roomname].OppoPowerUp[_FreezeKey]:
			PFreezeMechanism(h, roomname, id, myTeam, true, false, myProfile.MyGameProfile.FreezeLevel)
			ClashResetPowerUp(myOpponentID, roomname, _FreezeKey)
			BoardcastRemovePower(h, id, roomname, _FreezeKey)

		}
	}

	SingleSave(true, myOpponentID, roomname, myTeam, _StringSentinel_, challengeToken, _StringSentinel_, nil,
		false, false, true, false)

	h.broadcast <- BroadcastReq{Token: challengeDiscussion, RoomID: roomname}
	h.broadcast <- BroadcastReq{Token: _startGame, RoomID: roomname}
	h.broadcast <- BroadcastReq{Token: Unblock, RoomID: roomname}
	h.broadcast <- BroadcastReq{Token: waiting, RoomID: roomname}

	match1 := SumupGametokens(roomname, getClashNicknames[roomname][id], getClashNicknames[roomname][myOpponentID], getClashProfile[id][roomname].MyCredits.ImageURL, getClashProfile[myOpponentID][roomname].MyCredits.ImageURL)
	_c, _ := json.Marshal(match1)
	_match1 := "NYGMatchup: " + string(_c)
	for _, _id := range saveShuffle[roomname][myTeam] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: _match1, to: _id, roomname: roomname,
			_sleep: false,
		}
	}

	r1 := saveShuffle[roomname][myProfile.NextTeamname][0]
	b1 := saveShuffle[roomname][myProfile.NextTeamname][1]
	match2 := SumupGametokens(roomname, getClashNicknames[roomname][r1], getClashNicknames[roomname][b1], getClashProfile[r1][roomname].MyCredits.ImageURL, getClashProfile[b1][roomname].MyCredits.ImageURL)
	c_, _ := json.Marshal(match2)
	match2_ := "NYGMatchup: " + string(c_)
	for _, _id := range saveShuffle[roomname][myProfile.NextTeamname] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: match2_, to: _id, roomname: roomname,
			_sleep: false,
		}
	}
	ClashHandleGameTime(h, roomname, myProfile)

	re := map[string]bool{}
	re[roomname] = true
	clashResetCount <- re
}

func ClashHandleGameTime(h *Hub, roomname string, myProfile ClashFixtures) {
	tim := "10"

	for _, _id := range saveShuffle[roomname][myProfile.MyTeam] {
		if getClashProfile[_id][roomname].MyPenalties[myProfile.CurrentRound] == 1 {
			tim = "8"
			h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: _id, _sleep: false, token: gameTime + tim}
		} else {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: _id, _sleep: false, token: gameTime + tim}
		}
	}

	for _, _id := range saveShuffle[roomname][myProfile.NextTeamname] {
		if getClashProfile[_id][roomname].MyPenalties[myProfile.CurrentRound] == 1 {
			tim = "8"
			h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: _id, _sleep: false, token: gameTime + tim}
		} else {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: _id, _sleep: false, token: gameTime + tim}
		}
	}

}

type BetUsed struct {
	Betting bool     `json:"isBetting"`
	BetCups []string `json:"cups"`
}

// BroadcastBets id must be the id to sent token bets to
func BroadcastBets(h *Hub, id, teamname, roomname string, bets []string) {
	log.Println("in bets")
	x := BetUsed{}
	x.BetCups = bets
	x.Betting = true
	y, _ := json.Marshal(x)
	parcel := string(y)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		token: "ClashUnderClashBet: " + parcel, _sleep: false, to: id, roomname: roomname,
	}
}
