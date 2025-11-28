package server

import (
	"log"
	"math/rand/v2"
	"nyg/list"
	"nyg/validate"
	"strings"
)

type BetPattern struct {
	FirstCup  string `json:"firstCup"`
	SecondCup string `json:"secondCup"`
	ThirdCup  string `json:"thirdCup"`
}

// BetCups returns team name and bet pattern
func BetCups(book, dictionary string) (string, BetPattern) {
	log.Println("in bet cups")
	dict := dictionary

	dicts := list.FetchSports(book, dict).Pack
	rand.Shuffle(len(dicts), func(i, j int) {
		dicts[i], dicts[j] = dicts[j], dicts[i]
	})

	teamname := dicts[0]

	if strings.Contains("pakistan", strings.ToLower(teamname)) {
		teamname = dicts[1]
	}

	_conv := validate.Fetch(book, dict, teamname).Pack

	list := _conv[dict][teamname]

	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	firstCup := list[0]
	secondCup := list[1]
	thirdCup := list[2]
	de := BetPattern{FirstCup: firstCup, SecondCup: secondCup, ThirdCup: thirdCup}
	return teamname, de
}
