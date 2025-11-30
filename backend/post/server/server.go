package server

import (
	nygpostprotoc "app/nygprotoc"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Server() {
	log.Println("NYG POST SERVER RUNNING ⚙️")
	_p := fmt.Sprintf(":%d", *port)
	_net := "tcp"
	lsn, err := net.Listen(_net, _p)

	if err != nil {
		log.Println(err)
	}

	defer lsn.Close()

	NYGserver := grpc.NewServer()
	nygpostprotoc.RegisterNYGPostProfileServer(NYGserver, &Services{})

	if err := NYGserver.Serve(lsn); err != nil {
		log.Println(err)
		return
	}

}
