package server

import (
	"app/sqlmanager"
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type IProfile struct {
	ID string `json:"id"`

	TwoVTwoProfile IProfile_ `json:"twoVtwoProfile"`
	OneVOneProfile IProfile_ `json:"oneVoneProfile"`
	ClashProfile   IProfile_ `json:"clashProfile"`
}

type IProfile_ struct {
	// Record      int    `json:"record"`
	Gamesplayed int    `json:"gamesPlayed"`
	Tier        string `json:"tier"`
	Rating      int    `json:"rating"`
	Points      int    `json:"points"`
}

func (p *IProfile_) Value() (string, error) {
	x, err := json.Marshal(&p)
	return string(x), err
}
func GetClashProfile(id string) (*IProfile_, error) {
	var t IProfile_
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.CloseDB()

	_id := id
	_uuiD, err := uuid.Parse(_id)

	if err != nil {
		log.Println("id: ", _id, err)
		return nil, err
	}

	id_, err := _uuiD.MarshalBinary()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	db.ExtractSingleDataFromJSON("clashProfile", "profile", &id_, &t)

	return &t, nil
}

func GetOnevOneProfile(id string) (*IProfile_, error) {
	var t IProfile_
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.CloseDB()

	_id := id
	_uuiD, err := uuid.Parse(_id)

	if err != nil {
		log.Println("id: ", _id, err)
		return nil, err
	}

	id_, err := _uuiD.MarshalBinary()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	db.ExtractSingleDataFromJSON("oneVoneProfile", "profile", &id_, &t)

	return &t, nil

}
func GetTwovTwoProfile(id string) (*IProfile_, error) {
	var t IProfile_
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.CloseDB()

	_id := id
	_uuiD, err := uuid.Parse(_id)

	if err != nil {
		log.Println("id: ", _id, err)
		return nil, err
	}

	id_, err := _uuiD.MarshalBinary()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	db.ExtractSingleDataFromJSON("twoVtwoProfile", "profile", &id_, &t)

	return &t, nil
}
