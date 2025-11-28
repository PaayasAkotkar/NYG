package server

import (
	nygpostprotoc "app/nygprotoc"
	"app/sqlmanager"
	"log"

	"github.com/google/uuid"
)

type Credit struct {
	ImageURL       string                           `json:"imgULR"`
	Coin           string                           `json:"coins"`
	Spur           string                           `json:"spurs"`
	BoardTheme     string                           `json:"boardTheme"`
	GuessTheme     string                           `json:"guessTheme"`
	ID             string                           `json:"id"`
	Nexus          int32                            `json:"nexus"`
	Freeze         int32                            `json:"freeze"`
	ClashProfile   nygpostprotoc.TypeNYGGameProfile `json:"clashProfile"`
	OneVOneProfile nygpostprotoc.TypeNYGGameProfile `json:"oneVoneProfile"`
	TwoVTwoProfile nygpostprotoc.TypeNYGGameProfile `json:"twoVtwoProfile"`
}

func GetCredits(id string) *Credit {
	log.Println("in credits")
	conn := sqlmanager.ConnectSQL{}
	cfg := Env()
	token := Credit{}

	db, err := conn.Init("nygpatch", "_nygpatch_", cfg)
	if err != nil {
		log.Println(err)
		return nil
	}

	x, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return nil
	}

	BUID, err := x.MarshalBinary()
	if err != nil {
		log.Println(err)
		return nil
	}
	err = db.ExtractData("playerCredits", BUID, &token)
	if err != nil {
		log.Println(err)
		return nil
	}
	var twovtwo, clash, onevone IProfile_
	err = db.ExtractSingleDataFromJSON("clashProfile", "profile", &BUID, &clash)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = db.ExtractSingleDataFromJSON("oneVoneProfile", "profile", &BUID, &onevone)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = db.ExtractSingleDataFromJSON("twoVtwoProfile", "profile", &BUID, &twovtwo)
	if err != nil {
		log.Println(err)
		return nil
	}
	img := &IIMG{}
	err = db.ExtractData("img", &BUID, img)
	if err != nil {
		log.Println(err)
		return nil
	}
	var freeze, nexus int32 = 0, 0
	q := `select json_extract(powerUpgradesProgress,'$.powerProfile.freeze.level') from _nygpatch_ where id=?`
	err = db.Conn.QueryRow(q, &BUID).Scan(&freeze)
	if err != nil {
		log.Println(err)
		return nil
	}
	q = `select json_extract(powerUpgradesProgress,'$.powerProfile.nexus.level') from _nygpatch_ where id=?`
	err = db.Conn.QueryRow(q, &BUID).Scan(&nexus)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("nexus: ", nexus, "freeze: ", freeze)

	token.ClashProfile.GamesPlayed = int32(clash.Gamesplayed)
	token.ClashProfile.Rating = int32(clash.Rating)
	token.ClashProfile.Points = int32(clash.Points)
	token.ClashProfile.Tier = clash.Tier

	token.OneVOneProfile.GamesPlayed = int32(onevone.Gamesplayed)
	token.OneVOneProfile.Rating = int32(onevone.Rating)
	token.OneVOneProfile.Points = int32(onevone.Points)
	token.OneVOneProfile.Tier = onevone.Tier

	token.TwoVTwoProfile.GamesPlayed = int32(twovtwo.Gamesplayed)
	token.TwoVTwoProfile.Rating = int32(twovtwo.Rating)
	token.TwoVTwoProfile.Points = int32(twovtwo.Points)
	token.TwoVTwoProfile.Tier = twovtwo.Tier
	token.ImageURL = img.ImgURL
	token.Freeze = freeze
	token.Nexus = nexus

	return &token
}
