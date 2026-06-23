<script lang="ts">
	import { RegulateLogo } from '$lib/components/contents/logo/index';

	let { children } = $props();
</script>

<div class="flex min-h-svh flex-col items-center justify-center gap-6 bg-muted p-6 md:p-10">
	<div class="flex w-full max-w-sm flex-col gap-6">
		<a href="##" class="flex items-center justify-center gap-2 self-center font-medium">
			<RegulateLogo withTitle class="h-auto w-1/2" />
		</a>
		{@render children()}
	</div>
</div>
