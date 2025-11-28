package server

import "log"

func RestPower(clash bool, id, roomname, key string) {
	if clash {
		store := map[string]map[string]IPowerReset{}
		store[id] = map[string]IPowerReset{
			roomname: {Key: key, Clash: clash},
		}
		_resetPower <- store
	}
}

// UpdateProfile for clash the guess token must be matches and for locfiy it can be the current call
func UpdateProfile(clash bool, _in, roomname, against, myTeam, finalBossID string, eliminatedIDS []string,
	block, lock bool, myPartnerID string,
	guess string, updateStats map[int]string, round int) {

	store := map[string]map[string]UpdateFixtures{}
	store[_in] = map[string]UpdateFixtures{
		roomname: {
			Against:              against,
			FinalBossID:          finalBossID,
			EliminatedPlayersIDs: eliminatedIDS,
			MyTeam:               myTeam,
			Clash:                clash,
			Block:                block,
			Lock:                 lock,
			MyPartnerID:          myPartnerID,
			WholeGuess:           guess,
			UpdateStats:          updateStats,
			CurrentRound:         round,
		},
	}
	updateProfile <- store

}

// MasterSave note: setDictionary and bets are only for the locify
func MasterSave(clash bool, count int, eliminatedID, winnerID, losserID,
	roomname, winnerTeamname, FinalBossID, book, betID string,
	FinalBossFound, LastDance bool, redScore, blueScore int, prevWinner string, setDictionary string, bets []string, newList []string) {
	store := map[string]IMasterSave{}

	log.Println("bet: ", bets)

	store[roomname] = IMasterSave{
		Clash:              clash,
		BetID:              betID,
		FinalBossID:        FinalBossID,
		FinalBossFound:     FinalBossFound,
		LastDance:          LastDance,
		WinnerID:           winnerID,
		LosserID:           losserID,
		WinnerTeamName:     winnerTeamname,
		Count:              count,
		Book:               book,
		EliminatedPlayerID: eliminatedID,
		RedTeamScore:       redScore,
		BlueTeamScore:      blueScore,
		SetDictionary:      setDictionary,
		PrevList:           newList,
		Bets:               bets,
	}
	saveGlobal <- store

}

// SingleSave saves the tokens of requested id in the opponent profile
// this will update the token set by the opponent
// dont consider it as the id's set challenge or set dictionary
func SingleSave(clash bool, opponentID, roomname, myTeam, tossCalled, setChallenge,

	setBet string, powerUp map[string]bool, tossDone, dictionaryDone, challengeDone, betDone bool) {

	store := map[string]map[string]ISingleSave{}
	store[opponentID] = map[string]ISingleSave{
		roomname: {
			MyTeam:         myTeam,
			PowerUp:        powerUp,
			TossCalled:     tossCalled,
			SetChallenge:   setChallenge,
			SetBet:         setBet,
			TossDone:       tossDone,
			ChallengeDone:  challengeDone,
			DictionaryDone: dictionaryDone,
			BetDone:        betDone,
			Clash:          clash,
		},
	}
	saveSingle <- store

}

// SingleSaveStats saves the stats for the id
// or you can say to update this id profile
func SingleSaveStats(clash bool, _forID string, roomname string,
	currentChances, onFire int, guess string, points int, powersBin string, penalty map[int]int) {

	store := map[string]map[string]ISingleStatsSave{}
	store[_forID] = map[string]ISingleStatsSave{
		roomname: {
			OnFire:         onFire,
			CurrentChances: currentChances,
			Guess:          guess,
			Clash:          clash,
			Points:         points,
			PowersBin:      powersBin,
			Penalty:        penalty,
		},
	}
	saveStats <- store

}

func OneTimeSave(clash bool, roomname, teamname, dictionary string) {

	store := map[string]IOneTimeSave{}
	store[roomname] = IOneTimeSave{
		TeamName:      teamname,
		SetDictionary: dictionary,
		Clash:         clash,
	}
	saveOneTime <- store
}

// CommonSave fit for the clash because the players are mapped into one team unlike locify
func CommonSave(clash bool, roomname, teamname string, bets []string, tossDone, dictionaryDone, betDone, challengeDone bool) {
	if clash {
		store := map[string]ICommonSave{}
		store[roomname] = ICommonSave{
			TeamName:       teamname,
			TossDone:       tossDone,
			DictionaryDone: dictionaryDone,
			BetDone:        betDone,
			ChallengeDone:  challengeDone,
			Bets:           bets,
			Clash:          clash,
		}
		saveCommon <- store
	}
}
