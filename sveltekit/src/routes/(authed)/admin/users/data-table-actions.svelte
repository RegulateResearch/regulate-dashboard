<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import SquarePenIcon from '@lucide/svelte/icons/square-pen';
	import type { User } from './columns';

	let { data }: { data: User } = $props();
</script>

<div class="flex justify-center gap-1">
	<Button type="button" variant="ghost" size="icon-sm" href={`/admin/users/${data.id}`}
		><SquarePenIcon class="text-yellow-500"></SquarePenIcon></Button
	>
</div>
