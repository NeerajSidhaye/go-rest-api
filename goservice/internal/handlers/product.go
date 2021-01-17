package handlers

import (
	"encoding/json"
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
	d, e := json.Marshal(productList)
	if e!=nil {
		http.Error(rw, "error while marshalling procut list" , http.StatusInternalServerError)
	}

	rw.Write(d) // writing json back to the response.

}	