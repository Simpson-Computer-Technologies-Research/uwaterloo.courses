package main

import (
	"fmt"
	"strings"
)

// The CleaCourseTitle() function will remove spaces from the scraped course
// title, except for the double space in the title, it will only leave one
// space. The course title will be used for querying the database
func CleanCourseTitle(title string) string {
	// Iterate over the title string
	for i := 0; i < len(title); i++ {
		// Make sure the indexed value is a space and i is greater than 0
		// This is so we can check the title's previous index without getting
		// any errors
		if title[i] == ' ' && i > 0 {
			// Make sure the previous indexed value is not a space
			if title[i-1] != ' ' {
				// Remove the current index value
				title = title[:i] + title[i+1:]
			}
		}

	}
	// Return the new title in lowercase
	return strings.ToLower(title)
}

// The BasicSplitString() function is equivalent to the strings.Split() function
// The only different is that the BasicSplitString() function can only split by a
// single letter
//
// Notes: I decided to make my own function instead of using the strings module
// to have more control over what is happening inside the function
func BasicSplitString(s string, splitBy string) []string {
	// Declare variables
	// - str: string -> the string being appending to res
	// - res: []string -> the slice containing the split strings
	var (
		str string = ""
		res []string
	)
	// Iterate over the s: string variable
	for i := 0; i < len(s); i++ {
		// Convert st[i] from byte to string
		var strI string = string(s[i])
		// Check if the str[i] doesn't equal the string to split by
		if strI != splitBy {
			// Add the str[i] to the str
			str += strI
		} else {
			// Append the str to the result slice
			res = append(res, str)
			// Reset the str
			str = ""
		}
	}
	// Return the split strings slice
	return res
}

// The SliceContains() function returns whether or not the provided
// slice contains the provided string
func SliceContains(s []string, str string) bool {
	// Iterate over the slice
	for _, v := range s {
		// if the slice value equals the string then return true
		if v == str {
			return true
		}
	}
	// Else return false
	return false
}

// Convert the course subject info into categories
// For example it will convert MATH 235 LEC,IST,TUT 0.50
// to -> var subjectTitle, components, unit = MATH 235, LEC,IST,TUT, 0.50
func IndexCourseSubjectInfo(title string) (string, string, string) {
	// splitTitle: []string -> splits the given title by spaces
	// subjectTitle: string -> the subject name and catalog number (MATH 232)
	var (
		splitTitle   []string = BasicSplitString(title, " ")
		subjectTitle string   = fmt.Sprintf("%s %s", splitTitle[0], splitTitle[1])
	)
	// Return the Course Subject Title ex: CS 201
	// Return the Components (splitTitle[2]) ex: LAB,LEC,TST
	// Return the Unit (splitTitle[3]) ex: 0.50
	return subjectTitle, splitTitle[2], splitTitle[3]
}
