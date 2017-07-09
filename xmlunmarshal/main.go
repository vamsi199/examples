package main

import (
	"encoding/xml"
	"fmt"
)

func main() {

	type Result struct {
		XMLName xml.Name `xml:"book"`
		id    string   `xml:"id"`
		author   string
		title   string
		price  float32
		Phone string
		description string
		publish_date complex64

	}
	var v Result

	data := `
		<book id="bk">
      <author>Gambardella, Matthew</author>
      <title>XML Developer's Guide</title>
      <genre>Computer</genre>
      <price>44.95</price>
      <publish_date>2000-10-01</publish_date>
      <description>An in-depth look at creating applications
      with XML.</description>
   </book>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("id: %v\n", v.id)
	fmt.Printf("Phone: %v\n", v.Phone)
	fmt.Printf("price: %v\n", v.price)
	fmt.Printf("author: %q\n", v.author)
	fmt.Printf("description: %q\n", v.description)
	fmt.Printf("title: %v\n", v.title)


}

