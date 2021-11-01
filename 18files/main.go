package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")
	content := "This should go in a file"
	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("The length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text Data in the file is:\n ", string(databyte))
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
