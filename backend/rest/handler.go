package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hooneun/gomusic/backend/dblayer"
	"github.com/hooneun/gomusic/backend/models"
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
	db, err := dblayer.NewORM()
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
	// return new(Handler), nil
}

// GetProducts return products
func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return
	}

	products, err := h.db.GetAllProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetPromos return promotions
func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {
		return
	}

	promos, err := h.db.GetPromos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, promos)
}

// SiginIn User sigin
func (h *Handler) SiginIn(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.SignInUser(customer.Email, customer.Password)
	if err != nil {
		if err == dblayer.ErrINVALIDPASSWORD {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// AddUser user add
func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		return
	}

	var customer models.Customer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// SignOut user logginout
func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.db.SignOutUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// GetOrders orders all get
func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orders, err := h.db.GetCustomerOrdersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Charge Product Payment!!!!!!
func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}

	// 구조체 선언 동시에 초기화
	request := struct {
		models.Order
		Remember   bool `json:"rememberCard"`
		UseExiting bool `json:"useExisting"`
	}{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, request)
		return
	}
}

// func customMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Set("v", "123")
// 		c.Next()
// 		status := c.Writer.Status()
// 	}
// }

// MyCustomLogger !
func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("********************")
		c.Next()
		fmt.Println("********************")
	}
}
