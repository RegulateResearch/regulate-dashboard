<script lang="ts">
	import type { HTMLTableAttributes } from 'svelte/elements';
	import { cn, type WithElementRef } from '$lib/utils.js';
	import type { ClassValue } from 'clsx';

	let {
		ref = $bindable(null),
		class: className,
		wrapperClass,
		children,
		...restProps
	}: WithElementRef<HTMLTableAttributes> & { wrapperClass?: ClassValue } = $props();
</script>

<div data-slot="table-container" class={cn('relative w-full', wrapperClass)}>
	<table
		bind:this={ref}
		data-slot="table"
		class={cn('w-full caption-bottom text-sm', className)}
		{...restProps}
	>
		{@render children?.()}
	</table>
</div>
