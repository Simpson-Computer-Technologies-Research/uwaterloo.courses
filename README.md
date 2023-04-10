# uwaterloo.courses ![Stars](https://img.shields.io/github/stars/realTristan/uwaterloo.courses?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/uwaterloo.courses?label=Watchers)
<img width="1383" alt="Screen Shot 2022-08-23 at 7 41 53 PM" src="https://user-images.githubusercontent.com/75189508/186290354-a5ed2710-f1a5-43c8-ae19-c0252d874fa3.png">


# Challenge
My challenge for this project is to use solely native golang modules.
This also includes natively webscraping data from websites (god help my soul..)

### Challenge Exceptions

- mux module -> Used for api routing
- fasthttp module -> Used for http requests

# About
### Why Golang?

- Golang is fast, lightweight and easy to use for hosting API's. I have previous experience with golang and decided it was the best option for this project.

### Why Svelte?

- Svelte is 30% faster than other frameworks.
- Svelte is best when used for designing small apps

# API
<h3>Why make an API?</h3>
I decided to make an api because it will be used for refreshing the cache. I also chose to make an api because the university of waterloo's api only allows 5000 requests per month and the data is quite limited.

## Usage

### Get Subject Data (Computer Science)

```
$ curl -X GET http://localhost:8000/courses?q=computer+science
```

```json
"Example Response": {
    [
        "title": "CS 476",
        "name": "Numeric Computation for Financial Modeling",
        "desc": "The interaction of financial models, numerical methods, and computing environments. Basic  computational aspects of option pricing and hedging. Numerical methods for stochastic differential equations, strong and weak convergence. Generating correlated random numbers. Time-stepping methods. Finite difference methods for the Black-Scholes equation. Discretization, stability, convergence. Methods for portfolio optimization, effect of data errors on portfolio weights. ",
        "note": "Lab is not scheduled and students are expected to find time in open hours to complete their work. Students who receive a good grade in CS 335 may contact the instructor of CS 476 to seek admission without the formal prerequisites. Offered: W]",
        "pre_reqs": "(AMATH 242/CS 371 or CS 370) and STAT 231/241",
        "anti_reqs": "None",
        "co_reqs": "None",
        "unit":"0.50"
    ]
}
```

### Get Subjects

```
$ curl -X GET http://localhost:8000/subjects
```

```json
"Example Response": {
    "subjects": [
        "CS", "PHYS", "CHEM", "ART", "HLTH"
    ]
}
```

### Get Subjects and Names

```
$ curl -X GET http://localhost:8000/subjects/names
```

```json
"Example Response": {
    "actuarialscience":                       "ACTSC",
	"architecturalengineering":               "AE",
	"accountingfinancialmanagement":          "AFM",
	"appliedmathematics":                     "AMATH",
	"anthropology":                           "ANTH",
	"appliedlanguagestudies":                 "APPLS",
	"arabic":                                 "ARABIC",
	"artsandbusiness":                        "ARBUS",
	"headbodyarchitecture":                   "ARCH",
	"arts":                                   "ARTS"
}

```

# Acknowledgements
This project was inspired by Eric Zhang's (@ekzhang) Harvard course catalog. I want to point out that I did not steal any code from his project.

# License
MIT License

Copyright (c) 2022 Tristan Simpson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
