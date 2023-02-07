package main

import (
	"log"

	"github.com/ChiefGupta/go-fiber-postgres/controllers"
	"github.com/ChiefGupta/go-fiber-postgres/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env variables", err.Error())
	}

	initializers.ConnectDB(&config)
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env variables", err)
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	app.Route("/users", func(router fiber.Router) {
		router.Post("", controllers.CreateUserHandler)
		router.Get("", controllers.FindUsers)
	})

	app.Route("/users/:userId", func(router fiber.Router) {
		router.Get("", controllers.FindUserById)
		router.Delete("", controllers.DeleteUser)
		router.Patch("", controllers.UpdateUser)
	})

	log.Fatal(app.Listen(":" + config.ServerPort))
}
