package handler

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Good by Tarun"))
}
