<script>
	export let course;

    // Returns an empty space character multiplied
    // by the amount paramater
    function emptySpace(amount) {
        return "‏‏‎ ‎".repeat(amount)
    }

</script>

<main>
    <div class="class_div">
        <!-- Course Title -->
        <div class="course_div" style="display: flex; align-items: center;">
            <img 
                src="./static/images/waterloo_logo.png" 
                alt="" 
                width="40"
                height="40"
                style="margin-right: 10px;"
            >
            <p>
                <strong>{course.Title}: {course.Name}</strong>
            </p>
        </div>

        <!-- Course Description -->
        <div class="course_div">
            <p style="font-size: 14.5px;">
                {course.Description}
            </p>
        </div>

        <!-- Course Pre-Requisites -->
        <div class="course_div" style="font-size: 13px; margin-top: 1px;">
            <p>
                <strong>Pre Requisites </strong>{course["Pre Requisites"]}
            </p>
        </div>

        <!-- Course Anti-Requisites -->
        <div class="course_div" style="font-size: 13px;">
            <p>
                <strong>Anti Requisites </strong>{course["Anti Requisites"]}
            </p>
        </div>

        <!-- Official Course Website -->
        <div style="font-size: 12px;">
            <a 
                href="https://ucalendar.uwaterloo.ca/2223/COURSE/course-{course.Title.split(" ")[0]}.html#{course.Title.split(" ")[0]}{course.Title.split(" ")[1]}"
                target="_blank"
                rel="noreferrer noopener"
            >
                Official Website
                {emptySpace(4)} Components {course.Components}
                {emptySpace(4)} Unit {course.Unit}
                {emptySpace(4)} ID {course.ID}
            </a>
        </div>
    </div>
</main>

<style>
	.course_div {
		margin-top: -1%;
		padding: 0.0011%;
	}

    .class_div {
        margin-top: 1%;
        margin-left: 1.2%;
        margin-bottom: 2%; 
        padding: 1%; 
        background-color: #f9f9f9; 
        margin-right: 30%;
        padding-right: 2%;
        border-radius: 5px;
    }
</style>