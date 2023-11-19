package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	mux2 "github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"go-microservice/data"
	"go-microservice/handler"
	"google.golang.org/grpc"
	"grpc-go/protos/currency/protos"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	l := hclog.Default()
	v := data.NewValidation()

	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// create client
	cc := protos.NewCurrencyClient(conn)

	//Create productdb instance
	pdb := data.NewProductDB(l, cc)

	ph := handler.NewProducts(l, v, pdb)

	mux := mux2.NewRouter()

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/products", ph.GetProducts).Queries("currency", "{[A-Z]{3}}")

	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct).Queries("currency", "{[A-Z]{3}}")

	putRouter := mux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("products/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("products/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Addr:         ":9090",
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		Handler:      ch(mux),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Define a channel to signal when it's time to shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		l.Debug("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error: ", err)
		}
	}()

	// Block until a shutdown signal is received
	<-stop

	// Create a context with a timeout to force server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %v\n", err)
	}
}
