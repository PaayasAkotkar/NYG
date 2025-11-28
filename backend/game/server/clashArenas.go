package server

// GYM: random removes any one the power hand from the players deck of each player
// new GYM: no traits
// PSHYCIC: rewind can happended at any time
// NODE: the probability of power up happening is less
// FASTLANE: [on hold] every round it sends probailty of name of the player from the database collection
var (
	Arenas = []string{"gym", "pshycic", "node", "fastlane", "mysterium"}
	board  = make(map[string]string)
)

// CreateBoard returns the name of the board
func CreateBoard(arenaWinner string) string {
	board["gym"] = "clash-gym"
	board["pyshcic"] = "clash-pyshcic"
	board["node"] = "clash-node"
	board["fastlane"] = "clash-fastlane"
	board["mysterium"] = "clash-mysterium"

	return board[arenaWinner]
}
