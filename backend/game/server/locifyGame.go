package server

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

// LocifyGame todo make sure to get the teamname of fetched bet from the collection
func LocifyGame(h *Hub, conn *websocket.Conn, GR GameRoom, id string) {
	log.Println("Locify game")
	roomname := GR.RoomName

	myOpponentID := getLocifyProfile[id][roomname].Against
	myTeam_ := getLocifyProfile[id][roomname].MyTeam
	log.Println("current round: ", GR.Round)
	if GR.PowerActivated {
		if GR.Nexus {
			LocifyPowerManager(h, true, false, false, false, false, id, myOpponentID, myTeam_, roomname)
		} else if GR.Covert {
			LocifyPowerManager(h, false, true, false, false, false, id, myOpponentID, myTeam_, roomname)
		} else if GR.Bet {
			LocifyPowerManager(h, false, false, true, false, false, id, myOpponentID, myTeam_, roomname)
		} else if GR.Freeze {
			LocifyPowerManager(h, false, false, false, true, false, id, myOpponentID, myTeam_, roomname)
		} else if GR.Rewind {
			LocifyPowerManager(h, false, false, false, false, true, id, myOpponentID, myTeam_, roomname)
		} else if GR.Draw {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				token: drawOffer, _sleep: false,
				to: myOpponentID, roomname: roomname,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{to: id, token: CanUsePower, roomname: roomname, _sleep: false}
		} else if GR.Tag {
			PTagMechanism(h, roomname, id)
		}

	} else if GR.DTimeUp || GR.TimeUp {
		log.Println("time out")
		switch true {
		case !GR.Start && GR.DTimeUp:
			log.Println("disucssion timeout")
			winnerID, losserID := myOpponentID, id

			gameOver, pPoints := LocifyRecordTimeup(GR.Round, winnerID, losserID, roomname)
			if gameOver {
				r := getLocifyRoomSettings[roomname]
				LocifyGameOver(h, roomname, myOpponentID, id, pPoints, r, false)
				var lg LocfiyGameInfo
				lg.Session = false
				_c, err := json.Marshal(lg)
				if err != nil {
					log.Println(err)
					return
				}
				token := "LocifyGame: " + string(_c)

				h.broadcast <- BroadcastReq{RoomID: roomname, Token: token}
			} else if GR.ChallengeDiscussion {
				log.Println("yes")
				LocifyChallengeSession(h, id, roomname, GR.ChallengeToken, true)
			} else if GR.BetSession {
				LocifyBetSession(h, id, roomname, GR.BetOn, false)
			} else if GR.DictionarySession {
				LocifyDictionarySession(h, roomname, id, GR.DictionaryToken)
			} else if GR.TossSession {
				LocifyToss(h, id, roomname, GR.HeadTails)
			}

		case GR.Start && GR.TimeUp:
			winnerID, losserID := myOpponentID, id
			gameOver, pPoints := LocifyRecordDoublyTimeup(GR.Round, winnerID, losserID, roomname)
			if gameOver {
				r := getLocifyRoomSettings[roomname]
				LocifyGameOver(h, roomname, myOpponentID, id, pPoints, r, false)
				var lg LocfiyGameInfo
				lg.Session = false
				_c, err := json.Marshal(lg)
				if err != nil {
					log.Println(err)
					return
				}
				token := "LocifyGame: " + string(_c)

				h.broadcast <- BroadcastReq{RoomID: roomname, Token: token}
			} else {
				LocifyPlay(h, roomname, id, GR.Guess, GR.Round,
					GR.Set, GR.Time, GR.RedScoreCount, GR.BlueScoreCount, GR.TimeUp)
			}
		}
	} else {

		switch true {
		case GR.EmptyHand:
			winnerID, losserID := getLocifyProfile[id][roomname].Against, id
			pPoints := LocifyEmptyHand(GR.Round, winnerID, losserID, roomname)

			r := getLocifyRoomSettings[roomname]
			LocifyGameOver(h, roomname, myOpponentID, id, pPoints, r, true)
			var lg LocfiyGameInfo
			lg.Session = false
			_c, err := json.Marshal(lg)
			if err != nil {
				log.Println(err)
				return
			}
			token := "LocifyGame: " + string(_c)

			h.broadcast <- BroadcastReq{RoomID: roomname, Token: token}

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
			h.gameRoomBroadcast <- reqGameRoomBroadcast{to: getLocifyProfile[id][roomname].MyPatnerID,
				token: token, _sleep: false, roomname: roomname,
			}

		case GR.DrawSession:
			PDrawMechanism(h, roomname, id, GR.DrawAccept)

		case GR.Session && GR.Unfreeze:
			log.Println("unfreezing")

			var t ParcelFreezePackage
			t.CanUnfreeze = false
			t.FreezeUse = false
			t.Unfreeze = true
			c, _ := json.Marshal(t)
			token := "NYGFreeze: " + string(c)

			h.broadcast <- BroadcastReq{RoomID: roomname, Token: token}
		case !GR.Start && GR.TossSession:
			LocifyToss(h, id, roomname, GR.HeadTails)

		case !GR.Start && GR.DictionarySession:
			LocifyDictionarySession(h, roomname, id, GR.DictionaryToken)

		case !GR.Start && GR.ChallengeDiscussion:
			LocifyChallengeSession(h, id, roomname, GR.ChallengeToken, false)

		case !GR.Start && GR.BetSession:
			LocifyBetSession(h, id, roomname, GR.BetOn, false)

		case GR.Start && !GR.Unfreeze:
			LocifyPlay(h, roomname, id, GR.Guess, GR.Round,
				GR.Set, GR.Time, GR.RedScoreCount, GR.BlueScoreCount, GR.TimeUp)

		default:
			log.Println("not able to find play pattern")
		}
	}
}

