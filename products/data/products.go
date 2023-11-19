package data

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"grpc-go/protos/currency/protos"
	"time"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"`

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`

	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Products is a collection of Product
type Products []*Product

type ProductDB struct {
	log      hclog.Logger
	currency protos.CurrencyClient
}

func NewProductDB(l hclog.Logger, cc protos.CurrencyClient) *ProductDB {
	return &ProductDB{l, cc}
}

func (pDB *ProductDB) GetProducts(currency string) (Products, error) {

	if currency == "" {
		return productList, nil
	}

	rate, err := pDB.getRate(currency)
	if err != nil {
		pDB.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	pr := Products{}
	for _, p := range productList {
		np := *p
		np.Price = np.Price * rate
		pr = append(pr, &np)
	}

	return pr, nil

}

func (pDB *ProductDB) getRate(currency string) (float64, error) {
	// get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[currency]),
	}

	resp, err := pDB.currency.GetRate(context.Background(), rr)
	return resp.Rate, err
}

func (pDB *ProductDB) AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func (pDB *ProductDB) UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func DeleteProduct(id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	//productList = append(productList[:pos], productList[pos+1])
	productList = append(productList[:pos], productList[pos+1:]...)

	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	p := productList[len(productList)-1]
	return p.ID + 1
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func (pDB *ProductDB) GetProductByID(id int, currency string) (*Product, error) {
	p, id, _ := findProduct(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	if currency == "" {
		return p, nil
	}

	rate, err := pDB.getRate(currency)
	if err != nil {
		pDB.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	np := *p
	np.Price = np.Price * rate

	return &np, nil
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Coffee with milk",
		Price:       2.5,
		SKU:         "lt123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Black Coffee",
		Price:       1.5,
		SKU:         "es123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
