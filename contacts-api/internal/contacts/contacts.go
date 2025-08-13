package contacts

import (
	"context"
	"time"
)

type Contact struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CpfCnpj  string `json:"cpfCnpj"`
	Phone    string `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Store interface {
	List(ctx context.Context) ([]Contact, error)
	ListPaginated(ctx context.Context, page, perPage int) ([]Contact, int, error)
	Create(ctx context.Context, contact *CreateContactRequest) (error)
	Delete(ctx context.Context, id int64) (error)
}