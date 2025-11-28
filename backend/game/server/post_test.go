package server

import (
	"math/rand"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	teams := []string{
		"dummy", "dummy", "england", "australia", "netherland",
		"srilanka", "india",
	}

	rand.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})
	ip, ib := 0, 0
	hasIndia, doIt := false, false
	ii := 0
	for ix, i := range teams {
		if strings.ToLower(i) == "india" {
			hasIndia = true
			ii = ix
			break
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
	}

	for i, t := range teams {
		switch t {
		case "pakistan":
			ip = i
		case "bangladesh":
			ib = i
		}
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
		assert.Equal(t, "india", teams[0])
		assert.Equal(t, "bangladesh", teams[len(teams)-2])
		assert.Equal(t, "pakistan", teams[len(teams)-1], "pakistan")

	} else {
		assert.Equal(t, hasIndia, true)
		assert.Equal(t, "india", teams[0])
	}

}
