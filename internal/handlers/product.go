package handlers

import (
	"log"
	"net/http"
	"github.com/bethecodewithyou/restcrud/internal/data"
)

//Product struct wth a logger attribute
type Product struct {
	l *log.Logger
}

//NewProduct - creates a product handler with a given logger.
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

//handler is serving httpRequest. Returning product response JSON.
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// For any other methods, we are returning method not allowed
	rw.WriteHeader(http.StatusMethodNotAllowed)

}	

// private method to Product handler
func (p*Product) getProducts(rw http.ResponseWriter, r *http.Request) {

	productList := data.GetProducts()
	err :=productList.ToJSON(rw)
	if err!=nil {
		http.Error(rw, "error while marshalling procut list" , http.StatusInternalServerError)
	}

}