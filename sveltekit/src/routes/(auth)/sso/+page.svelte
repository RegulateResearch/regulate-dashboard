<script lang="ts">
	import AuthForm from '../authForm.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<AuthForm {formBody} />

{#snippet formBody()}
	<div class="flex flex-col items-center justify-center gap-4">
		<enhanced:img src="$lib/assets/ui-logo.png" alt="Logo UI" height="64" />
		{#if data.error}
			<p class="text-center text-red-500">
				Gagal menyambungkan dengan layanan SSO Universitas Indonesia. <br /> Mohon coba lagi nanti atau hubungi administrator jika masalah ini terus berlanjut. <br /> Error: {data.error}
			</p>
		{:else}
			<p class="text-center">
				Menghubungkan dengan layanan SSO Universitas Indonesia... <br /> Jika proses ini memakan waktu terlalu lama, pastikan Anda memiliki koneksi internet yang stabil dan coba lagi.
			</p>
		{/if}
	</div>
{/snippet}
