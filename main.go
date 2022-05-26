package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Id         int
	Name       string
	Dept       string
	Pin        int
	Rollnumber string
}

var students []Student = []Student{
	{
		Id:   121,
		Name: "sush",
		Dept: "EC",
		Pin:  123,
	},
	{
		Id:   122,
		Name: "vinu",
		Dept: "CS",
		Pin:  456,
	},
}

func main() {
	http.HandleFunc("/", roothandler)
	http.HandleFunc("/students", handleGetStudents)
	http.ListenAndServe("localhost:8000", nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><body><h1>Welcome to my college</h1></body</head></html>")
}

func handleGetStudents(w http.ResponseWriter, r *http.Request) {

	for i := range students {

		rollnumber := fmt.Sprintf("%s%d", students[i].Dept, students[i].Id)
		fmt.Println("debug", rollnumber)
		students[i].Rollnumber = rollnumber
	}

	studStr, _ := json.MarshalIndent(students, "", " ")
	fmt.Fprintf(w, string(studStr))
}
