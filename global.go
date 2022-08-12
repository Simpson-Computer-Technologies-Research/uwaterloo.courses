package main

// The slice that contains all the university of waterloo's subject codes
// I'm going to set them manually for now, but there is a webscrape section
// for getting them using info from https://classes.uwaterloo.ca/uwpcshtm.html
//
// Sorted them by alphabet because I was bored
var SubjectCodes []string = []string{
	// Alpha: A
	"ACTSC", "AE", "AFM", "AMATH", "ANTH", "APPLS", "ARABIC",
	"ARBUS", "ARCH", "ARTS", "ASL", "AVIA",
	// Alpha: B
	"BET", "BIOL", "BLKST", "BME", "BUS",
	// Alpha: C
	"CDNST", "CFM", "CHE", "CHEM", "CHINA", "CI", "CIVE", "CLAS", "CMW",
	"CO", "COMM", "CROAT", "CS",
	// Alpha: D
	"DAC", "DUTCH",
	// Alpha: E
	"EARTH", "EASIA", "ECE", "ECON", "EMLS", "ENBUS",
	"ENGL", "ENTR", "ENVE", "ENVS", "ERS",
	// Alpha: F
	"FINE", "FR",
	// Alpha: G
	"GBDA", "GENE", "GEOE", "GEOG", "GER", "GERON", "GRK", "GSJ",
	// Alpha: H
	"HEALTH", "HIST", "HLTH", "HRM", "HRTS", "HUMSC",
	// Alpha: I
	"INDENT", "INDEV", "INDG", "INTEG", "INTST", "ITAL", "ITALST",
	// Alpha: J
	"JAPAN", "JS",
	// Alpha: K
	"KIN", "KOREA",
	// Alpha: L
	"LAT", "LS",
	// Alpha: M
	"MATBUS", "MATH", "ME", "MEDVL", "MGMT", "MNS", "MOHAWK",
	"MSCI", "MTE", "MTHEL", "MUSIC",
	// Alpha: N, O
	"NE", "OPTOM",
	// Alpha: P
	"PACS", "PD", "PDARCH", "PHARM", "PHIL", "PHYS", "PLAN", "PMATH",
	"PSCI", "PSYCH",
	// Alpha: R
	"REC", "RS", "RSCH", "RUSS",
	// Alpha: S
	"SCBUS", "SCI", "SDS", "SE", "SFM", "SI", "SMF", "SOC",
	"SOCWK", "SPAN", "SPCOM", "STAT", "STV", "SYDE",
	// Alpha: T, U, V, W
	"THPERF", "UNIV", "VCULT", "WKRPT",
}

// The Subject Names for search indexing the course catalog
//
// I haven't added them yet because my wifi is brutal and I can't
// send any http requests, but I when I can, I will pre-insert them
var SubjectNames []string = []string{}

// The SliceContains() function returns whether or not the provided
// slice contains the provided string
func SliceContains(s *[]string, str string) bool {
	// Iterate over the slice
	for _, v := range *s {
		// if the slice value equals the string then return true
		if v == str {
			return true
		}
	}
	// Else return false
	return false
}
