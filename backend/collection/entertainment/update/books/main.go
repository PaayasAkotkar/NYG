// Package books implements the merge of total number of item in different keys
// all rights reserved, copyright 2025
package books

type BookPattern struct {
	EntertainmentBook []string `json:"EntertainmentBook"`
}

func Book() BookPattern {
	var token BookPattern
	list := []string{"MOVIES", "MUSIC", "SHOWS", "BOOKS"}
	token.EntertainmentBook = list
	return token
}
