package server

import (
	"encoding/json"
	"log"
	"math/rand/v2"
)

func LocifyResetPowerUp(id, roomname, key string) {
	log.Println("in reset")
	log.Println("reset key: ", key)
	log.Println("see: ", getLocifyProfile[id][roomname].OppoPowerUp)
	y := map[string]map[string]IResetPower{}
	var r IResetPower
	r.DoIt = true
	r.Key = key
	y[id] = map[string]IResetPower{
		roomname: r,
	}
	locifyresetPowerUp <- y

}

func FillUpLocifyTags(h *Hub, roomname string,
	teams map[string]map[string][]string, nicknamesViaID map[string]string, nicknames []string, profiles map[string]map[string]LocifyFixtures) {
	draftPowers := []string{}

	for power, include := range getLocifyRoomSettings[roomname].Powers {
		if include {
			draftPowers = append(draftPowers, power)
		}
	}

	rand.Shuffle(len(draftPowers), func(i, j int) {
		draftPowers[i], draftPowers[j] = draftPowers[j], draftPowers[i]
	})

	rp := draftPowers[:2]
	bp := draftPowers[2:]

	redTeamPower := map[string]bool{
		rp[0]: true,
		rp[1]: true,
	}

	blueTeamPower := map[string]bool{
		bp[0]: true,
		bp[1]: true,
	}

	var _gi LocfiyGameInfo
	_gi.BlueScore = 0
	_gi.RedScore = 0
	_gi.Powers = redTeamPower
	_gi.Round = 1
	_gi.Set = 1
	_gi.Session = true
	_gi.Mode = "locify"
	_gi.Roomname = roomname
	_gi.TeamName = "RED"
	_gi.PlayersStats = make(map[string]map[int]string)
	_gi.Nicknames = nicknames
	var _x = _gi.PlayersStats

	for _, _id := range teams[roomname][_TeamRedKey] {
		_x[nicknamesViaID[_id]] = map[int]string{
			1: "0",
		}
		if _, ok := profiles[_id][roomname]; ok {
			var src = profiles[_id][roomname]
			src.PowerDeck = make(map[string]bool)
			src.PowerDeck[rp[0]] = true
			src.PowerDeck[rp[1]] = true
			profiles[_id][roomname] = src
		}
	}

	for _, _id := range teams[roomname][_TeamBlueKey] {
		_x[nicknamesViaID[_id]] = map[int]string{
			1: "0",
		}
		if _, ok := profiles[_id][roomname]; ok {
			var src = profiles[_id][roomname]
			src.PowerDeck = make(map[string]bool)
			src.PowerDeck[bp[0]] = true
			src.PowerDeck[bp[1]] = true
			profiles[_id][roomname] = src
		}
	}

	c, _ := json.Marshal(_gi)

	token := "LocifyGame: " + string(c)

	for _, _id := range teams[roomname][_TeamRedKey] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token:    token,
			roomname: roomname, to: _id, _sleep: false,
		}
	}

	_gi.Powers = nil
	_gi.Powers = blueTeamPower
	_gi.TeamName = "BLUE"
	c, _ = json.Marshal(_gi)
	token = "LocifyGame: " + string(c)

	for _, _id := range teams[roomname][_TeamBlueKey] {
		h.gameRoomBroadcast <- reqGameRoomBroadcast{
			token:    token,
			roomname: roomname, to: _id, _sleep: false,
		}
	}

	BoardcastSession(h, roomname, false, true, false, false, false)

	h.wg.Go(func() {
		book := getLocifyRoomSettings[roomname].Book
		_, li := SendDictionary(book)
		sendDic := SetupDictionaryURL + li
		h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}
	})

}
