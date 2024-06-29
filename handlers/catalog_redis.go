package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
	"natthan.com/go-play/services"
)

type catalogHandlerRedis struct {
	catalogSrv  services.CatalogService
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(catalogSrv services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{catalogSrv: catalogSrv, redisClient: redisClient}
}

func (h catalogHandlerRedis) GetProducts(c fiber.Ctx) error {
	key := "handler::GetProducts"
	// Redis GET
	if responseJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}

	// Service
	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	// Redis SET
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, data, 10*time.Second)
	}
	fmt.Println("database")

	return c.JSON(response)
}
