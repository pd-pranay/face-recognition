package routes

import (
	"backend/controllers"

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

// AdminRoutes routes
func AdminRoutes(app fiber.Router, pc *controllers.AdminController) {
	// admin := app.Group("/admin")

	// admin.Post("/create", func(c *fiber.Ctx) error {
	// 	// return error
	// })

	// admin.Post("/login", func(c *fiber.Ctx) error {
	// 	// return error
	// })
}

// UsersRoutes routes
func UsersRoutes(app fiber.Router, pc *controllers.UsersController) {
	// users := app.Group("/users")

	// admin.Post("/", func(c *fiber.Ctx) error {
	// 	// return error
	// })
	// users.Get("/", func(c *fiber.Ctx) error {
	// 	// return error
	// })

}
