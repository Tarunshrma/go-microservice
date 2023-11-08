package handler

import (
	"go-microservice/data"
	"net/http"
)

func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle POST Add Product")
	prod := request.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}
