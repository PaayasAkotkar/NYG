package server

import (
	"app/nygprotoc"
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IGameOver struct {
	RatingIncrementedBy int  `json:"ratingIncrementedBy"`
	RatingDecrementedBy int  `json:"ratingDecrementedBy"`
	NewRating           int  `json:"newRating"`
	Alert               bool `json:"alert"`
}

// ClashGameOver note: to add the current playerPoints too id->...
// you can pass the eliminated id and eliminated by id
func ClashGameOver(h *Hub, roomname, winnerID, losserID string, playerPoints map[string]int, draw bool) {
	pack := map[string]*nygprotoc.TypeGCredentials{}
	wPrf := getClashProfile[winnerID][roomname] // only using to get the idle player id and powers used
	lPrf := getClashProfile[losserID][roomname]

	ratings := PrevRatings(wPrf.IDs, true, false)
	wRating, lRating := RatingUpdate(float64(ratings[winnerID]), float64(ratings[losserID]), draw)

	lPoints := playerPoints[losserID]
	wPoints := playerPoints[winnerID]
	wpUsed := wPrf.MyPowersBin
	lpUsed := lPrf.MyPowersBin
	// we'll only update result +1 for the winner ID while result 0
	var ud nygprotoc.TypeGCredentials

	// comman
	ud.GamesPlayed = 1 // +1
	// end comman

	ud.Points = int32(wPoints)
	ud.Rating = int32(wRating.NewRating)
	ud.Results = 1 // if won than +1
	ud.Powers = wpUsed
	ud.Decremented = int32(wRating.DecrementedBy)
	ud.Incremented = int32(wRating.IncrementedBy)
	pack[winnerID] = &ud

	ud.Points = int32(lPoints)
	ud.Rating = int32(lRating.NewRating)
	ud.Results = 0
	ud.Powers = lpUsed
	ud.Decremented = int32(lRating.DecrementedBy)
	ud.Incremented = int32(lRating.IncrementedBy)
	pack[losserID] = &ud

	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	client := nygprotoc.NewNYGGameCredentialsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.RecievedNYGGameCredentials(ctx, &nygprotoc.NYGGCredentials{PlayersGameCredits: pack, Clash: true, Onevone: false})
	if err != nil {
		log.Println(err)
	}
	defer cancel()
	log.Println("grpc NYG response: ", resp)
	var w IGameOver
	w.Alert = true
	w.RatingDecrementedBy = int(wRating.DecrementedBy)
	w.RatingIncrementedBy = int(wRating.IncrementedBy)
	w.NewRating = int(wRating.NewRating)
	t, err := json.Marshal(w)
	if err != nil {
		log.Println(err)
		return
	}

	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "NYGGameOver: " + string(t), roomname: roomname, to: winnerID, _sleep: false}
	w.RatingDecrementedBy = int(lRating.DecrementedBy)
	w.RatingIncrementedBy = int(lRating.IncrementedBy)
	w.NewRating = int(lRating.NewRating)

	t, err = json.Marshal(w)
	if err != nil {
		log.Println(err)
		return
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "NYGGameOver: " + string(t), roomname: roomname, to: losserID, _sleep: false}

}

// EliminatedClashGameOver only updates for the losserID
func EliminatedClashGameOver(h *Hub, roomname, winnerID, losserID string, playerPoints map[string]int, draw bool) {
	pack := map[string]*nygprotoc.TypeGCredentials{}

	ratings := PrevRatings([]string{losserID}, true, false)
	_, lRating := RatingUpdate(float64(ratings[winnerID]), float64(ratings[losserID]), draw)

	lPoints := playerPoints[losserID]
	lpUsed := getClashProfile[losserID][roomname].MyPowersBin

	// we'll only update result +1 for the winner ID while result 0
	var ud nygprotoc.TypeGCredentials

	ud.Points = int32(lPoints)
	ud.Rating = int32(lRating.NewRating)
	ud.Results = 0
	ud.Powers = lpUsed
	ud.Decremented = int32(lRating.DecrementedBy)
	ud.Incremented = int32(lRating.IncrementedBy)
	pack[losserID] = &ud

	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	client := nygprotoc.NewNYGGameCredentialsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.RecievedNYGGameCredentials(ctx, &nygprotoc.NYGGCredentials{PlayersGameCredits: pack, Clash: true, Onevone: false})
	if err != nil {
		log.Println(err)
	}
	defer cancel()
	log.Println("grpc NYG response: ", resp)

	var w IGameOver
	w.Alert = true

	w.RatingDecrementedBy = int(lRating.DecrementedBy)
	w.RatingIncrementedBy = int(lRating.IncrementedBy)
	t, err := json.Marshal(w)
	if err != nil {
		log.Println(err)
		return
	}
	h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "NYGGameOver: " + string(t), roomname: roomname, to: losserID, _sleep: false}

}
