package api

// Import packages
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	hermes "github.com/realTristan/Hermes/nocache"
)

// The ListenAndServe() function hosts the api
// And corresponds the request path to the correct
// api functions
func ListenAndServe(port string) {
	// Hermes cache
	// go get github.com/realTristan/Hermes
	var cache, err = hermes.InitWithJson("default_data.json", map[string]bool{
		"id":             false,
		"components":     false,
		"units":          false,
		"pre_requisites": false,
		"title":          false,
		"description":    true,
		"name":           true,
	})

	// Check for errors
	if err != nil {
		panic(err)
	}

	// Print the localhost url
	fmt.Printf(" >> Listening on: http://localhost%s\n", port)

	// Establish a new gorilla mux router
	var router *mux.Router = mux.NewRouter()

	// Home Page
	router.HandleFunc("/", HomePageHandler()).Methods("GET")

	// Dev Page
	router.HandleFunc("/dev", DevPageHandler()).Methods("GET")

	// Show course data with the paramter ?course={course_code}
	router.HandleFunc("/courses", CourseDataHandler(cache)).Methods("GET")

	// Show the list of subjects at the university of waterloo
	router.HandleFunc("/subjects", SubjectCodesHandler()).Methods("GET")

	// Show the list of subjects with their corresponding names
	// at the university of waterloo
	router.HandleFunc("/subjects/names", SubjectCodesWithNamesHandler()).Methods("GET")

	// Handle Router
	http.Handle("/", router)

	// Listen and Serve to the corresponding port
	http.ListenAndServe(port, nil)
}
