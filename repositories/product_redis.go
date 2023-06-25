package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryRedis{db, redisClient}
}

func (r productRepositoryRedis) GetProducts() (products []product, err error) {

	key := "repository::GetProducts"

	//Redis Get

	productsJson, err := r.redisClient.Get(context.Background(), key).Result()

	if err == nil {
		json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	//DataBase

	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error

	if err != nil {
		return nil, err
	}

	//Redis set
	data, err := json.Marshal(products)

	if err != nil {
		return nil, err
	}

	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*30).Err()

	if err != nil {
		return nil, err
	}

	fmt.Println("database")

	return products, err
}
