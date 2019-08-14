package main

import (
	"fmt"
	"html/template"
	"net/http"
	//"strings"
	"time"
)

func main() {
	fmt.Println("The Server runs with http://localhost:3000/")
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}

type Time struct {
	Year    int
	Month   string
	Day1    int
	Weekday string
	Hour    int
	Minute  int
	Second  int
}

func Handler(w http.ResponseWriter, r *http.Request) {

	t := Time{
		Year:    time.Now().Year(),
		Month:   time.Now().Month().String(),
		Day1:    time.Now().Day(),
		Hour:    time.Now().Hour(),
		Minute:  time.Now().Minute(),
		Second:  time.Now().Second(),
		Weekday: time.Now().Weekday().String(),
	}

	tmpl := template.Must(template.ParseFiles("./resources/index.html"))
	tmpl.Execute(w, t)

}
