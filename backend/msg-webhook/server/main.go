// Package server a webhook to fill the indox
// all rights reserved, copyright 2025
package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// IMessage author must be the one sending the msg to the client
type IMessage struct {
	Author string `json:"author"`
	Msg    string `json:"msg"`
	View   bool   `json:"view"`
}

func RunServer() {
	log.Println("webhook starts")

	port := "1500"
	http.HandleFunc("/nyg-welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Println("welcome")
		if r.Method == "POST" {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		// this is currently will broadcast to whole
		audience := IMessage{}
		audience.Author = "PAAYAS"
		audience.Msg = "Welcome to NYG!!!!!🤗"
		audience.View = false
		token, err := json.Marshal(audience)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(token)
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println(err)
		return
	}
}
