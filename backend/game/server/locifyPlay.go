package server

import (
	"encoding/json"
	"log"
	"maps"
	"nyg/dataset"
	"nyg/validate"
	"strings"
)

// LocifyPlay todo remove if the if else statement from the blockmanger and place it outside
func LocifyPlay(h *Hub, roomname string, id string, guessToken string, round, set, time, redScore, blueScore int, TimeUp bool) {
	log.Println("locify play")
	log.Println("round: ", round)

	myProfile := getLocifyProfile[id][roomname]
	Rewind, rID := IsRewind(id, roomname, guessToken, false)

	// winnerID, losserID := "", ""
	myTeam := myProfile.MyTeam
	// myOpponentID := myProfile.Against
	myOpponentSetChallenge := myProfile.OppoSetChallenge
	myGameDictionary := myProfile.SetDictionary
	point := false
	lists := myProfile.WholeGuess

	newList := []string{}
	dict, list := myGameDictionary, myOpponentSetChallenge
	_isBet := myProfile.OppoPowerUp[_BetKey]
	log.Println("bets: ", myProfile.OppoBets)

	if !TimeUp {
		winnerTeamname, losserTeamname := "", ""
		winnerID, losserID := "", ""
		RedScore, BlueScore := 0, 0
		if _isBet {

			// check if the bet is match the token
			for _, bet := range myProfile.OppoBets {
				if strings.Contains(guessToken, bet) || strings.HasPrefix(guessToken, bet) {
					point = true
				}
			}
			// if the curr not bet
			point = !strings.Contains(myProfile.OppoSetBet, guessToken)

			newList = append(newList, guessToken)
			var x Frame
			x.Frames = myProfile.OppoBets
			c, err := json.Marshal(x)
			if err != nil {
				log.Println(err)
				return
			}
			token := Frames + string(c)
			h.broadcast <- BroadcastReq{Token: token, RoomID: roomname}

		} else {
			point_, frames := validate.SportsValidate(getLocifyRoomSettings[roomname].Book, dict, list, guessToken, lists)
			point = point_
			newList = append(newList, frames...)
			a := []string{}
			a = append(a, newList...)
			a = append(a, myProfile.WholeGuess...)
			a = append(a, myProfile.OppoBets...)
			a = dataset.EraseDuplicate(a)

			log.Println("frames:", frames)
			var x Frame
			x.Frames = a
			c, err := json.Marshal(x)
			if err != nil {
				log.Println(err)
				return
			}
			token := Frames + string(c)
			h.broadcast <- BroadcastReq{Token: token, RoomID: roomname}
		}

		if Rewind {
			log.Println("rewind the process")
			PRewindMechanism(h, roomname, rID, myTeam, false)
			LocifyResetPowerUp(rID, roomname, _RewindKey)
			SingleSaveStats(false, id, roomname, _IntSentinel, _IntSentinel, guessToken, _IntSentinel, _StringSentinel_, nil)
			SingleSaveStats(false, myProfile.Against, roomname, _IntSentinel, _IntSentinel, guessToken, _IntSentinel, _StringSentinel_, nil)
		} else {
			switch point {
			case true:
				log.Println("correct guess")
				if myProfile.MyTeam == _TeamRedKey {
					RedScore = 1
					BlueScore = 0
				} else {
					BlueScore = 1
					RedScore = 0
				}
				winnerID = id
				losserID = myProfile.Against
				winnerTeamname = myTeam
				losserTeamname = myProfile.OppoTeamname
			case false:
				log.Println("incorrect guess")

				if myProfile.MyTeam == _TeamRedKey {
					RedScore = 0
					BlueScore = 1
				} else {
					RedScore = 1
					BlueScore = 0
				}
				winnerID = myProfile.Against
				losserID = id
				winnerTeamname = myProfile.OppoTeamname
				losserTeamname = myTeam
			}

			redCurrentScore := myProfile.RedTeamScore + RedScore
			bluCurrentScore := myProfile.BlueTeamScore + BlueScore
			wi := getLocifyProfile[winnerID][roomname]
			li := getLocifyProfile[losserID][roomname]

			go func() {
				MasterSave(false, _IntSentinel, _StringSentinel_,
					_StringSentinel_, _StringSentinel_, roomname, _StringSentinel_,
					_StringSentinel_, _StringSentinel_,
					_StringSentinel_, false, false,
					RedScore,
					BlueScore, _StringSentinel_, _StringSentinel_, nil, newList)
			}()

			lg := LocifyGameResultSummup(roomname, winnerID, losserID, myProfile, wi, li, redCurrentScore, bluCurrentScore, round)

			gameDone := (redScore+RedScore) >= 3 || (blueScore+BlueScore) >= 3 || round > 6
			log.Println(gameDone)

			switch gameDone {
			case true:
				lg.Session = false
				pPoints := map[string]int{}
				pPoints[winnerID] = wi.MyPoints + 1
				pPoints[wi.IdlePlayer] = getLocifyProfile[wi.IdlePlayer][roomname].MyPoints + 1
				pPoints[losserID] = li.MyPoints

				if myProfile.RoomSettings.Capacity > 2 {
					wl := getLocifyProfile[wi.IdlePlayer][roomname]
					ll := getLocifyProfile[li.IdlePlayer][roomname]
					pPoints[wl.MyID] = wi.MyPoints + 1
					pPoints[ll.MyID] = li.MyPoints
				}
				LocifyGameOver(h, roomname, winnerID, losserID, pPoints, myProfile.RoomSettings, false)

				_c, err := json.Marshal(lg)
				if err != nil {
					log.Println(err)
					return
				}
				token := "LocifyGame: " + string(_c)

				h.broadcast <- BroadcastReq{RoomID: roomname, Token: token}

			case false:
				BoardcastRefresh(h, roomname)

				BoardcastLocifyRoundResult(h, roomname, winnerID, losserID, winnerTeamname, losserTeamname, guessToken, lg, round, redCurrentScore, bluCurrentScore, myProfile, wi, li)
			}

		}
	} else {
		redCurrentScore := myProfile.RedTeamScore + 0
		bluCurrentScore := myProfile.BlueTeamScore + 0
		wi := getLocifyProfile[id][roomname]
		li := getLocifyProfile[myProfile.Against][roomname]

		go func() {
			MasterSave(false, _IntSentinel, _StringSentinel_,
				_StringSentinel_, _StringSentinel_, roomname, _StringSentinel_,
				_StringSentinel_, _StringSentinel_,
				_StringSentinel_, false, false,
				0,
				0, _StringSentinel_, _StringSentinel_, nil, newList)
		}()

		lg := LocifyGameResultSummup(roomname, id, myProfile.Against, myProfile, wi, li, redCurrentScore, bluCurrentScore, round)

		gameDone := (redScore) >= 3 || (blueScore) >= 3 || round > 6
		log.Println(gameDone)

		switch gameDone {
		case true:
			lg.Session = false
			pPoints := map[string]int{}
			pPoints[id] = wi.MyPoints
			pPoints[wi.IdlePlayer] = getLocifyProfile[wi.IdlePlayer][roomname].MyPoints
			pPoints[myProfile.Against] = li.MyPoints

			if myProfile.RoomSettings.Capacity > 2 {
				wl := getLocifyProfile[wi.IdlePlayer][roomname]
				ll := getLocifyProfile[li.IdlePlayer][roomname]
				pPoints[wl.MyID] = wi.MyPoints
				pPoints[ll.MyID] = li.MyPoints
			}
			LocifyGameOver(h, roomname, id, myProfile.Against, pPoints, myProfile.RoomSettings, true)

			_c, err := json.Marshal(lg)
			if err != nil {
				log.Println(err)
				return
			}
			token := "LocifyGame: " + string(_c)

			h.broadcast <- BroadcastReq{RoomID: roomname, Token: token}
			h.broadcast <- BroadcastReq{RoomID: roomname, Token: _CanUsePower}

		case false:

			BoardcastLocifyRoundResult(h, roomname, id, myProfile.Against, myTeam, myProfile.OppoTeamname, guessToken, lg, round, redCurrentScore, bluCurrentScore, myProfile, wi, li)
			BoardcastRefresh(h, roomname)
		}

	}
}

