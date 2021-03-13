package main

import (
	"log"
	"net/http"
	"time"

	mux "github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/v1/ideas", handler).Methods("GET")
	r.HandleFunc("/v1/users", handler).Methods("GET")
	r.HandleFunc("/v1/user/{id}", handler).Methods("GET")
	r.HandleFunc("/v1/idea/{id}", handler).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	http.Handle("/", r)
	log.Fatal(srv.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/v1/ideas" {
		w.Write([]byte("You are here"))
		return //if not here you will fall through and both strings will be added
	}
	w.Write([]byte("Path value checked"))
	return
}
