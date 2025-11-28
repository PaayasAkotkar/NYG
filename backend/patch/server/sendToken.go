package server

import (
	"encoding/json"
	"mime/multipart"
)

type IProfile struct {
	ID *string `json:"id"`

	TwoVTwoProfile *IProfile_ `json:"twoVtwoProfile"`
	OneVOneProfile *IProfile_ `json:"oneVoneProfile"`
	ClashProfile   *IProfile_ `json:"clashProfile"`
}
type IProfile_ struct {
	// Record      int    `json:"record"`
	Gamesplayed *int    `json:"gamesPlayed"`
	Tier        *string `json:"tier"`
	Rating      *int    `json:"rating"`
	Points      *int    `json:"points"`
}

type IIMG struct {
	// ID *string `json:"id"`
	// ImageURL  *string `json:"imageURL"`
	// ImageName *string `json:"imageName"`
	// ImageType *string `json:"imageType"`
	File *multipart.FileHeader `form:"img" binding:"required"`
	ID   *string               `form:"id" binding:"required"`
}

type IDeck struct {
	ID        *string `json:"id"`
	Frist     *string `json:"_first"`
	Second    *string `json:"_second"`
	IsDefault bool    `json:"isDefault"`
}

// ICreate for post
type ICreate struct {
	ID       *string `json:"id" binding:"required"`
	Nickname *string `json:"nickname"`
	Name     *string `json:"name"`
}

type IPowerUp struct {
	ID     *string `json:"id"`
	Covert *int    `json:"covert"` // int must be current level
	Nexus  *int    `json:"nexus"`
	Freeze *int    `json:"freeze"`
	Rewind *int    `json:"rewind"`
	Draw   *int    `json:"draw"`
	Tag    *int    `json:"tag"`
	Bet    *int    `json:"bet"`
}
type UpgradeProgress struct {
	Powers     map[string]PowerProfile `json:"powerProfile"` // power-name -> ...
	PowersKeys []string                `json:"powers"`       // keys to use to unlock the map
}

func (up *UpgradeProgress) Value() (string, error) {
	x, err := json.Marshal(up)
	return string(x), err
}

type PowerProfile struct {
	Power          string `json:"power"`   // name of the power
	DonatedSpur    int    `json:"spur"`    // donate spur
	CurrentCounter int    `json:"counter"` // current counter
	Level          int    `json:"level"`   // current level
	Allow          bool   `json:"allow"`   // to allow upgrade
	Display        string `json:"display"` // msg
}
type IName struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}
type INickname struct {
	ID       *string `json:"id"`
	Nickname *string `json:"nickname"`
}

type IGen struct {
	ID                   *string
	Name                 *string
	Nickname             *string
	Deck                 *IDeck
	Ugrades              *IPowerUp
	Profile              *IProfile
	Credentials          *Credit
	PowerUpgradeProgress *UpgradeProgress
}

func (d IProfile) Value() (string, error) {
	e, err := json.Marshal(d)
	return string(e), err
}
func (d IPowerUp) Value() (string, error) {
	e, err := json.Marshal(d)
	return string(e), err
}

func (d IDeck) Value() (string, error) {
	b, err := json.Marshal(d)
	return string(b), err
}

func (i IIMG) Value() (string, error) {
	b, err := json.Marshal(i)
	return string(b), err
}
