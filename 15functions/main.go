package main

import "fmt"

func main() {
	fmt.Println("Hello Functions")
	greeter()
	greeter2()
	result := adder(3, 5)
	fmt.Println("Result is: ", result)
	proresult, myMessage := proadder(1, 2, 492, 192)
	fmt.Println("Pro Result is", proresult)
	fmt.Println("Pro Message is", myMessage)
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func proadder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro Reslt Func"
}
func greeter2() {
	fmt.Println("Second func")
}
func greeter() {
	fmt.Println("Namaste from Golang")
}
