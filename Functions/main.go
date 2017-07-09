package main


import (
	"fmt"
        "triple"

func main() {
	m := 4
	n := 5
	mult(m, n)
	Tri(2)
}
func mult(a int, b int) (c int) { // func function name (input variables) (output variables)
	c = a * b
	fmt.Println(c)
	return
}
