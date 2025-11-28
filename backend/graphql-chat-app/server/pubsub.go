// Package server
// the pubsub system is robust build here
// it can be implmented in the direct request too nothing complex
// methods are implmented here for cleaner version
package server

import (
	"app/server/graph/model"
	"sync"
)

type NYGPubSub struct {
	mu          sync.Mutex
	subscribers map[string]map[chan *model.LatestMessage]struct{}
}

func InitPubSub() *NYGPubSub {
	return &NYGPubSub{
		subscribers: make(map[string]map[chan *model.LatestMessage]struct{}),
	}
}

func (p *NYGPubSub) Subscribe(room string) chan *model.LatestMessage {
	ch := make(chan *model.LatestMessage, 1)
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.subscribers[room]; !ok {
		p.subscribers[room] = make(map[chan *model.LatestMessage]struct{})

	}
	p.subscribers[room][ch] = struct{}{}
	return ch
}

func (p *NYGPubSub) Unsubscribe(room string, subscriber chan *model.LatestMessage) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if subs, ok := p.subscribers[room]; ok {
		if _, exists := p.subscribers[room][subscriber]; exists {
			delete(subs, subscriber)
			close(subscriber)
		}
		if (len(subs)) == 0 {
			delete(p.subscribers, room)
		}
	}
}

func (p *NYGPubSub) Publish(room string, latest *model.LatestMessage) {
	p.mu.Lock()
	defer p.mu.Unlock()
	subs := p.subscribers[room]
	for ch := range subs {
		select {
		case ch <- latest:
		default:
		}
	}
}

var (
	pubsub = InitPubSub()
)
