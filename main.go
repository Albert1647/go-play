package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"natthan.com/go-play/handlers"
	"natthan.com/go-play/repositories"
	"natthan.com/go-play/services"
)

func main() {
	db := initDatabse()
	redisClient := initRedis()
	_ = redisClient
	// Try Changing Adapter to load test performance
	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	productService := services.NewCatalogService(productRepo)
	productHandler := handlers.NewCatalogHandlerRedis(productService, redisClient)

	app := fiber.New()
	app.Get("/products", productHandler.GetProducts)
	app.Listen(":8000")
}

func initDatabse() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/infinitas")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
