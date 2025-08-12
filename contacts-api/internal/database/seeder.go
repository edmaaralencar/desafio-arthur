package database

import (
	"database/sql"
	"log"
)

func SeedContacts(db *sql.DB) error {
	contacts := []struct {
		Name    string
		Email   string
		CpfCnpj string
		Phone   string
	}{
		{"Alice Silva", "alice@example.com", "12345678900", "11999999999"},
		{"Bruno Souza", "bruno@example.com", "56379482000197", "11888888888"},
		{"Carla Lima", "carla@example.com", "11122233344", "11777777777"},
	}

	stmt, err := db.Prepare("INSERT OR IGNORE INTO contacts(name, email, cpf_cnpj, phone) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, c := range contacts {
		_, err := stmt.Exec(c.Name, c.Email, c.CpfCnpj, c.Phone)
		if err != nil {
				log.Printf("Failed to insert contact %s: %v\n", c.Name, err)
		}
	}

	return nil
}
