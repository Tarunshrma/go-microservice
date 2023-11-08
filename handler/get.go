package handler

import (
	"go-microservice/data"
	"net/http"
)

func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := data.ToJSON(lp, writer)
	if err != nil {
		http.Error(writer, "Error fetching products", http.StatusInternalServerError)
	}
}
