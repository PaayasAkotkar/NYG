package server

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

const (
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

func Env() *mysql.Config {
	addr := "127.0.0.1:3306"
	user := USER
	pass := PASSWORD
	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = pass
	cfg.Net = "tcp"
	cfg.Addr = addr
	cfg.DBName = "nygpatch"
	return cfg
}
