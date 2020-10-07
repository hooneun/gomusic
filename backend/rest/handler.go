package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/hooneun/gomusic/backend/dblayer"
)

// HandlerInterface ! Handler All func
type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SiginIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

// Handler read update dblayer access
type Handler struct {
	db dblayer.DBLayer
}

// NewHandler create construct
func NewHandler() (*Handler, error) {
	return new(Handler), nil
}
