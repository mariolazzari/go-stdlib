package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func main() {

	// marshal
	country := Country{
		Name:    "Italia",
		Capital: "Roma",
	}
	jsonData, err := json.Marshal(country)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshal", jsonData)

	// unmarshal
	var decodedCountry Country
	err = json.Unmarshal(jsonData, &decodedCountry)
	fmt.Println("Unmarshal", decodedCountry)

	// xml
	xmlData, err := xml.Marshal(country)
	fmt.Println("Marshal", xmlData)

}

type Country struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Population int    `json:"population,omitempty"`
}

type CounttyXml struct {
	XMLName xml.Name `xml:"country"`
	Name    string   `xml:"name"`
	Capital string   `xml:"capital"`
}
