<script lang="ts">
	import { cn } from '$lib/utils';
	import type { HTMLAttributes } from 'svelte/elements';

	let { label, class: className }: HTMLAttributes<HTMLDivElement> & { label: string } = $props();
</script>

<div class={cn('flex space-x-2 w-full min-w-125', className)}>
	<span class="w-full truncate">
		{label}
	</span>
</div>
