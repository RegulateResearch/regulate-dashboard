<script lang="ts">
	import { page } from '$app/state';
	import { Button } from '$lib/components/ui/button';
	import { RegulateLogo } from '$lib/components/contents/logo/index';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import DoorOpen from '@lucide/svelte/icons/door-open';

	const message = $derived(
		page.status === 404
			? { title: 'Halaman tidak ditemukan.', body: 'Periksa kembali alamat URL Anda.' }
			: {
					title: 'Terjadi kesalahan.',
					body: 'Tim kami sedang memperbaikinya. Coba kembali beberapa saat lagi.'
				}
	);
</script>

<div class="fixed top-0 right-0 bottom-0 left-0 flex items-center justify-center bg-yellow-50/25">
	<div class="flex flex-row items-center justify-center gap-6">
		<RegulateLogo active={false} class="w-20" />
		<div class="flex max-w-3xl flex-col">
			<h1 class="text-2xl text-neutral-500">{message.title}</h1>
			<p>
				{message.body}
				<br />
			</p>
			<p class="text-xs">
				Kembali ke
				<span>
					<Button variant="link" class="h-auto border-0 p-0 text-xs" href="/">Halaman Utama</Button>
				</span>
			</p>
		</div>
	</div>
</div>

<Tooltip.Root>
	<Tooltip.Trigger class="fixed top-4 left-4">
		<Button variant="ghost" class="h-12 w-12" href="/">
			<DoorOpen class="text-neutral-600 size-6" />
		</Button>
	</Tooltip.Trigger>
	<Tooltip.Content collisionPadding={16}>
		<p>Kembali ke Halaman Utama</p>
	</Tooltip.Content>
</Tooltip.Root>
