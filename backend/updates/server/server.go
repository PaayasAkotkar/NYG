package server

import (
	"app/nygprotoc"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Server() {
	log.Println("NYG UPDATE SERVER RUNNING ⚙️")

	_p := fmt.Sprintf(":%d", *port)
	_net := "tcp"
	lsn, err := net.Listen(_net, _p)
	if err != nil {
		log.Println(err)
	}
	defer lsn.Close()

	NYGserver := grpc.NewServer()

	nygprotoc.RegisterNYGGameCredentialsServer(NYGserver, &Services{})

	nygprotoc.RegisterNYGRatingServer(NYGserver, &Services{})

	if err := NYGserver.Serve(lsn); err != nil {
		panic(err)
	}
}
