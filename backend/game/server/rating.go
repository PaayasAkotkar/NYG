package server

import (
	"log"
	"math"
)

const (
	// point gain
	Kbeg         = 400 // players>1000
	kinter       = 200 // players<1000 && players <2000
	kmaster      = 20  // players>2000
	ksupermaster = 10  // players>2500

	// points on performance
	WINNER = 1
	LOSSER = 0
	DRAW   = 0.5
)

type IRating struct {
	NewRating     float64
	DecrementedBy float64
	IncrementedBy float64
}

// RatingUpdate returns the result of the game
func RatingUpdate(WinnerPreviousRating float64, LosserPreviousRating float64, isDraw bool) (IRating, IRating) {
	RwinnerR := IRating{}
	RlosserR := IRating{}

	WRating, LRating, K1, K2 := WinnerPreviousRating, LosserPreviousRating, 32.0, 1.0
	K3 := Kbeg
	// note: the porbability is generated upon the  highest rated player to be a winner
	switch true {
	case WRating < 1000:
		K1 = 96.16
		log.Println("<900")

	case WRating < 1400:
		K1 = 105.0
		K3 = kinter
		log.Println("<1400")

	case WRating < 2000:
		K1 = 108.14
		K3 = kinter
		log.Println("<2000")

	case WRating < 2200:
		K1 = 112.6
		K3 = kmaster
		log.Println("<2200")

	case WRating < 2500:
		K1 = 144.0
		K3 = kmaster
		log.Println("<2500")
	default:
		K1 = 36.0
		K3 = ksupermaster
	}

	// calculation of probablity of excepted win result
	strength := -(WRating - LRating) / float64(K3)
	pow := 1 + (math.Pow(10, strength))
	calc := 1 / pow
	ExcpetedScore := calc
	// end of calculation

	// upset // if the highest player looses
	ExcpetedScore2 := 1 - ExcpetedScore
	K2 = K1 - 5

	log.Println("k: ", K1)
	incrementedBy := math.Round(K1 * (WINNER - ExcpetedScore))
	decrementedBy := math.Round(K2 * (LOSSER - ExcpetedScore2))
	// constant: 32 can be change for different rating
	// case: for winner rating
	NewWRating := math.Round(WRating + K1*(WINNER-ExcpetedScore))
	NewLRating := math.Round(LRating + K2*(LOSSER-ExcpetedScore2))
	if isDraw {
		NewWRating = math.Round(WRating + K1*(DRAW-ExcpetedScore))
		NewLRating = math.Round(LRating + K2*(DRAW-ExcpetedScore2))
		incrementedBy = math.Round(K1 * (DRAW - ExcpetedScore))
		decrementedBy = math.Round(K2 * (DRAW - ExcpetedScore2))
	}
	log.Println("previous ratings: ", WRating, LRating)
	log.Println("new ratings: ", NewWRating, "🔥", NewLRating, "🔥")
	if NewLRating < 0 {
		NewLRating = 0
	}
	RwinnerR.IncrementedBy = incrementedBy
	RwinnerR.DecrementedBy = 0
	RwinnerR.NewRating = NewWRating

	RlosserR.DecrementedBy = decrementedBy
	RlosserR.NewRating = NewLRating
	RlosserR.IncrementedBy = 0
	
	return RwinnerR, RlosserR
}

// to update clash player ratings
// use the cycle of elimination
// eliminated players will get the updated ratings
// while the in game players will not get the rating
// in future it will also depend upon the kickout
// meaning if x eliminated y than x's rating will be saved
// and will be sum up at the end

// TwoPlayerAverage returns the average rating of the two players for winner
// concept: twovtwo-> a->[x,y]=> [a,z] where z is the avegrage of x+y/2
func TwoPlayerAverage(LosserPlayer1Rating float64, LosserPlayer2Rating float64) float64 {
	return (LosserPlayer1Rating + LosserPlayer2Rating) / 2
}

// FourPlayerAverage returns the average rating of the two players for winner
// concept: twovtwo-> a->[x,y,v,w]=> [a,z] where z is the avegrage of x+y+v+w/4
// note: this is only for 2v2
func FourPlayerAverage(LosserPlayer1Rating float64, LosserPlayer2Rating float64, LosserPlayer3Rating float64, LosserPlayer4Rating float64) float64 {
	return (LosserPlayer1Rating + LosserPlayer2Rating + LosserPlayer3Rating + LosserPlayer4Rating) / 4
}

// ThreePlayerAverage returns the average rating of the two players for winner
// concept: twovtwo-> a->[x,y,v]=> [a,z] where z is the avegrage of x+y+v/2
// note: this is only for 2v2
func ThreePlayerAverage(LosserPlayer1Rating float64, LosserPlayer2Rating float64, LosserPlayer3Rating float64) float64 {
	return (LosserPlayer1Rating + LosserPlayer2Rating + LosserPlayer3Rating) / 3
}
