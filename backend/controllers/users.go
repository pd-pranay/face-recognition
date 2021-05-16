package controllers

import (
	"backend/db"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UsersController is a controller and is defined here.
type UsersController struct {
	DB      *sql.DB
	Queries *db.Queries
}

// NewUsersController returns pointer to UsersController.
func NewUsersController(db *sql.DB, queries *db.Queries) *UsersController {
	return &UsersController{
		DB:      db,
		Queries: queries,
	}
}

func allowedFileType(ext string) bool {
	extensions := []string{"jpeg", "jpg", "png"}
	for _, extension := range extensions {
		if ext == extension {
			return true
		}
	}
	return false
}

func (uc *UsersController) CreateUser(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}

	fileType := strings.Split(file.Filename, ".")[1]

	if isAllowedExt := allowedFileType(fileType); !isAllowedExt {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": "File extension no allowed",
		})
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}

	file.Filename = uid.String() + "." + fileType

	imgPath := fmt.Sprintf("../ml/img/known/%s", file.Filename)
	if err := c.SaveFile(file, imgPath); err != nil {
		log.Println("err ", err)
	}

	body := db.CreateUserParams{
		Name:        sql.NullString{String: c.FormValue("name"), Valid: true},
		CollegeName: sql.NullString{String: c.FormValue("college_name"), Valid: true},
		MobileNo:    sql.NullInt32{Int32: 78, Valid: true},
		ImagePath:   file.Filename,
		ImageUid:    uid.String(),
	}

	users, err := uc.Queries.CreateUser(c.Context(), body)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code": 200,
			"data": users,
		})

}

func (uc *UsersController) ReadAllUsers(c *fiber.Ctx) error {
	allUsers, err := uc.Queries.ReadAllUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code": 200,
			"data": allUsers,
		})
}

func (uc *UsersController) ReadUserByID(c *fiber.Ctx) error {
	id, err := uuid.FromBytes([]byte(c.Params("id")))
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": "Empty ID",
		})
	}

	user, err := uc.Queries.ReadUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code": 200,
			"data": user,
		})
}

func (uc *UsersController) ReadFaceID(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": "empty id",
		})
	}
	match, err := uc.Queries.ReadUsersByFace(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code": 200,
			"data": match,
		})
}

func (uc *UsersController) DeleteUserByID(c *fiber.Ctx) error {
	id, err := uuid.FromBytes([]byte(c.Params("id")))
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": "Empty ID",
		})
	}

	user, err := uc.Queries.UpdateUserFlush(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code": 200,
			"data": user,
		})
}

// --name: ModifyUserByID :one
// UPDATE users SET name = ($2), college_name = ($3), address = ($4), mobile_no = ($5), image_path = ($6), image_uid = ($7) WHERE id = ($1) RETURNING *;
func (uc *UsersController) UpdateUserByID(c *fiber.Ctx) error {
	return nil
}
