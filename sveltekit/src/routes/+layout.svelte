<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { navigating } from '$app/state';
	import Loading from '$lib/components/ui/loading/loading.svelte';
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

{#if navigating.to}
	<Loading></Loading>
{/if}

<Tooltip.Provider>
	{@render children()}
</Tooltip.Provider>
