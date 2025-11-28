package server

import (
	"encoding/json"
	"log"
)

type SpectateToken struct {
	// Teamname & PlayerName must be the name of the team which is not block
	Teamname string `json:"teamName"`

	PlayerName string `json:"playerName"`
}

const (
	specate_ = "NYGspectateInfo: "
)

func GSpectate(teamname string, playerName string) string {
	var c SpectateToken
	c.PlayerName = playerName
	c.Teamname = teamname
	_x, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		return ""
	}
	token := string(_x)
	return token
}
