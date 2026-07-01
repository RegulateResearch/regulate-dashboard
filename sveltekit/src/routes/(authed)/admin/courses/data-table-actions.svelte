<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import SquarePenIcon from '@lucide/svelte/icons/square-pen';
	import ExternalLinkIcon from '@lucide/svelte/icons/external-link';
	import TrashIcon from '@lucide/svelte/icons/trash-2';
	import CopyIcon from '@lucide/svelte/icons/copy';

	let { id, url }: { id?: number; url?: string } = $props();
</script>

<div class="flex justify-center gap-1">
	<Button type="button" variant="ghost" size="icon-sm" href={`/admin/courses/${id}`}
		><SquarePenIcon class="text-yellow-500"></SquarePenIcon></Button
	>
	<Button type="button" variant="ghost" size="icon-sm" href={url} target="_blank" disabled={!url}
		><ExternalLinkIcon class="text-sky-500"></ExternalLinkIcon></Button
	>
	<Button type="button" variant="ghost" size="icon-sm"
		><TrashIcon class="text-red-500"></TrashIcon></Button
	>
	<Button type="button" variant="ghost" size="icon-sm"
		><CopyIcon class="text-purple-500"></CopyIcon></Button
	>
</div>
