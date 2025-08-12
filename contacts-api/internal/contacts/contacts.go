package contacts

import (
	"context"
	"time"
)

type Contact struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CpfCnpj  string `json:"cpf_cnpj"`
	Phone    string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Store interface {
	List(ctx context.Context) ([]Contact, error)
	ListPaginated(ctx context.Context, page, perPage int) ([]Contact, int, error)
	Create(ctx context.Context, contact *CreateContactRequest) (error)
}