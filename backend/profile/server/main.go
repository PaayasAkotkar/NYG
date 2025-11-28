// Package server fetches all the data associated with that id and send it to the client
// all rights reserved, copyright 2025
package server

import (
	"fmt"
	png "image/png"
	"os"
)

func Open(fs string) {
	fi, err := os.Open(fs)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	img, err := png.Decode(fi)
	if err != nil {
		panic(err)
	}

	fmt.Println(img.Bounds().Dx(), img.Bounds().Dy())
}

func Run() {
	SetEnv()
	Server()
}
