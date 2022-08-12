import requests


r = requests.get("https://ucalendar.uwaterloo.ca/2021/COURSE/course-CS.html")
print(r.text)