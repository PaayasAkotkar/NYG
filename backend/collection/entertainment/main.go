// Package entertainment implements the list of updated items
// all rights reserved, copyright 2025
package entertainment

import update "app/updateE"

type Parcel struct {
	Pack any
}

func UpdateEntertainment() Parcel {
	var token Parcel
	token.Pack = update.Update().EntertainmentBook
	return token
}

// UpdateMusic fully fletched cricket build
func UpdateMusic() []update.EntertainmentParcel {
	return update.BuildMusic().Pack
}

func UpdateMusicEvents() []string {
	return update.MusicEvents()
}

func UpdateMusicCategory() string {
	return update.MusicCategory()
}

func UpdateMusicBook() string {
	return update.MusicBook()
}

// UpdateMusicSheet all the associated key with their items
func UpdateMusicSheet() map[string][]string {
	return update.MusicSheet()
}

// UpdateMusicLists returns the updated events information
// for icc world cup 2011 -> teams-name->squad
func UpdateMusicLists(keyname string) []string {
	save := map[string][]string{}
	token := update.BuildMusic().Pack
	for _, i := range token {
		for keys := range i.MappedItems {
			save[i.Event] = append(save[i.Event], keys)
		}
	}
	return save[keyname]
}

// UpdateMusicValidation returns the updated squad as per the event and teamname
func UpdateMusicValidation(Event string, List string) map[string]map[string][]string {
	save := map[string]map[string][]string{}
	token := update.BuildMusic().Pack
	for _, search := range token {
		if search.Event == Event {
			for list, items := range search.MappedItems {
				if list == List {
					save[Event] = map[string][]string{
						List: items,
					}
				}

			}
		}
	}

	return save
}

// UpdateMovies fully fletched cricket build
func UpdateMovies() []update.EntertainmentParcel {
	return update.BuildMovies().Pack
}

func UpdateMoviesEvents() []string {
	return update.MoviesEvents()
}

func UpdateMoviesCategory() string {
	return update.MoviesCategory()
}

func UpdateMoviesBook() string {
	return update.MoviesBook()
}

// UpdateMoviesSheet all the associated key with their items
func UpdateMoviesSheet() map[string][]string {
	return update.MoviesSheet()
}

// UpdateMoviesLists returns the updated events information
// for icc world cup 2011 -> teams-name->squad
func UpdateMoviesLists(keyname string) []string {
	save := map[string][]string{}
	token := update.BuildMovies().Pack
	for _, i := range token {
		for keys := range i.MappedItems {
			save[i.Event] = append(save[i.Event], keys)
		}
	}
	return save[keyname]
}

// UpdateMoviesValidation returns the updated squad as per the event and teamname
func UpdateMoviesValidation(Event string, List string) map[string]map[string][]string {
	save := map[string]map[string][]string{}
	token := update.BuildMovies().Pack
	for _, search := range token {
		if search.Event == Event {
			for list, items := range search.MappedItems {
				if list == List {
					save[Event] = map[string][]string{
						List: items,
					}
				}

			}
		}
	}

	return save
}

// UpdateShows fully fletched cricket build
func UpdateShows() []update.EntertainmentParcel {
	return update.BuildShows().Pack
}

func UpdateShowsEvents() []string {
	return update.ShowsEvents()
}

func UpdateShowsCategory() string {
	return update.ShowsCategory()
}

func UpdateShowsBook() string {
	return update.ShowsBook()
}

// UpdateShowsSheet all the associated key with their items
func UpdateShowsSheet() map[string][]string {
	return update.ShowsSheet()
}

// UpdateShowsLists returns the updated events information
// for icc world cup 2011 -> teams-name->squad
func UpdateShowsLists(keyname string) []string {
	save := map[string][]string{}
	token := update.BuildShows().Pack
	for _, i := range token {
		for keys := range i.MappedItems {
			save[i.Event] = append(save[i.Event], keys)
		}
	}
	return save[keyname]
}

// UpdateShowsValidation returns the updated squad as per the event and teamname
func UpdateShowsValidation(Event string, List string) map[string]map[string][]string {
	save := map[string]map[string][]string{}
	token := update.BuildShows().Pack
	for _, search := range token {
		if search.Event == Event {
			for list, items := range search.MappedItems {
				if list == List {
					save[Event] = map[string][]string{
						List: items,
					}
				}

			}
		}
	}

	return save
}
