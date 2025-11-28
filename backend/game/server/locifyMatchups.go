package server

import (
	"log"
)

type Rounds struct {
	Round1 string `json:"round1"`
	Round2 string `json:"round2"`
	Round3 string `json:"round3"`
	Round4 string `json:"round4"`
	Round5 string `json:"round5"`
	Round6 string `json:"round6"`
}

type LocifyInit struct {
	Theme string `json:"theme"`
}

type LocfiyGameInfo struct {
	RedScore     int                       `json:"redScore"`
	BlueScore    int                       `json:"blueScore"`
	Round        int                       `json:"round"`
	Set          int                       `json:"set"`
	Powers       map[string]bool           `json:"powers"`
	Session      bool                      `json:"session"`
	Mode         string                    `json:"mode"`
	Roomname     string                    `json:"roomname"`
	TeamName     string                    `json:"teamname"`
	PlayersStats map[string]map[int]string `json:"stats"` // nickname->round->...
	Nicknames    []string                  `json:"nicknames"`
}

type BSession struct {
	Toss       bool `json:"toss"`
	Dictionary bool `json:"dictionary"`
	Challenge  bool `json:"challenge"`
	Game       bool `json:"game"`
	Clash      bool `json:"clash"`
}

// LocifyMatchUp returns the teams and matchups
func LocifyMatchUp(h *Hub, round int, roomname string,
	teams map[string]map[string][]string) (map[string]map[string][]string, LocifyMatch) {

	log.Println("in locify match up")

	log.Println("red team: ", teams[roomname][_TeamRedKey])
	log.Println("blue team: ", teams[roomname][_TeamBlueKey])

	R1 := teams[roomname][_TeamRedKey][0]
	R2 := teams[roomname][_TeamRedKey][1]

	B1 := teams[roomname][_TeamBlueKey][0]
	B2 := teams[roomname][_TeamBlueKey][1]

	m := map[int]LocifyMatch{
		1: {
			TeamRed: LockPlayerInfo{
				PlayingID: R1,
				IdleID:    R2,
			},
			TeamBlue: LockPlayerInfo{
				PlayingID: B1,
				IdleID:    B2,
			},
		},
		2: {
			TeamRed: LockPlayerInfo{
				PlayingID: R2,
				IdleID:    R1,
			},
			TeamBlue: LockPlayerInfo{
				PlayingID: B2,
				IdleID:    B1,
			},
		},
		3: {
			TeamRed: LockPlayerInfo{
				PlayingID: R1,
				IdleID:    R2,
			},
			TeamBlue: LockPlayerInfo{
				PlayingID: B2,
				IdleID:    B1,
			},
		},
		4: {
			TeamRed: LockPlayerInfo{
				PlayingID: R2,
				IdleID:    R1,
			},
			TeamBlue: LockPlayerInfo{
				PlayingID: B1,
				IdleID:    B1,
			},
		},
		5: {
			TeamRed: LockPlayerInfo{
				PlayingID: R1,
				IdleID:    R2,
			},
			TeamBlue: LockPlayerInfo{
				PlayingID: B1,
				IdleID:    B2,
			},
		},
		6: {
			TeamRed: LockPlayerInfo{
				PlayingID: R2,
				IdleID:    R1,
			},
			TeamBlue: LockPlayerInfo{
				PlayingID: B2,
				IdleID:    B1,
			},
		},
	}

	tRed := []string{m[round].TeamRed.PlayingID, m[round].TeamRed.IdleID}
	tblue := []string{m[round].TeamBlue.PlayingID, m[round].TeamBlue.IdleID}

	store := map[string]map[string][]string{}
	store[roomname] = map[string][]string{
		_TeamRedKey:  tRed,
		_TeamBlueKey: tblue,
	}

	return store, m[round]
}

