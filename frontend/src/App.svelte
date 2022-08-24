<script>
	// Component Imports
	import CourseInfo from './components/CourseInfo.svelte';
	import SiteHeader from './components/SiteHeader.svelte';

	// Track input last key up
	var queryTimer;

	// The time it took to query the subject
	let queryTime = 0;

	// The subject querying for
	let querySubject = "";

	// Track the amount of results
	let queryResultAmount = 0;

	// Hold the subject data
	let queryResult = [];

	// The _QuerySubjectData() function is used
	// to send the http request to the localhost api
	function _QuerySubjectData(query) {
		// Create a query starting time
		let startTime = Date.now();

		// Send the http request to the golang api
		fetch("http://127.0.0.1:8000/courses?q=" + query)
			.then((response) => response.json())
			.then((data) => {

				// Set the query data
				if (data == null) {
					queryResult = [];
					queryResultAmount = 0;
				} 
				// Data is not null
				else {
					queryResult = data;
					queryResultAmount = data.length;
				}
				// Set the query time variable
				queryTime = Date.now() - startTime;
			})
	}

	// The QuerySubjectData() function is called when the
	// user types a character into the course query bar
	//
	// It sets a timeout to prevent spam calling the api
	function QuerySubjectData(query) {
		querySubject = query;

		// Clear previous timer
		clearTimeout(queryTimer);

		// Fetch the courses if query length
		// is greater than 3
		if (query.length >= 3 && query.length <= 40) {
			queryTimer = setTimeout(_QuerySubjectData, 220, query);
		} else {
			// Reset query data and query time
			queryResult = [];
			queryTime = 0;
		}
	}
</script>

<main>
	<!-- When the user clicks the @code:cs -->
	<SiteHeader handleCodeClick={() => QuerySubjectData("@code:cs")}/>

	<!-- Input course to search for -->
	<div>
		<!-- svelte-ignore a11y-autofocus -->
		<input
			autofocus
			value={querySubject}
			placeholder="Search"
			class="course_input" 
			on:keyup={({ target: { value } }) => QuerySubjectData(value)} 
		/>
	</div>

	<!-- Result header -->
	<div class="result_div">
		<h3 class="result_header">
			{#if querySubject.length >= 3}
				{queryResultAmount} 
			{:else}
				0
			{/if}
				results 
			{#if querySubject.length > 0}
				for 
			{/if}
				{querySubject} in {queryTime}ms
		</h3>
	</div>
	
	<!-- List of courses and their info -->
	<!-- svelte-ignore empty-block -->
	{#each queryResult as course}
		<CourseInfo course={course}/>
	{/each}
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