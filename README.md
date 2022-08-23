# Challenge
My challenge for this project is to use solely native golang modules such as
strings, net/http, html/template, etc. This also includes natively webscraping
data from websites (god help my soul..)

<h3> Challenge Exceptions </h3>

- mux module -> Used for api routing
- fasthttp module -> Used for low memory http requests
- redis module > Used for Caching data
- External Services -> Ex: Svelte, Redis, Hosting, etc.

<br>

# API
<h3>Why make an API?</h3>
I decided to make an api because it will be used for refreshing
the redis database whenever somebody calls the /course endpoint.
I also chose to make an api because the university of waterloo
api only allows 5000 requests per month and the data is quite limited.
<br>
Although my api is not as fast and big as the university of waterloo's,
I just thought it'd be a good addition to the project

<h3>Usage</h3>

Example Response (Get Computer Science Data)
<br>

```
$ curl -X GET http://localhost:8000/courses?course=CS

or query with:

$ curl -X GET http://localhost:8000/courses?q=computer+science
```

```json
{
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

Example Response (Get Subjects List)
<br>

```
$ curl -X GET http://localhost:8000/subjects
```

```json
{
    "subjects": [
        "CS", "PHYS", "CHEM", "ART", "HLTH"
    ]
}
```

Example Response (Get Subjects and Names)
<br>

```
$ curl -X GET http://localhost:8000/subjects/names
```

```json
{
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

<br>

# Project Showcase
<img width="1122" alt="Screen Shot 2022-08-22 at 10 18 09 PM" src="https://user-images.githubusercontent.com/75189508/186062009-49d2782a-7e81-4f05-893c-cfb047223b7c.png">

<br>

# Acknowledgements
I was inspired to create this project after seeing Eric Zhang's (@ekzhang)
<br>
Harvard course catalog. I want to point out that I did not copy/paste any
<br>
code from his project.

<br>

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

