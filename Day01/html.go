package main

import (
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`my first website`))
}

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}

// 建立 http
// http://127.0.0.1:8080/
// use go mod init Day01 to create the go mod file
// then ctrl + ~ open the terminal and cd ~ to Day01 folder
// write go run . in terminal
// if you not use go mod, then use go run html.go to run the code.
