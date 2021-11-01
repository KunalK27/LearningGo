package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Pinepple"
	// fruitList[3] = "Peru"
	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Fruit list is", len(fruitList))
	var vegList = [3]string{"potato", "beet", "mushroom"}
	fmt.Println("Veg list is", len(vegList))
}
