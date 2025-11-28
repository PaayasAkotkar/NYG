package server

import (
	"encoding/json"
	"log"
	"nyg/validate"
	"strings"
)

// ClashPlay returns
// display the result and wait for other players to finish the game
// final boss mechanism:
// cases->first down
// cases->final boss found
// case->first down: if the current request resulted into one player down than the other player will be lock as the final boss
// rest of the players will be lock in clash 1v1
// case ->first down after: if the final boss is already found -> if the current request resulted into one player down
// than the other player and final boss will be locked in clash 1v1
// for clash max player count is 3
// note: during last dance make sure to broadcast imboss to false
// single stats point
// knockedOutPoints is the point that winner has deducted from his opponent
// note: only the prev winner and present winner points will be addded in the profile
func ClashPlay(h *Hub, roomname string, id string, guessToken string, round int, Time int, TimeUp bool, CurrentChance int, OnFire int, MaxPlayerCounnt int, lastDance bool) {
	log.Println("clash play")

	myProfile := getClashProfile[id][roomname]
	Rewind, rID := IsRewind(id, roomname, guessToken, true)

	// nextTeamWinnerID := myProfile.NextTeamWinner
	// nextTeamLosserID := myProfile.NextTeamLosser
	winnerID := ""
	losserID := ""
	myTeam := myProfile.MyTeam
	myOpponentID := myProfile.Against
	myOpponentSetChallenge := myProfile.OppoSetChallenge
	myGameDictionary := myProfile.SetDictionary
	point := false
	myGuess := myProfile.MyGuess
	Count := int(1)
	deduce := 0
	sessionDone := (myProfile.Count + Count) == int(MaxPlayerCounnt)

	WgameReward := ClashGameReward{}
	LgameReward := ClashGameReward{}

	finalBoss := myProfile.FinalBossFound

	nextTeam := myProfile.NextTeamname
	nWinner := myProfile.NextTeamWinner
	nLosser := myProfile.NextTeamLosser

	eliminatedIDs := myProfile.EliminatedPlayersIDs

	// for searching final boss
	losserCurrentChance := 0
	winnerCurrentChance := CurrentChance

	eliminated := false

	log.Println("my profile: ", myProfile)

	// log.Println("game info: ", getClashGameInfo)
	dict, list := myGameDictionary, myOpponentSetChallenge
	_isBet := myProfile.OppoPowerUp[_BetKey]
	if _isBet {
		point = strings.Contains(myProfile.OppoSetBet, guessToken)
		if point {
			// last check for incorrect
			for _, bet := range myProfile.OppoBets {
				point = !strings.Contains(bet, guessToken)
			}
		}
	} else {
		b := myProfile.Book
		point, _ = validate.SportsValidate(b, dict, list, guessToken, myGuess)
	}

	if point {
		// on fire will be set to 0 of the opp if true and dedcue will be on the basis of the formula+onFire
		log.Println("correct guess")

		// onFire mechanism
		// if false than it will be deduce by constant rather than deduce formula and onFire to be set to 0
		winnerID = id
		losserID = myOpponentID
		deduce = DeductMec(winnerID, roomname, Time)

		losserCurrentChance = deduce
		winnerCurrentChance = CurrentChance

		WgameReward.CurrentChance = winnerCurrentChance
		WgameReward.OnFire = getClashProfile[winnerID][roomname].ImOnFire + 1

		LgameReward.CurrentChance = losserCurrentChance
		LgameReward.OnFire = getClashProfile[winnerID][roomname].ImOnFire + 1

		h.wg.Go(func() {
			SingleSaveStats(true, winnerID, roomname, winnerCurrentChance, 1, guessToken, losserCurrentChance, _StringSentinel_, nil)
		})
		h.wg.Go(func() {
			SingleSaveStats(true, losserID, roomname, losserCurrentChance, 0, "", 0, _StringSentinel_, nil) // reset the on fire
		})
	} else {
		log.Println("incorrect guess")

		winnerID = myOpponentID
		losserID = id

		deduce = CurrentChance - Min(2, CurrentChance)

		losserCurrentChance = deduce
		winnerCurrentChance = getClashProfile[myOpponentID][roomname].MyCurrentChances

		WgameReward.CurrentChance = winnerCurrentChance
		WgameReward.OnFire = getClashProfile[winnerID][roomname].ImOnFire + 1

		LgameReward.CurrentChance = losserCurrentChance
		LgameReward.OnFire = getClashProfile[winnerID][roomname].ImOnFire + 1

		h.wg.Go(func() {
			SingleSaveStats(true, winnerID, roomname, winnerCurrentChance, 1, "", losserCurrentChance, _StringSentinel_, nil)
		})
		h.wg.Go(func() {
			SingleSaveStats(true, losserID, roomname, losserCurrentChance, 0, guessToken, 0, _StringSentinel_, nil) // reset the on fire
		})
	}

	switch true {
	// priority will always be rewind
	case Rewind:
		// trigger it
		PRewindMechanism(h, roomname, rID, myTeam, true)
		ClashResetPowerUp(rID, roomname, _RewindKey)
		SingleSaveStats(true, id, roomname, _IntSentinel, _IntSentinel, guessToken, _IntSentinel, _StringSentinel_, nil)
		SingleSaveStats(true, myProfile.Against, roomname, _IntSentinel, _IntSentinel, guessToken, _IntSentinel, _StringSentinel_, nil)

	case lastDance:
		log.Println("last dance")

		if losserCurrentChance == 0 {
			eliminated = true
			var j = PlayerElimatedMessage{}
			j.Clash = false
			j.GameBegin = false
			j.Session = false
			j.StartGame = false
			_s, _ := json.Marshal(j)
			token := "Eliminated: " + string(_s)
			pPoints := map[string]int{}
			pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints
			pPoints[winnerID] = getClashProfile[winnerID][roomname].MyPoints

			ClashGameOver(h, roomname, winnerID, losserID, pPoints, false)

			h.broadcast <- BroadcastReq{RoomID: roomname, Token: token}
			h.broadcast <- BroadcastReq{RoomID: roomname, Token: _Clash}
			BoardcastRefresh(h, roomname)

		} else {
			matchups := Clash1v1(h, false, eliminatedIDs, roomname,
				myProfile.FinalBossID, winnerID, losserID, winnerCurrentChance,
				losserCurrentChance, round)

			// keep broadcasting result
			_result := CalcResult(roomname, myTeam,
				winnerID, losserID,
				round, winnerCurrentChance,
				losserCurrentChance,
				myProfile, true)

			go func() {
				BroadcastPowerReset(h, roomname)
			}()
			BoardcastFinalResult(h, _result, roomname)

			go func() {
				storeShuffle.store <- matchups
			}()

		}

		j := ParcelMatchupUpdate{}
		j.FinalBoss = false
		j.LastDance = true
		j.IMBoss = false
		a, _ := json.Marshal(j)
		_c := string(a)

		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashMatchUpdate: " + _c, roomname: roomname, to: winnerID, _sleep: false}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashMatchUpdate: " + _c, roomname: roomname, to: losserID, _sleep: false}

		BroadcastGameReward(h, WgameReward, LgameReward, winnerID, losserID, roomname, winnerCurrentChance, losserCurrentChance)
		h.broadcast <- BroadcastReq{RoomID: roomname, Token: _CanUsePower}
		BoardcastRefresh(h, roomname)

	case finalBoss:
		// note: final boss is till the one of the player is not down
		// while the process is same
		// flow->
		// check for final boos
		// true-> distribute res
		// false -> distrbute res -> keep looking for final boss -> found store it
		log.Println("final boss")

		if losserCurrentChance == 0 {
			log.Println("final game losser count to 0")
			eliminated = true
			finalBoss = false
			lastDance = true

			var _j = PlayerElimatedMessage{}
			_j.Clash = false
			_j.GameBegin = false
			_j.Session = false
			_j.StartGame = false
			_s, _ := json.Marshal(_j)
			token := "Eliminated: " + string(_s)

			h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname,
				token: token, to: losserID, _sleep: false}

			eliminatedIDs = append(eliminatedIDs, losserID)

			red := winnerID
			blue := myProfile.FinalBossID

			bossCurrentChance := getClashProfile[blue][roomname].MyCurrentChances

			matchups := Clash1v1(h, true,
				eliminatedIDs,
				roomname, _StringSentinel_,
				red, blue, winnerCurrentChance,
				bossCurrentChance, round)

			go func() {
				log.Println("unregistering")
				var c = ClientIDUnregistration{}
				c.ID = losserID
				c.Roomname = roomname
				h.unregisterID <- c
			}()
			j := ParcelMatchupUpdate{}
			j.FinalBoss = finalBoss
			j.LastDance = lastDance
			j.IMBoss = false

			a, _ := json.Marshal(j)
			_c := string(a)

			go func() {
				storeShuffle.store <- matchups
			}()
			pPoints := map[string]int{}
			pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints

			EliminatedClashGameOver(h, roomname, winnerID, losserID, pPoints, false)

			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashMatchUpdate: " + _c, roomname: roomname, to: blue, _sleep: false}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _Clash, roomname: roomname, to: losserID, _sleep: false}

		} else {
			finalBoss = true
			lastDance = false

			matchups := Clash1v1(h, false,
				eliminatedIDs, roomname,
				myProfile.FinalBossID,
				winnerID, losserID,
				winnerCurrentChance, losserCurrentChance, round)

			go func() {
				storeShuffle.store <- matchups
			}()

		}

		j := ParcelMatchupUpdate{}
		j.FinalBoss = finalBoss
		j.LastDance = lastDance
		j.IMBoss = false

		a, _ := json.Marshal(j)
		_c := string(a)

		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashMatchUpdate: " + _c, roomname: roomname, to: winnerID, _sleep: false}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashMatchUpdate: " + _c, roomname: roomname, to: losserID, _sleep: false}

		_result := CalcClash3wayResults(roomname, myTeam, winnerID, losserID,
			round, winnerCurrentChance, losserCurrentChance, myProfile)

		go func() {
			BoardcastFinalResult(h, _result, roomname)
		}()

		go func() {
			BroadcastPowerReset(h, roomname)
		}()

		go func() {
			re := map[string]bool{}
			re[roomname] = true
			clashResetCount <- re
		}()

		BroadcastGameReward(h, WgameReward, LgameReward, winnerID, roomname, losserID, winnerCurrentChance, losserCurrentChance)
		h.broadcast <- BroadcastReq{RoomID: roomname, Token: _CanUsePower}
		BoardcastRefresh(h, roomname)

		// case before session done->
		// save the winner if loser eliminate and consider winner as final boss
		// case after session done->
		// note: no need to check for the final boss
		// save the winner if loser eliminate and consider winner as final boss
	default:

		log.Println("normal game")

		if sessionDone {
			log.Println("session done")

			_currentChances := map[string]int{}

			for _, _id := range saveShuffle[roomname][myProfile.NextTeamname] {
				if _, ok := _currentChances[_id]; !ok {
					_currentChances[_id] = int(getClashProfile[_id][roomname].MyCurrentChances)
				}
			}

			_currentChances[winnerID] = winnerCurrentChance
			_currentChances[losserID] = losserCurrentChance

			matchups := map[string]map[string][]string{}
			_create := map[string]map[string]ClashFixtures{}

			// if the sesssion done
			// now the prev match up will be the game
			if losserCurrentChance == 0 {
				eliminated = true
				finalBoss = true

				var _j = PlayerElimatedMessage{}
				_j.Clash = false
				_j.GameBegin = false
				_j.Session = false
				_j.StartGame = false
				_s, _ := json.Marshal(_j)
				token := "Eliminated: " + string(_s)

				h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, token: token, to: losserID, _sleep: false}

				prevLosser := getClashProfile[myProfile.NextTeamLosser][roomname]
				prevWinner := getClashProfile[myProfile.NextTeamWinner][roomname]

				matchups = Clash1v1(h, true, eliminatedIDs, roomname,
					myProfile.FinalBossID,
					prevWinner.MyID, prevLosser.MyID, prevWinner.MyCurrentChances,
					prevLosser.MyCurrentChances, round)

				pPoints := map[string]int{}
				pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints

				EliminatedClashGameOver(h, roomname, winnerID, losserID, pPoints, false)

				// finalBoss := BroadcastPlayerEliminated(h, winnerID, losserID, roomname, myTeam, true)
				// finalBoss := map[string]string{roomname: winnerID}
				// storeFinalBoss <- finalBoss

				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _Clash, roomname: roomname, to: losserID, _sleep: false}

				go func() {
					var c = ClientIDUnregistration{}
					c.ID = losserID
					c.Roomname = roomname
					h.unregisterID <- c
				}()

				MasterSave(true, _IntSentinel, losserID, _StringSentinel_, _StringSentinel_,
					roomname, _StringSentinel_, winnerID, _StringSentinel_,
					_StringSentinel_,
					true, false,
					_IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, nil, nil)

				go func() {
					storeShuffle.store <- matchups
				}()

				// ClashElimationMessage(h, roomname, losserID)
			} else {
				finalBoss = false
				matchups, _create = ClashMatchUp(h, round, roomname, saveShuffle, _currentChances, getClashNicknames)

				MasterSave(true, _IntSentinel, _StringSentinel_, _StringSentinel_, _StringSentinel_,
					roomname, _StringSentinel_, _StringSentinel_, _StringSentinel_,
					_StringSentinel_, false, false, _IntSentinel, _IntSentinel,
					_StringSentinel_, _StringSentinel_, nil, nil)
				go func() {
					storeShuffle.store <- matchups
				}()
				h.broadcast <- BroadcastReq{RoomID: roomname, Token: _DictionaryDiscussion}
			}

			j := ParcelMatchupUpdate{}
			j.FinalBoss = finalBoss
			j.LastDance = false
			a, _ := json.Marshal(j)
			_c := string(a)

			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashMatchUpdate: " + _c, roomname: roomname, to: winnerID, _sleep: false}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashMatchUpdate: " + _c, roomname: roomname, to: losserID, _sleep: false}

			_result := CalcClashResults(roomname, myTeam, winnerID, losserID,
				nLosser, nWinner, nextTeam, round, winnerCurrentChance,
				losserCurrentChance, matchups)

			go func() {
				BoardcastFinalResult(h, _result, roomname)
			}()

			go func() {
				BroadcastPowerReset(h, roomname)
			}()

			go func() {
				re := map[string]bool{}
				re[roomname] = true
				clashResetCount <- re
			}()

			go func() {
				createProfile <- _create
			}()

		} else {
			log.Println("waiting fo other")

			if losserCurrentChance == 0 {
				log.Println("player eliminiated")
				eliminated = true
				var _j = PlayerElimatedMessage{}
				_j.Clash = false
				_j.GameBegin = false
				_j.Session = false
				_j.StartGame = false
				_s, _ := json.Marshal(_j)
				token := "Eliminated: " + string(_s)

				h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, token: token, to: losserID, _sleep: false}

				// finalBoss := map[string]string{}
				// finalBoss[roomname] = winnerID

				// storeFinalBoss <- finalBoss

				j := ParcelMatchupUpdate{}
				j.FinalBoss = true
				j.LastDance = false
				j.IMBoss = true
				a, _ := json.Marshal(j)
				_c := string(a)

				h.wg.Go(func() {
					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashMatchUpdate: " + _c, roomname: roomname, to: winnerID, _sleep: false}
				})
				h.wg.Go(
					func() {
						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _Clash, roomname: roomname, to: losserID, _sleep: false}
					})
				pPoints := map[string]int{}
				pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints

				EliminatedClashGameOver(h, roomname, winnerID, losserID, pPoints, false)

				h.wg.Go(func() {
					MasterSave(true, int(Count), losserID, winnerID, losserID,
						roomname, myTeam, winnerID, _StringSentinel_,
						_StringSentinel_,
						true, false, _IntSentinel, _IntSentinel,
						_StringSentinel_, _StringSentinel_, nil, nil)
				})

			} else {

				h.wg.Go(func() {
					MasterSave(true, int(Count), _StringSentinel_, winnerID, losserID,
						roomname, myTeam, _StringSentinel_, _StringSentinel_,
						_StringSentinel_,
						false, false, _IntSentinel,
						_IntSentinel, _StringSentinel_, _StringSentinel_, nil, nil)
				})
			}
			BroadcastClashWait(h, roomname, myTeam)

		}
		// for winner only send onFire

		BroadcastGameReward(h, WgameReward, LgameReward, winnerID, roomname, losserID, winnerCurrentChance, losserCurrentChance)
		// dedcut the power
		for power, ok := range myProfile.OppoPowerUp {
			if ok {
				DeductPower(h, myOpponentID, roomname, power)
			}
		}
		for power, ok := range getClashProfile[myOpponentID][roomname].OppoPowerUp {
			if ok {
				DeductPower(h, id, roomname, power)
			}
		}
		BoardcastRefresh(h, roomname)

	}
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: _CanUsePower}

	log.Println("next team: ", myProfile.NextTeamname, "losserCurrentChance: ", getClashProfile[myProfile.NextTeamLosser][roomname].MyCurrentChances, "winnerCurrentChance: ", getClashProfile[myProfile.NextTeamWinner][roomname].MyCurrentChances)
	log.Println("myteam:", myTeam, "losserCurrentChance: ", losserCurrentChance, "winnerCurrentChance: ", winnerCurrentChance)
	// do this process at the end
	if eliminated {

		h.wg.Go(func() {
			log.Println("unregisering begnis")
			var c = ClientIDUnregistration{}
			c.ID = losserID
			c.Roomname = roomname
			h.unregisterID <- c

		})

	}
}
