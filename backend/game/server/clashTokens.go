package server

import (
	"errors"
	"nyg/profiles"
)

type ClashGameInfo struct {
	SetBook       string
	SetDictionary string
}

type IClash struct {
	storeLobbyKeys []string
	storeRoomKeys  []string
	storeClient    map[string][]string // lobby-key and client-ids
}

type ClashMatch struct {
	TeamRed  []string
	TeamBlue []string
}

type IPowers struct {
	Covert bool `json:"covert"`
	Nexus  bool `json:"nexus"`
	Freeze bool `json:"freeze"`
	Rewind bool `json:"rewind"`
	Draw   bool `json:"draw"`
	Tag    bool `json:"tag"`
	Bet    bool `json:"bet"`
}

// ClashFixtures implements the request monitor of the opponent
// if the req is sent by a than you look for b using this
type ClashFixtures struct {
	BroadcastID string // makes the work easy to broadcast to constant game ; because we only have to broadcast to the final boss
	// one time save
	MyID          string
	MyGameProfile GameProfile
	Clash         bool // if the fixtures is of clash
	// one time save
	IDs  []string // all ids of the players of both team red and team blue
	Book string
	// one time common save
	SetDictionary string
	MyCredits     profiles.Credit
	WholeGuess    []string
	// single save opponent details
	Against     string          // id of the player which is against
	MyTeam      string          // team name of the player which is against
	OppoPowerUp map[string]bool // power-up set of the opponent player .. note: not my power-up
	// challenges done by the opponent
	// how to store->
	// if the player request and the opponent has not completed yet
	// than store store[opp_id] meaning send for the update token for the opponent id
	// because the fixtures is all of that about
	OppoTossCalled   string
	OppoSetChallenge string
	OppoSetBet       string
	OppoBets         []string
	BetIDs           []string
	// single save and common save for fixtures meaning if avb than update will be done only for a and in b too
	OppoTossDone       bool
	OppoDictionaryDone bool
	OppoChallengeDone  bool
	OppoBetDone        bool
	CurrentRound       int
	MyPowersBin        []string
	// single save player stats save
	// current changes of the player
	MyCurrentChances int
	// if the oneFire>0-> than current onFire+chanceTodeduce
	ImOnFire int      // note:not opponent onFire
	MyGuess  []string // all the guess and lists validation made or seen by the player

	// globa save for all global ids
	FinalBossID    string // note: it can be any one from the match
	FinalBossFound bool   // so that we can do lock 1v1 fixture to face off the last boss
	LastDance      bool   // so that we can do lock 1v1 fixture final boss vs face off
	// note: NextTeamWinner and NextTeamLosser is important for the context of mapping via currentChances
	// opponent winner id
	NextTeamWinner string
	NextTeamname   string
	// opponent losser id
	NextTeamLosser string
	// if count==len(matches) than proceed
	// meaning there are 2 matches than count will be 4
	Count                int
	EliminatedPlayersIDs []string
	MyPoints             int
	MyPowerBin           []string
	MyPenalties          map[int]int // round->.. if ==2 than the game over
}

// IMasterSave implements the recommended tokens to save in all the ids of the respected room
type IMasterSave struct {
	Clash          bool
	FinalBossID    string // note: it can be any one from the match
	FinalBossFound bool   // so that we can do lock 1v1 fixture to face off the last boss
	LastDance      bool   // so that we can do lock 1v1 fixture final boss vs face off
	// note: NextTeamWinner and NextTeamLosser is important for the context of mapping via currentChances
	// opponent winner id
	WinnerID       string // for clash it will be nextTeamWinnerID
	WinnerTeamName string // for clash it will be nextTeamName
	// opponent losser id
	LosserID           string // for clash it will be nextTeamLosserID
	Count              int
	Book               string
	BetID              string
	EliminatedPlayerID string
	SetDictionary      string
	RedTeamScore       int
	BlueTeamScore      int
	PrevList           []string
	Bets               []string
}

