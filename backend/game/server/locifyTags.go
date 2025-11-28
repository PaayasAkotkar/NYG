package server

import (
	"encoding/json"
	"log"
	"math/rand"
	"nyg/profiles"
)

func LocifyTags(h *Hub, l Lobby, roomname string, ids []string) {

	log.Println("in locify tags")
	log.Println("ids: ", ids)

	isFriend := false

	for room, det := range getLocifyRoomSettings {
		if l.Code == det.Code {
			roomname = room
			isFriend = det.Friend
			break
		}
	}

	cap := getLocifyRoomSettings[roomname].Capacity

	log.Println("calculated cap: ", cap)

	switch cap {
	case 2:
		LocifyOneVOne(h, ids, roomname)

	case 4:
		switch isFriend {
		case true:
			fID := l.ID

			if l.Code == "" {
				fID = getLocifyRoomSettings[roomname].FriendID
			}

			LocfiyTwoVTwoFriend(h, fID, ids, roomname)

		case false:
			LocifyTwoVTwo(h, ids, roomname)

		default:
			log.Println("cant able to tag")

		}
	default:
		log.Println("cant able to tag")

	}
	log.Println("tag done")

}

func LocifyOneVOne(h *Hub, ids []string, roomname string) {
	log.Println("locify one v one")
	draft := ids
	rand.Shuffle(len(draft), func(i int, j int) {
		draft[i], draft[j] = draft[j], draft[i]
	})

	teams := map[string]map[string][]string{}
	teams[roomname] = map[string][]string{
		_TeamRedKey:  {draft[0]},
		_TeamBlueKey: {draft[1]},
	}

	adminID := getLocifyRoomSettings[roomname].RoomOwnerID
	book := getLocifyRoomSettings[roomname].Book
	roomCode := getLocifyRoomSettings[roomname].Code
	settings := getLocifyRoomSettings[roomname]

	_teams, matches := LocifyMatchUpOneVOne(h, roomname, teams)

	_profiles := LocifyProfile(h, 1, adminID,
		book, roomname, matches,
		roomCode, settings)

	_c := profiles.Fetch(ids)

	_createNickname := make(map[string]string)

	for _, _id := range ids {
		_createNickname[_id] = _c.NickNames[_id]
	}

	for _, _id := range ids {
		var src = _profiles[_id][roomname]
		src.NickNamesViaID = _createNickname
		_profiles[_id][roomname] = src
	}
	go func() {
		gp := GenerateGameProfile(ids, false, false)
		for id, g := range gp {
			if _, ok := _profiles[id]; ok {
				var src = _profiles[id][roomname]
				src.MyGameProfile = g
				_profiles[id][roomname] = src
			}
		}
		log.Println("gp: ", gp)
	}()

	for _, roomsDet := range _profiles {
		for _, det := range roomsDet {
			if det.ILock {
				log.Println("lock id: ", det.MyID, "team: ", det.MyTeam)
			}
		}
	}

	nn := []string{}
	for _, _nn := range _c.NickNames {
		nn = append(nn, _nn)
	}

	log.Println("ids: ", ids)
	credits := profiles.FetchCredits(ids, false, true)
	for _id, cred := range credits.PlayerCredits {
		cred.Name = _c.NickNames[_id]
		c, err := json.Marshal(cred)
		t := string(c)
		var src = _profiles[_id][roomname]
		src.MyCredits = cred
		_profiles[_id][roomname] = src
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("credits: ", cred)
		h.gameRoomBroadcast <- reqGameRoomBroadcast{to: _id, token: "NYGcredits: " + t, _sleep: false, roomname: roomname}
	}

	FillUpLocifyTags(h, roomname, teams, _c.NickNames, nn, _profiles)

	h.broadcast <- BroadcastReq{RoomID: roomname, Token: gameBegin}
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: Lock}

	go func() {
		createLocifyProfile <- _profiles
		storeShuffle.store <- _teams
	}()
}

