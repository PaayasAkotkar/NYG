package server

import (
	"app/sqlmanager"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Coin struct {
	Coin    string          `json:"coin"`
	ID      string          `json:"id"`
	Reciept PurchaseReciept `json:"reciept"`
}

func ServeCoin(ctx *gin.Context) {
	log.Println("coin server")
	var req Coin

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		panic(err)
	}
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.CloseDB()
	var src Credit
	i, err := uuid.Parse(req.ID)
	if err != nil {
		log.Println(err)
		return
	}
	id, err := i.MarshalBinary()

	if err != nil {
		log.Println(err)
		return
	}
	db.ExtractData("playerCredits", id, &src)
	var cleanAccount = strings.ReplaceAll(src.Coin, "C", "")
	var cleanPurchase = strings.ReplaceAll(req.Coin, "C", "")
	a, err := strconv.Atoi(cleanAccount)
	if err != nil {
		log.Println(err)
		return
	}
	b, err := strconv.Atoi(cleanPurchase)
	if err != nil {
		log.Println(err)
		return
	}
	newAmount := a + b
	na := strconv.Itoa(newAmount) + "C"
	err = db.UpdateSingleJSONentry(id, "coins", "playerCredits", na)
	if err != nil {
		log.Println(err)
		return
	}

	var c = req
	c.Coin = strconv.Itoa(newAmount) + "C"

	go func() {
		pass := map[Coin]bool{}
		pass[c] = true
		h.Coin <- pass
	}()

	// end updating credits

	// add recipets
	var p Purchase
	p.Purchases = append(p.Purchases, req.Reciept)
	exists, err := db.HasID(id)
	if err != nil {
		log.Println(err)
		return
	}
	db.ChangeTable("nygcredentials")
	if !exists {
		q := `
		INSERT INTO nygcredentials
		(id,purchase)
		VALUES (?, ?)`
		v, err := p.Value()
		if err != nil {
			log.Println(err)
			return
		}
		err = db.Prepare(q, req.ID, v)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		var p1 Purchase
		err = db.ExtractData("purhcase", id, p1)
		if err != nil {
			log.Println(err)
			return
		}
		p.Purchases = append(p.Purchases, p1.Purchases...)
		v, err := p.Value()
		if err != nil {
			log.Println(err)
			return
		}
		err = db.UpdateWholeJSONentry(id, "purchase", v)
		if err != nil {
			log.Println(err)
			return
		}
	}
	// end adding reciepts
}
