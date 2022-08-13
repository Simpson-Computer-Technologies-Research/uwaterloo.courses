package global

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
	"BET", "BIOL", "BME", "BUS",
	// Alpha: C
	"CDNST", "CHE", "CHEM", "CHINA", "CI", "CIVE", "CLAS", "CMW",
	"CO", "COMM", "CROAT", "CS",
	// Alpha: D
	"DAC", "DUTCH",
	// Alpha: E
	"EARTH", "EASIA", "ECE", "ECON", "EMLS", "ENBUS",
	"ENGL", "ENVE", "ENVS", "ERS",
	// Alpha: F
	"FINE", "FR",
	// Alpha: G
	"GBDA", "GENE", "GEOE", "GEOG", "GER", "GERON", "GRK", "GSJ",
	// Alpha: H
	"HIST", "HLTH", "HRM", "HRTS", "HUMSC",
	// Alpha: I
	"INDEV", "INDG", "INTEG", "INTST", "ITAL", "ITALST",
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
	"PACS", "PD", "PDARCH", "PHARM", "PHIL", "PHYS", "PLAN",
	"PMATH", "PSCI", "PSYCH",
	// Alpha: R
	"REC", "RS", "RUSS",
	// Alpha: S
	"SCBUS", "SCI", "SDS", "SE", "SI", "SMF", "SOC",
	"SOCWK", "SPAN", "SPCOM", "STAT", "STV", "SYDE",
	// Alpha: T, U, V, W
	"THPERF", "UNIV", "VCULT", "WKRPT",
}

// The Subject Names for search querying the course catalog
//
// I haven't added them yet because my wifi is brutal and I can't
// send any http requests, but I when I can, I will pre-insert them
var SubjectNames map[string]string = map[string]string{
	"actuarialscience":                       "ACTSC",
	"architecturalengineering":               "AE",
	"accountingfinancialmanagement":          "AFM",
	"appliedmathematics":                     "AMATH",
	"anthropology":                           "ANTH",
	"appliedlanguagestudies":                 "APPLS",
	"arabic":                                 "ARABIC",
	"artsandbusiness":                        "ARBUS",
	"headbodyarchitecture":                   "ARCH",
	"arts":                                   "ARTS",
	"americansignlanguage":                   "ASL",
	"aviation":                               "AVIA",
	"businessentrepreneurshipandtechnology":  "BET",
	"biology":                                "BIOL",
	"biomedicalengineering":                  "BME",
	"businessbrwilfridlaurieruniversity":     "BUS",
	"canadianstudies":                        "CDNST",
	"chemicalengineering":                    "CHE",
	"chemistry":                              "CHEM",
	"chinese":                                "CHINA",
	"culturalidentities":                     "CI",
	"civilengineering":                       "CIVE",
	"classicalstudies":                       "CLAS",
	"churchmusicandworship":                  "CMW",
	"combinatoricsandoptimization":           "CO",
	"commerce":                               "COMM",
	"croatian":                               "CROAT",
	"computerscience":                        "CS",
	"digitalartscommunication":               "DAC",
	"dutch":                                  "DUTCH",
	"earthsciences":                          "EARTH",
	"eastasian":                              "EASIA",
	"electricalandcomputerengineering":       "ECE",
	"economics":                              "ECON",
	"englishformultilingualspeakers":         "EMLS",
	"environmentandbusiness":                 "ENBUS",
	"english":                                "ENGL",
	"environmentalengineering":               "ENVE",
	"environmentalstudies":                   "ENVS",
	"environmentresourcesandsustainability":  "ERS",
	"finearts":                               "FINE",
	"frenchstudies":                          "FR",
	"globalbusinessanddigitalarts":           "GBDA",
	"generalengineering":                     "GENE",
	"geologicalengineering":                  "GEOE",
	"geographyandenvironmentalmanagement":    "GEOG",
	"german":                                 "GER",
	"gerontology":                            "GERON",
	"greek":                                  "GRK",
	"genderandsocialjustice":                 "GSJ",
	"history":                                "HIST",
	"health":                                 "HLTH",
	"humanresourcesmanagement":               "HRM",
	"humanrights":                            "HRTS",
	"humansciences":                          "HUMSC",
	"internationaldevelopment":               "INDEV",
	"indigenousstudies":                      "INDG",
	"knowledgeintegration":                   "INTEG",
	"internationalstudies":                   "INTST",
	"italian":                                "ITAL",
	"italianstudies":                         "ITALST",
	"japanese":                               "JAPAN",
	"jewishstudies":                          "JS",
	"kinesiology":                            "KIN",
	"korean":                                 "KOREA",
	"latin":                                  "LAT",
	"legalstudies":                           "LS",
	"mathematicalbusiness":                   "MATBUS",
	"mathematics":                            "MATH",
	"mechanicalengineering":                  "ME",
	"medievalstudies":                        "MEDVL",
	"management":                             "MGMT",
	"bodymaterialsandnanosciences":           "MNS",
	"mohawk":                                 "MOHAWK",
	"managementsciences":                     "MSCI",
	"mechatronicsengineering":                "MTE",
	"mathematicselectives":                   "MTHEL",
	"music":                                  "MUSIC",
	"nanotechnologyengineering":              "NE",
	"optometry":                              "OPTOM",
	"peaceandconflictstudies":                "PACS",
	"professionaldevelopment":                "PD",
	"professionaldevelopmentforarchitecture": "PDARCH",
	"pharmacy":                               "PHARM",
	"philosophy":                             "PHIL",
	"physics":                                "PHYS",
	"planning":                               "PLAN",
	"puremath":                               "PMATH",
	"politicalscience":                       "PSCI",
	"psychology":                             "PSYCH",
	"recreationandleisurestudies":            "REC",
	"religiousstudies":                       "RS",
	"russian":                                "RUSS",
	"scienceandbusiness":                     "SCBUS",
	"science":                                "SCI",
	"socialdevelopmentstudies":               "SDS",
	"softwareengineering":                    "SE",
	"studiesinislam":                         "SI",
	"sexualitymarriageandfamilystudies":      "SMF",
	"sociology":                              "SOC",
	"socialworkbrsocialdevelopmentstudies":   "SOCWK",
	"spanish":                                "SPAN",
	"speechcommunication":                    "SPCOM",
	"statistics":                             "STAT",
	"societytechnologyandvalues":             "STV",
	"systemsdesignengineering":               "SYDE",
	"theatreandperformance":                  "THPERF",
	"university":                             "UNIV",
	"visualculture":                          "VCULT",
	"worktermreport":                         "WKRPT",
}

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
