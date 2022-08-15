package global

import (
	"fmt"
	"math"
	"time"
)

// The HomePageSearchBar() function returns the home page
// search bar html string
func HomePageSearchBar() string {
	return `
	<div class="container">
	<img 
		src="static/images/waterloo_title_logo.png" 
		alt=""
		style="margin-top: -30%; display: block; margin-left: -12.5%; margin-right: auto; width: 120%;"
	>
        <form>
            <div class="input-field">
                <input
					minlength="5"
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

// The EndQueryTimer() function is used to return the
// query speed in an html div
func EndQueryTimer(resultCount int) string {
	return fmt.Sprintf(
		`
		<div class="container">
        <form>
            <div class="input-field">
                <input 
					minlength="5"
                    type="text" 
                    name="q" 
                    required="" 
                    id="name" 
                    formaction="/">
                <label>Search</label>
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
				<strong style="color: #FEDD00">%vs</strong>
			</div>
		</div>
		</div>
		`, resultCount,
		math.Round(time.Since(SearchTime).Seconds()*100)/100)
}
