package handler

import (
	"errors"
	"go-microservice/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(writer, request)
		return
	}

	if request.Method == http.MethodPost {
		p.addProduct(writer, request)
		return
	}

	if request.Method == http.MethodPut {
		p.l.Println("PUT", request.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(request.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(writer, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than two capture group")
			http.Error(writer, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(writer, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, writer, request)
		return
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(writer http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(writer)
	if err != nil {
		http.Error(writer, "Error fetching products", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle POST Add Product")
	prod := &data.Product{}
	err := prod.FromJSON(request)

	if err != nil {
		http.Error(writer, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
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
