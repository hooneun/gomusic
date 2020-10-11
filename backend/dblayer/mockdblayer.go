package dblayer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hooneun/gomusic/backend/models"
)

// MockDBLayer Test db layer
type MockDBLayer struct {
	err       error
	products  []models.Product
	customers []models.Customer
	orders    []models.Order
}

// NewMockDBLayer !
func NewMockDBLayer(products []models.Product, customers []models.Customer, orders []models.Order) *MockDBLayer {
	return &MockDBLayer{
		products:  products,
		customers: customers,
		orders:    orders,
	}
}

// NewMockDBLayerWithData Mock data
func NewMockDBLayerWithData() *MockDBLayer {
	PRODUCTS := `[
			{
					"ID": 1,
					"CreatedAt": "2018-08-14T07:54:19Z",
					"UpdatedAt": "2019-01-11T00:28:40Z",
					"DeletedAt": null,
					"img": "img/strings.png",
					"small_img": "img/img-small/strings.png",
					"imgalt": "string",
					"price": 100,
					"promotion": 0,
					"productname": "Strings",
					"Description": ""
			},
			{
					"ID": 2,
					"CreatedAt": "2018-08-14T07:54:20Z",
					"UpdatedAt": "2019-01-11T00:29:11Z",
					"DeletedAt": null,
					"img": "img/redguitar.jpeg",
					"small_img": "img/img-small/redguitar.jpeg",
					"imgalt": "redg",
					"price": 299,
					"promotion": 240,
					"productname": "Red Guitar",
					"Description": ""
			},
			{
					"ID": 3,
					"CreatedAt": "2018-08-14T07:54:20Z",
					"UpdatedAt": "2019-01-11T22:05:42Z",
					"DeletedAt": null,
					"img": "img/drums.jpg",
					"small_img": "img/img-small/drums.jpg",
					"imgalt": "drums",
					"price": 17000,
					"promotion": 0,
					"productname": "Drums",
					"Description": ""
			},
			{
					"ID": 4,
					"CreatedAt": "2018-08-14T07:54:20Z",
					"UpdatedAt": "2019-01-11T00:29:53Z",
					"DeletedAt": null,
					"img": "img/flute.jpeg",
					"small_img": "img/img-small/flute.jpeg",
					"imgalt": "flute",
					"price": 210,
					"promotion": 190,
					"productname": "Flute",
					"Description": ""
			},
			{
					"ID": 5,
					"CreatedAt": "2018-08-14T07:54:20Z",
					"UpdatedAt": "2019-01-11T00:30:12Z",
					"DeletedAt": null,
					"img": "img/blackguitar.jpeg",
					"small_img": "img/img-small/blackguitar.jpeg",
					"imgalt": "Black guitar",
					"price": 200,
					"promotion": 0,
					"productname": "Black Guitar",
					"Description": ""
			},
			{
					"ID": 6,
					"CreatedAt": "2018-08-14T07:54:20Z",
					"UpdatedAt": "2019-01-11T00:30:35Z",
					"DeletedAt": null,
					"img": "img/saxophone.jpeg",
					"small_img": "img/img-small/saxophone.jpeg",
					"imgalt": "Saxophone",
					"price": 1000,
					"promotion": 980,
					"productname": "Saxophone",
					"Description": ""
			}
	]
	`

	ORDERS := `[
		{
				"ID": 1,
				"CreatedAt": "2018-12-29T23:35:36Z",
				"UpdatedAt": "2018-12-29T23:35:36Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 1,
				"product_id": 1,
				"sell_price": 90,
				"purchase_date": "2018-12-29T23:34:32Z"
		},
		{
				"ID": 2,
				"CreatedAt": "2018-12-29T23:35:48Z",
				"UpdatedAt": "2018-12-29T23:35:48Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 1,
				"product_id": 2,
				"sell_price": 299,
				"purchase_date": "2018-12-29T23:34:53Z"
		},
		{
				"ID": 3,
				"CreatedAt": "2018-12-29T23:35:57Z",
				"UpdatedAt": "2018-12-29T23:35:57Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 1,
				"product_id": 3,
				"sell_price": 16000,
				"purchase_date": "2018-12-29T23:35:05Z"
		},
		{
				"ID": 4,
				"CreatedAt": "2018-12-29T23:36:18Z",
				"UpdatedAt": "2018-12-29T23:36:18Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 2,
				"product_id": 1,
				"sell_price": 95,
				"purchase_date": "2018-12-29T23:36:18Z"
		},
		{
				"ID": 5,
				"CreatedAt": "2018-12-29T23:36:39Z",
				"UpdatedAt": "2018-12-29T23:36:39Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 2,
				"product_id": 2,
				"sell_price": 299,
				"purchase_date": "2018-12-29T23:36:39Z"
		},
		{
				"ID": 6,
				"CreatedAt": "2018-12-29T23:38:13Z",
				"UpdatedAt": "2018-12-29T23:38:13Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 2,
				"product_id": 4,
				"sell_price": 205,
				"purchase_date": "2018-12-29T23:37:01Z"
		},
		{
				"ID": 7,
				"CreatedAt": "2018-12-29T23:38:19Z",
				"UpdatedAt": "2018-12-29T23:38:19Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 3,
				"product_id": 4,
				"sell_price": 210,
				"purchase_date": "2018-12-29T23:37:28Z"
		},
		{
				"ID": 8,
				"CreatedAt": "2018-12-29T23:38:28Z",
				"UpdatedAt": "2018-12-29T23:38:28Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 3,
				"product_id": 5,
				"sell_price": 200,
				"purchase_date": "2018-12-29T23:37:41Z"
		},
		{
				"ID": 9,
				"CreatedAt": "2018-12-29T23:38:32Z",
				"UpdatedAt": "2018-12-29T23:38:32Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 3,
				"product_id": 6,
				"sell_price": 1000,
				"purchase_date": "2018-12-29T23:37:54Z"
		},
		{
				"ID": 10,
				"CreatedAt": "2019-01-13T00:44:55Z",
				"UpdatedAt": "2019-01-13T00:44:55Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 19,
				"product_id": 6,
				"sell_price": 1000,
				"purchase_date": "2018-12-29T23:37:54Z"
		},
		{
				"ID": 11,
				"CreatedAt": "2019-01-14T06:03:08Z",
				"UpdatedAt": "2019-01-14T06:03:08Z",
				"DeletedAt": null,
				"img": "",
				"small_img": "",
				"imgalt": "",
				"price": 0,
				"promotion": 0,
				"productname": "",
				"Description": "",
				"name": "",
				"firstname": "",
				"lastname": "",
				"email": "",
				"password": "",
				"loggedin": false,
				"orders": null,
				"customer_id": 1,
				"product_id": 3,
				"sell_price": 17000,
				"purchase_date": "0001-01-01T00:00:00Z"
		}
	]
	`
	CUSTOMERS := `[
		{
				"ID": 1,
				"CreatedAt": "2018-08-14T07:52:54Z",
				"UpdatedAt": "2019-01-13T22:00:45Z",
				"DeletedAt": null,
				"name": "",
				"firstname": "Mal",
				"lastname": "Zein",
				"email": "mal.zein@email.com",
				"password": "$2a$10$ZeZI4pPPlQg89zfOOyQmiuKW9Z7pO9/KvG7OfdgjPAZF0Vz9D8fhC",
				"loggedin": true,
				"orders": null
		},
		{
				"ID": 2,
				"CreatedAt": "2018-08-14T07:52:55Z",
				"UpdatedAt": "2019-01-12T22:39:01Z",
				"DeletedAt": null,
				"name": "",
				"firstname": "River",
				"lastname": "Sam",
				"email": "river.sam@email.com",
				"password": "$2a$10$mNbCLmfCAc0.4crDg3V3fe0iO1yr03aRfE7Rr3vdfKMGVnnzovCZq",
				"loggedin": false,
				"orders": null
		},
		{
				"ID": 3,
				"CreatedAt": "2018-08-14T07:52:55Z",
				"UpdatedAt": "2019-01-13T21:56:05Z",
				"DeletedAt": null,
				"name": "",
				"firstname": "Jayne",
				"lastname": "Ra",
				"email": "jayne.ra@email.com",
				"password": "$2a$10$ZeZI4pPPlQg89zfOOyQmiuKW9Z7pO9/KvG7OfdgjPAZF0Vz9D8fhC",
				"loggedin": false,
				"orders": null
		},
		{
				"ID": 19,
				"CreatedAt": "2019-01-13T08:43:44Z",
				"UpdatedAt": "2019-01-13T15:12:25Z",
				"DeletedAt": null,
				"name": "",
				"firstname": "John",
				"lastname": "Doe",
				"email": "john.doe@bla.com",
				"password": "$2a$10$T4c8rmpbgKrUA0sIqtHCaO0g2XGWWxFY4IGWkkpVQOD/iuBrwKrZu",
				"loggedin": false,
				"orders": null
		}
	]
	`

	var products []models.Product
	var customers []models.Customer
	var orders []models.Order
	json.Unmarshal([]byte(PRODUCTS), &products)
	json.Unmarshal([]byte(CUSTOMERS), &customers)
	json.Unmarshal([]byte(ORDERS), &orders)

	return NewMockDBLayer(products, customers, orders)
}

