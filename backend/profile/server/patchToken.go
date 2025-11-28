package server

import (
	"app/server/graph/model"
)

type IProfile struct {
	ID string `json:"id"`

	TwoVTwoProfile INYGRecord `json:"twoVtwoProfile"`
	OneVOneProfile INYGRecord `json:"oneVoneProfile"`
	ClashProfile   INYGRecord `json:"clashProfile"`
}
type INYGRecord struct {
	// Record      int    `json:"record"`
	Gamesplayed int    `json:"gamesPlayed" graphql:"gamesPlayed"`
	Tier        string `json:"tier" graphql:"tier"`
	Rating      int    `json:"rating" graphql:"rating"`
	Points      int    `json:"points" graphql:"points"`
}
type IPowerUp struct {
	ID     string `json:"id"`
	Covert int    `json:"covert"`
	Nexus  int    `json:"nexus"`
	Freeze int    `json:"freeze"`
	Rewind int    `json:"rewind"`
	Draw   int    `json:"draw"`
	Tag    int    `json:"tag"`
	Bet    int    `json:"bet"`
}
type IDeck struct {
	ID        string `json:"id"`
	Frist     string `json:"_first"`
	Second    string `json:"_second"`
	IsDefault bool   `json:"isDefault"`
}

// Body implments all the data stored in the database
type Body struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	Nickname       string              `json:"nickname"`
	TwoVTwoProfile model.IProfile      `json:"twoVtwoProfile"`
	OneVOneProfile model.IProfile      `json:"oneVoneProfile"`
	ClashProfile   model.IProfile      `json:"clashProfile"`
	NYGProfile     model.IProfile      `json:"totalRecord"`
	ImageData      *model.IImage        `json:"imageData"`
	Upgrades       model.PowerProgress `json:"upgrades"`
	Deck           model.CurrentDeck   `json:"deck"`
	Locify         model.DailyEvents   `json:"locify"`
	Clash          model.DailyEvents   `json:"clash"`
	DEvents        model.DailyEvents   `json:"events"`
	Credits        model.ICredits      `json:"playerCredits"`
	Store          model.Store         `json:"store"`
	Profile        model.Display       `json:"display"`
}

type Credit struct {
	Coin  string `json:"coins"`
	Spur  string `json:"spurs"`
	Theme string `json:"theme"`
}
