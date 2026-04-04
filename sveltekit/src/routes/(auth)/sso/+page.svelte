<script lang="ts">
	import AuthForm from '../authForm.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<AuthForm {formBody} />

{#snippet formBody()}
	<div class="flex flex-col items-center justify-center gap-4">
		<enhanced:img src="$lib/assets/ui-logo.png" alt="Logo UI" height="64" />
		<p class="text-center">
			Sedang menyambungkan dengan layanan SSO Universitas Indoensia. <br /> Mohon menunggu.
		</p>
		<code class="text-sm text-gray-500">
			{data?.message}
		</code>
	</div>
{/snippet}
