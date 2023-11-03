package handler

import (
	"encoding/json"
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
	lp := data.GetProducts()
	res, err := json.Marshal(lp)
	if err != nil {
		http.Error(writer, "Error fetching products", http.StatusInternalServerError)
	}
	writer.Write(res)
}
