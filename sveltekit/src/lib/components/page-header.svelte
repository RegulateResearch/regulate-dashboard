<script lang="ts">
	import { cn } from '$lib/utils';
	import type { Snippet } from 'svelte';
	import type { HTMLAttributes } from 'svelte/elements';

	let {
		header,
		actionBtn,
		class: className,
		...restProps
	}: HTMLAttributes<HTMLDivElement> & { header?: Snippet; actionBtn?: Snippet } = $props();
</script>

<div
	class={cn('flex items-center justify-between gap-2 rounded-xl border p-4 h-16', className)}
	{...restProps}
>
	<div class="flex items-center gap-2">
		{#if header}
			{@render header()}
		{/if}
	</div>
	{#if actionBtn}
		{@render actionBtn()}
	{/if}
</div>
