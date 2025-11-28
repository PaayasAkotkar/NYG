package server

import (
	"app/server/graph/model"
	"context"
	"fmt"
	"log"
	"time"
)

func (*QServices) UpdatedBooks(ctx context.Context) (*model.IBooks, error) {
	log.Println("in books")
	var books = InitBooks()
	var b model.IBooks
	b.Entertainment = books.Entertainment
	b.Sports = books.Sports
	return &b, nil
}

func (*QServices) Login(ctx context.Context, id string) (*model.LoginReply, error) {
	log.Println("in login")
	body := ViewBody(id)
	var rply model.LoginReply
	nowLogged := time.Now()
	rply.Msg = "login : " + nowLogged.String()
	rply.CurrentDeck = &body.Deck
	rply.Name = body.Name
	rply.Nickname = body.Nickname
	rply.UpdatedDisplay = &body.Profile
	rply.UpdatedEvents = &body.DEvents
	rply.UpdatedPlayerCredits = &body.Credits
	rply.UpdatedPowerupUpgrades = &body.Upgrades
	rply.UpdatedStore = &body.Store
	return &rply, nil
}

func (*Services) Name(ctx context.Context, id string) (<-chan string, error) {
	log.Println("in name")

	ch := make(chan string, 1)

	go func() {
		defer close(ch)

		for {
			fmt.Println("Tick")
			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return

			case token := <-pubsub.pubUpdatedName:
				ch <- token

			}
		}
	}()

	return ch, nil

}

func (r *Services) Nickname(ctx context.Context, id string) (<-chan string, error) {
	log.Println("in nickname")

	ch := make(chan string, 1)

	go func() {
		defer close(ch)

		fmt.Println("Tick")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return

			case token := <-pubsub.pubUpdatedNickname:
				ch <- token

			}
		}
	}()

	return ch, nil

}

// UpdatedStore is the resolver for the updatedStore field.
func (r *Services) UpdatedStore(ctx context.Context, id string) (<-chan *model.Store, error) {
	log.Println("in store")

	ch := make(chan *model.Store, 1)

	go func() {
		defer close(ch)

		for {
			fmt.Println("Tick")
			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return

			case token := <-pubsub.pubUpdatedStore:
				ch <- token

			}
		}

	}()

	return ch, nil

}

// UpdatedPlayerCredits is the resolver for the updatedPlayerCredits field.
func (r *Services) UpdatedPlayerCredits(ctx context.Context, id string) (<-chan *model.ICredits, error) {
	log.Println("in playerCredits")

	ch := make(chan *model.ICredits, 1)

	go func() {
		defer close(ch)

		for {
			fmt.Println("Tick")
			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return

			case token := <-pubsub.pubUpdatedPlayerCredits:
				ch <- token

			}
		}

	}()

	return ch, nil

}

// UpdatedEvents is the resolver for the updatedEvents field.
func (r *Services) UpdatedEvents(ctx context.Context, id string) (<-chan *model.DailyEvents, error) {
	log.Println("in events")
	ch := make(chan *model.DailyEvents, 1)

	go func() {
		defer close(ch)

		for {
			fmt.Println("Tick")
			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return

			case token := <-pubsub.pubUpdatedEvents:
				ch <- token
			}
		}
	}()

	return ch, nil

}

// UpdatedDisplay is the resolver for the updatedDisplay field.
func (r *Services) UpdatedDisplay(ctx context.Context, id string) (<-chan *model.Display, error) {
	log.Println("in display")

	ch := make(chan *model.Display, 1)

	go func() {
		defer close(ch)

		for {
			fmt.Println("Tick")
			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return

			case token := <-pubsub.pubUpdatedProfile:

				ch <- token
			}
		}

	}()

	return ch, nil

}

// UpdatedPowerupUpgrades is the resolver for the updatedPowerupUpgrades field.
func (r *Services) UpdatedPowerupUpgrades(ctx context.Context, id string) (<-chan *model.PowerProgress, error) {
	log.Println("in PowerUpgrades")

	ch := make(chan *model.PowerProgress, 1)

	go func() {
		defer close(ch)
		for {

			fmt.Println("Tick")

			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return

			case token := <-pubsub.pubsubUpdatedPowerUpgrades:
				ch <- token

			}
		}
	}()

	return ch, nil

}

// CurrentDeck is the resolver for the currentDeck field.
func (r *Services) CurrentDeck(ctx context.Context, id string) (<-chan *model.CurrentDeck, error) {
	log.Println("in deck")

	ch := make(chan *model.CurrentDeck, 1)

	go func() {
		defer close(ch)

		for {
			fmt.Println("Tick")
			select {
			case <-ctx.Done():
				fmt.Println("Subscription Closed")
				return

			case token := <-pubsub.pubsubUpdatedDeck:
				ch <- token

			}
		}
	}()

	return ch, nil

}
