package server

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	decks "nyg/deck"
	"nyg/profiles"
	"strings"

	"github.com/gofiber/contrib/websocket"
)

func ClashTags(h *Hub, conn *websocket.Conn, roomname string) {
	log.Println("in clash tags")

	shuffle := getClashTokens.storeClient[roomname]

	// doing this so that the next room mode shuffle will be from b->a rather than shuffling default
	rand.Shuffle(len(Arenas), func(i, j int) {
		Arenas[i], Arenas[j] = Arenas[j], Arenas[i]
	})

	rand.Shuffle(len(shuffle), func(i, j int) {
		shuffle[i], shuffle[j] = shuffle[j], shuffle[i]
	})

	log.Println("shuffle: ", shuffle)

	teamRed := shuffle[:2]
	teamBlue := shuffle[2:]
	chances := map[string]int{}
	chances[teamRed[0]] = 10
	chances[teamRed[1]] = 10

	chances[teamBlue[0]] = 10
	chances[teamBlue[1]] = 10

	// round := 1
	teams := map[string]map[string][]string{}
	teams[roomname] = map[string][]string{
		_TeamRedKey:  teamRed,
		_TeamBlueKey: teamBlue,
	}

	ids := []string{shuffle[0], shuffle[1], shuffle[2], shuffle[3]}

	_c := profiles.Fetch(ids)
	log.Println("profile: ", _c)
	// just for testing
	// later it will be in random shuffle
	book := []string{"cricket", "music", "moves", "basketball"}
	rand.Shuffle(len(book), func(i, j int) {
		book[i], book[j] = book[j], book[i]
	})

	li, t := SendDictionary(book[0])
	sendDic := SetupDictionaryURL + t
	h.broadcast <- BroadcastReq{Token: sendDic, RoomID: roomname}

	// ClashMatchUp(h, round, roomname, teams)
	var pack ParcelClashInfo
	pack.Chances = 10
	pack.Round = 1
	pack.GameBegin = true
	pack.Mode = Arenas[0]
	pack.RoomName = roomname
	pack.Pairing = map[string]Pairing{}
	pack.IDs = shuffle
	pack.List = li
	pack.DiscussionTime = clashDiscussionTime
	pack.GameTime = clashGameTime
	pack.FinalBoss = false
	pack.LastDance = false
	createNickName := map[string]map[string]string{}
	createNickName[roomname] = map[string]string{}

	for _id, nickname := range _c.NickNames {
		var nn = nickname[:8]
		if _, ok := createNickName[roomname]; !ok {
			createNickName[roomname] = map[string]string{}
		}
		if _, ok := createNickName[roomname][_id]; !ok {
			createNickName[roomname][_id] = nn
		}
	}
	log.Println("created nickname: ", createNickName)
	storeClashNicknames <- createNickName

	matches, create_ := ClashMatchUp(h, 1, roomname, teams, chances, createNickName)
	for _id, nickname := range _c.NickNames {
		for _, id_ := range matches[roomname][_TeamRedKey] {
			if strings.Contains(_id, id_) {
				var nn = nickname[0:8]
				var temp = pack.Pairing[nn]
				temp.TeamName = "RED"
				temp.Chances = 10
				pack.Pairing[nn] = temp

			}
		}
	}

	for _id, nickname := range _c.NickNames {
		for _, id_ := range matches[roomname][_TeamBlueKey] {
			if strings.Contains(_id, id_) {
				var nn = nickname[0:8]
				var temp = pack.Pairing[nn]
				temp.TeamName = "BLUE"
				temp.Chances = 10
				pack.Pairing[nn] = temp
			}
		}
	}

	log.Println("pairings: ", pack.Pairing)

	_pack, _ := json.Marshal(pack)
	_token := "ClashOn: " + string(_pack)
	h.broadcast <- BroadcastReq{RoomID: roomname, Token: _token}
	h.wg.Go(func() {
		// to send:
		// created deck
		var powerParcel IPowers
		_decks, _ := decks.Fetch(ids)
		c, _ := json.Marshal(powerParcel)
		_tokens := "ClashPowers: " + string(c)
		log.Println("fetched: ", _decks)
		h.broadcast <- BroadcastReq{RoomID: roomname, Token: _tokens}
		log.Println("keys: ", _decks.PowerKeys)
		for _id, powerKeys := range _decks.PowerKeys {
			for _, powerKey := range powerKeys {
				switch powerKey {
				case "nexus":
					powerParcel.Nexus = true
				case "covert":
					powerParcel.Covert = true
				case "bet":
					powerParcel.Bet = true
				case "draw":
					powerParcel.Draw = true
				case "tag":
					powerParcel.Tag = true
				case "freeze":
					powerParcel.Freeze = true
				case "rewind":
					powerParcel.Rewind = true
				}
				m, _ := json.Marshal(powerParcel)
				token := string(m)
				log.Println("sending: ", token, "to: ", _id)
				h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "ClashPowers: " + token, roomname: roomname, _sleep: false, to: _id}
				log.Println("storing matches: ", matches)
			}
		}
	})

	credits := profiles.FetchCredits(ids, true, false)

	for _id, cred := range credits.PlayerCredits {
		cred.BoardTheme = CreateBoard(Arenas[0])
		cred.GuessTheme = "clash-text"

		cred.Name = _c.NickNames[_id]
		c, err := json.Marshal(cred)
		t := string(c)
		var src = create_[_id][roomname]
		src.MyCredits = cred
		create_[_id][roomname] = src
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("credits: ", cred)
		h.gameRoomBroadcast <- reqGameRoomBroadcast{to: _id, token: "NYGcredits: " + t, _sleep: false, roomname: roomname}
	}

	for id := range create_ {

		var _book = create_[id][roomname]
		_book.Book = book[0]
		create_[id][roomname] = _book
	}

	gp := GenerateGameProfile(ids, false, false)
	for id, g := range gp {
		if _, ok := create_[id]; ok {
			var src = create_[id][roomname]
			src.MyGameProfile = g
			create_[id][roomname] = src
		}
	}

	createProfile <- create_

	storeShuffle.store <- matches

	log.Println("clash tags done")
}
