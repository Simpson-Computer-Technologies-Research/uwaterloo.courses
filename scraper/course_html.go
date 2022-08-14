package scraper

import "fmt"

// The AppendHTML() function...
func (cs *CourseScrape) AppendHTML(key string, value string) {
	cs.HTML += fmt.Sprintf(
		`<div><strong> %v</strong> %v</div>`, key, value)
}

// The WrapHTML() function...
func (cs *CourseScrape) WrapHTML() string {
	return fmt.Sprintf(`<div class="course_div">%s</div>`, cs.HTML)
}
