package server

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

// ClashGame todo make sure to get the teamname of fetched bet from the collection
func ClashGame(h *Hub, conn *websocket.Conn, GR GameRoom, id string) {
	log.Println("clash game")
	roomname := GR.RoomName
	// // @IMPORTANT
	// ID := h.clients[conn].clientID

	myOpponentID := getClashProfile[id][roomname].Against
	myTeam_ := getClashProfile[id][roomname].MyTeam
	MaxPlayerCount := 4
	if GR.FinalBoss || GR.LastDance {
		MaxPlayerCount = 2 // for some reason number 3 doesnt work for challenge becuase i have to add 1 but for else statement it works fine

	} else if GR.DictionarySession || GR.Start {
		MaxPlayerCount = 3
	}

	if GR.PowerActivated {
		log.Println("power activated")
		if GR.Nexus {
			ClashPowerManager(h, true, false, false, false, false, id, myOpponentID, myTeam_, roomname)
		} else if GR.Covert {
			ClashPowerManager(h, false, true, false, false, false, id, myOpponentID, myTeam_, roomname)
		} else if GR.Bet {
			ClashPowerManager(h, false, false, true, false, false, id, myOpponentID, myTeam_, roomname)
		} else if GR.Freeze {
			ClashPowerManager(h, false, false, false, true, false, id, myOpponentID, myTeam_, roomname)
		} else if GR.Rewind {
			ClashPowerManager(h, false, false, false, false, true, id, myOpponentID, myTeam_, roomname)
		}

	} else if GR.TimeUp || GR.DTimeUp {

		switch true {
		case GR.EmptyHand:
			round, winnerID, losserID := GR.Round, getClashProfile[id][roomname].Against, id
			pPoints := ClashEmptyHand(round, winnerID, losserID, roomname)
			ClashGameOver(h, roomname, winnerID, losserID, pPoints, false)
			var lg LocfiyGameInfo
			lg.Session = false
			_c, err := json.Marshal(lg)
			if err != nil {
				log.Println(err)
				return
			}
			token := "ClashGame: " + string(_c)

			h.broadcast <- BroadcastReq{RoomID: roomname, Token: token}
		case !GR.Start && GR.DTimeUp:
			winnerID, losserID := getClashProfile[id][roomname].Against, id
			eliminated, pPoints := ClashRecordTimeup(GR.Round, winnerID, losserID, roomname)
			lastDance := getClashProfile[id][roomname].LastDance

			if lastDance && eliminated {
				ClashGameOver(h, roomname, winnerID, losserID, pPoints, false)
			} else if !lastDance && eliminated {
				EliminatedClashGameOver(h, roomname, winnerID, losserID, pPoints, false)
				h.wg.Go(func() {
					MasterSave(true, 1, losserID, winnerID, losserID,
						roomname, myTeam_, winnerID, _StringSentinel_,
						_StringSentinel_,
						true, false, _IntSentinel, _IntSentinel,
						_StringSentinel_, _StringSentinel_, nil, nil)
				})
			} else if GR.ChallengeDiscussion {
				ClashChallengeSession(h, id, roomname, GR.ChallengeToken, MaxPlayerCount, true)
				pe := map[int]int{GR.Round: 1}
				SingleSaveStats(true, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pe)
			} else if GR.BetSession {
				ClashBetSession(h, id, roomname, GR.BetOn, MaxPlayerCount, true)
				pe := map[int]int{GR.Round: 1}
				SingleSaveStats(true, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pe)
			} else if GR.DictionarySession {
				ClashDictionarySession(h, roomname, id, GR.DictionaryToken, MaxPlayerCount)
				pe := map[int]int{GR.Round: 1}
				SingleSaveStats(true, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pe)
			} else if GR.TossSession {
				ClashToss(h, id, roomname, GR.HeadTails, MaxPlayerCount)
				pe := map[int]int{GR.Round: 1}
				SingleSaveStats(true, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pe)
			}

		case GR.Start && GR.TimeUp:
			winnerID, losserID := getClashProfile[id][roomname].Against, id
			eliminated, pPoints := ClashRecordTimeup(GR.Round, winnerID, losserID, roomname)
			lastDance := getClashProfile[id][roomname].LastDance
			if lastDance && eliminated {
				ClashGameOver(h, roomname, winnerID, losserID, pPoints, false)
			} else if !lastDance && eliminated {
				EliminatedClashGameOver(h, roomname, winnerID, losserID, pPoints, false)
				h.wg.Go(func() {
					MasterSave(true, 1, losserID, winnerID, losserID,
						roomname, myTeam_, winnerID, _StringSentinel_,
						_StringSentinel_,
						true, false, _IntSentinel, _IntSentinel,
						_StringSentinel_, _StringSentinel_, nil, nil)
				})
			} else {
				ClashPlay(h, roomname, id, GR.Guess, GR.Round, GR.Time, true, GR.Chances, GR.OnFire, MaxPlayerCount, false)
				pe := map[int]int{GR.Round: 1}
				SingleSaveStats(true, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pe)

			}
		}

	} else {
		log.Println("power req not triggerd")
		switch true {
		case GR.Session && GR.Unfreeze:
			log.Println("unfreezing")

			var t ParcelFreezePackage
			t.CanUnfreeze = false
			t.FreezeUse = false
			t.Unfreeze = true
			c, _ := json.Marshal(t)
			token := "NYGFreeze: " + string(c)
			// if the requested freeze was already in the pocket of the opponent
			if getClashProfile[id][roomname].OppoPowerUp[_FreezeKey] {
				var t ParcelFreezePackage
				t.CanUnfreeze = false
				t.FreezeUse = false
				t.Unfreeze = false
				c, _ := json.Marshal(t)
				_token_ := "NYGFreeze: " + string(c)
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _token_, to: id, _sleep: false, roomname: roomname}
				DeductPower(h, id, roomname, _FreezeKey)
				RestPower(true, myOpponentID, roomname, _FreezeKey)
			} else {

				for _, _id := range saveShuffle[roomname][myTeam_] {
					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: _id, _sleep: false, roomname: roomname}
				}
			}

		case !GR.Start && GR.TossSession:
			ClashToss(h, id, roomname, GR.HeadTails, MaxPlayerCount)

		case !GR.Start && GR.DictionarySession:
			ClashDictionarySession(h, roomname, id, GR.DictionaryToken, MaxPlayerCount)

		case !GR.Start && GR.ChallengeDiscussion:
			ClashChallengeSession(h, id, roomname, GR.ChallengeToken, MaxPlayerCount, false)

		case !GR.Start && GR.BetSession:
			ClashBetSession(h, id, roomname, GR.BetOn, MaxPlayerCount, false)

		case GR.Start && !GR.Unfreeze:
			ClashPlay(h, roomname, id, GR.Guess, GR.Round,
				GR.Time, GR.TimeUp, GR.Chances, GR.OnFire,
				MaxPlayerCount, GR.LastDance)

		case GR.Start && GR.Spectate:

			if myTeam_ == _TeamRedKey {
				token := RedSpectate + GR.View.Word
				h.broadcast <- BroadcastReq{
					Token: token, RoomID: roomname,
				}

			} else {
				token := BlueSpectate + GR.View.Word
				h.broadcast <- BroadcastReq{
					Token: token, RoomID: roomname,
				}
			}

		case GR.Session && GR.Spectate:
			log.Println("spectating ")
			c, err := json.Marshal(GR.View)
			if err != nil {
				log.Println(err)
				return
			}
			token := View + string(c)
			h.gameRoomBroadcast <- reqGameRoomBroadcast{to: getClashProfile[id][roomname].FinalBossID,
				token: token, _sleep: false, roomname: roomname,
			}

		default:
			log.Println("not able to find play pattern")
		}
	}
}

