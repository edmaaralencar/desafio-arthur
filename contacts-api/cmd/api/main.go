package main

import (
	"log"

	"github.com/edmaaralencar/contacts-api/internal/contacts"
	"github.com/edmaaralencar/contacts-api/internal/database"
	"github.com/gofiber/fiber/v2"
)

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

  app.Get("/contacts", contacts.ListContacts(store))

  log.Fatal(app.Listen(":3000"))
}
