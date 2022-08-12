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


# License
Copyright (C) 2022 Tristan Simpson <heytristaann@gmail.com>

This file is part of the University of Waterloo Course Catalog project.

The University of Waterloo Course Catalog project can not be copied and/or distributed without the express permission of Tristan Simpson <heytristaann@gmail.com>. 
