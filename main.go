package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	makeDir()
	app.Static("/", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ping!")
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("upload")
		if err != nil {
			return err
		}
		err = c.SaveFile(file, "public/uploads/"+file.Filename)
		if err != nil {
			return c.JSON(err)
		}
		return c.SendString("Uploaded")
	})

	fmt.Println(os.Getenv("PORT"))

	app.Listen(":" + os.Getenv("PORT"))
}

func makeDir() error {
	if err := os.MkdirAll("public/uploads", os.ModePerm); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
