package controllers

import (
	"backend/db"
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// AdminController is a controller and is defined here.
type AdminController struct {
	DB      *sql.DB
	Queries *db.Queries
}

// NewAdminController returns pointer to AdminController.
func NewAdminController(db *sql.DB, queries *db.Queries) *AdminController {
	return &AdminController{
		DB:      db,
		Queries: queries,
	}
}

func (ac *AdminController) CreateAdmin(c *fiber.Ctx) error {
	body := db.CreateAdminParams{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": "Error reading body",
		})
	}

	bytePwd := []byte(body.Password)
	hashedPwd, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": "Error in encrypting password",
		})
	}

	body.Password = string(hashedPwd)

	admin, err := ac.Queries.CreateAdmin(c.Context(), body)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code": 200,
			"data": admin,
		})
}

func (ac *AdminController) Login(c *fiber.Ctx) error {

	type login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type adminModel struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	body := login{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": "Error reading body",
		})
	}

	admin, err := ac.Queries.Login(c.Context(), body.Email)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}

	byteHashedPwd := []byte(admin.Password)
	bytePlainPwd := []byte(body.Password)
	err = bcrypt.CompareHashAndPassword(byteHashedPwd, bytePlainPwd)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": "Invalid username or password",
		})
	}
	adminData := adminModel{admin.ID.String(), admin.Email, admin.Name}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = admin.Name
	claims["id"] = admin.ID
	claims["email"] = admin.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	accessToken, err := token.SignedString([]byte("wQ6u4Jt7vedu2UZughjcPzyo2"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code":  200,
			"data":  adminData,
			"token": accessToken,
		})
}

func (ac *AdminController) ListAdmins(c *fiber.Ctx) error {
	admins, err := ac.Queries.ListAllAdmin(c.Context())
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":  500,
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code": 200,
			"data": admins,
		})
}
