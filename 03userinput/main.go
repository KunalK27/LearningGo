package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Pizza Rating")

	//comma ok or error ok syntax, susbtitute for try catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of rating %T", input)
}
