package server

import (
	"log"
	"math/rand"
	"strconv"
)

// ClashToss returns to do  make sure to implement more complex later
// for clash toss it is 4
func ClashToss(h *Hub, id string, roomname string, HeadTails string, MaxPlayers int) {

	log.Println("in clash toss")

	myID := id
	myOpponentID := getClashProfile[myID][roomname].Against
	Count := int(1)
	winnerID := myID
	losserID := myOpponentID
	// if the my opp has done with toss
	proceed := getClashProfile[myID][roomname].OppoTossDone

	// if both the team has done with the session
	sessionDone := (getClashProfile[myID][roomname].Count) == int(MaxPlayers)

	log.Println("max player: ", MaxPlayers)

	youWon := TossMsg + "YOU WON NOW YOU CAN SET THE DICTIONARY"
	tossLost := TossMsg + "YOU LOST THE TOSS"

	Coin := []string{"HEADS", "TAILS"}

	rand.Shuffle(len(Coin), func(i int, j int) {
		Coin[i], Coin[j] = Coin[j], Coin[i]
	})

	log.Println("my profile: ", getClashProfile[id][roomname])
	log.Println("session done: ", sessionDone)
	log.Println("proceed: ", proceed)

	myTeam := getClashProfile[myID][roomname].MyTeam
	// the sent req won the toss
	IWon := getClashProfile[myID][roomname].OppoTossCalled == Coin[0]
	_body := "Heads: " + strconv.FormatBool(Coin[0] == "HEAD")

	// to set the face of the coin and to display the result after the coin toss
	log.Println("coin:", Coin[0], "called: ", getClashProfile[myID][roomname].OppoTossCalled)
	switch true {
	case sessionDone:
		log.Println("next")
		book := getClashProfile[id][roomname].Book
		_, li := SendDictionary(book)
		sendDictionary := DictionaryURL + li
		nextTeamWinnerID := getClashProfile[myID][roomname].NextTeamWinner
		nextTeamname := getClashProfile[myID][roomname].NextTeamname

		for _, _myID := range saveShuffle[roomname][nextTeamname] {
			if _myID != nextTeamWinnerID {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{
					roomname: roomname, to: _myID, _sleep: false, token: tossLost,
				}
			} else {

				h.gameRoomBroadcast <- reqGameRoomBroadcast{
					roomname: roomname, to: _myID, _sleep: false, token: youWon,
				}
			}
		}
		if IWon {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myOpponentID, _sleep: false, token: Block,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myOpponentID, _sleep: false, token: tossLost,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myID, _sleep: false, token: Unblock,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myID, _sleep: false, token: youWon,
			}
		} else {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myID, _sleep: false, token: Block,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myID, _sleep: false, token: tossLost,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myOpponentID, _sleep: false, token: Unblock,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myOpponentID, _sleep: false, token: youWon,
			}
		}

		h.broadcast <- BroadcastReq{Token: TossAlert, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: tossSession, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: sendDictionary, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: toss, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: tossCoin, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: _DictionaryDiscussion, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: waiting, RoomID: roomname}

		// to do reset the count
		Count = 0 // reset the count
		re := map[string]bool{}
		re[roomname] = true
		clashResetCount <- re

	case proceed:
		log.Println("proceed")

		if IWon {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, token: _DictionaryDiscussion, to: myID, _sleep: false,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myID, _sleep: false, token: Unblock,
			}
			// winner
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myOpponentID, _sleep: false, token: Block,
			}

		} else {
			winnerID = myOpponentID
			losserID = myID
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myID, _sleep: false, token: Block,
			}
			// winnner
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, token: _DictionaryDiscussion, to: myOpponentID, _sleep: false,
			}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: myOpponentID, _sleep: false, token: Unblock,
			}
			winnerID = myOpponentID
		}

		for _, _id := range saveShuffle[roomname][myTeam] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _waiting, roomname: roomname,
				_sleep: false, to: _id}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: _id,
				_sleep: false, token: _ClashWaitMessgae}

		}

		MasterSave(true, int(Count),
			_StringSentinel_, winnerID,
			losserID, roomname,
			myTeam, _StringSentinel_,
			_StringSentinel_, _StringSentinel_,
			false, false, _IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, nil, nil)
		CommonSave(true, roomname, myTeam, nil, true, false, false, false)
		SingleSave(true, myOpponentID, roomname, myTeam, HeadTails, _StringSentinel_, _StringSentinel_, nil, true, false, false, false)

	default:
		log.Println("default")

		MasterSave(true, int(Count),
			_StringSentinel_,
			_StringSentinel_,
			_StringSentinel_,
			roomname, myTeam,
			_StringSentinel_,
			_StringSentinel_,
			_StringSentinel_,
			false, false, _IntSentinel, _IntSentinel, _StringSentinel_, _StringSentinel_, nil, nil)
		SingleSave(true, myOpponentID, roomname, myTeam, HeadTails, _StringSentinel_, _StringSentinel_, nil, true, false, false, false)

		for _, _id := range saveShuffle[roomname][myTeam] {
			if myID == _id {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{
					roomname: roomname, to: _id, _sleep: false, token: Block,
				}
				// toss pick done
				h.gameRoomBroadcast <- reqGameRoomBroadcast{
					roomname: roomname, to: _id, _sleep: false, token: toss,
				}
				h.gameRoomBroadcast <- reqGameRoomBroadcast{
					roomname: roomname, to: _id, _sleep: false, token: tossSession,
				}

			} else {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{
					roomname: roomname, to: _id, _sleep: false, token: Unblock,
				}
				// toss
				h.gameRoomBroadcast <- reqGameRoomBroadcast{
					roomname: roomname, to: _id, _sleep: false, token: _toss,
				}
				h.gameRoomBroadcast <- reqGameRoomBroadcast{
					roomname: roomname, to: _id, _sleep: false, token: _tossCoin,
				}
			}
		}

		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: _body, _sleep: false,
			roomname: roomname, to: myOpponentID,
		}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token: _body, _sleep: false,
			roomname: roomname, to: myID,
		}
	}

}
