<script>
	import CourseInfo from './components/CourseInfo.svelte';
	import SiteHeader from './components/SiteHeader.svelte';

	// Used for pasting the @code:cs into the search
	// query when it is clicked on
	let handleCodeClick = () => subject_query = "@code:cs";

	// Promise for query http requests
	let promise = Promise.resolve([]);

	// The subject querying for
	let subject_query = "";

	// Track the amount of results
	let query_result_amount = 0;

	// The time it took to query the subject
	let query_time = 0;

	// The fetchCourses() function returns the response json
	// from the golang api. The golang api scrapes/grabs the
	// course information from the waterloo website/redis database
	async function fetchCourses(query) {
		// Create a query starting time
		let startTime = Date.now();

		// Send the http request to the golang api
		await self.fetch("http://127.0.0.1:8000/courses?q=" + query)
			.then((response) => response.json())
			.then((data) => {
				// Set the query result amount
				query_result_amount = data.length;

				// Set the query time variable
				query_time = Date.now() - startTime;

				// Set the promise to the data
				promise = data;
			})
	}

	// Handle the course input on key up
	function courseInputDebounce(query) {
		subject_query = query;

		// Fetch the courses if query length
		// is greater than 3
		if (query.length >= 3 && query.length <= 40) {
			fetchCourses(query);
		} else {
			// Reset course list and query time
			promise = Promise.resolve([]);
			query_time = 0;
		}
	}
</script>

<main>
	<SiteHeader handleCodeClick={handleCodeClick}/>

	<!-- Input course to search for -->
	<div>
		<!-- svelte-ignore a11y-autofocus -->
		<input
			autofocus
			value={subject_query}
			placeholder="Search"
			class="course_input" 
			on:keyup={({ target: { value } }) => courseInputDebounce(value)} 
		/>
	</div>

	<!-- Result header -->
	<div class="result_div">
		<h3 class="result_header">
			{#if subject_query.length >= 3}
				{query_result_amount} 
			{:else}
				0
			{/if}
				results 
			{#if subject_query.length > 0}
				for 
			{/if}
				{subject_query} in {query_time}ms
		</h3>
	</div>
	
	<!-- List of courses and their info -->
	<!-- svelte-ignore empty-block -->
	{#await promise}
		{:then courses}
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
		margin-left: 1%;
		padding: 0.5%;
		padding-left: 10px;
		font-size: 15px;
		margin-right: 60%;
	}

	.course_input {
		margin-left: 1.2%;
		margin-bottom: -1%;
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