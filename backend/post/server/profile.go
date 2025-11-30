package server

import (
	"app/sqlmanager"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type ParcelProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickname"`
}

func GetProfile(id string) ICreate {
	log.Println("in get profile")
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return ICreate{}
	}
	_uuid, _ := uuid.Parse(id)
	BUID, _ := _uuid.MarshalBinary()
	var prof ICreate

	exists, err := db.HasID(BUID)
	if err != nil {
		log.Println(err)
		return ICreate{}
	}

	// select values only if exists
	if exists {
		var name, nickname, _id []byte
		err = db.ExtractSingleData("name", BUID, &name)
		if err != nil {
			log.Println(err)
			return ICreate{}
		}

		err = db.ExtractSingleData("nickname", BUID, &nickname)
		if err != nil {
			log.Println(err)
			return ICreate{}
		}

		err = db.ExtractSingleData("id", BUID, &_id)
		if err != nil {
			log.Println(err)
			return ICreate{}
		}
		var n = string(name)
		prof.Name = n
		prof.Nickname = string(nickname)
		prof.ID = string(_id)

	} else {
		fmt.Println("NULLLLLLL 🤢 id : ", id)
		return ICreate{}
	}
	return prof
}
