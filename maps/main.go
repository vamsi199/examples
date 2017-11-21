package main

import ("fmt"
)

type data struct {
	username,email,contact map[string]string
}

func main() {
	var d data
	d.username = make(map[string]string)
	d.username["username"] = "vamsi"
	//d.email["email"]="vamsi@gmail.com"
	//d.contact["contact"]="phone number"
	fmt.Println(d)

}
