<script lang="ts">
	import * as Card from '$lib/components/ui/card/index';
	import type { Snippet } from 'svelte';
	import { Logo } from '$lib/components/contents/logo';
	import SSOButton from './ssoButton.svelte';

	type AuthFormProps = {
		formTitle?: String;
		formBody: Snippet;
		switchFormBtn?: Snippet;
		showSSOButton?: boolean;
		showLogo?: boolean;
	};

	let {
		formTitle,
		formBody,
		switchFormBtn,
		showSSOButton = false,
		showLogo = false
	}: AuthFormProps = $props();
</script>

<Card.Root class="-my-4 w-full max-w-sm py-10 shadow-none ring-0 md:shadow-xs md:ring-1 ">
	<Card.Header class="w-full justify-center gap-1 px-10">
		<Logo withTitle withSubTitle class="h-16"></Logo>
		<Card.Title class="text-center">{formTitle}</Card.Title>
	</Card.Header>
	<Card.Content class="px-10">
		{@render formBody()}
	</Card.Content>
	<Card.Footer class="flex-col gap-4 px-10">
		{#if switchFormBtn}
			{@render switchFormBtn()}
		{/if}
		{#if showSSOButton}
			<span class="text-neutral-300">atau</span>
			<SSOButton></SSOButton>
		{/if}
	</Card.Footer>
</Card.Root>
