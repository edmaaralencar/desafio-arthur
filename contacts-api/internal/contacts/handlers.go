package contacts

import (
	"database/sql"
	"encoding/json"
	"errors"

	apiError "github.com/edmaaralencar/contacts-api/internal/api-error"
	"github.com/edmaaralencar/contacts-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type ListContactsResponse struct {
	Total    int       `json:"total"`
	Contacts []Contact `json:"contacts"`
	Page     int       `json:"page"`
	PerPage  int       `json:"perPage"`
}

// ListContacts godoc
// @Summary      List contacts
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number (default 1)"
// @Param        per_page  query     int     false  "Items per page (default 10)"
// @Success      200  {object}  ListContactsResponse
// @Router       /contacts [get]
func ListContacts(store Store) fiber.Handler {
	return func (c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		perPage := c.QueryInt("per_page", 10)

		contacts, total, err := store.ListPaginated(c.Context(), page, perPage)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
		}

		if contacts == nil {
			contacts = []Contact{}
		}

		for i := range contacts {
			contacts[i].CpfCnpj = utils.FormatCpfCnpj(contacts[i].CpfCnpj)
			contacts[i].Phone = utils.FormatPhoneWithDDD(contacts[i].Phone)
		}

		response := fiber.Map{
			"total": total,
			"contacts": contacts,
			"page": page,
			"perPage": perPage,
		}

		return c.JSON(response)
	}
}

type CreateContactRequest struct {
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required"`
	CpfCnpj string `json:"cpfCnpj" validate:"required"`
}

type APIError struct {
	Error string `json:"error"`
}

// CreateContact godoc
// @Summary      Create a new contact
// @Description  Creates a contact with name, email, phone, and CPF/CNPJ
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        request  body      CreateContactRequest  true  "Contact data"
// @Success      201      {object}  CreateContactRequest
// @Failure      400      {object}  APIError
// @Router       /contacts [post]
func CreateContact(store Store) fiber.Handler {
	return func (c *fiber.Ctx) error {
		var body CreateContactRequest

		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return apiError.NewAPIError(fiber.StatusBadRequest, "Invalid JSON")
		}

		if errors := utils.ValidateStruct(body); errors != nil {
			return apiError.SendAPIError(c, fiber.StatusBadRequest, "Validation failed", errors)
		}

		isValid, cleanedDoc := utils.ValidateCpfCnpj(body.CpfCnpj)

		if !isValid {
			return apiError.NewAPIError(fiber.StatusBadRequest, "Invalid CPF or CNPJ")
		}

		exists, err := store.FindByCpfOrCnpj(c.Context(), body.CpfCnpj)

		if err != nil {
			return apiError.NewAPIError(fiber.StatusBadRequest, "Database error")
		}

		if exists {
			return apiError.NewAPIError(fiber.StatusBadRequest, "CPF/CNPJ already exists")
		}

		body.CpfCnpj = cleanedDoc

		store.Create(c.Context(), &body)
		
		return c.SendStatus(fiber.StatusCreated)
	}
}

// DeleteHandler godoc
// @Summary      Delete a contact
// @Description  Deletes a contact by ID
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Contact ID"
// @Success      204  "No Content"
// @Failure      400  {object}  APIError
// @Failure      404  {object}  APIError
// @Failure      500  {object}  APIError
// @Router       /contacts/{id} [delete]
func DeleteContact(store Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return apiError.NewAPIError(fiber.StatusBadRequest, "Invalid ID")
		}

		if err := store.Delete(c.Context(), int64(id)); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return apiError.NewAPIError(fiber.StatusNotFound, "Contact not found")
			}

			return apiError.NewAPIError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