func LocifyMatchUpOneVOne(h *Hub, roomname string,
	teams map[string]map[string][]string) (map[string]map[string][]string, LocifyMatch) {
	R1 := teams[roomname][_TeamRedKey][0]
	B1 := teams[roomname][_TeamBlueKey][0]
	m := map[int]LocifyMatch{
		1: {
			TeamRed: LockPlayerInfo{
				PlayingID: R1,
				IdleID:    _StringSentinel_,
			},
			TeamBlue: LockPlayerInfo{
				PlayingID: B1,
				IdleID:    _StringSentinel_,
			},
		},
	}
	store := map[string]map[string][]string{}
	store[roomname] = map[string][]string{
		_TeamRedKey:  {R1},
		_TeamBlueKey: {B1},
	}
	return store, m[1]
}

func LocifyProfile(h *Hub, round int, roomAdminID,
	book, roomname string, m LocifyMatch,
	roomCode string, roomSettings LocifyRoomSettings) map[string]map[string]LocifyFixtures {
	ids := []string{
		m.TeamBlue.PlayingID,
		m.TeamRed.PlayingID,
	}

	block, bIdle, unblock, uIdle, bteam, uteam := LocifyRandBlock(m, round, roomname)
	if roomSettings.Capacity == 4 {
		ids = append(ids, m.TeamRed.IdleID, m.TeamBlue.IdleID)
	}
	store := map[string]map[string]LocifyFixtures{}
	log.Println("matches: ", block, unblock)
	log.Println("matches partner: ", bIdle, uIdle)
	log.Println("team names: ", bteam, uteam)

	log.Println("ids: ", ids)

	store[block] = map[string]LocifyFixtures{
		roomname: {
			MyID:                block,
			IDs:                 ids,
			Book:                book,
			IdlePlayer:          bIdle,
			Against:             unblock,
			MyTeam:              bteam,
			MyPatnerID:          bIdle,
			OppoTeamname:        uteam,
			IBlock:              true,
			ILock:               true,
			MyPoints:            0,
			MyPowersBin:         nil,
			RoomAdminID:         roomAdminID,
			OppoPowerUp:         make(map[string]bool),
			MySheet:             make(map[int]string),
			NickNamesViaID:      make(map[string]string),
			OppoBets:            nil,
			BetIDs:              nil,
			OppoTossDone:        false,
			OppoDictionaryDone:  false,
			OppoChallengeDone:   false,
			OppoBetDone:         false,
			OppoPenalties:       0,
			MyPenalties:         nil,
			OppoTossCalled:      _StringSentinel_,
			OppoSetChallenge:    _StringSentinel_,
			OppoSetBet:          _StringSentinel_,
			DisconnectedID:      _StringSentinel_,
			PrevWinnerTeam:      _StringSentinel_,
			SetDictionary:       _StringSentinel_,
			DisconnectedIDFound: false,
			MyGuess:             nil,
			RedTeamScore:        0,
			BlueTeamScore:       0,
			RoomCode:            roomCode,
			RoomSettings:        roomSettings,
			WholeGuess:          nil,
			PowerDeck:           make(map[string]bool),
			MyGameProfile:       GameProfile{},
		},
	}

	store[unblock] = map[string]LocifyFixtures{
		roomname: {
			MyID:           unblock,
			IDs:            ids,
			Book:           book,
			IdlePlayer:     uIdle,
			Against:        block,
			MyTeam:         uteam,
			MyPatnerID:     uIdle,
			OppoTeamname:   bteam,
			IBlock:         false,
			ILock:          true,
			RoomAdminID:    roomAdminID,
			OppoPowerUp:    nil,
			MySheet:        make(map[int]string),
			NickNamesViaID: make(map[string]string),
			MyPoints:       0,
			MyPowersBin:    nil,

			OppoBets:           nil,
			BetIDs:             nil,
			OppoTossDone:       false,
			OppoDictionaryDone: false,
			OppoChallengeDone:  false,
			OppoBetDone:        false,
			OppoPenalties:      0,
			MyPenalties:        nil,

			DisconnectedIDFound: false,
			MyGuess:             nil,
			RedTeamScore:        0,
			BlueTeamScore:       0,
			RoomCode:            roomCode,
			RoomSettings:        roomSettings,
			WholeGuess:          nil,
			PrevWinnerTeam:      _StringSentinel_,
			SetDictionary:       _StringSentinel_,
			DisconnectedID:      _StringSentinel_,
			OppoTossCalled:      _StringSentinel_,
			OppoSetChallenge:    _StringSentinel_,
			OppoSetBet:          _StringSentinel_,
			PowerDeck:           make(map[string]bool),
			MyGameProfile:       GameProfile{},
		},
	}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname,
		token:    Lock, to: block, _sleep: false}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname,
		token:    Lock, to: unblock, _sleep: false}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname,
		token:    Block, to: block, _sleep: false}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname,
		token:    Unblock, to: unblock, _sleep: false}

	if roomSettings.Capacity > 2 {
		log.Println("cap == 4")
		store[bIdle] = map[string]LocifyFixtures{
			roomname: {
				MyID:           bIdle,
				IDs:            ids,
				Book:           book,
				IdlePlayer:     _StringSentinel_,
				Against:        _StringSentinel_,
				MyTeam:         bteam,
				MyPatnerID:     block,
				OppoTeamname:   uteam,
				IBlock:         true,
				ILock:          false,
				MySheet:        make(map[int]string),
				NickNamesViaID: make(map[string]string),
				MyPoints:       0,
				MyPowersBin:    nil,

				RoomAdminID:        roomAdminID,
				OppoPowerUp:        make(map[string]bool),
				OppoTossCalled:     _StringSentinel_,
				OppoSetChallenge:   _StringSentinel_,
				OppoSetBet:         _StringSentinel_,
				OppoBets:           nil,
				BetIDs:             nil,
				OppoTossDone:       false,
				OppoDictionaryDone: false,
				OppoChallengeDone:  false,
				OppoBetDone:        false,
				OppoPenalties:      0,
				MyPenalties:        nil,

				DisconnectedID:      _StringSentinel_,
				DisconnectedIDFound: false,
				MyGuess:             nil,
				SetDictionary:       _StringSentinel_,
				RedTeamScore:        0,
				BlueTeamScore:       0,
				RoomCode:            roomCode,
				RoomSettings:        roomSettings,
				PrevWinnerTeam:      _StringSentinel_,
				WholeGuess:          nil,
				PowerDeck:           make(map[string]bool),
				MyGameProfile:       GameProfile{},
			},
		}
		store[uIdle] = map[string]LocifyFixtures{
			roomname: {
				MyID:           uIdle,
				IDs:            ids,
				Book:           book,
				IdlePlayer:     _StringSentinel_,
				Against:        _StringSentinel_,
				MyTeam:         uteam,
				MyPatnerID:     unblock,
				OppoTeamname:   bteam,
				IBlock:         true,
				ILock:          false,
				MySheet:        make(map[int]string),
				NickNamesViaID: make(map[string]string),
				MyPoints:       0,
				MyPowersBin:    nil,

				RoomAdminID:        roomAdminID,
				OppoPowerUp:        make(map[string]bool),
				OppoTossCalled:     _StringSentinel_,
				OppoSetChallenge:   _StringSentinel_,
				OppoSetBet:         _StringSentinel_,
				OppoBets:           nil,
				BetIDs:             nil,
				OppoTossDone:       false,
				OppoDictionaryDone: false,
				OppoChallengeDone:  false,
				OppoBetDone:        false,
				OppoPenalties:      0,
				MyPenalties:        nil,

				DisconnectedID:      _StringSentinel_,
				DisconnectedIDFound: false,
				MyGuess:             nil,
				SetDictionary:       _StringSentinel_,
				RedTeamScore:        0,
				BlueTeamScore:       0,
				RoomCode:            roomCode,
				RoomSettings:        roomSettings,
				PrevWinnerTeam:      _StringSentinel_,
				WholeGuess:          nil,
				PowerDeck:           make(map[string]bool),
				MyGameProfile:       GameProfile{},
			},
		}
	}

	return store
}

