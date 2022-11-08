package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "成功",
		})
	})

	router.POST("/register", func(context *gin.Context) {
		//username password user_id
	})
	router.Run()

}
