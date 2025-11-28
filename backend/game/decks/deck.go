package decks

type Powers struct {
	Covert bool `json:"covert"`
	Nexus  bool `json:"nexus"`
	Freeze bool `json:"freeze"`
	Rewind bool `json:"rewind"`
	Draw   bool `json:"draw"`
	Tag    bool `json:"tag"`
	Bet    bool `json:"bet"`
}
