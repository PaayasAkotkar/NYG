package server

import "log"

// LocifyPowerManager stores the requested power in opponent profile and broadcasts for the usage to the request
// storing in opponent profile for easier broadcast
func LocifyPowerManager(h *Hub, nexus, covert, bet, freeze, rewind bool, myID, myOpponentID, myTeam, roomname string) {
	power := map[string]bool{}
	usedPower := ""
	switch true {
	case nexus:
		power[_NexusKey] = true
		usedPower = _NexusKey
	case bet:
		power[_BetKey] = true
		usedPower = _BetKey
		dict := getLocifyProfile[myID][roomname].SetDictionary

		if dict != _StringSentinel_ {
			b := getLocifyRoomSettings[roomname].Book

			bets := InitBet(h, myID, myTeam, roomname, b, dict, false)
			log.Println("fetched bets")
			go func() {
				// save opponent so that we can send the it to the opponent
				MasterSave(false, _IntSentinel, _StringSentinel_, _StringSentinel_,
					_StringSentinel_, roomname, myTeam,
					_StringSentinel_, _StringSentinel_,
					myOpponentID, false, false, _IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, bets, nil)
			}()
		}

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
	SingleSave(false, myOpponentID, roomname, myTeam,
		_StringSentinel_, _StringSentinel_, _StringSentinel_,
		power, false, false, false, false)
	SingleSaveStats(false, myID, roomname, 0, 0, _StringSentinel_, 0, usedPower, nil)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: myID, token: CanUsePower, roomname: roomname, _sleep: false}
}
