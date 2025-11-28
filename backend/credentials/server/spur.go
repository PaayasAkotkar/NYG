package server

import (
	"app/sqlmanager"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Spur struct {
	Spur    string          `json:"spur"`
	ID      string          `json:"id"`
	Reciept PurchaseReciept `json:"reciept"`
}

func ServeSpur(ctx *gin.Context) {
	log.Println("spur server")
	var req Spur

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		panic(err)
	}
	log.Println("body: ", req)

	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.CloseDB()
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

	// updating credits
	var src Credit
	if err := db.ExtractData("playerCredits", id, &src); err != nil {
		log.Println(err)
		return
	}
	log.Println("playersCredit: ", src)
	var cleanAccount = strings.ReplaceAll(src.Spur, "S", "")
	var cleanPurchase = strings.ReplaceAll(req.Spur, "S", "")
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
	na := strconv.Itoa(newAmount) + "S"
	err = db.UpdateSingleJSONentry(id, "spurs", "playerCredits", na)

	if err != nil {
		log.Println(err)
		return
	}

	var s = req
	s.Spur = strconv.Itoa(newAmount) + "S"
	go func() {
		pass := map[Spur]bool{}
		pass[s] = true
		h.Spur <- pass
	}()
	// end updating credits

	// store the reciept
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
	// end store reciept
}
