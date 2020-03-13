package main

import (
	"fmt"
	"gorm-canban/models"
)

func main() {
	var productModel models.ProductRepository
	productModel = models.NewProductRepository()
	// Create
	//now := time.Now()
	//product := entities.Product{
	//	Name:     "Bphone",
	//	Price:    37.032,
	//	Quantity: 50,
	//	Status:   true,
	//	Created:  &now,
	//}
	//err := productModel.Create(&product)
	//if err != nil {
	//	fmt.Println("Inserted Failed")
	//} else {
	//	fmt.Println(product.ToString())
	//}

	//Update
	//id := 20
	//product, err := productModel.FindById(id)
	//if err != nil {
	//	fmt.Println("Find erro", err.Error())
	//}
	//if product == nil {
	//	fmt.Println("Find not found product with id is ", id)
	//	return
	//}
	//product.Name = "Tu lanh 3 ngan panasonic"
	//product.Price = 1203.345
	//product.Status = true
	//product.Quantity = 20
	//
	//now := time.Now()
	//product.Created = &now
	//err = productModel.Update(product)
	//if err != nil {
	//	fmt.Println("Update faild")
	//} else {
	//	fmt.Println(product.ToString())
	//}

	//delete
	//id := 16
	//product, err := productModel.FindById(id)
	//if err != nil {
	//	fmt.Println("Find erro", err.Error())
	//}
	//if product == nil {
	//	fmt.Println("Find not found product with id is ", id)
	//	return
	//}
	//err = productModel.Delete(*product)
	//if err != nil {
	//	fmt.Println("Delete failed")
	//} else {
	//	fmt.Println("Done")
	//}

	// find all
	//products, _ := productModel.FindAll()
	//for _, product := range products {
	//	fmt.Println(product.ToString())
	//	fmt.Println("--------------------")
	//}

	//find product with conditions
	//products, _ := productModel.FindWithConditions(true, 10, 21)
	//for _, product := range products {
	//	fmt.Println(product.ToString())
	//	fmt.Println("-------------------")
	//}

	//find between
	//products, _ := productModel.FindProductsWithBetween(10, 20)
	//for _, product := range products {
	//	fmt.Println(product.ToString())
	//	fmt.Println("---------------------")
	//}

	////find like
	//products, _ := productModel.FindByWhereLike("tivi")
	//for _, product := range products {
	//	fmt.Println(product.ToString())
	//	fmt.Println("-----------------")
	//}

	//find and order by
	//products, _ := productModel.FindAndOrderBy("desc")
	//products, _ := productModel.OrderByAndConditions(false)
	//products, _ := productModel.LimitAndWhereAndOrderBy(true, 3)
	products, _ := productModel.SelectWithConditions(false)
	for _, product := range products {
		fmt.Println(product.ToString())
		fmt.Println("-----------------")
	}

	//products, _ := productModel.GroupBy()
	//for _, product := range products {
	//	fmt.Println("Status:", product.Status)
	//	fmt.Println("Count product:", product.Result1)
	//	fmt.Println("Sum quantity:", product.Result2)
	//	fmt.Println("Min price:", product.Result3)
	//	fmt.Println("Max price:", product.Result4)
	//	fmt.Println("Avg price:", product.Result5)
	//	fmt.Println("-----------------")
	//}
}