// GetMockProductData !
func (mock *MockDBLayer) GetMockProductData() []models.Product {
	return mock.products
}

// GetMockCustomersData !
func (mock *MockDBLayer) GetMockCustomersData() []models.Customer {
	return mock.customers
}

// GetMockOrdersData !
func (mock *MockDBLayer) GetMockOrdersData() []models.Order {
	return mock.orders
}

// SetError Setting Error
func (mock *MockDBLayer) SetError(err error) {
	mock.err = err
}

// GetAllProducts !
func (mock *MockDBLayer) GetAllProducts() ([]models.Product, error) {
	if mock.err != nil {
		return nil, mock.err
	}

	return mock.products, nil
}

// GetPromos !
func (mock *MockDBLayer) GetPromos() ([]models.Product, error) {
	if mock.err != nil {
		return nil, mock.err
	}

	promos := []models.Product{}
	for _, product := range mock.products {
		if product.Promotion > 0 {
			promos = append(promos, product)
		}
	}

	return promos, nil
}

// GetProduct !
func (mock *MockDBLayer) GetProduct(id int) (models.Product, error) {
	result := models.Product{}
	if mock.err != nil {
		return result, mock.err
	}

	for _, product := range mock.products {
		if product.ID == uint(id) {
			return product, nil
		}
	}

	return result, fmt.Errorf("Could not find product with id %d", id)
}

