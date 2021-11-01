package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for course-file
type Course struct {
	CourseId    string  `json:"courseid"`
	Coursname   string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
var courses []Course

//middleware or helpers -file
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.Coursname == ""
	return c.Coursname == ""
}
func main() {
	fmt.Println("API making")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "2", Coursname: "ReactJS", CoursePrice: 200, Author: &Author{Fullname: "Kunal", Website: "kunal.com"}})
	courses = append(courses, Course{CourseId: "4", Coursname: "MERN Stack", CoursePrice: 300, Author: &Author{Fullname: "Kunal", Website: "kunalk.com"}})
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers -file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from request
	params := mux.Vars(r)
	//loop through courses, find matching id and return respomse
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no courses found with given id")
	return

}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	//what if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	//what if data is like {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("no data inside json")
	}
	//TODO: Check only if  title is duplicate
	//loop, if title matches with course.coursename then display message

	//generate unique id and conversion to string
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first- grab id from request
	params := mux.Vars(r)

	//loop through value, once you get id, remove it, then add with myid (that is the new id)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//todo: response when id not found
}
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//loop, id, remove (index, index-1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
			//TODO: Send a confirm or deny response
		}
	}
}
