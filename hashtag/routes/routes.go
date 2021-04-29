package routes

import (
	"hashtag/controllers"

	"github.com/gofiber/fiber/v2"
)

// PostsRoutes routes
func PostsRoutes(app fiber.Router, pc *controllers.PostsController) {

	posts := app.Group("/posts")

	posts.Post("/", func(c *fiber.Ctx) error {
		return pc.CreatePost(c)
	})

	posts.Get("/", func(c *fiber.Ctx) error {
		return pc.ListAllPosts(c)
	})

	posts.Get("/:id", func(c *fiber.Ctx) error {
		return pc.ListPostByID(c)
	})

	posts.Get("/tag/:tag", func(c *fiber.Ctx) error {
		return pc.ListPostByTags(c)
	})

}
