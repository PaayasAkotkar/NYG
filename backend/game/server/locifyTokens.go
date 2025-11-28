package server

import "nyg/profiles"

type LocifyRoomSettings struct {
	SetupToss, Starter,
	Reverse, TwoVTwo bool
	Powers                                       map[string]bool
	RoomOwnerID, FriendID, Field, Category, Name string
	DecisionTime, GameTime                       int
	Book, Mode, Code                             string
	Public, Private, Friend, FriendJoin          bool
	Capacity                                     int
}
type GameProfile struct {
	FreezeLevel int
	NexusLevel  int
	Rating      int
	Points      int
	Gamesplayed int
	Pic         string
}

// LocifyFixtures implements the request monitor of the opponent
// if the req is sent by a than you look for b using this
type LocifyFixtures struct {
	// one time save
	MyID string
	// one time save
	IDs           []string // all ids of the players of both team red and team blue
	Book          string
	PowerDeck     map[string]bool // power key->true means can be use
	IBlock        bool
	ILock         bool
	MyGameProfile GameProfile

	// note idleplayer is the parnter but not playing for this round
	IdlePlayer string // idle player
	MyPatnerID string // if any else _StringSentinel_

	OppoTeamname string // oppo team
	RoomAdminID  string
	CurrentRound int

	// single save opponent details
	Against            string          // id of the player which is against
	MyTeam             string          // team name of the player which is against
	OppoPowerUp        map[string]bool // power-up set of the opponent player .. note: not my power-up
	OppoTossCalled     string
	OppoSetChallenge   string
	OppoSetBet         string
	OppoBets           []string
	BetIDs             []string
	OppoTossDone       bool
	OppoDictionaryDone bool
	OppoChallengeDone  bool
	OppoBetDone        bool
	OppoPenalties      int

	// global save for all global ids
	DisconnectedID      string   // note: it can be any one from the match
	DisconnectedIDFound bool     // so that we can do lock 1v1 fixture to face off the last boss
	MyGuess             []string // all the guess made by the player will be saved for both the players
	SetDictionary       string

	RedTeamScore  int
	BlueTeamScore int
	RoomCode      string

	RoomSettings LocifyRoomSettings

	NickNamesViaID map[string]string

	PrevWinnerTeam string
	WholeGuess     []string // combined guesses of the player

	// round -> ...
	MySheet     map[int]string // per round stat
	MyPoints    int            // increment by 1 if score point
	MyPowersBin []string       // powers used so far
	MyCredits   profiles.Credit
	MyPenalties map[int]int
}

type LockPatterns struct {
	// teamname id
	TeamRed  LockPlayerInfo
	TeamBlue LockPlayerInfo
}
type LockPlayerInfo struct {
	PlayingID string
	IdleID    string
}
type LocifyMatch struct {
	// teamname id
	TeamRed  LockPlayerInfo
	TeamBlue LockPlayerInfo
}
type RoomSettingsSave struct {
	Book      string
	DTime     int
	GTime     int
	FriendID  string
	SetupToss bool
	Reverse   bool
	Starter   bool
	Powers    map[string]bool
}
type LocifyOneTimeSave struct {
	SetDictionary string
	Teamname      string
}

var (
	createLocifyProfile     = make(chan map[string]map[string]LocifyFixtures)
	getLocifyProfile        = make(map[string]map[string]LocifyFixtures)
	storeLocifyRoomSettings = make(chan map[string]LocifyRoomSettings)
	getLocifyRoomSettings   = make(map[string]LocifyRoomSettings)

	// roomname-...
	saveRoomSettings = make(chan map[string]RoomSettingsSave)

	// room-name ->..
)

const (
	SessionKey = "NYGSessionUpdate:"
	LocifyMode = "locify"
)
