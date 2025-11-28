package server

import (
	"fmt"

	"github.com/tmaxmax/go-sse"
)

func RunHub(h *Hub, conn *sse.Server) {
	fmt.Println("🐎 running")
	for token := range h.trigger {
		fmt.Println("trigger 🏛️")
		fmt.Println("got: ", token)
		fmt.Println("key: ", h.key)

		// important to avoid coroutine race
		// h.wg.Add(1)
		// go func() {
		// 	defer h.wg.Done()

		switch true {

		// beginning of dictionaries channel
		case token[h.key][EventKey]:
			fmt.Println("🏏")
			switch h.bookKey {
			case "cricket":
				CompleteCricketEvents(conn, h)
			case "basketball":
				CompleteBasketballEvents(conn, h)
			case "football":
				CompleteFootballEvents(conn, h)
			case "movies":
				CompleteMoviesEvents(conn, h)
			case "music":
				CompleteMusicEvents(conn, h)
			case "shows":
				CompleteShowsEvents(conn, h)
			}
			// end of the dictionaries channel
			h.key = ""
			h.dictionaryKey = ""
			h.listKey = ""

			// beginning of items channel
		case token[h.key][ItemKey]:
			fmt.Println("◽")
			fmt.Println("dictionary key 🔑 ", h.dictionaryKey)

			switch h.bookKey {
			case "cricket", "CRICKET":
				CompleteCricketLists(conn, h, h.dictionaryKey)
			case "basketball":
				CompleteBasketballLists(conn, h, h.dictionaryKey)
			case "football":
				CompleteFootballLists(conn, h, h.dictionaryKey)
			case "movies":
				CompleteMoviesLists(conn, h, h.dictionaryKey)
			case "music":
				CompleteMusicLists(conn, h, h.dictionaryKey)
			case "shows":
				CompleteShowsLists(conn, h, h.dictionaryKey)
			}
			// end of the items channel
			h.key = ""
			h.dictionaryKey = ""
			h.listKey = ""

			// beginning of validation channel
		case token[h.key][ValidateKey]:
			fmt.Println("🤽‍♂️")
			fmt.Println("dictionary key 🔑 ", h.dictionaryKey)
			fmt.Println("list key 🔑 ", h.listKey)

			switch h.bookKey {
			case "cricket":
				CompleteCricketValidateSheet(conn, h, h.dictionaryKey, h.listKey)
			case "basketball":
				CompleteBasketballValidateSheet(conn, h, h.dictionaryKey, h.listKey)
			case "football":
				CompleteFootballValidateSheet(conn, h, h.dictionaryKey, h.listKey)

			case "movies":
				CompleteMoviesValidateSheet(conn, h, h.dictionaryKey, h.listKey)
			case "music":
				CompleteMusicValidateSheet(conn, h, h.dictionaryKey, h.listKey)
			case "shows":
				CompleteShowsValidateSheet(conn, h, h.dictionaryKey, h.listKey)
			}
			h.key = ""
			h.dictionaryKey = ""
			h.listKey = ""

			// end of the validation channel
		}
		// }()
	}
}
