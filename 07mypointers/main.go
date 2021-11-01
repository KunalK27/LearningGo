package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of pointer", ptr)  //accesses memory dress
	fmt.Println("value of pointer", *ptr) //accesses value
	*ptr = *ptr + 2
	fmt.Println("New value is", myNumber)
}
