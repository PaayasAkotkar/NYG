package server

import (
	"github.com/gin-gonic/gin"
)

type Credit struct {
	Coin       string `json:"coins"`
	Spur       string `json:"spurs"`
	BoardTheme string `json:"boardTheme"`
	GuessTheme string `json:"guessTheme"`
	ID         string `json:"id"`
}

func Server() {
	app := gin.New()
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// adding the purchase
	app.PUT("/coin/purchase/:id", ServeCoin)

	app.PUT("/spur/purchase/:id", ServeSpur)
	// end adding

	// updating the upgrades and deducting the coins, spurs
	app.PATCH("/upgrade/power/:id", ServeUpgrades)

	app.Run(":" + port)
}
