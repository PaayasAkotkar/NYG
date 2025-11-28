package server

import (
	"encoding/json"
	"log"
	"math/rand/v2"
)

type ClashMatchToken struct {
	Lock    bool   `json:"lock"`
	Block   bool   `json:"block"`
	Pairing string `json:"pairing"` // teamname
}

// dead blocking:
// first round -> random shuffle
// cases:
// both have same lives-> random shuffle for toss
// one with the highest lives-> toss the coin

// RandomClashDeadLock returns ids of the blocked one and unlock one
func RandomClashDeadLock(h *Hub, roomname string, TeamRed []string, TeamBlue []string, Round int, updatedProfile map[string]map[string]ClashFixtures) ([]string, []string) {
	// turn is basically nothing but passing the default RBlock
	// logic:
	// one of the TeamReded player will get RBlock
	// meaning that the RUnblock player will able to pick the toss
	log.Println("in dead RBlock")
	log.Println("room name: ", roomname)
	Block1 := _StringSentinel_
	Block2 := _StringSentinel_
	UBlock1 := _StringSentinel_
	UBlock2 := _StringSentinel_

	rand.Shuffle(len(TeamRed), func(i, j int) {
		TeamRed[i], TeamRed[j] = TeamRed[j], TeamRed[i]
	})

	rand.Shuffle(len(TeamBlue), func(i, j int) {
		TeamBlue[i], TeamBlue[j] = TeamBlue[j], TeamBlue[i]
	})

	if Round == 1 {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, token: Block, to: TeamRed[0], _sleep: false}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, token: Unblock, to: TeamRed[1], _sleep: false}

		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, token: Block, to: TeamBlue[0], _sleep: false}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, token: Unblock, to: TeamBlue[1], _sleep: false}
		Block1 = TeamRed[0]
		Block2 = TeamBlue[0]
		UBlock1 = TeamRed[1]
		UBlock2 = TeamBlue[1]

	} else {
		red1 := updatedProfile[TeamRed[0]][roomname].MyCurrentChances
		red2 := updatedProfile[TeamRed[1]][roomname].MyCurrentChances

		Block1, UBlock1 = ClashBlockManager(TeamRed[0], TeamRed[1], red1, red2)

		blue1 := updatedProfile[TeamBlue[0]][roomname].MyCurrentChances
		blue2 := updatedProfile[TeamBlue[1]][roomname].MyCurrentChances

		Block2, UBlock2 = ClashBlockManager(TeamBlue[0], TeamBlue[1], blue1, blue2)
	}

	_block := []string{Block1, Block2}
	_unblock := []string{UBlock1, UBlock2}

	return _block, _unblock
}

// if the length of the players are 2:
// go against one v one
// if the length of the player are 3:
// one player waits

// about the powers:
// if both the players used same powers:=
// smudge the power
// meaning rather than two calls only one call will be made to them

