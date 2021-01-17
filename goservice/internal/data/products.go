package data

import (
	"encoding/json"
	"io"

)

//Product : defines the attributes of running shoe
type Product struct {
	Brand      string  `json:"brand"` // struct tags or annotations to fields. This will be shown in the final JSON output.
	Stability  string  `json:"stability"`
	Coushining string  `json:"coushining"`
	Surface    string  `json:"surface"`
	Arch       string  `json:"arch"`
	Size       float32 `json:"size"`
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

		Brand:      "Saucony",
		Stability:  "Netural",
		Coushining: "Moderate",
		Surface:    "Road",
		Arch:       "flat",
		Size:       8.5,
		LaunchDate: "December 2006",
	},
	{

		Brand:      "Altra",
		Stability:  "Netural",
		Coushining: "High",
		Surface:    "Track",
		Arch:       "Medium",
		Size:       9.5,
		LaunchDate: "May 2017",
	},
}
