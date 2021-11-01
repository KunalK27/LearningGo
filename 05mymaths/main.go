package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to maths")

	// var mynumberone int =2
	// var mynumbertwo float64 =4.5
	//fmt.Println("The sum is",mynumberone+int(mynumbertwo))

	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("My random number", myRandomNumber)
}
