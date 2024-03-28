package adapters

import (
	"fmt"
	"hexagonal/core"

	"github.com/gofiber/fiber/v2"
)

// ------- [ handler adapter dones't have Interface (same as controller in MVC) ] -----

// Primary adapter - [ controller dependency ]
type HttpOrderHandler struct {
	service core.OrderService
}

// struct init - [ controller constructor for set default dependency value ]
func NewHttpOrderHandler(service core.OrderService) *HttpOrderHandler {
	return &HttpOrderHandler{service: service}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order core.Order
	if err := c.BodyParser(&order); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateOrder(order); err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
