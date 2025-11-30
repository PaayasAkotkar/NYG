// Package decks implements the fetching of the requested deck for stored id
// all rights reserved, copyright 2025
package decks

import (
	"app/nygpostprotoc"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Pattern struct {
	PowerKeys []string `json:"powerKeys"`
	ID        string   `json:"id"`
}
type Parcel struct {
	// id->keys
	PowerKeys map[string][]string `json:"powerKeys"`
}

// Fetch fetches the sports data key will be the name of the book
func Fetch(ids []string) (Parcel, error) {
	fmt.Println("IN DECK FETCH")
	parcel := Parcel{}
	parcel.PowerKeys = make(map[string][]string)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	client := nygpostprotoc.NewNYGPostProfileClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, _id := range ids {
		profile, err := client.RecieveNYGDeck(ctx, &nygpostprotoc.NYGDeckParam{Id: _id})
		if err != nil {
			panic(err)
		}
		if _, ok := parcel.PowerKeys[_id]; !ok {
			parcel.PowerKeys[_id] = append(parcel.PowerKeys[_id], profile.First, profile.Second)
		}
	}
	
	return parcel, nil
}
