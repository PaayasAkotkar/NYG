package server

import "app/server/graph/model"

var (
	pubsub = NewPubSub()
)

type PubSub struct {
	pub                        chan bool
	sub                        chan string // sub for id
	pubUpdatedName             chan string
	pubUpdatedNickname         chan string
	pubUpdatedStore            chan *model.Store
	pubUpdatedPlayerCredits    chan *model.ICredits
	pubUpdatedEvents           chan *model.DailyEvents
	pubUpdatedProfile          chan *model.Display
	pubsubUpdatedPowerUpgrades chan *model.PowerProgress
	pubsubUpdatedDeck          chan *model.CurrentDeck
	pubsubUpdatedStore         chan *model.Store
}

func NewPubSub() *PubSub {
	return &PubSub{
		pub:                        make(chan bool),
		sub:                        make(chan string),
		pubUpdatedName:             make(chan string),
		pubUpdatedNickname:         make(chan string),
		pubUpdatedStore:            make(chan *model.Store),
		pubUpdatedPlayerCredits:    make(chan *model.ICredits),
		pubUpdatedEvents:           make(chan *model.DailyEvents),
		pubUpdatedProfile:          make(chan *model.Display),
		pubsubUpdatedPowerUpgrades: make(chan *model.PowerProgress),
		pubsubUpdatedDeck:          make(chan *model.CurrentDeck),
		pubsubUpdatedStore:         make(chan *model.Store),
	}
}
