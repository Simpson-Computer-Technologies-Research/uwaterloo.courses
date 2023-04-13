<script>
	import VirtualList from './components/VirtualList.svelte';
	import CourseInfo from './components/CourseInfo.svelte';
	import SiteHeader from './components/SiteHeader.svelte';

	// The time it took to query the subject
	let queryTime = 0;

	// The subject querying for
	let querySubject = "";

	// Track the amount of results
	let queryResultAmount = 0;

	// Hold the subject data
	let queryResult = [];

	// The QuerySubjectData() function is used to send the http 
	// request to the backend api and fetch all of the courses
	function QuerySubjectData(query) {
		// If the query is less than 3 characters
		if (query.length < 3) {
			queryResult = [];
			queryResultAmount = 0;
			return;
		}

		// Send the http request
		fetch("http://127.0.0.1:8000/courses?q=" + query)
			.then((response) => response.json())
			.then((data) => {
				// Data is null
				if (data == null) {
					queryResult = [];
					queryResultAmount = 0;
					return;
				} 

				// Data is not null
				queryTime = data.time / 1000;
				queryResult = data.result;
				queryResultAmount = data.result.length;
			})
	}
</script>

<main style="width: 92%; padding: 1em; margin: 0 auto;">
	<!-- When the user clicks the @code:cs -->
	<SiteHeader handleHeaderClick={(query) => {
		QuerySubjectData(query);
		querySubject = query;
	}}/>

	<!-- Input course to search for -->
	<div>
		<!-- svelte-ignore a11y-autofocus -->
		<input
			autofocus
			value={querySubject}
			placeholder="Search"
			class="course_input" 
			on:keyup={({ target: { value } }) => {
				querySubject = value;
				QuerySubjectData(value);
			}} 
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

	<!-- The list of courses -->
	<div class="virtualList" style="width: 100%; height: 500px;">
		<VirtualList items={queryResult} let:item>
			<CourseInfo course={item}/>
		  </VirtualList>
	</div>
</main>

<style>
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

	/* Scroll bar styling */
	:root::-webkit-scrollbar { display: none; }
</style>