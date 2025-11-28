package server

import (
	"app/server/graph"
	"app/server/graph/model"
	"context"
	"log"
	"net/http"
)

// NYGServices implments the subscription and query services
type NYGServices struct {
	*NYGResolver
}

// NYGReplyServices implments the publish services
type NYGReplyServices struct {
	*NYGResolver
}
type NYGResolver struct{}

func (r *NYGResolver) Subscription() graph.SubscriptionResolver { return &NYGReplyServices{r} }
func (r *NYGResolver) Mutation() graph.MutationResolver         { return &NYGServices{r} }
func (r *NYGResolver) Query() graph.QueryResolver               { return &NYGServices{r} }

// Client inits the message incase of disconnection
func (*NYGServices) Client(ctx context.Context, id string) (*model.InitMessage, error) {
	log.Println("in client")
	body := ViewBody(id)

	if body != nil {
		var init model.InitMessage
		init.ID = body.latest.ID
		init.Msg = body.latest.Msg
		init.Name = body.latest.Name
		init.Roomname = body.latest.Roomname
		return &init, nil

	} else {
		return &model.InitMessage{Msg: "welcome"}, nil
	}
}

// Latest publishes the tokens
func (r *NYGReplyServices) Latest(ctx context.Context, room string) (<-chan *model.LatestMessage, error) {
	log.Println("in latest")
	ch := pubsub.Subscribe(room)
	go func() {
		<-ctx.Done()
		pubsub.Unsubscribe(room, ch)
	}()

	return ch, nil
}

// Post triggers the subscription
func (r *NYGServices) Post(ctx context.Context, client model.Put) (*model.Reply, error) {
	var latest model.LatestMessage
	latest.ID = client.ID
	latest.Msg = client.Msg
	latest.Roomname = client.Roomname
	latest.Name = client.Name
	store := map[string]model.LatestMessage{}
	store[client.Roomname] = latest
	room := client.Roomname

	Update(latest.ID, latest)

	pubsub.Publish(room, &latest)

	return &model.Reply{Status: http.StatusOK, Name: latest.Name, Msg: latest.Msg}, nil
}
