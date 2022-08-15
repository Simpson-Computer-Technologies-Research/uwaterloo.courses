package redis

// Import packages
import "fmt"

// The GenerateCourseHTML() function will use the course
// data map to generate an html string that is used for showing
// the list of courses on the home page.
func GenerateCourseHTML(data map[string]string) string {
	// Define Variables
	var (
		// Result: string -> The html result
		result string = ""

		// keys: []string -> They keys from the course data map
		keys []string = []string{
			"Components", "Unit", "ID", "Name",
			"Description", "Notes", "Pre Requisites",
			"Anti Requisites",
		}
	)

	// Iterate over the keys and add them
	// to the html result. The html result
	// will be returned inside styled divs
	for i := 0; i < len(keys); i++ {
		if len(data[keys[i]]) > 0 {
			result += fmt.Sprintf(`
			<div style="font-size:13px;">
				<strong>%s</strong> %v
			</div>
			`, keys[i], data[keys[i]])
		}
	}

	// Return the html wrapped in styled divs
	// If anyone has any suggestions for this part of the
	// code, please tell me!
	return fmt.Sprintf(`
	<div style="width: 100%s; margin-top: 5%s; margin-bottom: -1.3%s">
		<div class="course_div">
			<div style="font-size:13px;">
				<strong>%v</strong>
			</div>
			%s
		</div>
	</div>`, "%", "%", "%", data["Title"], result)
}

// Create html list from data
func GenerateHTML(data []map[string]string) string {
	// Html result string
	var html string = ""

	// Iterate over the course data slice
	for i := 0; i < len(data); i++ {
		html += GenerateCourseHTML(data[i])
	}
	// Return the html
	return html
}
