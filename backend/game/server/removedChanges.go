package server

// type StorageRoom struct {
// 	mu    sync.Mutex
// 	store chan map[string]int
// 	done  chan bool
// }

// type ClientIDStorageRoom struct {
// 	mu       sync.Mutex
// 	clientID chan string
// 	done     chan bool
// }
// type ProfileData struct {
// 	mu   sync.Mutex
// 	data chan map[string]string
// 	done chan bool
// }

// type RoomSettingsParcel struct {
// 	Book         string   `json:"book"`
// 	GameTime     string   `json:"gameTime"`
// 	DecisionTime string   `json:"decisionTime"`
// 	Powers       []string `json:"powers"`
// }
// store := make(map[string]int)
// store[l.RoomName] = l.RoomCapacity
// roomStore.store <- store

// log.Println("store: ", store)

// func BookSeat(l Lobby, h *Hub, conn *websocket.Conn, roomname string) {
// 	log.Println("in book seat")
// 	// logic: for 4ppl meaning for 2v2
// 	// 3 are allowed from 3 one is owner and other two are random
// 	// and to complete a seat has been booked
// 	saveSeat := false
// 	for _, det := range getLocifyRoomSettings {
// 		if l.Code == det.Code {
// 			saveSeat = true
// 			break
// 		}
// 	}

// 	if saveSeat {
// 		log.Println("save saet")
// 		clientCount := len(h.rooms[roomname])
// 		isFull := l.RoomCapacity <= clientCount
// 		if !isFull {
// 			log.Println("room joining")
// 			RegisterRoom(h, conn, roomname, l.ID, l.NickName, false)
// 			// active connection
// 			Conns := strconv.Itoa(clientCount + 1)

// 			// broadcast to all the clients
// 			h.broadcast <- BroadcastReq{Token: "ActiveConns: " + Conns}

// 			log.Println("after joining total client in the room: ", Conns)
// 		} else {
// 			log.Println("room is full")
// 		}
// 	} else {
// 		log.Println("joining code is wrong")
// 	}

// }

// func ChallengeSession(h *Hub, id string, roomname string, teamname string, opponentTeam string) {

// 	log.Println("challenge disucssion")
// 	log.Println("get power: ", getStoredPower)

// 	isTeamName, isTeamOpponent, proceed := false, false, false
// 	for _, r := range saveShuffle[roomname][teamname] {
// 		if r == id {
// 			isTeamName = true
// 		}
// 	}

// 	for _, r := range saveShuffle[roomname][opponentTeam] {
// 		if r == id {
// 			isTeamOpponent = true
// 		}
// 	}

// 	// store the visited team
// 	if isTeamName {
// 		store := map[string]map[string]bool{}
// 		store[roomname] = map[string]bool{
// 			opponentTeam: true,
// 		}
// 		Cwatch.done <- store
// 	} else {
// 		store := map[string]map[string]bool{}
// 		store[roomname] = map[string]bool{
// 			teamname: true,
// 		}
// 		Cwatch.done <- store
// 	}

// 	log.Println("challenge set: ", challengeToken)
// 	log.Println("get power: ", getStoredPower)
// 	log.Println("sessions: ", isSessionDone)
// 	switch true {
// 	case isTeamName && GR.ChallengeSet:
// 		log.Println("team red set challenge")
// 		// if the team B has arrived earlier
// 		// keep them on waiting and unblock the next team

// 		for _, _id := range saveShuffle[roomname][teamname] {
// 			if isSessionDone[_id][roomname][teamname][_ChallengeSessionKey] ||
// 				isSessionDone[_id][roomname][teamname][_BetSessionKey] {
// 				proceed = true
// 			}
// 		}
// 		// if isSessionDone[saveShuffle[roomname][teamname][0]][roomname][teamname][_ChallengeSessionKey] ||
// 		// 	isSessionDone[saveShuffle[roomname][teamname][1]][roomname][teamname][_ChallengeSessionKey] {
// 		// 	proceed = true
// 		// } else {
// 		// 	proceed = false
// 		// }
// 		// proceed to game room and note: to clear the cache of the stored team
// 		// if and only if the other team has come
// 		if proceed {
// 			ChallengeSessionTokens(h, teamname, opponentTeam, roomname)
// 		} else {

// 			// if isBetPowerUsed := IsBetSession(h, true, roomname); isBetPowerUsed {
// 			// 	ProceedBet(h, roomname, true)
// 			// } else if isBetPowerUsed := IsBetSession(h, false, roomname); isBetPowerUsed {
// 			// 	ProceedBet(h, roomname, false)
// 			// }
// 			go func() {
// 				store := map[string]map[string]map[string]map[string]bool{}
// 				store[id] = map[string]map[string]map[string]bool{roomname: {opponentTeam: {_ChallengeSessionKey: true}}}
// 				saveSessionDone.store <- store
// 			}()
// 			proceedBet := IsBetSession(h, true, roomname)
// 			betID := ""
// 			log.Println("proceed bet: ", proceedBet)
// 			// wait till the game start and store the set token
// 			for _, r := range saveShuffle[roomname][teamname] {
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: Unblock, to: r, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: waiting, to: r, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: RemoveChallenge + challengeToken, to: r, _sleep: false, roomname: roomname,
// 				}
// 				if isLock[teamname][roomname][r] {
// 					betID = r
// 				}
// 			}

// 			if proceedBet {
// 				dict := getStoredDict[roomname]
// 				SendBetTokens(h, betID, roomname, dict)

// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: BetPicker, to: betID, roomname: roomname, _sleep: false,
// 				}
// 			}

// 			for _, r := range saveShuffle[roomname][opponentTeam] {
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: Block, to: r, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: _waiting, to: r, _sleep: false, roomname: roomname,
// 				}
// 			}

// 			store := map[string]map[string]string{}
// 			store[opponentTeam] = map[string]string{
// 				roomname: challengeToken,
// 			}
// 			TChallengeSet.save <- store

// 			TrashChallenge(h, roomname, true, false)

// 		}

// 	case isTeamOpponent && GR.ChallengeSet:

// 		for _, _id := range saveShuffle[roomname][opponentTeam] {
// 			if isSessionDone[_id][roomname][opponentTeam][_ChallengeSessionKey] ||
// 				isSessionDone[_id][roomname][opponentTeam][_BetSessionKey] {
// 				proceed = true
// 			}
// 		}
// 		if proceed {
// 			ChallengeSessionTokens(h, teamname)
// 		} else {

// 			// if isBetPowerUsed := IsBetSession(h, true, roomname); isBetPowerUsed {
// 			// 	ProceedBet(h, roomname, true)
// 			// } else if isBetPowerUsed := IsBetSession(h, false, roomname); isBetPowerUsed {
// 			// 	ProceedBet(h, roomname, false)
// 			// }

// 			go func() {
// 				store := map[string]map[string]map[string]map[string]bool{}
// 				store[id] = map[string]map[string]map[string]bool{roomname: {teamname: {_ChallengeSessionKey: true}}}
// 				saveSessionDone.store <- store
// 			}()
// 			for _, r := range saveShuffle[roomname][teamname] {
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: Block, to: r, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: _waiting, to: r, _sleep: false, roomname: roomname,
// 				}
// 			}
// 			proceedBet := IsBetSession(h, false, roomname)
// 			betID := ""

// 			for _, r := range saveShuffle[roomname][opponentTeam] {
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: Unblock, to: r, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: waiting, to: r, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: RemoveChallenge + challengeToken, to: r, _sleep: false, roomname: roomname,
// 				}
// 				if isLock[opponentTeam][roomname][r] {
// 					betID = r
// 				}
// 			}

// 			if proceedBet {
// 				dict := getStoredDict[roomname]
// 				SendBetTokens(h, betID, roomname, dict)

// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: BetPicker, to: betID, roomname: roomname, _sleep: false,
// 				}
// 			}
// 			TrashChallenge(h, roomname, false, isTeamOpponent)

// 		}

// 		store := map[string]map[string]string{}
// 		store[teamname] = map[string]string{
// 			roomname: challengeToken,
// 		}
// 		TChallengeSet.save <- store
// 	}

// }

// ChallengeSession returns
// series1= toss won-> bet used->dictionary->bet
// series2= toss lost -> wait dict-> wait chal->bet
// series2= after set challenge check for the oppp
// series3= toss lost -> wait dict -> bet // according to the settings
// func ChallengeSession(h *Hub, id string, roomname string, teamname string, opponentTeam string, challengeToken string) {

// 	log.Println("challenge disucssion")
// 	log.Println("get power: ", getStoredPower)

// 	proceed, isTeamA := false, false

// 	for _, _id := range saveShuffle[roomname][teamname] {
// 		if _id == id {
// 			isTeamA = true
// 		}
// 	}

// 	if isTeamA {
// 		// store the visited team
// 		store := map[string]map[string]map[string]Session{}
// 		store[id] = map[string]map[string]Session{
// 			roomname: {teamname: Session{ChallengeDone: true}}}

// 		for _, _id := range saveShuffle[roomname][opponentTeam] {
// 			if getSessionUpdate[_id] != nil && getSessionUpdate[_id][roomname][opponentTeam].BetDone ||
// 				getSessionUpdate[_id][roomname][opponentTeam].ChallengeDone {
// 				proceed = true
// 			}
// 		}
// 		storeSessionUpdate <- store

// 	} else {
// 		store := map[string]map[string]map[string]Session{}
// 		store[id] = map[string]map[string]Session{
// 			roomname: {opponentTeam: Session{ChallengeDone: true}}}

// 		for _, _id := range saveShuffle[roomname][teamname] {
// 			if getSessionUpdate[_id] != nil && getSessionUpdate[_id][roomname][teamname].BetDone ||
// 				getSessionUpdate[_id][roomname][teamname].ChallengeDone {
// 				proceed = true
// 			}
// 		}

// 		storeSessionUpdate <- store

// 	}

// 	if proceed {
// 		ChallengeSessionTokens(h, id, roomname, challengeToken)
// 	} else {

// 		if isTeamA {
// 			// check if the oppponent has used the session update
// 			proceedBet := IsBetSession(h, opponentTeam, roomname)

// 			betID := ""

// 			log.Println("proceed bet: ", proceedBet)
// 			// wait till the game start and store the set token
// 			for _, _id := range saveShuffle[roomname][opponentTeam] {
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: Unblock, to: _id, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: waiting, to: _id, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: RemoveChallenge + challengeToken, to: _id, _sleep: false, roomname: roomname,
// 				}
// 				if getPlayerInfo[_id][roomname][teamname].isLock {
// 					betID = _id
// 				}
// 			}

// 			if proceedBet {
// 				dict := getGameInfo[roomname].SetDictionary
// 				SendBetTokens(h, betID, roomname, dict)

// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: BetPicker, to: betID, roomname: roomname, _sleep: false,
// 				}
// 			}

// 			for _, r := range saveShuffle[roomname][teamname] {
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: Block, to: r, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: _waiting, to: r, _sleep: false, roomname: roomname,
// 				}
// 			}

// 			store := map[string]map[string]map[string]Tokens{}
// 			store[id] = map[string]map[string]Tokens{
// 				roomname: {teamname: Tokens{Challenge: challengeToken}}}

// 			storeTokensUpdate <- store
// 			go func() {
// 				store := map[string]string{}
// 				store[roomname] = challengeToken

// 				addTeamRedChallenge <- store
// 			}()
// 		} else {
// 			proceedBet := IsBetSession(h, teamname, roomname)

// 			betID := ""

// 			log.Println("proceed bet: ", proceedBet)
// 			// wait till the game start and store the set token
// 			for _, _id := range saveShuffle[roomname][teamname] {
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: Unblock, to: _id, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: waiting, to: _id, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: RemoveChallenge + challengeToken, to: _id, _sleep: false, roomname: roomname,
// 				}
// 				if getPlayerInfo[_id][roomname][teamname].isLock {
// 					betID = _id
// 				}
// 			}

// 			if proceedBet {
// 				dict := getGameInfo[roomname].SetDictionary
// 				SendBetTokens(h, betID, roomname, dict)

// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: BetPicker, to: betID, roomname: roomname, _sleep: false,
// 				}
// 			}

// 			for _, r := range saveShuffle[roomname][opponentTeam] {
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: Block, to: r, _sleep: false, roomname: roomname,
// 				}
// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// 					token: _waiting, to: r, _sleep: false, roomname: roomname,
// 				}
// 			}

// 			store := map[string]map[string]map[string]Tokens{}
// 			store[id] = map[string]map[string]Tokens{
// 				roomname: {opponentTeam: Tokens{Challenge: challengeToken}}}

// 			storeTokensUpdate <- store
// 		}
// 		// go func() {
// 		// 	store := map[string]map[string]map[string]map[string]bool{}
// 		// 	store[id] = map[string]map[string]map[string]bool{roomname: {opponentTeam: {_ChallengeSessionKey: true}}}
// 		// 	saveSessionDone.store <- store
// 		// }()
// 		TrashChallenge(h, roomname, true, teamname, opponentTeam)
// 		go func() {
// 			store := map[string]string{}
// 			store[roomname] = challengeToken

// 			addTeamBlueChallenge <- store
// 		}()
// 	}

// }
// func CreateBetCups(h *Hub, RoomName string, isBetting bool, DictionaryName string) BetPattern {
// 	log.Println("in cups")

// 	log.Println("dictionary name: ", DictionaryName)

// 	dict := DictionaryName
// 	dicts := list.FetchSports(dict).Pack
// 	rand.Shuffle(len(dicts), func(i, j int) {
// 		dicts[i], dicts[j] = dicts[j], dicts[i]
// 	})

// 	_token := dicts[0]

// 	_conv := validate.FetchSports(dict, _token).Pack

// 	list := _conv[dict][_token]
// 	rand.Shuffle(len(list), func(i, j int) {
// 		list[i], list[j] = list[j], list[i]
// 	})

// 	log.Println("dictionary 📖 ", dict)
// 	firstCup := list[0]
// 	if !isBetting {
// 		firstCup = " " + getStoredBetValue[RoomName] + " "
// 	}
// 	secondCup := list[1]
// 	thirdCup := list[2]
// 	de := BetPattern{FirstCup: firstCup, SecondCup: secondCup, ThirdCup: thirdCup}

// 	go func() {
// 		KL := []string{list[0], list[1], list[2]}
// 		store := map[string][]string{}
// 		store[RoomName] = append(store[RoomName], KL...)
// 		storeCups.store <- store
// 	}()
// 	go func() {
// 		store := map[string][]string{}
// 		store[RoomName] = []string{firstCup, secondCup, thirdCup}
// 		storeCups.store <- store
// 	}()
// 	return de
// }

// //	type SessionWatcher struct {
// //		mu       sync.Mutex
// //		done     chan map[string]map[string]bool // first key must always be of room name
// //		complete chan bool
// //	}
// //
// //	type TrackSession struct {
// //		mu    sync.Mutex
// //		store chan map[string][]int
// //		done  chan bool
// //	}

// // type StoreClient struct {
// // 	mu    sync.Mutex
// // 	store chan int
// // 	done  chan bool
// // }
// // type trackLock struct {
// // 	mu     sync.Mutex
// // 	isLock chan map[string]map[string]map[string]bool
// // }

// // type TrackB struct {
// // 	store chan map[string][]string
// // 	done  chan bool
// // 	mu    sync.Mutex
// // }
// // type GameSession struct {
// // 	going chan map[string]bool
// // 	done  chan bool
// // }

// // type StoreRoomList struct {
// // 	store chan map[string]RoomList
// // 	done  chan bool
// // 	mu    sync.Mutex
// // 	wg    sync.WaitGroup
// // }

// // type StoreDict struct {
// // 	store chan map[string]string
// // 	done  chan bool
// // }

// // type TrackSessionDone struct {
// // 	store chan map[string]map[string]bool
// // 	done  chan bool
// // }

// var (

// // stores the name of the room and its limit
// // roomStore = StorageRoom{store: make(chan map[string]int), done: make(chan bool)}
// // gameSession        = GameSession{going: make(chan map[string]bool), done: make(chan bool)}
// // gameSessionStarted = make(map[string]bool)
// // rooms = []string{}
// // _tempGameSession   = map[string]bool{} // same as gameSession chann but this automatically starts the game

// // storeRoomList = StoreRoomList{store: make(chan map[string]RoomList), done: make(chan bool)}
// // getRoomList   = make(map[string]RoomList)
// // // checks if the passed token is not same as before
// // audit  = auditToken{token: make(chan []string), done: make(chan bool)}
// // _audit []string

// // storeCups     = TrackB{store: make(chan map[string][]string), done: make(chan bool)}
// // getStoredCups = map[string][]string{}

// // room-name: dict
// // storeDict     = StoreDict{store: make(chan map[string]string), done: make(chan bool)}
// // getStoredDict = map[string]string{}
// // checks if the passed token is same as before
// // mutualAudit  = mutualAuditToken{token: make(chan map[string][]string), done: make(chan bool)}
// // _mutualAudit = make(map[string][]string)

// // track the current lock
// // saves the team name, client id and if lock
// // TLock  = trackLock{isLock: make(chan map[string]map[string]map[string]bool)}
// // isLock = make(map[string]map[string]map[string]bool)

// // track the current block
// // saves the team name, client id and if block
// // TBlock  = make(chan map[string]map[string]bool)
// // isBlock = make(map[string]map[string]bool)

// // track if the lock process is done
// // __locked = make(chan bool)
// // _locked  bool

// // storing team name and coin side
// // Ttoss   = track{save: make(chan map[string]map[string]string), done: make(chan bool, 1)}
// // getToss = make(map[string]map[string]string)

// // storing team name and challenge set
// // TChallengeSet   = track{save: make(chan map[string]map[string]string), done: make(chan bool, 1)}
// // getChallengeSet = make(map[string]map[string]string)

// // team name: with their ids
// // todo make sure the room name also comes just like save shfuule
// // _REDTeam  = make(map[string][]string)
// // _BLUEteam = make(map[string][]string)

// // this makes sure that all players must complete this mutual session
// // Mwatch    = SessionWatcher{done: make(chan map[string]map[string]bool), complete: make(chan bool, 1)}
// // getMWatch = make(map[string]map[string]bool)

// // this makes sure that all players must complete this challenge session
// // Cwatch    = SessionWatcher{done: make(chan map[string]map[string]bool), complete: make(chan bool, 1)}
// // getCWatch = make(map[string]map[string]bool)

// // Dwatch    = SessionWatcher{done: make(chan map[string]map[string]bool), complete: make(chan bool, 1)}
// // getDWatch = make(map[string]map[string]bool)

// // // this makes sure that all players must complete this toss session
// // Twatch    = SessionWatcher{done: make(chan map[string]map[string]bool), complete: make(chan bool, 1)}
// // getTWatch = make(map[string]map[string]bool)

// // stores count of the round and set
// // note: to make things easy we are just adding 1 value to each
// // and using the length of these two to get the result
// // TtrackRound    = TrackSession{store: make(chan map[string][]int), done: make(chan bool)}
// // getTtrackRound = make(map[string]int)

// // TtrackSet    = TrackSession{store: make(chan map[string][]int), done: make(chan bool)}
// // getTtrackSet = make(map[string]int)

// // stores completion of session to make things easy for challenge and mutual
// // TTrackSessionDone    = make(chan map[string]map[string]bool) // room-name : challenge-name : true | false
// // getTTrackSessionDone = map[string]map[string]bool{}
// )
// // setToss              = "TossOff: true"               // remove the toss option from the game only if the setup toss is off
// // _DictionarySet = "DictionarySet: true" // singal to broadcast the dictionary event
// // _gameDone      = "GameDone: true"      // signal to game done
// // _roundOver     = "RoundOver: false"    // signal to start the round
// // start session
// // _tossSession          = "TossSession: true"          // on going toss session
// // _challengeDiscussion  = "ChallengeDiscussion: true"  // on going challenge set session
// // _gameBegin            = "GameBegin: false"           // on going game session

// // type BroadcastReq2 struct {
// // 	RoomID   string
// // 	Token    string
// // 	ClientID string
// // }
// // _ChallengeSessionKey = "ChallengeSesson"
// // _GameTimeKey     = "GameTime"
// // _DecisionTimeKey = "DecisionTime"
// // _Press = "Press: false"
// // _underTest = "UnderTest: false"
// // _BetPicker = "BetPicker: false"
// // _backClock = "ClockRestart: false"
// // _Betting   = "Betting: false"
// // for _, IDs := range saveShuffle[roomname][teamname] {
// // 	if getPlayerInfo[IDs][roomname][teamname].isLock {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			token: _tagIn, _sleep: false, to: IDs, roomname: roomname,
// // 		}
// // 		switched = IDs
// // 	} else {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			token: tagIn, _sleep: false, to: IDs, roomname: roomname,
// // 		}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _CanUsePower, _sleep: false,
// // 			to: IDs, roomname: roomname}

// // 		with_ = IDs
// // 	}
// // }

// // token := Alert + "Team Blue " + getNicknamesViaID[switched][roomname] + " Tagged " + getNicknamesViaID[with_][roomname]

// // for _, IDs := range saveShuffle[roomname][opponentTeamname] {
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, _sleep: false, to: IDs, roomname: roomname}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: TMsg, _sleep: false, to: IDs, roomname: roomname}
// // }

// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, to: ID, _sleep: false, roomname: roomname}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: TagUse, to: ID, _sleep: false, roomname: roomname}

// // locifyResetPower = make(chan map[string]map[string]IPowerReset) // resets the power of given id

// // locifyUpdateProfile = make(chan map[string]map[string]LocifyUpdateFixtures)

// // id->roomname->returns the key
// // locifyResetPowerUp = make(chan map[string]map[string]bool)

// // id->roomname->..
// // locifySaveStats  = make(chan map[string]map[string]LocfiySingleStatSave) // saves in profile id
// // locifySaveSingle = make(chan map[string]map[string]LocifySingleSave)     // saves in the oppoent id of the requested id

// // roomname->..
// // locifySaveCommon  = make(chan map[string]LocifyCommonSave)
// // locifySaveOneTime = make(chan map[string]LocifyOneTimeSave)
// // locifySaveGlobal  = make(chan map[string]LocifyMasterSave) // saves in all the ids of the global

// // // PBetMechanism returns done only left is sending stop request and some to do with the score
// // // teamname to send the token
// // func PBetMechanism(h *Hub, roomname, ID, teamname, opponentTeam, challengeToken, betOn string) {
// // 	// mechanism: the player gets the list of 3 items
// // 	//  from that three player ought to pick one
// // 	// and than these three items are send to the other opponent
// // 	// if he was to choosen the chosen words the player used the power wins the game

// // 	store := map[string]string{}
// // 	proceed := false

// // 	store[roomname] = betOn

// // 	// switch isTeamBlue {
// // 	// case true:
// // 	// 	store2[ID] = map[string]map[string]map[string]bool{
// // 	// 		roomname: {teamname: {_BetSessionKey: true}}}

// // 	// 	for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 	// 		log.Println("bet ids: ", _id)
// // 	// 		if isSessionDone[_id][roomname][teamname][_DictionarySessionKey] ||
// // 	// 			isSessionDone[_id][roomname][teamname][_ChallengeSessionKey] {
// // 	// 			proceed = true
// // 	// 		}
// // 	// 	}

// // 	// case false:
// // 	// 	store2[ID] = map[string]map[string]map[string]bool{
// // 	// 		roomname: {_TeamRedKey: {_BetSessionKey: true}}}
// // 	// 	for _, _id := range saveShuffle[roomname][teamname] {
// // 	// 		if isSessionDone[_id][roomname][teamname][_DictionarySessionKey] ||
// // 	// 			isSessionDone[_id][roomname][teamname][_ChallengeSessionKey] {
// // 	// 			proceed = true
// // 	// 		}
// // 	// 	}
// // 	// }

// // 	for _, _id := range saveShuffle[roomname][teamname] {

// // 		if getSessionUpdate[_id][roomname][teamname].DictionaryDone ||
// // 			getSessionUpdate[_id][roomname][teamname].ChallengeDone {
// // 			proceed = true
// // 		}
// // 	}

// // 	if !proceed {
// // 		for _, _id := range saveShuffle[roomname][opponentTeam] {
// // 			if getSessionUpdate[_id][roomname][opponentTeam].DictionaryDone ||
// // 				getSessionUpdate[_id][roomname][opponentTeam].ChallengeDone {
// // 				proceed = true
// // 			}
// // 		}
// // 	}

// // 	if proceed {
// // 		token := BetPattern{}
// // 		k := getStoredCups[roomname]
// // 		token.FirstCup = k[0]
// // 		token.SecondCup = k[1]
// // 		token.ThirdCup = k[2]
// // 		parcel, _ := json.Marshal(&token)

// // 		h.broadcast <- BroadcastReq{Token: ThreeBetCups + string(parcel), RoomID: roomname}

// // 		log.Println("parcel: ", string(parcel))
// // 		log.Println("stored cups: ", getStoredCups[roomname])

// // 		// if isTeamBlue {
// // 		// 	for _, r := range saveShuffle[roomname][teamname] {
// // 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Betting, to: r, roomname: roomname,
// // 		// 			_sleep: false}
// // 		// 		log.Println("id r: ", r)
// // 		// 	}

// // 		// } else {
// // 		// 	for _, r := range saveShuffle[roomname][_TeamBlueKey] {
// // 		// 		log.Println("id b: ", r)
// // 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Betting, to: r, roomname: roomname,
// // 		// 			_sleep: false}
// // 		// 	}
// // 		// }
// // 		for _, r := range saveShuffle[roomname][teamname] {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Betting, to: r, roomname: roomname,
// // 				_sleep: false}
// // 			log.Println("id r: ", r)
// // 		}

// // 		ChallengeSessionTokens(h, ID, roomname, challengeToken)
// // 	}

// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: BetUse, to: ID, _sleep: false, roomname: roomname}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _BetPicker, to: ID, _sleep: false, roomname: roomname}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, to: ID, _sleep: false, roomname: roomname}

// // 	storeBetValue.store <- store

// // }

// // GetPrivateRoom returns if found returns the name of the room

// // func GetPrivateRoom(code string) string {
// // 	Room := ""

// // 	for room, _code := range getPrivateRoomCode {
// // 		if code == _code {
// // 			Room = room
// // 		}
// // 	}
// // 	return Room
// // }

// // // PrivateRoomExists returns true if the private room exists via room name if specified else via matched code

// // func PrivateRoomExists(roomName string, _code string, viaCode bool) bool {
// // 	found := false
// // 	switch true {
// // 	case viaCode:
// // 		for _, code := range getPrivateRoomCode {
// // 			if code == _code {
// // 				found = true
// // 			}
// // 		}

// // 	case !viaCode:
// // 		if getPrivateRoom[roomName][_PrivateRoomKey] {
// // 			found = true
// // 		}
// // 	}

// // 	return found
// // }

// // GetPublicRoom returns  if found returns the name of the room

// // func GetPublicRoom(roomName string) string {
// // 	Room := ""
// // 	if PublicRoomExists(roomName) {
// // 		for room := range getPublicRoom {
// // 			if room == roomName {
// // 				Room = room
// // 			}
// // 		}
// // 	}
// // 	return Room
// // }

// // // PublicRoomExists returns true if the private room exists via room name if specified else via matched code
// // func PublicRoomExists(roomName string) bool {
// // 	found := getPublicRoom[roomName][_PublicRoomKey]

// // 	return found
// // }

// // // TwoVTwoRating returns winner Team new rating[both player-1 and player-2] and loser player rating respectively
// // // note: only a losser player that has lost the game
// // func TwoVTwoRating(WPreviousRating float64, LPreviousRating float64, WIdlePlayerRating float64) (float64, float64, float64) {
// // 	WRating, WIdleRating := CalcWinnerRating(WPreviousRating, LPreviousRating, WIdlePlayerRating)

// // 	LRating := CalcLosserRating(WPreviousRating, LPreviousRating)
// // 	return WRating, WIdleRating, LRating
// // }

// // func CalcLosserRating(WPreviousRating float64, LPreviousRating float64) float64 {

// // 	WRating, LRating := WPreviousRating, LPreviousRating

// // 	if WPreviousRating > LPreviousRating {
// // 		WRating = WPreviousRating
// // 		LRating = LPreviousRating
// // 	} else {
// // 		WRating = LPreviousRating
// // 		LRating = WPreviousRating
// // 	}
// // 	// note: the porbability is generated upon the  highest rated player to be a winner

// // 	// calculation of probablity of excepted win result
// // 	strength := -(WRating - LRating) / 400
// // 	pow := 1 + (math.Pow(10, strength))
// // 	calc := 1 / pow
// // 	ExcpetedScore := calc
// // 	// end of calculation

// // 	// upset // if the highest player looses
// // 	ExcpetedScore2 := 1 - ExcpetedScore

// // 	// case: for losser rating
// // 	NewLRating := math.Round(LRating + 32*(LOSSER-ExcpetedScore2))

// // 	return NewLRating
// // }

// // func CalcWinnerRating(WPreviousRating float64, LPreviousRating float64, WIdlePlayerRating float64) (float64, float64) {

// // 	WRating, LRating, IWRating := WPreviousRating, LPreviousRating, WIdlePlayerRating
// // 	if WPreviousRating > LPreviousRating {
// // 		WRating = WPreviousRating
// // 		LRating = LPreviousRating
// // 	} else {
// // 		WRating = LPreviousRating
// // 		LRating = WPreviousRating
// // 	}
// // 	// note: the porbability is generated upon the highest rated player to be a winner

// // 	// calculation of probablity of excepted win result
// // 	strength := -(WRating - LRating) / 400
// // 	pow := 1 + (math.Pow(10, strength))
// // 	calc := 1 / pow
// // 	ExcpetedScore := calc
// // 	// end of calculation

// // 	// constant: 32 can be change for different rating
// // 	// case: for winner rating
// // 	NewWRating := math.Round(WRating + 32*(WINNER-ExcpetedScore))

// // 	strength2 := -(IWRating - LRating) / 400
// // 	pow2 := 1 + (math.Pow(10, strength2))
// // 	calc2 := 1 / pow2
// // 	ExcpetedScore2 := calc2

// // 	NewW2Rating := math.Round(WIdlePlayerRating + 32*(WINNER-ExcpetedScore2))
// // 	return NewWRating, NewW2Rating
// // }

// // // OneVOneRating returns the winner and losser players new rating respectively
// // func OneVOneRating(WPreviousRating float64, LPreviousRating float64) (float64, float64) {

// // 	WRating, LRating := WPreviousRating, LPreviousRating
// // 	if WPreviousRating > LPreviousRating {
// // 		WRating = WPreviousRating
// // 		LRating = LPreviousRating
// // 	} else {
// // 		WRating = LPreviousRating
// // 		LRating = WPreviousRating
// // 	}
// // 	// note: the porbability is generated upon the  highest rated player to be a winner

// // 	// calculation of probablity of excepted win result
// // 	strength := -(WRating - LRating) / 400
// // 	pow := 1 + (math.Pow(10, strength))
// // 	calc := 1 / pow
// // 	ExcpetedScore := calc
// // 	// end of calculation

// // 	// upset // if the highest player looses
// // 	ExcpetedScore2 := 1 - ExcpetedScore

// // 	// constant: 32 can be change for different rating
// // 	// case: for winner rating
// // 	NewWRating := math.Round(WRating + 32*(WINNER-ExcpetedScore))
// // 	NewLRating := math.Round(LRating + 32*(LOSSER-ExcpetedScore2))
// // 	return NewWRating, NewLRating
// // }

// // // OneVOneRating_ returns the winner and losser players new rating respectively
// // func OneVOneRating_(WPreviousRating float64, LPreviousRating float64) (float64, float64) {

// // 	WRating, LRating, K1, K2 := WPreviousRating, LPreviousRating, 32.0, 1.0

// // 	// note: the porbability is generated upon the  highest rated player to be a winner
// // 	switch true {
// // 	case WRating < 900:
// // 		K1 = 96.16
// // 		log.Println("<900")

// // 	case WRating < 1400:
// // 		K1 = 105.0
// // 		log.Println("<1400")

// // 	case WRating < 2000:
// // 		K1 = 108.14
// // 		log.Println("<2000")

// // 	case WRating < 2200:
// // 		K1 = 112.6
// // 		log.Println("<2200")

// // 	case WRating < 2500:
// // 		K1 = 144.0
// // 		log.Println("<2500")
// // 	default:
// // 		K1 = 36.0
// // 	}

// // 	// calculation of probablity of excepted win result
// // 	strength := -(WRating - LRating) / 400
// // 	pow := 1 + (math.Pow(10, strength))
// // 	calc := 1 / pow
// // 	ExcpetedScore := calc
// // 	// end of calculation

// // 	// upset // if the highest player looses
// // 	ExcpetedScore2 := 1 - ExcpetedScore
// // 	K2 = K1 - 5

// // 	log.Println("k: ", K1)

// // 	// constant: 32 can be change for different rating
// // 	// case: for winner rating
// // 	NewWRating := math.Round(WRating + K1*(WINNER-ExcpetedScore))
// // 	NewLRating := math.Round(LRating + K2*(LOSSER-ExcpetedScore2))
// // 	log.Println("previous ratings: ", WRating, LRating)
// // 	log.Println("new ratings: ", NewWRating, "🔥", NewLRating, "🔥")

// // 	return NewWRating, NewLRating
// // }
// // func Result(h *Hub, Set int, TeamAScore int, TeamBScore int, GR GameRoom) {
// // 	log.Println("in result")

// // 	draw := Set == 4
// // 	winnerA := TeamAScore == 3
// // 	winnerB := TeamBScore == 3

// // 	switch true {
// // 	case draw:
// // 		h.wg.Go(func() {
// // 			// todo: send the new rating to both the teams and reregister
// // 			// match finished
// // 			h.broadcast <- BroadcastReq{Token: _gameBegin, RoomID: GR.RoomName}
// // 		})

// // 	case winnerA:
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			// todo: send the new rating to both the teams and reregister
// // 			// match finished
// // 			h.broadcast <- BroadcastReq{Token: _gameBegin, RoomID: GR.RoomName}
// // 		}()

// // 	case winnerB:
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			// todo: send the new rating to both the teams and reregister
// // 			// match finished
// // 			h.broadcast <- BroadcastReq{Token: _gameBegin, RoomID: GR.RoomName}
// // 		}()

// // 	default:
// // 		log.Println("game still going on")
// // 	}
// // }

// // func RoomsList(app *fiber.App, namespace string, h *Hub) {
// // 	log.Println("in room lists")
// // 	app.Use(func(c *fiber.Ctx) error {
// // 		if websocket.IsWebSocketUpgrade(c) {
// // 			c.Locals("allowed", true)
// // 			return c.Next()
// // 		}
// // 		// if c.Path() != "/ws" {
// // 		// 	return c.Next()
// // 		// }
// // 		return c.SendStatus(fiber.StatusUpgradeRequired)
// // 	})

// // 	app.Get(namespace, websocket.New(func(conn *websocket.Conn) {

// // 		log.Println("login successful")
// // 		defer func() {
// // 			conn.Close()
// // 		}()

// // 		token := RoomList{}
// // 		send, _ := json.Marshal(&token)
// // 		err := conn.WriteMessage(websocket.TextMessage, []byte(send))
// // 		if err != nil {

// // 			conn.WriteMessage(websocket.CloseMessage, []byte{})
// // 			conn.Close()
// // 		}
// // 		log.Println("token:", token)
// // 		log.Println("done broadcasting")

// // 	}))
// // }

// // var (
// // // stores nickname of associated room
// // // Snickname     = make(chan map[string]map[string]string) // room-name:id:nickname
// // // NickNamesList = make(map[string][]string)               // room-name: nicknames
// // )

// // type TFriend struct {
// // 	store chan map[string]map[string]bool // room-name: id: owner: true|false
// // 	mu    sync.Mutex
// // 	close chan bool
// // }

// // type StorePrivateRoom struct {
// // 	store      chan map[string]map[string]bool // room name: private room: true | false
// // 	roomCode   chan map[string]string          // room name: code
// // 	done       chan bool
// // 	removeCode chan bool
// // 	mu         sync.Mutex
// // }

// // type TURL struct {
// // 	store chan map[string]map[string]string
// // 	done  chan bool
// // 	mu    sync.Mutex
// // }
// // var (
// // storeRoomOwner = TFriend{store: make(chan map[string]map[string]bool), close: make(chan bool)}
// // getRoomOwner   = map[string]map[string]bool{}

// // storePrivateRoom   = StorePrivateRoom{store: make(chan map[string]map[string]bool), done: make(chan bool), removeCode: make(chan bool), roomCode: make(chan map[string]string)}
// // getPrivateRoom     = map[string]map[string]bool{}
// // getPrivateRoomCode = map[string]string{}

// // storePublicRoom = StoreFriendRoom{store: make(chan map[string]map[string]bool), done: make(chan bool), removeCode: make(chan bool), roomCode: make(chan map[string]string)}
// // getPublicRoom   = map[string]map[string]bool{}

// // storeRoomRequestedToken = TURL{store: make(chan map[string]map[string]string), done: make(chan bool)}
// // getRoomRequestedToken   = map[string]map[string]string{}
// // )

// // import "log"

// // func DeActivatePower(h *Hub, roomname string) {
// // 	log.Println("in deactive")
// // 	for _, ids := range saveShuffle[roomname][_TeamRedKey] {
// // 		switch true {
// // 		case getStoredPower[ids][roomname][_NexusKey]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _NexusKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][roomname][_DrawKey]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _DrawKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][roomname][_TagKey]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _TagKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][roomname][_FreezeKey]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _FreezeKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][roomname][_RewindKey]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _RewindKey + ": false", to: ids, _sleep: false, roomname: roomname}
// // 			h.broadcast <- BroadcastReq{Token: "RewindRestart: false", RoomID: roomname} // important

// // 		case getStoredPower[ids][roomname][_CovertKey]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _CovertKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][roomname][_BetKey]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _BetKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		}
// // 	}
// // 	for _, ids := range saveShuffle[roomname][_TeamBlueKey] {
// // 		switch true {
// // 		case getStoredPower[ids][_NexusKey][roomname]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _NexusKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][_DrawKey][roomname]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _DrawKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][_TagKey][roomname]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _TagKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][_FreezeKey][roomname]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _FreezeKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][_RewindKey][roomname]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _RewindKey + ": false", to: ids, _sleep: false, roomname: roomname}
// // 			h.broadcast <- BroadcastReq{Token: "RewindRestart: false", RoomID: roomname} // important

// // 		case getStoredPower[ids][_CovertKey][roomname]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _CovertKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		case getStoredPower[ids][roomname][_BetKey]:
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _BetKey + ": false", to: ids, _sleep: false, roomname: roomname}

// // 		}
// // 	}
// // }

// // // DeactiveAttribute deactive attributes such as bet picker
// // func DeactiveAttribute(h *Hub, roomname string) {

// // 	for _, teams := range saveShuffle {
// // 		for _, _ids := range teams {
// // 			for _, _id := range _ids {
// // 				switch true {

// // 				case getStoredPower[_id][roomname][_NexusKey]:
// // 					h.broadcast <- BroadcastReq{Token: NexusUse, RoomID: roomname}

// // 				case getStoredPower[_id][roomname][_FreezeKey]:
// // 					h.broadcast <- BroadcastReq{Token: _Press, RoomID: roomname}

// // 				case getStoredPower[_id][roomname][_CovertKey]:
// // 					h.broadcast <- BroadcastReq{Token: _underTest, RoomID: roomname}

// // 				case getStoredPower[_id][roomname][_RewindKey]:
// // 					h.broadcast <- BroadcastReq{Token: _backClock, RoomID: roomname}

// // 				case getStoredPower[_id][roomname][_BetKey]:
// // 					h.broadcast <- BroadcastReq{Token: _BetPicker, RoomID: roomname}
// // 					h.broadcast <- BroadcastReq{Token: _BetPicker, RoomID: roomname}
// // 					h.broadcast <- BroadcastReq{Token: _Betting, RoomID: roomname}
// // 				}
// // 			}
// // 		}
// // 	}

// // }

// // import (
// // 	"log"
// // )

// // // check if the given id has the betpower and if it is so send the tokens
// // // else search for the id and trigger it in instead of challenge

// // func DictonarySession(h *Hub, teamname string, opponentTeamName string, id string, roomname string, dictionaryToken string) {

// // 	log.Println("dictionary disucssion")
// // 	isTeamRed, isTeamBlue, RedBet, BlueBet := false, false, false, false

// // 	log.Println("sessions saved: ", getsaveSession)

// // 	for _, r := range saveShuffle[roomname][_TeamRedKey] {
// // 		if r == id {
// // 			isTeamRed = true

// // 		}
// // 	}

// // 	for _, r := range saveShuffle[roomname][_TeamBlueKey] {
// // 		if r == id {
// // 			isTeamBlue = true
// // 		}
// // 	}

// // 	RedBet = IsBetSession(h, teamname, roomname)
// // 	BlueBet = IsBetSession(h, teamname, roomname)

// // 	// go func() {
// // 	// 	store := map[string]string{}
// // 	// 	store[roomname] = dictionaryToken
// // 	// 	storeDictionaryEvent.store <- store
// // 	// }()
// // 	// if the first team to set the dictionary
// // 	// than second team can able to set the challenge
// // 	// or
// // 	// if the first team to set the dictornary
// // 	// than first team can able to set the challenge
// // 	if getSettings[roomname][_ReverseKey] {
// // 		// a1-b1 pattern: block the one that arrived and unblock the one that unable to make it and proceed him to set challenge
// // 		switch true {

// // 		case isTeamBlue:
// // 			if !RedBet {
// // 				for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _ChallengeSessionKey, roomname: roomname, _sleep: false, to: _id}
// // 				}
// // 			}
// // 			for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}

// // 			}
// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
// // 			}

// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: DictionaryEvent + dictionaryToken, roomname: roomname, _sleep: false, to: _id}
// // 			}
// // 			h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomID: roomname}
// // 			h.broadcast <- BroadcastReq{Token: _DictionarySet, RoomID: roomname}
// // 			// go func() {
// // 			// 	store := map[string]map[string]map[string]map[string]bool{}
// // 			// 	store[id] = map[string]map[string]map[string]bool{roomname: {_TeamBlueKey: {_DictionarySessionKey: true}}}
// // 			// 	saveSessionDone.store <- store
// // 			// }()

// // 		case isTeamRed:
// // 			if !BlueBet {
// // 				for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _ChallengeSessionKey, roomname: roomname, _sleep: false, to: _id}
// // 				}
// // 			}
// // 			for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
// // 			}
// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}
// // 			}
// // 			for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "DictionaryEvent: " + dictionaryToken, roomname: roomname, _sleep: false, to: _id}
// // 			}
// // 			h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomID: roomname}
// // 			h.broadcast <- BroadcastReq{Token: _DictionarySet, RoomID: roomname}
// // 			// go func() {
// // 			// 	store := map[string]map[string]map[string]map[string]bool{}
// // 			// 	store[id] = map[string]map[string]map[string]bool{roomname: {_TeamRedKey: {_DictionarySessionKey: true}}}
// // 			// 	saveSessionDone.store <- store
// // 			// }()
// // 		default:
// // 			log.Println("dictornay not able to decide")
// // 		}

// // 	} else {
// // 		// a1-a1 pattern: block the one that arriveed and block the one that unable to make it and proceed the one who has came here to set challenge
// // 		switch true {

// // 		case isTeamBlue:

// // 			if !BlueBet {
// // 				for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _challengeDiscussion, roomname: roomname, _sleep: false, to: _id}
// // 				}
// // 			}
// // 			for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
// // 			}
// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _challengeDiscussion, roomname: roomname, _sleep: false, to: _id}
// // 			}

// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "DictionaryEvent: " + dictionaryToken, roomname: roomname, _sleep: false, to: _id}
// // 			}

// // 			h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomID: roomname}
// // 			h.broadcast <- BroadcastReq{Token: _DictionarySet, RoomID: roomname}
// // 			h.broadcast <- BroadcastReq{Token: DictionaryEvent + dictionaryToken, RoomID: roomname}

// // 			// go func() {
// // 			// 	store := map[string]map[string]map[string]map[string]bool{}
// // 			// 	store[id] = map[string]map[string]map[string]bool{roomname: {_TeamBlueKey: {_DictionarySessionKey: true}}}
// // 			// 	saveSessionDone.store <- store
// // 			// }()
// // 		case isTeamRed:

// // 			if !RedBet {
// // 				for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _challengeDiscussion, roomname: roomname, _sleep: false, to: _id}
// // 				}
// // 			}
// // 			for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomname: roomname, _sleep: false, to: _id}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _challengeDiscussion, roomname: roomname, _sleep: false, to: _id}
// // 			}

// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomname: roomname, _sleep: false, to: _id}
// // 			}

// // 			h.broadcast <- BroadcastReq{Token: DictionaryEvent + dictionaryToken, RoomID: roomname}
// // 			h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomID: roomname}
// // 			h.broadcast <- BroadcastReq{Token: _DictionarySet, RoomID: roomname}

// // 			// go func() {
// // 			// 	store := map[string]map[string]map[string]map[string]bool{}
// // 			// 	store[id] = map[string]map[string]map[string]bool{roomname: {_TeamRedKey: {_DictionarySessionKey: true}}}
// // 			// 	saveSessionDone.store <- store
// // 			// }()
// // 		default:
// // 			log.Println("dictionay not able to decide")
// // 		}
// // 	}
// // 	log.Println("dictionary token: ", dictionaryToken)

// // 	token := SendList(dictionaryToken)
// // 	log.Println("token: ", token)

// // 	h.broadcast <- BroadcastReq{Token: ItemsURL + token, RoomID: roomname}

// // 	// go func() {
// // 	// 	store := map[string]string{}
// // 	// 	store[roomname] = dictionaryToken
// // 	// 	storeDict.store <- store
// // 	// }()

// // 	// if the bet used before the dictionary session
// // 	if !RedBet && !BlueBet {
// // 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: _challengeDiscussion}
// // 	} else {
// // 		if isTeamBlue {
// // 			ProceedBet(h, id, roomname, teamname, opponentTeamName, dictionaryToken)
// // 		} else {
// // 			ProceedBet(h, id, roomname, teamname, opponentTeamName, dictionaryToken)
// // 		}
// // 	}
// // 	go func() {
// // 		store := map[string]string{}
// // 		store[roomname] = dictionaryToken
// // 		addDictionary <- store
// // 	}()
// // 	go func() {
// // 		store := map[string]string{}
// // 		store[roomname] = token
// // 		addEvent <- store
// // 	}()

// // 	TrashDictionary(h, roomname, dictionaryToken)
// // 	// go func() {
// // 	// 	store := map[string]map[string]bool{}
// // 	// 	store[roomname] = map[string]bool{
// // 	// 		_DictionarySessionKey: true,
// // 	// 	}
// // 	// 	storeSessionDone.store <- store
// // 	// }()

// // }

// // func ASession(h *Hub, roomname string, toss, dictionary, challenge, game bool) {
// // 	var w ISessionUpdate
// // 	w.Challenge = challenge
// // 	w.Dictionary = dictionary
// // 	w.Game = game
// // 	c, err := json.Marshal(w)
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	token := SessionKey + string(c)
// // 	h.broadcast <- BroadcastReq{Token: token, RoomID: roomname}
// // }
// // // Monitor returns monitors if the room is full then starts the game
// // func Monitor(h *Hub, l Lobby, token []byte) {
// // 	log.Println("in monitor")

// // 	err := json.Unmarshal(token, &l)
// // 	if err != nil {
// // 		log.Println("err in lobby")
// // 		panic(err)
// // 	}
// // 	roomName, isPrivateRoom, isPublicRoom, isFriendRoom := l.To, false, true, false
// // 	for room, det := range getLocifyRoomSettings {
// // 		if l.Code == det.Code {
// // 			roomName = room
// // 			isPrivateRoom = det.Private
// // 			isFriendRoom = det.Friend
// // 		}
// // 	}

// // 	// switch true {

// // 	// case l.Code != "" && PrivateRoomExists("", l.Code, true):
// // 	// 	isPrivateRoom = true
// // 	// 	roomName = GetPrivateRoom(l.Code)

// // 	// case l.Code != "" && FriendRoomExists("", l.Code, true):
// // 	// 	isFriendRoom = true
// // 	// 	roomName = GetFriendRoom(l.Code)

// // 	// case l.Code == "" && FriendRoomExists(l.To, "", false):
// // 	// 	isFriendRoom = true
// // 	// }

// // 	switch true {
// // 	case isPrivateRoom:
// // 		log.Println("private room monitoring")
// // 		privateRoomCount := storeRoom[roomName]
// // 		privateMemberCount := CountRoomMembers(h, roomName)
// // 		log.Println("monitoring")
// // 		log.Println("storeroom[roomname]: ", storeRoom[roomName])
// // 		log.Println("store room: ", storeRoom)
// // 		log.Println("privar room count: ", privateRoomCount, "private member count: ", privateMemberCount)
// // 		if privateMemberCount >= 1 {
// // 			log.Println("client joined and his room capacity: ", l.RoomCapacity)
// // 		}
// // 		switch true {
// // 		case privateRoomCount == privateMemberCount:
// // 			// if player left between the session and other
// // 			// tried to join room than signal others that the game session is going on
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store := make(map[string]bool)
// // 				store[roomName] = true
// // 				gameSession.going <- store
// // 				h.broadcast <- BroadcastReq{RoomID: roomName, Token: gameBegin}
// // 			}()
// // 			_tempGameSession[roomName] = true

// // 			log.Println("game will start soon")

// // 		// case !PrivateRoomExists("", l.Code, true):
// // 		// 	log.Println("room doesnt exists")

// // 		case privateMemberCount > privateRoomCount:
// // 			log.Println("join is full and the players are in game session")

// // 		default:
// // 			log.Println("need", privateMemberCount-1, "members to start the game")
// // 			Conns = strconv.Itoa(privateMemberCount - 1)
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				h.broadcast <- BroadcastReq{RoomID: l.To, Token: _gameBegin}
// // 				h.broadcast <- BroadcastReq{RoomID: l.To, Token: "Searching: " + Conns}
// // 			}()
// // 		}
// // 	case isPublicRoom:

// // 		log.Println("public room monitoring")
// // 		publicRoomCount := storeRoom[roomName]
// // 		publicMemberCount := CountRoomMembers(h, roomName)
// // 		log.Println("monitoring")
// // 		log.Println("storeroom[roomname]: ", storeRoom[roomName])
// // 		log.Println("store room: ", storeRoom)
// // 		log.Println("public room name: ", roomName)
// // 		if publicMemberCount >= 1 {
// // 			log.Println("client joined and his room capacity: ", l.RoomCapacity)
// // 		}
// // 		switch true {
// // 		case publicRoomCount == publicMemberCount:
// // 			// if player left between the session and other
// // 			// tried to join room than signal others that the game session is going on
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store := make(map[string]bool)
// // 				store[roomName] = true
// // 				gameSession.going <- store
// // 				h.broadcast <- BroadcastReq{RoomID: roomName, Token: gameBegin}
// // 			}()
// // 			_tempGameSession[roomName] = true
// // 			log.Println("game will start soon")

// // 		// case !PrivateRoomExists("", l.Code, true):
// // 		// 	log.Println("room doesnt exists")

// // 		case publicMemberCount > publicRoomCount:
// // 			log.Println("join is full and the players are in game session")

// // 		default:
// // 			log.Println("need", publicMemberCount-1, "members to start the game")
// // 			Conns = strconv.Itoa(publicMemberCount - 1)
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				h.broadcast <- BroadcastReq{RoomID: l.To, Token: _gameBegin}
// // 				h.broadcast <- BroadcastReq{RoomID: l.To, Token: "Searching: " + Conns}
// // 			}()
// // 		}

// // 	case isFriendRoom:
// // 		log.Println("friend room monitoring")

// // 		friendRoomCount := storeRoom[roomName]
// // 		friendMemberCount := CountRoomMembers(h, roomName)
// // 		if friendMemberCount >= 1 {
// // 			log.Println("client joined and his room capacity: ", l.RoomCapacity)
// // 		}
// // 		switch true {
// // 		case friendRoomCount == friendMemberCount:
// // 			// if player left between the session and other
// // 			// tried to join room than signal others that the game session is going on
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store := make(map[string]bool)
// // 				store[roomName] = true
// // 				gameSession.going <- store
// // 				h.broadcast <- BroadcastReq{RoomID: roomName, Token: gameBegin}
// // 			}()
// // 			_tempGameSession[roomName] = true

// // 			log.Println("game will start soon")

// // 		// case !FriendRoomExists("", l.Code, true):
// // 		// 	log.Println("room doesnt exists")

// // 		case friendMemberCount > friendRoomCount:
// // 			log.Println("join is full and the players are in game session")

// // 		default:
// // 			log.Println("need", friendMemberCount-1, "members to start the game")
// // 			Conns = strconv.Itoa(friendMemberCount - 1)
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				h.broadcast <- BroadcastReq{RoomID: l.To, Token: _gameBegin}
// // 				h.broadcast <- BroadcastReq{RoomID: l.To, Token: "Searching: " + Conns}
// // 			}()
// // 		}
// // 	}

// // 	log.Println("monitor done")
// // }

// // /*************
// // * ALGORITHM FOR NEXUS POWER
// // ********/

// // // Nexus returns  @PARAM str char to replace in
// // // @PARAM replace any string that incldues no whitespace
// // // @RETURNS replaces the give char at any k limit at random rolls
// // // tip: you use replace="_" for better practice
// // // @EXAMPLE  input: king ouput: k _ n _ // or any other seq

// // func Nexus(str string, replace string) string {
// // 	m, _ := regexp.Match(`\s`, []byte(replace))
// // 	log.Println("m: ", m)
// // 	if m {
// // 		panic("ws not allowed")
// // 	}
// // 	// formula: replaces the char at any k limit of random rolls
// // 	_limit, rolls := []int{}, []int{}
// // 	for r := range len(str) {
// // 		rolls = append(rolls, r)
// // 		_limit = append(_limit, r+1)
// // 	}
// // 	rand.Shuffle(len(_limit), func(i int, j int) {
// // 		_limit[i], _limit[j] = _limit[j], _limit[i]
// // 	})
// // 	rand.Shuffle(len(rolls), func(i int, j int) {
// // 		rolls[i], rolls[j] = rolls[j], rolls[i]
// // 	})

// // 	for roll := range _limit[0] {
// // 		re := string(str[rolls[roll]]) // char to string conv: in-order to replace recursively
// // 		// r-1 is important
// // 		// reason: it helps in replacing the value at rolled
// // 		str = strings.Replace(str, re, replace, roll-1) // note: if you replace with " _ " it would cause index error
// // 	}

// // 	return strings.Join(strings.Split(str, ""), " ") // format into: king to k i n g

// // }

// // type SStoreNames struct {
// // 	viaRoom chan map[string]string            // room-name: list of nicknames
// // 	viaID   chan map[string]map[string]string // ID : room-name: nickname
// // 	done    chan bool
// // 	mu      sync.Mutex
// // }

// // var (
// // 	storeNickNames      = SStoreNames{viaRoom: make(chan map[string]string), done: make(chan bool), viaID: make(chan map[string]map[string]string)}
// // 	getNicknamesViaID   = make(map[string]map[string]string)
// // 	getNicknamesViaRoom = make(map[string][]string)
// // )

// // const (
// // 	PenaltyKey  = "Penalty: Max"
// // 	_PenaltyKey = "Penalty: "
// // 	_NewTimeKey = "NewTime: "
// // )

// // type TtrackPenalty struct {
// // 	store chan map[string]map[string]map[string]int
// // 	done  chan bool
// // 	mu    sync.Mutex
// // 	wg    sync.WaitGroup
// // }

// // var (
// // 	// id: room-name: team-name: penalty count
// // 	storePenalty    = TtrackPenalty{store: make(chan map[string]map[string]map[string]int), done: make(chan bool)}
// // 	getStorePenalty = make(map[string]map[string]map[string][]int)
// // )

// // func Penalty(h *Hub, roomname string, id string) {
// // 	isTeamBlue := false
// // 	_token := getRoomRequestedToken[roomname][_GameTimeKey]
// // 	gameTime, _ := strconv.Atoi(_token)

// // 	for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 		if id == _id {
// // 			isTeamBlue = true
// // 		}
// // 	}
// // 	switch isTeamBlue {
// // 	case true:
// // 		count := 0
// // 		for _, i := range getStorePenalty[id][roomname][_TeamBlueKey] {
// // 			count += i
// // 		}

// // 		switch count {
// // 		case 1:
// // 			gameTime -= 2 // -2
// // 			tokenx := _NewTimeKey + strconv.Itoa(gameTime)
// // 			// note: sending both to avoid complication for tag and draw
// // 			for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: tokenx, roomname: roomname,
// // 					to: _id, _sleep: false}
// // 			}

// // 		case 2:
// // 			gameTime = 3
// // 			tokenx := _NewTimeKey + strconv.Itoa(gameTime)
// // 			// note: sending both to avoid complication for tag and draw
// // 			for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: tokenx, roomname: roomname,
// // 					to: _id, _sleep: false}
// // 			}

// // 		default:
// // 			// max reached
// // 			for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PenaltyKey, roomname: roomname,
// // 					to: _id, _sleep: false}
// // 			}
// // 			// game done
// // 			// todo make sure to send the result too
// // 			h.broadcast <- BroadcastReq{Token: _gameDone, RoomID: roomname}
// // 		}
// // 		store := map[string]map[string]map[string]int{}
// // 		store[id] = map[string]map[string]int{
// // 			roomname: {
// // 				_TeamBlueKey: 1,
// // 			},
// // 		}
// // 		storePenalty.store <- store

// // 	case false:
// // 		count := 0
// // 		for _, i := range getStorePenalty[id][roomname][_TeamRedKey] {
// // 			count += i
// // 		}

// // 		switch count {
// // 		case 1:
// // 			gameTime -= 2 // -2
// // 			tokenx := _NewTimeKey + strconv.Itoa(gameTime)
// // 			// note: sending both to avoid complication for tag and draw
// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: tokenx, roomname: roomname,
// // 					to: _id, _sleep: false}
// // 			}

// // 		case 2:
// // 			gameTime = 3
// // 			tokenx := _NewTimeKey + strconv.Itoa(gameTime)
// // 			// note: sending both to avoid complication for tag and draw
// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: tokenx, roomname: roomname,
// // 					to: _id, _sleep: false}
// // 			}

// // 		default:
// // 			// max reached
// // 			for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PenaltyKey, roomname: roomname,
// // 					to: _id, _sleep: false}
// // 			}
// // 			// game done
// // 			// todo make sure to send the result too
// // 			h.broadcast <- BroadcastReq{Token: _gameDone, RoomID: roomname}
// // 		}
// // 		store := map[string]map[string]map[string]int{}
// // 		store[id] = map[string]map[string]int{
// // 			roomname: {
// // 				_TeamRedKey: 1,
// // 			},
// // 		}
// // 		storePenalty.store <- store
// // 	}
// // }

// // func Game(h *Hub, conn *websocket.Conn, GR GameRoom, id string) {
// // 	//h.wg.Wait()
// // 	log.Println("game")
// // 	// @TODO: make sure the challenge token is not same as the previous one
// // 	//
// // 	// pattern for tagging
// // 	count := 0
// // 	roomname := GR.RoomName
// // 	// @IMPORTANT
// // 	ID := h.clients[conn].clientID
// // 	// hasKey := valIDate.Has_Key(GR.ChallengeToken)
// // 	teamname, opponentTeamname := TokenFromTeam(id, roomname)
// // 	myteam := ""
// // 	_w := false
// // 	for _, _id := range saveShuffle[roomname][teamname] {
// // 		if _id == id {
// // 			_w = true
// // 		}
// // 	}
// // 	if _w {
// // 		myteam = teamname
// // 	} else {
// // 		myteam = opponentTeamname
// // 	}

// // 	// triggers only if the room session has started
// // 	if gameSessionStarted[roomname] {
// // 		log.Println("draw session: ", GR.DrawSession)
// // 		// saving the token at beginning to avoID loss of time
// // 		if GR.PowerActivated {
// // 			log.Println("power activated")
// // 			PocketPowers(h, ID, roomname, GR)
// // 			// send the offer asap
// // 			switch true {
// // 			case GR.Draw:
// // 				log.Println("draw used")

// // 				if _w {
// // 					for _, ID := range saveShuffle[roomname][opponentTeamname] {
// // 						if getPlayerInfo[ID][roomname][opponentTeamname].isLock {
// // 							h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 								token: drawOffer, _sleep: false, to: ID, roomname: roomname,
// // 							}
// // 						}
// // 					}
// // 				} else {
// // 					for _, ID := range saveShuffle[roomname][teamname] {
// // 						if getPlayerInfo[ID][roomname][teamname].isLock {
// // 							h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 								token: drawOffer, _sleep: false, to: ID, roomname: roomname,
// // 							}
// // 						}
// // 					}
// // 				}

// // 			case GR.Tag:

// // 				PTagMechanism(h, roomname, teamname, opponentTeamname, ID)

// // 			case GR.Covert:

// // 				store := map[string]map[string]map[string]bool{}
// // 				store[ID] = map[string]map[string]bool{
// // 					roomname: {_CovertKey: true},
// // 				}

// // 				storePower.store <- store
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: CovertUse, to: ID, roomname: roomname, _sleep: false,
// // 				}

// // 				// if used after the dictionary session
// // 			case GR.Bet:
// // 				log.Println("bet used")
// // 				if getSessionUpdate[id][roomname][teamname].DictionaryDone {

// // 					doneChallenge := false
// // 					switch _w {
// // 					case true:
// // 						for _, _id := range saveShuffle[roomname][opponentTeamname] {
// // 							if getSessionUpdate[_id][roomname][opponentTeamname].ChallengeDone {
// // 								doneChallenge = true
// // 							}
// // 						}

// // 					case false:
// // 						for _, _id := range saveShuffle[roomname][myteam] {
// // 							if getSessionUpdate[_id][roomname][myteam].ChallengeDone {
// // 								doneChallenge = true
// // 							}
// // 						}
// // 					}

// // 					if doneChallenge {
// // 						BetSession(h, ID, roomname, GR.ChallengeToken, GR.BetOn)
// // 						dict := getGameInfo[roomname].SetDictionary
// // 						SendBetTokens(h, ID, roomname, dict)
// // 						BetSession(h, ID, roomname, GR.ChallengeToken, GR.BetOn)
// // 					} else {
// // 						dict := getGameInfo[roomname].SetDictionary
// // 						ProceedBet(h, id, roomname, teamname, opponentTeamname, dict)
// // 					}
// // 				}

// // 			case GR.Session && GR.Draw:

// // 				isTeamBlue := false
// // 				for _, IDs := range saveShuffle[roomname][_TeamBlueKey] {
// // 					if ID == IDs {
// // 						isTeamBlue = true
// // 					}
// // 				}

// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: DrawUse, _sleep: false,
// // 					to: ID, roomname: roomname,
// // 				}

// // 				switch isTeamBlue {
// // 				case true:
// // 					for _, IDs := range saveShuffle[roomname][_TeamRedKey] {
// // 						if ID == IDs {
// // 							isTeamBlue = true
// // 						}
// // 						if getPlayerInfo[IDs][roomname][_TeamRedKey].isLock {
// // 							h.gameRoomBroadcast <- reqGameRoomBroadcast{token: drawOffer, _sleep: false,
// // 								to: IDs, roomname: roomname,
// // 							}
// // 						}
// // 					}
// // 				case false:
// // 					for _, IDs := range saveShuffle[roomname][_TeamBlueKey] {
// // 						if getPlayerInfo[IDs][roomname][_TeamBlueKey].isLock {
// // 							h.gameRoomBroadcast <- reqGameRoomBroadcast{token: drawOffer, _sleep: false,
// // 								to: IDs, roomname: roomname,
// // 							}
// // 						}
// // 					}
// // 				}

// // 			case GR.Session && GR.Unfreeze:
// // 				log.Println("unfreezing")
// // 				h.broadcast <- BroadcastReq{Token: Unfreeze, RoomID: roomname}
// // 				h.broadcast <- BroadcastReq{Token: _Press, RoomID: roomname}
// // 				// get the game time
// // 				GameTime := getRoomRequestedToken[roomname][_GameTimeKey]
// // 				h.broadcast <- BroadcastReq{Token: "RoomGameTime: " + GameTime, RoomID: roomname}

// // 			case GR.Session && GR.DrawSession && GR.DrawAccept:

// // 				isTeamBlue := false
// // 				for _, r := range saveShuffle[roomname][_TeamBlueKey] {
// // 					if ID == r {
// // 						isTeamBlue = true
// // 					}
// // 				}
// // 				PDrawMechanism(h, roomname, ID, isTeamBlue, true)

// // 			case GR.Session && GR.DrawSession && !GR.DrawAccept:

// // 				isTeamBlue := false
// // 				for _, r := range saveShuffle[roomname][_TeamBlueKey] {
// // 					if ID == r {
// // 						isTeamBlue = true
// // 					}
// // 				}
// // 				PDrawMechanism(h, roomname, ID, isTeamBlue, false)
// // 			}
// // 		} else {
// // 			switch true {

// // 			case !GR.Start && GR.TossSession:
// // 				log.Println("Heads tails: ", GR.HeadTails)
// // 				Toss(h, ID, GR.RoomName, GR.HeadTails)

// // 			case !GR.Start && GR.DictionarySession:
// // 				DictonarySession(h, teamname, opponentTeamname, id, GR.RoomName, GR.DictionaryToken)

// // 			case !GR.Start && GR.ChallengeDiscussion:
// // 				LocifyChallengeSession(h, ID, roomname, GR.ChallengeToken, count)

// // 			case !GR.Start && GR.BetSession && !GR.PowerActivated:
// // 				BetSession(h, ID, roomname, GR.ChallengeToken, GR.BetOn)

// // 			case !GR.Start && GR.DTimeUp:

// // 			case GR.Start && !GR.Unfreeze:
// // 				Score(h, count, ID, conn, GR)

// // 			default:
// // 				log.Println("not able to find play pattern")
// // 			}
// // 		}
// // 	} else {
// // 		log.Println("session not begin")

// // 	}
// // }

// // // TokenFromTeam returns fist team of the matched id and second team of the opponent
// // func TokenFromTeam(id string, roomname string) (string, string) {

// // 	FromTeamRed := slices.Contains(saveShuffle[roomname][_TeamBlueKey], id)
// // 	TeamMatch := ""
// // 	Opponent := ""

// // 	if FromTeamRed {
// // 		TeamMatch = _TeamRedKey
// // 		Opponent = _TeamBlueKey
// // 	} else {
// // 		TeamMatch = _TeamBlueKey
// // 		Opponent = _TeamRedKey
// // 	}

// // 	return TeamMatch, Opponent
// // }

// // // PatternIRA i= ID, r=room-name a=any [can be session name, power name or any other name]
// // type PatternIRA struct {
// // 	store chan map[string]map[string]map[string]bool // player-ID: room-name: power-name: true | false
// // 	mu    sync.Mutex
// // 	done  chan bool
// // }

// // var (
// // 	storePower     = PatternIRA{store: make(chan map[string]map[string]map[string]bool), done: make(chan bool)} // player-ID: room-name: power-name: true | false
// // 	getStoredPower = make(map[string]map[string]map[string]bool)
// // )

// // type SBet struct {
// // 	store chan map[string]map[string]map[string]string // room name: team-name: :dictionary-name: value
// // 	mu    sync.Mutex
// // 	done  chan bool
// // }
// // type SDictionary struct {
// // 	store chan map[string]string // room name: value
// // 	mu    sync.Mutex
// // 	done  chan bool
// // }

// // type SBetEvent struct {
// // 	store chan map[string]map[string]string // room name: team-name: value
// // 	mu    sync.Mutex
// // 	done  chan bool
// // }
// // type TrackBetUsed struct {
// // 	store chan map[string]map[string]map[string]bool
// // 	done  chan bool
// // }

// // type TrackCups struct {
// // 	store chan map[string][]string
// // 	done  chan bool
// // }

// // var (
// // 	// room-name: team-name: dictionary-name: gives bet-value
// // 	storeBetValue     = SDictionary{store: make(chan map[string]string), done: make(chan bool)}
// // 	getStoredBetValue = make(map[string]string)

// // 	// room-name: team-name: dictionary-name
// // 	// storeDictionaryEvent     = SDictionary{store: make(chan map[string]string), done: make(chan bool)}
// // 	// getStoredDictionaryEvent = make(map[string]string)

// // 	// room name: team-name: gives dictionary-name
// // 	storeBetEvent     = SDictionary{store: make(chan map[string]string), done: make(chan bool)}
// // 	getStoredBetEvent = make(map[string]string)
// // )

// // // Toss returns to do  make sure to implement more complex later
// // func Toss(h *Hub, id string, roomname string, HeadTails string) {

// // 	// logic:
// // 	// unblock player picks the coin side either heads or tails
// // 	// than player that is block will toss the coin
// // 	// if player blue toss the coin we will look into the cases of red becuase the side is pick by the red
// // 	// todo make sure to change the result of the toss

// // 	log.Println("toss")

// // 	isTeamRed := false
// // 	isTeamBlue := false
// // 	proceed := false
// // 	ishead := false
// // 	foundB := ""
// // 	se := Session{TossDone: true, ChallengeDone: false, DictionaryDone: false, BetDone: false, SessionDone: false}
// // 	roll := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}

// // 	rand.Shuffle(len(roll), func(i int, j int) {
// // 		roll[i], roll[j] = roll[j], roll[i]
// // 	})

// // 	if roll[0]%2 == 0 {
// // 		ishead = true
// // 	} else {
// // 		ishead = false
// // 	}

// // 	for _, r := range saveShuffle[roomname][_TeamRedKey] {
// // 		if r == id {
// // 			isTeamRed = true
// // 		}
// // 	}

// // 	for _, r := range saveShuffle[roomname][_TeamBlueKey] {
// // 		if r == id {
// // 			isTeamBlue = true
// // 		}
// // 	}

// // 	// if isTeamRed {
// // 	// 	// make sure that the coin has not tossed yet
// // 	// 	if GR.HeadTails != "" {
// // 	// 		h.wg.Add(1)
// // 	// 		go func() {
// // 	// 			defer h.wg.Done()
// // 	// 			store := map[string]map[string]bool{}
// // 	// 			store[roomname] = map[string]bool{
// // 	// 				_TeamRedKey: true,
// // 	// 			}
// // 	// 			Twatch.done <- store
// // 	// 		}()

// // 	// 	}
// // 	// } else {
// // 	// 	// make sure that the coin has not tossed yet
// // 	// 	if GR.HeadTails != "" {
// // 	// 		h.wg.Add(1)
// // 	// 		go func() {
// // 	// 			defer h.wg.Done()
// // 	// 			store := map[string]map[string]bool{}
// // 	// 			store[roomname] = map[string]bool{
// // 	// 				_TeamBlueKey: true,
// // 	// 			}
// // 	// 			Twatch.done <- store
// // 	// 		}()

// // 	// 		h.wg.Add(1)
// // 	// 		go func() {
// // 	// 			defer h.wg.Done()
// // 	// 			store := map[string]map[string]string{}
// // 	// 			store[roomname] = map[string]string{
// // 	// 				_TeamBlueKey: GR.HeadTails,
// // 	// 			}
// // 	// 			Ttoss.save <- store
// // 	// 		}()
// // 	// 	}
// // 	// }

// // 	TossBlueWin := getGameInfo[roomname].SetTossBody[_TeamBlueKey] == "HEADS" && ishead || getGameInfo[roomname].SetTossBody[_TeamBlueKey] == "TAILS" && !ishead
// // 	TossRedWin := getGameInfo[roomname].SetTossBody[_TeamRedKey] == "HEADS" && ishead || getGameInfo[roomname].SetTossBody[_TeamRedKey] == "TAILS" && !ishead

// // 	// log.Println("get watch: ", getTWatch)
// // 	log.Println("game info: ", getGameInfo[roomname])
// // 	switch true {
// // 	// coin toss by blue
// // 	case isTeamRed:
// // 		for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 			if getSessionUpdate[_id][roomname][_TeamBlueKey].TossDone {
// // 				proceed = true
// // 				break
// // 			}
// // 		}
// // 		if proceed {
// // 			book := getRoomRequestedToken[roomname][_BookKey]
// // 			sendDictionary := DictionaryURL + SendDictionary(book)

// // 			// checking if the body of the toss matches
// // 			switch TossBlueWin {

// // 			// matched
// // 			case true:
// // 				youWon := TossMsg + "YOU WON NOW YOU CAN SET THE DICTIONARY"

// // 				log.Println("not sending round 1 toss")
// // 				// toss loss than blue will set the dictionary
// // 				for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, token: _DictionaryDiscussion, to: _id, _sleep: false,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, to: _id, _sleep: false, token: Unblock,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, to: _id, _sleep: false, token: youWon,
// // 					}
// // 				}
// // 				tossWin := TossMsg + "WIN BY BLUE"
// // 				for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 					if getPlayerInfo[_id][roomname][_TeamRedKey].isBlock {
// // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 							roomname: roomname, to: _id, _sleep: false, token: Block,
// // 						}
// // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 							roomname: roomname, to: _id, _sleep: false, token: tossWin,
// // 						}
// // 					}
// // 				}

// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: TossAlert}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: tossSession}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: sendDictionary}

// // 				// not matched
// // 			default:
// // 				log.Println("default toss")
// // 				log.Println("foundB: ", foundB)
// // 				log.Println("not sending round 1 toss")
// // 				youWon := TossMsg + "YOU WON NOW YOU CAN SET THE DICTIONARY"

// // 				for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, token: _DictionaryDiscussion, to: _id, _sleep: false,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, to: _id, _sleep: false, token: Unblock,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, to: _id, _sleep: false, token: youWon,
// // 					}
// // 				}

// // 				tossWin := TossMsg + "WIN BY RED"
// // 				// block the toss loss team
// // 				for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 					if getPlayerInfo[_id][roomname][_TeamBlueKey].isBlock {
// // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 							roomname: roomname, to: _id, _sleep: false, token: Block,
// // 						}
// // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 							roomname: roomname, to: _id, _sleep: false, token: tossWin,
// // 						}
// // 					}
// // 				}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: TossAlert}

// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: sendDictionary}
// // 				// toss session done
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: tossSession}
// // 			}

// // 			h.broadcast <- BroadcastReq{Token: _toss, RoomID: roomname}
// // 		} else {
// // 			for _, r := range saveShuffle[roomname][_TeamBlueKey] {
// // 				log.Println("ids: ", r)
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: Unblock,
// // 				}
// // 				// toss
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: _toss,
// // 				}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: _tossCoin,
// // 				}
// // 			}
// // 			h.wg.Add(1)
// // 			go func() {
// // 				store := map[string]map[string]map[string]Session{}
// // 				store[id] = map[string]map[string]Session{
// // 					roomname: {_TeamRedKey: se},
// // 				}
// // 				storeSessionUpdate <- store
// // 			}()

// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store := map[string]GameInfo{}
// // 				ma := map[string]string{}
// // 				ma[_TeamRedKey] = HeadTails
// // 				store[roomname] = GameInfo{SetTossBody: ma}
// // 				storeGameInfo <- store

// // 			}()

// // 			// h.wg.Add(1)
// // 			// go func() {
// // 			// 	defer h.wg.Done()
// // 			// 	store := map[string]map[string]bool{}
// // 			// 	store[roomname] = map[string]bool{_TeamRedKey: true}
// // 			// 	Twatch.done <- store
// // 			// }()
// // 			// h.wg.Add(1)
// // 			// go func() {
// // 			// 	defer h.wg.Done()
// // 			// 	store := map[string]map[string]string{}
// // 			// 	store[roomname] = map[string]string{
// // 			// 		_TeamRedKey: GR.HeadTails,
// // 			// 	}
// // 			// 	Ttoss.save <- store
// // 			// }()

// // 			for _, r := range saveShuffle[roomname][_TeamRedKey] {
// // 				log.Println("ids: ", r)

// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: Block,
// // 				}
// // 				// toss pick done
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: toss,
// // 				}
// // 			}

// // 		}

// // 	case isTeamBlue:
// // 		for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 			if getSessionUpdate[_id][roomname][_TeamRedKey].TossDone {
// // 				proceed = true
// // 				break
// // 			}
// // 		}
// // 		if proceed {
// // 			sendDictionary := DictionaryURL + SendDictionary(roomname)

// // 			switch TossRedWin {
// // 			// toss win
// // 			case true:
// // 				log.Println("not sending round 1 toss")
// // 				youWon := TossMsg + "YOU WON NOW YOU CAN SET THE DICTIONARY"

// // 				for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, token: _DictionaryDiscussion, to: _id, _sleep: false,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, to: _id, _sleep: false, token: Unblock,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, to: _id, _sleep: false, token: youWon,
// // 					}
// // 				}
// // 				tossWin := TossMsg + "WIN BY RED"

// // 				// block the toss loss team
// // 				for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 					if getPlayerInfo[_id][roomname][_TeamBlueKey].isBlock {
// // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 							roomname: roomname, to: _id, _sleep: false, token: Block,
// // 						}
// // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 							roomname: roomname, to: _id, _sleep: false, token: tossWin,
// // 						}
// // 					}
// // 				}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: TossAlert}

// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: tossSession}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: sendDictionary}

// // 			// toss loss
// // 			default:
// // 				log.Println("default toss")
// // 				log.Println("foundB: ", foundB)
// // 				youWon := TossMsg + "YOU WON NOW YOU CAN SET THE DICTIONARY"
// // 				for _, _id := range saveShuffle[roomname][_TeamRedKey] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, token: _DictionaryDiscussion, to: _id, _sleep: false,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, to: _id, _sleep: false, token: Unblock,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						roomname: roomname, to: _id, _sleep: false, token: youWon,
// // 					}

// // 				}
// // 				tossWin := TossMsg + "WIN BY RED"

// // 				// block the toss loss team
// // 				for _, _id := range saveShuffle[roomname][_TeamBlueKey] {
// // 					if getPlayerInfo[_id][roomname][_TeamBlueKey].isBlock {
// // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 							roomname: roomname, to: _id, _sleep: false, token: Block,
// // 						}
// // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 							roomname: roomname, to: _id, _sleep: false, token: tossWin,
// // 						}
// // 					}
// // 				}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: TossAlert}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: tossSession}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: sendDictionary}

// // 				log.Println("Winners Team blue: ", saveShuffle[roomname][_TeamBlueKey])
// // 			}
// // 			h.broadcast <- BroadcastReq{Token: _toss, RoomID: roomname}
// // 		} else {
// // 			for _, r := range saveShuffle[roomname][_TeamRedKey] {
// // 				log.Println("ids: ", r)

// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: Unblock,
// // 				}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: _toss,
// // 				}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: _tossCoin,
// // 				}
// // 			}

// // 			h.wg.Add(1)
// // 			go func() {
// // 				store := map[string]map[string]map[string]Session{}
// // 				store[id] = map[string]map[string]Session{
// // 					roomname: {_TeamBlueKey: se},
// // 				}
// // 				storeSessionUpdate <- store
// // 			}()
// // 			// h.wg.Add(1)
// // 			// go func() {
// // 			// 	defer h.wg.Done()
// // 			// 	store := map[string]map[string]bool{}
// // 			// 	store[roomname] = map[string]bool{_TeamBlueKey: true}
// // 			// 	Twatch.done <- store
// // 			// }()
// // 			// h.wg.Add(1)
// // 			// go func() {
// // 			// 	defer h.wg.Done()
// // 			// 	store := map[string]map[string]string{}
// // 			// 	store[roomname] = map[string]string{
// // 			// 		_TeamBlueKey: GR.HeadTails,
// // 			// 	}
// // 			// 	Ttoss.save <- store
// // 			// }()

// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store := map[string]GameInfo{}
// // 				ma := map[string]string{}
// // 				ma[_TeamBlueKey] = HeadTails
// // 				store[roomname] = GameInfo{SetTossBody: ma}
// // 				storeGameInfo <- store
// // 			}()

// // 			for _, r := range saveShuffle[roomname][_TeamBlueKey] {
// // 				log.Println("ids: ", r)

// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: Block,
// // 				}
// // 				// toss pick done
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					roomname: roomname, to: r, _sleep: false, token: toss,
// // 				}
// // 			}
// // 		}
// // 	}
// // 	// to set the face of the coin and to display the result after the coin toss
// // 	t := "Heads: " + strconv.FormatBool(ishead)
// // 	h.broadcast <- BroadcastReq{Token: t, RoomID: roomname}

// // }

// // const (
// // 	IDLE    = "I"
// // 	POINT   = "1"
// // 	NOPOINT = "0"
// // )

// // type FrameUpdate struct {
// // 	Update []string `json:"update"`
// // }

// // func UpdateFrame(h *Hub, RoomName string, List []string) {
// // 	var u = FrameUpdate{Update: []string{}}
// // 	u.Update = List
// // 	_token, _ := json.Marshal(&u)
// // 	h.broadcast <- BroadcastReq{Token: string(_token), RoomID: RoomName}
// // 	log.Println("sending update: ", _token)
// // }

// // func Update(h *Hub, isTeamBlue bool, RoomName string) {
// // 	boardUpdateR := "UpdateBoardR: " // plus score
// // 	boardUpdateB := "UpdateBoardB: " // plus score

// // 	nicknames := []string{}
// // 	createSheet := PackSheet{Name: make([]string, 4), Sheet: make(map[string]map[string]string)}

// // 	// creating sheet
// // 	for _, id := range saveShuffle[RoomName][_TeamBlueKey] {
// // 		nicknames = append(nicknames, getNicknamesViaID[id][RoomName])
// // 	}
// // 	for _, id := range saveShuffle[RoomName][_TeamRedKey] {
// // 		nicknames = append(nicknames, getNicknamesViaID[id][RoomName])
// // 	}

// // 	round := func(_for map[string]int, roomName string) string {
// // 		temp := _for[roomName]
// // 		x := "round" + strconv.Itoa(temp)
// // 		return x
// // 	}(getTtrackRound, RoomName)

// // 	createSheet.Name = append(createSheet.Name, nicknames...)

// // 	switch isTeamBlue {
// // 	case true:

// // 		resB := POINT   // point
// // 		resR := NOPOINT // no point

// // 		for _, id := range saveShuffle[RoomName][_TeamBlueKey] {
// // 			if getPlayerInfo[id][RoomName][_TeamBlueKey].isLock {
// // 				createSheet.Sheet[getNicknamesViaID[id][RoomName]] = map[string]string{
// // 					round: resB,
// // 				}
// // 			} else {
// // 				createSheet.Sheet[getNicknamesViaID[id][RoomName]] = map[string]string{
// // 					round: IDLE, // idle
// // 				}
// // 			}
// // 		}

// // 		for _, id := range saveShuffle[RoomName][_TeamRedKey] {
// // 			if getPlayerInfo[id][RoomName][_TeamRedKey].isLock {
// // 				createSheet.Sheet[getNicknamesViaID[id][RoomName]] = map[string]string{
// // 					round: resR,
// // 				}
// // 			} else {
// // 				createSheet.Sheet[getNicknamesViaID[id][RoomName]] = map[string]string{
// // 					round: IDLE, // idle
// // 				}
// // 			}
// // 		}

// // 		h.broadcast <- BroadcastReq{Token: boardUpdateR + "0", RoomID: RoomName}
// // 		h.broadcast <- BroadcastReq{Token: boardUpdateB + "1", RoomID: RoomName}

// // 	case false:

// // 		resB := NOPOINT // point
// // 		resR := POINT   // no point

// // 		for _, id := range saveShuffle[RoomName][_TeamBlueKey] {
// // 			if getPlayerInfo[id][RoomName][_TeamBlueKey].isLock {
// // 				createSheet.Sheet[getNicknamesViaID[id][RoomName]] = map[string]string{
// // 					round: resB,
// // 				}
// // 			} else {
// // 				createSheet.Sheet[getNicknamesViaID[id][RoomName]] = map[string]string{
// // 					round: IDLE, // idle
// // 				}
// // 			}
// // 		}

// // 		for _, id := range saveShuffle[RoomName][_TeamRedKey] {
// // 			if getPlayerInfo[id][RoomName][_TeamRedKey].isLock {
// // 				createSheet.Sheet[getNicknamesViaID[id][RoomName]] = map[string]string{
// // 					round: resR,
// // 				}
// // 			} else {
// // 				createSheet.Sheet[getNicknamesViaID[id][RoomName]] = map[string]string{
// // 					round: IDLE, // idle
// // 				}
// // 			}
// // 		}
// // 		h.broadcast <- BroadcastReq{Token: boardUpdateR + "1", RoomID: RoomName}
// // 		h.broadcast <- BroadcastReq{Token: boardUpdateB + "0", RoomID: RoomName}
// // 	}

// // 	pack, _ := json.Marshal(&createSheet)

// // 	h.broadcast <- BroadcastReq{Token: "UpdateCheatSheet: " + string(pack), RoomID: RoomName}

// // 	go func() {
// // 		store := map[string]PackSheet{}
// // 		store[RoomName] = createSheet

// // 		storeSheetUpdate.store <- store
// // 	}()
// // }

// // import (
// // 	"log"

// // 	"github.com/gofiber/contrib/websocket"
// // )

// // func Game1v1(h *Hub, conn *websocket.Conn, l Lobby, GR GameRoom) {

// // 	//h.wg.Wait()

// // 	// @TODO: make sure the challenge token is not same as the previous one
// // 	//

// // 	// pattern for tagging
// // 	count := 0

// // 	// @IMPORTANT
// // 	id := h.clients[conn].clientId
// // 	// hasKey := validate.Has_Key(GR.ChallengeToken)

// // 	// triggers only if the room session has started
// // 	if gameSessionStarted[GR.RoomName] {

// // 		// saving the token at beginning to avoid loss of time
// // 		if GR.PowerActivated {
// // 			switch true {
// // 			case GR.Nexus:
// // 				go func() {
// // 					store := map[string]map[string]map[string]bool{}
// // 					store[id] = map[string]map[string]bool{GR.RoomName: {___NexusKey: true}}
// // 					storePower.store <- store
// // 				}()
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: GR.RoomName}
// // 			case GR.Rewind:

// // 				go func() {
// // 					store := map[string]map[string]map[string]bool{}
// // 					store[id] = map[string]map[string]bool{GR.RoomName: {___RewindKey: true}}
// // 					storePower.store <- store
// // 				}()
// // 				store2 := map[string]map[string]bool{}
// // 				store2[GR.RoomName] = map[string]bool{___RewindPowerKey: true}
// // 				storeRewindPower <- store2

// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: GR.RoomName}

// // 			case GR.Freeze:
// // 				go func() {
// // 					store := map[string]map[string]map[string]bool{}
// // 					store[id] = map[string]map[string]bool{GR.RoomName: {___FreezeKey: true}}
// // 					storePower.store <- store
// // 				}()
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: GR.RoomName}

// // 			case GR.Tag:
// // 				go func() {
// // 					store := map[string]map[string]map[string]bool{}
// // 					store[id] = map[string]map[string]bool{GR.RoomName: {___TagKey: true}}
// // 					storePower.store <- store
// // 				}()
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: GR.RoomName}

// // 			case GR.Draw:
// // 				go func() {
// // 					store := map[string]map[string]map[string]bool{}
// // 					store[id] = map[string]map[string]bool{GR.RoomName: {___DrawKey: true}}
// // 					storePower.store <- store
// // 				}()
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: GR.RoomName}
// // 			case GR.Covert:
// // 				go func() {
// // 					store := map[string]map[string]map[string]bool{}
// // 					store[id] = map[string]map[string]bool{GR.RoomName: {___CovertKey: true}}
// // 					storePower.store <- store
// // 				}()
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: GR.RoomName}

// // 			}
// // 		}

// // 		switch true {

// // 		case GR.TossSession && !GR.Start:
// // 			Toss(h, id)

// // 		case !GR.Start && GR.DictionarySession:
// // 			ChallengeDictonarySession(h, id)

// // 		case !GR.Start && GR.ChallengeDiscussion:
// // 			ChallengeSession(h, id)

// // 		case GR.Start:
// // 			Score(h, count, id, conn)

// // 		default:
// // 			log.Println("session started but game not started yet")
// // 		}

// // 	} else {
// // 		log.Println("session not begin")
// // 	}

// // }

// // // GetFriendRoom returns if found returns the name of the room
// // func GetFriendRoom(code string) string {
// // 	Room := ""

// // 	for room := range getFriendRoomCode {
// // 		if code == getFriendRoomCode[room] {
// // 			Room = room
// // 		}
// // 	}
// // 	return Room
// // }

// // // FriendRoomExists returns true if the friend room exists via room name if specified else via matched code
// // func FriendRoomExists(roomName string, _code string, viaCode bool) bool {
// // 	found := false
// // 	switch true {
// // 	case viaCode:
// // 		for _, code := range getFriendRoomCode {
// // 			if code == _code {
// // 				found = true
// // 			}
// // 		}

// // 	case !viaCode:
// // 		if getFriendRoom[roomName][_FriendRoomKey] {
// // 			found = true
// // 		}
// // 	}
// // 	return found
// // }

// // import (
// // 	"encoding/json"
// // 	"log"
// // )

// // // what we want:
// // // voting for mutual ----
// // // case agree and case disagree
// // // if agree proceed to toss and store the current value of agree
// // // if disagree proceed to toss and store the current value of disagree
// // // -----
// // // toss ----
// // // case winner and agree || winner and disagree and case losser and agree || losser and disagree
// // // if winner and agree than proceed to dictionary session
// // // if winner and diagree than pro.....
// // // same for losser
// // // ---
// // // dictionary session---
// // // case agree and case disagree
// // // if agree than proceed to agree session
// // // if disagree than proceed to challenge session
// // // ---
// // // callenge session ----
// // // proceed to game
// // // ---
// // // agree session ---
// // // proceed to game
// // // ---
// // // game session

// // type BroadProfDisp struct {
// // 	Matchup map[string]string `json:"matchup"` // key= red or blue + their assocaited nicknames
// // }

// // var (
// // 	saveSession    = make(chan map[string]map[string]map[string]bool) // room-name: team-name: session-name:true |false
// // 	getsaveSession = map[string]map[string]map[string]bool{}          // room-name: team-name: session-name:true |false
// // )

// // var (
// // 	saveAgreement    = make(chan map[string]map[string]map[string]bool) // player-id: room-name: team-name:true|false
// // 	getSaveAgreement = make(map[string]map[string]map[string]bool)
// // 	saveChatMes      = make(chan map[string]map[string]map[string][]string) // id: room: team:....
// // 	getChatMes       = make(map[string]map[string]map[string][]string)
// // )

// // // ChallengeSessionTokens sends the token for game start
// // // if the last call made by the non-nexus
// // func ChallengeSessionTokens(h *Hub, id string, roomname string,
// // 	challengeToken string) {
// // 	log.Println("challenge session tokens")

// // 	token_ := BroadProfDisp{Matchup: make(map[string]string)}
// // 	nexus := false
// // 	nID, token := "", ""
// // 	teamname, opponentTeamname := TokenFromTeam(id, roomname)
// // 	isTeamName := false // if the current team is from teamname or not teamname
// // 	dict := getGameInfo[roomname].SetDictionary
// // 	// searching if the caller was the nexus user
// // 	for _, _id := range saveShuffle[roomname][teamname] {
// // 		if _id == id {
// // 			isTeamName = true
// // 		}
// // 	}

// // 	if isTeamName {
// // 		for _, _id := range saveShuffle[roomname][teamname] {
// // 			if getStoredPower[_id][roomname][_NexusKey] {
// // 				nexus = true
// // 				nID = _id
// // 				break
// // 			}
// // 		}
// // 		if nexus {
// // 			for _, _id := range saveShuffle[roomname][opponentTeamname] {
// // 				if getPlayerInfo[_id][roomname][opponentTeamname].isLock {
// // 					token = getTokens[_id][roomname][teamname].Challenge
// // 					break
// // 				}
// // 			}
// // 			PNexusMechanism(h, roomname, nID, dict, token, 1)

// // 		} else {
// // 			// searching if it was not from the caller
// // 			for _, _id := range saveShuffle[roomname][opponentTeamname] {
// // 				if getStoredPower[_id][roomname][_NexusKey] {
// // 					nexus = true
// // 					nID = _id
// // 					break
// // 				}
// // 			}
// // 			if nexus {
// // 				token = challengeToken
// // 				PNexusMechanism(h, roomname, nID, dict, token, 1)
// // 			}

// // 		}
// // 	} else {
// // 		for _, _id := range saveShuffle[roomname][opponentTeamname] {
// // 			if getStoredPower[_id][roomname][_NexusKey] {
// // 				nexus = true
// // 				nID = _id
// // 			}
// // 		}
// // 		if nexus {
// // 			for _, _id := range saveShuffle[roomname][teamname] {
// // 				if getPlayerInfo[_id][roomname][teamname].isLock {
// // 					token = getTokens[_id][roomname][teamname].Challenge
// // 					break
// // 				}
// // 			}
// // 			PNexusMechanism(h, roomname, nID, dict, token, 1)

// // 		} else {
// // 			// searching if it was not from the caller
// // 			for _, _id := range saveShuffle[roomname][teamname] {
// // 				if getStoredPower[_id][roomname][_NexusKey] {
// // 					nexus = true
// // 					nID = _id
// // 					break
// // 				}
// // 			}
// // 			if nexus {
// // 				token = challengeToken
// // 				PNexusMechanism(h, roomname, nID, dict, token, 1)
// // 			}
// // 		}
// // 	}
// // 	isBetSessionUsedByTeamName, isBetSession := false, false
// // 	log.Println("checking for session update")
// // 	if isTeamName {

// // 		for _, _id := range saveShuffle[roomname][opponentTeamname] {
// // 			if getSessionUpdate[_id] != nil && getSessionUpdate[_id][roomname][opponentTeamname].BetDone {
// // 				log.Println("bet session found ")
// // 				isBetSession = true
// // 			}
// // 		}

// // 		store := map[string]map[string]map[string]Tokens{}
// // 		store[id] = map[string]map[string]Tokens{
// // 			roomname: {teamname: Tokens{Challenge: challengeToken}},
// // 		}
// // 		storeTokensUpdate <- store

// // 	} else {

// // 		for _, _id := range saveShuffle[roomname][teamname] {
// // 			if getSessionUpdate[_id] != nil && getSessionUpdate[_id][roomname][teamname].BetDone {
// // 				log.Println("bet session found")
// // 				isBetSession = true
// // 			}
// // 		}

// // 		store := map[string]map[string]map[string]Tokens{}
// // 		store[id] = map[string]map[string]Tokens{
// // 			roomname: {opponentTeamname: Tokens{Challenge: challengeToken}},
// // 		}

// // 		storeTokensUpdate <- store
// // 	}

// // 	log.Println("checking for lock")
// // 	for _, _id := range saveShuffle[roomname][teamname] {
// // 		if getPlayerInfo[_id][roomname][teamname].isLock {

// // 			token_.Matchup[teamname] = getNicknamesViaID[_id][roomname]
// // 		}
// // 	}
// // 	for _, _id := range saveShuffle[roomname][opponentTeamname] {
// // 		if getPlayerInfo[_id][roomname][teamname].isLock {

// // 			token_.Matchup[opponentTeamname] = getNicknamesViaID[_id][roomname]
// // 		}
// // 	}
// // 	_parcel2, _ := json.Marshal(&token_)
// // 	h.broadcast <- BroadcastReq{Token: "TeamAgainstx:" + string(_parcel2), RoomID: roomname}

// // 	if isBetSession {
// // 		// do not save anything it is begin saved in creatbetCups
// // 		token := BetPattern{}
// // 		k := getStoredCups[roomname]
// // 		token.FirstCup = k[0]
// // 		token.SecondCup = k[1]
// // 		token.ThirdCup = k[2]
// // 		log.Println("Stored cups: ", getStoredCups[roomname])
// // 		parcel, _ := json.Marshal(&token)
// // 		_token := ThreeBetCups + string(parcel)
// // 		log.Println("sending parcel: ", _token)

// // 		// h.broadcast <- BroadcastReq{Token: ThreeBetCups + string(parcel), RoomID: roomname}

// // 		if isBetSessionUsedByTeamName {
// // 			for _, r := range saveShuffle[roomname][opponentTeamname] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Betting, to: r, roomname: roomname,
// // 					_sleep: false}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _token, to: r, roomname: roomname,
// // 					_sleep: false}
// // 			}
// // 		} else {
// // 			for _, r := range saveShuffle[roomname][teamname] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Betting, to: r, roomname: roomname,
// // 					_sleep: false}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _token, to: r, roomname: roomname,
// // 					_sleep: false}
// // 			}
// // 		}
// // 	}

// // 	SendPowers(h, roomname, teamname, opponentTeamname)

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()

// // 		for _, _id := range saveShuffle[roomname][teamname] {
// // 			h.broadcast <- BroadcastReq{Token: "Display: " + getGameInfo[roomname].NicknameViaID[_id] + " " + "Team Red"}
// // 		}
// // 		for _, _id := range saveShuffle[roomname][teamname] {
// // 			h.broadcast <- BroadcastReq{Token: "Display: " + getGameInfo[roomname].NicknameViaID[_id] + " " + "Team Blue"}
// // 		}

// // 		log.Println("game info: ", getGameInfo)
// // 	}()

// // 	h.wg.Go(func() {
// // 		token := "ChallengeGuess: "
// // 		token2 := "ChallengeGuess: " + challengeToken

// // 		if isTeamName {
// // 			for _, _id := range saveShuffle[roomname][opponentTeamname] {
// // 				if _id != id {
// // 					token = "ChallengeGuess: " + getTokens[_id][roomname][opponentTeamname].Challenge
// // 					break
// // 				}
// // 			}
// // 		} else {
// // 			for _, _id := range saveShuffle[roomname][teamname] {
// // 				if _id != id {
// // 					token = "ChallengeGuess: " + getTokens[_id][roomname][teamname].Challenge
// // 					break
// // 				}
// // 			}
// // 		}

// // 		for _, ids := range saveShuffle[roomname][teamname] {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: ids, roomname: roomname, _sleep: false}
// // 		}

// // 		for _, ids := range saveShuffle[roomname][opponentTeamname] {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: ids, roomname: roomname, _sleep: false}
// // 		}

// // 		log.Println("room name: ", roomname)
// // 		log.Println(token)
// // 		log.Println(token2)

// // 		// saveSessionDone.done <- true
// // 	})
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		h.broadcast <- BroadcastReq{Token: challengeDiscussion, RoomID: roomname}
// // 		h.broadcast <- BroadcastReq{Token: _startGame, RoomID: roomname}
// // 		h.broadcast <- BroadcastReq{Token: Unblock, RoomID: roomname}
// // 	}()
// // 	TrashChallenge(h, roomname, true, teamname, opponentTeamname)

// // }

// // func Abundance(h *Hub, l Lobby) {
// // 	log.Println("changing the room settings")
// // 	go func() {
// // 		store := map[string]map[string]bool{}
// // 		store[l.RoomName] = map[string]bool{
// // 			_SetupTossKey: l.SetToss,
// // 			_StarterKey:   l.Starter,
// // 			_ReverseKey:   l.Reverse,

// // 			_DrawPowerKey:   l.DrawPower,
// // 			_NexusPowerKey:  l.NexusPower,
// // 			_TagPowerKey:    l.TagPower,
// // 			_RewindPowerKey: l.RewindPower,
// // 			_FreezePowerKey: l.FreezePower,
// // 			_CovertPowerKey: l.CovertPower,
// // 			_BetPowerKey:    l.BetPower,
// // 		}
// // 		settings.store <- store
// // 	}()

// // 	go func() {
// // 		store := map[string]map[string]string{}
// // 		store[l.RoomName] = map[string]string{
// // 			_FieldKey:        l.Field,
// // 			_CategoryKey:     l.Category,
// // 			_BookKey:         l.Book,
// // 			_DecisionTimeKey: l.DecisionTime,
// // 			_GameTimeKey:     l.GameTime,
// // 		}
// // 		storeRoomRequestedToken.store <- store
// // 	}()

// // }

// // func GenerateURLRequest(RoomName string, key string) string {
// // 	URL := getRoomRequestedToken[RoomName][key]
// // 	return URL // save it like map[room name][url]
// // }

// // RoomValidationTokensForJoined broadcasts the tokens after the room has been created or joined
// // func RoomValidationTokensForJoined(roomname string, h *Hub, isPrivate bool, is2v2 bool, isCreated bool, l Lobby) {
// // 	roomMode, token, sendPowers, IsEntertainment := "1v1", "", map[string]map[string]bool{}, false
// // 	fmt.Println("in tokens")
// // 	fmt.Println("room name: ", roomname)
// // 	if is2v2 {
// // 		roomMode = "2v2"
// // 	}

// // 	if l.Category == "entertainment" {
// // 		IsEntertainment = true
// // 	}
// // 	if IsEntertainment {
// // 		sendBook := books.PackEntertainment()
// // 		conv, err := json.Marshal(sendBook)
// // 		if err != nil {
// // 			panic("marshalling book error")
// // 		}
// // 		token = string(conv)

// // 	} else {
// // 		sendBook := books.PackSports()
// // 		conv, err := json.Marshal(sendBook)
// // 		if err != nil {
// // 			panic("marshalling book error")
// // 		}
// // 		token = string(conv)

// // 	}

// // 	fmt.Println("book token: ", token)

// // 	if isCreated {
// // 		sendPowers[l.RoomName] = map[string]bool{
// // 			_DrawPowerKey:   l.DrawPower,
// // 			_NexusPowerKey:  l.NexusPower,
// // 			_TagPowerKey:    l.TagPower,
// // 			_RewindPowerKey: l.RewindPower,
// // 			_FreezePowerKey: l.FreezePower,
// // 			_CovertPowerKey: l.CovertPower,
// // 			_BetPowerKey:    l.BetPower,
// // 		}

// // 		for power, included := range sendPowers[l.RoomName] {
// // 			if included {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ParcelRoomSettings:" + power + ": true", roomname: roomname, to: l.ID, _sleep: false}
// // 			}
// // 		}
// // 	} else {
// // 		sendP := ChoosenPowers(roomname)
// // 		for _, power := range sendP {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ParcelRoomSettings:" + power + ": true", roomname: roomname, to: l.ID, _sleep: false}
// // 		}
// // 	}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomDecisionTime: " + l.DecisionTime, roomname: roomname, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomGameTime: " + l.GameTime, roomname: roomname, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCategory: " + l.Category, roomname: roomname, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomBook: " + l.Book, roomname: roomname, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: roomname, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "IsEnterTainment: " + strconv.FormatBool(IsEntertainment), roomname: roomname, to: l.ID, _sleep: false}
// // 	// this is for the room-settings
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "FetchedBook: " + token, roomname: roomname, to: l.ID, _sleep: false}

// // }

// // // import (
// // // 	"fmt"
// // // )

// // // // players deciding to agree on setting the challenge mutualy or solelly
// // // // this is the first session that triggers in the angular session whihc basically decides
// // // // if player wish to do agree session or challenge session
// // // func MutualSession(h *Hub, id string) {

// // // 	isTeamRed := false
// // // 	isTeamBlue := false
// // // 	proceed := false

// // // 	for _, r := range saveShuffle[___TeamRedKey] {
// // // 		if r == id {
// // // 			isTeamRed = true
// // // 		}
// // // 	}
// // // 	for _, r := range saveShuffle[___TeamBlueKey] {
// // // 		if r == id {
// // // 			isTeamBlue = true
// // // 		}
// // // 	}

// // // 	// save the sesion to true if the player visited this session
// // // 	if isTeamRed {
// // // 		store := map[string]bool{}
// // // 		store[___TeamRedKey] = true
// // // 		Mwatch.going <- store
// // // 	} else {
// // // 		store := map[string]bool{}
// // // 		store[___TeamBlueKey] = true
// // // 		Mwatch.going <- store
// // // 	}

// // // 	go func() {
// // // 		// send the token back if and only if the player is agreed
// // // 		switch true {
// // // 		// proceed to the challenge dictionary session
// // // 		case GR.MutalCount == 0:

// // // 			switch true {

// // // 			case isTeamRed:

// // // 				// send back that the session is over
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // // 					roomName: l.To, token: mutalSession, to: saveShuffle[___TeamRedKey][0], _sleep: false,
// // // 				}
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // // 					roomName: l.To, token: mutalSession, to: saveShuffle[___TeamRedKey][1], _sleep: false,
// // // 				}

// // // 				// if team b already visited but this has a limitation tho for agree session count
// // // 				if getMWatch[___TeamBlueKey] {
// // // 					proceed = true
// // // 				}

// // // 				if proceed {
// // // 					fmt.Println("proceed")
// // // 					if getSettings[l.To][___SetupToss] {
// // // 						h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // // 					} else {
// // // 						if GR.Round == 1 { // note after round one the toss session has been stoped
// // // 							fmt.Println("sending round 1 toss")
// // // 							h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // // 						} else {
// // // 							fmt.Println("not sending round 1 toss")
// // // 							h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}
// // // 						}
// // // 					}
// // // 					Mwatch.done <- true
// // // 				}
// // // 				// this is the signal for the dictionary to pass token to respective team member to that session
// // // 				go func() {
// // // 					store2 := map[string]map[string]map[string]bool{}
// // // 					store2[___TeamRedKey] = map[string]map[string]bool{l.To: {___ChallengeSession: true}}
// // // 					saveSession <- store2
// // // 				}()
// // // 			case isTeamBlue:

// // // 				for _, _id := range saveShuffle[___TeamBlueKey] {
// // // 					// send back that the session is over
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // // 						roomName: l.To, token: mutalSession, to: _id, _sleep: false,
// // // 					}
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // // 						roomName: l.To, token: mutalSession, to: _id, _sleep: false,
// // // 					}
// // // 				}

// // // 				// if the previous team player visited
// // // 				if getMWatch[___TeamRedKey] {
// // // 					proceed = true
// // // 				}

// // // 				if proceed {
// // // 					if getSettings[l.To][___SetupToss] {
// // // 						h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // // 					} else {
// // // 						if GR.Round == 1 {
// // // 							h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // // 						} else {
// // // 							h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}
// // // 						}
// // // 					}
// // // 					Mwatch.done <- true
// // // 				}

// // // 				go func() {
// // // 					store2 := map[string]map[string]map[string]bool{}
// // // 					store2[___TeamBlueKey] = map[string]map[string]bool{l.To: {___ChallengeSession: true}}

// // // 					saveSession <- store2
// // // 				}()
// // // 			default:
// // // 				fmt.Println("none case mutual")
// // // 			}

// // // 		case GR.MutalCount == 1:

// // // 			switch true {
// // // 			case isTeamRed:
// // // 				if getMWatch[___TeamBlueKey] {
// // // 					proceed = true
// // // 				}

// // // 				if proceed {
// // // 					fmt.Println("proceed")
// // // 					if getSettings[l.To][___SetupToss] {

// // // 						h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // // 					} else {
// // // 						if GR.Round == 1 {
// // // 							fmt.Println("sending round 1 toss")

// // // 							h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // // 						} else {
// // // 							fmt.Println("not sending round 1 toss")
// // // 							h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}
// // // 						}
// // // 					}
// // // 					Mwatch.done <- true // important: clear cache
// // // 				}

// // // 				// store the session for which player has choosen
// // // 				go func() {
// // // 					store2 := map[string]map[string]map[string]bool{}
// // // 					store2[___TeamRedKey] = map[string]map[string]bool{l.To: {___AgressSession: true}}
// // // 					saveSession <- store2
// // // 				}()

// // // 			case isTeamBlue:
// // // 				// if the previous team player visited
// // // 				if getMWatch[___TeamRedKey] {
// // // 					proceed = true
// // // 				}

// // // 				if proceed {
// // // 					fmt.Println("proceed")
// // // 					if getSettings[l.To][___SetupToss] {

// // // 						h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // // 					} else {
// // // 						if GR.Round == 1 {
// // // 							h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // // 						} else {
// // // 							h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}
// // // 						}
// // // 					}
// // // 					Mwatch.done <- true
// // // 				}
// // // 				go func() {
// // // 					store2 := map[string]map[string]map[string]bool{}
// // // 					store2[___TeamBlueKey] = map[string]map[string]bool{l.To: {___AgressSession: true}}

// // // 					saveSession <- store2
// // // 					// Mwatch.done <- true
// // // 				}()
// // // 			}

// // // 		default:
// // // 			fmt.Println("pattern not found: ", GR.MutalCount)
// // // 		}
// // // 	}()
// // // }

// // // // what we need from this is basically whosoever is not block will first set the challenge that's it
// // // // proceed to challenge session and broadcast the dictionary token
// // // // in order to make this work you have to use mutual session
// // // // and in angular you have to add the token for mutual session
// // // func ChallengeDictonarySession(h *Hub, id string) {

// // // 	isTeamRed := false
// // // 	isTeamBlue := false
// // // 	for _, r := range saveShuffle[___TeamRedKey] {
// // // 		if r == id {
// // // 			isTeamRed = true
// // // 		}
// // // 	}

// // // 	for _, r := range saveShuffle[___TeamBlueKey] {
// // // 		if r == id {
// // // 			isTeamBlue = true
// // // 		}
// // // 	}

// // // 	// if the first team to set the dictionary
// // // 	// than second team can able to set the challenge
// // // 	// or
// // // 	// if the first team to set the dictornary
// // // 	// than first team can able to set the challenge
// // // 	// please note: reverse was here for exerimental the testing was only done for else case
// // // 	if getSettings[l.To][___ReverseKey] {
// // // 		// a1-b1 pattern: block the one that arrived and unblock the one that unable to make it and proceed him to set challenge
// // // 		switch true {

// // // 		case isTeamBlue:
// // // 			for _, __id := range saveShuffle[___TeamBlueKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomName: l.To, _sleep: false, to: __id}
// // // 			}
// // // 			for _, __id := range saveShuffle[___TeamRedKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomName: l.To, _sleep: false, to: __id}
// // // 			}

// // // 			h.broadcast <- BroadcastReq{Token: _challengeDiscussion, RoomId: l.To}
// // // 			h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomId: l.To}

// // // 		case isTeamRed:
// // // 			for _, __id := range saveShuffle[___TeamBlueKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomName: l.To, _sleep: false, to: __id}
// // // 			}
// // // 			for _, __id := range saveShuffle[___TeamRedKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomName: l.To, _sleep: false, to: __id}
// // // 			}
// // // 			h.broadcast <- BroadcastReq{Token: _challengeDiscussion, RoomId: l.To}
// // // 			h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomId: l.To}

// // // 		default:
// // // 			fmt.Println("dictornay not able to decide")
// // // 		}

// // // 	} else {
// // // 		// a1-a1 pattern: block the one that arriveed and block the one that unable to make it and proceed the one who has came here to set challenge
// // // 		switch true {

// // // 		case isTeamBlue:
// // // 			h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomId: l.To}

// // // 			for _, __id := range saveShuffle[___TeamBlueKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomName: l.To, _sleep: false, to: __id}
// // // 			}
// // // 			for _, __id := range saveShuffle[___TeamRedKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomName: l.To, _sleep: false, to: __id}
// // // 			}

// // // 			for _, _id := range saveShuffle[___TeamBlueKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "DictionaryEvent: " + GR.DictionaryToken, roomName: l.To, _sleep: false, to: _id}
// // // 			}

// // // 			// if the first session requested is for the challenge or agree
// // // 			// it's important to pass the to both the player becuase
// // // 			// only of the player will be setting the dictionary

// // // 			if getsaveSession[___TeamBlueKey][l.To][___ChallengeSession] {
// // // 				for _, __id := range saveShuffle[___TeamBlueKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _challengeDiscussion, roomName: l.To, _sleep: false, to: __id}
// // // 				}

// // // 			} else if getsaveSession[___TeamBlueKey][l.To][___AgressSession] {
// // // 				for _, __id := range saveShuffle[___TeamBlueKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _votingSession, roomName: l.To, _sleep: false, to: __id}
// // // 				}
// // // 			}

// // // 			if getsaveSession[___TeamRedKey][l.To][___ChallengeSession] {
// // // 				for _, __id := range saveShuffle[___TeamRedKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _challengeDiscussion, roomName: l.To, _sleep: false, to: __id}
// // // 				}

// // // 			} else if getsaveSession[___TeamRedKey][l.To][___AgressSession] {
// // // 				for _, __id := range saveShuffle[___TeamRedKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _votingSession, roomName: l.To, _sleep: false, to: __id}
// // // 				}
// // // 			}

// // // 		case isTeamRed:
// // // 			h.broadcast <- BroadcastReq{Token: DictionaryDiscussion, RoomId: l.To}

// // // 			for _, __id := range saveShuffle[___TeamBlueKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, roomName: l.To, _sleep: false, to: __id}
// // // 			}
// // // 			for _, __id := range saveShuffle[___TeamRedKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, roomName: l.To, _sleep: false, to: __id}
// // // 			}

// // // 			for _, _id := range saveShuffle[___TeamRedKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "DictionaryEvent: " + GR.DictionaryToken, roomName: l.To, _sleep: false, to: _id}
// // // 			}

// // // 			if getsaveSession[___TeamBlueKey][l.To][___ChallengeSession] {
// // // 				for _, __id := range saveShuffle[___TeamBlueKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _challengeDiscussion, roomName: l.To, _sleep: false, to: __id}
// // // 				}

// // // 			} else if getsaveSession[___TeamBlueKey][l.To][___AgressSession] {
// // // 				for _, __id := range saveShuffle[___TeamBlueKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _votingSession, roomName: l.To, _sleep: false, to: __id}
// // // 				}
// // // 			}

// // // 			if getsaveSession[___TeamRedKey][l.To][___ChallengeSession] {
// // // 				for _, __id := range saveShuffle[___TeamRedKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _challengeDiscussion, roomName: l.To, _sleep: false, to: __id}
// // // 				}

// // // 			} else if getsaveSession[___TeamRedKey][l.To][___AgressSession] {
// // // 				for _, __id := range saveShuffle[___TeamRedKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _votingSession, roomName: l.To, _sleep: false, to: __id}
// // // 				}
// // // 			}

// // // 		default:
// // // 			fmt.Println("dictornay not able to decide")
// // // 		}
// // // 	}
// // // }

// // // // i= id, r=room-name a=any [can be session name, power name or any other name]
// // // type PatternIRA struct {
// // // 	store chan map[string]map[string]map[string]bool // player-id: room-name: power-name: true | false
// // // 	mu    sync.Mutex
// // // 	done  chan bool
// // // }

// // // var (
// // // 	storePower     = PatternIRA{store: make(chan map[string]map[string]map[string]bool), done: make(chan bool)} // player-id: room-name: power-name: true | false
// // // 	getStoredPower = make(map[string]map[string]map[string]bool)
// // // )

// // // func Game2v2(h *Hub, conn *websocket.Conn, l Lobby, GR GameRoom) {
// // // 	//h.wg.Wait()

// // // 	// @TODO: make sure the challenge token is not same as the previous one
// // // 	//

// // // 	// pattern for tagging
// // // 	count := 0

// // // 	// @IMPORTANT
// // // 	id := h.clients[conn].clientId
// // // 	// hasKey := validate.Has_Key(GR.ChallengeToken)

// // // 	if gameSessionStarted[l.To] {

// // // 		go func() {
// // // 			audit.mu.Lock()
// // // 			defer audit.mu.Unlock()
// // // 			pass := []string{}
// // // 			if GR.VotingSession {
// // // 				pass = append(pass, GR.MutualVote)
// // // 			} else {
// // // 				pass = append(pass, GR.ChallengeToken)
// // // 			}
// // // 			audit.token <- pass
// // // 		}()

// // // 		// store the power here rather than when game starts
// // // 		// the reason being the loss of time
// // // 		if GR.PowerActivated {
// // // 			switch true {
// // // 			case GR.Nexus:
// // // 				go func() {
// // // 					store := map[string]map[string]map[string]bool{}
// // // 					store[id] = map[string]map[string]bool{l.To: {___Nexus: true}}
// // // 					storePower.store <- store
// // // 				}()
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}
// // // 			case GR.Rewind:

// // // 				go func() {
// // // 					store := map[string]map[string]map[string]bool{}
// // // 					store[id] = map[string]map[string]bool{l.To: {___Rewind: true}}
// // // 					storePower.store <- store
// // // 				}()
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 			case GR.Freeze:
// // // 				go func() {
// // // 					store := map[string]map[string]map[string]bool{}
// // // 					store[id] = map[string]map[string]bool{l.To: {___Freeze: true}}
// // // 					storePower.store <- store
// // // 				}()
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 			}
// // // 		}
// // // /**
// // // * the session here is placed as per the position and will get trigger as per the case
// // // * meaning first is mutual than player goes to toss session than dictionary session than challenge or voting session and than game begins
// // // */
// // // 		switch true {
// // // 		// checking for the players are mutually deciding the challegne token
// // // 		case GR.MutualSession && !GR.Start:
// // // 			MutualSession(h, id)

// // // 		case GR.TossSession && !GR.Start:
// // // 			Toss(h, id)

// // // 		case !GR.Start && GR.DictionarySession:
// // // 			ChallengeDictonarySession(h, id)

// // // 		case !GR.Start && GR.ChallengeDiscussion || GR.VotingSession && !GR.MutualSession:
// // // 			Abundance(h, id)

// // // 		case GR.Start:
// // // 			Score(h, count, id, conn)

// // // 		default:
// // // 			fmt.Println("session started but game not started yet")
// // // 		}

// // // 	} else {
// // // 		fmt.Println("session not begin")
// // // 	}
// // // }

// // // players agrees on setting the challenge mutually
// // // what we need?!!!
// // // keep chatting till they dont agree on common count
// // // once the chances are 0 then proceed
// // // once agreed proceed
// // // func AgreeSession(h *Hub, id string, count int, newCount string) {
// // // 	isTeamRed := false
// // // 	isTeamBlue := false
// // // 	chatToken := "ChatToken: " + GR.MutualVote
// // // 	proceed := false

// // // 	fmt.Println("Agree session")
// // // 	for _, r := range saveShuffle[___TeamRedKey] {
// // // 		if r == id {
// // // 			isTeamRed = true
// // // 		}
// // // 	}

// // // 	for _, r := range saveShuffle[___TeamBlueKey] {
// // // 		if r == id {
// // // 			isTeamBlue = true
// // // 		}
// // // 	}

// // // 	// store the visited team
// // // 	if isTeamRed {
// // // 		store := map[string]bool{}
// // // 		store[___TeamRedKey] = true
// // // 		Cwatch.going <- store
// // // 	} else {
// // // 		store := map[string]bool{}
// // // 		store[___TeamBlueKey] = true
// // // 		Cwatch.going <- store
// // // 	}

// // // 	fmt.Println("agree token: ", GR.MutualVote)

// // // 	// save the key that matches agree
// // // 	if isTeamBlue {
// // // 		if GR.MutualVote == "agree" {
// // // 			go func() {
// // // 				store := map[string]map[string]map[string]bool{}
// // // 				store[id] = map[string]map[string]bool{l.To: {___TeamBlueKey: true}}
// // // 				saveAgreement <- store
// // // 			}()
// // // 		} else {
// // // 			go func() {
// // // 				store := map[string]map[string]map[string]bool{}
// // // 				store[id] = map[string]map[string]bool{l.To: {___TeamBlueKey: false}}
// // // 				saveAgreement <- store
// // // 			}()
// // // 		}
// // // 	} else if isTeamRed {
// // // 		if GR.MutualVote == "agree" {
// // // 			go func() {
// // // 				store := map[string]map[string]map[string]bool{}
// // // 				store[id] = map[string]map[string]bool{l.To: {___TeamRedKey: true}}
// // // 				saveAgreement <- store
// // // 			}()
// // // 		} else {
// // // 			go func() {
// // // 				store := map[string]map[string]map[string]bool{}
// // // 				store[id] = map[string]map[string]bool{l.To: {___TeamRedKey: false}}
// // // 				saveAgreement <- store
// // // 			}()
// // // 		}
// // // 	}

// // // 	// check if both the player agrees on a item to set for the challenge
// // // 	// the problem here is that player has to send the agree twice becuase of gorutine
// // // 	// in order to get real time update make sure to also check with the current value
// // // 	// so that if the previous key was true and this was true we can send the real-time update
// // // 	if isTeamBlue {
// // // 		if getSaveAgreement[saveShuffle[___TeamBlueKey][0]][l.To][___TeamBlueKey] && getSaveAgreement[saveShuffle[___TeamBlueKey][1]][l.To][___TeamBlueKey] {
// // // 			proceed = true
// // // 		}
// // // 	} else if isTeamRed {
// // // 		if getSaveAgreement[saveShuffle[___TeamRedKey][0]][l.To][___TeamRedKey] && getSaveAgreement[saveShuffle[___TeamRedKey][1]][l.To][___TeamRedKey] {
// // // 			proceed = true
// // // 		}
// // // 	}

// // // 	// if the player agrees or disagrees
// // // 	switch true {
// // // 	// agrees
// // // 	case proceed:
// // // 		switch true {
// // // 		case isTeamBlue:

// // // 			// if the previous player's session has been done
// // // 			if isSessionDone[saveShuffle[___TeamRedKey][0]][l.To][___TeamRedKey][___AgressSession] || isSessionDone[saveShuffle[___TeamRedKey][1]][l.To][___TeamRedKey][___AgressSession] {
// // // 				proceed = true
// // // 			} else if isSessionDone[saveShuffle[___TeamRedKey][0]][l.To][___TeamRedKey][___ChallengeSession] || isSessionDone[saveShuffle[___TeamRedKey][1]][l.To][___TeamRedKey][___ChallengeSession] {
// // // 				proceed = true
// // // 			} else {
// // // 				proceed = false
// // // 			}

// // // 			// proceed if and only if the previous team player has completed his session
// // // 			if proceed {

// // // 				h.broadcast <- BroadcastReq{Token: votingSession, RoomId: l.To}
// // // 				h.broadcast <- BroadcastReq{Token: challengeDiscussion, RoomId: l.To}
// // // 				h.broadcast <- BroadcastReq{Token: _startGame, RoomId: l.To}
// // // 				h.broadcast <- BroadcastReq{Token: Unblock, RoomId: l.To}
// // // 				saveSessionDone.done <- true

// // // 				// important to activate the power up now becuase the game starts
// // // 				switch true {
// // // 				case getStoredPower[id][l.To][___Nexus]:
// // // 					token := "Nexus: " + Nexus("VIRAT", "_")
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Nexus: " + token,
// // // 						to: id, roomName: l.To, _sleep: false}
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 				case getStoredPower[id][l.To][___Freeze]:
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: FreezePower,
// // // 						to: id, roomName: l.To, _sleep: false}
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 				case getStoredPower[id][l.To][___Rewind]:
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RewindPower,
// // // 						to: id, roomName: l.To, _sleep: false}
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}
// // // 				}
// // // 				h.wg.Add(1)
// // // 				go func() {
// // // 					defer h.wg.Done()
// // // 					newStore := map[string]string{}
// // // 					reverseNewStore := map[string]string{}
// // // 					for _, tag := range saveShuffle {
// // // 						for _, data := range tag {
// // // 							for Name, ID := range getProfData {
// // // 								if ID == data {
// // // 									newStore[Name] = data
// // // 									reverseNewStore[data] = Name
// // // 								}
// // // 							}
// // // 						}
// // // 					}
// // // 					// becuase we have saved the data of joiners separately
// // // 					for _, id := range reverseNewStore {
// // // 						for Ownername, Ownerid := range getOwnerProf {
// // // 							if Ownerid != id {
// // // 								reverseNewStore[Ownerid] = Ownername
// // // 								newStore[Ownername] = Ownerid // its obvious that if the id doenst match it means that it is not in the new store too
// // // 							}
// // // 						}
// // // 					}

// // // 					token1 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][0]]}
// // // 					token2 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][1]]}

// // // 					token3 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][0]]}
// // // 					token4 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][1]]}

// // // 					__RED := "Display: " + token1.NickName + " " + token1.TeamName
// // // 					__RED2 := "Display: " + token2.NickName + " " + token2.TeamName
// // // 					__BLUE := "Display: " + token3.NickName + " " + token3.TeamName
// // // 					__BLUE2 := "Display: " + token4.NickName + " " + token4.TeamName

// // // 					h.broadcast <- BroadcastReq{Token: __RED, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: __RED2, RoomId: l.To}

// // // 					h.broadcast <- BroadcastReq{Token: __BLUE, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: __BLUE2, RoomId: l.To}

// // // 					fmt.Println("new Store: ", newStore)
// // // 					fmt.Println("reverseNewStore: ", reverseNewStore)
// // // 					fmt.Println("getOwnerProf: ", getOwnerProf)

// // // 				}()

// // // 				h.wg.Add(1)
// // // 				go func() {
// // // 					defer h.wg.Done()

// // // 					token := "ChallengeGuess: " + __getChallengeSet[___TeamRedKey]
// // // 					token2 := "ChallengeGuess: " + GR.MutualVote

// // // 					for _, ids := range saveShuffle[___TeamRedKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: ids, roomName: l.To, _sleep: false}
// // // 					}

// // // 					for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: ids, roomName: l.To, _sleep: false}
// // // 					}

// // // 					fmt.Println("challenge set: ", __getChallengeSet)

// // // 					_trackChallengeSet.done <- true
// // // 				}()
// // // 			} else {
// // // 				// waiting for other team to complete
// // // 				fmt.Println("waiting blue team")
// // // 				for _, ids := range saveShuffle[___TeamRedKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, to: ids, roomName: l.To, _sleep: false}
// // // 				}

// // // 				for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, to: ids, roomName: l.To, _sleep: false}
// // // 				}

// // // 				// storing the session done
// // // 				go func() {
// // // 					store := map[string]map[string]map[string]map[string]bool{}
// // // 					store[id] = map[string]map[string]map[string]bool{l.To: {___TeamBlueKey: {___AgressSession: true}}}
// // // 					saveSessionDone.store <- store
// // // 				}()

// // // 				// savnig the messages
// // // 				// problem: it's hard to distinguish the message of team a and team B
// // // 				// you can use two different queueu to store too
// // // 				msg := ""
// // // 				go func() {
// // // 					for _, msgs := range getMsgs.Range() {
// // // 						if msgs != "agree" {
// // // 							msg = msgs
// // // 						}
// // // 					}
// // // 					store := map[string]string{}
// // // 					store[___TeamBlueKey] = msg
// // // 					_trackChallengeSet.save <- store

// // // 				}()
// // // 			}

// // // 		case isTeamRed:
// // // 			fmt.Println("Red team proceeed")
// // // 			if isSessionDone[saveShuffle[___TeamBlueKey][0]][l.To][___TeamBlueKey][___AgressSession] || isSessionDone[saveShuffle[___TeamBlueKey][1]][l.To][___TeamBlueKey][___AgressSession] {
// // // 				proceed = true
// // // 			} else if isSessionDone[saveShuffle[___TeamBlueKey][0]][l.To][___TeamBlueKey][___ChallengeSession] || isSessionDone[saveShuffle[___TeamBlueKey][1]][l.To][___TeamBlueKey][___ChallengeSession] {
// // // 				proceed = true
// // // 			} else {
// // // 				proceed = false
// // // 			}

// // // 			if proceed {
// // // 				saveSessionDone.done <- true
// // // 				h.broadcast <- BroadcastReq{Token: votingSession, RoomId: l.To}
// // // 				h.broadcast <- BroadcastReq{Token: challengeDiscussion, RoomId: l.To}
// // // 				h.broadcast <- BroadcastReq{Token: _startGame, RoomId: l.To}
// // // 				h.broadcast <- BroadcastReq{Token: Unblock, RoomId: l.To}
// // // 				switch true {
// // // 				case getStoredPower[id][l.To][___Nexus]:
// // // 					token := "Nexus: " + Nexus("VIRAT", "_")
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Nexus: " + token,
// // // 						to: id, roomName: l.To, _sleep: false}
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 				case getStoredPower[id][l.To][___Freeze]:
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: FreezePower,
// // // 						to: id, roomName: l.To, _sleep: false}
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 				case getStoredPower[id][l.To][___Rewind]:
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RewindPower,
// // // 						to: id, roomName: l.To, _sleep: false}
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}
// // // 				}
// // // 				h.wg.Add(1)
// // // 				go func() {
// // // 					defer h.wg.Done()
// // // 					newStore := map[string]string{}
// // // 					reverseNewStore := map[string]string{}
// // // 					for _, tag := range saveShuffle {
// // // 						for _, data := range tag {
// // // 							for Name, ID := range getProfData {
// // // 								if ID == data {
// // // 									newStore[Name] = data
// // // 									reverseNewStore[data] = Name
// // // 								}
// // // 							}
// // // 						}
// // // 					}
// // // 					// becuase we have saved the data of joiners separately
// // // 					for _, id := range reverseNewStore {
// // // 						for Ownername, Ownerid := range getOwnerProf {
// // // 							if Ownerid != id {
// // // 								reverseNewStore[Ownerid] = Ownername
// // // 								newStore[Ownername] = Ownerid // its obvious that if the id doenst match it means that it is not in the new store too
// // // 							}
// // // 						}
// // // 					}

// // // 					token1 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][0]]}
// // // 					token2 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][1]]}

// // // 					token3 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][0]]}
// // // 					token4 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][1]]}

// // // 					__RED := "Display: " + token1.NickName + " " + token1.TeamName
// // // 					__RED2 := "Display: " + token2.NickName + " " + token2.TeamName
// // // 					__BLUE := "Display: " + token3.NickName + " " + token3.TeamName
// // // 					__BLUE2 := "Display: " + token4.NickName + " " + token4.TeamName

// // // 					h.broadcast <- BroadcastReq{Token: __RED, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: __RED2, RoomId: l.To}

// // // 					h.broadcast <- BroadcastReq{Token: __BLUE, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: __BLUE2, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: waiting, RoomId: l.To}

// // // 					fmt.Println("new Store: ", newStore)
// // // 					fmt.Println("reverseNewStore: ", reverseNewStore)
// // // 					fmt.Println("getOwnerProf: ", getOwnerProf)

// // // 				}()

// // // 				h.wg.Add(1)
// // // 				go func() {
// // // 					defer h.wg.Done()

// // // 					token := "ChallengeGuess: " + __getChallengeSet[___TeamBlueKey]
// // // 					token2 := "ChallengeGuess: " + GR.MutualVote

// // // 					for _, ids := range saveShuffle[___TeamRedKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: ids, roomName: l.To, _sleep: false}
// // // 					}

// // // 					for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: ids, roomName: l.To, _sleep: false}
// // // 					}

// // // 					fmt.Println("challenge set: ", __getChallengeSet)

// // // 					_trackChallengeSet.done <- true
// // // 				}()
// // // 			} else {
// // // 				fmt.Println("waiting red team")
// // // 				for _, ids := range saveShuffle[___TeamRedKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, to: ids, roomName: l.To, _sleep: false}
// // // 				}
// // // 				for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, to: ids, roomName: l.To, _sleep: false}
// // // 				}
// // // 				go func() {
// // // 					store := map[string]map[string]map[string]map[string]bool{}
// // // 					store[id] = map[string]map[string]map[string]bool{l.To: {___TeamRedKey: {___AgressSession: true}}}
// // // 					saveSessionDone.store <- store
// // // 				}()
// // // 				msg := ""
// // // 				go func() {
// // // 					for _, msgs := range getMsgs.Range() {
// // // 						if msgs != "agree" {
// // // 							msg = msgs
// // // 						}
// // // 					}
// // // 					store := map[string]string{}
// // // 					store[___TeamRedKey] = msg
// // // 					_trackChallengeSet.save <- store
// // // 				}()
// // // 			}
// // // 		}
// // // 	// disagree
// // // 	case !proceed:
// // // 		switch true {

// // // 		// decide till 2 chances
// // // 		case isTeamRed:
// // // 			for _, ids := range saveShuffle[___TeamRedKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: chatToken, to: ids, roomName: l.To, _sleep: false}
// // // 			}
// // // 			if GR.Chances != 2 {
// // // 				go func() {
// // // 					msg := dataset.Queue[string]{}
// // // 					msg.Front(GR.MutualVote)
// // // 					saveMesages.store <- msg
// // // 				}()
// // // 				for _, ids := range saveShuffle[___TeamRedKey] {
// // // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "CurrentChance: 1", to: ids, roomName: l.To, _sleep: false}
// // // 				}
// // // 			} else {
// // // 				fmt.Println("red chances done")

// // // 				// rest is same as done for the proceed
// // // 				if isSessionDone[saveShuffle[___TeamBlueKey][0]][l.To][___TeamBlueKey][___AgressSession] || isSessionDone[saveShuffle[___TeamBlueKey][1]][l.To][___TeamBlueKey][___AgressSession] {
// // // 					proceed = true
// // // 				} else if isSessionDone[saveShuffle[___TeamBlueKey][0]][l.To][___TeamBlueKey][___ChallengeSession] || isSessionDone[saveShuffle[___TeamBlueKey][1]][l.To][___TeamBlueKey][___ChallengeSession] {
// // // 					proceed = true
// // // 				} else {
// // // 					proceed = false
// // // 				}

// // // 				if proceed {
// // // 					switch true {
// // // 					case getStoredPower[id][l.To][___Nexus]:
// // // 						token := "Nexus: " + Nexus("VIRAT", "_")
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Nexus: " + token,
// // // 							to: id, roomName: l.To, _sleep: false}
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 					case getStoredPower[id][l.To][___Freeze]:
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: FreezePower,
// // // 							to: id, roomName: l.To, _sleep: false}
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 					case getStoredPower[id][l.To][___Rewind]:
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RewindPower,
// // // 							to: id, roomName: l.To, _sleep: false}
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}
// // // 					}
// // // 					saveSessionDone.done <- true
// // // 					h.broadcast <- BroadcastReq{Token: votingSession, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: challengeDiscussion, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: _startGame, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: Unblock, RoomId: l.To}
// // // 					token := "ChallengeGuess: " + __getChallengeSet[___TeamBlueKey]
// // // 					token2 := "ChallengeGuess: " + GR.MutualVote

// // // 					for _, ids := range saveShuffle[___TeamRedKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: ids, roomName: l.To, _sleep: false}
// // // 					}

// // // 					for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: ids, roomName: l.To, _sleep: false}
// // // 					}

// // // 					fmt.Println("challenge set: ", __getChallengeSet)
// // // 				} else {
// // // 					for _, ids := range saveShuffle[___TeamRedKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, to: ids, roomName: l.To, _sleep: false}
// // // 					}
// // // 					for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, to: ids, roomName: l.To, _sleep: false}
// // // 					}
// // // 					go func() {
// // // 						store := map[string]map[string]map[string]map[string]bool{}
// // // 						store[id] = map[string]map[string]map[string]bool{l.To: {___TeamRedKey: {___AgressSession: true}}}
// // // 						saveSessionDone.store <- store
// // // 					}()
// // // 					msg := ""
// // // 					go func() {
// // // 						for _, msgs := range getMsgs.Range() {
// // // 							if msgs != "agree" {
// // // 								msg = msgs
// // // 							}
// // // 						}
// // // 						store := map[string]string{}
// // // 						store[___TeamRedKey] = msg
// // // 						_trackChallengeSet.save <- store
// // // 					}()
// // // 				}
// // // 			}

// // // 		case isTeamBlue:
// // // 			for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: chatToken, to: ids, roomName: l.To, _sleep: false}
// // // 			}
// // // 			if GR.Chances != 2 {
// // // 				go func() {
// // // 					msg := dataset.Queue[string]{}
// // // 					msg.Front(GR.MutualVote)
// // // 					store := map[string]map[string]map[string][]string{}
// // // 					store[id] = map[string]map[string][]string{l.To: {___TeamBlueKey: {msg.Pop()}}}
// // // 					saveChatMes <- store
// // // 				}()
// // // 			} else {
// // // 				fmt.Println("red chances done")

// // // 				if isSessionDone[saveShuffle[___TeamRedKey][0]][l.To][___TeamRedKey][___AgressSession] || isSessionDone[saveShuffle[___TeamRedKey][1]][l.To][___TeamRedKey][___AgressSession] {
// // // 					proceed = true
// // // 				} else if isSessionDone[saveShuffle[___TeamRedKey][0]][l.To][___TeamRedKey][___ChallengeSession] || isSessionDone[saveShuffle[___TeamRedKey][1]][l.To][___TeamRedKey][___ChallengeSession] {
// // // 					proceed = true
// // // 				} else {
// // // 					proceed = false
// // // 				}

// // // 				if proceed {
// // // 					saveSessionDone.done <- true
// // // 					token := "ChallengeGuess: " + __getChallengeSet[___TeamRedKey]
// // // 					token2 := "ChallengeGuess: " + GR.MutualVote
// // // 					switch true {
// // // 					case getStoredPower[id][l.To][___Nexus]:
// // // 						token := "Nexus: " + Nexus("VIRAT", "_")
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Nexus: " + token,
// // // 							to: id, roomName: l.To, _sleep: false}
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 					case getStoredPower[id][l.To][___Freeze]:
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: FreezePower,
// // // 							to: id, roomName: l.To, _sleep: false}
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}

// // // 					case getStoredPower[id][l.To][___Rewind]:
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RewindPower,
// // // 							to: id, roomName: l.To, _sleep: false}
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: PowerActivate, to: id, _sleep: false, roomName: l.To}
// // // 					}
// // // 					for _, ids := range saveShuffle[___TeamRedKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: ids, roomName: l.To, _sleep: false}
// // // 					}

// // // 					for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: ids, roomName: l.To, _sleep: false}
// // // 					}

// // // 					h.broadcast <- BroadcastReq{Token: votingSession, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: challengeDiscussion, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: _startGame, RoomId: l.To}
// // // 					h.broadcast <- BroadcastReq{Token: Unblock, RoomId: l.To}
// // // 				} else {
// // // 					for _, ids := range saveShuffle[___TeamRedKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, to: ids, roomName: l.To, _sleep: false}
// // // 					}
// // // 					for _, ids := range saveShuffle[___TeamBlueKey] {
// // // 						h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, to: ids, roomName: l.To, _sleep: false}
// // // 					}
// // // 					go func() {
// // // 						store := map[string]map[string]map[string]map[string]bool{}
// // // 						store[id] = map[string]map[string]map[string]bool{l.To: {___TeamBlueKey: {___AgressSession: true}}}
// // // 						saveSessionDone.store <- store
// // // 					}()
// // // 					msg := ""
// // // 					go func() {
// // // 						for _, msgs := range getMsgs.Range() {
// // // 							if msgs != "agree" {
// // // 								msg = msgs
// // // 							}
// // // 						}
// // // 						store := map[string]string{}
// // // 						store[___TeamBlueKey] = msg
// // // 						_trackChallengeSet.save <- store
// // // 					}()
// // // 				}
// // // 			}
// // // 		}
// // // 	default:
// // // 		fmt.Println("proceed not able to find pattern")
// // // 	}

// // // }

// // import (
// // 	"fmt"
// // 	"strconv"
// // )

// // // whether to set challenge via nutually or indivually
// // func Abundance2(h *Hub, id string) {
// // 	isTeamRed := false
// // 	isTeamBlue := false
// // 	// count for proceeding to the playing hall
// // 	count := GR.MutualVoteCount + GR.ChallengeCount // this will indirectly will turn into 2
// // 	newCount := "CommonCount: " + strconv.Itoa(1)

// // 	//has := dataset.SearchSystem_{}.Constructor()
// // 	// _lower := strings.ToLower(GR.MutualVote)
// // 	// validM := has.Has_Key(GR.MutualVote)//
// // 	// validC := has.Has_Key(GR.ChallengeToken)
// // 	validM := GR.MutualVote != ""     // experimental
// // 	validC := GR.ChallengeToken != "" // experimental

// // 	fmt.Println("Count: ", count)
// // 	fmt.Println("Counts: ", GR.MutualVoteCount, GR.ChallengeCount)

// // 	// setChallenge := "SetChallenge: true"
// // 	// for voting sesion; player get's 3 chances to mutaully decide on challenge set
// // 	// chances := 3

// // 	fmt.Println("Abundance Session")
// // 	fmt.Println("startGame: ", startGame)
// // 	// toPlayingSession := "PlayingSession: true"
// // 	// discussionDone := "DoneDiscussion: true"

// // 	for _, r := range saveShuffle[___TeamRedKey] {
// // 		if r == id {
// // 			isTeamRed = true
// // 		}
// // 	}

// // 	for _, r := range saveShuffle[___TeamBlueKey] {
// // 		if r == id {
// // 			isTeamBlue = true
// // 		}
// // 	}
// // 	fmt.Println("valid vote: ", validM, "valid chall: ", validC)

// // 	// important: this will help in broadcasting the challenge set token
// // 	// store the voting session token
// // 	switch true {
// // 	case isTeamRed:
// // 		storeToken := map[string][]string{}
// // 		storeToken[___TeamRedKey] = append(storeToken[___TeamRedKey], GR.MutualVote)
// // 		mutualAudit.token <- storeToken
// // 	case isTeamBlue:
// // 		storeToken := map[string][]string{}
// // 		storeToken[___TeamBlueKey] = append(storeToken[___TeamBlueKey], GR.MutualVote)
// // 		mutualAudit.token <- storeToken
// // 	default:
// // 		fmt.Println("not able to store the mutaul token")
// // 	}

// // 	switch true {
// // 	case GR.VotingSession:
// // 		// for case:  voting session vs challenge session
// // 		// do the chatting unless players agrees with the value
// // 		// if they unable to agree
// // 		// the last input will be set for the challenge
// // 		// note: to remove the agree token from the option
// // 		// unblocking and blocking
// // 		fmt.Println("track lock: ", isLock)
// // 		fmt.Println("agree: ", GR.Agree)
// // 		// chatToken := "ChatToken: " + GR.MutualVote

// // 		AgreeSession(h, id, count, newCount)

// // 	case GR.ChallengeDiscussion:

// // 		ChallengeSession(h, id)
// // 		fmt.Println("Stored audit: ", _audit)
// // 	default:
// // 		fmt.Println("no mutually")
// // 	}
// // 	fmt.Println("challenge tokens: ", _audit)
// // 	fmt.Println("Stored audit: ", _audit)
// // 	fmt.Println("teamA: ", isTeamRed, "teamB: ", isTeamBlue, "teamA:", saveShuffle[___TeamRedKey], "teamB: ", saveShuffle[___TeamBlueKey])
// // 	fmt.Println("challenge tokens: ", _audit)
// // 	fmt.Println("Stored audit: ", _audit)
// // }

// // switch string(token) {
// // // cases here will be send once the one of the player left the room

// // case "Hosts":
// // 	Hosts = HostRooms(h)
// // 	for r := range Hosts {
// // 		if Hosts[r] != "" {
// // 			broadcast <- "Hosts:" + Hosts[r]
// // 		}
// // 	}

// // case "Specification": // requirements of the room
// // 	log.Println("room sepcification")
// // 	HostFor, Cap := RoomSpecification(h)
// // 	// only if the room exists
// // 	if HostFor != "" {
// // 		// broadcast to the specific client only
// // 		broadcast <- "Sepc:" + HostFor + strconv.Itoa(Cap)
// // 	} else {
// // 		broadcast <- "Spec:" + " no room as been created"
// // 	}
// // }
// // go func() { Game1v1(h, c, l, GR) }() // testing left in the go func
// // Game1v1(h, c, l, GR) // testing done working fine
// // go func() { Game2v2(h, c, l, GR) }()
// // if getSettings[l.To][_TwoVTwoKey] {
// // 	Game2v2(h, c, l, GR)
// // } else {
// // 	Game1v1(h, c, l, GR)

// // }
// // Game(h, c, GR, id)
// // if !l.JoinRoom {
// // 	// todo boarcast to client only
// // } else if l.JoinRoom {

// // } else {
// // 	log.Println("theres no room to be monitored")
// // }
// // if l.Set && l.JoinRoom {
// // 	log.Println("storing c data")
// // 	store := map[string]string{}
// // 	store[l.NickName] = l.ID
// // 	ProfData.data <- store
// // }
// // if l.Set && l.CreateRoom {

// // 	log.Println("storing c data")

// // }
// // go func() {
// // 	clientID.mu.Lock()
// // 	defer clientID.mu.Unlock()
// // 	clientID.clientID <- h.clients[conn].clientID
// // }()

// // h.broadcast <- BroadcastReq{RoomID: RoomID, Token: "HostID: " + Host}
// // h.broadcast <- BroadcastReq{Token: "MyRoom: " + RoomID, RoomID: RoomID}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "MyNickName: " + NickName, to: ID, roomname: RoomID, _sleep: false}
// // storing IDs
// // storeR := make(map[string]string)
// // storeI := make(map[string]map[string]string)
// // storeI[ID] = map[string]string{
// // 	RoomID: NickName,
// // }

// // storeNickNames.viaID <- storeI
// // storeNickNames.viaRoom <- storeR

// // this is working fine
// // count := -1
// // for r := range rooms {
// // 	if rooms[r] == RoomID {
// // 		count = 0
// // 		count = r
// // 	}
// // }
// // if count >= 0 {
// // 	rooms = dataset.EraseOnPos(rooms, count)
// // }
// // RemoveHosts returns remove specific room
// // func RemoveHosts(h *Hub, conn *websocket.Conn) {
// // 	log.Println("removing hosts")
// // 	for r := range Hosts {
// // 		log.Println("gosts inx: ", r)
// // 		log.Println("hosts:", Hosts)
// // 		if Hosts[r] == h.clients[conn].roomID {
// // 			Hosts = dataset.EraseOnPos(Hosts, r)
// // 		}
// // 	}
// // }

// // delete(storeRoom, RoomID)
// // hold
// // RemoveHosts(h, conn) // removes specific room not all
// //delete(h.clients, conn)
// // StopSession(RoomID) // remove the session
// // storeRoomList.done <- true
// // func CountRoomMembers(h *Hub, roomName string) int {
// // 	count := 0
// // 	count = len(h.rooms[roomName])
// // 	log.Println("Count: ", count)
// // 	return count
// // }

// // RoomSpecification returns @RETURN room name and specification of the room; right now it is capacity
// // func RoomSpecification(h *Hub) (string, int) {
// // 	_room := ""
// // 	_capacity := -1
// // 	for room := range storeRoom {
// // 		_room = room
// // 	}
// // 	for _, capacity := range storeRoom {
// // 		_capacity = capacity
// // 	}
// // 	return _room, _capacity
// // }

// // HostRooms returns list of all the rooms
// // func HostRooms(h *Hub) []string {
// // 	for rooms := range storeRoom {
// // 		Hosts = append(Hosts, rooms)
// // 	}
// // 	return Hosts
// // }

// // RoomExists returns true if exists
// // func RoomExists(name string) bool {
// // 	found := false
// // 	for room := range storeRoom {
// // 		if name == room {
// // 			found = true
// // 		}
// // 	}
// // 	return found
// // }

// // import (
// // 	"log"
// // 	"nyg/validate"

// // 	"github.com/gofiber/contrib/websocket"
// // )

// // func Score(h *Hub, count int, id string, conn *websocket.Conn, GR GameRoom) {
// // 	log.Println("in score")

// // 	log.Println("current score 🏉 ", GR.RedScoreCount, GR.BlueScoreCount)

// // 	isTeamRed := false
// // 	isTeamBlue := false
// // 	point := false

// // 	redScore, blueScore := 0, 0

// // 	for _, _id := range saveShuffle[GR.RoomName][_TeamRedKey] {
// // 		if id == _id {
// // 			isTeamRed = true
// // 		}
// // 	}
// // 	for _, _id := range saveShuffle[GR.RoomName][_TeamBlueKey] {
// // 		if id == _id {
// // 			isTeamBlue = true
// // 		}
// // 	}
// // 	log.Println("gameinfo: ", getGameInfo)
// // 	redBet := IsBetSession(h, _TeamRedKey, GR.RoomName)
// // 	blueBet := IsBetSession(h, _TeamBlueKey, GR.RoomName)
// // 	dict, list := getGameInfo[GR.RoomName].SetDictionary, getGameInfo[GR.RoomName].TeamRedSetChallenge
// // 	cups := getPreviousValidateToken[GR.RoomName] // note includes bet values too but this will get neglect if and only if the bet is not used

// // 	if redBet || blueBet {
// // 		list = getGameInfo[GR.RoomName].TeamRedSetChallenge
// // 	} else if !isTeamBlue {
// // 		list = getGameInfo[GR.RoomName].TeamBlueSetChallenge
// // 	}

// // 	log.Println("list: ", list)
// // 	point, Vstore := validate.SportsValidate(dict, list, GR.Guess, cups)
// // 	// Round1 SheetUpdate Nickname: result
// // 	// note: to check on the opppsite team rather than own
// // 	go func() {
// // 		store := map[string][]string{}
// // 		store[GR.RoomName] = Vstore
// // 		storeValidateToken.store <- store
// // 	}()

// // 	log.Println("previousValues: ", getPreviousValidateToken)

// // 	// note: not to consider the bet value same others when used bet power
// // 	if redBet || blueBet {
// // 		// blue score
// // 		point = BetValidation(h, GR.BetOn, GR.RoomName)
// // 		Vstore = append(Vstore, GR.BetOn)

// // 	} else {
// // 		Vstore = append(Vstore, cups...)
// // 		UpdateFrame(h, GR.RoomName, Vstore)
// // 	}

// // 	isRewindSet := false

// // 	log.Println("getstored power: ", getStoredPower)

// // 	for _, rooms := range saveShuffle {
// // 		for _, ids := range rooms {
// // 			for _, id := range ids {
// // 				if getStoredPower[id][GR.RoomName][_RewindKey] {
// // 					isRewindSet = true
// // 				}
// // 			}
// // 		}
// // 	}

// // 	log.Println("isRewind: ", isRewindSet)

// // 	go func() {
// // 		store := map[string][]string{}
// // 		store[GR.RoomName] = []string{GR.Guess}
// // 		storeCups.store <- store
// // 	}()
// // 	if isRewindSet {
// // 		// todo: reset the round and broadcast the value send by the other player
// // 		// once the gr.rewind has been used wait till the next input
// // 		// next input cases:
// // 		// case the player run's out of time // still the rewind happens
// // 		// case the player's input reached // rewind happens
// // 		PRewindMechanism(h, GR.RoomName, id, "", false)
// // 		PowersCleanUp() // cleanup the power too
// // 	} else {

// // 		// if none of the powers has used
// // 		switch point && !GR.TimeUp {

// // 		case true:
// // 			log.Println("right guess")

// // 			log.Println("right guess point: ", point)
// // 			for _, _id := range saveShuffle[GR.RoomName][_TeamRedKey] {
// // 				if id == _id {
// // 					isTeamRed = true
// // 				}
// // 			}
// // 			for _, _id := range saveShuffle[GR.RoomName][_TeamBlueKey] {
// // 				if id == _id {
// // 					isTeamBlue = true
// // 				}
// // 			}

// // 			// if the token matches any one of these than switch to that room

// // 			if id == saveShuffle[GR.RoomName][_TeamRedKey][0] || id == saveShuffle[GR.RoomName][_TeamBlueKey][0] {
// // 				count = 0
// // 			} else {
// // 				count = 1
// // 			}

// // 			log.Println("is team A: ", isTeamRed, "is team B: ", isTeamBlue)

// // 			switch true {
// // 			case isTeamBlue:
// // 				log.Println("team Blue right case")

// // 				blueScore = GR.BlueScoreCount + 1
// // 				redScore = GR.RedScoreCount

// // 				Update(h, true, GR.RoomName)
// // 			case isTeamRed:
// // 				log.Println("team Red right case")

// // 				blueScore = GR.BlueScoreCount
// // 				redScore = GR.BlueScoreCount + 1

// // 				Update(h, false, GR.RoomName)

// // 			default:
// // 				log.Println("none")
// // 			}

// // 		case false:
// // 			log.Println("wrong guess")

// // 			log.Println("wrong guess point: ", point)

// // 			log.Println("right guess point: ", point)
// // 			for _, _id := range saveShuffle[GR.RoomName][_TeamRedKey] {
// // 				if id == _id {
// // 					isTeamRed = true
// // 				}
// // 			}
// // 			for _, _id := range saveShuffle[GR.RoomName][_TeamBlueKey] {
// // 				if id == _id {
// // 					isTeamBlue = true
// // 				}
// // 			}

// // 			// if the token matches any one of these than switch to that room

// // 			if id == saveShuffle[GR.RoomName][_TeamRedKey][0] || id == saveShuffle[GR.RoomName][_TeamBlueKey][0] {
// // 				count = 0
// // 			} else {
// // 				count = 1
// // 			}

// // 			log.Println("is team A: ", isTeamRed, "is team B: ", isTeamBlue)

// // 			switch true {
// // 			case isTeamRed:
// // 				log.Println("team Red wrong")

// // 				blueScore = GR.BlueScoreCount + 1
// // 				redScore = GR.RedScoreCount

// // 				Update(h, true, GR.RoomName)
// // 			case isTeamBlue:
// // 				log.Println("team Blue wrong")

// // 				blueScore = GR.BlueScoreCount
// // 				redScore = GR.RedScoreCount + 1

// // 				Update(h, false, GR.RoomName)

// // 			default:
// // 				log.Println("none")
// // 			}

// // 		}
// // 		if GR.TimeUp {
// // 			// same as wrong answer
// // 			switch isTeamBlue {
// // 			case true:
// // 				Update(h, false, GR.RoomName)

// // 			case false:
// // 				Update(h, true, GR.RoomName)
// // 			}
// // 		}

// // 		// important else we won't be able to get the sync value to test if the game is over or not in the restart function
// // 		Restart(h, count, id, conn, redScore, blueScore, GR)

// // 	}
// // 	log.Println("locking at count: ", count)
// // 	log.Println("teams: ", saveShuffle[GR.RoomName][_TeamRedKey], saveShuffle[GR.RoomName][_TeamBlueKey])

// // }

// // type Validation struct {
// // 	store chan map[string][]string
// // 	done  chan bool
// // }

// // var (
// // 	// room-name :tokens
// // 	storeValidateToken       = Validation{store: make(chan map[string][]string), done: make(chan bool)}
// // 	getPreviousValidateToken = make(map[string][]string)
// // )

// // validation mechanism:
// // talking about without power:
// // rules for scoring:
// // 1. the token must not be the match for previous token
// // 2. the token must be in the given challenge set
// // mechanism:
// // 1. the token passed will be saved in the common dictionary
// // 2. the token passed will remove all the similary keys too
// // for example: in the dictionary of X there are 5 similar tokens named jimmy but their middle and surname is different
// // since the token is jimmy than all the jimmy's will be disassocaited from the list and
// // will be considered in the list of the previous token match

// // type TeamData struct {
// // 	TeamName string
// // 	NickName string
// // }

// // var (
// // 	latest bool
// // )

// // ChallengeCount int `json:"challengeCount"` // track if both of the players has set the challenge
// // VotingSession       bool `json:"votingSession"`       // on going voting session
// // MutualSession       bool `json:"mutalSession"`        // on going vote agreeing for mutual voting
// //	MutualVoteCount int `json:"mutualVoteCount"` // tracking the current voting
// //	MutalCount      int `json:"mutalCount"`      // agress for the discussion
// //	MutualVote      string `json:"mutualVote"`      // vote to set the challenge value
// // Chances        int `json:"chanceLeft"`     // track the current chance left
// // Agree               bool `json:"agree"`               // if the player mutually decided the token

// // type WLpattern struct {
// // 	store chan map[string]map[string]bool // team name: room-name: anything : true|false
// // 	close chan bool                       // reseting
// // }

// // type WLpattern2 struct {
// // 	store chan map[string]map[string]map[string]bool // team name: room-name: anything : true|false
// // 	close chan bool                                  // reseting
// // }

// // var (
// // 	// stores room and its room setting if the set toss is on or off
// // 	settings    = WLpattern{store: make(chan map[string]map[string]bool)} // example room-name:{SetupToss=false} or room-name:{winners:false}
// // 	getSettings = make(map[string]map[string]bool)

// // 	starterSettings   = WLpattern2{store: make(chan map[string]map[string]map[string]bool)}
// // 	getStarterSetting = map[string]map[string]map[string]bool{}
// // )

// // store := map[string]map[string]string{}
// // store[l.RoomName] = map[string]string{
// // 	l.NickName: l.ID,
// // }
// // SaveOwnerProf <- store
// // if _tempGameSession[l.To] || _tempGameSession[GetFriendRoom(l.Code)] || _tempGameSession[GetPrivateRoom(l.Code)] || gameSessionStarted[l.To] || gameSessionStarted[GetFriendRoom(l.Code)] || gameSessionStarted[GetPrivateRoom(l.Code)] {

// // 	// setup the environament
// // } else {
// // 	log.Println("team has not created yet waiting for players")
// // }

// // case l.LeaveRoom:
// // 	{
// // 		log.Println("leaving room")
// // 		h.unregister <- conn
// // 	}

// // _tempShuffle = map[string]map[string][]string{}     // for-safety so that we can use the saveShuffle in next round
// // keeps track of id's that are in the playing hall
// // roomStoreClientID = ClientIDStorageRoom{clientID: make(chan string)}
// // cIDstorage        = []string{}
// // CountFriend       = make(chan int) // tracks the current count of the friend
// // _CountFriend      int

// // storing-joiner: the client name with its id
// // ProfData    = ProfileData{data: make(chan map[string]string)} // saves the profile data
// // getProfData = make(map[string]string)

// // // storing-creator: room name, nickname and id
// // SaveOwnerProf = make(chan map[string]map[string]string) // saves the profile data
// // getOwnerProf  = make(map[string]string)                 // returning nickname with its associaetd id

// // storeRewindPower = make(chan map[string]map[string]bool)
// // getRewindPower   = make(map[string]map[string]bool)
// // type TimePattern struct {
// // 	GameTime       int `json:"gameTime"`
// // 	DiscussionTime int `json:"discussionTime"`
// // }

// // func SendRoomTime(h *Hub, RoomName string) {
// // 	gameTime := getRoomRequestedToken[RoomName][_GameTimeKey]
// // 	discussionTime := getRoomRequestedToken[RoomName][_DecisionTimeKey]
// // 	var tp TimePattern
// // 	d, _ := strconv.Atoi(discussionTime)
// // 	g, _ := strconv.Atoi(gameTime)

// // 	tp.DiscussionTime = d
// // 	tp.GameTime = g
// // 	token, _ := json.Marshal(&tp)
// // 	_for := "RoomTime:"
// // 	h.broadcast <- BroadcastReq{RoomID: RoomName, Token: _for + string(token)}
// // }
// // // TrashChallenge broadcasts to remove the value from list
// // // note: to use this function at the end of all the service done
// // func TrashChallenge(h *Hub, RoomName string, isSessionDone bool, teamname string, opponentTeamname string) {
// // 	TeamChallenge := ""
// // 	TeamOpponentChallenge := ""
// // 	for _, _id := range saveShuffle[RoomName][teamname] {
// // 		if getPlayerInfo[_id][RoomName][teamname].isLock {
// // 			TeamChallenge = getTokens[_id][RoomName][teamname].Challenge
// // 		}
// // 	}
// // 	for _, _id := range saveShuffle[RoomName][opponentTeamname] {
// // 		if getPlayerInfo[_id][RoomName][opponentTeamname].isLock {
// // 			TeamOpponentChallenge = getTokens[_id][RoomName][opponentTeamname].Challenge
// // 		}
// // 	}
// // 	RCtoken := RemoveChallenge + TeamChallenge
// // 	BCtoken := RemoveChallenge + TeamOpponentChallenge

// // 	if isSessionDone {
// // 		// case true:
// // 		h.broadcast <- BroadcastReq{RoomID: RoomName, Token: RCtoken}
// // 		h.broadcast <- BroadcastReq{RoomID: RoomName, Token: BCtoken}

// // 		// case false:
// // 		// 	// broadcast to the teams about the set challenge and cannot used
// // 		// 	if isTeamBlue {
// // 		// 		h.broadcast <- BroadcastReq{RoomID: RoomName, Token: BCtoken}
// // 		// 	} else {
// // 		// 		h.broadcast <- BroadcastReq{RoomID: RoomName, Token: RCtoken}
// // 		// 	}
// // 	}
// // }

// // // TrashDictionary broadcasts to remove the value from list
// // // note: to use this function at the end of all the service done
// // func TrashDictionary(h *Hub, RoomName string, event string) {
// // 	SetDictionary := event
// // 	token := RemoveDictionary + SetDictionary
// // 	h.broadcast <- BroadcastReq{RoomID: RoomName, Token: token}
// // }

// // import (
// // 	"encoding/json"
// // 	"log"
// // )

// // // -----
// // // bet to be use instead of challenge thats simple

// // // BetValidation validation is based on stored cups and betted value
// // // true if the token != betted value
// // func BetValidation(h *Hub, token string, roomname string) bool {
// // 	va := getStoredBetValue[roomname]

// // 	cups := getStoredCups[roomname]
// // 	point := false
// // 	if va == token {
// // 		point = false
// // 	}
// // 	if !point {
// // 		for _, cup := range cups {
// // 			if cup != va {
// // 				if token == cup {
// // 					point = true
// // 				}
// // 			}
// // 		}
// // 	}

// // 	return point
// // }

// // func BetSession(h *Hub, id string, roomname string, challengeToken string, betOn string) {
// // 	isTeamName := false
// // 	proceed := false
// // 	// ID_ := ""
// // 	log.Println("in bet session")
// // 	teamname, opponentTeamname := TokenFromTeam(id, roomname)
// // 	// to do search for the true bet key
// // 	// if true send back the picker done
// // 	// if the previous player has done his challenge session than
// // 	// proceed to game hall
// // 	// else send back the game start
// // 	for _, _id := range saveShuffle[roomname][teamname] {
// // 		if _id == id {
// // 			isTeamName = true
// // 		}
// // 	}

// // 	// search for the previous session done
// // 	switch isTeamName {
// // 	case true:
// // 		for _, _id := range saveShuffle[roomname][opponentTeamname] {
// // 			if getSessionUpdate[_id][roomname][opponentTeamname].ChallengeDone {
// // 				proceed = true
// // 				//ID_ = _id
// // 			}
// // 		}
// // 		for _, r := range saveShuffle[roomname][teamname] {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 				token: Block, to: r, _sleep: false, roomname: roomname,
// // 			}
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 				token: _waiting, to: r, _sleep: false, roomname: roomname,
// // 			}
// // 		}
// // 		for _, r := range saveShuffle[roomname][opponentTeamname] {
// // 			log.Println("team red: ", r)

// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 				token: Unblock, to: r, _sleep: false, roomname: roomname,
// // 			}

// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 				token: waiting, to: r, _sleep: false, roomname: roomname,
// // 			}
// // 		}

// // 		go func() {
// // 			store := map[string]map[string]map[string]Session{}
// // 			store[id] = map[string]map[string]Session{roomname: {teamname: Session{BetDone: true}}}
// // 			storeSessionUpdate <- store
// // 		}()

// // 	case false:
// // 		for _, _id := range saveShuffle[roomname][teamname] {
// // 			if getSessionUpdate[_id][roomname][teamname].ChallengeDone {
// // 				proceed = true
// // 				//ID_ = _id
// // 			}
// // 		}

// // 		// wait till the game start and store the set token
// // 		for _, r := range saveShuffle[roomname][teamname] {
// // 			log.Println("team blue: ", r)
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 				token: Unblock, to: r, _sleep: false, roomname: roomname,
// // 			}

// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 				token: waiting, to: r, _sleep: false, roomname: roomname,
// // 			}
// // 		}

// // 		for _, r := range saveShuffle[roomname][opponentTeamname] {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 				token: Block, to: r, _sleep: false, roomname: roomname,
// // 			}
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 				token: _waiting, to: r, _sleep: false, roomname: roomname,
// // 			}
// // 		}
// // 		store := map[string]map[string]Session{}
// // 		se := Session{
// // 			TossDone:       true,
// // 			ChallengeDone:  true,
// // 			DictionaryDone: true,
// // 			BetDone:        true,
// // 			SessionDone:    false,
// // 		}
// // 		store[roomname] = map[string]Session{
// // 			opponentTeamname: se,
// // 		}
// // 	}
// // 	log.Println("proceed: ", proceed)

// // 	// trigger the power
// // 	PBetMechanism(h, roomname, id, teamname, opponentTeamname, challengeToken, betOn)
// // }

// // // IsBetSession  note: teamname must be the teamname to search if the bet is used
// // func IsBetSession(h *Hub, teamname string, roomname string) bool {
// // 	log.Println("in finding if the client has used bet session")
// // 	log.Println("stored power: ", getStoredPower)

// // 	betPowerSignal := false

// // 	for _, _ID := range saveShuffle[roomname][teamname] {
// // 		if getStoredPower[_ID][roomname][_BetKey] {
// // 			betPowerSignal = true
// // 		}
// // 	}

// // 	return betPowerSignal
// // }

// // func ProceedBet(h *Hub, ID string, roomname string, teamname string, opponenTeamname string, DictionaryName string) {
// // 	log.Println("in proceed bet")
// // 	ID_ := ""
// // 	x := DictionaryName
// // 	token := CreateBetCups(h, roomname, true, x)
// // 	parcel, _ := json.Marshal(&token)
// // 	log.Println("dict name: ", x)
// // 	isTeamName := false
// // 	for _, _id := range saveShuffle[roomname][teamname] {
// // 		if ID == _id {
// // 			isTeamName = true
// // 			break
// // 		}
// // 	}

// // 	go func() {
// // 		store := map[string][]string{}
// // 		store[roomname] = append(store[roomname], token.FirstCup, token.SecondCup, token.ThirdCup)
// // 		storeValidateToken.store <- store
// // 	}()

// // 	log.Println("parcel: ", parcel)
// // 	if isTeamName {
// // 		for _, _ID := range saveShuffle[roomname][teamname] {
// // 			if getStoredPower[_ID][roomname][_BetKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: BetPicker, roomname: roomname, _sleep: false, to: _ID}
// // 				ID_ = _ID

// // 			}
// // 		}
// // 	} else {
// // 		for _, _ID := range saveShuffle[roomname][opponenTeamname] {
// // 			if getStoredPower[_ID][roomname][_BetKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: BetPicker, roomname: roomname, _sleep: false, to: _ID}
// // 				ID_ = _ID
// // 			}
// // 		}
// // 	}
// // 	log.Println("sending bet to: ", ID_)
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: BetPicker, roomname: roomname, _sleep: false, to: ID_}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: challengeDiscussion, roomname: roomname, _sleep: false, to: ID_}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: BetItem + string(parcel), roomname: roomname, _sleep: false, to: ID_}
// // 	}()
// // }

// // func SendBetTokens(h *Hub, ID string, roomname string, DictionaryName string) {
// // 	log.Println("in send bet tokens")

// // 	x := DictionaryName

// // 	h.wg.Go(func() {
// // 		token := CreateBetCups(h, roomname, true, x)
// // 		parcel, _ := json.Marshal(&token)
// // 		log.Println("parcel: ", string(parcel))
// // 		log.Println("dict name: ", DictionaryName)
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			token: BetItem + string(parcel), _sleep: false, to: ID, roomname: roomname,
// // 		}

// // 		go func() {
// // 			store := map[string][]string{}
// // 			store[roomname] = append(store[roomname], token.FirstCup, token.SecondCup, token.ThirdCup)
// // 			storeValidateToken.store <- store
// // 		}()

// // 		log.Println("token: ", token)
// // 	})

// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 		token: BetUse, to: ID, roomname: roomname, _sleep: false,
// // 	}

// // 	go func() {
// // 		store := map[string]map[string]map[string]bool{}
// // 		store[ID] = map[string]map[string]bool{
// // 			roomname: {_BetKey: true},
// // 		}

// // 		storePower.store <- store
// // 	}()

// // 	go func() {
// // 		store := map[string]string{}
// // 		store[roomname] = getGameInfo[roomname].SetEvent
// // 		storeBetEvent.store <- store
// // 	}()

// // }

// // for storing the room setting
// // _StarterKey     = "Starter"
// // _SetupTossKey   = "SetupToss"
// // _ReverseKey     = "ReverseKey"
// // _FriendRoomKey  = "FriendRoom"
// // _PrivateRoomKey = "PrivateRoom"
// // _PublicRoomKey  = "PublicRoom"
// // _FieldKey       = "Field"
// // _CategoryKey    = "Category"
// // _BookKey        = "BookKey"
// // _TwoVTwoKey     = "TwoVTwoKey"
// // _OneVOneKey     = "OnevOneKey"

// // for broadcasting the powers in playing room
// // _NexusPowerKey  = "PowerNexus"
// // _TagPowerKey    = "PowerTag"
// // _RewindPowerKey = "PowerRewind"
// // _FreezePowerKey = "PowerFreeze"
// // _DrawPowerKey   = "PowerDraw"
// // _CovertPowerKey = "PowerCovert"
// // _BetPowerKey    = "PowerBet"

// // removes the room once the session is done
// // case done := <-roomStore.done:
// // 	if done {
// // 		log.Println("closing ROOM")

// // 	} else {
// // 		log.Println("not closing ROOM")
// // 	}
// // type GameInfo struct {
// // 	NickNames            []string
// // 	Ids                  []string
// // 	NicknameViaID        map[string]string // nickname->..
// // 	IDViaNickname        map[string]string // id->..
// // 	PlayersRating        map[string]string // id->...
// // 	PlayersGamePlayed    map[string]int    // id->...
// // 	PlayersTier          map[string]string // id->..
// // 	TeamRedSetChallenge  string
// // 	TeamBlueSetChallenge string
// // 	RoomName             string
// // 	SetTossBody          map[string]string // teamname -> ....
// // 	SetDictionary        string
// // 	SetEvent             string
// // }

// // type Playerinfo struct {
// // 	isLock       bool
// // 	isBlock      bool
// // 	hasUsedPower bool
// // }
// // type Session struct {
// // 	TossDone       bool
// // 	ChallengeDone  bool
// // 	DictionaryDone bool
// // 	BetDone        bool
// // 	SessionDone    bool
// // }
// // type Tokens struct {
// // 	Toss           string
// // 	Dictionary     string
// // 	Challenge      string
// // 	Bet            string
// // 	Count          int16
// // 	FinalBossID    string
// // 	CurrentChances int
// // 	OnFire         int
// // 	Guess          string
// // }

// // var (
// // // roomname->..
// // // storeGameInfo = make(chan map[string]GameInfo)
// // // getGameInfo   = make(map[string]GameInfo)
// // // room-name->
// // // addDictionary = make(chan map[string]string)
// // // // room-name->..
// // // addTeamRedChallenge  = make(chan map[string]string)
// // // addTeamBlueChallenge = make(chan map[string]string)
// // // // room-name->..
// // // addEvent = make(chan map[string]string)
// // // // roomn-name->team-name->..
// // // addToss = make(chan map[string]map[string]string)

// // // // removes set challenges, set dictionary and toss from the gameinfo of the requested room
// // // resetTokens = make(chan map[string]bool)

// // // id-> roomname->teamname->..
// // // storePlayerInfo = make(chan map[string]map[string]map[string]Playerinfo)
// // // getPlayerInfo   = make(map[string]map[string]map[string]Playerinfo)

// // // saveBlock = make(chan map[string]map[string]map[string]Playerinfo)
// // // saveLock  = make(chan map[string]map[string]map[string]Playerinfo)

// // // id->roomname->teamname->...
// // // storeSessionUpdate = make(chan map[string]map[string]map[string]Session)
// // // getSessionUpdate   = make(map[string]map[string]map[string]Session)

// // // resets the session to false of provided room
// // // roomname->...
// // // resetSesssion = make(chan map[string]bool)
// // // id->roomname->teamname->...
// // // storeTokensUpdate = make(chan map[string]map[string]map[string]Tokens)
// // // getTokens         = make(map[string]map[string]map[string]Tokens)
// // )

// // type StoreFriendRoom struct {
// // 	store      chan map[string]map[string]bool // room name: friend : true | false
// // 	roomCode   chan map[string]string          // room name: code
// // 	done       chan bool
// // 	removeCode chan bool
// // 	mu         sync.Mutex
// // }

// // var (
// // 	storeFriendRoom = StoreFriendRoom{store: make(chan map[string]map[string]bool), roomCode: make(chan map[string]string), done: make(chan bool),
// // 		removeCode: make(chan bool)}
// // 	getFriendRoomCode = map[string]string{}
// // 	getFriendRoom     = map[string]map[string]bool{}

// // 	// saves room name and if the player has joined or not
// // 	PlayerJoined    = make(chan map[string]bool) // room name:true | true
// // 	hasPlayerJoined = map[string]bool{}

// // 	saveJoinFriendID     = make(chan map[string]map[string]string) // room name: friend: player id
// // 	getSavedJoinFriendID = map[string]map[string]string{}

// // 	saveOwnerFriendID     = make(chan map[string]map[string]string)
// // 	getSavedOwnerFriendID = map[string]map[string]string{}

// // 	// saveSessionDone = CommonPattern{store: make(chan map[string]map[string]map[string]map[string]bool), done: make(chan bool)}
// // 	// isSessionDone   = map[string]map[string]map[string]map[string]bool{}

// // 	// session name: room-name : true| false
// // 	// storeSessionDone = GameSessionTrack{store: make(chan map[string]map[string]bool), done: make(chan bool)}
// // 	// getSessionDone   = map[string]map[string]bool{}
// // )

// // type GameSessionTrack struct {
// // 	store chan map[string]map[string]bool
// // 	done  chan bool
// // }

// // type CommonPattern struct {
// // 	store chan map[string]map[string]map[string]map[string]bool // id: room-name: team-name: any: true|false
// // 	done  chan bool
// // 	mu    sync.Mutex
// // }

// // // TrackPattern stores team a id and other required tokens
// // type TrackPattern struct {
// // 	store chan map[string]map[string]string
// // 	done  chan bool
// // 	mu    sync.Mutex
// // 	wg    sync.WaitGroup
// // }

// // type TrackMsgs struct {
// // 	store chan dataset.Queue[string]
// // 	done  chan bool
// // 	mu    sync.Mutex
// // 	wg    sync.WaitGroup
// // }

// // // FUpdateSheet saves the sheet
// // func FUpdateSheet(storage map[string]PackSheet, save map[string]map[string]map[string]string) {
// // 	temp := storage
// // 	roomNames := reflect.ValueOf(temp).MapKeys()

// // 	storeSheetKeys := []string{}
// // 	// temp2 := map[string]map[string]map[string]string{}
// // 	for _, key := range roomNames {
// // 		for res := range temp[key.String()].Sheet {
// // 			storeSheetKeys = append(storeSheetKeys, res)
// // 		}
// // 	}

// // 	for _, key := range roomNames {
// // 		for _, key2 := range storeSheetKeys {
// // 			for keyx, res := range temp[key.String()].Sheet[key2] {
// // 				if _, ok := save[key.String()]; !ok {
// // 					save[key.String()] = map[string]map[string]string{}
// // 				}
// // 				if _, ok := save[key.String()][key2]; !ok {
// // 					save[key.String()][key2] = map[string]string{}
// // 				}

// // 				// Assign value without overwriting previous content
// // 				save[key.String()][key2][keyx] = res
// // 			}
// // 		}
// // 	}

// // }

// // import (
// // 	"log"
// // 	"math/rand/v2"
// // )

// // // DeadLock returns blocking player so that the toss the coin and unblock player will pick the toss
// // func DeadLock(h *Hub, turn int, roomname string, shuffle map[string]map[string][]string) {
// // 	// turn is basically nothing but passing the default block
// // 	// logic:
// // 	// one of the locked player will get block
// // 	// meaning that the unblock player will able to pick the toss
// // 	log.Println("in dead block")
// // 	log.Println("room name: ", roomname)

// // 	roll := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
// // 	rand.Shuffle(len(roll), func(i int, j int) {
// // 		roll[i], roll[j] = roll[j], roll[i]
// // 	})
// // 	// _block := true
// // 	// _unblock := true

// // 	isTeamRed := false
// // 	isTeamBlue := true

// // 	switch turn {
// // 	case 0:
// // 		log.Println("case 0 blocking")
// // 		log.Println("shuffle: ", shuffle)
// // 		if roll[0]%2 == 0 {
// // 			store := map[string]map[string]map[string]Playerinfo{}

// // 			for _, ID := range shuffle[roomname][_TeamRedKey] {
// // 				if _, ok := store[ID]; !ok {
// // 					store[ID] = map[string]map[string]Playerinfo{}
// // 				}
// // 				if _, ok := store[ID][roomname]; !ok {
// // 					store[ID][roomname] = map[string]Playerinfo{}
// // 				}
// // 				if _, ok := store[ID][roomname][_TeamRedKey]; !ok {
// // 					store[ID][roomname][_TeamRedKey] = Playerinfo{isBlock: true}
// // 				}

// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					to: ID, roomname: roomname, _sleep: false, token: Block,
// // 				}
// // 			}

// // 			for _, ID := range shuffle[roomname][_TeamBlueKey] {
// // 				if _, ok := store[ID]; !ok {
// // 					store[ID] = map[string]map[string]Playerinfo{}
// // 				}
// // 				if _, ok := store[ID][roomname]; !ok {
// // 					store[ID][roomname] = map[string]Playerinfo{}
// // 				}
// // 				if _, ok := store[ID][roomname][_TeamBlueKey]; !ok {
// // 					store[ID][roomname][_TeamBlueKey] = Playerinfo{isBlock: false}
// // 				}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					to: ID, roomname: roomname, _sleep: false, token: Unblock,
// // 				}
// // 			}
// // 			log.Println("send done")
// // 			storePlayerInfo <- store

// // 		} else {

// // 			log.Println("sending")
// // 			store := map[string]map[string]map[string]Playerinfo{}

// // 			for _, ID := range shuffle[roomname][_TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					to: ID, roomname: roomname, _sleep: false, token: Unblock,
// // 				}
// // 				if _, ok := store[ID]; !ok {
// // 					store[ID] = map[string]map[string]Playerinfo{}
// // 				}
// // 				if _, ok := store[ID][roomname]; !ok {
// // 					store[ID][roomname] = map[string]Playerinfo{}
// // 				}
// // 				if _, ok := store[ID][roomname][_TeamRedKey]; !ok {
// // 					store[ID][roomname][_TeamRedKey] = Playerinfo{isBlock: false}
// // 				}
// // 			}
// // 			for _, ID := range shuffle[roomname][_TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					to: ID, roomname: roomname, _sleep: false, token: Block,
// // 				}
// // 				if _, ok := store[ID]; !ok {
// // 					store[ID] = map[string]map[string]Playerinfo{}
// // 				}
// // 				if _, ok := store[ID][roomname]; !ok {
// // 					store[ID][roomname] = map[string]Playerinfo{}
// // 				}
// // 				if _, ok := store[ID][roomname][_TeamBlueKey]; !ok {
// // 					store[ID][roomname][_TeamBlueKey] = Playerinfo{isBlock: true}
// // 				}
// // 			}
// // 			log.Println("send done")
// // 			storePlayerInfo <- store

// // 		}
// // 	case 1:
// // 		store := map[string]map[string]map[string]Playerinfo{}

// // 		// if the player blocked was from team A than block team B player and vice versa
// // 		for _, ID := range shuffle[roomname][_TeamRedKey] {
// // 			if getPlayerInfo[ID][roomname][_TeamRedKey].isBlock {
// // 				isTeamRed = true
// // 			}

// // 		}
// // 		for _, _ID := range shuffle[roomname][_TeamBlueKey] {
// // 			if getPlayerInfo[_ID][roomname][_TeamBlueKey].isBlock {
// // 				isTeamBlue = true
// // 			}
// // 		}

// // 		switch true {
// // 		case isTeamRed:
// // 			// here we basically finding the block player and unblocking him and vice versa
// // 			// x := shuffle[roomname][_TeamBlueKey][0]
// // 			// y := shuffle[roomname][_TeamBlueKey][1]

// // 			/// // // //
// // 			//  if player from team A was block
// // 			//  unblock the team B player
// // 			//  note: this is done for fairplay
// // 			//  we are just playing around the lock players
// // 			// // // /

// // 			// unblock locked b player
// // 			// note: only all of them
// // 			// because we dont know who was the first one to get block
// // 			storeInfo := map[string]map[string]map[string]Playerinfo{}
// // 			for _, ID := range saveShuffle[roomname][_TeamBlueKey] {
// // 				if getPlayerInfo[ID][roomname][_TeamBlueKey].isLock {

// // 					if _, ok := store[ID]; !ok {
// // 						store[ID] = map[string]map[string]Playerinfo{}
// // 					}
// // 					if _, ok := store[ID][roomname]; !ok {
// // 						store[ID][roomname] = map[string]Playerinfo{}
// // 					}
// // 					if _, ok := store[ID][roomname][_TeamBlueKey]; !ok {
// // 						store[ID][roomname][_TeamBlueKey] = Playerinfo{isBlock: false}
// // 					}

// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: Unblock, to: ID, _sleep: false, roomname: roomname,
// // 					}
// // 				}
// // 			}
// // 			for _, ID := range shuffle[roomname][_TeamBlueKey] {
// // 				if getPlayerInfo[ID][roomname][_TeamBlueKey].isBlock {
// // 					isTeamBlue = true
// // 				}
// // 			}
// // 			// block the team A players
// // 			for _, _ID := range saveShuffle[roomname][_TeamRedKey] {
// // 				if getPlayerInfo[_ID][roomname][_TeamRedKey].isLock {
// // 					if _, ok := storeInfo[_ID]; ok {
// // 						storeInfo[_ID] = map[string]map[string]Playerinfo{}
// // 					}
// // 					if _, ok := storeInfo[_ID][roomname]; ok {
// // 						storeInfo[_ID][roomname] = map[string]Playerinfo{}
// // 					}
// // 					if _, ok := storeInfo[_ID]; ok {
// // 						storeInfo[_ID][roomname][_TeamRedKey] = Playerinfo{isBlock: true}
// // 					}

// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: Block, to: _ID, _sleep: false, roomname: roomname,
// // 					}
// // 				}
// // 			}
// // 			saveBlock <- storeInfo

// // 		case isTeamBlue:

// // 			/// // // //
// // 			//  if player from team B was block
// // 			//  unblock the team A player
// // 			//  note: this is done for fairplay
// // 			//  we are just playing around the lock players
// // 			// // // /

// // 			// unblock locked b player
// // 			// note: only all of them
// // 			// because we dont know who was the first one to get block

// // 			storeInfo := map[string]map[string]map[string]Playerinfo{}
// // 			for _, _ID := range saveShuffle[roomname][_TeamBlueKey] {
// // 				if getPlayerInfo[_ID][roomname][_TeamBlueKey].isLock {
// // 					if _, ok := storeInfo[_ID]; ok {
// // 						storeInfo[_ID] = map[string]map[string]Playerinfo{}
// // 					}
// // 					if _, ok := storeInfo[_ID][roomname]; ok {
// // 						storeInfo[_ID][roomname] = map[string]Playerinfo{}
// // 					}
// // 					if _, ok := storeInfo[_ID]; ok {
// // 						storeInfo[_ID][roomname][_TeamBlueKey] = Playerinfo{isBlock: true}
// // 					}

// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: Block, to: _ID, _sleep: false, roomname: roomname,
// // 					}
// // 				}
// // 			}
// // 			for _, _ID := range shuffle[roomname][_TeamBlueKey] {
// // 				if getPlayerInfo[_ID][roomname][_TeamBlueKey].isBlock {
// // 					isTeamBlue = true
// // 				}
// // 			}
// // 			// block the team A players
// // 			for _, _ID := range saveShuffle[roomname][_TeamRedKey] {
// // 				if getPlayerInfo[_ID][roomname][_TeamRedKey].isLock {
// // 					if _, ok := storeInfo[_ID]; ok {
// // 						storeInfo[_ID] = map[string]map[string]Playerinfo{}
// // 					}
// // 					if _, ok := storeInfo[_ID][roomname]; ok {
// // 						storeInfo[_ID][roomname] = map[string]Playerinfo{}
// // 					}
// // 					if _, ok := storeInfo[_ID]; ok {
// // 						storeInfo[_ID][roomname][_TeamRedKey] = Playerinfo{isBlock: false}
// // 					}

// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: Unblock, to: _ID, _sleep: false, roomname: roomname,
// // 					}
// // 				}
// // 			}
// // 			saveBlock <- storeInfo

// // 		}
// // 	default:
// // 		log.Println("error in blocking")
// // 	}

// // }

// // // PadLock returns  these will be going against each other
// // // note: to send the save shuffle here
// // func PadLock(h *Hub, Set int, Round int, isTeamBlue bool, roomname string, shuffle map[string]map[string][]string) {
// // 	// lock the untagged players
// // 	log.Println("in pad lock")
// // 	log.Println("set: ", Set, "round: ", Round)
// // 	log.Println("room name: ", roomname)

// // 	// lo := Playerinfo{isLock: true, isBlock: false}
// // 	// unlo := Playerinfo{isLock: false, isBlock: false}
// // 	// if getSettings[roomname][_TwoVTwoKey] {

// // 	// 	go func() {

// // 	// 		MatchUp := map[int]LockPatterns{
// // 	// 			1: {

// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][1],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][1],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			2: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][0],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][0],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			3: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][1],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][0],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			4: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][0],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][1],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			5: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][1],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][1],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			6: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][0],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][0],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 		}

// // 	// 		NotMatchUp := map[int]LockPatterns{
// // 	// 			1: {

// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][1],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][1],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			2: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][0],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][0],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			3: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][1],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][0],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			4: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][0],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][1],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			5: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][1],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][1],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 			6: {
// // 	// 				TeamRed: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamRedKey][0],
// // 	// 					Team: _TeamRedKey},
// // 	// 				TeamBlue: LockPlayerInfo{
// // 	// 					ID:   shuffle[roomname][_TeamBlueKey][0],
// // 	// 					Team: _TeamBlueKey},
// // 	// 			},
// // 	// 		}

// // 	// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, MatchUp[Round].TeamRed.ID, false}
// // 	// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, MatchUp[Round].TeamBlue.ID, false}
// // 	// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, NotMatchUp[Round].TeamBlue.ID, false}
// // 	// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, NotMatchUp[Round].TeamRed.ID, false}
// // 	// 		store := map[string]map[string]map[string]Playerinfo{}
// // 	// 		store[MatchUp[Round].TeamBlue.ID] = map[string]map[string]Playerinfo{
// // 	// 			roomname: {
// // 	// 				MatchUp[Round].TeamBlue.Team: lo,
// // 	// 			},
// // 	// 		}
// // 	// 		store[MatchUp[Round].TeamRed.ID] = map[string]map[string]Playerinfo{
// // 	// 			roomname: {
// // 	// 				MatchUp[Round].TeamRed.Team: lo,
// // 	// 			},
// // 	// 		}
// // 	// 		store[NotMatchUp[Round].TeamBlue.ID] = map[string]map[string]Playerinfo{
// // 	// 			roomname: {
// // 	// 				NotMatchUp[Round].TeamBlue.Team: unlo,
// // 	// 			},
// // 	// 		}
// // 	// 		store[NotMatchUp[Round].TeamRed.ID] = map[string]map[string]Playerinfo{
// // 	// 			roomname: {
// // 	// 				NotMatchUp[Round].TeamRed.Team: unlo,
// // 	// 			},
// // 	// 		}

// // 	// 		if Round == 1 {
// // 	// 			BlockManager(h, 0, 1, isTeamBlue, roomname, shuffle)
// // 	// 			storePlayerInfo <- store
// // 	// 		} else {
// // 	// 			BlockManager(h, 0, -1, isTeamBlue, roomname, shuffle)
// // 	// 			saveLock <- store
// // 	// 		}

// // 	// 		// switch Set {

// // 	// 		// // round1: 0v0
// // 	// 		// // round2: 1v1
// // 	// 		// case 1:
// // 	// 		// 	log.Println("in set 1")
// // 	// 		// 	switch Round {
// // 	// 		// 	case 1:
// // 	// 		// 		log.Println("a0 vs b0")
// // 	// 		// 		log.Println("round 1")
// // 	// 		// 		// 0 v 0
// // 	// 		// 		fromA := _tempShuffle[roomname][_TeamRedKey][0]
// // 	// 		// 		fromB := _tempShuffle[roomname][_TeamBlueKey][0]

// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromB, false}

// // 	// 		// 		log.Println("locked: ", _tempShuffle[roomname][_TeamRedKey][0], _tempShuffle[roomname][_TeamBlueKey][0])

// // 	// 		// 		fromA = _tempShuffle[roomname][_TeamRedKey][1]
// // 	// 		// 		fromB = _tempShuffle[roomname][_TeamBlueKey][1]

// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromB, false}

// // 	// 		// 		// tracking the locking player
// // 	// 		// 		locked := map[string]map[string]map[string]bool{}

// // 	// 		// 		locked[_TeamRedKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {_tempShuffle[roomname][_TeamRedKey][0]: true,
// // 	// 		// 				_tempShuffle[roomname][_TeamRedKey][1]: false,
// // 	// 		// 			}}

// // 	// 		// 		locked[_TeamBlueKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {_tempShuffle[roomname][_TeamBlueKey][0]: true,
// // 	// 		// 				_tempShuffle[roomname][_TeamBlueKey][1]: false,
// // 	// 		// 			}}

// // 	// 		// 		// store the current set and round here
// // 	// 		// 		go func() {
// // 	// 		// 			store := map[string][]int{}
// // 	// 		// 			store[roomname] = []int{1}
// // 	// 		// 			TtrackRound.store <- store
// // 	// 		// 		}()

// // 	// 		// 		TLock.isLock <- locked

// // 	// 		// 		BlockManager(h, 0, 1, isTeamBlue, roomname)

// // 	// 		// 	// no need for second round cause default is round 1 for set 1
// // 	// 		// 	case 2:
// // 	// 		// 		log.Println("a1 vs b1")
// // 	// 		// 		log.Println("round 2")
// // 	// 		// 		// 1 v 1
// // 	// 		// 		fromA := shuffle[roomname][_TeamRedKey][1]
// // 	// 		// 		fromB := shuffle[roomname][_TeamBlueKey][1]

// // 	// 		// 		// players that are going to face each other
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromB, false}

// // 	// 		// 		fromA = shuffle[roomname][_TeamRedKey][0]
// // 	// 		// 		fromB = shuffle[roomname][_TeamBlueKey][0]
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromB, false}

// // 	// 		// 		log.Println("locked: ", fromA, fromB)

// // 	// 		// 		// tracking the locking player
// // 	// 		// 		locked := map[string]map[string]map[string]bool{}

// // 	// 		// 		locked[_TeamRedKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamRedKey][0]: false,
// // 	// 		// 				shuffle[roomname][_TeamRedKey][1]: true,
// // 	// 		// 			}}

// // 	// 		// 		locked[_TeamBlueKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamBlueKey][0]: false,
// // 	// 		// 				shuffle[roomname][_TeamBlueKey][1]: true,
// // 	// 		// 			}}
// // 	// 		// 		// note: we are storing set after each round completed
// // 	// 		// 		go func() {
// // 	// 		// 			store := map[string][]int{}
// // 	// 		// 			store2 := map[string][]int{}
// // 	// 		// 			store[roomname] = []int{1}
// // 	// 		// 			store2[roomname] = []int{1}

// // 	// 		// 			TtrackRound.store <- store
// // 	// 		// 			TtrackSet.store <- store2
// // 	// 		// 		}()
// // 	// 		// 		TLock.isLock <- locked
// // 	// 		// 		BlockManager(h, 0, -1, isTeamBlue, roomname)

// // 	// 		// 	default:
// // 	// 		// 		log.Println("unable to decIDe round")
// // 	// 		// 	}

// // 	// 		// 	// round1:1v0
// // 	// 		// 	// round2: 0v1
// // 	// 		// case 2:
// // 	// 		// 	log.Println("set 2")
// // 	// 		// 	switch Round {
// // 	// 		// 	case 3:
// // 	// 		// 		log.Println("a1 vs b0")
// // 	// 		// 		log.Println("round 3")
// // 	// 		// 		fromA := shuffle[roomname][_TeamRedKey][1]
// // 	// 		// 		fromB := shuffle[roomname][_TeamBlueKey][0]

// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromB, false}

// // 	// 		// 		log.Println("locked: ", shuffle[roomname][_TeamRedKey][1], shuffle[roomname][_TeamBlueKey][0])

// // 	// 		// 		fromA = shuffle[roomname][_TeamRedKey][0]
// // 	// 		// 		fromB = shuffle[roomname][_TeamBlueKey][1]

// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromB, false}

// // 	// 		// 		// tracking the locking player
// // 	// 		// 		locked := map[string]map[string]map[string]bool{}

// // 	// 		// 		locked[_TeamRedKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamRedKey][0]: false,
// // 	// 		// 				shuffle[roomname][_TeamRedKey][1]: true,
// // 	// 		// 			}}

// // 	// 		// 		locked[_TeamBlueKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamBlueKey][0]: true,
// // 	// 		// 				shuffle[roomname][_TeamBlueKey][1]: false,
// // 	// 		// 			}}

// // 	// 		// 		go func() {
// // 	// 		// 			store := map[string][]int{}
// // 	// 		// 			store[roomname] = []int{1}

// // 	// 		// 			TtrackRound.store <- store

// // 	// 		// 		}()
// // 	// 		// 		TLock.isLock <- locked
// // 	// 		// 		BlockManager(h, 0, -1, isTeamBlue, roomname)

// // 	// 		// 	case 4:
// // 	// 		// 		log.Println("a0 vs b1")
// // 	// 		// 		log.Println("round 4")
// // 	// 		// 		// a0 vs b1
// // 	// 		// 		fromA := shuffle[roomname][_TeamRedKey][0]
// // 	// 		// 		fromB := shuffle[roomname][_TeamBlueKey][1]

// // 	// 		// 		// players that are going to face each other
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromB, false}

// // 	// 		// 		fromA = shuffle[roomname][_TeamRedKey][1]
// // 	// 		// 		fromB = shuffle[roomname][_TeamBlueKey][0]
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromB, false}

// // 	// 		// 		log.Println("locked: ", fromA, fromB)

// // 	// 		// 		// tracking the locking player
// // 	// 		// 		locked := map[string]map[string]map[string]bool{}

// // 	// 		// 		locked[_TeamRedKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamRedKey][0]: true,
// // 	// 		// 				shuffle[roomname][_TeamRedKey][1]: false,
// // 	// 		// 			}}

// // 	// 		// 		locked[_TeamBlueKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamBlueKey][0]: false,
// // 	// 		// 				shuffle[roomname][_TeamBlueKey][1]: true,
// // 	// 		// 			}}

// // 	// 		// 		go func() {
// // 	// 		// 			store := map[string][]int{}
// // 	// 		// 			store2 := map[string][]int{}
// // 	// 		// 			store[roomname] = []int{1}
// // 	// 		// 			store2[roomname] = []int{1}

// // 	// 		// 			TtrackRound.store <- store
// // 	// 		// 			TtrackSet.store <- store2
// // 	// 		// 		}()
// // 	// 		// 		TLock.isLock <- locked
// // 	// 		// 		BlockManager(h, 0, -1, isTeamBlue, roomname)

// // 	// 		// 	default:
// // 	// 		// 		log.Println("unable to decIDe round")
// // 	// 		// 	}

// // 	// 		// 	// round1: 1v1
// // 	// 		// 	// round4: 0v0
// // 	// 		// case 3:
// // 	// 		// 	log.Println("set 3")
// // 	// 		// 	switch Round {
// // 	// 		// 	case 5:
// // 	// 		// 		log.Println("round 5")
// // 	// 		// 		log.Println("a1 vs b1")
// // 	// 		// 		// 1 v 1
// // 	// 		// 		fromA := shuffle[roomname][_TeamRedKey][1]
// // 	// 		// 		fromB := shuffle[roomname][_TeamBlueKey][1]

// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromB, false}

// // 	// 		// 		log.Println("locked: ", shuffle[roomname][_TeamRedKey][1], shuffle[roomname][_TeamBlueKey][1])

// // 	// 		// 		fromA = shuffle[roomname][_TeamRedKey][0]
// // 	// 		// 		fromB = shuffle[roomname][_TeamBlueKey][0]

// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromB, false}

// // 	// 		// 		// tracking the locking player
// // 	// 		// 		locked := map[string]map[string]map[string]bool{}

// // 	// 		// 		locked[_TeamRedKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamRedKey][0]: false,
// // 	// 		// 				shuffle[roomname][_TeamRedKey][1]: true,
// // 	// 		// 			}}

// // 	// 		// 		locked[_TeamBlueKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamBlueKey][0]: false,
// // 	// 		// 				shuffle[roomname][_TeamBlueKey][1]: true,
// // 	// 		// 			}}

// // 	// 		// 		go func() {
// // 	// 		// 			store := map[string][]int{}
// // 	// 		// 			store[roomname] = []int{1}

// // 	// 		// 			TtrackRound.store <- store
// // 	// 		// 		}()
// // 	// 		// 		TLock.isLock <- locked
// // 	// 		// 		BlockManager(h, 0, -1, isTeamBlue, roomname)

// // 	// 		// 	case 6:
// // 	// 		// 		log.Println("a0 vs b0")
// // 	// 		// 		log.Println("round 6")
// // 	// 		// 		// 0 v 0
// // 	// 		// 		fromA := shuffle[roomname][_TeamRedKey][0]
// // 	// 		// 		fromB := shuffle[roomname][_TeamBlueKey][0]

// // 	// 		// 		// players that are going to face each other
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Lock, fromB, false}

// // 	// 		// 		fromA = shuffle[roomname][_TeamRedKey][1]
// // 	// 		// 		fromB = shuffle[roomname][_TeamBlueKey][1]
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromA, false}
// // 	// 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname, Unlock, fromB, false}

// // 	// 		// 		log.Println("locked: ", fromA, fromB)

// // 	// 		// 		// tracking the locking player
// // 	// 		// 		locked := map[string]map[string]map[string]bool{}

// // 	// 		// 		locked[_TeamRedKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamRedKey][0]: true,
// // 	// 		// 				shuffle[roomname][_TeamRedKey][1]: false,
// // 	// 		// 			}}

// // 	// 		// 		locked[_TeamBlueKey] = map[string]map[string]bool{
// // 	// 		// 			roomname: {shuffle[roomname][_TeamBlueKey][0]: true,
// // 	// 		// 				shuffle[roomname][_TeamBlueKey][1]: false,
// // 	// 		// 			}}
// // 	// 		// 		TLock.isLock <- locked
// // 	// 		// 		BlockManager(h, 0, -1, isTeamBlue, roomname)

// // 	// 		// 		go func() {
// // 	// 		// 			store := map[string][]int{}
// // 	// 		// 			store2 := map[string][]int{}
// // 	// 		// 			store[roomname] = []int{1}
// // 	// 		// 			store2[roomname] = []int{1}

// // 	// 		// 			TtrackRound.store <- store
// // 	// 		// 			TtrackSet.store <- store2
// // 	// 		// 		}()
// // 	// 		// 	default:
// // 	// 		// 		log.Println("unable to decIDe round")
// // 	// 		// 	}

// // 	// 		// default:
// // 	// 		// 	log.Println("Game over or set =", len(getTtrackSet))
// // 	// 		// }
// // 	// 	}()
// // 	// } else {

// // 	// 	go func() {
// // 	// 		log.Println("1v1")
// // 	// 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: Lock}
// // 	// 		if Round == 1 {
// // 	// 			log.Println("round 1 lock")

// // 	// 			// locked := map[string]map[string]map[string]bool{}

// // 	// 			// // todo: make sure to clean the temps full beuase it will be used only one time
// // 	// 			// locked[_TeamRedKey] = map[string]map[string]bool{
// // 	// 			// 	roomname: {_tempShuffle[roomname][_TeamRedKey][0]: true}}

// // 	// 			// locked[_TeamBlueKey] = map[string]map[string]bool{
// // 	// 			// 	roomname: {_tempShuffle[roomname][_TeamBlueKey][0]: true}}

// // 	// 			// store the current set and round here
// // 	// 			go func() {
// // 	// 				store := map[string][]int{}
// // 	// 				store[roomname] = []int{1}
// // 	// 				TtrackRound.store <- store
// // 	// 			}()
// // 	// 			store := map[string]map[string]map[string]Playerinfo{}
// // 	// 			store[shuffle[roomname][_TeamRedKey][0]] = map[string]map[string]Playerinfo{
// // 	// 				roomname: {_TeamRedKey: lo},
// // 	// 			}
// // 	// 			store[shuffle[roomname][_TeamBlueKey][0]] = map[string]map[string]Playerinfo{
// // 	// 				roomname: {_TeamBlueKey: lo},
// // 	// 			}
// // 	// 			storePlayerInfo <- store

// // 	// 			BlockManager(h, 0, 1, isTeamBlue, roomname, shuffle)
// // 	// 		} else {
// // 	// 			log.Println("round 2 lock")

// // 	// 			// locked := map[string]map[string]map[string]bool{}

// // 	// 			// locked[_TeamRedKey] = map[string]map[string]bool{
// // 	// 			// 	roomname: {_tempShuffle[roomname][_TeamRedKey][0]: true}}

// // 	// 			// locked[_TeamBlueKey] = map[string]map[string]bool{
// // 	// 			// 	roomname: {_tempShuffle[roomname][_TeamBlueKey][0]: true}}

// // 	// 			switch Round {
// // 	// 			case 1:
// // 	// 				// store the current set and round here
// // 	// 				go func() {
// // 	// 					store := map[string][]int{}
// // 	// 					store[roomname] = []int{1}

// // 	// 					TtrackRound.store <- store
// // 	// 				}()
// // 	// 			case 2:
// // 	// 				go func() {
// // 	// 					store := map[string][]int{}
// // 	// 					store2 := map[string][]int{}
// // 	// 					store[roomname] = []int{1}
// // 	// 					store2[roomname] = []int{1}

// // 	// 					TtrackRound.store <- store
// // 	// 					TtrackSet.store <- store2
// // 	// 				}()
// // 	// 			case 3:
// // 	// 				// store the current set and round here
// // 	// 				go func() {
// // 	// 					store := map[string][]int{}
// // 	// 					store[roomname] = []int{1}

// // 	// 					TtrackRound.store <- store
// // 	// 				}()
// // 	// 			case 4:
// // 	// 				go func() {
// // 	// 					store := map[string][]int{}
// // 	// 					store2 := map[string][]int{}
// // 	// 					store[roomname] = []int{1}
// // 	// 					store2[roomname] = []int{1}

// // 	// 					TtrackRound.store <- store
// // 	// 					TtrackSet.store <- store2
// // 	// 				}()
// // 	// 			case 5:
// // 	// 				// store the current set and round here
// // 	// 				go func() {
// // 	// 					store := map[string][]int{}
// // 	// 					store[roomname] = []int{1}
// // 	// 					TtrackRound.store <- store
// // 	// 				}()
// // 	// 			case 6:
// // 	// 				go func() {
// // 	// 					store := map[string][]int{}
// // 	// 					store2 := map[string][]int{}
// // 	// 					store[roomname] = []int{1}
// // 	// 					store2[roomname] = []int{1}

// // 	// 					TtrackRound.store <- store
// // 	// 					TtrackSet.store <- store2
// // 	// 				}()
// // 	// 			}
// // 	// 			store := map[string]map[string]map[string]Playerinfo{}
// // 	// 			store[shuffle[roomname][_TeamRedKey][0]] = map[string]map[string]Playerinfo{
// // 	// 				roomname: {_TeamRedKey: lo},
// // 	// 			}
// // 	// 			store[shuffle[roomname][_TeamBlueKey][0]] = map[string]map[string]Playerinfo{
// // 	// 				roomname: {_TeamBlueKey: lo},
// // 	// 			}

// // 	// 			// TLock.isLock <- locked
// // 	// 			saveLock <- store
// // 	// 			BlockManager(h, 0, 1, isTeamBlue, roomname, shuffle)
// // 	// 			// TLock.isLock <- locked
// // 	// 			go func() {
// // 	// 				BlockManager(h, 0, -1, isTeamBlue, roomname, shuffle)
// // 	// 			}()
// // 	// 		}
// // 	// 	}()
// // 	// }

// // 	log.Println("lock done")

// // }

// // func LocifyOneVOne(h *Hub, ids []string, roomname string) {
// // 	draft := ids

// // 	rand.Shuffle(len(draft), func(i int, j int) {
// // 		draft[i], draft[j] = draft[j], draft[i]
// // 	})

// // 	teamRed := []string{draft[0]}
// // 	teamBlue := []string{draft[1]}
// // 	Nicknames := profiles.Fetch(ids)
// // 	draftPowers := []string{}
// // 	for power, _include := range getLocifyRoomSettings[roomname].Powers {
// // 		if _include {
// // 			draftPowers = append(draftPowers, power)
// // 		}
// // 	}

// // 	rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 		draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 	})

// // 	teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 	teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 	for _, powers := range teamRedPowers {
// // 		for _, IDs := range teamRed {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 				_sleep: false, to: IDs, roomname: roomname}
// // 		}
// // 	}

// // 	for _, powers := range teamBluePowers {
// // 		for _, IDs := range teamBlue {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 				_sleep: false, to: IDs, roomname: roomname}
// // 		}
// // 	}
// // 	log.Println("team red: ", teamRed)
// // 	log.Println("team blue: ", teamBlue)
// // 	log.Println("team red powers", teamRedPowers)
// // 	log.Println("team blue powers", teamBluePowers)

// // 	if len(teamRed) > 0 && len(teamBlue) > 0 {
// // 		log.Println("teamRed: ", teamRed[0])
// // 		log.Println("teamBlue: ", teamBlue[0])
// // 		// broadcasting nicknames of teams and broadcasting their team name
// // 		log.Println("sending det")

// // 		TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			for _, IDs := range teamBlue {
// // 				h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 			}
// // 			for _, IDs := range teamRed {
// // 				h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 			}
// // 			for _, IDs := range teamRed {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 					_sleep: false, roomname: roomname}
// // 			}
// // 			for _, IDs := range teamBlue {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 					_sleep: false, roomname: roomname}
// // 			}
// // 			for _, _ID := range Nicknames.NickNames {
// // 				h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 			}
// // 		}()
// // 		log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 		SendRoomTime(h, roomname)
// // 		X := strconv.Itoa(len(h.rooms[roomname]))
// // 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + X}
// // 		roomMode := getRoomList[roomname].Type
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}

// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			Conns = strconv.Itoa(len(h.rooms[roomname]))
// // 			h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}
// // 			h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 		}()

// // 		nicknames := []string{}

// // 		// creating sheet
// // 		for _, id := range teamRed {
// // 			nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 		}
// // 		for _, id := range teamBlue {
// // 			nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 		}

// // 		createSheet := PackSheet{Name: make([]string, 4), Sheet: make(map[string]map[string]string)}
// // 		createSheet.Name = append(createSheet.Name, nicknames...)
// // 		log.Println("nicknames: ", createSheet.Name)

// // 		for _, names := range nicknames {
// // 			createSheet.Sheet[names] = map[string]string{
// // 				"round1": "",
// // 			}
// // 		}

// // 		log.Println("nickanems: ", getNicknamesViaID)

// // 		pack, _ := json.Marshal(&createSheet)

// // 		h.broadcast <- BroadcastReq{Token: "UpdateCheatSheet: " + string(pack), RoomID: roomname}

// // 		log.Println("created sheet: ", string(pack))

// // 		go func() {
// // 			store := map[string]PackSheet{}
// // 			store[roomname] = createSheet

// // 			storeSheetUpdate.store <- store
// // 		}()

// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			bookName := getRoomRequestedToken[roomname][_BookKey]
// // 			sendDic := SetupDictionaryURL + SendDictionary(bookName)

// // 			h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}
// // 		}()
// // 		h.broadcast <- BroadcastReq{Token: _Clash, RoomID: roomname}

// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			// meaning for next one set 1 round 2
// // 			store, store2 := map[string][]int{}, map[string][]int{}
// // 			store[l.RoomName] = []int{1}
// // 			store2[l.RoomName] = []int{1}

// // 			TtrackSet.store <- store2
// // 			TtrackRound.store <- store
// // 		}()
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			store3 := map[string]map[string][]string{}
// // 			store3[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 			storeShuffle.store <- store3
// // 		}()
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			latest = true
// // 			storePublicRoom.done <- true
// // 		}()
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			store := map[string]map[string][]string{}
// // 			store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 			PadLock(h, 1, 1, false, roomname, store)
// // 		}()
// // 	} else {
// // 		log.Println("problem in sending")
// // 	}
// // }

// //cap := FriendRoomCap(l) - _CountFriend // signal that one of the friend has joined

// // friendCount := _CountFriend // 1 || 0

// // IsEntertainment := false

// // if l.Category == "entertainment" {
// // 	IsEntertainment = true
// // }
// // delete(h.clients, connection)
// // delete(storeRoom, client.roomID)
// // }

// // case conn := <-h.unregisterSession:

// // 		UnregisterSession(h, _conn, )
// // 		delete(h.clients, _conn)

// // case ID := <-roomStoreClientID.clientID:
// // 	cIDstorage = append(cIDstorage, ID)

// // case done := <-roomStoreClientID.done:
// // 	if done {
// // 		cIDstorage = nil
// // 		if _, ok := <-roomStoreClientID.clientID; !ok {
// // 			close(roomStoreClientID.clientID)
// // 		}
// // 	}
// // case token := <-storeRoomList.store:
// // 	maps.Copy(getRoomList, token)
// // 	log.Println("STORED ROOM LIST: ", getRoomList)

// // case done := <-storeRoomList.done:
// // 	if done {
// // 		for r := range getRoomList {
// // 			delete(getRoomList, r)
// // 		}
// // 	}
// // broadcast to client's in the room
// // stores created room
// // case room := <-roomStore.store:
// // 	for r := range room {
// // 		storeRoom[r] = room[r]
// // 	}
// // 	// append the rooms
// // 	for r := range storeRoom {
// // 		rooms = append(rooms, r)
// // 	}
// // 	rooms = dataset.EraseDuplicate(rooms)

// // 	log.Println("STORED:", storeRoom)
// // case tokens := <-storeValidateToken.store:
// // 	for name, token := range tokens {
// // 		for _, _tok := range token {
// // 			if _, exists := getPreviousValidateToken[name]; !exists {
// // 				getPreviousValidateToken[name] = []string{}
// // 			}
// // 			getPreviousValidateToken[name] = append(getPreviousValidateToken[name], _tok)
// // 		}
// // 	}
// // signals session whether started or ended
// // case session := <-gameSession.going:
// // 	gameSessionStarted = session
// // 	for r := range session {
// // 		gameSessionStarted[r] = session[r]
// // 	}

// // case done := <-gameSession.done:
// // 	if done {
// // 		_, ok := <-gameSession.going
// // 		if !ok {
// // 			gameSession.going = nil
// // 			close(gameSession.done)
// // 			close(gameSession.going)
// // 		} else {
// // 			log.Println("game session running")
// // 		}
// // 	}
// // case lock := <-TLock.isLock:
// // 	maps.Copy(isLock, lock)
// // 	log.Println("lock view: ", lock)

// // case block := <-TBlock:
// // 	maps.Copy(isBlock, block)
// // 	log.Println("block view: ", block)

// // case toss := <-Ttoss.save:
// // 	maps.Copy(getToss, toss)

// // case _close := <-Ttoss.done:
// // 	log.Println("close: ", _close)
// // 	if _close {
// // 		// close(Ttoss.saveToss)
// // 		for r := range getToss {
// // 			delete(getToss, r)
// // 		}
// // 	} else {
// // 		log.Println("not able to close the toss channel")
// // 	}

// // case client := <-ProfData.data:
// // 	for r := range client {
// // 		getProfData[r] = client[r]
// // 	}

// // case owner := <-SaveOwnerProf:
// // 	for r := range owner {
// // 		maps.Copy(getOwnerProf, owner[r])
// // 	}

// // case challenge := <-TChallengeSet.save:
// // 	maps.Copy(getChallengeSet, challenge)

// // case _close := <-TChallengeSet.done:
// // 	log.Println("close: ", _close)
// // 	if _close {
// // 		// close(Ttoss.saveToss)
// // 		for r := range getChallengeSet {
// // 			delete(getChallengeSet, r)
// // 		}
// // 	} else {
// // 		log.Println("not able to close the toss channel")
// // 	}

// // case visited := <-Mwatch.done:
// // 	maps.Copy(getMWatch, visited)

// // case _close := <-Mwatch.complete:
// // 	if _close {
// // 		log.Println("closing")
// // 		for r := range getMWatch {
// // 			log.Println("deleting get watch")
// // 			delete(getMWatch, r)
// // 		}
// // 	} else {
// // 		log.Println("not able to close the watch channel")
// // 	}

// // case visited := <-Cwatch.done:
// // maps.Copy(getCWatch, visited)

// // case _close := <-Cwatch.complete:
// // 	if _close {
// // 		log.Println("closing")
// // 		for r := range getCWatch {
// // 			log.Println("deleting get watch")
// // 			delete(getCWatch, r)
// // 		}
// // 	} else {
// // 		log.Println("not able to close the watch channel")
// // 	}

// // case visited := <-Twatch.done:
// // 	maps.Copy(getTWatch, visited)

// // case _close := <-Twatch.complete:
// // 	if _close {
// // 		log.Println("closing")
// // 		for r := range getTWatch {
// // 			log.Println("deleting get watch")
// // 			delete(getTWatch, r)
// // 		}
// // 	} else {
// // 		log.Println("not able to close the watch channel")
// // 	}

// // case store := <-storePenalty.store:
// // 	for ids := range store {
// // 		for room := range store[ids] {
// // 			for teamname := range store[ids][room] {
// // 				for count := range store[ids][room][teamname] {
// // 					getStorePenalty[ids][room][teamname] = append(getStorePenalty[ids][room][teamname], count)
// // 				}

// // 			}
// // 		}
// // 	}

// // case store := <-TtrackRound.store:
// // 	for name := range store {

// // 		for _, r := range store[name] {
// // 			if _, ok := getTtrackRound[name]; !ok {
// // 				getTtrackRound[name] = r
// // 			} else {
// // 				getTtrackRound[name] += r
// // 			}
// // 		}
// // 	}
// // 	log.Println("STORED ROUND: ", store)
// // 	log.Println("getTrackround: ", getTtrackRound)
// // case done := <-TtrackRound.done:
// // 	if done {
// // 		for name := range getTtrackRound {
// // 			delete(getTtrackRound, name)
// // 		}
// // 	}

// // case store := <-TtrackSet.store:
// // 	for name := range store {
// // 		for _, r := range store[name] {
// // 			if _, ok := getTtrackSet[name]; !ok {
// // 				getTtrackSet[name] = r
// // 			} else {
// // 				getTtrackSet[name] += r
// // 			}
// // 		}
// // 	}
// // 	log.Println("STORED SET: ", store)
// // 	log.Println("get track set: ", getTtrackSet)

// // case done := <-TtrackSet.done:
// // 	if done {
// // 		for name := range getTtrackSet {
// // 			delete(getTtrackSet, name)
// // 		}
// // 	}

// // case store := <-storeCups.store:
// // 	maps.Copy(getStoredCups, store)
// // 	log.Println("STOING CURP: ", store)
// // 	log.Println("STORED CUPS: ", getStoredCups)

// // case store := <-storeDict.store:
// // 	maps.Copy(getStoredDict, store)

// // case done := <-storeDict.done:
// // 	if done {
// // 		for r := range getStoredDict {
// // 			delete(getStoredDict, r)
// // 		}
// // 	}

// // case done := <-storeCups.done:
// // 	if done {
// // 		for r := range getStoredCups {
// // 			delete(getStoredCups, r)
// // 		}
// // 	}

// // case store := <-storeSessionDone.store:
// // 	maps.Copy(getSessionDone, store)

// // case done := <-storeSessionDone.done:
// // 	if done {
// // 		for r := range getSessionDone {
// // 			delete(getSessionDone, r)
// // 		}
// // 	}

// // case store := <-storeSheetUpdate.store:

// // 	FUpdateSheet(store, getStoredSheetUpdate)

// // case get := <-settings.store:
// // 	maps.Copy(getSettings, get)
// // 	log.Println("settings saved: ", getSettings)

// // case token := <-TTrackSessionDone:
// // 	maps.Copy(getTTrackSessionDone, token)

// // case token := <-saveSession:
// // 	maps.Copy(getsaveSession, token)
// // case token := <-saveAgreement:
// // 	maps.Copy(getSaveAgreement, token)

// // // case token := <-Dwatch.done:
// // // 	maps.Copy(getDWatch, token)
// // // case done := <-Dwatch.complete:
// // // 	if done {
// // // 		for r := range getDWatch {
// // // 			delete(getDWatch, r)
// // // 		}
// // // 	}
// // case token := <-saveChatMes:
// // 	getChatMes[""] = map[string]map[string][]string{"": {"": {""}}}
// // 	for ID, innerRoomMap := range token {
// // 		for roomname, roomnameDet := range innerRoomMap {
// // 			for TeamName, Msgs := range roomnameDet {
// // 				getChatMes[ID][roomname][TeamName] = append(getChatMes[ID][roomname][TeamName], Msgs...)
// // 			}
// // 		}
// // 	}
// // 	maps.Copy(getChatMes, token)

// // case power := <-storePower.store:
// // 	maps.Copy(getStoredPower, power)

// // case done := <-storePower.done:
// // 	if done {
// // 		log.Println("closing")
// // 		for r := range getStoredPower {
// // 			delete(getStoredPower, r)
// // 		}
// // 	} else {
// // 		log.Println("not able to close the watch channel")
// // 	}

// // case token := <-saveSessionDone.store:
// // 	maps.Copy(isSessionDone, token)
// // case done := <-saveSessionDone.done:
// // 	if done {
// // 		for r := range isSessionDone {
// // 			delete(isSessionDone, r)
// // 		}
// // 	}
// // case token := <-starterSettings.store:
// // 	maps.Copy(getStarterSetting, token)

// // case token := <-storeRoomOwner.store:
// // 	maps.Copy(getRoomOwner, token)

// // case code := <-storeFriendRoom.roomCode:
// // 	maps.Copy(getFriendRoomCode, code)

// // case room := <-storeFriendRoom.store:
// // 	maps.Copy(getFriendRoom, room)

// // case eraseCode := <-storeFriendRoom.removeCode:
// // 	if eraseCode {
// // 		for r := range getFriendRoomCode {
// // 			delete(getFriendRoomCode, r)
// // 		}
// // 	}
// // case token := <-PlayerJoined:
// // 	maps.Copy(hasPlayerJoined, token)
// // case token := <-saveJoinFriendID:
// // 	maps.Copy(getSavedJoinFriendID, token)
// // case token := <-saveOwnerFriendID:
// // 	maps.Copy(getSavedOwnerFriendID, token)
// // case done := <-storeFriendRoom.done:
// // 	if done {
// // 		for r := range getFriendRoom {
// // 			delete(getFriendRoom, r)
// // 		}
// // 	}

// // case token := <-storePrivateRoom.store:
// // 	maps.Copy(getPrivateRoom, token)

// // case code := <-storePrivateRoom.roomCode:
// // 	maps.Copy(getPrivateRoomCode, code)
// // 	log.Println("STORED PRIVATE CODE: ", code)

// // case done := <-storePrivateRoom.removeCode:
// // 	if done {
// // 		for r := range getPrivateRoomCode {
// // 			delete(getPrivateRoomCode, r)
// // 		}
// // 	}
// // case done := <-storePrivateRoom.done:
// // 	if done {
// // 		for r := range getPrivateRoom {
// // 			delete(getPrivateRoom, r)
// // 		}
// // 	}

// // case token := <-storePublicRoom.store:
// // 	maps.Copy(getPublicRoom, token)

// // case done := <-storePublicRoom.done:
// // 	if done {
// // 		for r := range getPublicRoom {
// // 			delete(getPublicRoom, r)
// // 		}
// // 	}
// // case token := <-storeRoomRequestedToken.store:
// // 	maps.Copy(getRoomRequestedToken, token)
// // 	log.Println("STORED: ", token)

// // case done := <-storeRoomRequestedToken.done:
// // 	if done {
// // 		log.Println("DELETING store room requet ")

// // 		for r := range getRoomRequestedToken {
// // 			delete(getRoomRequestedToken, r)
// // 		}
// // 	}
// // case token := <-storeRewindPower:
// // 	maps.Copy(getRewindPower, token)

// // case names := <-storeNickNames.viaID:
// // 	maps.Copy(getNicknamesViaID, names)

// // case names := <-storeNickNames.viaRoom:
// // 	for rooms, names := range names {
// // 		getNicknamesViaRoom[rooms] = append(getNicknamesViaRoom[rooms], names)
// // 	}

// // case token := <-storeBetValue.store:
// // 	maps.Copy(getStoredBetValue, token)

// // case done := <-storeBetValue.done:
// // 	if done {
// // 		for r := range getStoredBetValue {
// // 			delete(getStoredBetEvent, r)
// // 		}
// // 	}

// // case token := <-storeDictionaryEvent.store:
// // 	maps.Copy(getStoredDictionaryEvent, token)

// // case done := <-storeDictionaryEvent.done:
// // 	if done {
// // 		for r := range getStoredDictionaryEvent {
// // 			delete(getStoredDictionaryEvent, r)
// // 		}
// // 	}

// // case token := <-storeBetEvent.store:
// // 	maps.Copy(getStoredBetEvent, token)

// // case done := <-storeBetEvent.done:
// // 	if done {
// // 		for r := range getStoredBetEvent {
// // 			delete(getStoredBetEvent, r)
// // 		}
// // 	}

// // case token := <-storeSessionUpdate:
// // 	maps.Copy(getSessionUpdate, token)
// // 	for id, roomnameTeamNameDetails := range token {
// // 		for roomname, TeamNameDetails := range roomnameTeamNameDetails {
// // 			for teamname, info := range TeamNameDetails {
// // 				if _, ok := getSessionUpdate[id]; !ok {
// // 					getSessionUpdate[id] = map[string]map[string]Session{}
// // 				}
// // 				if _, ok := getSessionUpdate[id][roomname]; !ok {
// // 					getSessionUpdate[id][roomname] = map[string]Session{}
// // 				}
// // 				if _, ok := getSessionUpdate[id][roomname][teamname]; !ok {
// // 					getSessionUpdate[id][roomname][teamname] = info
// // 				}

// // 			}
// // 		}
// // 	}
// // case token := <-resetSesssion:
// // 	for _roomname, reset := range token {
// // 		for id, roomsTeamsDet := range getSessionUpdate {
// // 			for roomname, teamsdet := range roomsTeamsDet {
// // 				for teamname := range teamsdet {
// // 					if _roomname == roomname {
// // 						if reset {
// // 							temp := getSessionUpdate[id][roomname][teamname]
// // 							temp.BetDone = false
// // 							temp.ChallengeDone = false
// // 							temp.DictionaryDone = false
// // 							temp.TossDone = false
// // 							getSessionUpdate[id][roomname][teamname] = temp
// // 						}
// // 					}
// // 				}
// // 			}
// // 		}
// // 	}
// // case token := <-storeTokensUpdate:
// // 	maps.Copy(getTokens, token)
// // case token := <-storePlayerInfo:
// // 	for id, roomnameTeamNameDetails := range token {
// // 		for roomname, TeamNameDetails := range roomnameTeamNameDetails {
// // 			for teamname, info := range TeamNameDetails {
// // 				if _, ok := getPlayerInfo[id]; !ok {
// // 					getPlayerInfo[id] = map[string]map[string]Playerinfo{}
// // 				}
// // 				if _, ok := getPlayerInfo[id][roomname]; !ok {
// // 					getPlayerInfo[id][roomname] = map[string]Playerinfo{}
// // 				}
// // 				if _, ok := getPlayerInfo[id][roomname][teamname]; !ok {
// // 					getPlayerInfo[id][roomname][teamname] = info
// // 				}

// // 			}
// // 		}
// // 	}

// // case token := <-storeGameInfo:
// // 	for roomname := range token {
// // 		if _, ok := getGameInfo[roomname]; !ok {
// // 			getGameInfo[roomname] = token[roomname]
// // 		}
// // 	}

// // case token := <-addDictionary:
// // 	for roomname, dictionary := range token {
// // 		if _, ok := getGameInfo[roomname]; ok {
// // 			temp := getGameInfo[roomname]
// // 			temp.SetDictionary = dictionary
// // 			getGameInfo[roomname] = temp
// // 		}
// // 	}
// // case token := <-addEvent:
// // 	for roomname, event := range token {
// // 		if _, ok := getGameInfo[roomname]; ok {
// // 			temp := getGameInfo[roomname]
// // 			temp.SetEvent = event
// // 			getGameInfo[roomname] = temp
// // 		}
// // 	}
// // case token := <-addTeamRedChallenge:
// // 	for roomname, challenge := range token {
// // 		if _, ok := getGameInfo[roomname]; ok {
// // 			temp := getGameInfo[roomname]
// // 			temp.TeamRedSetChallenge = challenge
// // 			getGameInfo[roomname] = temp
// // 		}
// // 	}
// // case token := <-addTeamBlueChallenge:
// // 	for roomname, challenge := range token {
// // 		if _, ok := getGameInfo[roomname]; ok {
// // 			temp := getGameInfo[roomname]
// // 			temp.TeamBlueSetChallenge = challenge
// // 			getGameInfo[roomname] = temp
// // 		}
// // 	}

// // case token := <-addToss:
// // 	for roomname, teams := range token {
// // 		for teamname := range teams {
// // 			if _, ok := getGameInfo[roomname]; ok {
// // 				temp := getGameInfo[roomname]
// // 				temp.SetTossBody[teamname] = token[roomname][teamname]
// // 				getGameInfo[roomname] = temp
// // 			}
// // 		}
// // 	}

// // case token := <-resetTokens:
// // 	for roomname, erase := range token {
// // 		if _, ok := getGameInfo[roomname]; ok {
// // 			if erase {
// // 				temp := getGameInfo[roomname]
// // 				temp2 := map[string]string{}
// // 				temp2[roomname] = ""
// // 				temp.SetDictionary = ""
// // 				temp.TeamBlueSetChallenge = ""
// // 				temp.TeamRedSetChallenge = ""
// // 				temp.SetTossBody = temp2
// // 				temp.SetEvent = ""

// // 				getGameInfo[roomname] = temp
// // 			}
// // 		}
// // 	}

// // case token := <-saveBlock:
// // 	for id, roomnameTeamNameDetails := range token {
// // 		for roomname, TeamNameDetails := range roomnameTeamNameDetails {
// // 			for teamname := range TeamNameDetails {
// // 				if view, ok := token[id][roomname][teamname]; ok {

// // 					_repl := getPlayerInfo[id][roomname][teamname]
// // 					_repl.isBlock = view.isBlock
// // 					getPlayerInfo[id][roomname][teamname] = _repl
// // 				}
// // 			}
// // 		}
// // 	}

// // case token := <-saveLock:
// // 	for id, roomnameTeamNameDetails := range token {
// // 		for roomname, TeamNameDetails := range roomnameTeamNameDetails {
// // 			for teamname := range TeamNameDetails {
// // 				if view, ok := token[id][roomname][teamname]; ok {
// // 					_repl := getPlayerInfo[id][roomname][teamname]
// // 					_repl.isBlock = view.isLock
// // 					getPlayerInfo[id][roomname][teamname] = _repl
// // 				}
// // 			}
// // 		}
// // 	}
// // draft := cRoom[roomname]
// // input:
// // consIDer capacity of 4:
// // do the swapping of h.clients[conn]
// // @TODO: do the token passing
// // count := len(h.rooms[lo.To]) // make the change if for 1v1
// // cap := -1
// // _REDTeam := map[string][]string{}
// // _BLUEteam := map[string][]string{}
// // ids := cRoom[roomname]

// // switch true {
// // case l.Code != "" && PrivateRoomExists("", l.Code, true):
// // 	isPrivate = true
// // 	roomname = GetPrivateRoom(l.Code)

// // case l.Code != "" && FriendRoomExists("", l.Code, true):
// // 	isFriend = true
// // 	joinCode = true
// // 	roomname = GetFriendRoom(l.Code)

// // case l.Code == "" && FriendRoomExists(roomname, "", false):
// // 	isFriend = true
// // default:
// // 	isPublic = true
// // }
// // log.Println("private cap: ", storeRoom[GetPrivateRoom(l.Code)])
// // log.Println("friend cap: ", storeRoom[GetFriendRoom(l.Code)])

// // func ChoosenPowers(roomname string) []string {
// // 	getPowers := []string{}
// // 	for powers, chosen := range getSettings[roomname] {
// // 		if chosen && powers != _TwoVTwoKey {
// // 			getPowers = append(getPowers, powers)
// // 		}
// // 	}

// // 	log.Println("powers got: ", getPowers)
// // 	return getPowers
// // }

// // rand.Shuffle(len(draft), func(i int, j int) {
// // 	draft[i], draft[j] = draft[j], draft[i]
// // })

// // teamRed := dataset.EraseAfter(draft, len(draft)/2)
// // teamBlue := dataset.EraseBefore(draft, len(draft)/2)

// // draftPowers := ChoosenPowers(roomname)

// // rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 	draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // })

// // teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // for _, powers := range teamRedPowers {
// // 	for _, IDs := range teamRed {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 			_sleep: false, to: IDs, roomname: roomname}
// // 	}
// // }

// // for _, powers := range teamBluePowers {
// // 	for _, IDs := range teamBlue {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 			_sleep: false, to: IDs, roomname: roomname}
// // 	}
// // }
// // if len(teamRed) == 2 && len(teamBlue) == 2 {

// // 	log.Println("teamRed: ", teamRed)
// // 	log.Println("teamBlue: ", teamBlue)
// // 	log.Println("lens of red and blue: ", len(teamRed), len(teamBlue))
// // 	log.Println("croom: ", cRoom[roomname])
// // 	if len(teamRed) == 2 && len(teamBlue) == 2 {
// // 		log.Println("teamRed: ", teamRed[0])
// // 		log.Println("teamBlue: ", teamBlue[0])
// // 		// broadcasting nicknames of teams and broadcasting their team name
// // 		go func() {
// // 			log.Println("sending det")

// // 			TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name

// // 			for _, IDs := range teamBlue {
// // 				h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 			}
// // 			for _, IDs := range teamRed {
// // 				h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 			}
// // 			draftPowers := ChoosenPowers(roomname)

// // 			rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 				draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 			})

// // 			teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 			teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 			for _, powers := range teamRedPowers {
// // 				for _, IDs := range teamRed {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 						_sleep: false, to: IDs, roomname: roomname}
// // 				}
// // 			}

// // 			for _, powers := range teamBluePowers {
// // 				for _, IDs := range teamBlue {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 						_sleep: false, to: IDs, roomname: roomname}
// // 				}
// // 			}
// // 			log.Println("team red powers", teamRedPowers)
// // 			log.Println("team blue powers", teamBluePowers)

// // 			for _, IDs := range teamRed {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 					_sleep: false, roomname: roomname}
// // 			}
// // 			for _, IDs := range teamBlue {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 					_sleep: false, roomname: roomname}
// // 			}
// // 			for _, _ID := range Nicknames[roomname] {
// // 				h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 			}

// // 			Conns = strconv.Itoa(len(h.rooms[roomname]))
// // 			h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 			h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}

// // 			log.Println("nickname list in ", roomname, ": ", Nicknames)
// // 			log.Println("team red: ", teamRed, "team blue: ", teamBlue)
// // 			go func() {
// // 				// meaning for next one set 1 round 2
// // 				store, store2 := map[string][]int{}, map[string][]int{}
// // 				store[l.RoomName] = []int{1}
// // 				store2[l.RoomName] = []int{1}

// // 				TtrackSet.store <- store2
// // 				TtrackRound.store <- store
// // 				PadLock(h, 1, 1, false, roomname)
// // 			}()
// // 		}()
// // 		go func() {
// // 			store := map[string]map[string][]string{}
// // 			store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}
// // 			storeShuffle.store <- store
// // 		}()

// // 		_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 		latest = true

// // 		// to clean up room code
// // 		storePrivateRoom.done <- true

// // 	} else {
// // 		log.Println("problem in sending")
// // 	}
// // }
// // if len(teamRed) > 0 && len(teamBlue) > 0 {
// // 	log.Println("teamRed: ", teamRed[0])
// // 	log.Println("teamBlue: ", teamBlue[0])
// // 	// broadcasting nicknames of teams and broadcasting their team name
// // 	log.Println("sending det")

// // 	TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		for _, IDs := range teamBlue {
// // 			h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 		}
// // 		for _, IDs := range teamRed {
// // 			h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 		}
// // 		for _, IDs := range teamRed {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 				_sleep: false, roomname: roomname}
// // 		}
// // 		for _, IDs := range teamBlue {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 				_sleep: false, roomname: roomname}
// // 		}
// // 		for _, _ID := range Nicknames[roomname] {
// // 			h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 		}
// // 	}()
// // 	log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 	SendRoomTime(h, roomname)
// // 	X := strconv.Itoa(len(h.rooms[roomname]))
// // 	h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + X}
// // 	roomMode := getRoomList[l.RoomName].Type
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		Conns = strconv.Itoa(len(h.rooms[roomname]))

// // 		_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 		h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}
// // 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 	}()
// // 	log.Println("team red: ", teamRed, "team blue: ", teamBlue)

// // 	nicknames := []string{}

// // 	// creating sheet
// // 	for _, id := range teamRed {
// // 		nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 	}
// // 	for _, id := range teamBlue {
// // 		nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 	}

// // 	createSheet := PackSheet{Name: make([]string, 4), Sheet: make(map[string]map[string]string)}
// // 	createSheet.Name = append(createSheet.Name, nicknames...)
// // 	log.Println("nicknames: ", createSheet.Name)

// // 	for _, names := range nicknames {
// // 		createSheet.Sheet[names] = map[string]string{
// // 			"round1": "",
// // 		}
// // 	}

// // 	log.Println("nickanems: ", getNicknamesViaID)

// // 	pack, _ := json.Marshal(&createSheet)

// // 	h.broadcast <- BroadcastReq{Token: "UpdateCheatSheet: " + string(pack), RoomID: roomname}

// // 	log.Println("created sheet: ", string(pack))

// // 	go func() {
// // 		store := map[string]PackSheet{}
// // 		store[roomname] = createSheet

// // 		storeSheetUpdate.store <- store
// // 	}()

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		bookName := getRoomRequestedToken[roomname][_BookKey]
// // 		sendDic := SetupDictionaryURL + SendDictionary(bookName)

// // 		h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}
// // 	}()

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		// meaning for next one set 1 round 2
// // 		store, store2 := map[string][]int{}, map[string][]int{}
// // 		store[l.RoomName] = []int{1}
// // 		store2[l.RoomName] = []int{1}

// // 		TtrackSet.store <- store2
// // 		TtrackRound.store <- store
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store3 := map[string]map[string][]string{}
// // 		store3[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 		storeShuffle.store <- store3
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		latest = true
// // 		storePublicRoom.done <- true
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store := map[string]map[string][]string{}
// // 		store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 		PadLock(h, 1, 1, false, roomname, store)
// // 	}()
// // } else {
// // 	log.Println("problem in sending")
// // }
// // rand.Shuffle(len(draft), func(i int, j int) {
// // 	draft[i], draft[j] = draft[j], draft[i]
// // })

// // teamRed := dataset.EraseAfter(draft, len(draft)/2)
// // teamBlue := dataset.EraseBefore(draft, len(draft)/2)

// // draftPowers := ChoosenPowers(roomname)

// // rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 	draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // })

// // teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // for _, powers := range teamRedPowers {
// // 	for _, IDs := range teamRed {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 			_sleep: false, to: IDs, roomname: roomname}
// // 	}
// // }

// // for _, powers := range teamBluePowers {
// // 	for _, IDs := range teamBlue {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 			_sleep: false, to: IDs, roomname: roomname}
// // 	}
// // }
// // if len(teamRed) == 2 && len(teamBlue) == 2 {
// // 	log.Println("teamRed: ", teamRed[0])
// // 	log.Println("teamBlue: ", teamBlue[0])
// // 	// broadcasting nicknames of teams and broadcasting their team name
// // 	go func() {
// // 		log.Println("sending det")

// // 		TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name

// // 		for _, IDs := range teamBlue {
// // 			h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 		}
// // 		for _, IDs := range teamRed {
// // 			h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 		}
// // 		draftPowers := ChoosenPowers(roomname)

// // 		rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 			draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 		})

// // 		teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 		teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 		for _, powers := range teamRedPowers {
// // 			for _, IDs := range teamRed {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 					_sleep: false, to: IDs, roomname: roomname}
// // 			}
// // 		}

// // 		for _, powers := range teamBluePowers {
// // 			for _, IDs := range teamBlue {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 					_sleep: false, to: IDs, roomname: roomname}
// // 			}
// // 		}
// // 		log.Println("team red powers", teamRedPowers)
// // 		log.Println("team blue powers", teamBluePowers)

// // 		for _, IDs := range teamBlue {
// // 			h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 		}
// // 		for _, IDs := range teamRed {
// // 			h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 		}
// // 		for _, _ID := range Nicknames[roomname] {
// // 			h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 		}

// // 		log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 		Conns = strconv.Itoa(len(h.rooms[roomname]))
// // 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 		log.Println("team red: ", teamRed, "team blue: ", teamBlue)
// // 		go func() {
// // 			// meaning for next one set 1 round 2
// // 			store, store2 := map[string][]int{}, map[string][]int{}
// // 			store[l.RoomName] = []int{1}
// // 			store2[l.RoomName] = []int{1}

// // 			TtrackSet.store <- store2
// // 			TtrackRound.store <- store
// // 			PadLock(h, 1, 1, false, roomname)
// // 		}()
// // 	}()

// // 	go func() {
// // 		store := map[string]map[string][]string{}
// // 		store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 		storeShuffle.store <- store
// // 	}()

// // 	_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 	h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}

// // 	latest = true
// // 	storePublicRoom.done <- true
// // } else {
// // 	log.Println("problem in sending")
// // }
// // if len(teamRed) > 0 && len(teamBlue) > 0 {
// // 	log.Println("teamRed: ", teamRed[0])
// // 	log.Println("teamBlue: ", teamBlue[0])
// // 	// broadcasting nicknames of teams and broadcasting their team name
// // 	log.Println("sending det")

// // 	TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		for _, IDs := range teamBlue {
// // 			h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 		}
// // 		for _, IDs := range teamRed {
// // 			h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 		}
// // 		for _, IDs := range teamRed {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 				_sleep: false, roomname: roomname}
// // 		}
// // 		for _, IDs := range teamBlue {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 				_sleep: false, roomname: roomname}
// // 		}
// // 		for _, _ID := range Nicknames[roomname] {
// // 			h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 		}
// // 	}()
// // 	log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 	SendRoomTime(h, roomname)
// // 	X := strconv.Itoa(len(h.rooms[roomname]))
// // 	h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + X}
// // 	roomMode := getRoomList[l.RoomName].Type
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		Conns = strconv.Itoa(len(h.rooms[roomname]))

// // 		_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 		h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}
// // 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 	}()
// // 	log.Println("team red: ", teamRed, "team blue: ", teamBlue)

// // 	nicknames := []string{}

// // 	// creating sheet
// // 	for _, id := range teamRed {
// // 		nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 	}
// // 	for _, id := range teamBlue {
// // 		nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 	}

// // 	createSheet := PackSheet{Name: make([]string, 4), Sheet: make(map[string]map[string]string)}
// // 	createSheet.Name = append(createSheet.Name, nicknames...)
// // 	log.Println("nicknames: ", createSheet.Name)

// // 	for _, names := range nicknames {
// // 		createSheet.Sheet[names] = map[string]string{
// // 			"round1": "",
// // 		}
// // 	}

// // 	log.Println("nickanems: ", getNicknamesViaID)

// // 	pack, _ := json.Marshal(&createSheet)

// // 	h.broadcast <- BroadcastReq{Token: "UpdateCheatSheet: " + string(pack), RoomID: roomname}

// // 	log.Println("created sheet: ", string(pack))

// // 	go func() {
// // 		store := map[string]PackSheet{}
// // 		store[roomname] = createSheet

// // 		storeSheetUpdate.store <- store
// // 	}()

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		bookName := getRoomRequestedToken[roomname][_BookKey]
// // 		sendDic := SetupDictionaryURL + SendDictionary(bookName)

// // 		h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}
// // 	}()

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		// meaning for next one set 1 round 2
// // 		store, store2 := map[string][]int{}, map[string][]int{}
// // 		store[l.RoomName] = []int{1}
// // 		store2[l.RoomName] = []int{1}

// // 		TtrackSet.store <- store2
// // 		TtrackRound.store <- store
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store3 := map[string]map[string][]string{}
// // 		store3[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 		storeShuffle.store <- store3
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		latest = true
// // 		storePublicRoom.done <- true
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store := map[string]map[string][]string{}
// // 		store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 		PadLock(h, 1, 1, false, roomname, store)
// // 	}()
// // } else {
// // 	log.Println("problem in sending")
// // }
// // // todo: first check if it is a friend room
// // switch true {
// // // last call made by the one who joined without code
// // case isFriend && !joinCode:
// // 	// if not than search only for the owner of the room and his friend
// // 	fID, oID, teamBlue, teamRed := "", "", []string{}, []string{}
// // 	for _, _ID := range getSavedJoinFriendID[roomname] {
// // 		fID = _ID
// // 	}
// // 	for _, _ID := range getSavedOwnerFriendID[roomname] {
// // 		oID = _ID
// // 	}
// // 	if fID != "" && oID != "" {
// // 		for _, _ID := range cRoom[roomname] {
// // 			if _ID != fID && oID != _ID {
// // 				teamBlue = append(teamBlue, _ID)
// // 			}
// // 		}
// // 		teamRed = append(teamRed, fID, oID)

// // 		// if len(teamRed) == 2 && len(teamBlue) == 2 {
// // 		// 	log.Println("teamRed: ", teamRed[0])
// // 		// 	log.Println("teamBlue: ", teamBlue[0])
// // 		// 	// broadcasting nicknames of teams and broadcasting their team name
// // 		// 	go func() {
// // 		// 		log.Println("sending det")

// // 		// 		TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name

// // 		// 		for _, IDs := range teamBlue {
// // 		// 			h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 		// 		}
// // 		// 		for _, IDs := range teamRed {
// // 		// 			h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 		// 		}
// // 		// 		draftPowers := ChoosenPowers(roomname)

// // 		// 		rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 		// 			draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 		// 		})

// // 		// 		teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 		// 		teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 		// 		for _, powers := range teamRedPowers {
// // 		// 			for _, IDs := range teamRed {
// // 		// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 		// 					_sleep: false, to: IDs, roomname: roomname}
// // 		// 			}
// // 		// 		}

// // 		// 		for _, powers := range teamBluePowers {
// // 		// 			for _, IDs := range teamBlue {
// // 		// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 		// 					_sleep: false, to: IDs, roomname: roomname}
// // 		// 			}
// // 		// 		}
// // 		// 		log.Println("team red powers", teamRedPowers)
// // 		// 		log.Println("team blue powers", teamBluePowers)

// // 		// 		for _, IDs := range teamRed {
// // 		// 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 		// 				_sleep: false, roomname: roomname}
// // 		// 		}
// // 		// 		for _, IDs := range teamBlue {
// // 		// 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 		// 				_sleep: false, roomname: roomname}
// // 		// 		}
// // 		// 		for _, _ID := range Nicknames[roomname] {
// // 		// 			h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 		// 		}

// // 		// 		log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 		// 		Conns = strconv.Itoa(len(h.rooms[roomname]))
// // 		// 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 		// 		log.Println("team red: ", teamRed, "team blue: ", teamBlue)
// // 		// 	}()

// // 		// 	go func() {
// // 		// 		// meaning for next one set 1 round 2
// // 		// 		store, store2 := map[string][]int{}, map[string][]int{}
// // 		// 		store[l.RoomName] = []int{1}
// // 		// 		store2[l.RoomName] = []int{1}

// // 		// 		TtrackSet.store <- store2
// // 		// 		TtrackRound.store <- store
// // 		// 		PadLock(h, 1, 1, false, roomname)
// // 		// 	}()

// // 		// 	h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}

// // 		// 	go func() {
// // 		// 		store := map[string]map[string][]string{}
// // 		// 		store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}
// // 		// 		storeShuffle.store <- store
// // 		// 	}()

// // 		// 	_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 		// 	h.broadcast <- BroadcastReq{RoomID: roomname, Token: "Friend: true"}

// // 		// 	latest = true

// // 		// 	// to clean up room code
// // 		// 	storeFriendRoom.done <- true
// // 		// } else {
// // 		// 	log.Println("problem in sending")
// // 		// }

// // 		draftPowers := ChoosenPowers(roomname)

// // 		rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 			draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 		})

// // 		teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 		teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 		for _, powers := range teamRedPowers {
// // 			for _, IDs := range teamRed {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 					_sleep: false, to: IDs, roomname: roomname}
// // 			}
// // 		}

// // 		for _, powers := range teamBluePowers {
// // 			for _, IDs := range teamBlue {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 					_sleep: false, to: IDs, roomname: roomname}
// // 			}
// // 		}
// // 		if len(teamRed) > 0 && len(teamBlue) > 0 {
// // 			log.Println("teamRed: ", teamRed[0])
// // 			log.Println("teamBlue: ", teamBlue[0])
// // 			// broadcasting nicknames of teams and broadcasting their team name
// // 			log.Println("sending det")

// // 			TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				for _, IDs := range teamBlue {
// // 					h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 				}
// // 				for _, IDs := range teamRed {
// // 					h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 				}
// // 				for _, IDs := range teamRed {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 						_sleep: false, roomname: roomname}
// // 				}
// // 				for _, IDs := range teamBlue {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 						_sleep: false, roomname: roomname}
// // 				}
// // 				for _, _ID := range Nicknames[roomname] {
// // 					h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 				}
// // 			}()
// // 			log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 			SendRoomTime(h, roomname)
// // 			X := strconv.Itoa(len(h.rooms[roomname]))
// // 			h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + X}
// // 			roomMode := getRoomList[l.RoomName].Type
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}

// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				Conns = strconv.Itoa(len(h.rooms[roomname]))

// // 				_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 				h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 			}()
// // 			log.Println("team red: ", teamRed, "team blue: ", teamBlue)

// // 			nicknames := []string{}

// // 			// creating sheet
// // 			for _, id := range teamRed {
// // 				nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 			}
// // 			for _, id := range teamBlue {
// // 				nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 			}

// // 			createSheet := PackSheet{Name: make([]string, 4), Sheet: make(map[string]map[string]string)}
// // 			createSheet.Name = append(createSheet.Name, nicknames...)
// // 			log.Println("nicknames: ", createSheet.Name)

// // 			for _, names := range nicknames {
// // 				createSheet.Sheet[names] = map[string]string{
// // 					"round1": "",
// // 				}
// // 			}

// // 			log.Println("nickanems: ", getNicknamesViaID)

// // 			pack, _ := json.Marshal(&createSheet)

// // 			h.broadcast <- BroadcastReq{Token: "UpdateCheatSheet: " + string(pack), RoomID: roomname}

// // 			log.Println("created sheet: ", string(pack))

// // 			go func() {
// // 				store := map[string]PackSheet{}
// // 				store[roomname] = createSheet

// // 				storeSheetUpdate.store <- store
// // 			}()

// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				bookName := getRoomRequestedToken[roomname][_BookKey]
// // 				sendDic := SetupDictionaryURL + SendDictionary(bookName)

// // 				h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}
// // 			}()

// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				// meaning for next one set 1 round 2
// // 				store, store2 := map[string][]int{}, map[string][]int{}
// // 				store[l.RoomName] = []int{1}
// // 				store2[l.RoomName] = []int{1}

// // 				TtrackSet.store <- store2
// // 				TtrackRound.store <- store
// // 			}()
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store3 := map[string]map[string][]string{}
// // 				store3[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 				storeShuffle.store <- store3
// // 			}()
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				latest = true
// // 				storePublicRoom.done <- true
// // 			}()
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store := map[string]map[string][]string{}
// // 				store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 				PadLock(h, 1, 1, false, roomname, store)
// // 			}()
// // 		} else {
// // 			log.Println("problem in sending")
// // 		}
// // 	} else {
// // 		log.Println("friend ID and owner ID is not found", getSavedJoinFriendID[roomname])
// // 	}

// // // last call made by friend by joining with code
// // case isFriend && joinCode:
// // 	log.Println("joined code")
// // 	log.Println("room name: ", roomname)
// // 	fID, oID, teamBlue, teamRed := "", "", []string{}, []string{}
// // 	for _, _ID := range getSavedJoinFriendID[roomname] {
// // 		fID = _ID
// // 	}
// // 	for _, _ID := range getSavedOwnerFriendID[roomname] {
// // 		oID = _ID
// // 	}
// // 	if fID != "" && oID != "" {
// // 		for _, _ID := range cRoom[roomname] {
// // 			if _ID != fID && oID != _ID {
// // 				teamBlue = append(teamBlue, _ID)
// // 			}
// // 			log.Println("_ID: ", _ID)
// // 		}
// // 		log.Println("fID: ", fID, "oID: ", oID)
// // 		teamRed = append(teamRed, fID, oID)

// // 		draftPowers := ChoosenPowers(roomname)

// // 		rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 			draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 		})

// // 		teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 		teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 		for _, powers := range teamRedPowers {
// // 			for _, IDs := range teamRed {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 					_sleep: false, to: IDs, roomname: roomname}
// // 			}
// // 		}

// // 		for _, powers := range teamBluePowers {
// // 			for _, IDs := range teamBlue {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 					_sleep: false, to: IDs, roomname: roomname}
// // 			}
// // 		}

// // 		log.Println("teamRed: ", teamRed)
// // 		log.Println("teamBlue: ", teamBlue)
// // 		log.Println("lens of red and blue: ", len(teamRed), len(teamBlue))
// // 		log.Println("croom: ", cRoom[roomname])
// // 		// if len(teamRed) == 2 && len(teamBlue) == 2 {
// // 		// 	// broadcasting nicknames of teams and broadcasting their team name
// // 		// 	go func() {
// // 		// 		log.Println("sending det")

// // 		// 		TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name

// // 		// 		for _, IDs := range teamBlue {
// // 		// 			h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 		// 		}
// // 		// 		for _, IDs := range teamRed {
// // 		// 			h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 		// 		}
// // 		// 		draftPowers := ChoosenPowers(roomname)

// // 		// 		rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 		// 			draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 		// 		})

// // 		// 		teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 		// 		teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 		// 		for _, powers := range teamRedPowers {
// // 		// 			for _, IDs := range teamRed {
// // 		// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 		// 					_sleep: false, to: IDs, roomname: roomname}
// // 		// 			}
// // 		// 		}

// // 		// 		for _, powers := range teamBluePowers {
// // 		// 			for _, IDs := range teamBlue {
// // 		// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 		// 					_sleep: false, to: IDs, roomname: roomname}
// // 		// 			}
// // 		// 		}

// // 		// 		for _, IDs := range teamRed {
// // 		// 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 		// 				_sleep: false, roomname: roomname}
// // 		// 		}
// // 		// 		for _, IDs := range teamBlue {
// // 		// 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 		// 				_sleep: false, roomname: roomname}
// // 		// 		}
// // 		// 		for _, _ID := range Nicknames[roomname] {
// // 		// 			h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 		// 		}

// // 		// 		Conns = strconv.Itoa(len(h.rooms[roomname]))
// // 		// 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 		// 		h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}

// // 		// 		log.Println("nickname list in ", roomname, ": ", Nicknames)
// // 		// 		log.Println("team red: ", teamRed, "team blue: ", teamBlue)
// // 		// 		go func() {
// // 		// 			// meaning for next one set 1 round 2
// // 		// 			store, store2 := map[string][]int{}, map[string][]int{}
// // 		// 			store[l.RoomName] = []int{1}
// // 		// 			store2[l.RoomName] = []int{1}

// // 		// 			TtrackSet.store <- store2
// // 		// 			TtrackRound.store <- store
// // 		// 			PadLock(h, 1, 1, false, roomname)
// // 		// 		}()
// // 		// 	}()
// // 		// 	go func() {
// // 		// 		store := map[string]map[string][]string{}
// // 		// 		store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}
// // 		// 		storeShuffle.store <- store
// // 		// 	}()

// // 		// 	_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 		// 	h.broadcast <- BroadcastReq{RoomID: roomname, Token: "Friend: true"}
// // 		// 	latest = true

// // 		// 	// to clean up room code
// // 		// 	storeFriendRoom.done <- true

// // 		// } else {
// // 		// 	log.Println("problem in sending")
// // 		// }
// // 		if len(teamRed) > 0 && len(teamBlue) > 0 {
// // 			log.Println("teamRed: ", teamRed[0])
// // 			log.Println("teamBlue: ", teamBlue[0])
// // 			// broadcasting nicknames of teams and broadcasting their team name
// // 			log.Println("sending det")

// // 			TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				for _, IDs := range teamBlue {
// // 					h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 				}
// // 				for _, IDs := range teamRed {
// // 					h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 				}
// // 				for _, IDs := range teamRed {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 						_sleep: false, roomname: roomname}
// // 				}
// // 				for _, IDs := range teamBlue {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 						_sleep: false, roomname: roomname}
// // 				}
// // 				for _, _ID := range Nicknames[roomname] {
// // 					h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 				}
// // 			}()
// // 			log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 			SendRoomTime(h, roomname)
// // 			X := strconv.Itoa(len(h.rooms[roomname]))
// // 			h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + X}
// // 			roomMode := getRoomList[l.RoomName].Type
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}

// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				Conns = strconv.Itoa(len(h.rooms[roomname]))

// // 				_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 				h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}
// // 				h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 			}()
// // 			log.Println("team red: ", teamRed, "team blue: ", teamBlue)

// // 			nicknames := []string{}

// // 			// creating sheet
// // 			for _, id := range teamRed {
// // 				nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 			}
// // 			for _, id := range teamBlue {
// // 				nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 			}

// // 			createSheet := PackSheet{Name: make([]string, 4), Sheet: make(map[string]map[string]string)}
// // 			createSheet.Name = append(createSheet.Name, nicknames...)
// // 			log.Println("nicknames: ", createSheet.Name)

// // 			for _, names := range nicknames {
// // 				createSheet.Sheet[names] = map[string]string{
// // 					"round1": "",
// // 				}
// // 			}

// // 			log.Println("nickanems: ", getNicknamesViaID)

// // 			pack, _ := json.Marshal(&createSheet)

// // 			h.broadcast <- BroadcastReq{Token: "UpdateCheatSheet: " + string(pack), RoomID: roomname}

// // 			log.Println("created sheet: ", string(pack))

// // 			go func() {
// // 				store := map[string]PackSheet{}
// // 				store[roomname] = createSheet

// // 				storeSheetUpdate.store <- store
// // 			}()

// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				bookName := getRoomRequestedToken[roomname][_BookKey]
// // 				sendDic := SetupDictionaryURL + SendDictionary(bookName)

// // 				h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}
// // 			}()

// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				// meaning for next one set 1 round 2
// // 				store, store2 := map[string][]int{}, map[string][]int{}
// // 				store[l.RoomName] = []int{1}
// // 				store2[l.RoomName] = []int{1}

// // 				TtrackSet.store <- store2
// // 				TtrackRound.store <- store
// // 			}()
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store3 := map[string]map[string][]string{}
// // 				store3[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 				storeShuffle.store <- store3
// // 			}()
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				latest = true
// // 				storePublicRoom.done <- true
// // 			}()
// // 			h.wg.Add(1)
// // 			go func() {
// // 				defer h.wg.Done()
// // 				store := map[string]map[string][]string{}
// // 				store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 				PadLock(h, 1, 1, false, roomname, store)
// // 			}()
// // 		} else {
// // 			log.Println("problem in sending")
// // 		}
// // 	} else {
// // 		log.Println("friend ID and owner ID is not found", getSavedJoinFriendID[roomname], GetFriendRoom(l.Code))
// // 	}
// // default:
// // 	log.Println("none of it")
// // }
// // draft := cRoom[roomname]

// // rand.Shuffle(len(draft), func(i int, j int) {
// // 	draft[i], draft[j] = draft[j], draft[i]
// // })

// // teamRed := []string{draft[0]}
// // teamBlue := []string{draft[1]}
// // draftPowers := ChoosenPowers(roomname)

// // rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 	draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // })

// // teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // for _, powers := range teamRedPowers {
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 		_sleep: false, to: teamRed[0], roomname: roomname}
// // }

// // for _, powers := range teamBluePowers {
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 		_sleep: false, to: teamBlue[0], roomname: roomname}
// // }
// // log.Println("team red powers", teamRedPowers)
// // log.Println("team blue powers", teamBluePowers)
// // if len(teamRed) > 0 && len(teamBlue) > 0 {
// // 	log.Println("teamRed: ", teamRed[0])
// // 	log.Println("teamBlue: ", teamBlue[0])
// // 	// broadcasting nicknames of teams and broadcasting their team name
// // 	log.Println("sending det")

// // 	TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		for _, IDs := range teamBlue {
// // 			h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 		}
// // 		for _, IDs := range teamRed {
// // 			h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 		}
// // 		for _, IDs := range teamRed {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 				_sleep: false, roomname: roomname}
// // 		}
// // 		for _, IDs := range teamBlue {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 				_sleep: false, roomname: roomname}
// // 		}
// // 		for _, _ID := range Nicknames[roomname] {
// // 			h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 		}
// // 	}()
// // 	log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 	SendRoomTime(h, roomname)
// // 	X := strconv.Itoa(len(h.rooms[roomname]))
// // 	h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + X}
// // 	roomMode := getRoomList[l.RoomName].Type
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		Conns = strconv.Itoa(len(h.rooms[roomname]))

// // 		_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 		h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}
// // 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 	}()
// // 	log.Println("team red: ", teamRed, "team blue: ", teamBlue)

// // 	nicknames := []string{}

// // 	// creating sheet
// // 	for _, id := range teamRed {
// // 		nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 	}
// // 	for _, id := range teamBlue {
// // 		nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 	}

// // 	createSheet := PackSheet{Name: make([]string, 4), Sheet: make(map[string]map[string]string)}
// // 	createSheet.Name = append(createSheet.Name, nicknames...)
// // 	log.Println("nicknames: ", createSheet.Name)

// // 	for _, names := range nicknames {
// // 		createSheet.Sheet[names] = map[string]string{
// // 			"round1": "",
// // 		}
// // 	}

// // 	log.Println("nickanems: ", getNicknamesViaID)

// // 	pack, _ := json.Marshal(&createSheet)

// // 	h.broadcast <- BroadcastReq{Token: "UpdateCheatSheet: " + string(pack), RoomID: roomname}

// // 	log.Println("created sheet: ", string(pack))

// // 	go func() {
// // 		store := map[string]PackSheet{}
// // 		store[roomname] = createSheet

// // 		storeSheetUpdate.store <- store
// // 	}()

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		bookName := getRoomRequestedToken[roomname][_BookKey]
// // 		sendDic := SetupDictionaryURL + SendDictionary(bookName)

// // 		h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}
// // 	}()

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		// meaning for next one set 1 round 2
// // 		store, store2 := map[string][]int{}, map[string][]int{}
// // 		store[l.RoomName] = []int{1}
// // 		store2[l.RoomName] = []int{1}

// // 		TtrackSet.store <- store2
// // 		TtrackRound.store <- store
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store3 := map[string]map[string][]string{}
// // 		store3[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 		storeShuffle.store <- store3
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		latest = true
// // 		storePublicRoom.done <- true
// // 	}()
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store := map[string]map[string][]string{}
// // 		store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 		PadLock(h, 1, 1, false, roomname, store)
// // 	}()
// // } else {
// // 	log.Println("problem in sending")
// // }

// // todo: first check if it is a friend room
// // switch true {

// // // last call made by friend by joining with code
// // case isFriend && joinCode:
// // 	log.Println("joined code")
// // 	log.Println("room name: ", roomname)
// // 	fID, oID := "", ""

// // 	for _, _ID := range getSavedJoinFriendID[roomname] {
// // 		fID = _ID
// // 	}
// // 	for _, _ID := range getSavedOwnerFriendID[roomname] {
// // 		oID = _ID
// // 	}

// // 	rnd := []int{1, 2, 3, 5}
// // 	rand.Shuffle(len(rnd), func(i int, j int) {
// // 		rnd[i], rnd[j] = rnd[j], rnd[i]
// // 	})

// // 	teamRed := []string{fID}
// // 	teamBlue := []string{oID}

// // 	if rnd[0]%2 == 0 {
// // 		teamRed = []string{oID}
// // 		teamBlue = []string{fID}
// // 	}

// // 	draftPowers := ChoosenPowers(roomname)

// // 	rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 		draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 	})

// // 	teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 	teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 	for _, powers := range teamRedPowers {
// // 		for _, IDs := range teamRed {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 				_sleep: false, to: IDs, roomname: roomname}
// // 		}
// // 	}

// // 	for _, powers := range teamBluePowers {
// // 		for _, IDs := range teamBlue {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 				_sleep: false, to: IDs, roomname: roomname}
// // 		}
// // 	}

// // 	// if fID != "" && oID != "" {

// // 	// 	log.Println("fID: ", fID, "oID: ", oID)
// // 	// 	// so that each player get's new team
// // 	// 	if rnd[0]%2 == 0 {
// // 	// 		teamRed = []string{fID}
// // 	// 		teamBlue = []string{oID}
// // 	// 	} else {
// // 	// 		teamRed = []string{oID}
// // 	// 		teamBlue = []string{fID}
// // 	// 	}
// // 	// 	draftPowers := ChoosenPowers(roomname)

// // 	// 	rand.Shuffle(len(draftPowers), func(i int, j int) {
// // 	// 		draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
// // 	// 	})

// // 	// 	teamRedPowers := dataset.EraseAfter(draftPowers, len(draftPowers)/2)
// // 	// 	teamBluePowers := dataset.EraseBefore(draftPowers, len(draftPowers)/2)

// // 	// 	for _, powers := range teamRedPowers {
// // 	// 		for _, IDs := range teamRed {
// // 	// 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 	// 				_sleep: false, to: IDs, roomname: roomname}
// // 	// 		}
// // 	// 	}

// // 	// 	for _, powers := range teamBluePowers {
// // 	// 		for _, IDs := range teamBlue {
// // 	// 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: powers + ": true",
// // 	// 				_sleep: false, to: IDs, roomname: roomname}
// // 	// 		}
// // 	// 	}
// // 	// 	log.Println("team red powers", teamRedPowers)
// // 	// 	log.Println("team blue powers", teamBluePowers)
// // 	// 	log.Println("teamRed: ", teamRed)
// // 	// 	log.Println("teamBlue: ", teamBlue)
// // 	// 	log.Println("lens of red and blue: ", len(teamRed), len(teamBlue))
// // 	// 	log.Println("croom: ", cRoom[roomname])
// // 	// 	if len(teamRed) > 0 && len(teamBlue) > 0 {
// // 	// 		log.Println("teamRed: ", teamRed[0])
// // 	// 		log.Println("teamBlue: ", teamBlue[0])
// // 	// 		// broadcasting nicknames of teams and broadcasting their team name
// // 	// 		go func() {
// // 	// 			log.Println("sending det")

// // 	// 			TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name

// // 	// 			for _, IDs := range teamBlue {
// // 	// 				h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 	// 			}
// // 	// 			for _, IDs := range teamRed {
// // 	// 				h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 	// 			}
// // 	// 			for _, IDs := range teamRed {
// // 	// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 	// 					_sleep: false, roomname: roomname}
// // 	// 			}
// // 	// 			for _, IDs := range teamBlue {
// // 	// 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 	// 					_sleep: false, roomname: roomname}
// // 	// 			}
// // 	// 			for _, _ID := range Nicknames[roomname] {
// // 	// 				h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 	// 			}

// // 	// 			Conns = strconv.Itoa(len(h.rooms[roomname]))
// // 	// 			h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 	// 			h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}

// // 	// 			log.Println("nickname list in ", roomname, ": ", Nicknames)
// // 	// 			log.Println("team red: ", teamRed, "team blue: ", teamBlue)
// // 	// 			go func() {
// // 	// 				// meaning for next one set 1 round 2
// // 	// 				store, store2 := map[string][]int{}, map[string][]int{}
// // 	// 				store[l.RoomName] = []int{1}
// // 	// 				store2[l.RoomName] = []int{1}

// // 	// 				TtrackSet.store <- store2
// // 	// 				TtrackRound.store <- store
// // 	// 				PadLock(h, 1, 1, false, roomname)
// // 	// 			}()
// // 	// 		}()
// // 	// 		go func() {
// // 	// 			store := map[string]map[string][]string{}
// // 	// 			store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}
// // 	// 			storeShuffle.store <- store
// // 	// 		}()

// // 	// 		sendDic := SetupDictionaryURL + SendDictionary(roomname)
// // 	// 		h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}

// // 	// 		_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 	// 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "Friend: true"}
// // 	// 		latest = true

// // 	// 		// to clean up room code
// // 	// 		storeFriendRoom.done <- true

// // 	// 	} else {
// // 	// 		log.Println("problem in sending")
// // 	// 	}

// // 	// } else {
// // 	// 	log.Println("friend ID and owner ID is not found", getSavedJoinFriendID[roomname], GetFriendRoom(l.Code))
// // 	// }
// // 	if len(teamRed) > 0 && len(teamBlue) > 0 {
// // 		log.Println("teamRed: ", teamRed[0])
// // 		log.Println("teamBlue: ", teamBlue[0])
// // 		// broadcasting nicknames of teams and broadcasting their team name
// // 		log.Println("sending det")

// // 		TBlue, TRed := "TeamBlue BLUE ", "TeamRed RED " // nick name + team name
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			for _, IDs := range teamBlue {
// // 				h.broadcast <- BroadcastReq{Token: TBlue + getNicknamesViaID[IDs][roomname] + ": BLUE", RoomID: roomname}
// // 			}
// // 			for _, IDs := range teamRed {
// // 				h.broadcast <- BroadcastReq{Token: TRed + getNicknamesViaID[IDs][roomname] + ": RED", RoomID: roomname}
// // 			}
// // 			for _, IDs := range teamRed {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: RED", to: IDs,
// // 					_sleep: false, roomname: roomname}
// // 			}
// // 			for _, IDs := range teamBlue {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Team: BLUE", to: IDs,
// // 					_sleep: false, roomname: roomname}
// // 			}
// // 			for _, _ID := range Nicknames[roomname] {
// // 				h.broadcast <- BroadcastReq{Token: "NicknameLists: " + _ID, RoomID: roomname}
// // 			}
// // 		}()
// // 		log.Println("nickname list in ", roomname, ": ", Nicknames)

// // 		SendRoomTime(h, roomname)
// // 		X := strconv.Itoa(len(h.rooms[roomname]))
// // 		h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + X}
// // 		roomMode := getRoomList[l.RoomName].Type
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}

// // 		h.wg.Go(func() {
// // 			Conns = strconv.Itoa(len(h.rooms[roomname]))

// // 			_REDTeam["RED"], _BLUEteam["BLUE"] = teamRed, teamBlue
// // 			h.broadcast <- BroadcastReq{Token: "TO: " + roomname, RoomID: roomname}
// // 			h.broadcast <- BroadcastReq{RoomID: roomname, Token: "ActiveConns: " + Conns}
// // 		})
// // 		log.Println("team red: ", teamRed, "team blue: ", teamBlue)

// // 		nicknames := []string{}

// // 		// creating sheet
// // 		for _, id := range teamRed {
// // 			nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 		}
// // 		for _, id := range teamBlue {
// // 			nicknames = append(nicknames, getNicknamesViaID[id][roomname])
// // 		}

// // 		createSheet := PackSheet{Name: make([]string, 4), Sheet: make(map[string]map[string]string)}
// // 		createSheet.Name = append(createSheet.Name, nicknames...)
// // 		log.Println("nicknames: ", createSheet.Name)

// // 		for _, names := range nicknames {
// // 			createSheet.Sheet[names] = map[string]string{
// // 				"round1": "",
// // 			}
// // 		}

// // 		log.Println("nicknames: ", getNicknamesViaID)

// // 		pack, _ := json.Marshal(&createSheet)

// // 		h.broadcast <- BroadcastReq{Token: "UpdateCheatSheet: " + string(pack), RoomID: roomname}

// // 		log.Println("created sheet: ", string(pack))

// // 		go func() {
// // 			store := map[string]PackSheet{}
// // 			store[roomname] = createSheet

// // 			storeSheetUpdate.store <- store
// // 		}()

// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			bookName := getRoomRequestedToken[roomname][_BookKey]
// // 			sendDic := SetupDictionaryURL + SendDictionary(bookName)

// // 			h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}
// // 		}()

// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			// meaning for next one set 1 round 2
// // 			store, store2 := map[string][]int{}, map[string][]int{}
// // 			store[l.RoomName] = []int{1}
// // 			store2[l.RoomName] = []int{1}

// // 			TtrackSet.store <- store2
// // 			TtrackRound.store <- store
// // 		}()
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			store3 := map[string]map[string][]string{}
// // 			store3[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 			storeShuffle.store <- store3
// // 		}()
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			latest = true
// // 			storePublicRoom.done <- true
// // 		}()
// // 		h.wg.Add(1)
// // 		go func() {
// // 			defer h.wg.Done()
// // 			store := map[string]map[string][]string{}
// // 			store[roomname] = map[string][]string{_TeamRedKey: teamRed, _TeamBlueKey: teamBlue}

// // 			PadLock(h, 1, 1, false, roomname, store)
// // 		}()
// // 	} else {
// // 		log.Println("problem in sending")
// // 	}
// // default:
// // 	log.Println("none of it")
// // }

// // if l.Code == "" {
// // 	isFriendRoom = FriendRoomExists(l.To, "", false) // if the joined room is a friend room
// // 	isPrivateRoom = PrivateRoomExists(l.To, "", false)
// // } else if l.Code != "" {
// // 	isFriendRoom = FriendRoomExists("", l.Code, true)
// // 	isPrivateRoom = PrivateRoomExists("", l.Code, true)
// // }

// // switch true {
// // case isFriendRoom:
// // 	exists = isFriendRoom
// // 	l.RoomCapacity = storeRoom[GetFriendRoom(l.Code)]
// // 	clientCount = CountRoomMembers(h, GetFriendRoom(l.Code))
// // 	isFull = l.RoomCapacity <= clientCount
// // 	roomname = GetFriendRoom(l.Code)
// // 	if storeRoom[GetFriendRoom(l.Code)] == 4 {
// // 		roomMode = "2v2"
// // 	} else {
// // 		roomMode = "1v1"
// // 	}
// // case isPrivateRoom:
// // 	exists = isPrivateRoom
// // 	l.RoomCapacity = storeRoom[GetPrivateRoom(l.Code)]
// // 	clientCount = CountRoomMembers(h, GetPrivateRoom(l.Code))
// // 	isFull = l.RoomCapacity <= clientCount
// // 	roomname = GetPrivateRoom(l.Code)

// // 	if storeRoom[GetPrivateRoom(l.Code)] == 4 {
// // 		roomMode = "2v2"
// // 	} else {
// // 		roomMode = "1v1"
// // 	}

// // 	// public room
// // default:
// // 	if storeRoom[l.RoomName] == 4 {
// // 		roomMode = "2v2"
// // 	} else {
// // 		roomMode = "1v1"
// // 	}
// // }
// // if !isFriendRoom {
// // 	log.Println("not friend")
// // 	isPrivateRoom = PrivateRoomExists("", l.Code, true)

// // } else if isFriendRoom {
// // 	exists = isFriendRoom
// // 	l.RoomCapacity = storeRoom[GetFriendRoom(l.Code)]
// // 	clientCount = CountRoomMembers(h, GetFriendRoom(l.Code))
// // 	isFull = l.RoomCapacity <= clientCount
// // 	if storeRoom[GetFriendRoom(l.Code)] == 4 {
// // 		roomMode = "2v2"
// // 	} else {
// // 		roomMode = "1v1"
// // 	}
// // }

// // if isPrivateRoom {
// // 	exists = isPrivateRoom
// // 	l.RoomCapacity = storeRoom[GetPrivateRoom(l.Code)]
// // 	clientCount = CountRoomMembers(h, GetPrivateRoom(l.Code))
// // 	isFull = l.RoomCapacity <= clientCount
// // 	if storeRoom[GetPrivateRoom(l.Code)] == 4 {
// // 		roomMode = "2v2"
// // 	} else {
// // 		roomMode = "1v1"
// // 	}
// // }

// // is2v2 := storeRoom[roomname] == 4

// // log.Println("is there a friend room", isFriendRoom)
// // log.Println("friend room list: ", getFriendRoom)
// // log.Println("friend room code: ", getFriendRoomCode)
// // log.Println("friend via code:", FriendRoomExists("", l.Code, true))
// // log.Println("friend via room:", FriendRoomExists(l.To, "", false))
// // log.Println("private room list: ", getPrivateRoom)
// // log.Println("private room code: ", getPrivateRoomCode)
// // log.Println("private via code:", PrivateRoomExists("", l.Code, true))
// // log.Println("is private room:", isPrivateRoom)
// // h.wg.Add(1)
// // go func() {
// // 	defer h.wg.Done()
// // 	store := map[string]map[string]string{}
// // 	store[l.To] = map[string]string{
// // 		l.ID: l.NickName,
// // 	}
// // 	Snickname <- store
// // }()
// // sendPowers := ChoosenPowers(l.To)

// // for _, power := range sendPowers {
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: power + "X" + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // }

// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomDecisionTime: " + l.DecisionTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomGameTime: " + l.GameTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCategory: " + l.Category, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomBook: " + l.Book, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.broadcast <- BroadcastReq{Token: "OnlineNickName: " + l.NickName, RoomID: l.RoomName}

// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomDecisionTime: " + l.DecisionTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomGameTime: " + l.GameTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCategory: " + l.Category, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomBook: " + l.Book, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}
// // sendPowers := ChoosenPowers(l.To)
// // h.broadcast <- BroadcastReq{Token: "OnlineNickName: " + l.NickName, RoomID: l.RoomName}

// // for _, power := range sendPowers {
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: power + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // }
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Friend: false" + l.To, roomname: l.To, to: l.ID, _sleep: false}

// // h.wg.Add(1)
// // go func() {
// // 	defer h.wg.Done()
// // 	store := map[string]map[string]string{}
// // 	store[l.To] = map[string]string{
// // 		l.ID: l.NickName,
// // 	}
// // 	Snickname <- store
// // }()
// // l.To = GetFriendRoom(l.Code)
// // h.broadcast <- BroadcastReq{Token: "ActiveConns: " + Conns, RoomID: l.To}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: GetFriendRoom(l.Code), to: conn.Params("id"),
// // 	_sleep: false, token: "Friend: true"}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: GetFriendRoom(l.Code), to: conn.Params("id"),
// // 	_sleep: false, token: "TO: " + GetFriendRoom(l.Code)} // set the to for this
// // h.broadcast <- BroadcastReq{Token: "OnlineNickName: " + l.NickName, RoomID: l.RoomName}

// // h.broadcast <- BroadcastReq{Token: "ActiveConns: " + Conns, RoomID: roomname}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: conn.Params("id"),
// // 	_sleep: false, token: "Friend: true"}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname, to: conn.Params("id"),
// // 	_sleep: false, token: "TO: " + roomname} // set the to for this
// // h.broadcast <- BroadcastReq{Token: "OnlineNickName: " + l.NickName, RoomID: l.RoomName}

// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "Friend: false" + l.To, roomname: roomname, to: l.ID, _sleep: false}
// // go func() {
// // 	store := map[string]bool{}
// // 	store2 := map[string]map[string]string{}
// // 	// storing friend id to pair with his friend
// // 	store2[roomname] = map[string]string{_FriendRoomKey: l.ID}
// // 	store[roomname] = true

// // 	PlayerJoined <- store
// // 	saveJoinFriendID <- store2
// // }()
// // h.wg.Add(1)
// // go func() {
// // 	defer h.wg.Done()
// // 	store := map[string]map[string]string{}
// // 	store[l.To] = map[string]string{
// // 		l.ID: l.NickName,
// // 	}
// // 	Snickname <- store
// // }()
// // store := map[string]map[string]string{}
// // store[roomname] = map[string]string{
// // 	l.ID: l.NickName,
// // }
// // Snickname <- store

// // sendPowers := ChoosenPowers(roomname)

// // for _, power := range sendPowers {
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: power + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // }
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomDecisionTime: " + l.DecisionTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomGameTime: " + l.GameTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCategory: " + l.Category, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomBook: " + l.Book, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}

// // h.broadcast <- BroadcastReq{Token: "OnlineNickName: " + l.NickName, RoomID: l.RoomName}

// // var _pack RoomSettingsParcel
// // _pack.Book = l.Book
// // _pack.GameTime = l.GameTime
// // _pack.DecisionTime = l.DecisionTime
// // powers := []string{}

// // sendPowers := map[string]map[string]bool{}
// // sendPowers[l.RoomName] = map[string]bool{
// // 	"DRAW":   l.DrawPower,
// // 	"NEXUS":  l.NexusPower,
// // 	"TAG":    l.TagPower,
// // 	"REWIND": l.RewindPower,
// // 	"FREEZE": l.FreezePower,
// // 	"COVERT": l.CovertPower,
// // 	"BET":    l.BetPower,
// // }

// // for name, yes := range sendPowers[l.RoomName] {
// // 	if yes {
// // 		powers = append(powers, name)
// // 	}
// // }
// // _pack.Powers = powers
// // pack, _ := json.Marshal(&_pack)
// // _token := string(pack)
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{to: l.ID, roomname: l.RoomName, _sleep: false, token: "ParcelRoomSettings: " + _token}

// // go func() {
// // 	log.Println("storing list")
// // 	var g = RoomList{RoomName: l.RoomName, Category: l.Category,
// // 		Type: roomMode, Book: l.Book}
// // 	token, _ := json.Marshal(&g)
// // 	log.Println("unmarshalled: ", string(token))
// // 	h.Cbroadcast <- BroadcastReq{Token: "RoomLists: " + string(token), RoomID: l.RoomName}
// // }()

// // h.wg.Add(1)
// // go func() {
// // 	defer h.wg.Done()
// // 	store := make(map[string]int)
// // 	store[l.RoomName] = l.RoomCapacity
// // 	roomStore.store <- store
// // 	log.Println("store: ", store)

// // 	go func() {
// // 		store := map[string]map[string]bool{}
// // 		if l.RoomCapacity == 4 {
// // 			store[l.RoomName] = map[string]bool{
// // 				_SetupTossKey:   l.SetToss,
// // 				_StarterKey:     l.Starter,
// // 				_ReverseKey:     l.Reverse,
// // 				_TwoVTwoKey:     true,
// // 				_DrawPowerKey:   l.DrawPower,
// // 				_NexusPowerKey:  l.NexusPower,
// // 				_TagPowerKey:    l.TagPower,
// // 				_RewindPowerKey: l.RewindPower,
// // 				_FreezePowerKey: l.FreezePower,
// // 				_CovertPowerKey: l.CovertPower,
// // 				_BetPowerKey:    l.BetPower,
// // 			}
// // 			sendPowers := map[string]map[string]bool{}
// // 			sendPowers[l.RoomName] = map[string]bool{
// // 				_DrawPowerKey:   l.DrawPower,
// // 				_NexusPowerKey:  l.NexusPower,
// // 				_TagPowerKey:    l.TagPower,
// // 				_RewindPowerKey: l.RewindPower,
// // 				_FreezePowerKey: l.FreezePower,
// // 				_CovertPowerKey: l.CovertPower,
// // 				_BetPowerKey:    l.BetPower,
// // 			}

// // 			for power, included := range sendPowers[l.RoomName] {
// // 				if included {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: power + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // 				}
// // 			}
// // 			roomMode = "2v2"

// // 		} else {
// // 			store[l.RoomName] = map[string]bool{
// // 				_SetupTossKey:   l.SetToss,
// // 				_StarterKey:     l.Starter,
// // 				_ReverseKey:     l.Reverse,
// // 				_TwoVTwoKey:     false,
// // 				_DrawPowerKey:   l.DrawPower,
// // 				_NexusPowerKey:  l.NexusPower,
// // 				_TagPowerKey:    l.TagPower,
// // 				_RewindPowerKey: l.RewindPower,
// // 				_FreezePowerKey: l.FreezePower,
// // 				_CovertPowerKey: l.CovertPower,
// // 				_BetPowerKey:    l.BetPower,
// // 			}
// // 			sendPowers := map[string]map[string]bool{}
// // 			sendPowers[l.RoomName] = map[string]bool{
// // 				_DrawPowerKey:   l.DrawPower,
// // 				_NexusPowerKey:  l.NexusPower,
// // 				_TagPowerKey:    l.TagPower,
// // 				_RewindPowerKey: l.RewindPower,
// // 				_FreezePowerKey: l.FreezePower,
// // 				_CovertPowerKey: l.CovertPower,
// // 				_BetPowerKey:    l.BetPower,
// // 			}

// // 			for power, included := range sendPowers[l.RoomName] {
// // 				if included {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: power + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // 				}
// // 			}
// // 			roomMode = "1v1"
// // 		}
// // 		settings.store <- store
// // 	}()

// // 	go func() {
// // 		store := map[string]map[string]string{}
// // 		store[l.RoomName] = map[string]string{
// // 			_FieldKey:        l.Field,
// // 			_CategoryKey:     l.Category,
// // 			_BookKey:         l.Book,
// // 			_DecisionTimeKey: l.DecisionTime,
// // 			_GameTimeKey:     l.GameTime,
// // 		}
// // 		storeRoomRequestedToken.store <- store
// // 	}()
// // 	// storing room owner id
// // 	// so that we can pass him the control of changing the settings of the room

// // 	store3 := map[string]string{}
// // 	store3[l.RoomName] = code
// // 	storePrivateRoom.roomCode <- store3
// // 	log.Println("storing: ", store3)

// // 	store4 := map[string]map[string]bool{}
// // 	store4[l.RoomName] = map[string]bool{_PrivateRoomKey: true}
// // 	storePrivateRoom.store <- store4
// // 	log.Println("storing: ", store4)

// // 	store5 := map[string]map[string]string{}
// // 	store5[l.RoomName] = map[string]string{
// // 		l.ID: l.NickName,
// // 	}
// // 	Snickname <- store5
// // }()

// // h.wg.Add(1)
// // go func() {
// // 	defer h.wg.Done()
// // 	store := map[string]RoomList{}
// // 	store[l.RoomName] = RoomList{
// // 		RoomName: l.RoomName,
// // 		Category: l.Category,
// // 		Book:     l.Book,
// // 		Type:     "1v1",
// // 	}
// // 	store[l.RoomName] = RoomList{Type: roomMode}

// // 	storeRoomList.store <- store

// // }()

// // h.wg.Add(1)
// // go func() {
// // 	defer h.wg.Done()
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomDecisionTime: " + l.DecisionTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomGameTime: " + l.GameTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCategory: " + l.Category, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomBook: " + l.Book, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCode: " + code, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "IsEnterTainment: " + strconv.FormatBool(IsEntertainment), roomname: l.RoomName, to: l.ID, _sleep: false}
// // 	h.broadcast <- BroadcastReq{Token: "OnlineNickName: " + l.NickName, RoomID: l.RoomName}
// // }()
// // go func() {
// // 	RegisterRoom(h, conn, l.RoomName, l.ID, l.NickName, false)

// // 	store := make(map[string]int)
// // 	store[l.RoomName] = l.RoomCapacity
// // 	roomStore.store <- store

// // 	log.Println("store: ", store)

// // go func() {
// // 	store := map[string]map[string]bool{}
// // 	if l.RoomCapacity == 4 {
// // 		store[l.RoomName] = map[string]bool{
// // 			_SetupTossKey:   l.SetToss,
// // 			_StarterKey:     l.Starter,
// // 			_ReverseKey:     l.Reverse,
// // 			_TwoVTwoKey:     true,
// // 			_DrawPowerKey:   l.DrawPower,
// // 			_NexusPowerKey:  l.NexusPower,
// // 			_TagPowerKey:    l.TagPower,
// // 			_RewindPowerKey: l.RewindPower,
// // 			_FreezePowerKey: l.FreezePower,
// // 			_CovertPowerKey: l.CovertPower,
// // 			_BetPowerKey:    l.BetPower,
// // 		}
// // 		sendPowers := map[string]map[string]bool{}
// // 		sendPowers[l.RoomName] = map[string]bool{
// // 			_DrawPowerKey:   l.DrawPower,
// // 			_NexusPowerKey:  l.NexusPower,
// // 			_TagPowerKey:    l.TagPower,
// // 			_RewindPowerKey: l.RewindPower,
// // 			_FreezePowerKey: l.FreezePower,
// // 			_CovertPowerKey: l.CovertPower,
// // 			_BetPowerKey:    l.BetPower,
// // 		}

// // 		// for power, included := range sendPowers[l.RoomName] {
// // 		// 	if included {
// // 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "X: " + power + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // 		// 	}
// // 		// }

// // 		roomMode = "2v2"

// // 	} else {
// // 		store[l.RoomName] = map[string]bool{
// // 			_SetupTossKey:   l.SetToss,
// // 			_StarterKey:     l.Starter,
// // 			_ReverseKey:     l.Reverse,
// // 			_TwoVTwoKey:     false,
// // 			_DrawPowerKey:   l.DrawPower,
// // 			_NexusPowerKey:  l.NexusPower,
// // 			_TagPowerKey:    l.TagPower,
// // 			_RewindPowerKey: l.RewindPower,
// // 			_FreezePowerKey: l.FreezePower,
// // 			_CovertPowerKey: l.CovertPower,
// // 			_BetPowerKey:    l.BetPower,
// // 		}
// // 		sendPowers := map[string]map[string]bool{}
// // 		sendPowers[l.RoomName] = map[string]bool{
// // 			_DrawPowerKey:   l.DrawPower,
// // 			_NexusPowerKey:  l.NexusPower,
// // 			_TagPowerKey:    l.TagPower,
// // 			_RewindPowerKey: l.RewindPower,
// // 			_FreezePowerKey: l.FreezePower,
// // 			_CovertPowerKey: l.CovertPower,
// // 			_BetPowerKey:    l.BetPower,
// // 		}

// // 		// for power, included := range sendPowers[l.RoomName] {
// // 		// 	if included {
// // 		// 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: power + "X" + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // 		// 	}
// // 		// }

// // 		roomMode = "1v1"
// // 	}
// // 	settings.store <- store
// // }()

// // go func() {
// // 	store := map[string]map[string]string{}
// // 	store[l.RoomName] = map[string]string{
// // 		l.ID: l.NickName,
// // 	}
// // 	Snickname <- store
// // }()

// // storing room owner id
// // so that we can pass him the control of changing the settings of the room
// //go func() {
// // store := map[string]map[string]bool{}
// // store[l.RoomName] = map[string]bool{l.ID: true}

// // store3 := map[string]map[string]bool{}
// // store3[l.RoomName] = map[string]bool{_PublicRoomKey: true}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomDecisionTime: " + l.DecisionTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomGameTime: " + l.GameTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCategory: " + l.Category, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomBook: " + l.Book, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}
// // h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "IsEnterTainment: " + strconv.FormatBool(IsEntertainment), roomname: l.RoomName, to: l.ID, _sleep: false}

// // storePublicRoom.store <- store3
// // storeRoomOwner.store <- store
// //}()

// // go func() {
// // 	store := map[string]map[string]string{}
// // 	store[l.RoomName] = map[string]string{
// // 		_FieldKey:        l.Field,
// // 		_CategoryKey:     l.Category,
// // 		_BookKey:         l.Book,
// // 		_DecisionTimeKey: l.DecisionTime,
// // 		_GameTimeKey:     l.GameTime,
// // 	}
// // 	storeRoomRequestedToken.store <- store
// // }()

// // go func() {
// // 	log.Println("storing list")
// // _type := "1v1"
// // if is2v2 {
// // 	_type = "2v2"
// // }
// // store := map[string]RoomList{}
// // store[l.RoomName] = RoomList{
// // 	RoomName: l.RoomName,
// // 	Category: l.Category,
// // 	Book:     l.Book,
// // 	Type:     _type,
// // }

// // storeRoomList.store <- store

// // 		var g = RoomList{RoomName: l.RoomName, Category: l.Category,
// // 			Type: roomMode, Book: l.Book}
// // 		token, _ := json.Marshal(&g)
// // 		log.Println("stored list: ", store)
// // 		log.Println("unmarsahlled: ", string(token))
// // 		h.Cbroadcast <- BroadcastReq{Token: "RoomLists: " + string(token), RoomID: l.RoomName}
// // //	}()
// // h.broadcast <- BroadcastReq{Token: "OnlineNickName: " + l.NickName, RoomID: l.RoomName}

// // RoomValidationTokensForJoined(l.RoomName, h, false, is2v2, true)
// // broadcast <- "Hosts: " + l.RoomName
// //				}()
// // RegisterRoom(h, conn, l.RoomName, l.ID, l.NickName, false)

// // h.wg.Add(1)
// // log.Println("friend room")
// // // store the capacity
// // go func() {
// // 	defer h.wg.Done()
// // 	store := make(map[string]int)
// // 	store[l.RoomName] = l.RoomCapacity
// // 	roomStore.store <- store
// // 	log.Println("store: ", store)

// // 	// store the room setting
// // 	go func() {
// // 		store := map[string]map[string]bool{}
// // 		if l.RoomCapacity == 4 {
// // 			store[l.RoomName] = map[string]bool{
// // 				_SetupTossKey:   l.SetToss,
// // 				_StarterKey:     l.Starter,
// // 				_ReverseKey:     l.Reverse,
// // 				_TwoVTwoKey:     true,
// // 				_DrawPowerKey:   l.DrawPower,
// // 				_NexusPowerKey:  l.NexusPower,
// // 				_TagPowerKey:    l.TagPower,
// // 				_RewindPowerKey: l.RewindPower,
// // 				_FreezePowerKey: l.FreezePower,
// // 				_CovertPowerKey: l.CovertPower,
// // 				_BetPowerKey:    l.BetPower,
// // 			}
// // 			sendPowers := map[string]map[string]bool{}
// // 			sendPowers[l.RoomName] = map[string]bool{
// // 				_DrawPowerKey:   l.DrawPower,
// // 				_NexusPowerKey:  l.NexusPower,
// // 				_TagPowerKey:    l.TagPower,
// // 				_RewindPowerKey: l.RewindPower,
// // 				_FreezePowerKey: l.FreezePower,
// // 				_CovertPowerKey: l.CovertPower,
// // 				_BetPowerKey:    l.BetPower,
// // 			}

// // 			for power, included := range sendPowers[l.RoomName] {
// // 				if included {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: power + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // 				}
// // 			}
// // 			roomMode = "2v2"

// // 		} else {
// // 			store[l.RoomName] = map[string]bool{
// // 				_SetupTossKey:   l.SetToss,
// // 				_StarterKey:     l.Starter,
// // 				_ReverseKey:     l.Reverse,
// // 				_TwoVTwoKey:     false,
// // 				_DrawPowerKey:   l.DrawPower,
// // 				_NexusPowerKey:  l.NexusPower,
// // 				_TagPowerKey:    l.TagPower,
// // 				_RewindPowerKey: l.RewindPower,
// // 				_FreezePowerKey: l.FreezePower,
// // 				_CovertPowerKey: l.CovertPower,
// // 				_BetPowerKey:    l.BetPower,
// // 			}
// // 			sendPowers := map[string]map[string]bool{}
// // 			sendPowers[l.RoomName] = map[string]bool{
// // 				_DrawPowerKey:   l.DrawPower,
// // 				_NexusPowerKey:  l.NexusPower,
// // 				_TagPowerKey:    l.TagPower,
// // 				_RewindPowerKey: l.RewindPower,
// // 				_FreezePowerKey: l.FreezePower,
// // 				_CovertPowerKey: l.CovertPower,
// // 				_BetPowerKey:    l.BetPower,
// // 			}

// // 			for power, included := range sendPowers[l.RoomName] {
// // 				if included {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: power + ": true", roomname: l.RoomName, to: l.ID, _sleep: false}
// // 				}
// // 			}
// // 			roomMode = "1v1"
// // 		}
// // 		settings.store <- store
// // 	}()
// // 	go func() {
// // 		store := map[string]map[string]string{}
// // 		store[l.RoomName] = map[string]string{
// // 			_FieldKey:        l.Field,
// // 			_CategoryKey:     l.Category,
// // 			_BookKey:         l.Book,
// // 			_DecisionTimeKey: l.DecisionTime,
// // 			_GameTimeKey:     l.GameTime,
// // 		}
// // 		storeRoomRequestedToken.store <- store
// // 	}()
// // 	// store the nickname
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store := map[string]map[string]string{}
// // 		store[l.RoomName] = map[string]string{
// // 			l.ID: l.NickName,
// // 		}
// // 		Snickname <- store
// // 	}()
// // 	// so that we can pass him the control of changing the settings of the room
// // 	// and team him with his friend
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store := map[string]map[string]bool{}
// // 		store[l.RoomName] = map[string]bool{l.ID: true}

// // 		storeRoomOwner.store <- store
// // 	}()
// // 	go func() {
// // 		store := map[string]map[string]string{}
// // 		// storing friend id to pair with his friend
// // 		store[l.RoomName] = map[string]string{_FriendRoomKey: l.ID}
// // 		saveOwnerFriendID <- store
// // 	}()

// // 	// store the friend room
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()

// // 		store := map[string]map[string]bool{}
// // 		store[l.RoomName] = map[string]bool{_FriendRoomKey: true}

// // 		store2 := map[string]string{}
// // 		store2[l.RoomName] = code

// // 		storeFriendRoom.store <- store

// // 		storeFriendRoom.roomCode <- store2
// // 	}()

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		store := map[string]RoomList{}
// // 		store[l.RoomName] = RoomList{
// // 			RoomName: l.RoomName,
// // 			Category: l.Category,
// // 			Book:     l.Book,
// // 			Type:     roomMode,
// // 		}
// // 		storeRoomList.store <- store
// // 	}()

// // 	// send the room code
// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomDecisionTime: " + l.DecisionTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomGameTime: " + l.GameTime, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCategory: " + l.Category, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomBook: " + l.Book, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomMode: " + roomMode, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 		// note roomvalid...... hasnt applied for room code
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "RoomCode: " + code, roomname: l.RoomName, to: l.ID, _sleep: false}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "IsEnterTainment: " + strconv.FormatBool(IsEntertainment), roomname: l.RoomName, to: l.ID, _sleep: false}
// // 	}()
// // 	// broadcast <- "Hosts: " + l.RoomName
// // }()
// // h.broadcast <- BroadcastReq{Token: "OnlineNickName: " + l.NickName, RoomID: l.RoomName}

// // import (
// // 	"log"
// // 	"strconv"

// // 	"github.com/gofiber/contrib/websocket"
// // )

// // // func Restart(h *Hub, count int, id string, conn *websocket.Conn, redScore int, blueScore int, GR GameRoom) {
// // // 	log.Println("in restart")

// // // 	currentSet := func(_for map[string]int, roomName string) int {
// // // 		temp := _for[roomName] + 1
// // // 		return temp
// // // 	}(getTtrackSet, GR.RoomName)

// // // 	currentRound := func(_for map[string]int, roomName string) int {
// // // 		temp := _for[roomName] + 1
// // // 		return temp
// // // 	}(getTtrackRound, GR.RoomName)

// // // 	doneSession := redScore == 3 || blueScore == 3 || currentRound == 6

// // // 	log.Println("set token: ", currentSet, "round token: ", currentRound)
// // // 	log.Println("score update 👁️ team red vs team blue: ", redScore, blueScore)
// // // 	log.Println("is session done: ", doneSession)

// // // 	// working fine
// // // 	if doneSession {
// // // 		Result(h, currentSet, redScore, blueScore, GR)
// // // 	} else {
// // // 		log.Println("start over")
// // // 		if redScore != 0 {
// // // 			PadLock(h, currentSet, currentRound, false, GR.RoomName, saveShuffle)
// // // 		} else {
// // // 			PadLock(h, currentSet, currentRound, true, GR.RoomName, saveShuffle)
// // // 		}

// // // 		Set := "Set: " + strconv.Itoa(currentSet)       // to track the current set
// // // 		Round := "Round: " + strconv.Itoa(currentRound) // to track the current round

// // // 		// re-run the process for challenge
// // // 		h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: startGame}
// // // 		h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: challengeDiscussion}
// // // 		h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: waiting}
// // // 		h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: resetCount}
// // // 		h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: Set}
// // // 		h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: Round}
// // // 		h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: _roundOver}
// // // 		h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: _CanUsePower}

// // // 		DeActivatePower(h, GR.RoomName)
// // // 		DeactiveAttribute(h, GR.RoomName)
// // // 		SessionCleanUp()
// // // 		PowersCleanUp()

// // // 		if getSettings[GR.RoomName][_SetupTossKey] {
// // // 			log.Println("setting has toss on")
// // // 			h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: tossCoin}
// // // 			h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: _toss}
// // // 			h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: _tossSession}
// // // 		} else {
// // // 			h.broadcast <- BroadcastReq{RoomID: GR.RoomName, Token: setToss}
// // // 		}
// // // 		log.Println("current set: ", currentSet, "current round: ", currentRound)
// // // 		store := map[string]bool{}
// // // 		store[GR.RoomName] = true
// // // 		resetTokens <- store
// // // 		resetSesssion <- store
// // // 	}

// // // 	log.Println("round over")
// // // 	log.Println("current count: ", count)
// // // 	log.Println("done session ;", doneSession)
// // // }

// // func SendPowers(h *Hub, roomname string, teamname string, opponentTeamname string) {
// // 	log.Println("in send powers")
// // 	log.Println("stored power: ", getStoredPower)
// // 	log.Println("room name: ", roomname)

// // 	for _, id := range saveShuffle[roomname][teamname] {
// // 		switch true {

// // 		case getStoredPower[id][roomname][_FreezeKey]:
// // 			log.Println("freeze")
// // 			PFreezeMechanism(h, roomname, id, teamname, false, false)

// // 		case getStoredPower[id][roomname][_CovertKey]:
// // 			PCovertMechanism(h, roomname, id)
// // 		}
// // 	}
// // 	for _, id := range saveShuffle[roomname][opponentTeamname] {
// // 		switch true {

// // 		case getStoredPower[id][roomname][_FreezeKey]:
// // 			log.Println("freeze")
// // 			PFreezeMechanism(h, roomname, id, teamname, false, false)

// // 		case getStoredPower[id][roomname][_CovertKey]:
// // 			PCovertMechanism(h, roomname, id) // make sure the red team goes to the test
// // 		}
// // 	}
// // }

// // // PocketPowers returns  helps in storing the powers
// // func PocketPowers(h *Hub, id string, roomname string, GR GameRoom) {
// // 	switch true {
// // 	case GR.Nexus:
// // 		log.Println("NEXUS stored")
// // 		go func() {
// // 			store := map[string]map[string]map[string]bool{}
// // 			if _, ok := store[id]; !ok {
// // 				store[id] = map[string]map[string]bool{}
// // 			}
// // 			if _, ok := store[id][roomname]; !ok {
// // 				store[id][roomname] = map[string]bool{}
// // 			}
// // 			if _, ok := store[id][roomname][_NexusKey]; !ok {
// // 				store[id][roomname][_NexusKey] = true
// // 			}
// // 			store[id] = map[string]map[string]bool{roomname: {_NexusKey: true}}
// // 			storePower.store <- store
// // 		}()
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: NexusUse, to: id, _sleep: false, roomname: roomname}

// // 	case GR.Rewind:
// // 		log.Println("REWIND stored")
// // 		go func() {
// // 			store := map[string]map[string]map[string]bool{}
// // 			store[id] = map[string]map[string]bool{roomname: {_RewindKey: true}}
// // 			storePower.store <- store
// // 		}()
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RewindUse, to: id, _sleep: false, roomname: roomname}

// // 	case GR.Freeze:
// // 		log.Println("fREEZE STORED")
// // 		go func() {
// // 			store := map[string]map[string]map[string]bool{}
// // 			store[id] = map[string]map[string]bool{roomname: {_FreezeKey: true}}
// // 			storePower.store <- store
// // 		}()

// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: FreezeUse, to: id, _sleep: false, roomname: roomname}

// // 	case GR.Covert:
// // 		log.Println("COVERT STORED")

// // 		go func() {
// // 			store := map[string]map[string]map[string]bool{}
// // 			store[id] = map[string]map[string]bool{roomname: {_CovertKey: true}}
// // 			storePower.store <- store
// // 		}()
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CovertUse, to: id, _sleep: false, roomname: roomname}

// // 	case GR.Draw:
// // 		go func() {
// // 			store := map[string]map[string]map[string]bool{}
// // 			store[id] = map[string]map[string]bool{roomname: {_DrawKey: true}}
// // 			storePower.store <- store
// // 		}()
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: DrawUse, to: id, _sleep: false, roomname: roomname}

// // 	case GR.Tag:
// // 		go func() {
// // 			store := map[string]map[string]map[string]bool{}
// // 			store[id] = map[string]map[string]bool{roomname: {_TagKey: true}}
// // 			storePower.store <- store
// // 		}()
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: TagUse, to: id, _sleep: false, roomname: roomname}

// // 	case GR.Bet:
// // 		go func() {
// // 			store := map[string]map[string]map[string]bool{}
// // 			store[id] = map[string]map[string]bool{roomname: {_BetKey: true}}
// // 			storePower.store <- store
// // 		}()
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: BetUse, to: id, _sleep: false, roomname: roomname}
// // 	}
// // }

// // import (
// // 	"fmt"
// // 	"strconv"

// // 	"github.com/gofiber/contrib/websocket"
// // 	"github.com/gofiber/fiber/v2"
// // )

// // var (
// // 	ClientLens = 0
// // )

// // var (
// // 	clients    = make(map[*websocket.Conn]*client)
// // 	register   = make(chan *websocket.Conn)
// // 	unregister = make(chan *websocket.Conn)
// // )

// // func HubRun() {
// // 	for {
// // 		select {
// // 		case connection := <-register:
// // 			clients[connection] = &client{}

// // 			ClientLens = len(clients) // total number of client including the request one too

// // 		case token := <-broadcast:
// // 			for connection, c := range clients {
// // 				go func(conn *websocket.Conn, c *client) {
// // 					c.mu.Lock()
// // 					defer c.mu.Unlock()
// // 					if c.isClosing {
// // 						return
// // 					}
// // 					err := conn.WriteMessage(websocket.TextMessage, []byte(token))
// // 					if err != nil {
// // 						conn.WriteMessage(websocket.CloseMessage, []byte{})
// // 						conn.Close()
// // 						unregister <- conn
// // 					}
// // 				}(connection, c)
// // 			}
// // 		case connection := <-unregister:
// // 			delete(clients, connection)
// // 			fmt.Println("client disconnect")
// // 		}
// // 	}
// // }
// // func Watch(namespace string, app *fiber.App) {
// // 	app.Use(func(c *fiber.Ctx) error {
// // 		if websocket.IsWebSocketUpgrade(c) {
// // 			c.Locals("allowed", true)
// // 			return c.Next()
// // 		}
// // 		return c.SendStatus(fiber.StatusUpgradeRequired)
// // 	})
// // 	app.Get(namespace, websocket.New(func(c *websocket.Conn) {
// // 		fmt.Println("websocket connected")
// // 		defer func() {
// // 			unregister <- c
// // 			c.Close()
// // 		}()
// // 		register <- c
// // 		for {
// // 			tokenType, token, err := c.ReadMessage()
// // 			if err != nil {
// // 				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// // 					fmt.Println("read error:", err)
// // 				}
// // 				return
// // 			}
// // 			/**
// // 			*@TODO make sure to keep the name of the item that have used before
// // 			**/
// // 			if tokenType == websocket.TextMessage && string(rune(tokenType)) != "" {
// // 				fmt.Println("recieved message: ", string(token))
// // 				//	approved = tokenValidate.Valid(string(token))
// // 				Token := strconv.FormatBool(approved) // conversion from bool to string
// // 				switch string(token) {
// // 				case "connections":
// // 					{
// // 						// in send token the  live connection=0 which is removed
// // 						// Token = strconv.Itoa(LiveConnection)
// // 						// broadcast <- Token
// // 						// fmt.Println("live connections: ", Token)
// // 					}
// // 				case "createRoom":
// // 					{
// // 						// @TODO change live connection to id of the player
// // 						// if the client request for creating the room create that room

// // 						fmt.Println("room joined")
// // 					}
// // 				case "joinRoom":
// // 					{

// // 					}
// // 				default:
// // 					// validation
// // 					{
// // 						broadcast <- string(Token)
// // 					}

// // 				}

// // 				//	fmt.Println("approved: ", approved)
// // 			} else {
// // 				fmt.Println("invalid message type")
// // 			}
// // 		}
// // 	}, websocket.Config{Origins: []string{"http://localhost:4200"}}))
// // }

// // if GR.Chances != 2 {
// // 	if proceed {

// // 	} else {
// // 		chatMes[l.To][].Front(GR.MutualVote) // keep pushing the newly message
// // 	}
// // } else {
// // 	switch true {
// // 	case isTeamBlue:
// // 		get := ""
// // 		__temp := []string{}
// // 		for _, msgs := range chatMes.Range() {
// // 			if msgs != "AGREE" {
// // 				__temp = append(__temp, msgs)
// // 			}
// // 		}
// // 		get = __temp[0]
// // 		for _, _ids := range saveShuffle[___TeamRedKey] {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ChallengeSet: " + get, to: _ids, roomName: l.To, _sleep: false}
// // 		}

// // 	case isTeamRed:
// // 		get := ""
// // 		__temp := []string{}
// // 		for _, msgs := range chatMes.Range() {
// // 			if msgs != "AGREE" {
// // 				__temp = append(__temp, msgs)
// // 			}
// // 		}
// // 		get = __temp[0] // front value
// // 		for _, _ids := range saveShuffle[___TeamBlueKey] {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ChallengeSet: " + get, to: _ids, roomName: l.To, _sleep: false}
// // 		}
// // 	}
// // }

// // if GR.Chances != 2 {
// // 	switch true {
// // 	case isTeamRed && GR.Agree:
// // 		fmt.Println("A lock: ", isLock[___TeamRedKey],
// // 			"sav shuflle A", saveShuffle[___TeamRedKey])
// // 		// proceed to play in

// // 		fmt.Println("teamA to playing")

// // 		// signal to other players that you have to vote
// // 		switch count {
// // 		case 1:
// // 			fmt.Println("count 1")

// // 			// send the faceoff player to the playing hall
// // 			for _, r := range saveShuffle[___TeamRedKey] {
// // 				if isLock[___TeamRedKey][r] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: waiting, _sleep: false, roomName: l.To, to: r,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: newCount, _sleep: false, roomName: l.To, to: r,
// // 					}
// // 				}
// // 			}

// // 			for _, r := range saveShuffle[___TeamBlueKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{roomName: l.To, token: Unblock, to: r, _sleep: false}
// // 				//h.gameRoomBroadcast <- reqGameRoomBroadcast{roomName: l.To, token: setChallenge, to: r, _sleep: false}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: newCount, _sleep: false, roomName: l.To, to: r}
// // 			}

// // 			fmt.Println("new count:", newCount)
// // 		case 2:
// // 			fmt.Println("count 2")
// // 			// directly to playing session
// // 			for _, r := range saveShuffle[___TeamRedKey] {
// // 				if isLock[___TeamRedKey][r] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: _startGame, _sleep: false, roomName: l.To, to: r,
// // 					}

// // 				}
// // 			}
// // 			for _, r := range saveShuffle[___TeamBlueKey] {
// // 				if isLock[___TeamBlueKey][r] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: _waiting, _sleep: false, roomName: l.To, to: r,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: _startGame, _sleep: false, roomName: l.To, to: r,
// // 					}
// // 				}
// // 			}

// // 			// voting session done
// // 			h.broadcast <- BroadcastReq{RoomId: l.To, Token: challengeDiscussion}
// // 			h.broadcast <- BroadcastReq{RoomId: l.To, Token: votingSession}
// // 			newStore := map[string]string{}
// // 			reverseNewStore := map[string]string{}
// // 			for _, tag := range saveShuffle {
// // 				for _, data := range tag {
// // 					for Name, ID := range getProfData {
// // 						if ID == data {
// // 							newStore[Name] = data
// // 							reverseNewStore[data] = Name
// // 						}
// // 					}
// // 				}
// // 			}
// // 			// becuase we have saved the data of joiners separately
// // 			for _, id := range reverseNewStore {
// // 				for Ownername, Ownerid := range getOwnerProf {
// // 					if Ownerid != id {
// // 						reverseNewStore[Ownerid] = Ownername
// // 						newStore[Ownername] = Ownerid // its obvious that if the id doenst match it means that it is not in the new store too
// // 					}
// // 				}
// // 			}
// // 			fmt.Println("new Store: ", newStore)
// // 			fmt.Println("reverseNewStore: ", reverseNewStore)
// // 			fmt.Println("getOwnerProf: ", getOwnerProf)
// // 			token1 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][0]]}
// // 			token2 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][1]]}

// // 			token3 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][0]]}
// // 			token4 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][1]]}

// // 			__RED := "Display: " + token1.NickName + " " + token1.TeamName
// // 			__RED2 := "Display: " + token2.NickName + " " + token2.TeamName
// // 			__BLUE := "Display: " + token3.NickName + " " + token3.TeamName
// // 			__BLUE2 := "Display: " + token4.NickName + " " + token4.TeamName

// // 			h.broadcast <- BroadcastReq{Token: __RED, RoomId: l.To}
// // 			h.broadcast <- BroadcastReq{Token: __RED2, RoomId: l.To}

// // 			h.broadcast <- BroadcastReq{Token: __BLUE, RoomId: l.To}
// // 			h.broadcast <- BroadcastReq{Token: __BLUE2, RoomId: l.To}
// // 		default:
// // 			fmt.Println("still waiting")
// // 		}

// // 	case isTeamRed && !GR.Agree:
// // 		// keep chatting
// // 		fmt.Println("team A discussing")
// // 		for _, r := range saveShuffle[___TeamRedKey] {
// // 			if r != id {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: chatToken, to: r, _sleep: false, roomName: l.To,
// // 				}
// // 			}
// // 		}
// // 	case isTeamBlue && GR.Agree:

// // 		fmt.Println("teamB to playing")
// // 		fmt.Println("B lock: ", isLock[___TeamBlueKey],
// // 			"sav shuflle B", saveShuffle[___TeamBlueKey])

// // 		// signal to other players that you have to vote
// // 		switch count {
// // 		case 1:
// // 			fmt.Println("count 1")
// // 			// send the faceoff player to the playing hall
// // 			for _, r := range saveShuffle[___TeamBlueKey] {
// // 				if isLock[___TeamBlueKey][r] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: waiting, _sleep: false, roomName: l.To, to: r,
// // 					}
// // 				}
// // 			}

// // 			for _, r := range saveShuffle[___TeamRedKey] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{roomName: l.To, token: Unblock, to: r, _sleep: false}
// // 				//h.gameRoomBroadcast <- reqGameRoomBroadcast{roomName: l.To, token: setChallenge, to: r, _sleep: false}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: newCount, _sleep: false, roomName: l.To, to: r,
// // 				}
// // 			}
// // 		case 2:
// // 			fmt.Println("count 2")
// // 			// directly to playing session
// // 			for _, r := range saveShuffle[___TeamBlueKey] {
// // 				if isLock[___TeamBlueKey][r] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: _startGame, _sleep: false, roomName: l.To, to: r,
// // 					}
// // 				}
// // 			}

// // 			for _, r := range saveShuffle[___TeamRedKey] {
// // 				if isLock[___TeamRedKey][r] {
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: _waiting, _sleep: false, roomName: l.To, to: r,
// // 					}
// // 					h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 						token: _startGame, _sleep: false, roomName: l.To, to: r,
// // 					}
// // 				}
// // 			}
// // 			newStore := map[string]string{}
// // 			reverseNewStore := map[string]string{}
// // 			for _, tag := range saveShuffle {
// // 				for _, data := range tag {
// // 					for Name, ID := range getProfData {
// // 						if ID == data {
// // 							newStore[Name] = data
// // 							reverseNewStore[data] = Name
// // 						}
// // 					}
// // 				}
// // 			}
// // 			// becuase we have saved the data of joiners separately
// // 			for _, id := range reverseNewStore {
// // 				for Ownername, Ownerid := range getOwnerProf {
// // 					if Ownerid != id {
// // 						reverseNewStore[Ownerid] = Ownername
// // 						newStore[Ownername] = Ownerid // its obvious that if the id doenst match it means that it is not in the new store too
// // 					}
// // 				}
// // 			}
// // 			fmt.Println("new Store: ", newStore)
// // 			fmt.Println("reverseNewStore: ", reverseNewStore)
// // 			fmt.Println("getOwnerProf: ", getOwnerProf)
// // 			token1 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][0]]}
// // 			token2 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][1]]}

// // 			token3 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][0]]}
// // 			token4 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][1]]}

// // 			__RED := "Display: " + token1.NickName + " " + token1.TeamName
// // 			__RED2 := "Display: " + token2.NickName + " " + token2.TeamName
// // 			__BLUE := "Display: " + token3.NickName + " " + token3.TeamName
// // 			__BLUE2 := "Display: " + token4.NickName + " " + token4.TeamName

// // 			h.broadcast <- BroadcastReq{Token: __RED, RoomId: l.To}
// // 			h.broadcast <- BroadcastReq{Token: __RED2, RoomId: l.To}

// // 			h.broadcast <- BroadcastReq{Token: __BLUE, RoomId: l.To}
// // 			h.broadcast <- BroadcastReq{Token: __BLUE2, RoomId: l.To}
// // 			// set challenge discussion over and voting session done
// // 			h.broadcast <- BroadcastReq{RoomId: l.To, Token: challengeDiscussion}
// // 			h.broadcast <- BroadcastReq{RoomId: l.To, Token: votingSession}

// // 		default:
// // 			fmt.Println("still waiting")
// // 		}

// // 	case isTeamBlue && !GR.Agree:
// // 		fmt.Println("team A discussing")
// // 		// keep chatting
// // 		for _, r := range saveShuffle[___TeamBlueKey] {
// // 			if r != id {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: chatToken, to: r, _sleep: false, roomName: l.To,
// // 				}
// // 			}
// // 		}
// // 	}
// // } else {
// // 	if isTeamBlue {
// // 		fmt.Println("count 2")
// // 		// directly to playing session
// // 		for _, r := range saveShuffle[___TeamBlueKey] {
// // 			if isLock[___TeamBlueKey][r] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: _startGame, _sleep: false, roomName: l.To, to: r,
// // 				}
// // 			}
// // 		}

// // 		for _, r := range saveShuffle[___TeamRedKey] {
// // 			if isLock[___TeamRedKey][r] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: _waiting, _sleep: false, roomName: l.To, to: r,
// // 				}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: _startGame, _sleep: false, roomName: l.To, to: r,
// // 				}
// // 			}
// // 		}
// // 		newStore := map[string]string{}
// // 		reverseNewStore := map[string]string{}
// // 		for _, tag := range saveShuffle {
// // 			for _, data := range tag {
// // 				for Name, ID := range getProfData {
// // 					if ID == data {
// // 						newStore[Name] = data
// // 						reverseNewStore[data] = Name
// // 					}
// // 				}
// // 			}
// // 		}
// // 		// becuase we have saved the data of joiners separately
// // 		for _, id := range reverseNewStore {
// // 			for Ownername, Ownerid := range getOwnerProf {
// // 				if Ownerid != id {
// // 					reverseNewStore[Ownerid] = Ownername
// // 					newStore[Ownername] = Ownerid // its obvious that if the id doenst match it means that it is not in the new store too
// // 				}
// // 			}
// // 		}
// // 		fmt.Println("new Store: ", newStore)
// // 		fmt.Println("reverseNewStore: ", reverseNewStore)
// // 		fmt.Println("getOwnerProf: ", getOwnerProf)
// // 		token1 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][0]]}
// // 		token2 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][1]]}

// // 		token3 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][0]]}
// // 		token4 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][1]]}

// // 		__RED := "Display: " + token1.NickName + " " + token1.TeamName
// // 		__RED2 := "Display: " + token2.NickName + " " + token2.TeamName
// // 		__BLUE := "Display: " + token3.NickName + " " + token3.TeamName
// // 		__BLUE2 := "Display: " + token4.NickName + " " + token4.TeamName

// // 		h.broadcast <- BroadcastReq{Token: __RED, RoomId: l.To}
// // 		h.broadcast <- BroadcastReq{Token: __RED2, RoomId: l.To}

// // 		h.broadcast <- BroadcastReq{Token: __BLUE, RoomId: l.To}
// // 		h.broadcast <- BroadcastReq{Token: __BLUE2, RoomId: l.To}
// // 		// set challenge discussion over and voting session done
// // 		h.broadcast <- BroadcastReq{RoomId: l.To, Token: challengeDiscussion}
// // 		h.broadcast <- BroadcastReq{RoomId: l.To, Token: votingSession}

// // 	} else {
// // 		fmt.Println("count 2")
// // 		// directly to playing session
// // 		for _, r := range saveShuffle[___TeamRedKey] {
// // 			if isLock[___TeamRedKey][r] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: _startGame, _sleep: false, roomName: l.To, to: r,
// // 				}

// // 			}
// // 		}
// // 		for _, r := range saveShuffle[___TeamBlueKey] {
// // 			if isLock[___TeamBlueKey][r] {
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: _waiting, _sleep: false, roomName: l.To, to: r,
// // 				}
// // 				h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 					token: _startGame, _sleep: false, roomName: l.To, to: r,
// // 				}
// // 			}
// // 		}

// // 		// voting session done
// // 		h.broadcast <- BroadcastReq{RoomId: l.To, Token: challengeDiscussion}
// // 		h.broadcast <- BroadcastReq{RoomId: l.To, Token: votingSession}
// // 		newStore := map[string]string{}
// // 		reverseNewStore := map[string]string{}
// // 		for _, tag := range saveShuffle {
// // 			for _, data := range tag {
// // 				for Name, ID := range getProfData {
// // 					if ID == data {
// // 						newStore[Name] = data
// // 						reverseNewStore[data] = Name
// // 					}
// // 				}
// // 			}
// // 		}
// // 		// becuase we have saved the data of joiners separately
// // 		for _, id := range reverseNewStore {
// // 			for Ownername, Ownerid := range getOwnerProf {
// // 				if Ownerid != id {
// // 					reverseNewStore[Ownerid] = Ownername
// // 					newStore[Ownername] = Ownerid // its obvious that if the id doenst match it means that it is not in the new store too
// // 				}
// // 			}
// // 		}
// // 		fmt.Println("new Store: ", newStore)
// // 		fmt.Println("reverseNewStore: ", reverseNewStore)
// // 		fmt.Println("getOwnerProf: ", getOwnerProf)
// // 		token1 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][0]]}
// // 		token2 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][1]]}

// // 		token3 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][0]]}
// // 		token4 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][1]]}

// // 		__RED := "Display: " + token1.NickName + " " + token1.TeamName
// // 		__RED2 := "Display: " + token2.NickName + " " + token2.TeamName
// // 		__BLUE := "Display: " + token3.NickName + " " + token3.TeamName
// // 		__BLUE2 := "Display: " + token4.NickName + " " + token4.TeamName

// // 		h.broadcast <- BroadcastReq{Token: __RED, RoomId: l.To}
// // 		h.broadcast <- BroadcastReq{Token: __RED2, RoomId: l.To}

// // 		h.broadcast <- BroadcastReq{Token: __BLUE, RoomId: l.To}
// // 		h.broadcast <- BroadcastReq{Token: __BLUE2, RoomId: l.To}
// // 	}
// // }
// // proceed to voting session
// // case GR.MutalCount == 2:
// // 	// check if the team has challenge session true

// // 	// if the players agreed the voted value will be set for the challenge
// // 	fmt.Println("players agreed")

// // 	switch true {
// // 	case isTeamRed:

// // 		fmt.Println("team a sending token", mcout)

// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			roomName: l.To, token: _votingSession, to: saveShuffle[___TeamRedKey][0], _sleep: false,
// // 		}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			roomName: l.To, token: _votingSession, to: saveShuffle[___TeamRedKey][1], _sleep: false,
// // 		}

// // 		// send back that the session is over
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			roomName: l.To, token: mutalSession, to: saveShuffle[___TeamRedKey][0], _sleep: false,
// // 		}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			roomName: l.To, token: mutalSession, to: saveShuffle[___TeamRedKey][1], _sleep: false,
// // 		}
// // 		if getMWatch[___TeamBlueKey] {
// // 			proceed = true
// // 		}

// // 		if proceed {
// // 			fmt.Println("proceed")
// // 			if getSettings[l.To][___SetupToss] {
// // 				h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}

// // 				h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // 			} else {
// // 				if GR.Round == 1 {
// // 					fmt.Println("sending round 1 toss")
// // 					h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}

// // 					h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // 				} else {
// // 					fmt.Println("not sending round 1 toss")
// // 					h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}
// // 				}
// // 			}
// // 		}
// // 		store := map[string]map[string]bool{}
// // 		store[l.To] = map[string]bool{
// // 			___AgressSession: true,
// // 		}
// // 		TTrackSessionDone <- store

// // 	case isTeamBlue:

// // 		fmt.Println("team b sending token", mcout)

// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			roomName: l.To, token: _votingSession, to: saveShuffle[___TeamBlueKey][0], _sleep: false,
// // 		}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			roomName: l.To, token: _votingSession, to: saveShuffle[___TeamBlueKey][1], _sleep: false,
// // 		}
// // 		// send back that the session is over
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			roomName: l.To, token: mutalSession, to: saveShuffle[___TeamBlueKey][0], _sleep: false,
// // 		}
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{
// // 			roomName: l.To, token: mutalSession, to: saveShuffle[___TeamBlueKey][1], _sleep: false,
// // 		}

// // 		if getMWatch[___TeamRedKey] {
// // 			proceed = true
// // 		}

// // 		if proceed {
// // 			fmt.Println("proceed")
// // 			if getSettings[l.To][___SetupToss] {
// // 				h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}

// // 				h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // 			} else {
// // 				if GR.Round == 1 {
// // 					fmt.Println("sending round 1 toss")
// // 					h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}

// // 					h.broadcast <- BroadcastReq{RoomId: l.To, Token: _tossSession}
// // 				} else {
// // 					fmt.Println("not sending round 1 toss")
// // 					h.broadcast <- BroadcastReq{RoomId: l.To, Token: _DictionaryDiscussion}
// // 				}
// // 			}

// // 			store := map[string]map[string]bool{}
// // 			store[l.To] = map[string]bool{
// // 				___AgressSession: true,
// // 			}
// // 			TTrackSessionDone <- store

// // 		}
// // 	default:
// // 		fmt.Println("none case mutal")
// // 	}
// // h.wg.Add(1)
// // go func() {
// // 	defer h.wg.Done()
// // 	// set challenge discussion over
// // 	h.broadcast <- BroadcastReq{RoomId: l.To, Token: votingSession}
// // 	newStore := map[string]string{}
// // 	reverseNewStore := map[string]string{}
// // 	for _, tag := range saveShuffle {
// // 		for _, data := range tag {
// // 			for Name, ID := range getProfData {
// // 				if ID == data {
// // 					newStore[Name] = data
// // 					reverseNewStore[data] = Name
// // 				}
// // 			}
// // 		}
// // 	}
// // 	// becuase we have saved the data of joiners separately
// // 	for _, id := range reverseNewStore {
// // 		for Ownername, Ownerid := range getOwnerProf {
// // 			if Ownerid != id {
// // 				reverseNewStore[Ownerid] = Ownername
// // 				newStore[Ownername] = Ownerid // its obvious that if the id doenst match it means that it is not in the new store too
// // 			}
// // 		}
// // 	}

// // 	token1 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][0]]}
// // 	token2 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][1]]}

// // 	token3 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][0]]}
// // 	token4 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][1]]}

// // 	__RED := "Display: " + token1.NickName + " " + token1.TeamName
// // 	__RED2 := "Display: " + token2.NickName + " " + token2.TeamName
// // 	__BLUE := "Display: " + token3.NickName + " " + token3.TeamName
// // 	__BLUE2 := "Display: " + token4.NickName + " " + token4.TeamName

// // 	h.broadcast <- BroadcastReq{Token: __RED, RoomId: l.To}
// // 	h.broadcast <- BroadcastReq{Token: __RED2, RoomId: l.To}

// // 	h.broadcast <- BroadcastReq{Token: __BLUE, RoomId: l.To}
// // 	h.broadcast <- BroadcastReq{Token: __BLUE2, RoomId: l.To}
// // 	h.broadcast <- BroadcastReq{Token: waiting, RoomId: l.To}

// // 	fmt.Println("new Store: ", newStore)
// // 	fmt.Println("reverseNewStore: ", reverseNewStore)
// // 	fmt.Println("getOwnerProf: ", getOwnerProf)

// // 	// clear cache
// // 	Cwatch.done <- true
// // }()

// // h.wg.Add(1)
// // go func() {
// // 	defer h.wg.Done()

// // 	token := "ChallengeGuess: " + __getChallengeSet[___TeamRedKey]

// // 	for _, r := range saveShuffle[___TeamBlueKey] {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: r, roomName: l.To, _sleep: false}
// // 	}

// // 	token2 := "ChallengeGuess: " + GR.ChallengeToken
// // 	to := saveShuffle[___TeamBlueKey][0]
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: to, roomName: l.To, _sleep: false}
// // 	to2 := saveShuffle[___TeamBlueKey][1]
// // 	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: to2, roomName: l.To, _sleep: false}

// // 	fmt.Println("challenge set: ", __getChallengeSet)

// // 	_trackChallengeSet.done <- true
// // }()

// // 	for _, ids := range saveShuffle[___TeamBlueKey] {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, to: ids, roomName: l.To, _sleep: false}
// // 	}
// // 	for _, ids := range saveShuffle[___TeamRedKey] {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, to: ids, roomName: l.To, _sleep: false}
// // 	}
// // }
// // if getCWatch[___TeamRedKey] {
// // 	proceed = true
// // } else {
// // 	proceed = false
// // }
// // if proceed {
// // 	h.broadcast <- BroadcastReq{RoomId: l.To, Token: votingSession}
// // 	h.broadcast <- BroadcastReq{RoomId: l.To, Token: _startGame}

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()
// // 		// set challenge discussion over
// // 		h.broadcast <- BroadcastReq{RoomId: l.To, Token: votingSession}
// // 		newStore := map[string]string{}
// // 		reverseNewStore := map[string]string{}
// // 		for _, tag := range saveShuffle {
// // 			for _, data := range tag {
// // 				for Name, ID := range getProfData {
// // 					if ID == data {
// // 						newStore[Name] = data
// // 						reverseNewStore[data] = Name
// // 					}
// // 				}
// // 			}
// // 		}
// // 		// becuase we have saved the data of joiners separately
// // 		for _, id := range reverseNewStore {
// // 			for Ownername, Ownerid := range getOwnerProf {
// // 				if Ownerid != id {
// // 					reverseNewStore[Ownerid] = Ownername
// // 					newStore[Ownername] = Ownerid // its obvious that if the id doenst match it means that it is not in the new store too
// // 				}
// // 			}
// // 		}

// // 		token1 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][0]]}
// // 		token2 := TeamData{TeamName: "Team Red", NickName: reverseNewStore[_______REDTeam["RED"][1]]}

// // 		token3 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][0]]}
// // 		token4 := TeamData{TeamName: "Team Blue", NickName: reverseNewStore[_______BLUEteam["BLUE"][1]]}

// // 		__RED := "Display: " + token1.NickName + " " + token1.TeamName
// // 		__RED2 := "Display: " + token2.NickName + " " + token2.TeamName
// // 		__BLUE := "Display: " + token3.NickName + " " + token3.TeamName
// // 		__BLUE2 := "Display: " + token4.NickName + " " + token4.TeamName

// // 		h.broadcast <- BroadcastReq{Token: __RED, RoomId: l.To}
// // 		h.broadcast <- BroadcastReq{Token: __RED2, RoomId: l.To}

// // 		h.broadcast <- BroadcastReq{Token: __BLUE, RoomId: l.To}
// // 		h.broadcast <- BroadcastReq{Token: __BLUE2, RoomId: l.To}
// // 		h.broadcast <- BroadcastReq{Token: waiting, RoomId: l.To}

// // 		fmt.Println("new Store: ", newStore)
// // 		fmt.Println("reverseNewStore: ", reverseNewStore)
// // 		fmt.Println("getOwnerProf: ", getOwnerProf)

// // 		// clear cache
// // 		Cwatch.done <- true
// // 	}()

// // 	h.wg.Add(1)
// // 	go func() {
// // 		defer h.wg.Done()

// // 		token := "ChallengeGuess: " + __getChallengeSet[___TeamRedKey]

// // 		for _, r := range saveShuffle[___TeamRedKey] {
// // 			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: r, roomName: l.To, _sleep: false}
// // 		}

// // 		token2 := "ChallengeGuess: " + GR.ChallengeToken
// // 		to := saveShuffle[___TeamBlueKey][0]
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: to, roomName: l.To, _sleep: false}
// // 		to2 := saveShuffle[___TeamBlueKey][1]
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token2, to: to2, roomName: l.To, _sleep: false}

// // 		fmt.Println("challenge set: ", __getChallengeSet)

// // 		_trackChallengeSet.done <- true
// // 	}()
// // } else {
// // 	for _, ids := range saveShuffle[___TeamRedKey] {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, to: ids, roomName: l.To, _sleep: false}
// // 	}
// // 	for _, ids := range saveShuffle[___TeamBlueKey] {
// // 		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, to: ids, roomName: l.To, _sleep: false}
// // 	}
// // }

// // package main

// // import (
// // 	"github.com/google/uuid"
// // )

// // var (
// // 	approved   = false
// // 	namespaces = []string{"/ws/:id", "/ws/:id/:room"}
// // 	Conns      = ""         // active connections
// // 	Host       = ""         // id of the host
// // 	Cap        = ""         // capacity of the room
// // 	HostFor    = ""         // name of the room
// // 	Hosts      = []string{} // list of rooms
// // 	Point      = ""         // experimental: a signal to give the point to
// // 	Items      = []string{}
// // )

// // type Lobby struct {
// // 	JoinRoom    bool   `json:"joinRoom"`    // if true
// // 	To          string `json:"to"`          // then join to this
// // 	CreateRoom  bool   `json:"createRoom"`  // if true
// // 	RoomCreated bool   `json:"roomCreated"` // if true than they can modify the room settings

// // 	/*room data*/
// // 	Category string `json:"category"` // set of category  [sports or entertainment]
// // 	Field    string `json:"field"`    // set of field for [international, national, domestic]
// // 	Book     string `json:"book"`     // name of the book set for the room [for example cricket, basketball, tennis...]
// // 	/**end of room data*/
// // 	Reverse        bool   `json:"reverse"`      // if to set the challenge dictionary as A1-B1 or A1-A1 [meaning: A1-B1 the one who sets the challenge dictionary the other one can first set the challenge than other, A1-A1 the one who sets the challenge dictionary than he  can first set the cahllenge than other]
// // 	RoomName       string `json:"roomName"`     // then create this room
// // 	RoomCapacity   int    `json:"roomCapacity"` // for 1v1 is 2 and for 2v2 is 4
// // 	Friend         bool   `json:"friend"`       // if the client is joint via code
// // 	Code           string `json:"code"`         // if friend true than join via code
// // 	NickName       string `json:"nickname"`     // to display to all connected user
// // 	GameTime       string `json:"gameTime"`
// // 	DecisionTime   string `json:"decisionTime"`
// // 	Set            bool   `json:"set"`            // required else the nickname wont be saved
// // 	ID             string `json:"id"`             // generated or saved id
// // 	SetToss        bool   `json:"setToss"`        // if to involve toss session
// // 	Starter        bool   `json:"starter"`        // if to involve starter system[winners first or lossers first]
// // 	PrivateRoom    bool   `json:"privateRoom"`    // if the room request is for the private room
// // 	ChangeSettings bool   `json:"changeSettings"` // to change the settings
// // 	NexusPower     bool   `json:"nexusPower"`     // sends the word from the dictionary but in mazed order [V_R A_]
// // 	TagPower       bool   `json:"tagPower"`       // locks the partner in-place of the chosen power
// // 	RewindPower    bool   `json:"rewindPower"`    // reset's the clock once the output of the opponent been used
// // 	FreezePower    bool   `json:"freezePower"`    // gives the control to start the clock under specific condition
// // 	DrawPower      bool   `json:"drawPower"`      // sends the request to the player to tag in their partner's inplace of them
// // 	CovertPower    bool   `json:"covertPower"`    // hides the text when the opponent writes something
// // 	// any one can use this bet power [if the dictionary setter or non-dictionary setter]
// // 	BetPower bool `json:"betPower"` // player get's the list of the current event and they ought to bet on any one word of the event; note: they cannot guess

// // }

// // type GameRoom struct {

// // 	// mutal chatting: both players will vote to mutally decide for the challenge token
// // 	// they wouldnt be able to chat rather they can pick
// // 	// if not pick under time than the battle ground player can pick for the win
// // 	// if even the battle ground player wouldnt be able to pick the point goes to the opponent team
// // 	// if both players weren't able to pick
// // 	// the warning will be given
// // 	// even after the warning the players wont pick the room will be closed

// // 	MutalCount      int `json:"mutalCount"`      // agress for the discussion
// // 	HomeScoreCount  int `json:"homeScoreCount"`  // current home score; note: it will always be less than 1 than current meaning if the score board is 2 than then score count send here will be 1
// // 	AwayScoreCount  int `json:"awayScoreCount"`  // current away score
// // 	MutualVoteCount int `json:"mutualVoteCount"` // tracking the current voting
// // 	Chances         int `json:"chanceLeft"`      // track the current chance left

// // 	RoomName        string `json:"roomName"`        // name of the room
// // 	ChallengeToken  string `json:"challengeToken"`  // token to challenge the rival
// // 	MutualVote      string `json:"mutualVote"`      // vote to set the challenge value
// // 	Guess           string `json:"guessToken"`      // guess the token as per the given guessin category
// // 	HeadTails       string `json:"headTails"`       // chosen side
// // 	DictionaryToken string `json:"dictionaryToken"` // chosen dictionary word

// // 	ChallengeCount int `json:"challegneCount"` // track if both of the players has set the challenge

// // 	MutualSession       bool `json:"mutalSession"`        // on going vote agreeing for mutual voting
// // 	Session             bool `json:"session"`             // on going game session
// // 	Start               bool `json:"start"`               // is the game has started
// // 	Agree               bool `json:"agree"`               // if the player mutually decided the token
// // 	TossSession         bool `json:"tossSession"`         // for toss session
// // 	VotingSession       bool `json:"votingSession"`       // on going voting session
// // 	ChallengeSet        bool `json:"challengeSet"`        // same as Agree
// // 	DictionarySession   bool `json:"dictionarySession"`   // on going dictionary session for setting challenge
// // 	ChallengeDiscussion bool `json:"challengeDiscussion"` // is in the challenge discussion session
// // 	TimeUp              bool `json:"timeUp"`              // if the player loss by time
// // 	PowerActivated      bool `json:"power"`               // if the player has decided to use the power

// // 	/** players attribute */
// // 	Freeze bool `json:"freeze"`
// // 	Nexus  bool `json:"nexus"`
// // 	Rewind bool `json:"Rewind"`
// // 	Tag    bool `json:"tag"`
// // 	Draw   bool `json:"draw"`
// // 	Covert bool `json:"covert"`
// // 	Bet    bool `json:"bet"`
// // 	/** end of players attribute **/

// // 	Unfreeze bool `json:"unfreeze"`

// // 	DrawSession bool   `json:"drawSession"` // this will be key to accept and reject the draw offer
// // 	DrawAccept  bool   `json:"drawAccept"`
// // 	Set         int    `json:"set"`        // current set
// // 	Round       int    `json:"round"`      // current Round
// // 	BetSession  bool   `json:"betSession"` // this will be the key for storing the bet value
// // 	BetOn       string `json:"betOn"`      // storing the bet value
// // }

// // var (
// // 	code = GenCode() // random code generated for a friend to join the group
// // )

// // func GenCode() string {
// // 	_gen := uuid.New().String()[0:4]
// // 	return _gen
// // }

// // type TeamData struct {
// // 	TeamName string
// // 	NickName string
// // }

// // var (
// // 	l      Lobby
// // 	GR     GameRoom
// // 	latest bool
// // )
