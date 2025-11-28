package books

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"resty.dev/v3"
)

func PackEntertainment() Parcel {
	var done = make(chan struct{})
	var run sync.Once
	var parcel Parcel
	App, Conn := resty.New(), resty.NewEventSource().SetURL(FBURLS[1])

	// temp in-future they have to logged in
	register := Register{ID: "12x", RoomName: "GX", UserAgent: "", Location: "INDIA"}
	// important to use the go func to avoid race condition
	go func() {
		resp, _ := App.R().
			SetHeader("Content-Type", "application/json").SetBody(register).SetResult(&Parcel{}).Post(FBURLS[1])
		if !resp.IsSuccess() {
			return
		}
	}()

	Conn.OnOpen(func(url string) {
		fmt.Println("🔥 connected to url: ", url)
	})

	// note: if you print the result it will give you empty but the result is fetched perfectly
	Conn.OnMessage(func(a any) {
		fmt.Println("🎧 on message triggered")
		token := []byte(a.(*resty.Event).Data)

		if err := json.Unmarshal(token, &parcel); err != nil {
			panic(err)
		}
		run.Do(func() {
			close(done)
		})
		Conn.Close() // important to fetch only one time
		fmt.Println("✅ on message done")

	}, nil)
	go func() {
		if err := Conn.Get(); err != nil {
			close(done)
			panic(err)
		}
	}()

	<-done
	select {
	case <-done:
		return parcel
	case <-time.After(1 * time.Second):
		Conn.Close()

		return parcel
	}
}
