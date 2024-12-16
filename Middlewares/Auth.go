package Middlewares

import (
	"go-gin-mysql-boilerplate/Config"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Services"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func IsClientAuthenticated(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	c.Next()
	return
	if hasAuth {
		var clientUser Schema.OAuthClient
		if err := Config.DB.
			Raw("select * from o_auth_clients where name = ? and secret = ?", user, password).
			Scan(&clientUser).
			Error; err != nil {
			Services.Unauthorized(c, "Authentication Failed!", nil)
			return
		}
		c.Next()
	} else {
		Services.Unauthorized(c, "Authentication Failed!", nil)
	}
}

func IsUserAuthenticated(c *gin.Context) {
	c.Next()
	return
	if c.Request.Header.Get("Authorization") == "" {
		Services.Unauthorized(c, "Authentication Failed!", nil)
		return
	}

	req := c.Request.Header.Get("Authorization")
	var access Schema.OAuthAccessToken
	if err := Config.DB.
		Raw(
			"select * from o_auth_access_tokens where access_token = ? and expired_at > ? and revoked = ?",
			strings.Split(req, " ")[1],
			time.Now().Format("2006-01-02 15:04:05"),
			false,
		).
		Scan(&access).
		Error; err != nil {
		Services.Unauthorized(c, "Invalid Token!", nil)
		return
	}
	c.Next()
}
