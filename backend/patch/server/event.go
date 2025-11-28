package server

import (
	"app/sqlmanager"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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

type Events struct {
	Clash  []IDailyEvent `json:"clash"`
	Locify []IDailyEvent `json:"locify"`
}

func (e *Events) Value() (string, error) {
	j, err := json.Marshal(e)
	return string(j), err
}
func PutEvent(ctx *gin.Context) {
	log.Println("in put event")
	type s struct {
		ID string `json:"id"`
	}
	var store s
	if err := ctx.ShouldBindBodyWithJSON(&store); err != nil {
		log.Println(err)
	}

	userID, _ := uuid.Parse(store.ID)
	uuID := userID[:]

	if userID == uuid.Nil {
		log.Println("nil id")
		return
	}
	GenerateEvent(uuID)
}

func GenerateEvent(id any) {

	m := sqlmanager.ConnectSQL{}

	cfg := Env()
	db, err := m.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return

	}

	defer db.CloseDB()

	games := 5

	clash := make([]IDailyEvent, 3)
	locify := make([]IDailyEvent, 3)

	var d IDailyEvent
	d.EventName = "CLASH ON"
	d.Completed = 0
	d.Progress = 0
	d.IsClash = true

	d.Reward = "50C"
	d.ToComplete = games
	d.Description = fmt.Sprintf("WIN %d GAME", games)
	clash[0] = d

	games = 7
	d.ToComplete = games
	d.Reward = "60C"
	d.Description = fmt.Sprintf("WIN %d GAME", games)
	clash[1] = d

	games = 10
	d.ToComplete = games
	d.Reward = "80C"
	d.Description = fmt.Sprintf("WIN %d GAME", games)
	clash[2] = d

	games = 5
	d.ToComplete = games
	d.Description = fmt.Sprintf("WIN %d GAME", games)
	d.IsClash = false
	d.EventName = "SPORTS"
	d.Reward = "70C"
	locify[0] = d

	games = 7
	d.ToComplete = games
	d.Description = fmt.Sprintf("WIN %d GAME", games)
	d.EventName = "ENTERTAINMENT"
	locify[1] = d

	games = 10
	d.Reward = "90C"
	d.ToComplete = games
	d.EventName = "SPORTS"
	d.Description = fmt.Sprintf("WIN %d GAME", games)
	locify[2] = d

	var IsSet []byte
	var _token Events
	_token.Clash = []IDailyEvent{}
	_token.Locify = []IDailyEvent{}

	_token.Clash = clash
	_token.Locify = locify
	v, err := _token.Value()
	log.Println("clash: ", _token.Clash)
	if err != nil {
		log.Println(err)
		return
	}

	if err := db.ExtractSingleData("eventSet", id, &IsSet); err != nil {
		log.Println(err)
		return
	}

	if set_, err := strconv.ParseBool(string(IsSet)); !set_ && err == nil {

		log.Println("event send")
		q := `update _nygpatch_ set event = ? where id =?`

		err := db.Prepare(q, v, id)
		if err != nil {
			log.Println(err)
			return

		}

		q = `update _nygpatch_ set eventSet = ? where id=?`

		err = db.Prepare(q, true, id)
		if err != nil {
			log.Println(err)
			return

		}

	} else {
		log.Println("event already been set")
		return

	}
}
