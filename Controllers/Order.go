package Controllers

import (
	"go-gin-mysql-boilerplate/Models"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Services"
	"go-gin-mysql-boilerplate/Validations"
	"strings"

	"github.com/gin-gonic/gin"
)

func OrderFetchAll(c *gin.Context) {
	var order []Schema.Order
	err := Models.OrderFetchAll(&order)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, nil, order)
	}
}

func OrderFetchSingle(c *gin.Context) {
	userId := c.Param("id")

	var order Schema.Order
	err := Models.OrderFetchSingle(&order, userId)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, nil, order)
	}
}

func OrderCreate(c *gin.Context) {
	var request Validations.OrderCreate
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	// request.Password = Services.MD5Hash(request.Password)
	saveErr := Models.OrderCreate(&request)
	if saveErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", saveErr)
	} else {
		Services.Success(c, "Created", request)
	}
}

func OrderUpdate(c *gin.Context) {
	userId := c.Param("id")

	var request Validations.OrderUpdate
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	updateErr := Models.OrderUpdate(&request, userId)
	if updateErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", updateErr)
	} else {
		Services.Success(c, "Updated", request)
	}
}

func OrderDelete(c *gin.Context) {
	userId := c.Param("id")

	var order Schema.Order
	err := Models.OrderDelete(&order, userId)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, "Deleted", nil)
	}
}
