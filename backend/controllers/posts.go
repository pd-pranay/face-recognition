package controllers

import (
	"backend/db"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PostsController is a controller and is defined here.
type PostsController struct {
	DB      *sql.DB
	Queries *db.Queries
}

// NewPostsController returns pointer to PostsController.
func NewPostsController(db *sql.DB, queries *db.Queries) *PostsController {
	return &PostsController{
		DB:      db,
		Queries: queries,
	}
}

// CreatePost single
func (pc *PostsController) CreatePost(c *fiber.Ctx) error {

	body := db.CreatePostParams{}

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}

	post, err := pc.Queries.CreatePost(c.Context(), body)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"success": true,
			"data":    post,
		})

}

// ListAllPosts single
func (pc *PostsController) ListAllPosts(c *fiber.Ctx) error {

	posts, err := pc.Queries.GetAllPosts(c.Context())
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    []db.GetAllPostsRow(posts),
	})
}

// ListPostByID single
func (pc *PostsController) ListPostByID(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is empty",
		})
	}

	uid := uuid.MustParse(c.Params("id"))

	posts, err := pc.Queries.GetPostByID(c.Context(), uid)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    db.GetPostByIDRow(posts),
	})
}

// ListPostByTags single
func (pc *PostsController) ListPostByTags(c *fiber.Ctx) error {

	tag := c.Params("tag")
	if tag == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "tag is empty",
		})
	}

	posts, err := pc.Queries.GetPostsByTag(c.Context(), []string{tag})

	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    []db.GetPostsByTagRow(posts),
	})
}
