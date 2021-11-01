package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Verb video")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	var responsestring strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responsestring.Write(content)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responsestring.String())
	// fmt.Println(content)
	// fmt.Println(string(content))
}
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"
	//fake json payload
	requestbody := strings.NewReader(`
	{
		"coursename": "Let's go with golang",
		"price": 0,
		"platform": "learnCodeonline.in"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestbody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstName", "Kunal")
	data.Add("lastName", "Kulkarni")
	data.Add("email", "Kunal@gmail.com")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
