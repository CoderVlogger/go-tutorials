package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	mdb := NewMemDb()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, mdb.ListUsers())
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusCreated, mdb.AddUser(&user))
		}
	})

	r.GET("/users/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.JSON(http.StatusOK, mdb.GetUser(username))
	})

	r.DELETE("/users/:username", func(c *gin.Context) {
		username := c.Param("username")
		mdb.DeleteUser(username)
		c.JSON(http.StatusOK, gin.H{})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
