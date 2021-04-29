package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	const URL = "http://localhost:8000/"

	// Routes
	app.Post("/fileupload", func(c *fiber.Ctx) error {

		file, err := c.FormFile("file1")
		if err != nil {
			return err
		}
		// file1 := c.FormValue("name")

		fileType := strings.Split(file.Filename, ".")[1]

		if isAllowedExt := allowedFileType(fileType); !isAllowedExt {
			return c.SendString("no type")
		}

		uid, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
		}

		file.Filename = uid.String() + "." + fileType

		if err := c.SaveFile(file, fmt.Sprintf("../ml/img/known/%s", file.Filename)); err != nil {
			log.Println("err ", err)
		}
		return c.SendString("done")
	})

	app.Listen(":3000")
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
