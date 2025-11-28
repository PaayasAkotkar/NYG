package server

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCal(t *testing.T) {
	type Pairing struct {
		Chances  int
		TeamName string
	}
	type ClashFixtures struct {
		MyCurrentChances     int
		EliminatedPlayersIDs []string
	}
	type ParcelClashResult struct {
		Pairing map[string]Pairing
	}
	gameResult := ParcelClashResult{}
	roomname := "a"
	myProfile := map[string]map[string]ClashFixtures{}
	myProfile["a"] = map[string]ClashFixtures{
		roomname: {
			MyCurrentChances: 100,
		},
	}
	myProfile["b"] = map[string]ClashFixtures{
		roomname: {
			MyCurrentChances: 100,
		},
	}
	profile := ClashFixtures{
		EliminatedPlayersIDs: []string{"a", "b"},
	}

	gameResult.Pairing = map[string]Pairing{}
	nicknames := map[string]map[string]string{}
	nicknames[roomname] = map[string]string{
		"a": "king",
		"b": "queen",
	}
	for _, _id := range profile.EliminatedPlayersIDs {
		if _, ok := gameResult.Pairing[nicknames[roomname][_id]]; !ok {
			src := gameResult.Pairing[nicknames[roomname][_id]]
			src.Chances = myProfile[_id][roomname].MyCurrentChances
			src.TeamName = "GREY"
			gameResult.Pairing[nicknames[roomname][_id]] = src
		}
	}
	assert.Equal(t, 100, gameResult.Pairing[nicknames[roomname]["a"]].Chances)

}
func TestFinalBoss(t *testing.T) {
	currentChance := 2
	guess := false
	deduce := uint(0)
	id := "axa" // loser id
	r := "s"
	fid := map[string]string{}
	_max := func(a, b uint) uint {
		if a > b {
			return a
		}
		return b
	}
	if !guess {
		deduce = _max(uint(currentChance), 2) - uint(currentChance)
	}
	if deduce == 0 {
		fid[r] = id
	}

	assert.Equal(t, "axa", fid[r])
}

func TestStatsSave(t *testing.T) {
	roomname := "a"
	id := "sfss"
	actual := map[string]map[string]ClashFixtures{}
	actual[id] = map[string]ClashFixtures{
		roomname: {},
	}
	actual["fsf"] = map[string]ClashFixtures{
		roomname: {},
	}
	except := map[string]map[string]ISingleStatsSave{}
	except[id] = map[string]ISingleStatsSave{
		roomname: {
			OnFire:         0,
			CurrentChances: 8,
			Guess:          "",
		},
	}
	token := except

	for id, roomsDet := range token {
		for roomname, det := range roomsDet {
			// safety

			var src = actual[id][roomname]
			var paste = det
			if paste.CurrentChances != _IntSentinel {
				src.MyCurrentChances = paste.CurrentChances
			}
			if paste.OnFire != _IntSentinel {
				src.ImOnFire += paste.OnFire
			}
			if paste.Guess != _StringSentinel_ {
				src.MyGuess = append(src.MyGuess, paste.Guess)
			}
			actual[id][roomname] = src
			log.Println("saved stats: ", getClashProfile[id][roomname])
		}
	}
	assert.Equal(t, 8, actual[id][roomname].MyCurrentChances)
	assert.Equal(t, 0, actual["fsf"][roomname].MyCurrentChances)

}
