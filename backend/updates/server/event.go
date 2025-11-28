package server

import (
	"app/sqlmanager"
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type IDailyEvent struct {
	EventName   string `json:"eventName"`
	Reward      string `json:"reward"`
	Progress    int    `json:"progress"`
	ToComplete  int    `json:"toComplete"`
	Completed   int    `json:"completed"`
	IsClash     bool   `json:"isClash"`
	Description string `json:"description"`
}

func (d *IDailyEvent) Value() (string, error) {
	x, err := json.Marshal(&d)
	return string(x), err
}

type Events struct {
	Clash  []IDailyEvent `json:"clash"`
	Locify []IDailyEvent `json:"locify"`
}

type LocifyGameProfile struct {
	// note: these must be in capital letter
	Category string `json:"category"`
	Book     string `json:"book"`
	ID       string `json:"id"`
}

func UpdateLocifyEvent(id string, body LocifyGameProfile) error {

	conn := sqlmanager.ConnectSQL{}
	var e Events

	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.CloseDB()
	_id, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return err
	}
	buid, err := _id.MarshalBinary()
	if err != nil {
		log.Println(err)
		return err
	}
	if err := db.ExtractData("event", buid, &e); err != nil {
		log.Println(err)
		return err
	}

	for i, event := range e.Locify {
		if event.EventName == body.Category {
			var src = e.Locify[i]
			src.Completed += 1
			src.Progress = (src.Completed / src.ToComplete) * 100
			e.Locify[i] = src
		}
	}

	if err := db.UpdateSingleJSONentry(buid, "locify", "event", e.Locify); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateClash(id string) error {
	var e Events
	conn := sqlmanager.ConnectSQL{}

	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.CloseDB()
	_id, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return err
	}
	buid, err := _id.MarshalBinary()
	if err != nil {
		log.Println(err)
		return err
	}
	if err := db.ExtractData("event", buid, &e); err != nil {
		log.Println(err)
		return err
	}

	for i := range e.Clash {
		var src = e.Clash[i]
		src.Completed += 1
		src.Progress = (src.Completed / src.ToComplete) * 100
		e.Locify[i] = src
	}
	if err := db.UpdateSingleJSONentry(buid, "clash", "event", e); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
