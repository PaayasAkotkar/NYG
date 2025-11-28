package server

import (
	"encoding/json"
	"log"
	"nyg/dictionary"
	"nyg/list"
	"nyg/profiles"
	"sort"
	"strings"
)

func GenerateGameProfile(ids []string, clash, onevone bool) map[string]GameProfile {
	log.Println("in generate game profile")
	gprofiles := map[string]GameProfile{}
	profiles := profiles.FetchCredits(ids, clash, onevone)

	for id, p := range profiles.PlayerCredits {
		if _, ok := gprofiles[id]; !ok {
			var src = gprofiles[id]
			src.Points = p.Profile.Points
			src.Rating = p.Profile.Rating
			src.Gamesplayed = p.Profile.Gamesplayed
			src.NexusLevel = int(p.Nexus)
			src.FreezeLevel = int(p.Freeze)
			src.Pic = p.ImageURL
			gprofiles[id] = src
		}
	}
	log.Println("game profiles: ", gprofiles)

	return gprofiles

}

func SendDictionary(book string) ([]string, string) {
	log.Println("in sending dict")

	conv := dictionary.Fetch(book)
	tok, _ := json.Marshal(conv)
	token := string(tok)
	return conv.Pack, token
}

func GetList(book, dictionary string) []string {
	conv := list.FetchSports(book, dictionary)
	return conv.Pack
}
func SendList(book, dictionary string) string {
	log.Println("in sending list")

	conv := list.FetchSports(book, dictionary)
	teams := conv.Pack

	ip, ib := 0, 0
	doIt := false
	hasIndia := false
	ii := 0
	for ix, i := range teams {
		if strings.ToLower(i) == "india" {
			hasIndia = true
			ii = ix
		}
	}

	for i, t := range teams {
		if strings.Contains(strings.ToLower(t), "pakistan") || strings.Contains(strings.ToLower(t), "bangladesh") {
			doIt = true
			switch strings.ToLower(t) {
			case "pakistan":
				ip = i
			case "bangladesh":
				ib = i
			}
		}
	}

	if hasIndia {
		sort.StringSlice(teams).Swap(ii, 0)
		conv.Pack = teams
	}

	if doIt {

		var remove = func(s []string, i int) []string {
			return append(s[:i], s[i+1:]...)
		}

		if ip > ib {
			if ip != -1 {
				teams = remove(teams, ip)
			}
			if ib != -1 {
				teams = remove(teams, ib)
			}
		} else {
			if ib != -1 {
				teams = remove(teams, ib)
			}
			if ip != -1 {
				teams = remove(teams, ip)
			}
		}

		teams = append(teams, "bangladesh")
		teams = append(teams, "pakistan")

		conv.Pack = teams
	}

	tok, _ := json.Marshal(conv)

	token := string(tok)

	return token
}
