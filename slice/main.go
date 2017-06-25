package main

import "fmt"

func main() {
	var myslice1 []int16
	myslice1 = make([]int16, 5, 10)
	myslice1[0] = 100
	myslice1[1] = 101
	var myslice2 []int16
	myslice2 = make([]int16, 3, 10)
	myslice2[0] = 200
	myslice2[1] = 201
	fn1(myslice1)
	fn1(myslice2)
	f := append(myslice1, myslice2...)
	fmt.Println(f)
	copy(myslice1, myslice2)
	fmt.Println(myslice1)

}

func fn1(slice []int16) {
	fmt.Println(slice, len(slice), cap(slice))
}
