package server

var (
	BooksURLs = []string{"/update/sport-book", "/update/entertainment-book"}

	// broadcast dictionaries such as icc world cup 2011
	DictionariesURL = "/ocs-ps/event/:book"

	// boardcast list such as india, australia....
	ListsURLs = "/ocs-ps/items/:book/:dictionary"

	// boardcast items such as team india squad from 2011
	ItemsURLs = "/ocs-ps/validate/items/:book/:dictionary/:list"
)
