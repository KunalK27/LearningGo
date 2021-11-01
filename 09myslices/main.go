package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var fruitList = []string{"A", "B", "C"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)
	fruitList = append(fruitList, "C", "D")
	fmt.Println(fruitList)
	fruitList = append(fruitList)
	fmt.Println(fruitList)
	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	highScore = append(highScore, 555, 666, 321)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove value from slice based on index
	var courses = []string{"react", "js", "python", "Swift", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //append to remove values
	fmt.Println(courses)
}
