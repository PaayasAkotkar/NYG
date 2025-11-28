package server

import (
	"app/sqlmanager"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IStore struct {
	Price string `json:"Price"`
	Title string `json:"Title"`
}
type Store struct {
	Coin []IStore `json:"coin"`
	Spur []IStore `json:"spur"`
}

func (s *Store) Value() (string, error) {
	n, err := json.Marshal(s)
	return string(n), err
}
func GenerateStore() {
	log.Println("in store")
	var _spurs = []IStore{
		{Price: "100S", Title: "SPUR"},
		{Price: "250S", Title: "SPUR"},
		{Price: "450S", Title: "SPUR"},
		{Price: "500S", Title: "SPUR"},
	}
	var _coins = []IStore{
		{Price: "100C", Title: "COIN"},
		{Price: "250C", Title: "COIN"},
		{Price: "450C", Title: "COIN"},
		{Price: "500C", Title: "COIN"},
	}
	store := Store{}
	store.Coin = _coins
	store.Spur = _spurs
	cfg := Env()
	conn := sqlmanager.ConnectSQL{}
	db, err := conn.Init("nygpatch", "nygstore", cfg)
	if err != nil {
		log.Println(err)
		return
	}
	q := `select updated from nygstore`
	stm, err := db.Conn.Query(q)
	if err != nil {
		log.Println(err)
		return
	}
	var updated bool
	for stm.Next() {
		if err := stm.Scan(&updated); err != nil {
			log.Println(err)
			return
		}
	}
	_myID := uuid.NewString()

	x, err := uuid.Parse(_myID)
	if err != nil {
		log.Println(err)
		return
	}
	myID := x[:]
	val, err := store.Value()
	if err != nil {
		log.Println(err)
	}
	if !updated {
		q := `insert into nygstore (store,offers,id,updated) values(?, NULL, ?, ?)`
		if err := db.Prepare(q, val, myID, true); err != nil {
			log.Println(err)
			return
		}
	} else {
		log.Println("already updated")
	}
}

func CleanName(name string) string {
	re := regexp.MustCompile("_|-|[^a-zA-Z]")
	var a = re.ReplaceAllString(name, "")
	if len(a) > 9 {
		return a[:8]
	}
	return a
}

type CookieProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

var (
	mu sync.Mutex
)

func InitProfile(ctx *gin.Context) {
	log.Println("init profie")
	mu.Lock()
	defer mu.Unlock()

	var store ICreate
	if err := ctx.ShouldBindBodyWithJSON(&store); err != nil {
		log.Println("form binding error:", err)
		ctx.JSON(400, gin.H{"error": "Invalid upload"})
		return
	}

	userID, _ := uuid.Parse(*store.ID)
	uuID := userID[:]

	if userID == uuid.Nil {
		log.Println("nil id", store.ID)
		log.Println("nil id", string(uuID))
		return
	}

	f := "nexus"
	s := "bet"
	max := 100
	_d := 0
	_r := 400
	_t := "novice"

	var create IGen
	create.ID = store.ID

	var n string
	n = CleanName(*store.Name)
	create.Name = &n
	n = CleanName(*store.Nickname)
	create.Nickname = &n
	log.Println("new nickname and name: ", *create.Nickname, *create.Name)

	create.Deck = &IDeck{}
	create.Deck.Frist = &f
	create.Deck.Second = &s
	create.Deck.ID = store.ID

	create.Ugrades = &IPowerUp{}
	create.Ugrades.ID = store.ID
	var pp PowerProfile
	storeUP := map[string]PowerProfile{}

	pp.Allow = false
	pp.CurrentCounter = 0
	pp.Display = "MAX"
	pp.DonatedSpur = 100
	pp.Level = 3

	pp.Power = "COVERT"
	storeUP["covert"] = pp
	pp.Power = "DRAW"
	storeUP["draw"] = pp
	pp.Power = "TAG"
	storeUP["tag"] = pp
	pp.Power = "BET"
	storeUP["bet"] = pp
	pp.Power = "REWIND"
	storeUP["rewind"] = pp

	pp.Allow = true
	pp.CurrentCounter = 0
	pp.Display = "S"
	pp.DonatedSpur = 0
	pp.Level = 1
	pp.Power = "FREEZE"
	storeUP["freeze"] = pp
	pp.Power = "NEXUS"
	storeUP["nexus"] = pp
	keys := []string{"nexus", "covert", "freeze", "rewind", "draw", "tag", "bet"}

	create.PowerUpgradeProgress = &UpgradeProgress{}
	create.PowerUpgradeProgress.Powers = storeUP
	create.PowerUpgradeProgress.PowersKeys = keys
	create.Ugrades.Bet = &max
	create.Ugrades.Draw = &max
	create.Ugrades.Covert = &max
	create.Ugrades.Rewind = &max
	create.Ugrades.Tag = &max
	create.Ugrades.Freeze = &_d
	create.Ugrades.Nexus = &_d

	create.Profile = &IProfile{}
	create.Profile.ClashProfile = &IProfile_{}
	create.Profile.OneVOneProfile = &IProfile_{}
	create.Profile.TwoVTwoProfile = &IProfile_{}

	create.Profile.ClashProfile.Gamesplayed = &_d
	create.Profile.ClashProfile.Points = &_d
	create.Profile.ClashProfile.Rating = &_r
	create.Profile.ClashProfile.Tier = &_t

	create.Profile.OneVOneProfile.Gamesplayed = &_d
	create.Profile.OneVOneProfile.Points = &_d
	create.Profile.OneVOneProfile.Rating = &_r
	create.Profile.OneVOneProfile.Tier = &_t

	create.Profile.TwoVTwoProfile.Gamesplayed = &_d
	create.Profile.TwoVTwoProfile.Points = &_d
	create.Profile.TwoVTwoProfile.Rating = &_r
	create.Profile.TwoVTwoProfile.Tier = &_t

	create.Credentials = &Credit{}
	create.Credentials.ID = *store.ID

	create.Credentials.Coin = "150C"
	create.Credentials.Spur = "5S"
	create.Credentials.BoardTheme = "text11x"
	create.Credentials.GuessTheme = "text11x"

	log.Println("connected 🔥")
	log.Println("req id: ", *store.ID)
	m := sqlmanager.ConnectSQL{}

	cfg := Env()
	db, err := m.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return
	}

	defer db.CloseDB()

	// imp: in order to not disturb the primary key

	exists, err := db.HasID(uuID)
	if err != nil {
		log.Println(err)
		return
	}

	GenerateGameCredits(*store.ID)

	log.Println("exists: ", exists)
	if !exists {
		log.Println("inserting")
		q := `
		INSERT INTO _nygpatch_
		(id, name, nickname, deck, profile, upgrades,
		 img,playerCredits,powerUpgradesProgress)
		VALUES (?, ?, ?, ?, ?, ?, NULL,?,?)`

		v, err := create.Deck.Value()

		if err != nil {
			log.Println(err)
			return
		}

		v1, err := create.Profile.Value()

		if err != nil {
			log.Println(err)
			return
		}

		v2, err := create.Ugrades.Value()

		if err != nil {
			log.Println(err)
			return
		}

		v3, err := create.Credentials.Value()

		if err != nil {
			log.Println(err)
			return
		}

		v4, err := create.PowerUpgradeProgress.Value()

		if err != nil {
			log.Println(err)
			return
		}

		err = db.Prepare(q, uuID, create.Name, create.Nickname, v, v1, v2, v3, v4)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("storing id: ", string(uuID))
		GenerateStore()

	} else {
		log.Println("client already exists with 🆔 ", string(uuID))
		return
	}

	log.Println("profile done")
}

