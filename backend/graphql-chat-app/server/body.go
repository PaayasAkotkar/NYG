package server

import (
	"app/server/graph/model"
	"app/sqlmanager"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type StoreBody struct {
	ID       string `json:"id"`
	Msg      string `json:"msg"`
	Roomname string `json:"roomname"`
	Name     string `json:"name"`
}
type Body struct {
	latest model.LatestMessage
}

func (s *StoreBody) Value() (string, error) {
	e, err := json.Marshal(s)
	return string(e), err
}

func Update(id string, msg model.LatestMessage) {
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "nygchat", cfg)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.CloseDB()
	tb := "nygchat"
	col := "messages"
	var latest StoreBody
	latest.ID = msg.ID
	latest.Msg = msg.Msg
	latest.Name = msg.Name
	latest.Roomname = msg.Roomname
	v, err := latest.Value()
	if err != nil {
		log.Println(err)
		return
	}
	x, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return
	}
	uuID, err := x.MarshalBinary()
	if err != nil {
		log.Println(err)
		return
	}
	if exists, err := db.HasID(uuID); !exists && err == nil {
		q := fmt.Sprintf("insert into %s (id, %s) values (?,?)", tb, col)
		if err := db.Prepare(q, &v); err != nil {
			log.Println(err)
			return
		}
	} else if err != nil {

		return
	} else {
		if err := db.UpdateSingleJSONentry(uuID, "msg", col, &latest.Msg); err != nil {
			log.Println(err)
			return
		}
		if err := db.UpdateSingleJSONentry(uuID, "name", col, &latest.Name); err != nil {
			log.Println(err)
			return
		}
		if err := db.UpdateSingleJSONentry(uuID, "roomname", col, &latest.Roomname); err != nil {
			log.Println(err)
			return
		}
	}
}

func ViewBody(id string) *Body {
	var (
		latest StoreBody
		body   Body
	)
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "nygchat", cfg)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.CloseDB()
	col := "messages"

	x, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return nil
	}
	uuID, err := x.MarshalBinary()
	if err != nil {
		log.Println(err)
		return nil
	}
	if exists, err := db.HasID(uuID); exists && err == nil {
		if err = db.ExtractData(col, uuID, &latest); err != nil {
			log.Println(err)
			return nil
		}
	} else {
		return nil
	}
	body.latest.ID = latest.ID
	body.latest.Msg = latest.Msg
	body.latest.Roomname = latest.Roomname
	return &body
}
