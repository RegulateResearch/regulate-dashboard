<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import SquarePenIcon from '@lucide/svelte/icons/square-pen';
	import ExternalLinkIcon from '@lucide/svelte/icons/external-link';
	import TrashIcon from '@lucide/svelte/icons/trash-2';
	import CopyIcon from '@lucide/svelte/icons/copy';
	import type { CourseWithId } from '$lib/schema';
	import { getCreateCourseModalState, getDeleteCourseModalState } from './modal-state.svelte.ts';

	let { data }: { data: CourseWithId } = $props();
	let deleteCourseModalState = getDeleteCourseModalState();
	let createCourseModalState = getCreateCourseModalState();
</script>

<div class="flex justify-center gap-1">
	<Button type="button" variant="ghost" size="icon-sm" href={`/admin/courses/${data.id}`}
		><SquarePenIcon class="text-yellow-500"></SquarePenIcon></Button
	>
	<Button
		type="button"
		variant="ghost"
		size="icon-sm"
		href={data.url}
		target="_blank"
		disabled={!data.url}><ExternalLinkIcon class="text-sky-500"></ExternalLinkIcon></Button
	>
	<Button
		type="button"
		variant="ghost"
		size="icon-sm"
		onclick={() => deleteCourseModalState.open(data.id)}
		><TrashIcon class="text-red-500"></TrashIcon></Button
	>
	<Button
		type="button"
		variant="ghost"
		size="icon-sm"
		onclick={() =>
			createCourseModalState.open({
				name: `${data.name} (Copy)`,
				year: data.year.toString(),
				term: data.term,
				url: data.url
			})}><CopyIcon class="text-purple-500"></CopyIcon></Button
	>
</div>
