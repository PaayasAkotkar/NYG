// Package profiles fetches profiles from the server
// all rights reserved, copyright 2025
package profiles

import (
	"app/nygpostprotoc"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Parcel struct {
	NickNames map[string]string `json:"NickNames"`
}

type Pattern struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickname"`
}

// Fetch fetches the sports data key will be the name of the book
func Fetch(ids []string) Parcel {
	// todo: we need deck[] because we will be fetching to all of them
	fmt.Println("IN PROFILE FETCH")
	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// if this is not used its impossible to tell coroutine to actually go for the signal
	// var cond = sync.NewCond(&mu) // create a new condition for_
	parcel := Parcel{}
	parcel.NickNames = map[string]string{}
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("ids: ", ids)

	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	client := nygpostprotoc.NewNYGPostProfileClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, _id := range ids {
		profile, err := client.RecieveNYGProfile(ctx, &nygpostprotoc.NYGProfileParam{Id: _id})
		if err != nil {
			panic(err)
		}
		if _, ok := parcel.NickNames[_id]; !ok {
			parcel.NickNames[_id] = profile.Nickname
		}
	}

	return parcel
}

type IProfile_ struct {
	// Record      int    `json:"record"`
	Gamesplayed int    `json:"gamesPlayed"`
	Tier        string `json:"tier"`
	Rating      int    `json:"rating"`
	Points      int    `json:"points"`
}

type Credit struct {
	Name       string    `json:"name"`
	Coin       string    `json:"coins"`
	Spur       string    `json:"spurs"`
	BoardTheme string    `json:"boardTheme"`
	GuessTheme string    `json:"guessTheme"`
	ID         string    `json:"id"`
	Nexus      int32     `json:"nexus"`
	Freeze     int32     `json:"freeze"`
	Profile    IProfile_ `json:"credits"`
	ImageURL   string    `json:"imageURL"`
}

type ParcelCredits struct {
	PlayerCredits map[string]Credit // id->...
}

func FetchCredits(ids []string, clash bool, onevone bool) ParcelCredits {
	log.Println("in fetch credits")
	parcel := ParcelCredits{}
	parcel.PlayerCredits = map[string]Credit{}
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("ids: ", ids)

	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	client := nygpostprotoc.NewNYGPostProfileClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, _id := range ids {
		profile, err := client.RecieveNYGCredits(ctx, &nygpostprotoc.NYGPlayerCreditsParam{Id: _id})
		if err != nil {
			panic(err)
		}
		var c Credit
		c.ID = _id
		c.Coin = profile.Coins
		c.ImageURL = profile.ImgURL
		c.Spur = profile.Spurs
		c.BoardTheme = profile.BoardTheme
		c.GuessTheme = profile.GuessTheme
		c.Freeze = profile.Freeze
		c.Nexus = profile.Nexus
		var src IProfile_
		if profile.ClashProfile != nil && profile.OnevoneProfile != nil && profile.TwovtwoProfile != nil {
			switch true {
			case clash:
				src.Gamesplayed = int(profile.ClashProfile.GamesPlayed)
				src.Points = int(profile.ClashProfile.Rating)
				src.Rating = int(profile.ClashProfile.Rating)
				src.Tier = profile.ClashProfile.Tier

			case onevone:
				src.Gamesplayed = int(profile.OnevoneProfile.GamesPlayed)
				src.Points = int(profile.OnevoneProfile.Rating)
				src.Rating = int(profile.OnevoneProfile.Rating)
				src.Tier = profile.OnevoneProfile.Tier
			case !onevone:
				src.Gamesplayed = int(profile.TwovtwoProfile.GamesPlayed)
				src.Points = int(profile.TwovtwoProfile.Rating)
				src.Rating = int(profile.TwovtwoProfile.Rating)
				src.Tier = profile.TwovtwoProfile.Tier
			}
		}
		c.Profile = src
		if _, ok := parcel.PlayerCredits[_id]; !ok {
			parcel.PlayerCredits[_id] = c
		}
	}
	return parcel
}
