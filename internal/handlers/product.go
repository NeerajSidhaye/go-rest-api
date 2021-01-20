package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/bethecodewithyou/product/internal/data"
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
		p.l.Println("handling GET")
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.l.Println("handling POST")
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		
		p.l.Println("handling PUT")
	    
	   regex := regexp.MustCompile(`/([0-9]+)`)
	   g := regex.FindAllStringSubmatch(r.URL.Path,  -1)
	   p.l.Printf("regex g group %q\n", g)

	   if len(g)!=1 {  // if true, means there are more than one id passed in the URI.
			p.l.Println("invalid product id in the URI")
			http.Error(rw, "Invalid Request URI", http.StatusBadRequest)
			return
	   }

	   if len(g[0]) > 2 {
		 http.Error(rw, "Invalid URI", http.StatusBadRequest)
		 return
	   }
	   
		productID := g[0][1]
		idString, err := strconv.Atoi(productID)
	     if err!=nil {
			 http.Error(rw, "Inavalid id value in URI", http.StatusBadRequest)
			 return
		 }

		p.l.Println("updating product for ID ->", idString)
		p.updateProduct(idString, rw, r)
		return

	}

	// For any other methods, we are returning method not allowed
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

// private method to Product handler
func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {

	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "error while marshalling procut list", http.StatusInternalServerError)
	}
}

// private method - this will add a new product coming from POST request into existing list of products.
func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {

	newProduct := &data.Product{} // this prod will have address of Product struct

	err := newProduct.FromJSONtoProduct(r.Body)

	if err != nil {
		http.Error(rw, "error while adding new product", http.StatusInternalServerError)
	}

	data.AddProduct(newProduct)

}

func (p*Product) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {

	prod := &data.Product{}
	
	err := prod.FromJSONtoProduct(r.Body)
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
