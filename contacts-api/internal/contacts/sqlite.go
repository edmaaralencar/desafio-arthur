package contacts

import (
	"context"
	"database/sql"
)

type sqliteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) Store {
	return &sqliteStore{db: db}
}

func (s *sqliteStore) List(ctx context.Context) ([]Contact, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, email, cpf_cnpj, phone FROM contacts")
	if err != nil {
			return nil, err
	}
	defer rows.Close()

	var list []Contact
	for rows.Next() {
			var c Contact
			if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.CpfCnpj, &c.Phone); err != nil {
					return nil, err
			}
			list = append(list, c)
	}
	return list, nil
}

func (s *sqliteStore) ListPaginated(ctx context.Context, page, perPage int) ([]Contact, int, error) {
	offset := (page - 1) * perPage

	var total int
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM contacts").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := s.db.QueryContext(ctx,
		`SELECT id, name, email, cpf_cnpj, phone, created_at, updated_at
		 FROM contacts
		 LIMIT ? OFFSET ?`,
		perPage, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []Contact
	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.CpfCnpj, &c.Phone, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, 0, err
		}
		list = append(list, c)
	}

	return list, total, nil
}

func (s *sqliteStore) Create(ctx context.Context, contact *CreateContactRequest) (error) {
	_, err := s.db.ExecContext(ctx, 
		"INSERT INTO contacts(name, email, cpf_cnpj, phone) VALUES (?, ?, ?, ?)",
		contact.Name, contact.Email, contact.CpfCnpj, contact.Phone,
	)

	if err != nil {
		return err
	}

	return nil
}