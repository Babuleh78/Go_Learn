package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

//	func home_page(page http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(page, "Thunder")
//	}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.name, u.age, u.money)
}

func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{"Babuleh", 18, -1000, 5.0, 1.0}
	// fmt.Fprintf(page, "User name is "+bob.name)
	//fmt.Fprint(page, `<b> Main Text </b> <h1> Main Text</h1>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(page, bob)
}

func babuleh_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Babuleh is here")
}

func HandleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/babuleh/", babuleh_page)
	http.ListenAndServe(":8080", nil)
}
func main() {
	//var bob User = ....
	//bob := User{name: "babuleh", age:18, money:-1000, avg_grades:5, happiness: 1}

	HandleRequest()
}
