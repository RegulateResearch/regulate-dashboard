<script lang="ts">
	import LayoutDashboardIcon from '@lucide/svelte/icons/layout-dashboard';
	import BlocksIcon from '@lucide/svelte/icons/blocks';
	import * as Empty from '$lib/components/ui/empty';
	import PageHeader from '$lib/components/page-header.svelte';
	import RoundedTabs from '$lib/components/ui/rounded-tabs/rounded-tabs.svelte';
</script>

<PageHeader class="shrink-0">
	{#snippet header()}
		<LayoutDashboardIcon class="fill-yellow-400 text-6xl" />
		<h1 class="text-xl">Dashboard</h1>
	{/snippet}
</PageHeader>

<RoundedTabs class="min-h-0 grow flex justify-center">
	<Empty.Root class="h-full">
		<Empty.Header>
			<Empty.Media variant="icon">
				<BlocksIcon />
			</Empty.Media>
			<Empty.Title>Fitur masih dalam pengembangan</Empty.Title>
			<Empty.Description>
				Tidak ada yang ditampilkan di sini. Fitur masih dalam tahap pengembangan.
			</Empty.Description>
		</Empty.Header>
	</Empty.Root>
</RoundedTabs>