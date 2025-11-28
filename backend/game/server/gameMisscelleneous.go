package server

// IsRewind returns true and id of the used rewind resets the count for the rewind and stores the guessMade in the requested id
func IsRewind(id string, roomname string, guessMade string, clash bool) (bool, string) {
	if clash {
		myProf := getClashProfile[id][roomname]
		opponentProf := getClashProfile[myProf.Against][roomname]
		if myProf.OppoPowerUp[_RewindKey] {

			SingleSaveStats(true, id, roomname, _IntSentinel, _IntSentinel, guessMade, 0, _StringSentinel_, nil)
			return true, id
		} else if opponentProf.OppoPowerUp[_RewindKey] {
			SingleSaveStats(true, id, roomname, _IntSentinel, _IntSentinel, guessMade, 0, _StringSentinel_, nil)
			return true, myProf.Against
		}
	} else {
		myProf := getLocifyProfile[id][roomname]
		opponentProf := getLocifyProfile[myProf.Against][roomname]
		if myProf.OppoPowerUp[_RewindKey] {
			SingleSaveStats(false, id, roomname, _IntSentinel, _IntSentinel, guessMade, 0, _StringSentinel_, nil)
			return true, id
		} else if opponentProf.OppoPowerUp[_RewindKey] {
			SingleSaveStats(false, id, roomname, _IntSentinel, _IntSentinel, guessMade, 0, _StringSentinel_, nil)
			return true, myProf.Against
		}

	}
	return false, ""
}
