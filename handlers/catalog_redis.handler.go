package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type catalogRedisHandler struct {
	catalogService services.CatalogService
	redisClient    *redis.Client
}

func NewCatalogRedisHandler(catalogService services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogRedisHandler{catalogService, redisClient}
}

func (r catalogRedisHandler) GetProducts(c *fiber.Ctx) error {

	key := "handler:GetProducts"

	// Redis Get

	reponseJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {

		fmt.Print("Handler Redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(reponseJson)

	}

	// Service Get

	products, err := r.catalogService.GetProducts()

	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	//Redis Set

	if data, err := json.Marshal(response); err == nil {
		r.redisClient.Set(context.Background(), key, string(data), time.Second*30)
	}

	fmt.Print("Service DataBase")

	return c.JSON(response)
}
