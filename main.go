package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	// These lines are added inside the newRouter() function before returning r
	r.HandleFunc("/dog", getDogHandler).Methods("GET")
	r.HandleFunc("/dog", createDogHandler).Methods("POST")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":4200", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
