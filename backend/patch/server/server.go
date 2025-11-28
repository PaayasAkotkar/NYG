package server

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Server() {
	app := gin.New()
	log.Println("in server 🪛")

	// setting user name and password for mysql
	err := os.Setenv("root", "root")
	if err != nil {
		log.Fatalf("Error setting MYSQL_USER: %v", err)
	}

	err = os.Setenv("kingp12", "kingp12")
	if err != nil {
		log.Fatalf("Error setting MYSQL_PASS: %v", err)
	}
	strict := cors.Config{
		AllowOrigins:     allowOrgs,
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowCredentials: true,
	}
	app.Use(cors.New(strict))
	// app.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH,")
	// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// 	c.Writer.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(200)
	// 		return
	// 	}
	// 	c.Next()
	// })

	app.PUT("/event/update", PutEvent)

	// to create new set of deck
	app.PUT("/deck/update", PatchDeck)

	// app.PATCH("/profile/update", PatchProfile)

	app.PATCH("/coin/update", PatchCoin)
	app.PATCH("/spur/update", PatchSpur)

	app.PATCH("/credits/update", PatchCredits)

	app.PATCH("/theme/update", PatchTheme)

	app.PATCH("/img/update", PatchImage)

	// app.PATCH("/name/update", PatchName)

	// app.PATCH("/nickname/update", PatchNickname)

	// current power up upgrades
	app.PATCH("/power-up/update", PatchPowerUp)

	// creates profile
	app.POST("/profile/init", InitProfile)

	app.Static("/uploads/", "./server/uploads") // ./server because the root of the main is naked

	app.Run(":" + port)
}
