package Routes

import (
	"go-gin-mysql-boilerplate/Controllers"
	"go-gin-mysql-boilerplate/Middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	route.Use(cors.New(config))

	route.POST("auth/signin", Middlewares.IsClientAuthenticated, Controllers.AuthSignin)
	route.DELETE("auth/signout", Middlewares.IsUserAuthenticated, Controllers.AuthSignout)

	route.GET("users", Middlewares.IsUserAuthenticated, Controllers.UserFetchAll)
	route.GET("users/:id", Middlewares.IsUserAuthenticated, Controllers.UserFetchSingle)
	route.POST("users", Middlewares.IsClientAuthenticated, Controllers.UserCreate)
	route.PUT("users/:id", Middlewares.IsUserAuthenticated, Controllers.UserUpdate)
	route.DELETE("users/:id", Middlewares.IsUserAuthenticated, Controllers.UserDelete)
	route.GET("products", Middlewares.IsUserAuthenticated, Controllers.ProductFetchAll)
	route.GET("products/:id", Middlewares.IsUserAuthenticated, Controllers.ProductFetchSingle)
	route.POST("products", Middlewares.IsClientAuthenticated, Controllers.ProductCreate)
	route.PUT("products/:id", Middlewares.IsUserAuthenticated, Controllers.ProductUpdate)
	route.DELETE("products/:id", Middlewares.IsUserAuthenticated, Controllers.ProductDelete)
	return route
}
