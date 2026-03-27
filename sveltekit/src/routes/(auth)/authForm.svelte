<script lang="ts">
	import * as Card from '$lib/components/ui/card/index';
	import type { Snippet } from 'svelte';
	import { Button } from '$lib/components/ui/button';
	import { Logo } from '$lib/components/ui/logo';

	type AuthFormContent = {
		formTitle: String;
		formBody: Snippet;
		switchFormBtn: Snippet;
	};

	let { formTitle, formBody, switchFormBtn }: AuthFormContent = $props();

	const handleClick = async () => {
		const resp = await popUpLogin();
		console.log('RESP', resp);
	};

	const popUpLogin = () => {
		const SSOWindow = window.open(
			new URL(
				'https://sso.ui.ac.id/cas2/login?service=http%3A%2F%2Flocalhost%3A5173%2Fsso'
			).toString(),
			'SSO UI Login',
			'left=50, top=50, width=480, height=480'
		);

		return new Promise(function (resolve, reject) {
			window.addEventListener(
				'message',
				(e) => {
					if (SSOWindow) {
						SSOWindow.close();
					}
					const data = e.data;
					resolve(data);
				},
				{ once: true }
			);
		});
	};
</script>

<Card.Root class="-my-4 w-full max-w-sm py-10 shadow-none ring-0 md:shadow-xs md:ring-1 ">
	<Card.Header class="w-full justify-center gap-1 px-10">
		<Logo withTitle withSubTitle class="w-48"></Logo>
		<Card.Title class="text-center">{formTitle}</Card.Title>
	</Card.Header>
	<Card.Content class="px-10">
		{@render formBody()}
	</Card.Content>
	<Card.Footer class="flex-col gap-4 px-10">
		{@render switchFormBtn()}
		<span class="text-neutral-300">atau</span>
		<Button variant="outline" class="w-full" onclick={handleClick}>
			<enhanced:img src="$lib/assets/ui-logo.png" alt="Logo UI" height="24" />
			Masuk dengan SSO UI
		</Button>
	</Card.Footer>
</Card.Root>
