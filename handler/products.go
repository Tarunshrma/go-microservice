package handler

import (
	"go-microservice/data"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(writer, request)
		return
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(writer http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(writer)
	if err != nil {
		http.Error(writer, "Error fetching products", http.StatusInternalServerError)
	}
}
