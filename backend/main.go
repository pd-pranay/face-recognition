package main

import (
	"backend/controllers"
	"backend/db"
	"backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	dbEngine := "postgres"
	dbUser := "hashtag"
	dbPassword := "hashtag"
	dbName := "hashtag"
	dbHost := "localhost"
	dbPort := "7777"
	dbSSLMode := "disable"

	// Connect to Database
	dbCon, err := db.Connect(dbEngine, dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)
	if err != nil {
		log.Fatal(err)
	}
	defer dbCon.Close()

	queries := db.New(dbCon)

	postController := controllers.NewPostsController(dbCon, queries)
	adminController := controllers.NewAdminController(dbCon, queries)
	usersController := controllers.NewUsersController(dbCon, queries)

	app := fiber.New()

	v1 := app.Group("/v1/api")

	routes.PostsRoutes(v1, postController)
	routes.AdminRoutes(v1, adminController)
	routes.UsersRoutes(v1, usersController)

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Use(cors.New(cors.Config{
		// AllowCredentials: false,
		AllowOrigins: "http://localhost:3000/",
	}))

	log.Fatal(app.Listen(":5000"))

}
