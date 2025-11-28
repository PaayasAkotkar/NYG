package server

import (
	"app/sqlmanager"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PatchDeck(ctx *gin.Context) {
	log.Println("deck")

	var store IDeck
	if err := ctx.ShouldBindBodyWithJSON(&store); err != nil {
		panic(err)
	}
	userID := store.ID
	_uuid, _ := uuid.Parse(*userID)
	BUID, _ := _uuid.MarshalBinary()
	log.Println("for id: ", string(BUID))
	m := sqlmanager.ConnectSQL{}

	cfg := Env()
	db, err := m.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	if !store.IsDefault {
		query := "UPDATE _nygpatch_ SET  deck = ? WHERE id = ?"
		va, _ := store.Value()
		db.Prepare(query, va, BUID)
		log.Println("for id: ", string(BUID))
	}
	log.Println("PROCESS DONE")
}
