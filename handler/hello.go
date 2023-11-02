package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) Hello {
	return Hello{l}
}

func (h Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	d, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(writer, "Some Error...", http.StatusBadRequest)
		return
	}

	h.l.Println("Hello from handler..")
	fmt.Fprintf(writer, "%s", d)
}
