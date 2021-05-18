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

	imgPath := fmt.Sprintf("../ml/api/images/images_test/%s", file.Filename)
	if err := c.SaveFile(file, imgPath); err != nil {
		log.Println("err ", err)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}

	body := db.CreateUserParams{
		Name:        sql.NullString{String: c.FormValue("name"), Valid: true},
		CollegeName: sql.NullString{String: c.FormValue("college_name"), Valid: true},
		Address:     sql.NullString{String: c.FormValue("address"), Valid: true},
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
	log.Println("c.Params", c.Params("id"))
	id := (c.Params("id"))
	if id == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": " err.Error()",
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

func (uc *UsersController) UpdateUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

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

	imgPath := fmt.Sprintf("../ml/api/images/images_test/%s", file.Filename)
	if err := c.SaveFile(file, imgPath); err != nil {
		log.Println("err ", err)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}

	body := db.UpdateUserParams{
		ImageUid:    id,
		ImageUid_2:  uid.String(),
		Name:        sql.NullString{String: c.FormValue("name"), Valid: true},
		CollegeName: sql.NullString{String: c.FormValue("college_name"), Valid: true},
		Address:     sql.NullString{String: c.FormValue("address"), Valid: true},
		MobileNo:    sql.NullInt32{Int32: 78, Valid: true},
		ImagePath:   file.Filename,
	}

	users, err := uc.Queries.UpdateUser(c.Context(), body)
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

func (uc *UsersController) DeleteUserByID(c *fiber.Ctx) error {
	if err := uc.Queries.DeleteUsersById(c.Context(), c.Params("id")); err != nil {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"code": 500,
				"data": err.Error(),
			})
	}
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code": 200,
			"data": "Success",
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

	tags := strings.Split(id, ",")
	values := make([]string, 0, len(tags))
	for _, s := range tags {
		values = append(values, s)
	}

	match, err := uc.Queries.ReadUsersByFace(c.Context(), values)
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
