package validate

import (
	"fmt"
	"nyg/dataset"
	"slices"
	"strings"
)

// ThreeFold prases the word into fold
// input: [king p, queen p]
// output: [[king, p], [queen,p]]
func ThreeFold(_from []string) [][]string {
	parse := [][]string{}
	for _, r := range _from {
		_split := strings.Split(r, " ")
		parse = append(parse, _split)
	}
	return parse
}

// Fold union of parent
func Fold(_parent [][]string, _search []string) []string {
	var match []string
	for _, inner := range _parent {
		search := strings.Join(inner, " ")

		for _, token := range _search {
			// search via shortcut, prefix and if contains
			if token == search || strings.HasPrefix(search, token) || strings.Contains(search, token) {
				match = append(match, inner...) // note: to append all the associated id with it
			}
		}
	}
	if match == nil {
		match = append(match, _search...)
	}
	return match
}

// Match simple search
func Match(list []string, search string) bool {
	found := false
	for _, token := range list {
		// token==search just for safety
		if strings.Contains(token, search) || token == search {
			found = true
		}
	}
	return found
}

// SportsValidate searches in two way for old and new and returns the list of all the associated keys
// conditon: if new and !old=> true
func SportsValidate(book string, dictionary string, list string, newSearch string, Oldsearch []string) (bool, []string) {
	token := Fetch(book, dictionary, list).Pack
	exam := []string{}
	valid := false
	exam = append(exam, token[dictionary][list]...)
	match := dataset.ParseWords(exam)
	valid = Match(exam, newSearch)

	// keep searching in the database if it contains the name
	for _, name := range match {

		if name == newSearch {
			valid = true
		}
	}

	fmt.Println("searched for 🗺️: ", newSearch, "found ❓ ", valid)

	temp := []string{newSearch}
	X := dataset.ParseWords(Oldsearch) // go through the first and last name too

	// keep searching in the previous-database if it contains the name
	for _, se := range temp {
		if slices.Contains(X, se) {
			valid = false
		}
	}

	Oldsearch = append(Oldsearch, newSearch)
	threeFolder := ThreeFold(exam)
	fold := Fold(threeFolder, Oldsearch)

	fmt.Println("examine data: ", exam)
	fmt.Println("fold: ", fold)
	fmt.Println("examine in oldsearch 🗺️: ", Oldsearch, "found ❓ ", valid)

	return valid, fold
}
