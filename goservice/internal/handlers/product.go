package handlers

import (
	"log"
	"net/http"
	"github.com/bethecodewithyou/gohttpwebservice/internal/data"
)

//Product is a hander
type Product struct {
	l *log.Logger
}

//NewProduct - creates a product handler with a given logger.
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

//handler is serving httpRequest. Returning product response JSON.
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	productList := data.GetProducts()
	e := productList.ToJSON(rw)
	if e!=nil {
		http.Error(rw, "error while marshalling procut list" , http.StatusInternalServerError)
	}
}	