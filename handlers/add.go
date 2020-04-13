package handlers

import (
	"net/http"

	"github.com/sharmarajdaksh/microservices_in_Go/data"
)

// AddProduct adds a product on a POST request
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