// LocfiyTwoVTwoFriend  current id of the request token
func LocfiyTwoVTwoFriend(h *Hub, id string, ids []string, roomname string) {
	friendID := id
	if getLocifyRoomSettings[roomname].FriendID != _StringSentinel_ {
		friendID = getLocifyRoomSettings[roomname].FriendID
	}
	teamRed := []string{friendID, getLocifyRoomSettings[roomname].RoomOwnerID}
	teamBlue := []string{}
	for _, _id := range ids {
		if _id != friendID && _id != getLocifyRoomSettings[roomname].RoomOwnerID {
			teamBlue = append(teamBlue, _id)
		}
	}
	teams := map[string]map[string][]string{}
	teams[roomname] = map[string][]string{
		_TeamRedKey:  teamRed,
		_TeamBlueKey: teamBlue,
	}
	_teams, matches := LocifyMatchUp(h, 1, roomname, teams)

	adminID := getLocifyRoomSettings[roomname].RoomOwnerID
	book := getLocifyRoomSettings[roomname].Book
	roomCode := getLocifyRoomSettings[roomname].Code
	settings := getLocifyRoomSettings[roomname]
	_profiles := LocifyProfile(h, 1, adminID,
		book, roomname, matches, roomCode, settings)

	_c := profiles.Fetch(ids)

	_createNickname := make(map[string]string)

	for _, _id := range ids {
		_createNickname[_id] = _c.NickNames[_id]
	}

	for _, _id := range ids {
		var src = _profiles[_id][roomname]
		src.NickNamesViaID = _createNickname
		_profiles[_id][roomname] = src
	}

	for _, roomsDet := range _profiles {
		for _, det := range roomsDet {
			if det.ILock {
				log.Println("lock id: ", det.MyID, "team: ", det.MyTeam)
			}
		}
	}
	go func() {
		gp := GenerateGameProfile(ids, false, false)
		for id, g := range gp {
			if _, ok := _profiles[id]; ok {
				var src = _profiles[id][roomname]
				src.MyGameProfile = g
				_profiles[id][roomname] = src
			}
		}
		log.Println("gp: ", gp)
	}()

	nn := []string{}
	for _, _nn := range _c.NickNames {
		nn = append(nn, _nn)
	}
	FillUpLocifyTags(h, roomname, teams, _c.NickNames, nn, _profiles)

	credits := profiles.FetchCredits(ids, false, false)
	for _id, cred := range credits.PlayerCredits {
		c, err := json.Marshal(cred)

		t := string(c)
		if err != nil {
			log.Println(err)
			return
		}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{to: _id, token: "NYGcredits: " + t, _sleep: false, roomname: roomname}
	}

	for _, _id := range ids {
		var src = _profiles[_id][roomname]
		if src.ILock {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{roomname: roomname,
				token: Lock, to: _id, _sleep: false}
		}
	}
	go func() {
		storeShuffle.store <- _teams
		createLocifyProfile <- _profiles
	}()
}

func LocifyTwoVTwo(h *Hub, ids []string, roomname string) {
	draft := ids
	log.Println("twov two ids: ", ids)
	rand.Shuffle(len(draft), func(i int, j int) {
		draft[i], draft[j] = draft[j], draft[i]
	})

	teamRed := draft[:2]
	teamBlue := draft[2:]
	teams := map[string]map[string][]string{}
	teams[roomname] = map[string][]string{
		_TeamRedKey:  teamRed,
		_TeamBlueKey: teamBlue,
	}

	store, matches := LocifyMatchUp(h, 1, roomname, teams)

	adminID := getLocifyRoomSettings[roomname].RoomOwnerID

	book := getLocifyRoomSettings[roomname].Book
	roomCode := getLocifyRoomSettings[roomname].Code
	settings := getLocifyRoomSettings[roomname]

	_profiles := LocifyProfile(h, 1, adminID, book,
		roomname, matches, roomCode, settings)
	_c := profiles.Fetch(ids)

	_createNickname := make(map[string]string)
	credits := profiles.FetchCredits(ids, false, false)
	for _id, cred := range credits.PlayerCredits {
		c, err := json.Marshal(cred)
		t := string(c)
		if err != nil {
			log.Println(err)
			return
		}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{to: _id, token: "NYGcredits: " + t, _sleep: false, roomname: roomname}
	}

	for _, _id := range ids {
		_createNickname[_id] = _c.NickNames[_id]
	}
	go func() {
		gp := GenerateGameProfile(ids, false, false)
		for id, g := range gp {
			if _, ok := _profiles[id]; ok {
				var src = _profiles[id][roomname]
				src.MyGameProfile = g
				_profiles[id][roomname] = src
			}
		}
		log.Println("gp: ", gp)
	}()

	for _, _id := range ids {
		var src = _profiles[_id][roomname]
		src.NickNamesViaID = _createNickname
		_profiles[_id][roomname] = src
	}
	nn := []string{}
	for _, _nn := range _c.NickNames {
		nn = append(nn, _nn)
	}

	FillUpLocifyTags(h, roomname, teams, _c.NickNames, nn, _profiles)

	go func() {
		storeShuffle.store <- store
	}()

	go func() {
		createLocifyProfile <- _profiles
	}()

	h.broadcast <- BroadcastReq{RoomID: roomname, Token: gameBegin}

}
