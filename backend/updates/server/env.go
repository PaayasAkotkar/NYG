package server

import (
	"flag"

	"github.com/go-sql-driver/mysql"
)

var (
	port = flag.Int("port", 6070, "server port")
)

const (
	USER     = "root"
	PASSWORD = ""
)

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
