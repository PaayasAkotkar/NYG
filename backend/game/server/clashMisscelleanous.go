package server

import (
	"encoding/json"
	"log"
)

// CalcResult use this for final action
// meaning 1v1
// fromLastDance is imp for 3way
func CalcResult(roomname, myTeam, winnerID, losserID string,
	round, winnerCurrentChance, losserCurrentChance int,
	profile ClashFixtures, fromLastDance bool) ParcelClashResult {
	log.Println("in calc clash results")

	gameResult := ParcelClashResult{}

	gameResult.RoomName = roomname
	gameResult.Round = round // because the frontend sends via +1
	// gameResult.FinalBoss = false
	// gameResult.LastDance = true

	gameResult.Pairing = map[string]Pairing{}
	var Losser = gameResult.Pairing[getClashNicknames[roomname][losserID]]
	var Winner = gameResult.Pairing[getClashNicknames[roomname][winnerID]]

	Losser.Chances = losserCurrentChance
	Winner.Chances = winnerCurrentChance

	Winner.TeamName = "RED"
	Losser.TeamName = "RED"

	for _, _id := range profile.EliminatedPlayersIDs {
		if _, ok := gameResult.Pairing[getClashNicknames[roomname][_id]]; !ok {
			src := gameResult.Pairing[getClashNicknames[roomname][_id]]
			src.Chances = getClashProfile[_id][roomname].MyCurrentChances
			src.TeamName = "GREY"
			gameResult.Pairing[getClashNicknames[roomname][_id]] = src
		}
	}
	if !fromLastDance {
		var x = gameResult.Pairing[getClashNicknames[roomname][profile.FinalBossID]]
		x.Chances = getClashProfile[roomname][profile.FinalBossID].MyCurrentChances
		gameResult.Pairing[getClashNicknames[roomname][winnerID]] = x
	}

	gameResult.Pairing[getClashNicknames[roomname][winnerID]] = Winner
	gameResult.Pairing[getClashNicknames[roomname][losserID]] = Losser

	log.Println("game result: ", gameResult)
	return gameResult
}

// CalcClash3wayResults use this after player elimination
func CalcClash3wayResults(roomname, myTeam, winnerID, losserID string,
	round, winnerCurrentChance, losserCurrentChance int,
	profile ClashFixtures) ParcelClashResult {
	log.Println("in calc clash results")

	gameResult := ParcelClashResult{}

	gameResult.RoomName = roomname
	gameResult.Round = round // because the frontend sends via +1
	// gameResult.FinalBoss = true
	// gameResult.LastDance = false

	gameResult.Pairing = map[string]Pairing{}
	var Losser = gameResult.Pairing[getClashNicknames[roomname][losserID]]
	var Winner = gameResult.Pairing[getClashNicknames[roomname][winnerID]]

	Losser.Chances = losserCurrentChance
	Winner.Chances = winnerCurrentChance
	Winner.TeamName = "RED"
	Losser.TeamName = "RED"

	for _, _id := range profile.IDs {
		if _id != winnerID && _id != losserID {
			if _, ok := gameResult.Pairing[getClashNicknames[roomname][_id]]; !ok {
				src := gameResult.Pairing[getClashNicknames[roomname][_id]]
				src.Chances = getClashProfile[_id][roomname].MyCurrentChances
				if _id == profile.FinalBossID {
					src.TeamName = "BLACK" // final boss
				} else {
					src.TeamName = "GREY" // eliminators
				}
				gameResult.Pairing[getClashNicknames[roomname][_id]] = src
			}
		}
	}

	gameResult.Pairing[getClashNicknames[roomname][winnerID]] = Winner
	gameResult.Pairing[getClashNicknames[roomname][losserID]] = Losser

	log.Println("game result: ", gameResult)
	return gameResult
}

type IResetPower struct {
	Key  string
	DoIt bool
}

func ClashResetPowerUp(id, roomname, key string) {
	y := map[string]map[string]IResetPower{}
	var r IResetPower
	r.DoIt = true
	r.Key = key
	y[id] = map[string]IResetPower{
		roomname: r,
	}
	clashresetPowerUp <- y
}

func Min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

// func ClashElimationMessage(h *Hub, roomname string, kickoutID string) {
// 	h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 		token: _gameBegin, to: kickoutID,
// 		_sleep: false, roomname: roomname,
// 	}
// }

