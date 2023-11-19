package handler

import (
	"context"
	"go-microservice/data"
	"net/http"
)

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Error("deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Error("validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