// ClashMatchUp returns the matches division into red team and blue team where red vs red and blue vs blue all the time but ids will be different
// CurrentChances must be id->..
// faceOffID must not be the id of the final boss
func ClashMatchUp(h *Hub, round int, roomname string,
	teams map[string]map[string][]string,
	CurrentChances map[string]int,
	NickName map[string]map[string]string) (map[string]map[string][]string, map[string]map[string]ClashFixtures) {
	// TeamRed the untagged players
	log.Println("in pad TeamRed")
	log.Println("round: ", round)
	log.Println("room name: ", roomname)
	fromA1 := teams[roomname][_TeamRedKey][0]
	fromA2 := teams[roomname][_TeamRedKey][1]

	fromB1 := teams[roomname][_TeamBlueKey][0]
	fromB2 := teams[roomname][_TeamBlueKey][1]

	matches := map[int]ClashMatch{
		1: {TeamRed: []string{fromA1, fromA2}, TeamBlue: []string{fromB1, fromB2}},
		2: {TeamRed: []string{fromA1, fromB1}, TeamBlue: []string{fromA2, fromB2}},
		3: {TeamRed: []string{fromA1, fromB2}, TeamBlue: []string{fromA2, fromB1}},
		// repeat back
		4: {TeamRed: []string{fromA1, fromA2}, TeamBlue: []string{fromB1, fromB2}},
		5: {TeamRed: []string{fromA1, fromB1}, TeamBlue: []string{fromB2, fromA2}},
		6: {TeamRed: []string{fromA2, fromB1}, TeamBlue: []string{fromB2, fromA1}},
	}

	store := map[string]map[string][]string{}
	store[roomname] = map[string][]string{
		_TeamBlueKey: matches[round].TeamBlue,
		_TeamRedKey:  matches[round].TeamRed,
	}

	store2 := map[string]map[string]ClashFixtures{}
	_ids := []string{}
	_ids = append(_ids, matches[round].TeamBlue...)
	_ids = append(_ids, matches[round].TeamRed...)

	if round > 1 {
		UpdateProfile(true, matches[round].TeamRed[0], roomname, matches[round].TeamRed[1], _TeamRedKey, _StringSentinel_, nil, false, false, _StringSentinel_, _StringSentinel_, nil, round)
		UpdateProfile(true, matches[round].TeamBlue[0], roomname, matches[round].TeamBlue[1], _TeamRedKey, _StringSentinel_, nil, false, false, _StringSentinel_, _StringSentinel_, nil, round)
		UpdateProfile(true, matches[round].TeamRed[1], roomname, matches[round].TeamRed[0], _TeamRedKey, _StringSentinel_, nil, false, false, _StringSentinel_, _StringSentinel_, nil, round)
		UpdateProfile(true, matches[round].TeamBlue[1], roomname, matches[round].TeamBlue[0], _TeamRedKey, _StringSentinel_, nil, false, false, _StringSentinel_, _StringSentinel_, nil, round)

		var a = getClashProfile[matches[round].TeamRed[0]][roomname]
		a.MyCurrentChances = CurrentChances[matches[round].TeamRed[0]]
		var b = getClashProfile[matches[round].TeamRed[1]][roomname]
		a.MyCurrentChances = CurrentChances[matches[round].TeamRed[1]]
		var c = getClashProfile[matches[round].TeamBlue[0]][roomname]
		a.MyCurrentChances = CurrentChances[matches[round].TeamBlue[0]]
		var d = getClashProfile[matches[round].TeamBlue[1]][roomname]
		a.MyCurrentChances = CurrentChances[matches[round].TeamBlue[0]]

		store2[matches[round].TeamRed[0]] = map[string]ClashFixtures{
			roomname: a,
		}
		store2[matches[round].TeamRed[1]] = map[string]ClashFixtures{
			roomname: b,
		}
		store2[matches[round].TeamBlue[0]] = map[string]ClashFixtures{
			roomname: c,
		}
		store2[matches[round].TeamBlue[1]] = map[string]ClashFixtures{
			roomname: d,
		}
	} else {
		store2[matches[round].TeamRed[0]] = map[string]ClashFixtures{
			roomname: {
				BroadcastID:        _W,
				IDs:                _ids,
				Clash:              true,
				Against:            matches[round].TeamRed[1],
				MyTeam:             _TeamRedKey,
				MyID:               matches[round].TeamRed[0],
				OppoPowerUp:        make(map[string]bool),
				OppoTossDone:       false,
				OppoDictionaryDone: false,
				OppoChallengeDone:  false,
				OppoBetDone:        false,
				OppoTossCalled:     _StringSentinel_,
				SetDictionary:      _StringSentinel_,
				OppoSetChallenge:   _StringSentinel_,
				OppoSetBet:         _StringSentinel_,
				Book:               _StringSentinel_,
				Count:              1,
				NextTeamWinner:     "none",
				NextTeamname:       _TeamBlueKey,
				MyCurrentChances:   int(CurrentChances[matches[round].TeamRed[1]]),
			}}

		store2[matches[round].TeamRed[1]] = map[string]ClashFixtures{
			roomname: {
				BroadcastID:        _X,
				IDs:                _ids,
				Clash:              true,
				Against:            matches[round].TeamRed[0],
				MyTeam:             _TeamRedKey,
				MyID:               matches[round].TeamRed[1],
				OppoTossDone:       false,
				OppoDictionaryDone: false,
				OppoChallengeDone:  false,
				OppoBetDone:        false,
				OppoPowerUp:        make(map[string]bool),
				OppoTossCalled:     _StringSentinel_,
				SetDictionary:      _StringSentinel_,
				OppoSetChallenge:   _StringSentinel_,
				OppoSetBet:         _StringSentinel_,
				Book:               _StringSentinel_,
				Count:              1,
				NextTeamWinner:     "none",
				NextTeamname:       _TeamBlueKey,
				MyCurrentChances:   int(CurrentChances[matches[round].TeamRed[0]]),
			}}

		store2[matches[round].TeamBlue[0]] = map[string]ClashFixtures{
			roomname: {
				BroadcastID:        _Y,
				IDs:                _ids,
				Clash:              true,
				Against:            matches[round].TeamBlue[1],
				MyTeam:             _TeamBlueKey,
				MyID:               matches[round].TeamBlue[0],
				OppoTossDone:       false,
				OppoDictionaryDone: false,
				OppoChallengeDone:  false,
				OppoBetDone:        false,
				OppoPowerUp:        make(map[string]bool),
				OppoTossCalled:     _StringSentinel_,
				SetDictionary:      _StringSentinel_,
				OppoSetChallenge:   _StringSentinel_,
				OppoSetBet:         _StringSentinel_,
				Book:               _StringSentinel_,
				Count:              1,
				NextTeamWinner:     "none",
				NextTeamname:       _TeamRedKey,
				MyCurrentChances:   CurrentChances[matches[round].TeamBlue[0]],
			}}

		store2[matches[round].TeamBlue[1]] = map[string]ClashFixtures{
			roomname: {
				BroadcastID:        _Z,
				IDs:                _ids,
				Clash:              true,
				Against:            matches[round].TeamBlue[0],
				MyID:               matches[round].TeamBlue[1],
				MyTeam:             _TeamBlueKey,
				OppoTossDone:       false,
				OppoDictionaryDone: false,
				OppoChallengeDone:  false,
				OppoBetDone:        false,
				OppoPowerUp:        make(map[string]bool),
				OppoTossCalled:     _StringSentinel_,
				SetDictionary:      _StringSentinel_,
				OppoSetChallenge:   _StringSentinel_,
				OppoSetBet:         _StringSentinel_,
				Book:               _StringSentinel_,
				Count:              1,
				NextTeamWinner:     "none",
				NextTeamname:       _TeamRedKey,
				MyCurrentChances:   int(CurrentChances[matches[round].TeamBlue[1]]),
			}}

	}

	TeamRed := matches[round].TeamRed
	TeamBlue := matches[round].TeamBlue

	_block, _unblock := RandomClashDeadLock(h, roomname, TeamRed, TeamBlue,
		round, store2)
	for _, _id := range _block {
		c := ClashMatchToken{}
		c.Block = true
		c.Lock = true
		_t, _ := json.Marshal(&c)
		_x := string(_t)
		_token := "ClashMatchUp: " + _x
		h.gameRoomBroadcast <- reqGameRoomBroadcast{to: _id, _sleep: false, roomname: roomname, token: _token}
	}
	for _, _id := range _unblock {
		c := ClashMatchToken{}
		c.Block = false
		c.Lock = true
		_t, _ := json.Marshal(&c)
		_x := string(_t)
		_token := "ClashMatchUp: " + _x
		h.gameRoomBroadcast <- reqGameRoomBroadcast{to: _id, _sleep: false, roomname: roomname, token: _token}
	}

	return store, store2
}

