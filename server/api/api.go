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
	fmt.Printf(" >> Listening on: http://localhost/%s\n", port)

	// Establish a new gorilla mux router
	var router *mux.Router = mux.NewRouter()

	// Home Page
	router.HandleFunc("/", HomePageHandler()).Methods("GET")

	// Dev Page
	router.HandleFunc("/dev", DevPageHandler()).Methods("GET")

	// Show course data with the paramter ?course={course_code}
	router.HandleFunc("/courses", CourseDataHandler()).Methods("GET")

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
