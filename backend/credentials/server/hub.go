package server

import (
	"encoding/json"
	"log"
	"os"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func RunHub(h *Hub) {
	for {
		select {
		case coin := <-h.Coin:

			for send, ok := range coin {
				if ok {
					pack, err := json.Marshal(&send)
					if err != nil {
						log.Println(err)
						return
					}
					sigs := make(chan os.Signal, 1)
					done := make(chan bool, 1)

					go func() {
						<-sigs
						done <- true
					}()

					options := &mqtt.Options{}

					producer := mqtt.New(options)

					if err := producer.AddHook(new(auth.AllowHook), nil); err != nil {
						log.Println(err)
						return
					}

					tcp := listeners.NewWebsocket(listeners.Config{
						ID:      "ws1",
						Address: ":1883",
					})

					err = producer.AddListener(tcp)

					if err != nil {
						log.Println(err)
						return
					}

					go func() {
						if err := producer.Serve(); err != nil {
							log.Println(err)
							return
						}
					}()

					if err := producer.Publish("coin", pack, false, 0); err != nil {
						log.Println(err)
						return
					}

					<-done

					producer.Log.Warn("caught signal, stopping...")

					if err := producer.Close(); err != nil {
						return
					}

					producer.Log.Info("coin hub finished")

				}
			}
		case spur := <-h.Spur:
			for send, ok := range spur {
				if ok {
					pack, err := json.Marshal(&send)
					if err != nil {
						log.Println(err)
						return
					}
					sigs := make(chan os.Signal, 1)
					done := make(chan bool, 1)

					go func() {
						<-sigs
						done <- true
					}()

					options := &mqtt.Options{}

					producer := mqtt.New(options)

					if err := producer.AddHook(new(auth.AllowHook), nil); err != nil {
						log.Println(err)
						return
					}

					tcp := listeners.NewWebsocket(listeners.Config{
						ID:      "ws1",
						Address: ":1884",
					})

					err = producer.AddListener(tcp)

					if err != nil {
						log.Println(err)
						return
					}

					go func() {
						if err := producer.Serve(); err != nil {
							log.Println(err)
							return
						}
					}()

					if err := producer.Publish("spur", pack, false, 0); err != nil {
						log.Println(err)
						return
					}

					<-done

					producer.Log.Warn("caught signal, stopping...")

					if err := producer.Close(); err != nil {
						return
					}

					producer.Log.Info("spur hub finished")

				}
			}
		case upgrade := <-h.Upgarade:
			for send, ok := range upgrade {
				if ok {
					pack, err := json.Marshal(&send)
					if err != nil {
						log.Println(err)
						return
					}
					sigs := make(chan os.Signal, 1)
					done := make(chan bool, 1)

					go func() {
						<-sigs
						done <- true
					}()

					producer := mqtt.New(nil)

					if err := producer.AddHook(new(auth.AllowHook), nil); err != nil {
						log.Println(err)
						return
					}

					tcp := listeners.NewWebsocket(listeners.Config{
						ID:      "ws1",
						Address: ":1885",
					})

					err = producer.AddListener(tcp)

					if err != nil {
						log.Println(err)
						return
					}

					go func() {
						if err := producer.Serve(); err != nil {
							log.Println(err)
							return
						}
					}()

					if err := producer.Publish("upgrade", pack, true, 0); err != nil {
						log.Println(err)
						return
					}

					<-done

					producer.Log.Warn("caught signal, stopping...")

					if err := producer.Close(); err != nil {
						return
					}

					producer.Log.Info("upgrade hub finished")

				}
			}

		}
	}

}
