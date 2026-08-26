<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import SquarePenIcon from '@lucide/svelte/icons/square-pen';
	import type { User } from './columns';
	import { getEditRoleModalState } from './modal-state.svelte';

	let { data }: { data: User } = $props();
	let editRoleModalState = getEditRoleModalState();
</script>

<div class="flex justify-center gap-1">
	<Button
		type="button"
		variant="ghost"
		size="icon-sm"
		onclick={() =>
			editRoleModalState.open(data)}
		><SquarePenIcon class="text-yellow-500"></SquarePenIcon>
	</Button>
</div>
