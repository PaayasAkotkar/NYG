package server

import (
	"log"
	"nyg/dataset"
)

func RoomHub(h *Hub) {

	for {
		select {

		// registers conn
		case connection := <-h.register:

			client := &roomClient{conn: connection, clientID: ""}
			h.clients[connection] = client
			log.Println("registered done")

			// unregisters the rooms and client conn
		case connection := <-h.unregister:
			// if client, ok := h.clients[connection]; ok {
			log.Println("deleting client")
			// UnregisterRoom(h, connection, client.roomID)
			UnregisterRoom(h, connection)

		case client := <-h.unregisterID:
			log.Println("unregsitering client")
			UnregisterClientFromRoom(h, client.Roomname, client.ID)
		case conn := <-h.unregisterRoomSession:
			for _conn, det := range conn {
				UnregisterRoomSession(h, _conn, det.RoomName, det.ID)
			}
		case token := <-h.Cbroadcast:
			BroadcastAll(h, token.Token)

		case token := <-h.broadcast:
			log.Println("broadcast to all room clients")

			BroadcastToRoom(h, token.RoomID, string(token.Token))

		case token := <-h.gameRoomBroadcast:

			// GameRoomBroadcast(h, token.roomname, token.token, token.to, token._sleep)
			GameBroadcastToRoom(h, token.roomname, token.token, token.to)

		case store := <-cRommstore.rooms:
			for roomname, id := range store {
				if _, ok := cRoom[roomname]; !ok {
					cRoom[roomname] = []string{}
				}

				cRoom[roomname] = append(cRoom[roomname], id)

			}
			log.Println("STORED ASSOCIATED CLIENT: ", cRoom)
			log.Println("STORING ROOM: ", store)

		case storeTeams := <-storeShuffle.store:
			// maps.Copy(saveShuffle, storeTeams)
			for room, teamsIDs := range storeTeams {

				for teamname, ids := range teamsIDs {
					if _, ok := saveShuffle[room]; !ok {
						saveShuffle[room] = map[string][]string{}
					}
					if _, ok := saveShuffle[room][teamname]; !ok {
						saveShuffle[room][teamname] = append(saveShuffle[room][teamname], ids...)
					} else if _, ok := saveShuffle[room][teamname]; ok { // safety
						saveShuffle[room][teamname] = storeTeams[room][teamname]
					}
				}
			}

		case token := <-storeClashTokens:
			for lobby, clients := range token.storeClient {
				for _, id := range clients {
					slice_, exists := getClashTokens.storeClient[lobby]
					if !exists {
						log.Println("not exists ")
						getClashTokens.storeClient[lobby] = []string{id}
						getClashTokens.storeLobbyKeys = append(getClashTokens.storeLobbyKeys, lobby)
					} else {
						log.Println("exists ")
						getClashTokens.storeClient[lobby] = append(slice_, id)
					}
				}
			}
			log.Println("STOED DONE: ", getClashTokens)

		case token := <-createProfile:
			for id, roomsDets := range token {
				for roomname, det := range roomsDets {
					if det.Clash {
						if _, ok := getClashProfile[id]; !ok {
							getClashProfile[id] = map[string]ClashFixtures{}
						}
						if _, ok := getClashProfile[id][roomname]; !ok {
							getClashProfile[id][roomname] = det
						} else if _, ok := getClashProfile[id][roomname]; ok {
							getClashProfile[id][roomname] = det
						}
					}
				}
			}

		case token := <-saveRoomSettings:
			for room, paste := range token {
				if _, ok := getLocifyRoomSettings[room]; ok {
					src := getLocifyRoomSettings[room]
					src.Book = paste.Book
					src.DecisionTime = paste.DTime
					src.GameTime = paste.GTime
					src.SetupToss = paste.SetupToss
					src.Reverse = paste.Reverse
					src.Starter = paste.Starter
					src.Powers = paste.Powers
					getLocifyRoomSettings[room] = src
				}
			}
		case token := <-clashResetCount:
			for roomname, reset := range token {
				if reset {
					for id := range getClashProfile {
						if _, ok := getClashProfile[id][roomname]; ok {
							var temp = getClashProfile[id][roomname]
							temp.Count = 1
							getClashProfile[id][roomname] = temp
						}
					}
				}
			}

		case token := <-storeClashNicknames:
			for roomname, idsNicknames := range token {
				for id, nickname := range idsNicknames {
					if _, ok := getClashNicknames[roomname]; !ok {
						getClashNicknames[roomname] = map[string]string{}
					}
					if _, ok := getClashNicknames[roomname][id]; !ok {
						getClashNicknames[roomname][id] = nickname
					}
				}
			}

		// case token := <-storeFinalBoss:

		// 	for id, roomsDet := range getClashProfile {
		// 		for roomname := range roomsDet {
		// 			if _, ok := token[roomname]; ok {
		// 				var temp = getClashProfile[id][roomname]
		// 				temp.FinalBossID = token[roomname]
		// 				temp.FinalBossFound = true
		// 				getClashProfile[id][roomname] = temp
		// 				getFinalBoss[roomname] = id
		// 			}
		// 		}
		// 	}

		case token := <-saveGlobal:
			log.Println("saving global")
			for roomname, det := range token {
				log.Println("bets: ", det.Bets)
				if det.Clash {
					for id := range getClashProfile {

						var src = getClashProfile[id][roomname]
						var paste = det
						if paste.Book != _StringSentinel_ {
							src.Book = paste.Book
						}

						// dont use if else statment becuase they are not related

						if paste.FinalBossFound {
							src.FinalBossFound = paste.FinalBossFound
						}

						if paste.LastDance {
							src.LastDance = paste.LastDance
						}

						if paste.LosserID != _StringSentinel_ {
							src.NextTeamLosser = paste.LosserID
						}
						if paste.WinnerID != _StringSentinel_ {
							src.NextTeamWinner = paste.WinnerID
						}
						if paste.WinnerTeamName != _StringSentinel_ {
							src.NextTeamname = paste.WinnerTeamName
						}
						if paste.FinalBossID != _StringSentinel_ {
							src.FinalBossID = paste.FinalBossID
						}
						if paste.Count != _IntSentinel {
							src.Count += det.Count
						}
						if paste.BetID != _StringSentinel_ {
							src.BetIDs = append(src.BetIDs, paste.BetID)
						}
						if paste.EliminatedPlayerID != _StringSentinel_ {
							src.EliminatedPlayersIDs = append(src.EliminatedPlayersIDs, paste.EliminatedPlayerID)
						}
						if paste.PrevList != nil {
							src.WholeGuess = append(src.WholeGuess, paste.PrevList...)
							src.WholeGuess = dataset.EraseDuplicate(src.WholeGuess)
						}
						if paste.Bets != nil {
							log.Println("storing bets: ", paste.Bets)
							src.OppoBets = append(src.OppoBets, paste.Bets...)
						}
						getClashProfile[id][roomname] = src

					}
				} else {
					for id := range getLocifyProfile {
						var src = getLocifyProfile[id][roomname]
						var paste = det

						if paste.Book != _StringSentinel_ {
							src.Book = paste.Book
						}
						if paste.Bets != nil {
							log.Println("storing bets: ", paste.Bets)
							src.OppoBets = append(src.OppoBets, paste.Bets...)
						}
						if paste.BetID != _StringSentinel_ {
							src.BetIDs = append(src.BetIDs, paste.BetID)
						}
						if paste.SetDictionary != _StringSentinel_ {
							src.SetDictionary = paste.SetDictionary
						}
						if paste.BlueTeamScore != _IntSentinel {
							src.BlueTeamScore += det.BlueTeamScore
						}
						if paste.RedTeamScore != _IntSentinel {
							src.RedTeamScore += det.RedTeamScore
						}
						if paste.PrevList != nil {
							src.WholeGuess = append(src.WholeGuess, paste.PrevList...)
							src.WholeGuess = dataset.EraseDuplicate(src.WholeGuess)
						}
						getLocifyProfile[id][roomname] = src

					}
				}
			}

		case token := <-saveSingle:
			for id, roomsDet := range token {
				for roomname, paste := range roomsDet {
					if paste.Clash {
						if _, ok := getClashProfile[id][roomname]; ok {
							var src = getClashProfile[id][roomname]

							if src.MyTeam != _StringSentinel_ {
								src.MyTeam = paste.MyTeam
							}

							if paste.PowerUp != nil {
								log.Println("saved power: ", paste.PowerUp)
								src.OppoPowerUp = paste.PowerUp
							}

							if paste.TossCalled != _StringSentinel_ {
								src.OppoTossCalled = paste.TossCalled
							}

							if paste.SetBet != _StringSentinel_ {
								src.OppoSetBet = paste.SetBet
							}
							if paste.SetChallenge != _StringSentinel_ {
								src.OppoSetChallenge = paste.SetChallenge
							}

							switch true {
							case paste.TossDone:
								src.OppoTossDone = paste.TossDone
							case paste.DictionaryDone:
								src.OppoDictionaryDone = paste.DictionaryDone
							case paste.BetDone:
								src.OppoBetDone = paste.BetDone
							case paste.ChallengeDone:
								src.OppoChallengeDone = paste.ChallengeDone
							}

							getClashProfile[id][roomname] = src
						}
					} else {
						if _, ok := getLocifyProfile[id][roomname]; ok {

							var src = getLocifyProfile[id][roomname]

							if paste.PowerUp != nil {
								log.Println("saved power: ", paste.PowerUp)
								src.OppoPowerUp = paste.PowerUp
							}

							if paste.TossCalled != _StringSentinel_ {
								src.OppoTossCalled = paste.TossCalled
							}

							if paste.SetBet != _StringSentinel_ {
								src.OppoSetBet = paste.SetBet
							}

							if paste.SetChallenge != _StringSentinel_ {
								src.OppoSetChallenge = paste.SetChallenge
							}

							switch true {
							case paste.TossDone:
								src.OppoTossDone = paste.TossDone
							case paste.DictionaryDone:
								src.OppoDictionaryDone = paste.DictionaryDone
							case paste.BetDone:
								src.OppoBetDone = paste.BetDone
							case paste.ChallengeDone:
								src.OppoChallengeDone = paste.ChallengeDone
							}

							getLocifyProfile[id][roomname] = src
						}
					}
				}
			}

		case token := <-saveStats:
			for id, roomsDet := range token {
				for roomname, paste := range roomsDet {
					// safety
					if paste.Clash {
						var src = getClashProfile[id][roomname]
						if paste.CurrentChances != _IntSentinel {
							src.MyCurrentChances = paste.CurrentChances
						}
						if paste.OnFire != _IntSentinel {
							src.ImOnFire += paste.OnFire
						}
						if paste.Guess != _StringSentinel_ {
							src.MyGuess = append(src.MyGuess, paste.Guess)
							src.WholeGuess = append(src.WholeGuess, paste.Guess)
						}
						if paste.PowersBin != _StringSentinel_ {
							src.MyPowerBin = append(src.MyPowerBin, paste.PowersBin)
						}
						if paste.Points != _IntSentinel {
							src.MyPoints += paste.Points
						}
						if paste.Penalty != nil {
							for round, pen := range paste.Penalty {
								if _, ok := src.MyPenalties[round]; !ok {
									src.MyPenalties = map[int]int{}

									src.MyPenalties[round] = pen
								}
							}
						}
						getClashProfile[id][roomname] = src
					} else {
						var src = getLocifyProfile[id][roomname]

						if paste.Guess != _StringSentinel_ {
							src.MyGuess = append(src.MyGuess, paste.Guess)
						}
						if paste.PowersBin != _StringSentinel_ {
							src.MyPowersBin = append(src.MyPowersBin, paste.PowersBin)
						}
						if paste.Points != _IntSentinel {
							src.MyPoints += paste.Points
						}
						if paste.Guess != _StringSentinel_ {
							src.MyGuess = append(src.MyGuess, paste.Guess)
							src.WholeGuess = append(src.WholeGuess, paste.Guess)
						}
						if paste.Penalty != nil {
							for round, pen := range paste.Penalty {
								if _, ok := src.MyPenalties[round]; !ok {
									src.MyPenalties = map[int]int{}
									src.MyPenalties[round] = pen
								}
							}
						}
						getLocifyProfile[id][roomname] = src
					}
				}
			}
		case token := <-saveCommon:
			for roomname, det := range token {
				if det.Clash {
					for id := range getClashProfile {
						var src = getClashProfile[id][roomname]
						if src.MyTeam == det.TeamName {
							if det.TossDone {
								src.OppoTossDone = det.TossDone
							}
							if det.BetDone {
								src.OppoBetDone = det.BetDone
							}
							if det.ChallengeDone {
								src.OppoChallengeDone = det.ChallengeDone
							}
							if det.DictionaryDone {
								src.OppoDictionaryDone = det.DictionaryDone
							}
							if det.Bets != nil {
								src.OppoBets = det.Bets
							}
							getClashProfile[id][roomname] = src
						}
					}
				} else {
					for id := range getLocifyProfile {
						var src = getLocifyProfile[id][roomname]
						if src.MyTeam == det.TeamName {
							if det.TossDone {
								src.OppoTossDone = det.TossDone
							}
							if det.BetDone {
								src.OppoBetDone = det.BetDone
							}
							if det.ChallengeDone {
								src.OppoChallengeDone = det.ChallengeDone
							}
							if det.DictionaryDone {
								src.OppoDictionaryDone = det.DictionaryDone
							}
							if det.Bets != nil {
								src.OppoBets = det.Bets
							}
							getLocifyProfile[id][roomname] = src
						}
					}
				}
			}
		case token := <-saveOneTime:
			for roomname, det := range token {
				if det.Clash {
					for id := range getClashProfile {
						var src = getClashProfile[id][roomname]
						if _, ok := getClashProfile[id][roomname]; ok && src.MyTeam == det.TeamName {
							if det.SetDictionary != _StringSentinel_ {
								src.SetDictionary = det.SetDictionary
							}
							getClashProfile[id][roomname] = src
						}
					}
				} else {
					for id := range getLocifyProfile {
						var src = getLocifyProfile[id][roomname]
						if _, ok := getClashProfile[id][roomname]; ok {
							if det.SetDictionary != _StringSentinel_ {
								src.SetDictionary = det.SetDictionary
							}
							getLocifyProfile[id][roomname] = src
						}
					}
				}
			}

		case token := <-clashresetPowerUp:
			for id, roomsReset := range token {
				for roomname, det := range roomsReset {
					log.Println("reset power: ", det.Key)
					if det.DoIt {
						if _, ok := getClashProfile[id][roomname]; ok {
							var re = getClashProfile[id][roomname]
							log.Println("reset power: ", det.Key)
							delete(re.OppoPowerUp, det.Key)
							getClashProfile[id][roomname] = re

						}
					}
				}
			}

		case token := <-locifyresetPowerUp:
			for id, roomsReset := range token {
				for room, det := range roomsReset {
					if det.DoIt {
						if _, ok := getLocifyProfile[id][room]; ok {
							var re = getLocifyProfile[id][room]
							log.Println("reset power: ", det.Key)
							delete(re.OppoPowerUp, det.Key)
							getLocifyProfile[id][room] = re
							var src = getLocifyProfile[re.Against][room]
							src.PowerDeck[det.Key] = false
							getLocifyProfile[re.Against][room] = src
							log.Println("after reset: ", getLocifyProfile[id][room].OppoPowerUp[det.Key])
							log.Println("power deck: ", getLocifyProfile[re.Against][room].PowerDeck)
						}
					}

				}
			}

		case token := <-updateProfile:
			for id, roomsDet := range token {
				for roomname, paste := range roomsDet {
					if paste.Clash {
						var src = getClashProfile[id][roomname]
						if paste.EliminatedPlayersIDs != nil {
							src.EliminatedPlayersIDs = paste.EliminatedPlayersIDs
						}
						if paste.FinalBossID != _StringSentinel_ {
							src.FinalBossID = paste.FinalBossID
							src.FinalBossFound = true
						}
						if paste.MyTeam != _StringSentinel_ {
							src.MyTeam = paste.MyTeam
						}
						if paste.Against != _StringSentinel_ {
							src.Against = paste.Against
						}
						if paste.CurrentRound != _IntSentinel {
							src.CurrentRound = paste.CurrentRound
						}
						src.OppoChallengeDone = false
						src.OppoDictionaryDone = false
						src.OppoTossDone = false
						src.OppoBetDone = false
						getClashProfile[id][roomname] = src
					} else {
						var src = getLocifyProfile[id][roomname]

						src.Against = paste.Against

						src.IBlock = paste.Block
						src.ILock = paste.Lock
						src.MyGuess = append(src.MyGuess, paste.WholeGuess)
						src.IdlePlayer = paste.MyPartnerID

						src.OppoChallengeDone = false
						src.OppoDictionaryDone = false
						src.OppoTossDone = false
						src.OppoBetDone = false
						if paste.CurrentRound != _IntSentinel {
							src.CurrentRound = paste.CurrentRound
						}
						if paste.UpdateStats != nil {
							for round, res := range paste.UpdateStats {

								if _, ok := src.MySheet[round]; !ok {
									src.MySheet[round] = res
								}
							}
						}

						getLocifyProfile[id][roomname] = src
					}
				}
			}

		case token := <-_resetPower:
			for id, roomsdets := range token {
				for room, paste := range roomsdets {
					if paste.Clash {
						var src = getClashProfile[id][room]
						src.OppoPowerUp[paste.Key] = false
						getClashProfile[id][room] = src
					} else {
						var src = getLocifyProfile[id][room]
						src.OppoPowerUp[paste.Key] = false
						getLocifyProfile[id][room] = src
					}
				}
			}
		case token := <-storeLocifyRoomSettings:
			log.Println("storing settings : ", token)
			for room, det := range token {

				if _, ok := getLocifyRoomSettings[room]; !ok {
					log.Println("creating settings for room: ", room)
					getLocifyRoomSettings[room] = det
				}
			}
			log.Println("STORED ROOM SETTINGS: ", getLocifyRoomSettings)

		case token := <-createLocifyProfile:
			log.Println("creating profile")
			for id, roomsDets := range token {
				for roomname, det := range roomsDets {
					if _, ok := getLocifyProfile[id]; !ok {
						getLocifyProfile[id] = map[string]LocifyFixtures{}
					}
					if _, ok := getLocifyProfile[id][roomname]; !ok {
						getLocifyProfile[id][roomname] = det
					} else if _, ok := getLocifyProfile[id][roomname]; ok {
						getLocifyProfile[id][roomname] = det
					}
					log.Println("ids: ", det.IDs)
				}
			}

		case token := <-friendJoin:
			for room, join := range token {
				if _, ok := getLocifyRoomSettings[room]; ok {
					var src = getLocifyRoomSettings[room]
					src.FriendJoin = join.Join
					src.FriendID = join.ID
					log.Println("before manipulating: ", getLocifyRoomSettings[room].Book)
					getLocifyRoomSettings[room] = src
					log.Println("friend has been register here is the book of the room: ",
						room, getLocifyRoomSettings[room].Book)
				}
			}
		}
	}
}
