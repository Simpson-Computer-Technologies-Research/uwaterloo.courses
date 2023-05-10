# uwaterloo.courses ![Stars](https://img.shields.io/github/stars/realTristan/uwaterloo.courses?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/uwaterloo.courses?label=Watchers)
![image](https://user-images.githubusercontent.com/75189508/231576310-b455740c-7b00-4c89-8c33-9465a994ff97.png)

# Challenge
My challenge for this project is to use solely native go modules.
This also includes natively webscraping data from websites (god help my soul..)

### Challenge Exceptions
```
hermes => An extremely fast full-text search algorithm (developed by me, and it uses native go modules)
```
```
fiber => Used for fast and easy api routing
```
```
fasthttp => Used for fast http requests
```

# Hermes
Hermes is an extremely fast full-text search algorithm and caching system written in Go. It's designed for API implementations which can be used by wrappers in other languages.
Hermes has two notable algorithms. The first being the with-cache algorithm. When using the with-cache algorithm, you can set, get, store, etc. keys and values into the cache. The second
being a no-cache algorithm that reads data from a map, or json file, and uses and array to store the data. Both of these algorithms provide full-text search query times from 10µs to 300µs.

# API
## Usage
### Get Courses (Example: Computer Science)
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

# Inspiration
This project was inspired by Eric Zhang's (@ekzhang) Harvard course catalog. I want to point out that I did not take any code from his project.

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
