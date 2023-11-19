package handler

import (
	"go-microservice/data"
	"net/http"
)

// swagger:route GET /products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetProducts ListAll handles GET requests and returns all current products
func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	p.l.Debug("Get all products")
	writer.Header().Add("Content-Type", "application/json")

	cur := request.URL.Query().Get("currency")

	lp, err := p.productDB.GetProducts(cur)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, writer)
		return
	}

	err = data.ToJSON(lp, writer)
	if err != nil {
		http.Error(writer, "Error fetching products", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// GetProduct handles GET requests
func (p *Products) GetProduct(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(request)

	p.l.Debug("Get record id", id)

	cur := request.URL.Query().Get("currency")
	prod, err := p.productDB.GetProductByID(id, cur)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Error("fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Error("fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Error("serializing product", err)
	}
}
