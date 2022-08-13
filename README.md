# The University of Waterloo Course Catalog
This was made by Tristan Simpson, do not steal lmao
<br>
This project is unfinished but I wanted to put it on github for any suggestions on what I should add or any improvements to my code
<br>
The license will be converted to an Official MIT Copyright License once an official Github Release of this repository has been created :)


# Challenge
My challenge for this project is to use solely native golang modules
other than fasthttp which is used to host the api and send http requests.
- This challenge does not include hosting services, database services, etc.
- God help my soul for webscraping..

# Plan
My plan is to have a backend that will scrape all the course codes then use those codes
to scrape the course info whenever the university of waterloo's website changes.
It'll then replace the data inside the redis database with the new data.

If the process takes to long, I can increase the speed by storing the course codes in memory
instead of having to scrape them.


# API
<h3>Usage</h3>

```go
//////////////////////////////////////////////////////////////////////////////////////////
//                                                                                      //
// The University of Waterloo API doesn't provide much information about courses.       //
// Thus, I have decided to not only create a course catalog, but an open source api     //
// That is free to use, free of charge.                                                 //
//                                                                                      //
//////////////////////////////////////////////////////////////////////////////////////////

// Example Response (Get Computer Science Data)
// $ curl -X GET http://localhost:8000?course=CS
{
    "course id": {
        "title": "CS 476",
        "name": "Numeric Computation for Financial Modeling",
        "desc": "The interaction of financial models, numerical methods, and computing environments. Basic  computational aspects of option pricing and hedging. Numerical methods for stochastic differential equations, strong and weak convergence. Generating correlated random numbers. Time-stepping methods. Finite difference methods for the Black-Scholes equation. Discretization, stability, convergence. Methods for portfolio optimization, effect of data errors on portfolio weights. ",
        "note": "Lab is not scheduled and students are expected to find time in open hours to complete their work. Students who receive a good grade in CS 335 may contact the instructor of CS 476 to seek admission without the formal prerequisites. Offered: W]",
        "pre_reqs": "(AMATH 242/CS 371 or CS 370) and STAT 231/241",
        "anti_reqs": "...",
        "co_reqs": "...",
        "unit":"0.50",
    }...
}
```

# License
Copyright (C) 2022 Tristan Simpson <heytristaann@gmail.com>

This file is part of the University of Waterloo Course Catalog project.

The University of Waterloo Course Catalog project can not be copied and/or distributed without the express permission of Tristan Simpson <heytristaann@gmail.com>. 
