package server

import (
	"app/sqlmanager"
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

func TestHasID(t *testing.T) {
	id := "c8aaee39-9a55-4386-909a-28c4a5ac6321"
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("for id: ", id)
	_uuid, _ := uuid.Parse(id)
	BUID, _ := _uuid.MarshalBinary()

	exists, err := db.HasID(BUID)
	if err != nil {
		log.Println(err)
		return
	}
	var name []byte
	err = db.ExtractSingleData("name", BUID, &name)
	if err != nil {
		log.Println(err)
		return
	}
	type s struct {
		ID     string `json:"id"`
		Rating int    `json:"rating"`
	}
	var view s
	err = db.ExtractData("profile", BUID, &view)
	if err != nil {
		log.Println(err)
		return
	}

	assert.Equal(t, true, exists)
	// assert.Equal(t, "", name)
	assert.Equal(t, 400, view.Rating)
}