// GetCustomerByName !
func (mock *MockDBLayer) GetCustomerByName(first, last string) (models.Customer, error) {
	result := models.Customer{}
	if mock.err != nil {
		return result, mock.err
	}

	for _, customer := range mock.customers {
		if strings.EqualFold(customer.FirstName, first) && strings.EqualFold(customer.LastName, last) {
			return customer, nil
		}
	}

	return result, fmt.Errorf("Could not find user %s %s", first, last)
}

// AddUser !
func (mock *MockDBLayer) AddUser(customer models.Customer) (models.Customer, error) {
	if mock.err != nil {
		return models.Customer{}, mock.err
	}

	mock.customers = append(mock.customers, customer)
	return customer, nil
}

// SignInUser !
func (mock *MockDBLayer) SignInUser(email, password string) (models.Customer, error) {
	if mock.err != nil {
		return models.Customer{}, mock.err
	}
	for _, customer := range mock.customers {
		if strings.EqualFold(email, customer.Email) && customer.Password == password {
			customer.LoggedIn = true
			return customer, nil
		}
	}

	return models.Customer{}, fmt.Errorf("Could not sign in user %s", email)
}

// SignOutUserByID !
func (mock *MockDBLayer) SignOutUserByID(id int) error {
	if mock.err != nil {
		return mock.err
	}
	for _, customer := range mock.customers {
		if customer.ID == uint(id) {
			customer.LoggedIn = false
			return nil
		}
	}
	return fmt.Errorf("Cluld not sign out user %d", id)
}

// GetCustomerOrdersByID !
func (mock *MockDBLayer) GetCustomerOrdersByID(id int) ([]models.Order, error) {
	if mock.err != nil {
		return nil, mock.err
	}
	for _, customer := range mock.customers {
		if customer.ID == uint(id) {
			return customer.Orders, nil
		}
	}
	return nil, fmt.Errorf("Cluld not find customer id %d", id)
}

// AddOrder !
func (mock *MockDBLayer) AddOrder(order models.Order) error {
	if mock.err != nil {
		return mock.err
	}
	for _, customer := range mock.customers {
		if customer.ID == uint(order.CustomerID) {
			customer.Orders = append(customer.Orders, order)
			return nil
		}
	}
	return fmt.Errorf("Could not find customer id %d for order", order.CustomerID)
}

// GetCreditCardID !
func (mock *MockDBLayer) GetCreditCardID(id int) (string, error) {
	if mock.err != nil {
		return "", mock.err
	}
	return "", nil
}

// SaveCreditCardForCustomer !
func (mock *MockDBLayer) SaveCreditCardForCustomer(int, string) error {
	if mock.err != nil {
		return mock.err
	}
	return nil
}
