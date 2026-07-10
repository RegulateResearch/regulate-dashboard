<script lang="ts">
	import * as Card from '$lib/components/ui/card/index.js';
	import { FieldDescription } from '$lib/components/ui/field/index.js';
	import { cn } from '$lib/utils.js';
	import type { HTMLAttributes } from 'svelte/elements';
	import type { Snippet } from 'svelte';

	interface PropsType extends HTMLAttributes<HTMLDivElement> {
		formTitle: string;
		formBody: Snippet;
		switchFormBtn?: Snippet;
	}

	let { class: className, formTitle, formBody, switchFormBtn, ...restProps }: PropsType = $props();
</script>

<div class={cn('flex flex-col gap-6', className)} {...restProps}>
	<Card.Root>
		<Card.Header class="text-center">
			<Card.Title class="text-lg">{formTitle}</Card.Title>
		</Card.Header>
		<Card.Content>
			{@render formBody()}
			{#if switchFormBtn}
				{@render switchFormBtn()}
			{/if}
		</Card.Content>
	</Card.Root>
	<FieldDescription class="px-6 text-center">
		Aplikasi ini dibuat dalam rangka penelitian. Data Anda akan digunakan untuk keperluan penelitian
		dan tidak akan dibagikan kepada pihak ketiga. Segala bentuk pertanyaan atau kekhawatiran terkait
		privasi dapat disampaikan melalui email ke
		<a href="mailto:jahns.michael51@ui.ac.id">jahns.michael51@ui.ac.id</a>
	</FieldDescription>
</div>
