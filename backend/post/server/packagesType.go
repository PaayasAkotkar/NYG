package server

import (
	"database/sql/driver"
	"encoding/json"
)

type IProfile struct {
	ID string `json:"id"`

	TwoVTwoProfile IProfile_ `json:"twoVtwo"`
	OneVOneProfile IProfile_ `json:"oneVone"`
	ClashProfile   IProfile_ `json:"clashProfile"`
}

type IProfile_ struct {
	// Record      int    `json:"record"`
	Gamesplayed int    `json:"gamesPlayed"`
	Tier        string `json:"tier"`
	Rating      int    `json:"rating"`
	Points      int    `json:"points"`
}
type IIMG struct {
	ImgURL  string `json:"imgURL"`
	ImgName string `json:"imgName"`
	ID      string `json:"id"`
}

type IDeck struct {
	ID        string `json:"id"`
	Frist     string `json:"_first"`
	Second    string `json:"_second"`
	IsDefault bool   `json:"isDefault"`
}

// ICreate for post
type ICreate struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

type IPowerUp struct {
	ID     string `json:"id"`
	Covert int    `json:"covert"` // int must be current level
	Nexus  int    `json:"nexus"`
	Freeze int    `json:"freeze"`
	Rewind int    `json:"rewind"`
	Draw   int    `json:"draw"`
	Tag    int    `json:"tag"`
	Bet    int    `json:"bet"`
}

type IName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type INickname struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

type IGen struct {
	ID       string
	Name     string
	Nickname string
	Deck     IDeck
	Ugrades  IPowerUp
	Profile  IProfile
}

func (d IProfile) Value() (driver.Value, error) {
	return json.Marshal(d)
}
func (d IPowerUp) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d IDeck) Value() (string, error) {
	b, err := json.Marshal(d)
	return string(b), err
}