// LocifyGameResultSummup wi is winnerIdle partner and li is LosserIdle partner
func LocifyGameResultSummup(roomname, winnerID, losserID string, myProfile, wi, li LocifyFixtures, redCurrentScore, blueCurrentScore, round int) LocfiyGameInfo {
	var lg LocfiyGameInfo
	lg.PlayersStats = make(map[string]map[int]string)

	switch round {
	case 3:
		lg.Set = 2
	case 5:
		lg.Set = 3
	default:
		lg.Set = 1
	}

	if round != 2 {

		lg.PlayersStats = make(map[string]map[int]string)

		x := make(map[string]map[int]string)
		r1, r2 := make(map[int]string), make(map[int]string)

		// add the prev round
		maps.Copy(r1, wi.MySheet)
		maps.Copy(r2, li.MySheet)

		// add the current round
		r1[round-1] = "1" // winner
		r2[round-1] = "0" // losser

		x[myProfile.NickNamesViaID[winnerID]] = r1
		x[myProfile.NickNamesViaID[losserID]] = r2

		if myProfile.RoomSettings.Capacity > 2 {

			wl := getLocifyProfile[wi.IdlePlayer][roomname]
			ll := getLocifyProfile[li.IdlePlayer][roomname]

			var r3, r4 = make(map[int]string), make(map[int]string)

			// add the prev round
			maps.Copy(r3, wi.MySheet)
			maps.Copy(r4, li.MySheet)

			// add the current round
			r2[round-1] = "l" // winner
			r4[round-1] = "l" // losser

			x[wl.MyID] = r3
			x[ll.MyID] = r4
		}
		lg.PlayersStats = x

	} else {
		wi := getLocifyProfile[winnerID][roomname]
		li := getLocifyProfile[losserID][roomname]
		x := make(map[string]map[int]string)
		r1, r2 := make(map[int]string), make(map[int]string)

		// add the current round
		r1[round-1] = "1" // winner
		r2[round-1] = "0" // losser

		x[myProfile.NickNamesViaID[winnerID]] = r1

		x[myProfile.NickNamesViaID[losserID]] = r2

		if myProfile.RoomSettings.Capacity > 2 {

			wl := getLocifyProfile[wi.IdlePlayer][roomname]
			ll := getLocifyProfile[li.IdlePlayer][roomname]
			r3, r4 := make(map[int]string), make(map[int]string)

			// add the current round
			r3[round-1] = "l" // winner
			r4[round-1] = "l" // losser

			x[myProfile.NickNamesViaID[wl.MyID]] = r3
			x[myProfile.NickNamesViaID[ll.MyID]] = r4

		}
		lg.PlayersStats = x
	}

	lg.Session = true
	lg.Round = round // because from front-end round+1
	lg.BlueScore = blueCurrentScore
	lg.RedScore = redCurrentScore
	lg.Mode = LocifyMode
	lg.Roomname = roomname

	var nn []string
	for _, _nn := range myProfile.NickNamesViaID {
		nn = append(nn, _nn)
	}
	lg.Nicknames = nn
	return lg
}

