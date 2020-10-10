package rest

import (
	"github.com/gin-gonic/gin"
)

// RunAPI API !
func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}

	return RunAPIWithHandler(address, h)
}

// RunAPIWithHandler API List!
func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()
	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetProducts)
	// r.POST("/users/signin", h.SiginIn)
	// r.POST("/users", h.AddUser)
	// r.POST("/user/:id/signout", h.SignOut)
	// r.GET("/user/:id/orders", h.GetOrders)
	// r.POST("/users/charge", h.Charge)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SiginIn)
		usersGroup.POST("", h.AddUser)
	}

	return r.Run(address)
}
