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
	fmt.Println("ids: ", ids)

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
	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// // if this is not used its impossible to tell coroutine to actually go for the signal
	// var cond = sync.NewCond(&mu) // create a new condition for_
	// parcel := Parcel{}
	// parcel.PowerKeys = map[string][]string{}
	// fmt.Println("ids: ", ids)
	// var stream = make(chan map[string][]string, len(ids))

	// for _, id := range ids {
	// 	wg.Add(1)
	// 	go func(_id string) {
	// 		defer wg.Done()
	// 		url := URLs[0] + _id

	// 		Conn := resty.NewEventSource().SetURL(url)

	// 		Conn.OnOpen(func(url string) {
	// 			fmt.Println("🔥 connected to url: ", url)
	// 		})

	// 		Conn.OnMessage(func(a any) {
	// 			fmt.Println("🎧 on message triggered")
	// 			event := a.(*resty.Event)
	// 			token := []byte(event.Data)
	// 			var _get Pattern

	// 			if err := json.Unmarshal(token, &_get); err != nil {
	// 				fmt.Println("❌ JSON unmarshal error:", err)
	// 				return
	// 			}

	// 			create := map[string][]string{}
	// 			create[id] = _get.PowerKeys
	// 			stream <- create
	// 		}, nil)

	// 		if err := Conn.Get(); err != nil {
	// 			fmt.Println("errr")
	// 		}
	// 	}(id)

	// }

	// go func() {
	// 	wg.Wait()
	// 	close(stream)
	// }()
	// mu.Lock()
	// for len(parcel.PowerKeys) < len(ids) {
	// 	select {
	// 	case token, ok := <-stream:
	// 		if !ok {
	// 			return parcel, fmt.Errorf("unable to fetch all powerKeys")
	// 		}
	// 		for id, powerKeys := range token {
	// 			if _, exists := parcel.PowerKeys[id]; !exists {
	// 				parcel.PowerKeys[id] = powerKeys
	// 				cond.Signal() // notify waiters
	// 			}
	// 		}
	// 	case <-time.After(10 * time.Second):
	// 		return Parcel{}, fmt.Errorf("timeout fetching powerKeys")
	// 	}
	// }
	// mu.Unlock()

	return parcel, nil
}
