package server

import (
	"encoding/json"

	"github.com/tmaxmax/go-sse"
)

func RunHub(h *Hub) {
	for token := range h.Broadcast {
		SSEmessenger := &sse.Message{}
		convToken, _ := json.Marshal(&token.Token)
		sendToken := string(convToken)
		SSEmessenger.AppendData(sendToken)
		token.Conn.Publish(SSEmessenger)
	}
}
