package server

import (
	"app/sqlmanager"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type ParcelDeck struct {
	PowerKeys []string `json:"powerKeys"`
	ID        string   `json:"id"`
}

func GetDeck(id string) IDeck {
	log.Println("in deck")
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return IDeck{}
	}
	_uuid, _ := uuid.Parse(id)
	BUID, _ := _uuid.MarshalBinary()
	var deck IDeck
	fmt.Println("for id: ", string(BUID))

	exists, err := db.HasID(BUID)

	if err != nil {
		log.Println(err)
		return IDeck{}
	}

	// select values only if exists
	if exists {
		// err := db.QueryRow("select deck from _nygpatch_ where id = ?", BUID).Scan(&token)
		// if err != nil {
		// 	log.Println(err) ;return
		// }
		err = db.ExtractData("deck", BUID, &deck)
		if err != nil {
			log.Println(err)
			return IDeck{}
		}
	} else {
		fmt.Println("NULLLLLLL 🤢 id : ", id)
	}
	fmt.Println("deck: ", deck)
	return deck
}
