package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main()  {
	r := gin.Default()

	r.GET("/me", func (c *gin.Context)  {
		c.String(http.StatusOK, "I am David")
	})

	r.GET("/hello", getHello)
	r.GET("/:name", getParameter)

	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "david" || json.Password != "terralogic" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.Run(":8080")
}

func getHello(c *gin.Context)  {
	c.String(http.StatusOK, "Hello")
}

func getParameter(c *gin.Context)  {
	message := "Hello " + c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}