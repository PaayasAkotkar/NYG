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

func PrevRatings(ids []string, isClash bool, isOneVOne bool) map[string]int32 {
	_ratings := map[string]int32{}
	log.Println("ids: ", ids)

	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer conn.Close()
	client := nygprotoc.NewNYGRatingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, id := range ids {
		_rating, err := client.GetNYGRating(ctx, &nygprotoc.NYGpID{Id: id, Clash: isClash, Onevone: isOneVOne})
		if err != nil {
			log.Println(err)
			return nil
		}
		if _, ok := _ratings[id]; !ok {
			_ratings[id] = _rating.Rating
		}
	}
	return _ratings
}

// LocifyGameOver note: to add the current playerPoints too id->...
func LocifyGameOver(h *Hub, roomname, winnerID, losserID string, playerPoints map[string]int, r LocifyRoomSettings, draw bool) {
	pack := map[string]*nygprotoc.TypeGCredentials{}
	wPrf := getLocifyProfile[winnerID][roomname] // only using to get the idle player id and powers used
	lPrf := getLocifyProfile[losserID][roomname]
	onevone := getLocifyRoomSettings[roomname].Capacity == 2

	ratings := PrevRatings(wPrf.IDs, false, onevone)
	wRating, lRating := RatingUpdate(float64(ratings[winnerID]), float64(ratings[losserID]), draw)

	lPoints := playerPoints[losserID]
	wPoints := playerPoints[winnerID]
	wpUsed := wPrf.MyPowersBin
	lpUsed := lPrf.MyPowersBin
	cat := r.Category
	b := r.Book
	onevone = r.Capacity == 2
	// we'll only update result +1 for the winner ID while result 0
	var ud nygprotoc.TypeGCredentials

	// comman
	ud.Book = b
	ud.Category = cat
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

	if r.Capacity == 4 {

		iwPrf := getLocifyProfile[wPrf.IdlePlayer][roomname] // only using to get the idle player id and powers used
		ilPrf := getLocifyProfile[lPrf.IdlePlayer][roomname]
		wIrating, lIrating := RatingUpdate(float64(ratings[iwPrf.MyID]), float64(ratings[ilPrf.MyID]), false)

		ilPoints := playerPoints[losserID]
		iwPoints := playerPoints[winnerID]
		iwpUsed := iwPrf.MyPowersBin
		ilpUsed := ilPrf.MyPowersBin

		ud.Points = int32(iwPoints)
		ud.Rating = int32(wIrating.NewRating)
		ud.Results = 1 // if won than +1
		ud.Powers = iwpUsed
		ud.Decremented = int32(wIrating.DecrementedBy)
		ud.Incremented = int32(wIrating.IncrementedBy)
		pack[iwPrf.MyID] = &ud

		ud.Points = int32(ilPoints)
		ud.Rating = int32(lIrating.NewRating)
		ud.Results = int32(0)
		ud.Powers = ilpUsed
		ud.Decremented = int32(lIrating.DecrementedBy)
		ud.Incremented = int32(lIrating.IncrementedBy)
		pack[ilPrf.MyID] = &ud
		var w IGameOver
		w.Alert = true

		w.RatingDecrementedBy = int(wIrating.DecrementedBy)
		w.RatingIncrementedBy = int(wIrating.IncrementedBy)
		w.NewRating = int(wIrating.NewRating)
		t, err := json.Marshal(w)
		if err != nil {
			log.Println(err)
			return
		}

		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "NYGGameOver: " + string(t), roomname: roomname, to: winnerID, _sleep: false}
		w.RatingDecrementedBy = int(lIrating.DecrementedBy)
		w.RatingIncrementedBy = int(lIrating.IncrementedBy)
		w.NewRating = int(lIrating.NewRating)

		t, err = json.Marshal(w)
		if err != nil {
			log.Println(err)
			return
		}
		h.gameRoomBroadcast <- reqGameRoomBroadcast{token: "NYGGameOver: " + string(t), roomname: roomname, to: winnerID, _sleep: false}

	}

	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	client := nygprotoc.NewNYGGameCredentialsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.RecievedNYGGameCredentials(ctx, &nygprotoc.NYGGCredentials{PlayersGameCredits: pack, Clash: false, Onevone: onevone})
	if err != nil {
		log.Println(err)
	}
	defer cancel()

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

	log.Println("grpc NYG response: ", resp)
}
