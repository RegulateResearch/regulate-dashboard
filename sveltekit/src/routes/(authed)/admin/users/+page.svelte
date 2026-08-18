<script lang="ts">
	import PageHeader from '$lib/components/page-header.svelte';
	import UsersIcon from '@lucide/svelte/icons/users';
	import * as Empty from '$lib/components/ui/empty';
	import RoundedTabs from '$lib/components/ui/rounded-tabs/rounded-tabs.svelte';
	import FolderIcon from '@lucide/svelte/icons/folder';
	import DataTable from '$lib/components/ui/data-table/data-table.svelte';
	import { columns, columnsLabel } from './columns';

	let { data }: PageProps = $props();

	const roleOptions = [
		{
			value: 'admin',
			label: 'Admin'
		},
		{
			value: 'user',
			label: 'User'
		}
	];
	const academicRoleOptions = [
		{
			value: 'lecturer',
			label: 'Dosen'
		},
		{
			value: 'student',
			label: 'Mahasiswa'
		}
	];
</script>

<PageHeader class="shrink-0">
	{#snippet header()}
		<UsersIcon class="fill-yellow-400 text-6xl" />
		<h1 class="text-xl">Kelola Pengguna</h1>
	{/snippet}
</PageHeader>

<RoundedTabs class="min-h-0 grow">
	{#if data.users && data.users.length > 0}
		<DataTable
			{columns}
			{columnsLabel}
			data={data.users}
			columnToSearch="displayName"
			searchPlaceholder="Cari pengguna"
			categorialFilters={[
				{
					title: 'Peran (Sistem)',
					colName: 'role',
					options: roleOptions
				},
				{
					title: 'Peran (Akademik)',
					colName: 'academicRole',
					options: academicRoleOptions
				}
			]}
			class="h-full w-full"
		></DataTable>
	{:else}
		<Empty.Root>
			<Empty.Header>
				<Empty.Media variant="icon">
					<FolderIcon />
				</Empty.Media>
				<Empty.Title>Belum ada pengguna</Empty.Title>
				<Empty.Description>
					Belum ada pengguna. Pengguna baru akan muncul di sini setelah mereka mendaftar atau
					ditambahkan oleh admin.
				</Empty.Description>
			</Empty.Header>
		</Empty.Root>
	{/if}
</RoundedTabs>
