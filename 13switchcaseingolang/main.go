package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and Case in Golang")
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you can move two spots")
	case 3:
		fmt.Println("Dice value is 3 and you can open")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 and you can open")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 and you can open")
	case 6:
		fmt.Println("Dice value is 6 and you can open")
	default:
		fmt.Println("What was that!")
	}
}
