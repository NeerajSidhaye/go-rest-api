package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bethecodewithyou/gorest/gorilla/internal/data"
	"github.com/gorilla/mux"
)

//Product struct wth a logger attribute
type Product struct {
	l *log.Logger
}

//NewProduct - creates a product handler with a given logger.
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

//GetProducts : get list of all products
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {

	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "error while marshalling procut list", http.StatusInternalServerError)
	}
}

//AddProduct : this will add a new product coming from POST request into existing list of products.
func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {

	newProduct := &data.Product{} // this prod will have address of Product struct

	err := newProduct.FromJSONtoProduct(r.Body)

	if err != nil {
		http.Error(rw, "error while adding new product", http.StatusInternalServerError)
	}

	data.AddProduct(newProduct)

}

//UpdateProduct : updating a product
func (p*Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	uriParams := mux.Vars(r)
	id, err := strconv.Atoi(uriParams["id"])

	if err!=nil {
		http.Error(rw, "unable to parse uri parma product id", http.StatusBadRequest)
		return
	}

	prod := &data.Product{}
	
	err = prod.FromJSONtoProduct(r.Body)
	if err!=nil {
		p.l.Println("unable to deserialize input json", err)
		http.Error(rw, "invalid request", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err!=nil {
		http.Error(rw, "Internal server Error while updating product", http.StatusInternalServerError)
	}

}
