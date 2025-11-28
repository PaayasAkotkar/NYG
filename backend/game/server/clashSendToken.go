package server

const (
	clashGameTime       = 10
	clashDiscussionTime = 1
)

type Pairing struct {
	TeamName string `json:"teamname"`
	Chances  int    `json:"chances"`
}
type ClashGameReward struct {
	OnFire        int `json:"onFire"`
	CurrentChance int `json:"currentChance"`
}

type ParcelClashResult struct {
	RoomName string             `json:"roomname"`
	Pairing  map[string]Pairing `json:"pairing"`
	Round    int                `json:"round"`
	// FinalBoss bool               `json:"finalBoss"`
	// LastDance bool               `json:"lastDance"`
}

type ParcelClashInfo struct {
	RoomName  string `json:"roomname"`
	GameBegin bool   `json:"gameBegin"`
	Chances   int    `json:"chances"`
	Round     int    `json:"round"`
	Mode      string `json:"mode"`
	TeamName  string `json:"teamName"`
	// this is cruical for scoring
	Pairing        map[string]Pairing `json:"pairing"` // id-teamname
	IDs            []string           `json:"ids"`
	List           []string           `json:"list"`
	FinalBoss      bool               `json:"finalBoss"`
	LastDance      bool               `json:"lastDance"`
	GameTime       int                `json:"gameTime"`
	DiscussionTime int                `json:"discussionTime"`
}

type ParcelPowerReset struct {
	IsBetting bool `json:"isBetting"`
	Betting   bool `json:"betting"`
	Unfreeze  bool `json:"unfreeze"`
	UnderTest bool `json:"underTest"`
}

// ParcelMatchupUpdate incase found the last game or final boss
type ParcelMatchupUpdate struct {
	FinalBoss bool `json:"finalBoss"`
	LastDance bool `json:"lastDance"`
	IMBoss    bool `json:"imBoss"`
}