// ClashRecordDoublyTimeup handles game timeout
func ClashRecordDoublyTimeup(round int, winnerID, losserID, roomname string) (bool, map[string]int) {
	refWinner := getClashProfile[winnerID][roomname]
	refLosser := getClashProfile[losserID][roomname]
	// if not prefered for loop use round-1 because round comes from frontend via round + 1

	refLosserCount := 1 // will always 1
	for _, count := range refLosser.MyPenalties {
		refLosserCount += count
	}

	refWinnerCount := 1
	for _, count := range refWinner.MyPenalties {
		refWinnerCount += count
	}

	if refWinnerCount == 2 && refLosserCount == 2 {
		pPoints := map[string]int{}
		pPoints[winnerID] = getClashProfile[winnerID][roomname].MyPoints
		pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints

		return true, pPoints
	} else if refLosserCount == 2 {
		losserID = refLosser.MyID
		winnerID = refWinner.MyID
		pPoints := map[string]int{}
		pPoints[winnerID] = getClashProfile[winnerID][roomname].MyPoints + 1
		pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints

		return true, pPoints
	} else if refWinnerCount == 2 {

		losserID = refWinner.MyID
		winnerID = refLosser.MyID
		pPoints := map[string]int{}
		pPoints[winnerID] = getClashProfile[winnerID][roomname].MyPoints + 1
		pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints

		return true, pPoints
	}
	pen := map[int]int{round: 1}
	SingleSaveStats(false, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pen)

	return false, nil
}

// ClashRecordTimeup handling for disucssion timeout
func ClashRecordTimeup(round int, winnerID, losserID, roomname string) (bool, map[string]int) {
	log.Println("for round: ", round)
	pen := map[int]int{round: 1}
	cur := 1
	for _, count := range getClashProfile[losserID][roomname].MyPenalties {
		cur += count
	}
	log.Println("current penalty: ", cur)

	if cur == 2 {
		log.Println("game over via timeeout")
		pPoints := map[string]int{}
		pPoints[winnerID] = getClashProfile[winnerID][roomname].MyPoints + 1
		pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints

		return true, pPoints
	}
	SingleSaveStats(true, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pen)

	return false, nil
}

func ClashEmptyHand(round int, winnerID, losserID, roomname string) map[string]int {
	log.Println("for round: ", round)
	pen := map[int]int{}
	pen[round] = 1
	log.Println("game over via timeeout")

	pPoints := map[string]int{}
	pPoints[winnerID] = getClashProfile[winnerID][roomname].MyPoints + 1
	pPoints[losserID] = getClashProfile[losserID][roomname].MyPoints
	return pPoints

}
