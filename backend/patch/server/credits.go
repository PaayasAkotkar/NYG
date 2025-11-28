package server

import (
	"app/sqlmanager"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Credit struct {
	Coin       string `json:"coins"`
	Spur       string `json:"spurs"`
	BoardTheme string `json:"boardTheme"`
	GuessTheme string `json:"guessTheme"`

	ID string `json:"id"`
}

func (d Credit) Value() (string, error) {
	e, err := json.Marshal(d)
	return string(e), err
}

func PatchCoin(ctx *gin.Context) {
	var store Credit
	if err := ctx.ShouldBindBodyWithJSON(&store); err != nil {
		panic(err)
	}
	CCredits(store.ID, store, "_nygpatch_")
}
func PatchSpur(ctx *gin.Context) {
	var store Credit
	if err := ctx.ShouldBindBodyWithJSON(&store); err != nil {
		panic(err)
	}
	SCredits(store.ID, store, "_nygpatch_")
}
func PatchCredits(ctx *gin.Context) {
	var store Credit
	if err := ctx.ShouldBindBodyWithJSON(&store); err != nil {
		panic(err)
	}
	CCredits(store.ID, store, "_nygpatch_")
}

func PatchTheme(ctx *gin.Context) {
	log.Println("in patch theme")
	var store Credit
	if err := ctx.ShouldBindBodyWithJSON(&store); err != nil {
		panic(err)
	}

	_uuid, _ := uuid.Parse(store.ID)
	BUID, _ := _uuid.MarshalBinary()
	TCredits(BUID, store, "_nygpatch_")
}

// WCredits updates the whole json
func WCredits(id any, playerCredits Credit, table string) {

	c := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := c.Init("nygpatch", table, cfg)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	q := fmt.Sprintf("update %s set playerCredits = ?", table)
	db.Prepare(q, playerCredits)
}

func CCredits(id any, playerCredits Credit, table string) {
	c := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := c.Init("nygpatch", table, cfg)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()
	db.UpdateSingleJSONentry(id, "coins", "playerCredits", playerCredits.Coin)

}

func SCredits(id any, playerCredits Credit, table string) {
	c := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := c.Init("nygpatch", table, cfg)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()
	db.UpdateSingleJSONentry(id, "coins", "playerCredits", playerCredits.Spur)

}

func TCredits(id any, playerCredits Credit, table string) {
	// text11x and text22x are the themes to unlock the scoreboard and guess text
	c := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := c.Init("nygpatch", table, cfg)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	db.UpdateSingleJSONentry(id, "boardTheme", "playerCredits", playerCredits.BoardTheme)
	db.UpdateSingleJSONentry(id, "guessTheme", "playerCredits", playerCredits.GuessTheme)
}
