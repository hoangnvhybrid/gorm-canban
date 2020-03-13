package models

import (
	"gorm-canban/config"
	"gorm-canban/entities"
)

type ProductRepository interface {
	Create(product *entities.Product) error
	Update(product *entities.Product) error
	FindById(id int) (*entities.Product, error)
	Delete(product entities.Product) error
	FindAll() ([]entities.Product, error)
	FindWithConditions(status bool, min, max float64) ([]entities.Product, error)
	FindProductsWithBetween(min, max float64) ([]entities.Product, error)
	FindByWhereLike(keyword string) ([]entities.Product, error)
	FindAndOrderBy(typeOrderBy string) ([]entities.Product, error)
	OrderByAndConditions(status bool) ([]entities.Product, error)
	LimitAndWhereAndOrderBy(status bool, n int) ([]entities.Product, error)
	SelectWithConditions(status bool) ([]entities.ProductInfo, error)
	GroupBy() ([]entities.ProductGroup, error)
	FindAllStoredProcedureNoParameter() ([]entities.Product, error)
	FindByStoredProcedureHasParameter(min, max float64) ([]entities.Product, error)
	SumWithCalculate(status bool) (float64, error)
}

func NewProductRepository() ProductRepository {
	return &productModel{}
}

type productModel struct {
}

func (*productModel) SumWithCalculate(status bool) (float64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	}
	var result float64
	row := db.Table("product").Where("status = ?", status).Select("sum(price * quantity)").Row()
	row.Scan(&result)
	return result, nil
}
func (*productModel) FindAllStoredProcedureNoParameter() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	//DELIMITER $$
	//CREATE PROCEDURE sp_findAll()
	//BEGIN
	//SELECT * FROM product;
	//END $$
	//DELIMITER ;
	var products []entities.Product
	db.Raw("call sp_findAll()").Scan(&products)
	return products, nil
}
func (*productModel) FindByStoredProcedureHasParameter(min, max float64) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	//sql
	//DELIMITER $$
	//CREATE PROCEDURE findBetween(min decimal, max decimal)
	//BEGIN
	//SELECT * FROM product where price BETWEEN min and max;
	//END $$
	//DELIMITER ;
	var products []entities.Product
	db.Raw("call findBetween(?, ?)", min, max).Scan(&products)
	return products, nil
}
func (*productModel) GroupBy() ([]entities.ProductGroup, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var productGroups []entities.ProductGroup
	db.Table("product").
		Select("status, count(id) as result1, sum(quantity) as result2, min(price) as result3, max(price) as result4, avg(price) as result5").
		Group("status").
		Having("count(id) > 2").
		Scan(&productGroups)
	return productGroups, nil
}
func (*productModel) SelectWithConditions(status bool) ([]entities.ProductInfo, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entities.ProductInfo
	//TODO: chua phan biet duoc debug va ko debug
	if err = db.Debug().Table("product").Where("status = ?", status).Select("id, name, price").Find(&products).Error; err != nil {
		//if err = db.Table("product").Where("status = ?", status).Select("id, name, price").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (*productModel) LimitAndWhereAndOrderBy(status bool, n int) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	if err = db.Where("status = ?", status).Order("price desc").Limit(n).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (productModel *productModel) Create(product *entities.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	return db.Create(product).Error
}

func (productModel *productModel) Update(product *entities.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	return db.Save(&product).Error

}

//update nay ko khoa ban ghi lai, co the dan den xung dot
func (productModel *productModel) FindById(id int) (*entities.Product, error) {
	db, err := config.GetDB()
	var product entities.Product
	if err != nil {
		return nil, err
	}
	result := db.Where("id = ?", id).First(&product)
	if result.Error != nil {
		if result.RecordNotFound() {
			return nil, nil
		}
		return nil, result.Error
	}
	return &product, nil
}

//delete vat ly (khong phai delete mem bang cach chuyen truong deleted_at khac null)
func (productModel *productModel) Delete(product entities.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	return db.Delete(&product).Error
}

func (productModel *productModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (*productModel) FindWithConditions(status bool, min, max float64) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	db.Where("status = ? and price >? and price <= ?", status, min, max).Find(&products)
	return products, nil
}
func (*productModel) FindProductsWithBetween(min, max float64) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	if err = db.Where("price between ?  and ?", min, max).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (*productModel) FindByWhereLike(keyword string) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	if err = db.Where("name like ?", "%"+keyword+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (*productModel) FindAndOrderBy(typeOrderBy string) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	if err = db.Order("price " + typeOrderBy).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (*productModel) OrderByAndConditions(status bool) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	if err = db.Where("status = ?", status).Order("price desc").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