func BoardcastLocifyRoundResult(h *Hub, roomname, winnerID, losserID, winnerTeamname, losserTeamname, guessToken string, lg LocfiyGameInfo, round, redCurrentScore, bluCurrentScore int, myProfile, wi, li LocifyFixtures) {
	LocifyRoundOverMessage(h, winnerTeamname, roomname)
	h.wg.Go(func() {
		myTeam := myProfile.MyTeam
		if winnerTeamname == _TeamRedKey {
			lg.TeamName = "RED"
		} else {
			lg.TeamName = "BLUE"
		}

		_c, err := json.Marshal(lg)
		if err != nil {
			log.Println(err)
			return
		}

		token := "LocifyGame: " + string(_c)

		for _, _id := range saveShuffle[roomname][winnerTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				token: token, roomname: roomname, to: _id,
				_sleep: false,
			}
		}

		if lg.TeamName == "RED" {
			lg.TeamName = "BLUE"
		} else {
			lg.TeamName = "RED"
		}
		_c, err = json.Marshal(lg)
		if err != nil {
			log.Println(err)
			return
		}

		token = "LocifyGame: " + string(_c)

		for _, _id := range saveShuffle[roomname][losserTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				token: token, roomname: roomname, to: _id,
				_sleep: false,
			}
		}

		updateStats := map[string]map[int]string{}

		updateStats[winnerID] = map[int]string{
			round - 1: "1",
		}
		updateStats[losserID] = map[int]string{
			round - 1: "0",
		}

		if myProfile.RoomSettings.Capacity == 4 {
			wi := getLocifyProfile[winnerID][roomname].IdlePlayer
			li := getLocifyProfile[losserID][roomname].IdlePlayer
			updateStats[wi] = map[int]string{
				round - 1: "I",
			}
			updateStats[li] = map[int]string{
				round - 1: "I",
			}
		}

		if myProfile.RoomSettings.Capacity == 2 {
			log.Println("1 v 1")
			_, m := LocifyMatchUpOneVOne(h, roomname, saveShuffle)

			p := LocifyUpdateProfile(h, m, round, roomname,
				redCurrentScore, bluCurrentScore, myTeam,
				myProfile.OppoTeamname, guessToken,
				myProfile.RoomSettings, updateStats, 1)

			go func() {
				SingleSaveStats(false, winnerID, roomname, 0, 0, guessToken, 1, _StringSentinel_, nil)
			}()
			go func() {
				SingleSaveStats(false, losserID, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			LocifyBlockManager(h, round, saveShuffle, p, roomname, winnerTeamname, losserTeamname, false)

		} else {
			log.Println("2 v 2")
			_, m := LocifyMatchUp(h, round, roomname, saveShuffle)

			p := LocifyUpdateProfile(h, m, round, roomname,
				redCurrentScore, bluCurrentScore,
				myTeam, myProfile.OppoTeamname,
				guessToken, myProfile.RoomSettings, updateStats, 1)
			go func() {
				SingleSaveStats(false, winnerID, roomname, 0, 0, guessToken, 1, _StringSentinel_, nil)
			}()
			go func() {
				SingleSaveStats(false, losserID, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			go func() {
				SingleSaveStats(false, wi.IdlePlayer, roomname, 0, 0, guessToken, 1, _StringSentinel_, nil)
			}()
			go func() {
				SingleSaveStats(false, li.IdlePlayer, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			LocifyBlockManager(h, round, saveShuffle, p, roomname, winnerTeamname, losserTeamname, false)

		}
	})
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: _CanUsePower}

}

func LocifyRoundOverDrawMessage(h *Hub, roomname string) {
	msg := "game draw the toss will be played"
	var t AlertMessage
	t.Alert = true
	t.Message = msg
	c, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.broadcast <- BroadcastReq{Token: Message + string(c), RoomID: roomname}
}

func BoardcastLocifyRoundDrawResult(h *Hub, roomname, myID, oppoID, myTeam, oppoTeam, guessToken string, lg LocfiyGameInfo, round, redCurrentScore, bluCurrentScore int, myProfile, wi, li LocifyFixtures) {
	LocifyRoundOverDrawMessage(h, roomname)
	h.wg.Go(func() {
		myTeam := myProfile.MyTeam
		if myTeam == _TeamRedKey {
			lg.TeamName = "RED"
		} else {
			lg.TeamName = "BLUE"
		}

		_c, err := json.Marshal(lg)
		if err != nil {
			log.Println(err)
			return
		}

		token := "LocifyGame: " + string(_c)

		for _, _id := range saveShuffle[roomname][myTeam] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				token: token, roomname: roomname, to: _id,
				_sleep: false,
			}
		}

		if lg.TeamName == "RED" {
			lg.TeamName = "BLUE"
		} else {
			lg.TeamName = "RED"
		}

		_c, err = json.Marshal(lg)
		if err != nil {
			log.Println(err)
			return
		}

		token = "LocifyGame: " + string(_c)

		for _, _id := range saveShuffle[roomname][oppoTeam] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				token: token, roomname: roomname, to: _id,
				_sleep: false,
			}
		}

		updateStats := map[string]map[int]string{}

		updateStats[myID] = map[int]string{
			round - 1: "0",
		}
		updateStats[oppoID] = map[int]string{
			round - 1: "0",
		}

		if myProfile.RoomSettings.Capacity == 4 {
			wi := getLocifyProfile[myID][roomname].IdlePlayer
			li := getLocifyProfile[oppoID][roomname].IdlePlayer
			updateStats[wi] = map[int]string{
				round - 1: "I",
			}
			updateStats[li] = map[int]string{
				round - 1: "I",
			}
		}

		if myProfile.RoomSettings.Capacity == 2 {
			log.Println("1 v 1")
			_, m := LocifyMatchUpOneVOne(h, roomname, saveShuffle)

			p := LocifyUpdateProfile(h, m, round, roomname,
				redCurrentScore, bluCurrentScore, myTeam,
				myProfile.OppoTeamname, guessToken,
				myProfile.RoomSettings, updateStats, 0)

			go func() {
				SingleSaveStats(false, myID, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			go func() {
				SingleSaveStats(false, oppoID, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			LocifyBlockManager(h, round, saveShuffle, p, roomname, myTeam, oppoTeam, true)

		} else {
			log.Println("2 v 2")
			_, m := LocifyMatchUp(h, round, roomname, saveShuffle)

			p := LocifyUpdateProfile(h, m, round, roomname,
				redCurrentScore, bluCurrentScore,
				myTeam, myProfile.OppoTeamname,
				guessToken, myProfile.RoomSettings, updateStats, 0)
			go func() {
				SingleSaveStats(false, myID, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			go func() {
				SingleSaveStats(false, oppoID, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			go func() {
				SingleSaveStats(false, wi.IdlePlayer, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			go func() {
				SingleSaveStats(false, li.IdlePlayer, roomname, 0, 0, guessToken, 0, _StringSentinel_, nil)
			}()
			LocifyBlockManager(h, round, saveShuffle, p, roomname, myID, oppoID, true)

		}
	})
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: _CanUsePower}

}
