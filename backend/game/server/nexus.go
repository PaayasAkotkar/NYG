package server

import (
	"log"
	"math/rand/v2"
	"strings"
)

// Nexus returns first and last name respectively
// both: returns the replacement of the both words
// max: returns the replacement of the one word while keep next word as it is
func Nexus(str string, replace string, both bool, max bool) (string, string) {
	log.Println("in nexus")
	var manipulate = func(s string) string {
		// formula: replaces the char at any k limit of
		//          random rolls
		_limit, rolls := []int{}, []int{}
		for r := 1; r < len(s); r++ {
			rolls = append(rolls, r)
			_limit = append(_limit, r)
		}
		rand.Shuffle(len(_limit), func(i int, j int) {
			_limit[i], _limit[j] = _limit[j], _limit[i]
		})
		rand.Shuffle(len(rolls), func(i int, j int) {
			rolls[i], rolls[j] = rolls[j], rolls[i]
		})

		for roll := range _limit[0] {
			re := string(s[rolls[roll]]) // char to string conv: in-order to replace recursively
			// r-1 is important
			// reason: it helps in replacing the value at rolled
			s = strings.Replace(s, re, replace, roll-1) // note: if you replace with " _ " it would cause index error
		}
		// format into: king queen to k i n g
		s = strings.Join(strings.Split(s, ""), " ")

		return s
	}

	f, l := LastNameFirstName(str)
	_f := manipulate(f)

	switch true {
	case both && !max:
		l = "|" + manipulate(l) // for clarity
		return _f, l
	case max:
		x := []string{f, l}
		rand.Shuffle(len(x), func(i, j int) {
			x[i], x[j] = x[j], x[i]
		})

		if x[0] == f {
			a := f
			_f = "|" + manipulate(l)
			return a, _f
		} else {
			a := manipulate(f)
			_f = "|" + l
			return a, _f
		}
	default:
		return _f, ""

	}

}
func LastNameFirstName(s string) (string, string) {
	// 	re := regexp.MustCompile(`\s+`)
	// 	i := re.FindIndex([]byte(str))
	firstname := ""
	lastname := ""
	// 	for _, inx := range i {
	// 		firstname = str[:inx]
	// 		lastname = str[inx:]
	// 	}
	// 	return firstname, lastname
	// }
	s = strings.TrimSpace(s)
	if s == "" {
		return "", ""
	}
	parts := strings.SplitN(s, " ", 2)
	firstname = parts[0]
	if len(parts) == 1 {
		lastname = ""
	} else {
		lastname = parts[1]
	}
	return firstname, lastname
}
