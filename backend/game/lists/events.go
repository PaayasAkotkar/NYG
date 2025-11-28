package list

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"resty.dev/v3"
)

// FetchSports fetches the sports data key will be the name of the book
func FetchSports(book, dictionary string) Parcel {
	var (
		resultChan = make(chan Parcel, 1)
		wg         sync.WaitGroup
		once       sync.Once
		postDone   = make(chan struct{})
	)
	url := URLs + book + "/" + dictionary
	fmt.Println("url: ", url)
	fmt.Println("book: ", book)
	fmt.Println("dictionary: ", dictionary)

	App, Conn := resty.New(), resty.NewEventSource().SetURL(url)

	// temp in-future they have to logged in
	register := Register{ID: "12x", RoomName: "GX", UserAgent: "", Location: "INDIA"}

	Conn.OnOpen(func(url string) {
		fmt.Println("🔥 connected to url: ", url)
	})
	Conn.OnMessage(func(a any) {
		fmt.Println("🎧 on message triggered")
		event := a.(*resty.Event)
		token := []byte(event.Data)
		var parcel Parcel

		if err := json.Unmarshal(token, &parcel); err != nil {
			fmt.Println("❌ JSON unmarshal error:", err)
			return
		}
		fmt.Println("fetcehd: ", parcel)
		once.Do(func() {
			resultChan <- parcel
			fmt.Println("✅ on message done")
		})
	}, nil)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := Conn.Get(); err != nil {
			fmt.Println("errr")
		}
	}()

	// important to use the go func to avoid race condition
	go func() {
		resp, _ := App.R().
			SetHeader("Content-Type", "application/json").SetBody(register).SetResult(&Parcel{}).Post(url)
		if !resp.IsSuccess() {
			close(postDone)
			return
		} else {
			close(postDone)
		}
	}()

	select {
	case parcel := <-resultChan:
		<-postDone
		Conn.Close()
		return parcel
	case <-time.After(10 * time.Second):
		<-postDone
		Conn.Close()
		return Parcel{}
	}
}
