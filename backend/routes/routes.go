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
func AdminRoutes(app fiber.Router, ac *controllers.AdminController) {
	admin := app.Group("/admin")

	admin.Post("/create", func(c *fiber.Ctx) error {
		return ac.CreateAdmin(c)
	})

	admin.Post("/login", func(c *fiber.Ctx) error {
		return ac.Login(c)
	})

	admin.Get("/", func(c *fiber.Ctx) error {
		return ac.ListAdmins(c)
	})
}

// UsersRoutes routes
func UsersRoutes(app fiber.Router, ac *controllers.UsersController) {
	users := app.Group("/users")

	users.Post("/", func(c *fiber.Ctx) error {
		return ac.CreateUser(c)
	})
	users.Get("/", func(c *fiber.Ctx) error {
		return ac.ReadAllUsers(c)
	})
	users.Get("/read/:id", func(c *fiber.Ctx) error {
		return ac.ReadUserByID(c)
	})

	users.Put("/:id", func(c *fiber.Ctx) error {
		return ac.UpdateUserByID(c)
	})

	users.Delete("/:id", func(c *fiber.Ctx) error {
		return ac.DeleteUserByID(c)
	})

	users.Get("/ml/:id", func(c *fiber.Ctx) error {
		return ac.ReadFaceID(c)
	})

}
