package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["Py"] = "python"
	languages["Ja"] = "javat"
	fmt.Println("List of all languages", languages)
	fmt.Println("JS short form of", languages["JS"])
	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	//loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)

	}
}
