package server

type Frame struct {
	Frames []string `json:"frames"`
}

// note: last update of these variables will be their final value
// make sure to take care of that
// for-safty i have created stop and start var
// so that we dont have to be worry of them changing the value
const (

	// stop on going sessions
	Gamesession         = "GameSession: false"         // on going game session
	tossSession         = "TossSession: false"         // on going toss session
	challengeDiscussion = "ChallengeDiscussion: false" // on going challenge set session

	// DictionaryDiscussion  = "DictionaryDiscussion: false" // on going challenge dictornay dicussion set
	gameBegin             = "GameBegin: true"            // on going game session
	_DictionaryDiscussion = "DictionaryDiscussion: true" // on going challenge dictornay dicussion set
	Spectate              = "NYGSpectate: "

	// for clash it will be different
	// _RedSpectate  = "NYGredSpectate: "  // helpful for text
	// _BlueSpectate = "NYGblueSpectate: " // helpful for text
	// start signals
	_startGame = "StartGame: true" // signal to start the game
	_waiting   = "Waiting: true"   // signal to wait until the challenge is set
	_toss      = "Toss: true"      // signal to start the toss
	_tossCoin  = "TossCoin: true"  // signal to toss the coin

	// stop signals
	startGame     = "StartGame: false"     // signal to start the game
	waiting       = "Waiting: false"       // signal to wait until the challenge is set
	toss          = "Toss: false"          // signal to toss done
	tossCoin      = "TossCoin: false"      // signal to stop the coin tos
	DictionarySet = "DictionarySet: false" // singal to stop broadcast the dictionary event

	RemoveChallenge  = "ChallengeRemove: "
	RemoveDictionary = "DictionaryRemove: "

	// deadblocking
	Block   = "Block: true"  // signal to block the player or team from current session
	Unblock = "Block: false" // signal to unblock the player or team to play the  session

	DictionaryURL      = "DURL: "  // to send the dictionary after the toss done to set the dictionary
	ItemsURL           = "CURL: "  // to send the items after the dictionary event to set the challenge
	SetupDictionaryURL = "DURLT: " // for the setup

	// locking
	Lock   = "Lock: true"  // signal to lock the player from current session
	Unlock = "Lock: false" // signal to unlock the player to play the current session

	DictionaryEvent = "DictionaryEvent: "
	BlueSpectate    = "NYGblueSpectate: "
	RedSpectate     = "NYGredSpectate: "

	Frames = "NYGFrames: "

	// for now only 2 players can visit the session
	// one player is from team blue and other team red
	// limit for session per people:  [1,2]
	// resetCount = "ResetCount: 1"

	// power has been in used
	// FreezePower = "PowerFreeze: true"
	// NexusPower  = "PowerNexus: true"
	// RewindPower = "PowerRewind: true"
	// CovertPower = "CovertPower: true"

	// power use done
	CanUsePower  = "PowerActivated: true"
	_CanUsePower = "PowerActivated: false"
	View         = "NYGView: "
)
