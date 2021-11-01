package main

import "fmt"

const LoginToken string = "sqhdbsjcb" //First letter is capital so public variable

func main() {
	var username string = "Kunal"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.2828201290130
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and aliases

	var anothervar int
	fmt.Println(anothervar)
	fmt.Printf("Variable is of type: %T \n", anothervar)

	//implicit variable declaration

	var website = "kunal.com"
	fmt.Println(website)

	//no var style

	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)

	//public
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
