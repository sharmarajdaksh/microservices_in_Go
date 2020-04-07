package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello struct
type Hello struct {
	l *log.Logger
}

// NewHello creates a Hello handler instance with a custom logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello handler executed")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Hello %s\n", b)
}

// The http.Handler interface is a simple interface which has only ONE method on it: ServeHTTP(ResponseWriter, *Request)

// For tetstability, we won't use a global logger. Instead, we'll define a field on our Hello handler
