package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tmaxmax/go-sse"
)

// Server note: books has been removed from here
func Server() {
	app := gin.New()
	h := NewHub()
	conn := &sse.Server{}
	fmt.Println("ROOT SERVER RUNNING ⚙️")

	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Next()
	})

	// we can do this with for-loop too
	app.GET(DictionariesURL, func(ctx *gin.Context) {
		fmt.Println("sending to client")
		conn.ServeHTTP(ctx.Writer, ctx.Request)
	})

	app.POST(DictionariesURL, func(ctx *gin.Context) {
		store := map[string]map[string]bool{}
		h.key = "Room"
		fmt.Println("POST 🤝 successfully ")
		h.bookKey = strings.ToLower(ctx.Param("book"))
		store[h.key] = map[string]bool{EventKey: true}
		h.trigger <- store
	})

	app.GET(ListsURLs, func(ctx *gin.Context) {
		fmt.Println("sending to client")

		conn.ServeHTTP(ctx.Writer, ctx.Request)
	})

	app.POST(ListsURLs, func(ctx *gin.Context) {
		store := map[string]map[string]bool{}
		h.key = "Room"
		h.dictionaryKey = ctx.Param("dictionary")
		h.bookKey = strings.ToLower(ctx.Param("book"))
		fmt.Println("POST 🤝 successfully ")

		store[h.key] = map[string]bool{ItemKey: true}
		h.trigger <- store
	})

	app.GET(ItemsURLs, func(ctx *gin.Context) {
		fmt.Println("sending to client")

		conn.ServeHTTP(ctx.Writer, ctx.Request)
	})

	app.POST(ItemsURLs, func(ctx *gin.Context) {
		store := map[string]map[string]bool{}
		h.key = "Room"
		fmt.Println("POST 🤝 successfully ")
		h.bookKey = strings.ToLower(ctx.Param("book"))
		h.dictionaryKey = ctx.Param("dictionary")
		h.listKey = ctx.Param("list")

		store[h.key] = map[string]bool{ValidateKey: true}
		h.trigger <- store
	})

	go RunHub(h, conn)

	app.Run(":9120")

}
