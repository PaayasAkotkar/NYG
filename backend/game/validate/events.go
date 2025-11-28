package validate

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"resty.dev/v3"
)

// Fetch fetches the sports data key will be the name of the book
func Fetch(book, dictionary string, List string) Parcel {
	var (
		resultChan = make(chan Parcel)
		// wg         sync.WaitGroup
		once     sync.Once
		postDone = make(chan struct{})
	)
	url := URL + book + "/" + dictionary + "/" + List
	App, Conn := resty.New(), resty.NewEventSource().SetURL(url)
	fmt.Println("url 🎴 ", url)
	fmt.Println("dictionary 🛩️ ", dictionary)
	fmt.Println("list 🏷️", List)
	// temp in-future they have to logged in
	register := Register{ID: "12x", RoomName: "GX", UserAgent: "", Location: "INDIA"}

	Conn.OnOpen(func(url string) {
		fmt.Println("🔥 connected to url: ", url)
	})

	// note: if you print the result it will give you empty but the result is fetched perfectly
	Conn.OnMessage(func(a any) {
		fmt.Println("🎧 on message triggered")
		token := []byte(a.(*resty.Event).Data)
		var parcel Parcel

		if err := json.Unmarshal(token, &parcel); err != nil {
			return
		}
		once.Do(func() {
			resultChan <- parcel
			fmt.Println("✅ on message done")
		})

	}, nil)

	go func() {
		if err := Conn.Get(); err != nil {
			fmt.Println("error")
		}
	}()

	// important to use the go func to avoid race condition
	go func() {
		resp, err := App.R().
			SetHeader("Content-Type", "application/json").SetBody(register).SetResult(&Parcel{}).Post(url)
		if err != nil {
			close(postDone)
			return
		}
		if !resp.IsSuccess() {
			close(postDone)
			return
		}

		close(postDone)

	}()

	select {
	case parcel := <-resultChan:
		<-postDone
		Conn.Close()
		return parcel
	case <-time.After(15 * time.Second):
		<-postDone
		Conn.Close()
		return Parcel{}
	}
}
