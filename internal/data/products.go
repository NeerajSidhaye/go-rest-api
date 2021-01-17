package data

import (
	"encoding/json"
	"io"
)

//Product : defines the attributes of shoe
type Product struct {
	ID         string  `json:"id"` // struct tags or annotations to fields. This will be shown in the final JSON output.
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

//Products : is a collection of running shoe
type Products []*Product

//GetProducts : returns list of all running shoes
func GetProducts() Products {
	return runningShoeList
}

//ToJSON : serializes collection of runningshoes to JSON
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// example data source - creating hard coded list of shoes for CRUD oprations purpose.
var runningShoeList = []*Product{

	{

		ID:         "1111",
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

		ID:         "2222",
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