type PlayerElimatedMessage struct {
	GameBegin bool `json:"gameBegin"`
	StartGame bool `json:"startGame"`
	Clash     bool `json:"clash"`
	Session   bool `json:"session"`
}

type IOneTimeSave struct {
	SetDictionary string
	TeamName      string
	Clash         bool
}
type IPowerReset struct {
	Key   string
	Clash bool
}

// UpdateFixtures no need to update current chances and on fire becuase it is updated via single stats
type UpdateFixtures struct {
	Against, MyTeam, FinalBossID, WholeGuess string
	EliminatedPlayersIDs                     []string
	Clash                                    bool
	Block, Lock                              bool
	MyPartnerID                              string
	UpdateStats                              map[int]string
	CurrentRound                             int
}

type ISingleSave struct {
	PowerUp        map[string]bool
	MyTeam         string
	TossCalled     string
	SetChallenge   string
	SetBet         string
	TossDone       bool
	DictionaryDone bool
	ChallengeDone  bool
	BetDone        bool
	Clash          bool
	Bets           []string
}

type ICommonSave struct {
	TossDone       bool
	DictionaryDone bool
	ChallengeDone  bool
	BetDone        bool
	TeamName       string // to save for the team
	Bets           []string
	Clash          bool
}

type ISingleStatsSave struct {
	OnFire         int
	CurrentChances int
	Guess          string // done guess will be append later
	Clash          bool
	Points         int // total points earned via locify
	PowersBin      string
	Penalty        map[int]int
}

type IPlayerProfile struct {
	PowerUpLevel map[string]int
	Rating       int
	Clash        bool
}
type ClashRoom struct {
	RoomName string
	Book     string
	Arena    string
}

var (
	// id->roomname->..
	_resetPower = make(chan map[string]map[string]IPowerReset) // resets the power of given id

	updateProfile = make(chan map[string]map[string]UpdateFixtures)

	// id->roomname->returns the key
	clashresetPowerUp  = make(chan map[string]map[string]IResetPower)
	locifyresetPowerUp = make(chan map[string]map[string]IResetPower)

	// id->roomname->..
	saveStats  = make(chan map[string]map[string]ISingleStatsSave) // saves in opponent id
	saveSingle = make(chan map[string]map[string]ISingleSave)      // saves in the profile of the requested id

	// roomname->..
	saveCommon  = make(chan map[string]ICommonSave)
	saveOneTime = make(chan map[string]IOneTimeSave)
	saveGlobal  = make(chan map[string]IMasterSave) // saves in all the ids of the global

	// roomname->id->..
	storeClashNicknames = make(chan map[string]map[string]string)
	getClashNicknames   = make(map[string]map[string]string)

	// roomname->returns id
	// storeFinalBoss = make(chan map[string]string)
	// getFinalBoss   = make(map[string]string)

	// roomname->
	clashResetCount = make(chan map[string]bool)

	// id->roomname->...
	createProfile = make(chan map[string]map[string]ClashFixtures)
	// updateProfileToken = make(chan map[string]map[string]Tokens)
	getClashProfile = make(map[string]map[string]ClashFixtures)

	// room-name->...
	// getClashSessionDone = make(map[string]int16)

	storeClashTokens = make(chan IClash)
	getClashTokens   IClash
)

func AppendLimit(max int, slice []string, token string) ([]string, error) {
	if len(slice) >= max {
		return nil, errors.New("max limit reached")
	}
	return append(slice, token), nil
}

var (
	Clash  = "Clash:true"
	_Clash = "Clash:false"

	Chances           = "Chances:10"
	Mode              = "Mode:"
	FinalBoss         = "FinalBoss: true"
	_ClashWaitMessgae = "ClashResultWait: Hang in there!!!! we'll tally everything up once everyone's through"
)

const (
	_StringSentinel_ = "not request for this so dont worry" // to compare if the string is not sent empty
	_IntSentinel     = -1
)
