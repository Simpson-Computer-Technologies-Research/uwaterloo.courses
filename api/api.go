package api

// Import packages
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// The ListenAndServe() function hosts the api
// And corresponds the request path to the correct
// api functions
func ListenAndServe(port string) {
	// Print the localhost url
	fmt.Printf("Listening on: http://localhost%s", port)

	// Establish a new gorilla mux router
	var router *mux.Router = mux.NewRouter()

	// Show course data with the paramter ?course={course_code}
	router.HandleFunc("/courses", CourseDataHandler()).Methods("GET")

	// Show the home page of the course catalog
	// This is the area where you can search for courses
	router.HandleFunc("/", HomePageHandler()).Methods("GET")

	// Show the list of subjects at the university of waterloo
	router.HandleFunc("/subjects", SubjectCodesHandler()).Methods("GET")

	// Show the list of subjects with their corresponding names
	// at the university of waterloo
	router.HandleFunc("/subjects/name", SubjectCodesWithNamesHandler()).Methods("GET")

	// Developement Testing
	router.HandleFunc("/dev", DevTestingHandler()).Methods("GET")

	// Handle Router
	http.Handle("/", router)

	// Serve Static Files: html, css, images, etc.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Listen and Serve to the corresponding port
	http.ListenAndServe(port, nil)
}
