package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kunalk27/mongoapi/router"
)

func main() {
	fmt.Println("MONGODB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening at port 5000...")

}
