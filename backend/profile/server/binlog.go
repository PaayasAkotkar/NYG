package server

import (
	"log"
	"reflect"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/google/uuid"
)

type EVT struct {
	canal.DummyEventHandler
}

type ChangeDetectionRef struct {
	ActionColName  string
	ActionColIndex int
	Datas          []string
}

func (h *EVT) OnRow(e *canal.RowsEvent) error {
	var id string
	inx := 0
	cdr := ChangeDetectionRef{}

	for i, col := range e.Table.Columns {
		cdr.Datas = append(cdr.Datas, col.Name)
		if col.Name == "id" {
			inx = i
		}
	}

	for _, row := range e.Rows {
		switch v := row[inx].(type) {
		case []byte:
			u, err := uuid.FromBytes(v)
			if err != nil {
				panic(err)
			}
			id = u.String()
		case string:
			id = v
		}
	}
	_id := []byte(id)
	uID, err := uuid.FromBytes(_id)
	if err != nil {
		panic(err)
	}

	body := ViewBody(uID.String())

	switch e.Table.Name {
	case "_nygpatch_":
		if e.Action == canal.UpdateAction {
			for i := 0; i < len(e.Rows); i += 2 {
				before := e.Rows[i]
				after := e.Rows[i+1]

				for j, col := range e.Table.Columns {
					beforeValue := before[j]
					afterValue := after[j]
					if !reflect.DeepEqual(beforeValue, afterValue) {
						cdr.ActionColName = col.Name
					}

				}
			}

			switch cdr.ActionColName {
			case "playerCredits":
				pubsub.pubUpdatedPlayerCredits <- &body.Credits
			case "name":
				pubsub.pubUpdatedName <- body.Name
			case "nickname":
				pubsub.pubUpdatedNickname <- body.Nickname
			case "deck":
				pubsub.pubsubUpdatedDeck <- &body.Deck
			case "profile":
				pubsub.pubUpdatedProfile <- &body.Profile
			case "powerUpgradesProgress":
				pubsub.pubsubUpdatedPowerUpgrades <- &body.Upgrades
			case "events":
				pubsub.pubUpdatedEvents <- &body.DEvents

			}

		}
	case "nygstore":
		if e.Action == canal.UpdateAction {
			for i := 0; i < len(e.Rows); i += 2 {
				before := e.Rows[i]
				after := e.Rows[i+1]

				for _, col := range e.Table.Columns {
					if !reflect.DeepEqual(before, after) {
						cdr.ActionColName = col.Name
					}

				}
			}
			switch cdr.ActionColName {
			case "store":
				pubsub.pubUpdatedStore <- &body.Store
			case "offer":
			default:
			}
		}
	}
	return nil
}

func Watch() {
	mu.Lock()
	defer mu.Unlock()

	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = USER
	cfg.Password = PASSWORD
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "nygpatch"
	cfg.Dump.Tables = []string{"_nygpatch_"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	c.SetEventHandler(&EVT{})
	go func() {
		if err := c.Run(); err != nil {
			return
		}
	}()
	wg.Wait()

}
