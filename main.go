package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		d, err := io.ReadAll(request.Body)

		if err != nil {
			http.Error(writer, "Some Error...", http.StatusBadRequest)
			return
		}

		log.Println("Hello Tarun")
		fmt.Fprintf(writer, "%s", d)
	})

	http.HandleFunc("/sanchit", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello Sanchit")
	})

	http.ListenAndServe(":9090", nil)
}
