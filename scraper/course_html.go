package scraper

import "fmt"

// The AppendHTML() function takes the course key and the
// course info and puts it into an html div.
// This html div is used as a value in the final html list
func (cs *CourseScrape) AppendHTML(key string, value string) {
	cs.HTML += fmt.Sprintf(
		`<div style="font-size: 13px;"><strong> %v</strong> %v</div>`, key, value)
}

// The WrapHTML() function wraps the final course data list html
// into a div that is then styled using the style.css static file
func (cs *CourseScrape) WrapHTML() string {
	return fmt.Sprintf(`<div class="course_div">%s</div>`, cs.HTML)
}
