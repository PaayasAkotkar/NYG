package server

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Server() {
	app := gin.New()
	log.Println("NYG PATCH SERVER RUNNING ⚙️")

	// setting user name and password for mysql
	err := os.Setenv(USER, USER)
	if err != nil {
		log.Fatalf("Error setting MYSQL_USER: %v", err)
	}

	err = os.Setenv(PASSWORD, PASSWORD)
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

	app.PUT("/event/update", PutEvent)

	// to create new set of deck
	app.PUT("/deck/update", PatchDeck)

	app.PATCH("/coin/update", PatchCoin)
	app.PATCH("/spur/update", PatchSpur)

	app.PATCH("/credits/update", PatchCredits)

	app.PATCH("/theme/update", PatchTheme)

	app.PATCH("/img/update", PatchImage)

	// current power up upgrades
	app.PATCH("/power-up/update", PatchPowerUp)

	// creates profile
	app.POST("/profile/init", InitProfile)

	app.Static("/uploads/", "./server/uploads") // ./server because the root of the main is naked

	app.Run(":" + port)
}
