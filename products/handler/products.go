package handler

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"go-microservice/data"
	"net/http"
	"strconv"
)

type Products struct {
	l         hclog.Logger
	v         *data.Validation
	productDB *data.ProductDB
}

func NewProducts(l hclog.Logger, v *data.Validation, pdb *data.ProductDB) *Products {
	return &Products{l, v, pdb}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

type KeyProduct struct{}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
