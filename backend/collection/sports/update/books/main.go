// Package books implements the merge of total number of item in different keys
// all rights reserved, copyright 2025
package books

type BookPattern struct {
	SportsBook []string `json:"SportsBook"`
}

func Book() BookPattern {
	var token BookPattern
	list := []string{"BASEBALL", "BASKETBALL", "CRICKET", "FOOTBALL", "HOCKEY", "RUGBY"}
	token.SportsBook = list
	return token
}
