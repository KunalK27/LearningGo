package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")
	Kunal := User{"Kunal", "kunal@gmal.com", true, 19}
	fmt.Println(Kunal)
	fmt.Printf("Name is %+v and email is %+v.\n", Kunal.Name, Kunal.Email)
	Kunal.GetStatus()
	Kunal.NewMail()
	fmt.Printf("Name is %+v and email is %+v.\n", Kunal.Name, Kunal.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
