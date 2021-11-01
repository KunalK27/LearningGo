package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//no inheritance or super or parent in golang

	Kunal := User{"Kunal", "kunal@gmal.com", true, 19}
	fmt.Println(Kunal)
	fmt.Printf("Name is %+v and email is %+v.", Kunal.Name, Kunal.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
