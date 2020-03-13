package entities

import (
	"fmt"
	"time"
)

type Product struct {
	Id       int `gorm:"primary_key, AUTO_INCREMENT"`
	Name     string
	Price    float64
	Quantity int
	Status   bool
	Created  *time.Time
}

func (product *Product) TableName() string {
	return "product"
}
func (product Product) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\nprice: %0.1f\nquantity: %d\nstatus: %t\ncreated: %s", product.Id, product.Name, product.Price, product.Quantity, product.Status, product.Created.Format("02/01/2006"))
}

type ProductInfo struct {
	Id    int
	Name  string
	Price float64
}

func (productInfo ProductInfo) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\nPrice: %0.1f", productInfo.Id, productInfo.Name, productInfo.Price)
}
