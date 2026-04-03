<script lang="ts">
	import DoorOpen from '@lucide/svelte/icons/door-open';
	import { Button } from '$lib/components/ui/button';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import { SRL } from '$lib/components/contents/srl';

	let { children } = $props();
</script>

<Tooltip.Root>
	<Tooltip.Trigger class="fixed top-4 left-4">
		<Button variant="ghost" class="h-12 w-12" href="/">
			<DoorOpen class="size-6 text-neutral-600" />
		</Button>
	</Tooltip.Trigger>
	<Tooltip.Content collisionPadding={16}>
		<p>Kembali ke Halaman Utama</p>
	</Tooltip.Content>
</Tooltip.Root>

<section
	class="flex h-screen w-full flex-row items-center justify-center gap-16 lg:gap-32 lg:px-16"
>
	<div class="regulate-logo-content hidden lg:block">
		<SRL class="w-0.8 max-h-0.8 h-100"></SRL>
	</div>
	{@render children()}
</section>
