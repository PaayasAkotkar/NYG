package server

import (
	"log"
	"math/rand/v2"
)

func LocifyBlockManager(h *Hub, round int, teams map[string]map[string][]string, profiles map[string]map[string]LocifyFixtures, roomname, winnerTeamname, losserTeamname string, draw bool) {
	log.Println("in locify block manager")
	_set := profiles[winnerTeamname][roomname].RoomSettings

	if !_set.Starter {
		// unblock the winner
		for _, id := range saveShuffle[roomname][winnerTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, to: id, roomname: roomname, _sleep: false}
		}
		for _, id := range saveShuffle[roomname][losserTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, to: id, roomname: roomname, _sleep: false}
		}
	} else if !draw {
		// block the winner
		for _, id := range saveShuffle[roomname][losserTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Unblock, to: id, roomname: roomname, _sleep: false}
		}
		for _, id := range saveShuffle[roomname][winnerTeamname] {
			h.gameRoomBroadcast <- reqGameRoomBroadcast{token: Block, to: id, roomname: roomname, _sleep: false}
		}
	}
	if _set.SetupToss || draw {
		BoardcastSession(h, roomname,
			false, true, false, false, false)

	} else {
		BoardcastSession(h, roomname, false,
			false, false, true, false)
	}

}

// LocifyRandBlock note: it just works with the current playing id and nothing to do with idle id
// returns the block and unblock ids of the respective team with their idle ids
func LocifyRandBlock(matchUp LocifyMatch, round int, roomname string) (string, string, string, string, string, string) {
	roll := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	rand.Shuffle(len(roll), func(i int, j int) {
		roll[i], roll[j] = roll[j], roll[i]
	})
	var block, unblock string
	var bIdle, uIdle string
	var bteamname, uteamname string
	if roll[0]%2 == 0 {
		unblock = matchUp.TeamBlue.PlayingID
		uIdle = matchUp.TeamBlue.IdleID
		uteamname = _TeamBlueKey

		block = matchUp.TeamRed.PlayingID
		bIdle = matchUp.TeamRed.IdleID
		bteamname = _TeamRedKey

	} else {

		unblock = matchUp.TeamRed.PlayingID
		uIdle = matchUp.TeamRed.IdleID
		uteamname = _TeamRedKey

		block = matchUp.TeamBlue.PlayingID
		bIdle = matchUp.TeamBlue.IdleID
		bteamname = _TeamBlueKey

	}

	return block, bIdle, unblock, uIdle, bteamname, uteamname
}
