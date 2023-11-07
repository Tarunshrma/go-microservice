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
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(writer)
	if err != nil {
		http.Error(writer, "Error fetching products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle POST Add Product")
	prod := &data.Product{}
	err := prod.FromJSON(request)

	if err != nil {
		http.Error(writer, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

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

	prod := &data.Product{}

	err := prod.FromJSON(r)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if errors.Is(err, data.ErrProductNotFound) {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
