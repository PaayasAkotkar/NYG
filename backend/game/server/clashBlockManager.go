package server

import "math/rand/v2"

// ClashBlockManager returns the block and unblock id respectively
func ClashBlockManager(redID string, blueID string, redCurrentChances int, blueCurrentChances int) (string, string) {

	red := redCurrentChances > blueCurrentChances
	blue := blueCurrentChances > redCurrentChances
	box := []string{redID, blueID}
	if red {
		return blueID, redID
	} else if blue {
		return blueID, redID
	} else {
		rand.Shuffle(len(box), func(i, j int) {
			box[i], box[j] = box[j], box[i]
		})
	}

	return box[1], box[0]
}
