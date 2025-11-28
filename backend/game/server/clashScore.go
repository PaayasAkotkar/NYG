package server

import "log"

const (
	DEDUCE = 2
)

type ClashMatchResult struct {
	Pairing map[string]Pairing `json:"pairing"`
}

// DeductMec returns the left chances after calc
func DeductMec(winnerID string, roomname string, Time int) int {
	WinnerStats := getClashProfile[winnerID][roomname]
	LoserStats := getClashProfile[WinnerStats.Against][roomname]
	leftChances := Deduct(int(LoserStats.MyCurrentChances), Time) - int(WinnerStats.ImOnFire)
	log.Println("decued: ", leftChances)
	return leftChances
}

func Deduct(currentChance int, underTime int) int {
	deduce := int(0)

	switch underTime {
	case 10:
		deduce = Max(5, int(currentChance)) - currentChance
	case 9:
		deduce = Max(4, int(currentChance)) - currentChance
	case 8:
		deduce = Max(3, int(currentChance)) - currentChance
	default:
		deduce = Max(2, int(currentChance)) - currentChance
	}
	return deduce
}

func Max(a, b int) int {
	if a > b {
		return int(a)
	}
	return int(b)
}
