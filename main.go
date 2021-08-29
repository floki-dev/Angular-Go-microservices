package main

import "github.com/gofiber/fiber/v2"

func main() {
  _, err := gorm.Open(postgres.Open("root:root@tcp(postgres:5432)/ambassador"), &gorm.Config{})

  if err != nil {
    panic("Could not connect with the database!")
  }

  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  app.Listen(":8000")
}
