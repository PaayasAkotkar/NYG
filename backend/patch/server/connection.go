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
	user := "root"
	pass := "kingp12"
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv(user)
	cfg.Passwd = os.Getenv(pass)
	cfg.Net = "tcp"
	cfg.Addr = addr
	cfg.DBName = "nygpatch"
	return cfg
}

// InitClient this is done for better performance
// func InitClient() *sql.DB {

// 	once.Do(func() {
// 		addr := "127.0.0.1:3306"
// 		user := "root"
// 		pass := "kingp12"
// 		cfg := mysql.NewConfig()
// 		cfg.User = os.Getenv(user)
// 		cfg.Passwd = os.Getenv(pass)
// 		cfg.Net = "tcp"
// 		cfg.Addr = addr
// 		cfg.DBName = "nygpatch"
// 		var err error

// 		conn, err = sql.Open("mysql", cfg.FormatDSN())
// 		if err != nil {
// 			log.Fatalf("Could not open DB: %v", err)
// 		}
// 		if err = conn.Ping(); err != nil {
// 			log.Fatalf("Could not ping DB: %v", err)
// 		}

// 		// set connection pool parameters
// 		conn.SetMaxOpenConns(25)
// 		conn.SetMaxIdleConns(5)
// 	})
// 	return conn
// }
