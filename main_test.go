package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(t *testing.T) {
	//Creating a new request object
	//Body left as an empty string for now
	req, err := http.NewRequest("GET", "", nil)

	//If error is found during creating the request, terminate process
	if err != nil {
		t.Fatal(err)
	}

	//Creating a new recorder object
	//This acts as a mini browser for us to pass requests to
	recorder := httptest.NewRecorder()

	//Creating a new http handler object
	//HTTP handlers handle http requests. This is our rest API.
	//We pass in our handler function to the new handler object
	hf := http.HandlerFunc(handler) //HTTP handler
	hf.ServeHTTP(recorder, req)     //Serves our request to the recorder

	//Checks that the status of the recorder is what we expect
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: expected %v but got %v", status, http.StatusOK)
	}

	fmt.Print(recorder.Body.String())
	//Checks that the handler function did what it is expcted to do
	expected := `hello world`
	actual := recorder.Body.String()
	if expected != actual {
		t.Errorf("handler returned unexpected body: expected %v but got %v", expected, actual)
	}
}

func testRouter(t *testing.T) {
	r := newRouter()
	server := httptest.NewServer(r)
	resp, err := http.Get(server.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status Error: expected %v but got %v", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// We want to hit the `GET /assets/` route to get the index.html file response
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// It isn't wise to test the entire content of the HTML file.
	// Instead, we test that the content-type header is "text/html; charset=utf-8"
	// so that we know that an html file has been served
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}
