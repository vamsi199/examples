package main

import "fmt"

func main() {
	m := 4
	n := 5
	mult(m, n)
}
func mult(a int, b int) (c int) { // func function name (input variables) (output variables)
	c = a * b
	fmt.Println(c)
	return
}
