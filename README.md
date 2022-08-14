# The University of Waterloo Course Catalog
This was made by Tristan Simpson
<br>
This project is unfinished but I wanted to put it on github for any suggestions on what I should add or any improvements to my code
<br>
The license will be converted to an Official MIT Copyright License once an official Github Release of this repository has been created :)

# Challenge
My challenge for this project is to use solely native golang modules such as
strings, net/http, html/template, etc. This also includes natively webscraping
data from websites. God help my soul for webscraping..

<h3> Challenge Exceptions </h3>

- mux module -> Used for api routing
- fasthttp module -> Used for low memory http requests
- redis module > Used for Caching data
- External Services -> Ex: Svelte, Redis, Fly.io, etc.

# Notes
- I'm not the best at frontend developement, so go easy on me

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
    "003352": {
        "title": "CS 476",
        "name": "Numeric Computation for Financial Modeling",
        "desc": "The interaction of financial models, numerical methods, and computing environments. Basic  computational aspects of option pricing and hedging. Numerical methods for stochastic differential equations, strong and weak convergence. Generating correlated random numbers. Time-stepping methods. Finite difference methods for the Black-Scholes equation. Discretization, stability, convergence. Methods for portfolio optimization, effect of data errors on portfolio weights. ",
        "note": "Lab is not scheduled and students are expected to find time in open hours to complete their work. Students who receive a good grade in CS 335 may contact the instructor of CS 476 to seek admission without the formal prerequisites. Offered: W]",
        "pre_reqs": "(AMATH 242/CS 371 or CS 370) and STAT 231/241",
        "anti_reqs": "None",
        "co_reqs": "None",
        "unit":"0.50"
    }
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

# Project Showcase
<img width="1311" alt="Screen Shot 2022-08-14 at 4 56 40 PM" src="https://user-images.githubusercontent.com/75189508/184554649-373345f0-502d-4a16-8579-0f6f9bfa2b5d.png">



# License
Copyright (C) 2022 Tristan Simpson <heytristaann@gmail.com>

This file is part of the University of Waterloo Course Catalog project.

The University of Waterloo Course Catalog project can not be copied and/or distributed without the express permission of Tristan Simpson <heytristaann@gmail.com>. 
