package handler

import (
	"go-microservice/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := data.ToJSON(lp, writer)
	if err != nil {
		http.Error(writer, "Error fetching products", http.StatusInternalServerError)
	}
}
