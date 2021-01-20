package data

import (
	"encoding/json"
	"fmt"
	"io"
)

//Product : defines the attributes of shoe
type Product struct {
	ID         int     `json:"id"` // struct tags or annotations to fields. This will be shown in the final JSON output.
	Sport      string  `json:"sport"`
	Type       string  `json:"type"`
	Brand      string  `json:"brand"`
	Colour     string  `json:"colour"`
	Terrain    string  `json:"terrain"`
	Feature    string  `json:"feature"`
	Size       float32 `json:"size"`
	Price      string  `json:"price"`
	LaunchDate string  `json:"-"` // fields which has struct tag with dash ( - ), won't be added to the resulsting JSON.
}

//Products : is a collection of products OR slice of product
type Products []*Product

//GetProducts : returns list of all running shoes
func GetProducts() Products {
	return productList
}

//ToJSON : serialize collection of products to JSON
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

//FromJSONtoProduct : deserialize incoming json to as new product.
func (p *Product) FromJSONtoProduct(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

//AddProduct : adding new products to an existing product list
func AddProduct(p*Product) {
	p.ID = nextProductID()
	productList = append(productList, p)
}

//UpdateProduct : updating existing product entry.
func UpdateProduct(id int, p*Product) error {
	
	pos, err := getProductPosition(id)
	
	if err !=nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

//ErrProductNotFound : global error variable
var ErrProductNotFound = fmt.Errorf("product not found")

func getProductPosition(id int) (int, error) {

	for index, p:= range productList {
		if p.ID == id {
			return index, nil
		}
	}
	return -1, ErrProductNotFound

}

func nextProductID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

// example data source - creating hard coded list of shoes for CRUD oprations purpose.
var productList = []*Product{

	{

		ID:         1,
		Sport:      "Running",
		Type:       "Netural",
		Brand:      "Saucony",
		Colour:     "Blue",
		Terrain:    "Road",
		Feature:    "Lightweight",
		Size:       8.5,
		Price:      "£90",
		LaunchDate: "Dec-2006",
	},
	{

		ID:         2,
		Sport:      "Running",
		Type:       "Trail",
		Brand:      "Altra",	
		Colour:     "Green",
		Terrain:    "Trail",
		Feature:    "Breathable",
		Size:       9.5,
		Price:      "£110",
		LaunchDate: "Jan-2020",
	},
}
