package server

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"nyg/validate"
)

type ParcelFreezePackage struct {
	Unfreeze    bool `json:"unfreeze"`
	CanUnfreeze bool `json:"canUnfreeze"`
	FreezeUse   bool `json:"freezeUse"`
	FreezeTime  int  `json:"freezeTime"`
}
type AlertMessage struct {
	Message string `json:"message"`
	Alert   bool   `json:"alert"`
}

// PFreezeMechanism returns the delay of game
// freezeLevel must be the progress freeze upgrade of the requested id power
func PFreezeMechanism(h *Hub, roomname string, ID string, teamname string, isClash bool, both bool, freezeLevel int) {
	log.Println("freeze mechanism")
	/**
	* block's the time and hides the guess button
	* sends the control of pressing the button to the power's owner
	***/
	if !isClash {
		var t ParcelFreezePackage
		t.CanUnfreeze = true
		t.FreezeUse = true
		t.Unfreeze = false
		switch freezeLevel {
		case 1:
			t.FreezeTime = 7
		case 2:
			t.FreezeTime = 9
		case 3:
			t.FreezeTime = 12
		}
		c, _ := json.Marshal(t)
		token := "NYGFreeze: " + string(c)
		t.CanUnfreeze = false
		t.FreezeUse = true
		j, _ := json.Marshal(t)
		token_ := "NYGFreeze: " + string(j)
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, to: ID, _sleep: false, roomname: roomname}
		x := getLocifyProfile[ID][roomname].Against
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token_, to: x, _sleep: false, roomname: roomname}
		log.Println("freezing id: ", x, "control freeze id: ", ID)
	} else {

		if both {
			var t ParcelFreezePackage
			t.CanUnfreeze = true
			t.FreezeUse = true
			t.Unfreeze = false
			t.CanUnfreeze = false
			t.FreezeUse = true

			id1 := getClashProfile[ID][roomname].Against
			// send this to id1
			x := getClashProfile[ID][roomname].MyGameProfile.FreezeLevel

			switch x {
			case 1:
				t.FreezeTime = 7
			case 2:
				t.FreezeTime = 9
			case 3:
				t.FreezeTime = 12
			}

			t.CanUnfreeze = true
			jx, _ := json.Marshal(t)
			_token_ := "NYGFreeze: " + string(jx)

			// send this to id
			y := getClashProfile[id1][roomname].MyGameProfile.FreezeLevel

			switch y {
			case 1:
				t.FreezeTime = 7
			case 2:
				t.FreezeTime = 9
			case 3:
				t.FreezeTime = 12
			}
			t.CanUnfreeze = true
			jx, _ = json.Marshal(t)
			_token2 := "NYGFreeze: " + string(jx)

			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _token_, _sleep: false, to: id1, roomname: roomname}
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _token2, _sleep: false, to: ID, roomname: roomname}

		} else {
			var t ParcelFreezePackage
			t.CanUnfreeze = true
			t.FreezeUse = true
			t.Unfreeze = false
			switch freezeLevel {
			case 1:
				t.FreezeTime = 7
			case 2:
				t.FreezeTime = 9
			case 3:
				t.FreezeTime = 12
			}
			c, _ := json.Marshal(t)
			token := "NYGFreeze: " + string(c)

			c, _ = json.Marshal(t)
			t.CanUnfreeze = false
			token_ := "NYGFreeze: " + string(c)
			for _, id := range saveShuffle[roomname][teamname] {

				if id == ID {
					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token, _sleep: false, to: id, roomname: roomname}
				} else {
					h.gameRoomBroadcast <- reqGameRoomBroadcast{token: token_, _sleep: false, to: id, roomname: roomname}
				}
			}
		}
	}
}

// PCovertMechanism returns done only left is sending the stop request; signals to hide the text display
// id must be the id of the opponent team
// meaning if x used the covert than y must be the id to be passed
func PCovertMechanism(h *Hub, roomname string, ID string) {
	log.Println("covert mechanism", roomname)

	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: underTest, _sleep: false, to: ID, roomname: roomname}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, to: ID, _sleep: false, roomname: roomname}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CovertUse, to: ID, _sleep: false, roomname: roomname}

}

