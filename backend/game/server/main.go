// Package server implements the gameplay
// all rights reserved, copyright 2025
package server

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run() {
	app := fiber.New()
	hub := NewHub()
	log.SetFlags(log.Lshortfile)
	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	Room(namespaces[1], app, hub)
	go RoomHub(hub)
	log.Println("🍀: http://localhost:3000")
	app.Listen(":3000")
}
