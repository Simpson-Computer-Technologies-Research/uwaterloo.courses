<script>
	export let course;

    // Returns an empty space character multiplied
    // by the amount paramater
    function emptySpace(amount) {
        return "‏‏‎ ‎".repeat(amount)
    }

</script>

<div class="class_div">
    <!-- Course Title -->
    <div class="course_div" style="display: flex; align-items: center;">
        <img 
            src="./images/waterloo_logo.png" 
            alt="" 
            width="40"
            height="40"
            style="margin-right: 10px;"
        >
        <p>
            <strong>{course.title}: {course.name}</strong>
        </p>
    </div>

    <!-- Course Description -->
    <div class="course_div">
        <p style="font-size: 14.5px;">
            {course.description}
        </p>
    </div>

    <!-- Course Pre-Requisites -->
    <div class="course_div" style="font-size: 13px; margin-top: 1px;">
        <p>
            <strong>Pre Requisites </strong>{course.pre_requisites}
        </p>
    </div>

    <!-- Course Anti-Requisites -->
    <div class="course_div" style="font-size: 13px;">
        <p>
            <strong>Anti Requisites </strong>{course["anti_requisites"]}
        </p>
    </div>

    <!-- Official Course Website -->
    {#if course.title != undefined}
    {@const splitTitle = course.title.split(" ")}
        <div style="font-size: 12px;">
            <a 
                href="https://ucalendar.uwaterloo.ca/2223/COURSE/course-{splitTitle[0]}.html#{splitTitle[0]}{splitTitle[1]}"
                target="_blank"
                rel="noreferrer noopener"
            >
                Official Website
                {emptySpace(4)} Components {course.components}
                {emptySpace(4)} Unit {course.unit}
                {emptySpace(4)} ID {course.id}
            </a>
        </div>
    {/if}
</div>

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