package main

import (
	"github.com/gin-gonic/gin"
	"tgin/bootstrap"
	"tgin/global"
	"net/http"
)

func main() {

	bootstrap.InitializeConfig()

	r := gin.Default()

	// do test route
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	//r.Run(":8088")
	r.Run(":" + global.App.Config.App.Port)

}


