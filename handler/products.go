package handler

import (
	"errors"
	"github.com/gorilla/mux"
	"go-microservice/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
	v *data.Validation
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := data.ToJSON(lp, writer)
	if err != nil {
		http.Error(writer, "Error fetching products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle POST Add Product")
	prod := request.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}

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

type KeyProduct struct{}
