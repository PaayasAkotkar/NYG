package server

import (
	"github.com/google/uuid"
)

type Lobby struct {
	JoinRoom    bool   `json:"joinRoom"`    // if true
	To          string `json:"to"`          // then join to this
	CreateRoom  bool   `json:"createRoom"`  // if true
	RoomCreated bool   `json:"roomCreated"` // if true than they can modify the room settings

	/*room data*/
	Category string `json:"category"` // set of category  [sports or entertainment]
	Field    string `json:"field"`    // set of field for [international, national, domestic]
	Book     string `json:"book"`     // name of the book set for the room [for example cricket, basketball, tennis...]
	/**end of room data*/
	Reverse      bool   `json:"reverse"`      // if to set the challenge dictionary as A1-B1 or A1-A1 [meaning: A1-B1 the one who sets the challenge dictionary the other one can first set the challenge than other, A1-A1 the one who sets the challenge dictionary than he  can first set the cahllenge than other]
	RoomName     string `json:"roomName"`     // then create this room
	RoomCapacity int    `json:"roomCapacity"` // for 1v1 is 2 and for 2v2 is 4
	Friend       bool   `json:"friend"`       // if the client is joint via code
	Code         string `json:"code"`         // if friend true than join via code
	// NickName       string `json:"nickname"`     // to display to all connected user
	GameTime       int    `json:"gameTime"`
	DecisionTime   int    `json:"decisionTime"`
	Set            bool   `json:"set"`            // required else the nickname wont be saved
	ID             string `json:"id"`             // generated or saved id
	SetToss        bool   `json:"setToss"`        // if to involve toss session
	Starter        bool   `json:"starter"`        // if to involve starter system[winners first or lossers first]
	PrivateRoom    bool   `json:"privateRoom"`    // if the room request is for the private room
	ChangeSettings bool   `json:"changeSettings"` // to change the settings
	NexusPower     bool   `json:"nexusPower"`     // sends the word from the dictionary but in mazed order [V_R A_]
	TagPower       bool   `json:"tagPower"`       // locks the partner in-place of the chosen power
	RewindPower    bool   `json:"rewindPower"`    // reset's the clock once the output of the opponent been used
	FreezePower    bool   `json:"freezePower"`    // gives the control to start the clock under specific condition
	DrawPower      bool   `json:"drawPower"`      // sends the request to the player to tag in their partner's inplace of them
	CovertPower    bool   `json:"covertPower"`    // hides the text when the opponent writes something
	// any one can use this bet power [if the dictionary setter or non-dictionary setter]
	BetPower bool `json:"betPower"` // player get's the list of the current event and they ought to bet on any one word of the event; note: they cannot guess
	Clash    bool `json:"clash"`
}

type GameRoom struct {

	// mutal chatting: both players will vote to mutally decide for the challenge token
	// they wouldnt be able to chat rather they can pick
	// if not pick under time than the battle ground player can pick for the win
	// if even the battle ground player wouldnt be able to pick the point goes to the opponent team
	// if both players weren't able to pick
	// the warning will be given
	// even after the warning the players wont pick the room will be closed

	RedScoreCount  int `json:"redScoreCount"`  // current home score; note: it will always be less than 1 than current meaning if the score board is 2 than then score count send here will be 1
	BlueScoreCount int `json:"blueScoreCount"` // current away score
	Chances        int `json:"chances"`        // player's life
	OnFire         int `json:"onFire"`

	RoomName        string `json:"roomName"`        // name of the room
	ChallengeToken  string `json:"challengeToken"`  // token to challenge the rival
	Guess           string `json:"guessToken"`      // guess the token as per the given guessin category
	HeadTails       string `json:"headTails"`       // chosen side
	DictionaryToken string `json:"dictionaryToken"` // chosen dictionary word

	Session             bool      `json:"session"`             // on going game session
	Start               bool      `json:"start"`               // is the game has started
	TossSession         bool      `json:"tossSession"`         // for toss session
	ChallengeSet        bool      `json:"challengeSet"`        // same as Agree
	DictionarySession   bool      `json:"dictionarySession"`   // on going dictionary session for setting challenge
	ChallengeDiscussion bool      `json:"challengeDiscussion"` // is in the challenge discussion session
	PowerActivated      bool      `json:"power"`               // if the player has decided to use the power
	FinalBoss           bool      `json:"finalBoss"`           // if the final boss has been found
	LastDance           bool      `json:"lastDance"`           // final match
	Clash               bool      `json:"clash"`
	TimeUp              bool      `json:"timeUp"`           // if the player loss by time
	DTimeUp             bool      `json:"discussionTimeUp"` // if the player hasnt completed the session in under this time
	Spectate            bool      `json:"spectate"`
	View                ISpectate `json:"myView"`
	EmptyHand           bool      `json:"emptyHand"` // if the player endedup with 0 ditionaries

	/** players attribute */
	Freeze bool `json:"freeze"`
	Nexus  bool `json:"nexus"`
	Rewind bool `json:"Rewind"`
	Tag    bool `json:"tag"`
	Draw   bool `json:"draw"`
	Covert bool `json:"covert"`
	Bet    bool `json:"bet"`
	/** end of players attribute **/

	Unfreeze bool `json:"unfreeze"`

	DrawSession bool   `json:"drawSession"` // this will be key to accept and reject the draw offer
	DrawAccept  bool   `json:"drawAccept"`
	Set         int    `json:"set"`        // current set
	Round       int    `json:"round"`      // current Round
	BetSession  bool   `json:"betSession"` // this will be the key for storing the bet value
	BetOn       string `json:"betOn"`      // storing the bet value

	Time int `json:"time"`
}

// ISpectate about powers
// about powers:
// for clash -finalBoss: we can broadcast the boss everything for each round which is easy
// brodcasting everything is same as playing while it is not playing it is watching
// for locify: there are two stags:
// -afterGame and -beforeGame
// -beforeGame: we must only broadcast everything to the not locked player only of their teammate
// -afterGame: we must only broadcast everything to the not locked player of both the teams
type ISpectate struct {
	ScrollPos     float64 `json:"scrollPos"`     // current scroll pos
	OnClick       string  `json:"onClick"`       // current clicked event
	Word          string  `json:"word"`          // current word written on the guest
	ScrollHappend bool    `json:"scrollHappend"` // if event was scroll
}

type SignOut struct {
	RoomName string `json:"roomName"`
	ID       string `json:"id"`
	Logout   bool   `json:"loggout"`
}

func GenCode() string {
	return uuid.New().String()[0:4]

}
