package main

import (
	"fmt"
	"gorm-canban/entities"
	"gorm-canban/models"
	"time"
)
func PointAvailableTransactionInsert() {
	var patModel models.PointAvailableTransactionRepository
	patModel = models.NewPointAvailableTransactionRepository()
	count := 0
	var walletId uint64
	for i := 1; i < 5027646; i++ {
		walletId = uint64(i)
		if count % 2 == 0 {
			walletId = uint64(i)-1
		}
		count++
		// Create
		pat := entities.PointAvailableTransaction{

			WalletID:    walletId,
			Point:       1,
			BeforePoint: 0,
			AfterPoint:  1,
			Type:        2,
			TypeID:      0,
			OriginalID:  nil,
			CreatedAt:   time.Now(),
		}
		err := patModel.Create(&pat)
		if err != nil {
			fmt.Println("Inserted PointAvailableTransaction Failed")
		} else {
			fmt.Println("Insert PointAvailableTransaction success")
		}
	}
}
func PointPendingTransactionInsert() {
	var pptModel models.PointPendingTransactionRepository
	pptModel = models.NewPointPendingTransactionRepository()
	for i := 1; i < 5027646; i++ {
		// Create
		now := time.Now()
		ppt := entities.PointPendingTransaction{

			WalletID:      uint64(i),
			Point:         1,
			Rate:          0.01,
			ProcessedType: 2,
			CreatedAt:     now,
			UpdatedAt:     now,
		}
		err := pptModel.Create(&ppt)
		if err != nil {
			fmt.Println("Inserted PointPendingTransaction Failed")
		} else {
			fmt.Println("Insert PointPendingTransaction success")
		}
	}
}
func main() {
	//Postgres

	//PointAvailableTransactionInsert()
	//PointPendingTransactionInsert()

	var walletModel models.WalletRepository
	walletModel = models.NewWalletRepository()
	for i := 0; i < 5027646; i++ {
		// Create
		wallet := entities.Wallet{

			CurrencyID:          0,
			UserID:              uint64(i),
			BalanceAccumulative: 0,
			BalanceCurrent:      100,
			BalanceKYCApproach:  0,
			NegativeAmount:      0,
			CreatedAt:           time.Now(),
		}
		err := walletModel.Create(&wallet)
		if err != nil {
			fmt.Println("Inserted Failed")
		} else {
			fmt.Println("Insert success")
		}
	}
	//begin := time.Now()
	//wallets, _ := walletModel.FindAll()
	//end := time.Now()
	//fmt.Println("Time execute begin: ", (begin))
	//fmt.Println("Time execute end: ", (end))
	//fmt.Println("wallets", len(wallets))
	//for _, wallet := range wallets {
	//	fmt.Println(wallet)
	//	fmt.Println("--------------------")
	//}

	//MySQL
	//var productModel models.ProductRepository
	//productModel = models.NewProductRepository()
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
	//products, _ := productModel.SelectWithConditions(false)
	//products, _ := productModel.FindAllStoredProcedureNoParameter()
	//products, _ := productModel.FindByStoredProcedureHasParameter(10, 20)
	//for _, product := range products {
	//	fmt.Println(product.ToString())
	//	fmt.Println("-----------------")
	//}

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

	//result, _ := productModel.SumWithCalculate(true)
	//fmt.Println("Sum: ", result)

	//var facultyModel models.FacultyModel
	//faculties, _ := facultyModel.FindAll()
	//for _, faculty := range faculties {
	//	fmt.Println(faculty.ToString())
	//	fmt.Println("Students: ", len((faculty.Students)))
	//	if len(faculty.Students) > 0 {
	//		for _, student := range faculty.Students {
	//			fmt.Println(student.ToString())
	//			fmt.Println("==============")
	//		}
	//	}
	//	fmt.Println("-----------------------")
	//}
}
