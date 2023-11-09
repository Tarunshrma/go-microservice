package handler

import (
	"errors"
	"github.com/gorilla/mux"
	"go-microservice/data"
	"net/http"
	"strconv"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.Atoi(vars["id"])

	if e != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err := data.UpdateProduct(id, prod)
	if errors.Is(err, data.ErrProductNotFound) {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
