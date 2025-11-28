package server

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var (
	port     = flag.Int("port", 3434, "server port")
	USER     = ""
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
	user := "root"
	pass := PASSWORD
	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = pass
	cfg.Net = "tcp"
	cfg.Addr = addr
	cfg.DBName = "nygpatch"
	return cfg
}

func SetCookie(ctx *gin.Context) {
	words := []string{"a", "b"}
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	_crp := sha256.Sum256([]byte(words[0]))
	id := hex.EncodeToString(_crp[:])[:4]
	limit := int((365 * 100) * 24 * time.Hour / time.Second)
	path := "/" // so that to allow angular client to use one single cookie
	domain := "localhost"
	name := "genID"
	ctx.SetCookie(name, id, limit, path, domain, true, false)
}
