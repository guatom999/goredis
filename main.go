package main

import (
	"fmt"
	"goredis/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := initDatabase()

	db.AutoMigrate(&product{})
	productRepo := repositories.NewProductRepositoryRedis(db)

	products, err := productRepo.GetProducts()

	if err != nil {
		panic(err)
	}

	fmt.Println(products)

	// app := fiber.New()

	// app.Get("/hello", func(c *fiber.Ctx) error {
	// 	time.Sleep(10 * time.Millisecond)
	// 	return c.SendString("Hello World")
	// })
	// app.Listen(":8000")

}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:Bossza555@tcp(localhost:3306)/redistest")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
