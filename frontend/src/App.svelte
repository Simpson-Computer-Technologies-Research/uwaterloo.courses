<script>
	import CourseInfo from './components/CourseInfo.svelte';
	import SiteHeader from './components/SiteHeader.svelte';

	// Promise for query http requests
	let promise = Promise.resolve([]);

	// The subject querying for
	let subject_query = "";

	// The time it took to query the subject
	let query_time = 0;

	// The fetchCourses() function returns the response json
	// from the golang api. The golang api scrapes/grabs the
	// course information from the waterloo website/redis database
	async function fetchCourses(subject) {
		// Create a query starting time
		let startTime = Date.now();

		// Send the http request to the golang api
		const response = await self.fetch("http://127.0.0.1:8000/courses?q=" + subject);

		// Set the query time variable
		query_time = Date.now() - startTime

		// Return the response json
		if (response.ok) {
  			return response.json();
		}
	}

	// Handle the course input on key up
	function courseInputDebounce(query) {
		subject_query = query;

		// Fetch the courses if query length
		// is greater than 3
		if (query.length >= 3) {
			promise = fetchCourses(query);
		}
	}
</script>

<main>
	<SiteHeader/>

	<!-- Input course to search for -->
	<div>
		<input 
			class="course_input" 
			on:keyup={({ target: { value } }) => courseInputDebounce(value)} 
		/>
	</div>
	
	<!-- svelte-ignore empty-block -->
	{#await promise}
		{:then courses}

		<!-- Result header -->
		{#if subject_query.length > 0}
			<div class="result_div">
				<h3 class="result_header">
					{courses.length} results for {subject_query} in {query_time}ms
				</h3>
			</div>
		{/if}
		
		<!-- List of courses and their info -->
	  	{#each courses as course}
		  	<CourseInfo course={course}/>
		{/each}
	{/await}
</main>

<style global lang="postcss">
	@tailwind base;
	@tailwind components;
	@tailwind utilities;

	main {
		width: 92%;
		padding: 1em;
		margin: 0 auto;
	}

	.result_div {
		color: #969696;
	}

	.result_header {
		font-weight: 300;
		border-radius: 3px;
		margin-left: 1.2%;
		background-color: #d5ffdf;
		padding: 0.5%;
		padding-left: 10px;
		font-size: 15px;
		margin-right: 60%;
	}

	.course_input {
		margin-left: 1.2%;
		margin-bottom: -0.05%;
		outline: none;
		border-top: 0;
		border-right: 0;
		border-left: 0;
		border-bottom-width: 2px;
		border-bottom-color: #969696;
		background-color: #f9f9f9;
		border-radius: 3px;
		height: 40px;
		width: 45%;
		color:#525252;
	}

	:root::-webkit-scrollbar {
		width: 20px;
		height: 20px;
	}

	/* Track */
	:root::-webkit-scrollbar-track {
		background: #f1f1f1; 
	}
	
	/* Handle */
	:root::-webkit-scrollbar-thumb {
		background: #6366f1;
	}

	/* Handle on hover */
	:root::-webkit-scrollbar-thumb:hover {
		background: #474af2;
	}
</style>