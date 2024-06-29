package handlers

import (
	"github.com/gofiber/fiber/v3"
	"natthan.com/go-play/services"
)

type catalogHandler struct {
	catalogSrv services.CatalogService
}

func NewCatalogHandler(catalogSrv services.CatalogService) CatalogHandler {
	return catalogHandler{catalogSrv: catalogSrv}
}

func (h catalogHandler) GetProducts(c fiber.Ctx) error {
	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	return c.JSON(response)
}
