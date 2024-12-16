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

func OrderFetchSingle(order *Schema.Order, userId string) (err error) {
	if err = Config.DB.Where("id = ?", userId).First(order).Error; err != nil {
		return err
	}
	return nil
}

func OrderCreate(request *Validations.OrderCreate) (err error) {
	if err = Config.DB.Table("orders").Create(request).Error; err != nil {
		return err
	}
	return nil
}

func OrderUpdate(request *Validations.OrderUpdate, userId string) (err error) {
	if err = Config.DB.Table("orders").Where("id = ?", userId).Update(request).Error; err != nil {
		return err
	}
	return nil
}

func OrderDelete(order *Schema.Order, userId string) (err error) {
	if err = Config.DB.Where("id = ?", userId).Delete(order).Error; err != nil {
		return err
	}
	return nil
}

func OrderFetchWithEmail(order *Schema.Order, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(order).Error; err != nil {
		return err
	}
	return nil
}
