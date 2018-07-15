package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/touhid91/ems/auth"
	"github.com/touhid91/ems/conf"
)


func main() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	api := r.Group("/api/v1")
	{
		authAPI := api.Group("/auth")
		{
			authAPI.POST("/token", auth.TokenHandler)
		}
	}

	env := conf.New()

	r.Run(fmt.Sprintf(":%d", env.Ports.App))
}
