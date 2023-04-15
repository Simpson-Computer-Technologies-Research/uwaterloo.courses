<script lang="ts">
	import VirtualList from './components/VirtualList.svelte';
	import CourseInfo from './components/CourseInfo.svelte';
	import SiteHeader from './components/SiteHeader.svelte';
  import { select_value } from 'svelte/internal';

	// Query variables
	let queryTime: number = 0;
	let querySubject: string = "";
	let queryResultAmount: number = 0;
	let queryResult: any[] = [];

	// The querySubjectData() function is used to send the http 
	// request to the backend api and fetch all of the courses
	function querySubjectData(query: string): void {
		// If the query is less than 3 characters
		if (query.length < 3) {
			[queryTime, queryResult, queryResultAmount] = [0, [], 0];
			return;
		}

		// Send the http request
		fetch("http://127.0.0.1:8000/courses?q=" + query)
			.then((response) => response.json())
			.then((data) => {
				if (data == null) {
					[queryTime, queryResult, queryResultAmount] = [0, [], 0];
				} else {
					[queryTime, queryResult, queryResultAmount] = [data.time / 1000, data.result, data.result.length];
				}
			})
			.catch((error) => console.log(error));
	}

	// The onSearch() function is used to update the
	// querySubject variable when the user types in the search bar
	// and then call the querySubjectData() function to send the http request
	// to the backend api.
	function onSearch(event: any): void {
		querySubject = event.target.value;
		querySubjectData(querySubject);
	}

	// Method to handle the header click.
	// This method is passed to the SiteHeader component
	function handleHeaderClick(query: string): void {
		querySubject = query;
    querySubjectData(query);
	}
</script>

<main style="width: 92%; padding: 1em; margin: 0 auto;">
	<!-- When the user clicks the @code:cs -->
	<SiteHeader handleHeaderClick={handleHeaderClick}/>

	<!-- Input course to search for -->
	<div>
		<!-- svelte-ignore a11y-autofocus -->
		<input autofocus value={querySubject} placeholder="Search" class="course_input" 
			on:keyup={(event) => onSearch(event)}
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