// LocifyRecordDoublyTimeup handles game timeout
func LocifyRecordDoublyTimeup(round int, winnerID, losserID, roomname string) (bool, map[string]int) {
	refWinner := getLocifyProfile[winnerID][roomname]
	refLosser := getLocifyProfile[losserID][roomname]
	// if not prefered for loop use round-1 because round comes from frontend via round + 1

	refLosserCount := 1
	for _, count := range refLosser.MyPenalties {
		refLosserCount += count
	}

	refWinnerCount := 1
	for _, count := range refWinner.MyPenalties {
		refWinnerCount += count
	}

	if refWinnerCount == 2 && refLosserCount == 2 {
		pPoints := map[string]int{}
		pPoints[winnerID] = getLocifyProfile[winnerID][roomname].MyPoints
		pPoints[losserID] = getLocifyProfile[losserID][roomname].MyPoints
		r := getLocifyRoomSettings[roomname]
		if r.Capacity == 4 {
			wl := getLocifyProfile[getLocifyProfile[winnerID][roomname].MyPatnerID][roomname]
			ll := getLocifyProfile[getLocifyProfile[losserID][roomname].MyPatnerID][roomname]
			pPoints[wl.MyID] = wl.MyPoints + 1
			pPoints[ll.MyID] = ll.MyPoints + 1
		}
		return true, pPoints
	} else if refLosserCount == 2 {
		losserID = refLosser.MyID
		winnerID = refWinner.MyID
		pPoints := map[string]int{}
		pPoints[winnerID] = getLocifyProfile[winnerID][roomname].MyPoints + 1
		pPoints[losserID] = getLocifyProfile[losserID][roomname].MyPoints
		r := getLocifyRoomSettings[roomname]
		if r.Capacity == 4 {
			wl := getLocifyProfile[getLocifyProfile[winnerID][roomname].MyPatnerID][roomname]
			ll := getLocifyProfile[getLocifyProfile[losserID][roomname].MyPatnerID][roomname]
			pPoints[wl.MyID] = wl.MyPoints + 1
			pPoints[ll.MyID] = ll.MyPoints + 1
		}
		return true, pPoints
	} else if refWinnerCount == 2 {

		losserID = refWinner.MyID
		winnerID = refLosser.MyID
		pPoints := map[string]int{}
		pPoints[winnerID] = getLocifyProfile[winnerID][roomname].MyPoints + 1
		pPoints[losserID] = getLocifyProfile[losserID][roomname].MyPoints
		r := getLocifyRoomSettings[roomname]
		if r.Capacity == 4 {
			wl := getLocifyProfile[getLocifyProfile[winnerID][roomname].MyPatnerID][roomname]
			ll := getLocifyProfile[getLocifyProfile[losserID][roomname].MyPatnerID][roomname]
			pPoints[wl.MyID] = wl.MyPoints + 1
			pPoints[ll.MyID] = ll.MyPoints + 1
		}
		return true, pPoints
	}
	pen := map[int]int{round: 1}
	SingleSaveStats(false, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pen)

	return false, nil
}

// LocifyRecordTimeup handling for disucssion timeout
func LocifyRecordTimeup(round int, winnerID, losserID, roomname string) (bool, map[string]int) {
	log.Println("for round: ", round)

	pen := map[int]int{}
	pen[round] = 1

	cur := 1 // for current
	for _, count := range getLocifyProfile[losserID][roomname].MyPenalties {
		cur += count
	}

	log.Println("current penalty: ", cur)

	if cur == 2 {
		log.Println("game over via timeeout")
		pPoints := map[string]int{}
		pPoints[winnerID] = getLocifyProfile[winnerID][roomname].MyPoints + 1
		pPoints[losserID] = getLocifyProfile[losserID][roomname].MyPoints
		r := getLocifyRoomSettings[roomname]
		if r.Capacity == 4 {
			wl := getLocifyProfile[getLocifyProfile[winnerID][roomname].MyPatnerID][roomname].MyPoints + 1
			ll := getLocifyProfile[getLocifyProfile[losserID][roomname].MyPatnerID][roomname].MyPoints + 1
			pPoints[getLocifyProfile[losserID][roomname].MyPatnerID] = ll
			pPoints[getLocifyProfile[winnerID][roomname].MyPatnerID] = wl
		}
		return true, pPoints
	}
	SingleSaveStats(false, losserID, roomname, _IntSentinel, _IntSentinel, _StringSentinel_, _IntSentinel, _StringSentinel_, pen)

	return false, nil
}

func LocifyEmptyHand(round int, winnerID, losserID, roomname string) map[string]int {
	log.Println("for round: ", round)

	log.Println("game over via timeeout")
	pPoints := map[string]int{}
	pPoints[winnerID] = getLocifyProfile[winnerID][roomname].MyPoints + 1
	pPoints[losserID] = getLocifyProfile[losserID][roomname].MyPoints + 1
	r := getLocifyRoomSettings[roomname]
	if r.Capacity == 4 {
		wl := getLocifyProfile[getLocifyProfile[winnerID][roomname].MyPatnerID][roomname].MyPoints + 1
		ll := getLocifyProfile[getLocifyProfile[losserID][roomname].MyPatnerID][roomname].MyPoints + 1
		pPoints[getLocifyProfile[losserID][roomname].MyPatnerID] = ll
		pPoints[getLocifyProfile[winnerID][roomname].MyPatnerID] = wl
	}
	return pPoints

}
