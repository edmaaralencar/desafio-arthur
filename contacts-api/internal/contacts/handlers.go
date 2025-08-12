package contacts

import (
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