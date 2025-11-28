package server

import (
	"log"
	"math/rand"
	"strconv"
)

// LocifyToss returns to do  make sure to implement more complex later
// for Locify toss it is 4
func LocifyToss(h *Hub, id string, roomname string, HeadTails string) {

	log.Println("in Locify toss")

	myID := id
	myOpponentID := getLocifyProfile[myID][roomname].Against
	myProfile := getLocifyProfile[id][roomname]
	winnerTeam := myProfile.MyTeam
	losserTeam := myProfile.OppoTeamname

	// if both the team has done with the session
	sessionDone := getLocifyProfile[myID][roomname].OppoTossDone

	tossWon := TossMsg + "YOU WON NOW YOU CAN SET THE DICTIONARY"
	tossLost := TossMsg + "YOU LOST THE TOSS"

	Coin := []string{"HEADS", "TAILS"}

	rand.Shuffle(len(Coin), func(i int, j int) {
		Coin[i], Coin[j] = Coin[j], Coin[i]
	})

	myTeam := getLocifyProfile[myID][roomname].MyTeam
	// the sent req won the toss
	IWon := getLocifyProfile[myID][roomname].OppoTossCalled == Coin[0]
	_body := "Heads: " + strconv.FormatBool(Coin[0] == "HEAD")
	log.Println("game: ", myID, myOpponentID)
	log.Println("team: ", myTeam, "oppo team: ", myProfile.OppoTeamname)
	log.Println("my team squad: ", myProfile.MyID, myProfile.IdlePlayer)
	log.Println("oppo team squad: ", getLocifyProfile[myOpponentID][roomname].MyID, getLocifyProfile[myOpponentID][roomname].MyPatnerID)
	winnerID := id
	losserID := myOpponentID

	// to set the face of the coin and to display the result after the coin toss
	switch true {
	case sessionDone:
		log.Println("next")
		book := getLocifyRoomSettings[roomname].Book
		log.Println("book: ", book)
		_, li := SendDictionary(book)
		sendDictionary := DictionaryURL + li
		if !IWon {
			winnerTeam = myProfile.OppoTeamname
			losserTeam = myProfile.MyTeam
			winnerID = myOpponentID
			losserID = myID
		}

		for _, _id := range saveShuffle[roomname][winnerTeam] {

			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: _id, _sleep: false, token: Unblock,
			}

			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: _id, _sleep: false, token: tossWon,
			}
		}

		for _, _id := range saveShuffle[roomname][losserTeam] {

			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: _id, _sleep: false, token: Block,
			}

			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				roomname: roomname, to: _id, _sleep: false, token: tossLost,
			}
		}
		LocifyTossDoneMessage(h, roomname, losserID, winnerID, winnerTeam)

		h.broadcast <- BroadcastReq{Token: TossAlert, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: tossSession, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: sendDictionary, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: toss, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: tossCoin, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: _DictionaryDiscussion, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: waiting, RoomID: roomname}

	default:
		log.Println("default")
		LocifyTossMessage(h, roomname, id)
		SingleSave(false, myOpponentID, roomname, _StringSentinel_,
			HeadTails, _StringSentinel_,
			_StringSentinel_, nil, true, false, false, false)

		for _, _id := range saveShuffle[roomname][myTeam] {

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
		}

		for _, _id := range saveShuffle[roomname][myProfile.OppoTeamname] {
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
