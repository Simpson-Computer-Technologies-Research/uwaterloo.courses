<script lang="ts">
	import { onMount, tick } from 'svelte';
	// props
	export let items: any;
	export let height: any = "100%";
	export let itemHeight: any = undefined;
	// read-only, but visible to consumers via bind:start
	export let start: number = 0;
	export let end: number = 0;
	// local state
	let height_map: any[] = [];
	let rows: any;
	let viewport: any;
	let contents: any;
	let viewport_height: number = 0;
	let visible: any[];
	let mounted: boolean;
	let top: number = 0;
	let bottom: number = 0;
	let average_height: number;
	$: visible = items.slice(start, end).map((data: any, i: number) => {
		return { index: i + start, data };
	});
	// whenever `items` changes, invalidate the current heightmap
	$: if (mounted) refresh(items, viewport_height, itemHeight);
	async function refresh(items: any[], viewport_height: number, itemHeight: number) {
		const { scrollTop } = viewport;
		await tick(); // wait until the DOM is up to date
		let content_height = top - scrollTop;
		let i = start;
		while (content_height < viewport_height && i < items.length) {
			let row = rows[i - start];
			if (!row) {
				end = i + 1;
				await tick(); // render the newly visible row
				row = rows[i - start];
			}
			const row_height = height_map[i] = itemHeight || row.offsetHeight;
			content_height += row_height;
			i += 1;
		}
		end = i;
		const remaining = items.length - end;
		average_height = (top + content_height) / end;
		bottom = remaining * average_height;
		height_map.length = items.length;
	}
	async function handle_scroll() {
		const { scrollTop } = viewport;
		const old_start = start;
		for (let v = 0; v < rows.length; v += 1) {
			height_map[start + v] = itemHeight || rows[v].offsetHeight;
		}
		let i = 0;
		let y = 0;
		while (i < items.length) {
			const row_height = height_map[i] || average_height;
			if (y + row_height > scrollTop) {
				start = i;
				top = y;
				break;
			}
			y += row_height;
			i += 1;
		}
		while (i < items.length) {
			y += height_map[i] || average_height;
			i += 1;
			if (y > scrollTop + viewport_height) break;
		}
		end = i;
		const remaining = items.length - end;
		average_height = y / end;
		while (i < items.length) height_map[i++] = average_height;
		bottom = remaining * average_height;
		// prevent jumping if we scrolled up into unknown territory
		if (start < old_start) {
			await tick();
			let expected_height = 0;
			let actual_height = 0;
			for (let i = start; i < old_start; i +=1) {
				if (rows[i - start]) {
					expected_height += height_map[i];
					actual_height += itemHeight || rows[i - start].offsetHeight;
				}
			}
			const d = actual_height - expected_height;
			viewport.scrollTo(0, scrollTop + d);
		}
		// TODO if we overestimated the space these
		// rows would occupy we may need to add some
		// more. maybe we can just call handle_scroll again?
	}
	// trigger initial refresh
	onMount(() => {
		rows = contents.getElementsByTagName('svelte-virtual-list-row');
		mounted = true;
	});
</script>

<style>
	svelte-virtual-list-viewport {
		position: relative;
		overflow-y: auto;
		-webkit-overflow-scrolling: none;
		display: block;
	}
	svelte-virtual-list-viewport::-webkit-scrollbar { width: 20px; }
	svelte-virtual-list-viewport::-webkit-scrollbar-track { background: #f1f1f1; }
	svelte-virtual-list-viewport::-webkit-scrollbar-thumb { background: #6366f1; }
	svelte-virtual-list-viewport::-webkit-scrollbar-thumb:hover { background: #474af2; }
	svelte-virtual-list-contents, svelte-virtual-list-row {
		display: block;
	}
	svelte-virtual-list-row {
		overflow: hidden;
	}
</style>

<svelte-virtual-list-viewport
	bind:this={viewport}
	bind:offsetHeight={viewport_height}
	on:scroll={handle_scroll}
	style="height: {height};"
>
	<svelte-virtual-list-contents
		bind:this={contents}
		style="padding-top: {top}px; padding-bottom: {bottom}px;"
	>
		{#each visible as row (row.index)}
			<svelte-virtual-list-row>
				<slot item={row.data}>Missing template</slot>
			</svelte-virtual-list-row>
		{/each}
	</svelte-virtual-list-contents>
</svelte-virtual-list-viewport>