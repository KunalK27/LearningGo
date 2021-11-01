package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")
	logincount := 10
	var result string
	if logincount < 10 {
		result = "Regular user"
	} else if logincount > 10 {
		result = "Watchout"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is more than 10")
	}
	// if err !=nil{

	// }
}
