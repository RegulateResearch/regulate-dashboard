<script lang="ts">
	import PageHeader from '$lib/components/page-header.svelte';
	import UsersIcon from '@lucide/svelte/icons/users';
	import * as Empty from '$lib/components/ui/empty';
	import BlocksIcon from '@lucide/svelte/icons/blocks';
	import RoundedTabs from '$lib/components/ui/rounded-tabs/rounded-tabs.svelte';
</script>

<PageHeader class="shrink-0">
	{#snippet header()}
		<UsersIcon class="fill-yellow-400 text-6xl" />
		<h1 class="text-xl">Kelola Pengguna</h1>
	{/snippet}
</PageHeader>

<RoundedTabs class="flex min-h-0 grow justify-center">
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