// CalcClashResults returns the currentChances of losser, gameresult, and matchup
func CalcClashResults(roomname, myTeam, winnerID,
	losserID, nextTeamLosserID, nextTeamWinnerID, nextTeamname string,
	round, winnerCurrentChance, losserCurrentChance int,
	Newmatchups map[string]map[string][]string) ParcelClashResult {
	log.Println("in calc clash results")

	gameResult := ParcelClashResult{}

	gameResult.RoomName = roomname
	gameResult.Round = round // because the frontend sends via +1
	// gameResult.FinalBoss = finalBoss
	// gameResult.LastDance = lastDance

	// note: prevMatchLosser and prevMatchWinner is not related to this result
	matches := Newmatchups

	gameResult.Pairing = map[string]Pairing{}
	var Winner = gameResult.Pairing[getClashNicknames[roomname][winnerID]]
	var Losser = gameResult.Pairing[getClashNicknames[roomname][losserID]]
	var prevMatchLosser = gameResult.Pairing[getClashNicknames[roomname][nextTeamLosserID]]
	var prevMatchWinner = gameResult.Pairing[getClashNicknames[roomname][nextTeamWinnerID]]

	Winner.Chances = winnerCurrentChance
	Losser.Chances = losserCurrentChance

	prevMatchLosser.Chances = getClashProfile[nextTeamLosserID][roomname].MyCurrentChances
	prevMatchWinner.Chances = getClashProfile[nextTeamWinnerID][roomname].MyCurrentChances

	for _, _id := range Newmatchups[roomname][_TeamRedKey] {
		if winnerID == _id {
			Winner.TeamName = "RED"
		} else if losserID == _id {
			Losser.TeamName = "RED"
		} else if nextTeamWinnerID == _id {
			prevMatchWinner.TeamName = "RED"
		} else if nextTeamLosserID == _id {
			prevMatchLosser.TeamName = "RED"
		}
	}

	for _, _id := range matches[roomname][_TeamBlueKey] {
		if winnerID == _id {
			Winner.TeamName = "BLUE"
		} else if losserID == _id {
			Losser.TeamName = "BLUE"
		} else if nextTeamWinnerID == _id {
			prevMatchWinner.TeamName = "BLUE"
		} else if nextTeamLosserID == _id {
			prevMatchLosser.TeamName = "BLUE"
		}
	}
	gameResult.Pairing[getClashNicknames[roomname][winnerID]] = Winner
	gameResult.Pairing[getClashNicknames[roomname][losserID]] = Losser
	gameResult.Pairing[getClashNicknames[roomname][nextTeamWinnerID]] = prevMatchWinner
	gameResult.Pairing[getClashNicknames[roomname][nextTeamLosserID]] = prevMatchLosser

	return gameResult
}

func BroadcastClashWait(h *Hub, roomname, myTeam string) {
	for _, _id := range saveShuffle[roomname][myTeam] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _waiting, roomname: roomname, to: _id, _sleep: false}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _ClashWaitMessgae, roomname: roomname, to: _id, _sleep: false}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: startGame, roomname: roomname, to: _id, _sleep: false}
	}
}

func BoardcastFinalResult(h *Hub, gameResult ParcelClashResult, roomname string) {
	_c, _ := json.Marshal(&gameResult)
	_token := string(_c)
	h.broadcast <- BroadcastReq{Token: startGame, RoomID: roomname}
	h.broadcast <- BroadcastReq{Token: waiting, RoomID: roomname}
	h.broadcast <- BroadcastReq{Token: "ClashGameResult: " + _token, RoomID: roomname}
}

func BroadcastPowerReset(h *Hub, roomname string) {
	_re := ParcelPowerReset{}
	_re.Betting = false
	_re.IsBetting = false
	_re.Unfreeze = true
	_re.UnderTest = false
	_c, _ := json.Marshal(_re)
	token := string(_c)
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ClashPowerRest: " + token}
}

func BroadcastGameReward(h *Hub, WgameReward ClashGameReward, LgameReward ClashGameReward, winnerID, roomname, losserID string, winnerCurrentChance, losserCurrentChance int) {
	go func() {
		_c, _ := json.Marshal(&WgameReward)
		_token := string(_c)
		WgameReward.OnFire = int(getClashProfile[winnerID][roomname].ImOnFire) + 1
		WgameReward.CurrentChance = winnerCurrentChance
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashGameReward: " + _token, roomname: roomname, to: winnerID, _sleep: false}

		LgameReward.OnFire = 0
		LgameReward.CurrentChance = losserCurrentChance
		_c, _ = json.Marshal(&LgameReward)
		_token = string(_c)
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashGameReward: " + _token, roomname: roomname, to: losserID, _sleep: false}
	}()

}
