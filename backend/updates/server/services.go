package server

import (
	"app/nygprotoc"
	"context"
	"log"
)

type Services struct {
	nygprotoc.UnimplementedNYGGameCredentialsServer
	nygprotoc.UnimplementedNYGRatingServer
}

func (s *Services) RecievedNYGGameCredentials(_ context.Context, ngc *nygprotoc.NYGGCredentials) (*nygprotoc.NYGGCredentialsReply, error) {
	log.Println("in nygame")
	if err := Update(ngc.PlayersGameCredits, ngc.Clash, ngc.Onevone); err != nil {
		log.Println(err)
		return nil, err
	}
	return &nygprotoc.NYGGCredentialsReply{Reached: true, Msg: "The database will update"}, nil
}

func (s *Services) GetNYGRating(_ context.Context, id *nygprotoc.NYGpID) (*nygprotoc.NYGpRating, error) {
	log.Println("in rating")

	rating := -1
	if id.Clash {
		c, err := GetClashProfile(id.Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		rating = c.Rating
	} else if id.Onevone {
		c, err := GetOnevOneProfile(id.Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		rating = c.Rating
	} else {
		c, err := GetTwovTwoProfile(id.Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		rating = c.Rating
	}
	return &nygprotoc.NYGpRating{Rating: int32(rating)}, nil
}
