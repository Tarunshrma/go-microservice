package handler

import (
	"go-microservice/data"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle POST Add Product")
	prod := request.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}
