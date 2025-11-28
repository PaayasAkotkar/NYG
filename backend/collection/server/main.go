// Package server runs the server
// all rights reserved, copyright 2025
package server

type ParcelSheet struct {
	Pack map[string][]string `json:"Pack"`
}
type ParcelValidationSheet struct {
	Pack map[string]map[string][]string `json:"Pack"`
}

type ParcelEvents struct {
	Pack []string `json:"Pack"`
}

func Start() {
	Server()
}