type IDraw struct {
	DrawOffer         bool   `json:"drawOffer"`
	DrawOfferDeclined bool   `json:"drawOfferDeclined"`
	Draw              bool   `json:"draw"`
	DrawMeetingDone   bool   `json:"drawMeetingDone"`
	DrawMsg           string `json:"drawMsg"`
}

// PDrawMechanism returns done only left is sending the stop request
func PDrawMechanism(h *Hub, roomname string, ID string, accept bool) {
	log.Println("in draw mechanism")
	myProfile := getLocifyProfile[ID][roomname]
	oppoTeam := myProfile.OppoTeamname
	myTeam := myProfile.MyTeam
	matchups := map[string]string{}
	if accept {

		// switch the players
		for _, _id := range saveShuffle[roomname][myTeam] {
			// h.gameRoomBroadcast <- reqGameRoomBroadcast{
			// 	token: drawOffer, _sleep: false, to: IDs, roomname: roomname,
			// }
			if !getLocifyProfile[_id][roomname].ILock {
				matchups[myTeam] = getLocifyProfile[_id][roomname].NickNamesViaID[_id]
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Lock, _sleep: false, to: _id, roomname: roomname}
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, _sleep: false, to: _id, roomname: roomname}
			} else {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unlock, _sleep: false, to: _id, roomname: roomname}
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, _sleep: false, to: _id, roomname: roomname}
			}
		}
		for _, _id := range saveShuffle[roomname][oppoTeam] {

			if !getLocifyProfile[_id][roomname].ILock {
				matchups[oppoTeam] = getLocifyProfile[_id][roomname].NickNamesViaID[_id]
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Lock, _sleep: false, to: _id, roomname: roomname}
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, _sleep: false, to: _id, roomname: roomname}
			} else {
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unlock, _sleep: false, to: _id, roomname: roomname}
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, _sleep: false, to: _id, roomname: roomname}
			}
		}

		msg := fmt.Sprintf("%s %s decided to agree on draw!!! a substitution from team %s %s checks in and from team %s %s checks in", myProfile.NickNamesViaID[myProfile.MyID], myProfile.NickNamesViaID[myProfile.Against], myTeam, myProfile.MyPatnerID, myProfile.OppoTeamname, getLocifyProfile[myProfile.Against][roomname].MyPatnerID)
		var t AlertMessage

		t.Alert = true
		t.Message = msg
		c, err := json.Marshal(t)
		if err != nil {
			log.Println(err)
			return
		}
		h.broadcast <- BroadcastReq{RoomID: roomname, Token: Message + string(c)}

		// update the profile
		p := getLocifyProfile[ID][roomname]
		o := getLocifyProfile[p.Against][roomname]

		UpdateProfile(false, ID, roomname, _StringSentinel_,
			myTeam, _StringSentinel_, nil, true, false,
			p.IdlePlayer, _StringSentinel_, nil, _IntSentinel)

		UpdateProfile(false, p.Against, roomname,
			_StringSentinel_, oppoTeam, _StringSentinel_,
			nil, true, false,
			o.IdlePlayer, _StringSentinel_, nil, _IntSentinel)

		UpdateProfile(false, o.IdlePlayer, roomname,
			p.IdlePlayer, oppoTeam, _StringSentinel_,
			nil, false, true,
			o.MyID, _StringSentinel_, nil, _IntSentinel)

		UpdateProfile(false, p.IdlePlayer, roomname,
			o.IdlePlayer, myTeam, _StringSentinel_,
			nil, false, true,
			p.MyID, _StringSentinel_, nil, _IntSentinel)

		h.broadcast <- BroadcastReq{Token: CanUsePower, RoomID: roomname}
		h.broadcast <- BroadcastReq{Token: DMsg + "Draw Offer Accepted", RoomID: roomname}

	} else {
		// send back the offer from whom the request was send for draw as declined
		for _, _id := range saveShuffle[roomname][myTeam] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: _decline, _sleep: false, to: _id, roomname: roomname}
		}
		msg := "draw offer declined"
		var t AlertMessage

		t.Alert = true
		t.Message = msg
		c, err := json.Marshal(t)
		if err != nil {
			log.Println(err)
			return
		}
		h.broadcast <- BroadcastReq{RoomID: roomname, Token: Message + string(c)}
	}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, to: ID, _sleep: false, roomname: roomname}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: DrawUse, to: ID, _sleep: false, roomname: roomname}
	h.broadcast <- BroadcastReq{Token: DrawMeetingDone, RoomID: roomname}
}

