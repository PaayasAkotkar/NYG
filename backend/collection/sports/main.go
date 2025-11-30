// Package sports implements the list of updated items
// all rights reserved, copyright 2025
package sports

import (
	update "app/updateS"
)

type Parcel struct {
	Pack any
}

func UpdateBook() Parcel {
	var token Parcel
	token.Pack = update.Update().SportsBook
	return token
}

// UpdateCricket fully fletched cricket build
func UpdateCricket() []update.SportsParcel {
	return update.BuildCricket().Pack
}

func UpdateCricketEvents() []string {
	return update.CricketEvents()
}

func UpdateCricketCategory() string {
	return update.CricketCategory()
}

func UpdateCricketBook() string {
	return update.CricketBook()
}

// UpdateCricketSheet all the associated key with their items
func UpdateCricketSheet() map[string][]string {
	return update.CricketSheet()
}

// UpdateCricketLists returns the updated events information
// for icc world cup 2011 -> teams-name->squad
func UpdateCricketLists(keyname string) []string {
	save := map[string][]string{}
	token := update.BuildCricket().Pack
	for _, i := range token {
		for keys := range i.MappedItems {
			save[i.Event] = append(save[i.Event], keys)
		}
	}
	return save[keyname]
}

// UpdateCricketValidation returns the updated squad as per the event and teamname
func UpdateCricketValidation(Event string, List string) map[string]map[string][]string {
	save := map[string]map[string][]string{}
	token := update.BuildCricket().Pack
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

// UpdateFootball fully fletched cricket build
func UpdateFootball() []update.SportsParcel {
	return update.BuildFootball().Pack
}

func UpdateFootballEvents() []string {
	return update.FootballEvents()
}

func UpdateFootballCategory() string {
	return update.FootballCategory()
}

func UpdateFootballBook() string {
	return update.FootballBook()
}

// UpdateFootballSheet all the associated key with their items
func UpdateFootballSheet() map[string][]string {
	return update.FootballSheet()
}

// UpdateFootballLists returns the updated events information
// for icc world cup 2011 -> teams-name->squad
func UpdateFootballLists(keyname string) []string {
	save := map[string][]string{}
	token := update.BuildFootball().Pack
	for _, i := range token {
		for keys := range i.MappedItems {
			save[i.Event] = append(save[i.Event], keys)
		}
	}
	return save[keyname]
}

// UpdateFootballValidation returns the updated squad as per the event and teamname
func UpdateFootballValidation(Event string, List string) map[string]map[string][]string {
	save := map[string]map[string][]string{}
	token := update.BuildFootball().Pack
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

// UpdateBasketball fully fletched cricket build
func UpdateBasketball() []update.SportsParcel {
	return update.BuildBasketball().Pack
}

func UpdateBasketballEvents() []string {
	return update.BasketballEvents()
}

func UpdateBasketballCategory() string {
	return update.BasketballCategory()
}

func UpdateBasketballBook() string {
	return update.BasketballBook()
}

// UpdateBasketballSheet all the associated key with their items
func UpdateBasketballSheet() map[string][]string {
	return update.BasketballSheet()
}

// UpdateBasketballLists returns the updated events information
// for icc world cup 2011 -> teams-name->squad
func UpdateBasketballLists(keyname string) []string {
	save := map[string][]string{}
	token := update.BuildBasketball().Pack
	for _, i := range token {
		for keys := range i.MappedItems {
			save[i.Event] = append(save[i.Event], keys)
		}
	}
	return save[keyname]
}

// UpdateBasketballValidation returns the updated squad as per the event and teamname
func UpdateBasketballValidation(Event string, List string) map[string]map[string][]string {
	save := map[string]map[string][]string{}
	token := update.BuildBasketball().Pack
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
