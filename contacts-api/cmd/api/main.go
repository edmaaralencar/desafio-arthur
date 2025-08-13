package main

import (
	"log"

	_ "github.com/edmaaralencar/contacts-api/docs"
	apiError "github.com/edmaaralencar/contacts-api/internal/api-error"
	"github.com/edmaaralencar/contacts-api/internal/contacts"
	"github.com/edmaaralencar/contacts-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Contacts API
// @version 1.0
// @description Contacts API built with Go
// @host localhost:8080
// @BasePath /
func main() {
  db, err := database.ConnectAndMigrate()
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  err = database.SeedContacts(db)
  if err != nil {
    log.Fatal("Failed to seed database:", err)
  }

  store := contacts.NewSQLiteStore(db)

  app := fiber.New(fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
      var apiErr *apiError.APIError

      if ok := apiError.AsAPIError(err, &apiErr); ok {
        return c.Status(apiErr.Code).JSON(apiErr)
      }

      if fe, ok := err.(*fiber.Error); ok {
        return c.Status(fe.Code).JSON(&apiError.APIError{
          Code:    fe.Code,
          Message: fe.Message,
        })
      }

      return c.Status(fiber.StatusInternalServerError).JSON(&apiError.APIError{
        Code:    fiber.StatusInternalServerError,
        Message: "Internal server error",
      })
    },
  })

	app.Get("/docs/*", swagger.HandlerDefault)

  app.Get("/contacts", contacts.ListContacts(store))
  app.Post("/contacts", contacts.CreateContact(store))
  app.Delete("/contacts/:id", contacts.DeleteContact(store))

  log.Fatal(app.Listen(":8080"))
}
