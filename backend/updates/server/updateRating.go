package server

import (
	"app/nygprotoc"
	"app/sqlmanager"
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type UpdateGameCredits struct {
	ID                  string   `json:"id"`
	Books               []string `json:"books"`
	Categories          []string `json:"categories"`
	Powers              []string `json:"powers"`
	Points              int32    `json:"points"`
	GamesWon            int32    `json:"gamesWon"`
	GamesPlayed         int32    `json:"gamesPlayed"`
	Rating              int32    `json:"rating"`
	Tier                string   `json:"tier"`
	RatingIncrementedBy int32    `json:"ratingIncremented"`
	RatingDecrementedBy int32    `json:"ratingDecremented"`
}

func (u UpdateGameCredits) Value() (string, error) {
	b, err := json.Marshal(u)
	return string(b), err
}

func Update(playersCredits map[string]*nygprotoc.TypeGCredentials, clash, onevone bool) error {
	m := sqlmanager.ConnectSQL{}
	cfg := Env()
	db, err := m.Init("nygpatch", "nyg_gamecredentials", cfg)

	if err != nil {
		log.Println(err)
		return err
	}

	defer db.CloseDB()

	for id, det := range playersCredits {
		var updated UpdateGameCredits
		_uuid, _ := uuid.Parse(id)
		BUID, _ := _uuid.MarshalBinary()
		err = db.ExtractData("credits", BUID, &updated)

		if err != nil {
			log.Println(err)
			return err
		}

		if exists, err := db.HasID(BUID); exists && err == nil {
			db.ChangeTable("nyg_gamecredentials")

			if det.Powers != nil {
				updated.Powers = append(updated.Powers, det.Powers...)
			}

			updated.Points += det.Points
			updated.Books = append(updated.Books, det.Book)
			updated.GamesPlayed += det.GamesPlayed
			updated.GamesWon += det.Results
			updated.Categories = append(updated.Categories, det.Category)
			updated.Rating = det.Rating
			updated.RatingIncrementedBy = det.Incremented
			updated.RatingDecrementedBy = det.Decremented

			switch true {
			case det.Rating < 1400:
				updated.Tier = "novice"
				// to reach grandmaster
			case det.Rating >= 1400 && det.Rating < 2000:
				updated.Tier = "master"
			// to reach super grandmaster
			case det.Rating >= 2000 && det.Rating < 2500:
				updated.Tier = "grandmaster"
			case det.Rating >= 2500:
				updated.Tier = "super grandmaster"
			default:
				updated.Tier = "novice"
			}

			v, err := updated.Value()

			if err != nil {
				panic(err)
			}

			db.ChangeTable("nyg_gamecredentials")
			if err := db.UpdateWholeJSONentry(BUID, "credits", v); err != nil {
				log.Println(err)
				return err
			}
			var profile IProfile_
			profile.Rating = int(updated.Rating)
			profile.Points = int(updated.Points)
			profile.Tier = updated.Tier
			profile.Gamesplayed = int(updated.GamesPlayed)
			_v, err := profile.Value()
			if err != nil {
				log.Println(err)
				return err
			}
			db.ChangeTable("_nygpatch_")
			switch true {
			case clash:
				if err := db.UpdateSingleJSONentry(BUID, "clashProfile", "profile", &_v); err != nil {
					log.Println(err)
					return err
				}
			case onevone:
				if err := db.UpdateSingleJSONentry(BUID, "oneVoneProfile", "profile", &_v); err != nil {
					log.Println(err)
					return err
				}
			case !onevone:
				if err := db.UpdateSingleJSONentry(BUID, "twoVtwoProfile", "profile", &_v); err != nil {
					log.Println(err)
					return err
				}
			}

		} else {
			return err
		}
		db.ChangeTable("nyg_gamecredentials")

	}
	return nil
}
