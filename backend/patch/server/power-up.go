package server

import (
	"app/sqlmanager"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PatchPowerUp(ctx *gin.Context) {
	log.Println("powerup")

	var store IPowerUp

	if err := ctx.ShouldBindBodyWithJSON(&store); err != nil {
		panic(err)
	}
	typ := reflect.TypeOf(store)
	val := reflect.ValueOf(store)
	var setClauses []string
	var args []any

	for i := range val.NumField() {
		n := typ.Field(i).Tag.Get("json")
		v := val.Field(i)
		setClauses = append(setClauses, n+" =?")
		args = append(args, v)
	}
	userID, _ := uuid.Parse(*store.ID)
	uuID := userID[:]

	query := "UPDATE _nygpatch_ SET " + strings.Join(setClauses, ", ") + " WHERE id = ?"
	m := sqlmanager.ConnectSQL{}

	cfg := Env()
	db, err := m.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		panic(err)
	}
	db.Prepare(query, uuID, args)
}
