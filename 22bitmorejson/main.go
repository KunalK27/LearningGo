package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json Video")
	// EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	lcoCourses := []course{
		{"RectJs Bootcamp", 299, "Learncodeonline.in", "abcd123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Learncodeonline.in", "ahabcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "Learncodeonline.in", "abcdasxs123", nil},
	}
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "RectJs Bootcamp",
		"Price": 299,
		"website":"Learncodeonline.in",
		"tags": ["web-dev",	"js"]
	}
	`)
	var lcocourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where you want to add data to key value pair
	var myonlinedata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlinedata)
	fmt.Printf("%#v\n", myonlinedata)

	for k, v := range myonlinedata {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