// PNexusMechanism returns only left is sending the stop request
func PNexusMechanism(h *Hub, roomname, ID, book, dictionay, ChallengeToken string, nexusLevel int) {
	log.Println("in nexus")
	dict := dictionay

	_token := ChallengeToken
	log.Println("dictionary ", dict, "challenge token: ", ChallengeToken)
	_conv := validate.Fetch(book, dict, _token).Pack

	list := _conv[dict][_token]

	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	log.Println("nexus word: ", list[0])

	max := false // meaning for both
	word := list[0]
	both := false
	switch nexusLevel {
	case 1:
		max = false
	case 2:
		max = false
		both = true
	case 3:
		max = true
		both = true
	}
	firstname, lastname := Nexus(word, "_", both, max)

	token := nexusWord + firstname
	if nexusLevel == 2 {
		token = nexusWord + firstname + lastname
	}

	log.Println("token: ", token)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{to: ID, roomname: roomname, _sleep: false, token: token}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, to: ID, _sleep: false, roomname: roomname}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: NexusUse, to: ID, _sleep: false, roomname: roomname}
}

// PRewindMechanism returns done only left is sending the stop request
// id must of the used id in-order to deactive the power from used
func PRewindMechanism(h *Hub, roomname, ID, teamName string, isClash bool) {
	if !isClash {
		h.broadcast <- BroadcastReq{
			Token: backClock, RoomID: roomname,
		}
	} else {
		for _, id := range saveShuffle[roomname][teamName] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: backClock, to: id, _sleep: false, roomname: roomname}
		}
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, to: ID, _sleep: false, roomname: roomname}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: RewindUse, to: ID, _sleep: false, roomname: roomname}
}

type TagMessage struct {
	TaggedIn  string `json:"taggedIn"`
	TaggedOut string `json:"taggedOut"`
}

// PTagMechanism returns tag done only left is sending stop request
// to do make sure to boardcast the alert the players
func PTagMechanism(h *Hub, roomname, ID string) {
	// switched, with_ := "", ""
	myProfile := getLocifyProfile[ID][roomname]
	myTeam := myProfile.MyTeam
	tagged := map[string]string{}
	for _, _id := range saveShuffle[roomname][myProfile.MyTeam] {
		if getLocifyProfile[_id][roomname].ILock {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				token: _tagIn, _sleep: false, to: _id, roomname: roomname,
			}
			tagged[TaggedOut] = getLocifyProfile[_id][roomname].NickNamesViaID[_id]
		} else {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{
				token: tagIn, _sleep: false, to: _id, roomname: roomname,
			}
			tagged[TaggedIn] = getLocifyProfile[_id][roomname].NickNamesViaID[_id]

		}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: CanUsePower, _sleep: false,
			to: _id, roomname: roomname}
	}

	msg := fmt.Sprintf("%s tags %s !!! a substitution from team %s %s checks in", myProfile.NickNamesViaID[myProfile.MyID], myProfile.NickNamesViaID[myProfile.Against], myTeam, myProfile.MyPatnerID)
	var t AlertMessage

	t.Alert = true
	t.Message = msg
	c, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: Message + string(c)}

	p := getLocifyProfile[ID][roomname]
	o := getLocifyProfile[p.Against][roomname]

	UpdateProfile(false, ID, roomname, _StringSentinel_,
		myTeam, _StringSentinel_, nil, true, false,
		p.IdlePlayer, _StringSentinel_, nil, _IntSentinel)

	UpdateProfile(false, p.Against, roomname,
		p.IdlePlayer, o.MyTeam, _StringSentinel_,
		nil, false, true,
		o.IdlePlayer, _StringSentinel_, nil, _IntSentinel)

	UpdateProfile(false, p.IdlePlayer, roomname, p.Against, myTeam, _StringSentinel_,
		nil, false, true,
		p.MyID, _StringSentinel_, nil, _IntSentinel)

}
