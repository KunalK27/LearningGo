package main

import "fmt"

func main() {
	fmt.Println("Loops")
	days := []string{"A", "B", "C", "D", "E", "F", "G"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			break
		}
		fmt.Println("value is : ", rogueValue)
		rogueValue++
	}
lco:
	fmt.Println("Jumping at learncode")
}
