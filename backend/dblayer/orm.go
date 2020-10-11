package dblayer

import (
	"errors"

	"github.com/hooneun/gomusic/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBORM database orm
type DBORM struct {
	*gorm.DB
}

// ErrINVALIDPASSWORD ,
var ErrINVALIDPASSWORD = errors.New("Invalid password")

// NewORM !
func NewORM() (*DBORM, error) {
	dns := "root@tcp(127.0.0.1:3306)/todo_list?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	return &DBORM{
		DB: db,
	}, err
}

// GetAllProducts !
func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

// GetPromos get promotions
func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NILL").Find(&products).Error
}

// GetCustomerByName !
func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}

// GetCustomerByID !
func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

// GetProduct !
func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

// AddUser create User
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Password)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

// SignInUser signin
func (db *DBORM) SignInUser(email, password string) (customer models.Customer, err error) {

	result := db.Table("Cutomers").Where(&models.Customer{Email: email})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}
	if !checkPassword(customer.Password, password) {
		return customer, ErrINVALIDPASSWORD
	}

	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}

	return customer, result.Find(&customer).Error
}

// SignOutUserByID SinOut!
func (db *DBORM) SignOutUserByID(id int) error {
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	return db.Table("Cutomers").Where(&customer).Update("loggedin", 0).Error
}

// GetCustomerOrdersByID !
func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id = ?", id).Scan(&orders).Error
}

// AddOrder !
func (db *DBORM) AddOrder(order models.Order) error {
	return db.Create(&order).Error
}

// GetCreditCardCID !
func (db *DBORM) GetCreditCardCID(id int) (string, error) {
	customerWithCCID := struct {
		models.Customer
		CCID string `gorm:"column:cc_customerid"`
	}{}
	return customerWithCCID.CCID, db.First(&customerWithCCID).Error
}

//SaveCreditCardForCustomer !
func (db *DBORM) SaveCreditCardForCustomer(id int, ccid string) error {
	result := db.Table("customers").Where("id=?", id)
	return result.Update("cc_customerid", ccid).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	sBytes := []byte(*s)
	hashed, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*s = string(hashed[:])
	return nil
}

func checkPassword(existingHash, incomingPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPassword)) == nil
}
