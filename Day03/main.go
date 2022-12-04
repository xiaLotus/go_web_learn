package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("temp.html")

	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	number := 10
	t.Execute(w, number)
}

func main() {
	http.HandleFunc("/", helloHandleFunc)
	http.ListenAndServe(":8080", nil)
}
