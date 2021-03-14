package data

import (
	"encoding/json"
	"fmt"
	"io"
	"github.com/jinzhu/copier"
	
)

//Product : defines the attributes of Shoe product
type Product struct {
	ID         int     `json:"id"` // struct tags or annotations to fields. This will be shown in the final JSON output.
	Sport      string  `json:"sport"`
	Brand      string  `json:"brand"`
	Colour     string  `json:"colour"`
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

//UpdateProductAttribute : Handing Partial update on Product attribute
func UpdateProductAttribute(id int, p*Product) error {

	pos, err := getProductPosition(id)
	if err!=nil {
		return err
	}

	copier.CopyWithOption(productList[pos], p, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	
	return nil
}

//DeleteProductByID : 
func DeleteProductByID(id int) error {

	pos, err := getProductPosition(id)
	if err != nil {
		return err
	}

	productList = append(productList[:pos], productList[pos+1:]...)

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

// example data source - creating hard coded list of product for CRUD oprations purpose.
var productList = []*Product{

	{

		ID:         1,
		Sport:      "Running",
		Brand:      "Altra",
		Colour:     "Blue",
		LaunchDate: "Dec-2006",
	},
	{

		ID:         2,
		Sport:      "Hiking",
		Brand:      "NorthFace",	
		Colour:     "Green",
		LaunchDate: "Jan-2020",
	},
}
