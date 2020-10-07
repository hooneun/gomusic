package rest

import (
	"github.com/gin-gonic/gin"
)

// RunAPI API !
func RunAPI(address string) error {
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {

	})
	r.GET("/promos", func(c *gin.Context) {

	})
	r.POST("/users/signin", func(c *gin.Context) {

	})
	r.POST("/users", func(c *gin.Context) {

	})
	r.POST("/user/:id/signout", func(c *gin.Context) {

	})
	r.GET("/user/:id/orders", func(c *gin.Context) {

	})
	r.POST("/users/charge", func(c *gin.Context) {

	})
}
