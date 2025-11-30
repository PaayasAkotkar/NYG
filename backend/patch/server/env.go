package server

import (
	"log"
	"os"
)

var (
	port      = "5556"
	allowOrgs = []string{
		"http://localhost:4200",
		"http://172.26.128.1:4200",
		"http://192.168.1.7:4200",
	}

	USER     = "root"
	PASSWORD = ""
)

func SetEnv() {

	err := os.Setenv(USER, USER)
	if err != nil {
		log.Fatalf("Error setting MYSQL_USER: %v", err)
	}
	err = os.Setenv(PASSWORD, PASSWORD)
	if err != nil {
		log.Fatalf("Error setting MYSQL_PASS: %v", err)
	}
}
