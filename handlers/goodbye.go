package handlers

import (
	"log"
	"net/http"
)

// Goodbye struct
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye creates a Goodbye handler instance with a custom logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Hello handler executed")
	rw.Write([]byte("Goodbye!\n"))
}
