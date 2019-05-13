package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

func main() {
	// router, like a signal-slot, create a mapping relationship
	// the route will trigger corresponding function
	http.HandleFunc("/", sayHello)

	// start listening to the port
	http.ListenAndServe(":8888", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	user := User{Username: "tong"}
	tpl, _ := template.ParseFiles("templates/index.html")
	tpl.Execute(w, &user)
}