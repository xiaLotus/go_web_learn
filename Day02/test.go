package main

import (
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`my first website`))
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`my first website1`))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`my first website2`))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", test)
	mux.HandleFunc("/hi", hi)
	mux.HandleFunc("/hi/index", index)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// http.ListenAndServe(":8080", nil)
}
