package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
	//first non defers print and then for defer last in first out
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
