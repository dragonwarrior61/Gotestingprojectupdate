package Controllers

import (
	"go-gin-mysql-boilerplate/Models"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Services"
	"go-gin-mysql-boilerplate/Validations"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProductFetchAll(c *gin.Context) {
	var product []Schema.Product
	err := Models.ProductFetchAll(&product)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, nil, product)
	}
}

func ProductFetchSingle(c *gin.Context) {
	userId := c.Param("id")

	var product Schema.Product
	err := Models.ProductFetchSingle(&product, userId)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, nil, product)
	}
}

func ProductCreate(c *gin.Context) {
	var request Validations.ProductCreate
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	// request.Password = Services.MD5Hash(request.Password)
	saveErr := Models.ProductCreate(&request)
	if saveErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", saveErr)
	} else {
		Services.Success(c, "Created", request)
	}
}

func ProductUpdate(c *gin.Context) {
	userId := c.Param("id")

	var request Validations.ProductUpdate
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	updateErr := Models.ProductUpdate(&request, userId)
	if updateErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", updateErr)
	} else {
		Services.Success(c, "Updated", request)
	}
}

func ProductDelete(c *gin.Context) {
	userId := c.Param("id")

	var product Schema.Product
	err := Models.ProductDelete(&product, userId)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, "Deleted", nil)
	}
}
