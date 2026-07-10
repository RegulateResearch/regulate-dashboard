<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { navigating } from '$app/state';
	import Loading from '$lib/components/ui/loading/loading.svelte';
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';
	import { page } from '$app/state';

	let { children } = $props();

	let pageTitle = $derived(
		page.data.pageTitle ? `${page.data.pageTitle} | Regulate Dashboard` : 'Regulate Dashboard'
	);
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>{pageTitle}</title>
</svelte:head>

<Tooltip.Provider>
	{@render children()}
</Tooltip.Provider>

{#if navigating.to}
	<Loading></Loading>
{/if}
