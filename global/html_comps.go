package global

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// The EmptyDiv() function returns
// a div with empty spaces
func EmptyDiv() string {
	return `
		<div style="font-size:13px;">
			‏‏‎ ‎
		</div>
	`
}

// The HomePageSearchBar() function returns the home page
// search bar html string
func HomePageSearchBar() string {
	return `
	<div class="container">
	<img 
		src="static/images/waterloo_title_logo.png" 
		alt=""
		style="margin-top: -30%; display: block; margin-left: -13%; margin-right: auto; width: 120%;"
	>
        <form>
            <div class="input-field">
                <input
					minlength="3"
                    type="text" 
                    name="q" 
                    required="" 
                    id="name" 
                    formaction="/">
                <label style="font-size: 21px;">Search</label>
                <span></span>
            </div>
        </form>
	</div>
	`
}

// The SearchPageMenu() function is used to return the
// query speed in an html div
func SearchPageMenu(resultCount int) string {
	return fmt.Sprintf(
		`
		<div class="container">
        <form>
            <div class="input-field">
                <input 
					minlength="3"
                    type="text" 
                    name="q" 
                    required="" 
                    id="name" 
                    formaction="/">
                <label style="font-size: 21px;">Search</label>
                <span></span>
            </div>
        </form>
		<div 
			style="
				background-color: #fff; 
				padding: 1.2%%; 
				border-radius: 7px;
				margin: 5%%;
				text-align: center;
			"
		>
			<div>
				<strong>%d</strong> 
					results in 
				<strong style="color: #ffd54f">%vs</strong>
			</div>
		</div>
		</div>
		`, resultCount,
		math.Round(time.Since(SearchTime).Seconds()*100)/100)
}

// The GenerateCourseHTML() function will use the course
// data map to generate an html string that is used for showing
// the list of courses on the home page.
func GenerateCourseHTML(data map[string]string) string {
	// Define Variables
	var (
		// result: string -> The html result
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
			// Add the div with the key and value
			// to the html result
			result += fmt.Sprintf(`
				<div style="font-size:13px;">
					<strong>%s</strong> %v
				</div>
			`, keys[i], data[keys[i]])

			// Seperate the Title, Components, Unit and ID
			// from the Name, Description, and so on
			if i == 2 {
				result += EmptyDiv()
			}
		}
	}
	// Append an empty div to seperate the href
	// urls from the course data
	result += EmptyDiv()

	// Split the course title to get the subject code
	// and the course code
	var info []string = strings.Split(data["Title"], " ")

	// Append the official url to the course to the result html
	result += fmt.Sprintf(`
		<div style="font-size: 13px;">
			<a 
				href="https://ucalendar.uwaterloo.ca/2223/COURSE/course-%s.html#%s%s">
				Official
			</a>
		</div>
	`, info[0], info[0], info[1])

	// Return the html wrapped in styled divs
	// If anyone has any suggestions for this part of the
	// code, please tell me!
	return fmt.Sprintf(`
		<div style="width: 100%%; margin-top: 5%%; margin-bottom: -1.3%%">
			<div class="course_div" style="position: relative;">
			<img 
				src="static/images/waterloo_logo.png" 
				alt="" 
				width="80"
				height="80"
				style="position: absolute; top: 8.3%%; right: 1.7%%;"
			>
				<div style="font-size:13px;">
					<strong>%v</strong>
				</div>
				%s
			</div>
		</div>`, data["Title"], result)
}
