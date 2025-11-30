package server

// var (
// 	once sync.Once
// 	conn *sql.DB
// )
import (
	"os"

	"github.com/go-sql-driver/mysql"
)

func Env() *mysql.Config {
	addr := "127.0.0.1:3306"
	user := USER
	pass := PASSWORD
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv(user)
	cfg.Passwd = os.Getenv(pass)
	cfg.Net = "tcp"
	cfg.Addr = addr
	cfg.DBName = "nygpatch"
	return cfg
}
