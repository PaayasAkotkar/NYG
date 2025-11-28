package server

import "log"

// ClashChallengeSession returns
// proceed to game if and only if team red and team blue has completed the challenge
// first checking if the teamred->a->b==done note: currentID!=teamredID as we are counting it as true
// if true than check for teamblue->a->b==done
// if both the true -> proceed
// for else mostly we are going look in the same team storage
// for clash it is 4
func ClashChallengeSession(h *Hub, id string, roomname string, challengeToken string, MaxPlayerCount int, fromTimeup bool) {
	log.Println("clash challenge disucssion")

	myTeam, myOpponentID := getClashProfile[id][roomname].MyTeam, getClashProfile[id][roomname].Against
	myProfile := getClashProfile[id][roomname]

	Count := int(1)
	oppoProfile := getClashProfile[myOpponentID][roomname]
	sessionDone := getClashProfile[id][roomname].Count == MaxPlayerCount

	proceed := getClashProfile[id][roomname].OppoChallengeDone || getClashProfile[id][roomname].OppoBetDone

	token := "ChallengeGuess: " + challengeToken
	fromOpp := "ChallengeGuess: " + getClashProfile[id][roomname].OppoSetChallenge

	log.Println("session done: ", sessionDone)

	log.Println("freeze power: ", getClashProfile[id][roomname].OppoPowerUp[_FreezeKey])
	log.Println("freeze power of oppo: ", getClashProfile[myOpponentID][roomname].OppoPowerUp[_FreezeKey])

	log.Println("max count: ", MaxPlayerCount, "current count: ", getClashProfile[id][roomname].Count+1)
	log.Println("procced: ", proceed)
	_fromOpp := getClashProfile[id][roomname].OppoSetChallenge
	log.Println("challenge token: ", challengeToken)
	log.Println("my team: ", getClashProfile[id][roomname].MyTeam)
	switch true {
	case sessionDone:
		log.Println("seesionDone!!!!")
		bets := getClashProfile[id][roomname]
		SumupTokensGame(h, id, roomname, challengeToken, false,fromTimeup, bets.OppoBets)

	case proceed:
		log.Println("proceed!!!!")

		for _, _id := range saveShuffle[roomname][myTeam] {
			if _id != id {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, roomname: roomname, _sleep: false, to: _id}
			} else if !getClashProfile[id][roomname].OppoPowerUp[_BetKey] {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: fromOpp, roomname: roomname, _sleep: false, to: _id}
			}
			msg := "Hang in there!!!! we'll tally everything up once everyone's through"
			BWait(h, msg, roomname, myTeam, true)
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RemoveChallenge + challengeToken, roomname: roomname, _sleep: false, to: _id}

			// h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _waiting, roomname: roomname, _sleep: false, to: _id}
			// h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _ClashWaitMessgae, roomname: roomname, _sleep: false, to: _id}
		}
		if getClashProfile[id][roomname].OppoPowerUp[_FreezeKey] && getClashProfile[myOpponentID][roomname].OppoPowerUp[_FreezeKey] {
			log.Println("Both")
			PFreezeMechanism(h, roomname, myOpponentID, myTeam, true, true, 0)
			go func() {
				ClashResetPowerUp(myOpponentID, roomname, _FreezeKey)
			}()
			go func() {
				ClashResetPowerUp(id, roomname, _FreezeKey)
			}()
			BoardcastRemovePower(h, id, roomname, _FreezeKey)
			BoardcastRemovePower(h, myOpponentID, roomname, _FreezeKey)

		} else {
			switch true {
			case getClashProfile[id][roomname].OppoPowerUp[_NexusKey]:
				dictionary := getClashProfile[id][roomname].SetDictionary
				book := getClashProfile[id][roomname].Book
				PNexusMechanism(h, roomname, myOpponentID, book, dictionary, challengeToken, int(oppoProfile.MyGameProfile.NexusLevel))
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

		if getClashProfile[id][roomname].OppoPowerUp[_BetKey] {
			a := getClashProfile[id][roomname].OppoBets
			BroadcastBets(h, id, myTeam, roomname, a)
			ClashResetPowerUp(id, roomname, _BetKey)
			BoardcastRemovePower(h, myOpponentID, roomname, _BetKey)

		}
		go func() {
			MasterSave(true, int(Count), _StringSentinel_, _StringSentinel_, _StringSentinel_,
				roomname, myTeam, _StringSentinel_, _StringSentinel_,
				_StringSentinel_, false, false,
				_IntSentinel, _IntSentinel,
				_StringSentinel_, _StringSentinel_, nil, nil)
		}()
		go func() {
			SingleSave(true, myOpponentID, roomname, myTeam, _StringSentinel_,
				challengeToken, _StringSentinel_, nil, false, false, true, false)
		}()
		log.Println("opp prof: ", getClashProfile[myOpponentID][roomname])
		log.Println("challenge token: ", challengeToken)

	default:
		log.Println("default!!!!")
		for _, _id := range saveShuffle[roomname][myTeam] {
			if id != _id {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, roomname: roomname, _sleep: false, to: _id}
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
			} else {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RemoveChallenge + challengeToken, roomname: roomname, _sleep: false, to: _id}
		}

		if getClashProfile[id][roomname].OppoPowerUp[_BetKey] {
			dictionary := getClashProfile[id][roomname].SetDictionary
			b := getClashProfile[id][roomname].Book
			InitBet(h, myOpponentID, myTeam, roomname, b, dictionary, true)
			ClashResetPowerUp(id, roomname, _BetKey)
			BoardcastRemovePower(h, myOpponentID, roomname, _BetKey)
		}

		go func() {
			MasterSave(true, int(Count), _StringSentinel_, _StringSentinel_,
				_StringSentinel_,
				roomname, myTeam, _StringSentinel_, _StringSentinel_,
				_StringSentinel_, false, false,
				_IntSentinel, _IntSentinel,
				_StringSentinel_, _StringSentinel_, nil, nil)
		}()

		go func() {
			SingleSave(true, myOpponentID, roomname, myTeam, _StringSentinel_,
				challengeToken, _StringSentinel_, nil, false, false,
				true, false)
		}()
	}
}
