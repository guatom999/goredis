package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type productRepositoeyDB struct {
	db *gorm.DB
}

func NewProductRepositoeyDB(db *gorm.DB) productRepositoeyDB {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoeyDB{db: db}
}

func mockData(db *gorm.DB) error {

	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	products := []product{}
	for i := 0; i < 5000; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("Product %v", i+1),
			Quantity: random.Intn(100),
		})
	}

	return db.Create(&products).Error

}

func (r productRepositoeyDB) GetProducts() (products []product, err error) {

	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error

	return products, err
}
