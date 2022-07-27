package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Dog struct {
	name  string `json:"name"`
	owner string `json:"owner"`
}

var dogs []Dog

func getDogHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "birds" variable to json
	dogListBytes, err := json.Marshal(dogs)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(dogListBytes)
}

func createDogHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Dog
	dog := Dog{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	dog.name = r.Form.Get("name")
	dog.owner = r.Form.Get("owner")

	// Append our existing list of birds with a new entry
	dogs = append(dogs, dog)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
