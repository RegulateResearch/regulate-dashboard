<script lang="ts">
	import type { ComponentProps } from 'svelte';
	import ChevronUpDownIcon from '@lucide/svelte/icons/chevrons-up-down';
	import { Button } from '$lib/components/ui/button/index.js';

	let {
		label,
		variant = 'ghost',
		size ='xs',
		class: className,
		...restProps
	}: ComponentProps<typeof Button> & { label: string } = $props();
</script>

<div class={className}>
	<Button {variant} {size} {...restProps}>
		{label}
		<ChevronUpDownIcon class="ms-2" />
	</Button>
</div>
