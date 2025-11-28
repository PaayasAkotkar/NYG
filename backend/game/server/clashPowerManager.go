package server

// ClashPowerManager stores the requested power in opponent profile and broadcasts for the usage to the request
// storing in opponent profile for easier broadcast
func ClashPowerManager(h *Hub, nexus, covert, bet, freeze, rewind bool, myID, myOpponentID, myTeam, roomname string) {
	power := map[string]bool{}
	usedPower := ""
	switch true {
	case nexus:
		power[_NexusKey] = true
		usedPower = _NexusKey
	case bet:
		power[_BetKey] = true
		usedPower = _BetKey
		dict := getClashProfile[myID][roomname].SetDictionary
		if dict != _StringSentinel_ {
			b := getClashProfile[myID][roomname].Book

			InitBet(h, myID, myTeam, roomname, b, dict, true)
		}
		// save opponent so that we can send the it to the opponent
		MasterSave(true, _IntSentinel, _StringSentinel_, _StringSentinel_,
			_StringSentinel_, roomname, myTeam,
			_StringSentinel_, _StringSentinel_,
			myOpponentID, false, false, _IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, nil, nil)
	case freeze:
		usedPower = _FreezeKey
		power[_FreezeKey] = true
	case rewind:
		usedPower = _RewindKey
		power[_RewindKey] = true
	case covert:
		usedPower = _CovertKey
		power[_CovertKey] = true
	}
	SingleSave(true, myOpponentID, roomname, myTeam,
		_StringSentinel_, _StringSentinel_, _StringSentinel_,
		power, false, false, false, false)
	SingleSaveStats(true, myID, roomname, 0, 0, _StringSentinel_, 0, usedPower, nil)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: myID, token: CanUsePower, roomname: roomname, _sleep: false}
}
