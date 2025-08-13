package main

import (
	"log"

	_ "github.com/edmaaralencar/contacts-api/docs"
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

  app := fiber.New()

	app.Get("/docs/*", swagger.HandlerDefault)

  app.Get("/contacts", contacts.ListContacts(store))
  app.Post("/contacts", contacts.CreateContact(store))
  app.Delete("/contacts/:id", contacts.DeleteContact(store))

  log.Fatal(app.Listen(":8080"))
}