// LocifyUpdateProfile updates the profile and returns the dummy profile
// note: it is a dummy profile only made changes recommended through param
// updateStats must be passed with the id->current_round->result
// the points only will be added in the winner's profile id
func LocifyUpdateProfile(h *Hub, m LocifyMatch, round int,
	roomname string,
	RedTeamScore, BlueTeamScore int,
	myTeamname, OppoTeamname,
	guess string, r LocifyRoomSettings,
	updateStats map[string]map[int]string, points int) map[string]map[string]LocifyFixtures {

	log.Println("in updating profile")

	block, bIdle, unblock, uIdle, bteam, uteam := LocifyRandBlock(m, round, roomname)

	if bteam != myTeamname {
		uteam = myTeamname
		bteam = OppoTeamname
	}
	if r.Capacity > 2 {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			roomname: roomname,
			token:    Unlock, to: bIdle, _sleep: false}

		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			roomname: roomname,
			token:    Unlock, to: uIdle, _sleep: false}
	}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname,
		token:    Lock, to: block, _sleep: false}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		roomname: roomname,
		token:    Lock, to: unblock, _sleep: false}

	log.Println("playing id: ", block, unblock)
	log.Println("their teams: ", bteam, uteam)

	go func() {
		UpdateProfile(false, block, roomname, unblock,
			_StringSentinel_, _StringSentinel_, nil, true, true,
			bIdle, guess, updateStats[block], round)
	}()

	go func() {
		UpdateProfile(false, unblock, roomname,
			block, _StringSentinel_, _StringSentinel_, nil,
			true, true,
			uIdle, guess, updateStats[unblock], round)
	}()

	if r.Capacity > 2 {
		go func() {
			UpdateProfile(false, bIdle, roomname,
				_StringSentinel_,
				_StringSentinel_, _StringSentinel_, nil,
				true,
				false, block, guess, updateStats[bIdle], round)
		}()
		go func() {
			UpdateProfile(false, uIdle, roomname,
				_StringSentinel_, _StringSentinel_, _StringSentinel_, nil,

				true, false, unblock, guess, updateStats[uIdle], round)
		}()
	}

	store := map[string]map[string]LocifyFixtures{}

	store[block] = map[string]LocifyFixtures{
		roomname: {
			MyID:         block,
			IDs:          nil,
			Book:         _StringSentinel_,
			IdlePlayer:   bIdle,
			Against:      unblock,
			MyTeam:       bteam,
			MyPatnerID:   bIdle,
			OppoTeamname: uteam,
			IBlock:       true,
			ILock:        true,

			RoomAdminID:        _StringSentinel_,
			OppoPowerUp:        make(map[string]bool),
			OppoTossCalled:     _StringSentinel_,
			OppoSetChallenge:   _StringSentinel_,
			OppoSetBet:         _StringSentinel_,
			OppoBets:           nil,
			BetIDs:             nil,
			OppoTossDone:       false,
			OppoDictionaryDone: false,
			OppoChallengeDone:  false,
			OppoBetDone:        false,
			OppoPenalties:      0,
			MyPenalties:        nil,
			MyPowersBin:        nil,

			DisconnectedID:      _StringSentinel_,
			DisconnectedIDFound: false,
			MyGuess:             nil,
			SetDictionary:       _StringSentinel_,
			RedTeamScore:        RedTeamScore,
			BlueTeamScore:       BlueTeamScore,
			RoomCode:            _StringSentinel_,
			RoomSettings:        r,
			PrevWinnerTeam:      _StringSentinel_,
			WholeGuess:          nil,
		},
	}
	store[unblock] = map[string]LocifyFixtures{
		roomname: {
			MyID:         unblock,
			IDs:          nil,
			Book:         _StringSentinel_,
			IdlePlayer:   uIdle,
			Against:      block,
			MyTeam:       uteam,
			MyPatnerID:   uIdle,
			OppoTeamname: bteam,
			IBlock:       false,
			ILock:        true,

			RoomAdminID:        _StringSentinel_,
			OppoPowerUp:        make(map[string]bool),
			OppoTossCalled:     _StringSentinel_,
			OppoSetChallenge:   _StringSentinel_,
			OppoSetBet:         _StringSentinel_,
			OppoBets:           nil,
			BetIDs:             nil,
			OppoTossDone:       false,
			OppoDictionaryDone: false,
			OppoChallengeDone:  false,
			OppoBetDone:        false,
			OppoPenalties:      0,
			MyPenalties:        nil,
			MyPowersBin:        nil,

			DisconnectedID:      _StringSentinel_,
			DisconnectedIDFound: false,
			MyGuess:             nil,
			SetDictionary:       _StringSentinel_,
			RedTeamScore:        RedTeamScore,
			BlueTeamScore:       BlueTeamScore,
			RoomCode:            _StringSentinel_,
			RoomSettings:        r,
			PrevWinnerTeam:      _StringSentinel_,
			WholeGuess:          nil,
		},
	}

	store[bIdle] = map[string]LocifyFixtures{
		roomname: {
			MyID:         bIdle,
			IDs:          nil,
			Book:         _StringSentinel_,
			IdlePlayer:   _StringSentinel_,
			Against:      _StringSentinel_,
			MyTeam:       bteam,
			MyPatnerID:   block,
			OppoTeamname: uteam,
			IBlock:       true,
			ILock:        false,

			RoomAdminID:         _StringSentinel_,
			OppoPowerUp:         nil,
			OppoTossCalled:      _StringSentinel_,
			OppoSetChallenge:    _StringSentinel_,
			OppoSetBet:          _StringSentinel_,
			OppoBets:            nil,
			BetIDs:              nil,
			OppoTossDone:        false,
			OppoDictionaryDone:  false,
			OppoChallengeDone:   false,
			OppoBetDone:         false,
			OppoPenalties:       0,
			MyPenalties:         nil,
			MyPowersBin:         nil,
			DisconnectedID:      _StringSentinel_,
			DisconnectedIDFound: false,
			MyGuess:             nil,
			SetDictionary:       _StringSentinel_,
			RedTeamScore:        RedTeamScore,
			BlueTeamScore:       BlueTeamScore,
			RoomCode:            _StringSentinel_,
			RoomSettings:        r,
			PrevWinnerTeam:      _StringSentinel_,
			WholeGuess:          nil,
		},
	}
	store[uIdle] = map[string]LocifyFixtures{
		roomname: {
			MyID:         uIdle,
			IDs:          nil,
			Book:         _StringSentinel_,
			IdlePlayer:   _StringSentinel_,
			Against:      _StringSentinel_,
			MyTeam:       uteam,
			MyPatnerID:   unblock,
			OppoTeamname: bteam,
			IBlock:       true,
			ILock:        false,

			RoomAdminID:        _StringSentinel_,
			OppoPowerUp:        make(map[string]bool),
			OppoTossCalled:     _StringSentinel_,
			OppoSetChallenge:   _StringSentinel_,
			OppoSetBet:         _StringSentinel_,
			OppoBets:           nil,
			BetIDs:             nil,
			OppoTossDone:       false,
			OppoDictionaryDone: false,
			OppoChallengeDone:  false,
			OppoBetDone:        false,
			OppoPenalties:      0,
			MyPenalties:        nil,
			MyPowersBin:        nil,

			DisconnectedID:      _StringSentinel_,
			DisconnectedIDFound: false,
			MyGuess:             nil,
			SetDictionary:       _StringSentinel_,
			RedTeamScore:        RedTeamScore,
			BlueTeamScore:       BlueTeamScore,
			RoomCode:            _StringSentinel_,
			RoomSettings:        r,
			PrevWinnerTeam:      _StringSentinel_,
			WholeGuess:          nil,
		},
	}

	return store
}
