package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveCommonMap(t *testing.T) {
	_roomname := "a"

	token := map[string]map[string]ISingleSave{}
	token["ID"] = map[string]ISingleSave{
		_roomname: {
			SetChallenge: "INDIA",
		},
	}

	_src := map[string]map[string]ClashFixtures{}
	_src["ID"] = map[string]ClashFixtures{
		_roomname: {
			OppoSetChallenge: _StringSentinel_,
		},
	}
	for id, roomsDet := range token {
		for roomname, det := range roomsDet {
			if _, ok := _src[id][roomname]; ok {
				var paste = det
				var src = _src[id][roomname]

				if src.MyTeam != _StringSentinel_ {
					src.MyTeam = paste.MyTeam
				}

				if paste.PowerUp != nil {
					src.OppoPowerUp = paste.PowerUp
				}

				if paste.TossCalled != _StringSentinel_ {
					src.OppoTossCalled = paste.TossCalled
				}
				if paste.SetBet != _StringSentinel_ {
					src.OppoSetChallenge = paste.SetChallenge
				}
				if paste.SetChallenge != _StringSentinel_ {
					src.OppoSetBet = paste.SetBet
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
				_src[id][roomname] = src
			}
		}
	}
	assert.Equal(t, "INDIA", _src["ID"][_roomname].OppoSetChallenge)

}

func TestSaveSingleMap(t *testing.T) {
	ma := map[string]map[string]ISingleSave{}
	actual := map[string]map[string]ClashFixtures{}
	token := ma
	_to := actual
	for id, roomsDet := range token {
		for roomname, det := range roomsDet {
			if _, ok := _to[id][roomname]; ok {
				var paste = det
				var src = _to[id][roomname]

				if src.MyTeam != _StringSentinel_ {
					src.MyTeam = paste.MyTeam
				}

				if len(paste.PowerUp) != 0 {
					src.OppoPowerUp = paste.PowerUp
				}

				switch true {
				case paste.TossCalled != _StringSentinel_:
					src.OppoTossCalled = paste.TossCalled
				case paste.SetBet != _StringSentinel_:
					src.OppoSetChallenge = paste.SetChallenge
				case paste.SetChallenge != _StringSentinel_:
					src.OppoSetBet = paste.SetBet
				}
				_to[id][roomname] = src
			}
		}
	}
}
