<script lang="ts">
	import AuthForm from '$lib/components/auth-form.svelte';
	import { Field, FieldDescription } from '$lib/components/ui/field';
	import type { PageProps } from './$types';
	import { resolve } from '$app/paths';
	import { UILogo } from '$lib/components/contents/logo';

	let { data }: PageProps = $props();
</script>

<AuthForm formTitle="Masuk dengan Single Sign On Universitas Indonesia" {formBody} {switchFormBtn} />

{#snippet formBody()}
	<div class="flex flex-col items-center justify-center gap-4">
		<UILogo class="h-auto w-1/4" />
		{#if data.error}
			<p class="text-center text-red-500">
				Gagal menyambungkan dengan layanan SSO Universitas Indonesia. <br /> Mohon coba lagi nanti
				atau hubungi administrator jika masalah ini terus berlanjut. <br /> Error: {data.error}
			</p>
		{:else}
			<p class="text-center">
				Menghubungkan dengan layanan SSO Universitas Indonesia... <br /> Jika proses ini memakan waktu
				terlalu lama, pastikan Anda memiliki koneksi internet yang stabil dan coba lagi.
			</p>
		{/if}
	</div>
{/snippet}

{#snippet switchFormBtn()}
	<Field class="p-2">
		<FieldDescription class="text-center">
			Masuk dengan cara lain? <a href={resolve('/login')}>Masuk</a>
		</FieldDescription>
	</Field>
{/snippet}
