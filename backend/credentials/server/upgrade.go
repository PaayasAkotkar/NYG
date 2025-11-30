package server

import (
	"app/sqlmanager"
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PowerProfile struct {
	Power   string `json:"power"`
	Spur    int32  `json:"spur"`
	Counter int32  `json:"counter"`
	Level   int32  `json:"level"`
	Allow   bool   `json:"allow"`
	Display string `json:"display"`
}

type UpgradeProgress struct {
	Powers     map[string]PowerProfile `json:"powerProfile"` // power-name -> ...
	PowersKeys []string                `json:"powers"`       // keys to use to unlock the map
}
type Upgarade struct {
	Spur     string `json:"spur"`
	Coin     string `json:"coin"`
	Power    string `json:"power"` // must be the key of the power
	ID       string `json:"id"`
	NewLevel int    `json:"newLevel"`
}
type PowerUp struct {
	ID     string `json:"id"`
	Covert int    `json:"covert"` // int must be current level
	Nexus  int    `json:"nexus"`
	Freeze int    `json:"freeze"`
	Rewind int    `json:"rewind"`
	Draw   int    `json:"draw"`
	Tag    int    `json:"tag"`
	Bet    int    `json:"bet"`
}

func (up *UpgradeProgress) Value() (string, error) {
	x, err := json.Marshal(up)
	return string(x), err
}
func ServeUpgrades(ctx *gin.Context) {
	log.Println("NYG UPDATE SERVER RUNNING ⚙️")

	var req Upgarade
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		log.Println(err)
		return
	}
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.CloseDB()
	// updating the upgrades

	var pu PowerUp
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

	// upgrading credits
	var _src Credit

	err = db.ExtractData("playerCredits", id, &_src)
	if err != nil {
		log.Println(err)
		return
	}
	var cleanAccount = strings.ReplaceAll(_src.Spur, "S", "")
	var cleanPurchase = strings.ReplaceAll(req.Spur, "S", "")

	a, err := strconv.Atoi(cleanAccount)

	if err != nil {
		log.Println(err, "for src spur: ", _src.Spur)
		return
	}
	b, err := strconv.Atoi(cleanPurchase)
	if err != nil {
		log.Println(err, "for req spur: ", req.Spur)
		return
	}
	OCSlvl := 0
	if b < 60 {
		OCSlvl = 1
	} else if b >= 60 && b < 99 {
		OCSlvl = 2
	} else if b == 100 {
		OCSlvl = 3
	}

	newSpurAmount := a - b
	if newSpurAmount < 0 {
		newSpurAmount *= -1
	}

	cleanAccount = strings.ReplaceAll(_src.Coin, "C", "")
	cleanPurchase = strings.ReplaceAll(req.Coin, "C", "")
	a, err = strconv.Atoi(cleanAccount)

	if err != nil {
		log.Println(err, "for src coin: ", _src.Coin)
		return
	}
	b, err = strconv.Atoi(cleanPurchase)
	if err != nil {
		log.Println(err, "for req coin: ", req.Coin)
		return
	}

	var newCoinAmount = a - b
	if newCoinAmount < 0 {
		newCoinAmount *= -1
	}

	ns := strconv.Itoa(newSpurAmount) + "S"
	nc := strconv.Itoa(newCoinAmount) + "C"
	err = db.UpdateSingleJSONentry(id, "spurs", "playerCredits", ns)
	if err != nil {
		log.Println(err)
		return
	}
	err = db.UpdateSingleJSONentry(id, "coins", "playerCredits", nc)

	if err != nil {
		log.Println(err)
		return
	}

	err = db.ExtractData("upgrades", id, &pu)
	if err != nil {
		log.Println(err)
		return
	}

	Update(&pu, req.Power)
	var pp UpgradeProgress

	err = db.ExtractData("powerUpgradesProgress ", id, &pp)
	if err != nil {
		log.Println(err)
		return
	}

	for power, det := range pp.Powers {
		if strings.Contains(power, req.Power) {
			var src = det
			if req.NewLevel == 100 {
				src.Level = int32(OCSlvl)
				src.Allow = false
				src.Display = "MAX"
				src.Spur = 100
				src.Counter = 0
			} else {
				var s, err = strconv.Atoi(strings.ReplaceAll(req.Spur, "S", ""))
				if err != nil {
					return
				}
				src.Level = int32(OCSlvl)
				src.Spur = int32(s)
				src.Display = strconv.Itoa(req.NewLevel) + "S"
				src.Counter = int32(s)
			}

			pp.Powers[power] = src
			va, err := pp.Value()
			if err != nil {
				log.Println(err)
				return
			}

			err = db.UpdateWholeJSONentry(id, "powerUpgradesProgress", va)
			if err != nil {
				log.Println(err)
				return
			}
			break
		}
	}

	err = db.UpdateSingleJSONentry(id, req.Power, "upgrades", req.NewLevel)
	if err != nil {
		log.Println(err)
		return
	}

	// end updating the upgrades

}

// Update updates any struct field that matches the given field key
func Update(x *PowerUp, key string) {

	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Pointer {
		return
	}
	elemRef := v.Elem()
	stType := elemRef.Type()

	for t := range elemRef.NumField() {
		tag := stType.Field(t).Tag.Get("json")
		if tag == key {
			va := elemRef.Field(t)
			if va.CanSet() {
				if va.Kind() == reflect.String {
					va.SetString("crash")
					return
				}
			}
		}
	}
}
