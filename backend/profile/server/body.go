package server

import (
	"app/server/graph/model"
	"app/sqlmanager"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type IIMG struct {
	ImgURL  string `json:"imgURL"`
	ImgName string `json:"imgName"`
	ID      string `json:"id"`
}
type UpgradeProgress struct {
	Powers     map[string]model.PowerProfile `json:"powerProfile"` // power-name -> ...
	PowersKeys []string                      `json:"powers"`       // keys to use to unlock the map
}

func ViewBody(id string) Body {
	var (
		body           Body
		img            *IIMG
		upgrades       UpgradeProgress
		devents        model.DailyEvents
		playerCredits  model.ICredits
		name           []byte
		nickname       []byte
		clashProfile   model.IProfile
		onevoneProfile model.IProfile
		twovtwoProfile model.IProfile
		bestRecord     model.IProfile
		store          model.Store
		deck           model.CurrentDeck
	)
	var refImg IIMG

	img = &refImg
	log.Println("req for id: ", id)
	_id, err := uuid.Parse(id)

	if err != nil {
		panic(err)
	}

	buid, err := _id.MarshalBinary()

	if err != nil {
		panic(err)
	}

	manager := sqlmanager.ConnectSQL{}

	addr := "127.0.0.1:3306"
	user := "root"
	pass := "kingp12"
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv(user)
	cfg.Passwd = os.Getenv(pass)
	cfg.Net = "tcp"
	cfg.Addr = addr
	cfg.DBName = "nygpatch"

	db, err := manager.Init("nygpatch", "_nygpatch_", cfg)

	if err != nil {
		panic(err)
	}
	defer db.CloseDB()
	if exists, err := db.HasID(buid); exists && err == nil {
		err = db.ExtractData("event", buid, &devents)
		if err != nil {
			panic(err)
		}

		err = db.ExtractData("playerCredits", buid, &playerCredits)
		if err != nil {
			panic(err)
		}

		err = db.ExtractData("img", buid, img)
		if err != nil {
			panic(err)
		}
		err = db.ExtractData("deck", buid, &deck)
		if err != nil {
			panic(err)
		}
		err = db.ExtractData("powerUpgradesProgress", buid, &upgrades)
		if err != nil {
			panic(err)
		}

		err = db.ExtractSingleData("name", buid, &name)
		if err != nil {
			panic(err)
		}

		err = db.ExtractSingleData("nickname", buid, &nickname)
		if err != nil {
			panic(err)
		}
		err = db.ExtractSingleDataFromJSON("twoVtwoProfile", "profile", buid, &twovtwoProfile)
		if err != nil {
			panic(err)
		}

		err = db.ExtractSingleDataFromJSON("oneVoneProfile", "profile", buid, &onevoneProfile)
		if err != nil {
			panic(err)
		}

		err = db.ExtractSingleDataFromJSON("clashProfile", "profile", buid, &clashProfile)
		if err != nil {
			panic(err)
		}
		db.ChangeTable("nygstore")
		smt, err := db.Conn.Query("select store from nygstore")
		if err != nil {
			panic(err)
		}
		var str string
		for smt.Next() {
			if err := smt.Scan(&str); err != nil {
				panic(err)
			}
		}
		if err := json.Unmarshal([]byte(str), &store); err != nil {
			panic(err)
		}

		body.ClashProfile = model.IProfile{}
		body.OneVOneProfile = model.IProfile{}
		body.TwoVTwoProfile = model.IProfile{}
		body.ClashProfile = clashProfile
		body.OneVOneProfile = onevoneProfile
		body.TwoVTwoProfile = twovtwoProfile
		body.Store = store
		bestRating := int32(0)
		body.Deck = deck
		if clashProfile.Rating > onevoneProfile.Rating || clashProfile.Rating > twovtwoProfile.Rating {
			bestRating = clashProfile.Rating
		} else if onevoneProfile.Rating > twovtwoProfile.Rating || onevoneProfile.Rating > clashProfile.Rating {
			bestRating = onevoneProfile.Rating
		} else if twovtwoProfile.Rating > onevoneProfile.Rating || twovtwoProfile.Rating > clashProfile.Rating {
			bestRating = twovtwoProfile.Rating
		} else if onevoneProfile.Rating == clashProfile.Rating || twovtwoProfile.Rating == onevoneProfile.Rating {
			bestRating = onevoneProfile.Rating
		}

		bestTier := ""
		totalGamesPlayed := clashProfile.GamesPlayed + onevoneProfile.GamesPlayed + twovtwoProfile.GamesPlayed
		totalPoints := clashProfile.Points + onevoneProfile.Points + twovtwoProfile.Points
		switch true {
		case bestRating < 1000:
			bestTier = "NOVICE"

		case bestRating >= 1000 && bestRating < 2000:
			bestTier = "MASTER"

		case bestRating >= 2000 && bestRating < 2500:
			bestTier = "GRAND MASTER"
		default:
			bestTier = "SUPER GRAND MASTER"
		}
		bestRecord.Tier = bestTier
		bestRecord.Rating = int32(bestRating)
		bestRecord.GamesPlayed = int32(totalGamesPlayed)
		bestRecord.Points = totalPoints
		// name = name[:8]
		body.DEvents = devents
		log.Println("SS: ", img.ImgName)
		if img.ID != "" {
			i := model.IImage{}
			i.ID = img.ID
			i.ImgName = img.ImgName
			i.ImgURL = img.ImgURL
			body.ImageData = &i
		} else {
			var de = "DEFAULT"
			i := model.IImage{}
			i.ID = id
			i.ImgName = de
			i.ImgURL = de
			body.ImageData = &i

		}
		body.Name = string(name)
		body.Nickname = string(nickname)
		body.ClashProfile.Tier = strings.ToUpper(clashProfile.Tier)
		body.OneVOneProfile.Tier = strings.ToUpper(onevoneProfile.Tier)
		body.TwoVTwoProfile.Tier = strings.ToUpper(twovtwoProfile.Tier)
		body.NYGProfile = bestRecord
		var temp []*model.PowerProfile
		for _, _key := range upgrades.PowersKeys {
			var z = upgrades.Powers[_key]
			temp = append(temp, &z)
		}
		body.Upgrades.Upgrades = temp
		body.Credits = playerCredits

		fill := model.Display{}

		fill.ClashProfile = &body.ClashProfile

		fill.NygProfile = &body.NYGProfile

		fill.OnevoneProfile = &body.OneVOneProfile

		fill.TwovtwoProfile = &body.TwoVTwoProfile

		fill.Name = body.Name
		fill.Nickname = body.Nickname

		if body.ImageData != nil {
			fill.Image = body.ImageData
		}

		fill.ID = body.ID
		body.Profile = fill
		log.Println("updagrades: ", upgrades.Powers)
		fmt.Println("body done")
		log.Println("deck: ", body.Deck)
		return body
	}
	return Body{}
}

// DecodeImageURL writes the image on the storage based on the location and given name
// note: the file must contain the extension
func DecodeImageURL(url []byte, createLocationAndname string) error {
	b64String := string(url)
	re := regexp.MustCompile(`\s+`)
	shave := re.ReplaceAllString(b64String, "")
	img, err := base64.StdEncoding.DecodeString(shave)
	if err != nil {
		return err
	}
	if err := os.WriteFile(createLocationAndname, img, 0644); err != nil {
		return err
	}
	return nil
}
