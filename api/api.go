package api

// Import packages
import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Main function
func ListenAndServe(port string) {
	// Print the localhost url
	fmt.Printf("Listening on: http://localhost%s", port)

	// Listen and Server the port
	fasthttp.ListenAndServe(port, func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		// Show course data with the paramter ?course={course_code}
		case "/courses":
			CourseDataHandler(ctx)

		// Show the list of subjects at the university of waterloo
		case "/subjects":
			SubjectCodesHandler(ctx)

		// Show the list of subjects with their corresponding names
		// at the university of waterloo
		case "/subjects/names":

		// Invalid path error
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	})
}
