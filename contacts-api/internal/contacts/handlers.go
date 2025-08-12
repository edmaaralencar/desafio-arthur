package contacts

import (
	"encoding/json"

	"github.com/edmaaralencar/contacts-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func ListContacts(store Store) fiber.Handler {
	return func (c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		perPage := c.QueryInt("per_page", 10)

		contacts, total, err := store.ListPaginated(c.Context(), page, perPage)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
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

func CreateContact(store Store) fiber.Handler {
	return func (c *fiber.Ctx) error {
		var body CreateContactRequest

		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid json")
		}

		if errors := utils.ValidateStruct(body); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": errors,
			})
		}

		isValid, cleanedDoc := utils.ValidateCpfCnpj(body.CpfCnpj)

		if !isValid {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid CPF or CNPJ"})
		}

		body.CpfCnpj = cleanedDoc

		store.Create(c.Context(), &body)
		
		return nil
	}
}