type UpdateGameCredits struct {
	ID          string   `json:"id"`
	Books       []string `json:"books"`
	Categories  []string `json:"categories"`
	Powers      []string `json:"powers"`
	Points      int32    `json:"points"`
	GamesWon    int32    `json:"gamesWon"`
	GamesPlayed int32    `json:"gamesPlayed"`
	Rating      int32    `json:"rating"`
}

func (u UpdateGameCredits) GetValue() (string, error) {
	b, err := json.Marshal(u)
	return string(b), err
}

func GenerateGameCredits(id string) bool {
	log.Println("in generate game credits")
	var ugc UpdateGameCredits
	ugc.ID = id
	ugc.Books = nil
	ugc.Categories = nil
	ugc.Powers = nil
	ugc.Points = 0
	ugc.GamesWon = 0
	ugc.GamesPlayed = 0
	ugc.Rating = 400

	userID, _ := uuid.Parse(id)
	uuID := userID[:]

	if userID == uuid.Nil {
		log.Println("nil id")
		return false
	}

	cfg := Env()
	tb := "nyg_gamecredentials"

	conn := sqlmanager.ConnectSQL{}
	db, err := conn.Init("nygpatch", tb, cfg)
	if err != nil {
		log.Println(err)
	}

	if exists, err := db.HasID(&uuID); !exists && err == nil {
		q := fmt.Sprintf("INSERT INTO %s (id, credits) VALUES (?, ?)", tb)
		v, err := ugc.GetValue()
		if err != nil {
			log.Println(err)
			return false
		}
		if err := db.Prepare(q, uuID, v); err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	return false

}