// Clash1v1 implements
// locks 2 players into game hall
// returns matchup and teamname
// not to make it complicated
// both the players are dropped in same team
func Clash1v1(h *Hub, lastDance bool, eliminatedIDs []string, roomname, finalBossID,
	winnerID, losserID string,
	winnerCurrentChance, losserCurrentChance, round int) map[string]map[string][]string {

	store := map[string]map[string][]string{}
	store[roomname] = map[string][]string{
		_TeamRedKey: {winnerID, losserID},
	}

	UpdateProfile(true, winnerID, roomname, losserID, _TeamRedKey, finalBossID, eliminatedIDs, false, false, _StringSentinel_, _StringSentinel_, nil, round)
	UpdateProfile(true, losserID, roomname, winnerID, _TeamRedKey, finalBossID, eliminatedIDs, false, false, _StringSentinel_, _StringSentinel_, nil, round)

	// if redPlayerCurrentChance > bluePlayerCurrentChance {
	// 	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: bluePlayerID, token: Block, roomname: roomname, _sleep: false}
	// 	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: redPlayerID, token: Unblock, roomname: roomname, _sleep: false}
	// } else if redPlayerCurrentChance == bluePlayerCurrentChance {
	// 	_ids_ := store[roomname][_TeamRedKey]
	// 	rand.Shuffle(len(_ids_), func(i, j int) {
	// 		_ids_[i], _ids_[j] = _ids_[j], _ids_[i]
	// 	})
	// 	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: _ids_[0], token: Block, roomname: roomname, _sleep: false}
	// 	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: _ids_[1], token: Unblock, roomname: roomname, _sleep: false}

	// } else {
	// 	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: bluePlayerID, token: Unblock, roomname: roomname, _sleep: false}
	// 	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: redPlayerID, token: Block, roomname: roomname, _sleep: false}
	// }

	b, ub := ClashBlockManager(winnerID, losserID, winnerCurrentChance, losserCurrentChance)

	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: ub, token: Unblock, roomname: roomname, _sleep: false}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: b, token: Block, roomname: roomname, _sleep: false}

	// h.gameRoomBroadcast <- reqGameRoomBroadcast{to: winnerID, token: _DictionaryDiscussion, roomname: roomname, _sleep: false}
	// h.gameRoomBroadcast <- reqGameRoomBroadcast{to: losserID, token: _DictionaryDiscussion, roomname: roomname, _sleep: false}
	h.broadcast <- BroadcastReq{Token: _DictionaryDiscussion, RoomID: roomname}
	return store

}
