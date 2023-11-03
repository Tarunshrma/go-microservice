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

// NewHello The reason for using *Hello in return and &Hello is that it return the referene of same object instead of creating new everytime
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	d, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(writer, "Some Error...", http.StatusBadRequest)
		return
	}

	h.l.Println("Hello from handler..")
	fmt.Fprintf(writer, "%s", d)
}
