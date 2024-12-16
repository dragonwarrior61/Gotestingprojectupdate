package Models

import (
	"go-gin-mysql-boilerplate/Config"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Validations"

	_ "github.com/go-sql-driver/mysql"
)

func ProductFetchAll(product *[]Schema.Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func ProductFetchSingle(product *Schema.Product, userId string) (err error) {
	if err = Config.DB.Where("id = ?", userId).First(product).Error; err != nil {
		return err
	}
	return nil
}

func ProductCreate(request *Validations.ProductCreate) (err error) {
	if err = Config.DB.Table("products").Create(request).Error; err != nil {
		return err
	}
	return nil
}

func ProductUpdate(request *Validations.ProductUpdate, userId string) (err error) {
	if err = Config.DB.Table("products").Where("id = ?", userId).Update(request).Error; err != nil {
		return err
	}
	return nil
}

func ProductDelete(product *Schema.Product, userId string) (err error) {
	if err = Config.DB.Where("id = ?", userId).Delete(product).Error; err != nil {
		return err
	}
	return nil
}

func ProductFetchWithEmail(product *Schema.Product, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(product).Error; err != nil {
		return err
	}
	return nil
}
