<script lang="ts">
	import PageHeader from '$lib/components/page-header.svelte';
	import { Button, buttonVariants } from '$lib/components/ui/button/index';
	import DataTable from '$lib/components/ui/data-table/data-table.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Empty from '$lib/components/ui/empty';
	import * as Form from '$lib/components/ui/form/index';
	import RoundedTabs from '$lib/components/ui/rounded-tabs/rounded-tabs.svelte';
	import * as Select from '$lib/components/ui/select';
	import FolderIcon from '@lucide/svelte/icons/folder';
	import UsersIcon from '@lucide/svelte/icons/users';
	import { untrack } from 'svelte';
	import { toast } from 'svelte-sonner';
	import { superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import type { PageProps } from './$types';
	import { columns, columnsLabel } from './columns';
	import { setEditRoleModalState } from './modal-state.svelte.ts';
	import { editUserRoleSchema } from './schema';
	import { Input } from '$lib/components/ui/input/index';

	let { data }: PageProps = $props();
	let editRoleModalState = setEditRoleModalState();

	const form = superForm(
		untrack(() => data.form),
		{
			validators: zod4Client(editUserRoleSchema),
			onError({ result }) {
				if (result.type === 'error' || result.type === 'failure') {
					toast.error(
						'Gagal memperbarui peran pengguna. Silakan periksa kembali data yang dimasukkan.'
					);
				}
			},
			onUpdate({ result }) {
				if (result.type === 'failure') {
					toast.error(
						'Gagal memperbarui peran pengguna. Silakan periksa kembali data yang dimasukkan.'
					);
				}
			},
			onUpdated() {
				toast.success('Peran pengguna berhasil diperbarui');
				editRoleModalState.close();
			}
		}
	);

	const { form: formData, enhance } = form;

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
	const roleTriggerContent = $derived(
		roleOptions.find((t) => t.value === $formData.role)?.label ?? 'Pilih peran sistem'
	);
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
	const academicRoleTriggerContent = $derived(
		academicRoleOptions.find((t) => t.value === $formData.academicRole)?.label ??
			'Pilih peran akademik'
	);

	$effect(() => {
		if (editRoleModalState.isOpen && editRoleModalState.rowData) {
			$formData.id = editRoleModalState.rowData.id;
			$formData.role = editRoleModalState.rowData.role;
			$formData.academicRole = editRoleModalState.rowData.academicRole;
		}
	});
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

<Dialog.Root bind:open={editRoleModalState.isOpen}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Edit Peran Pengguna</Dialog.Title>
		</Dialog.Header>
		<form
			method="POST"
			use:enhance
			action="?/editSingleUserRole"
			class="grid gap-4"
			id="edit-user-role-form"
		>
			<Form.Field {form} name="id">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<Form.Label>ID</Form.Label>
							<Input {...props} bind:value={$formData.id} disabled />
						</div>
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<p>Username</p>
							<p>{editRoleModalState.rowData.username}</p>
						</div>
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<p>Nama Pengguna</p>
							<p>{editRoleModalState.rowData.displayName}</p>
						</div>
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<p>Email</p>
							<p>{editRoleModalState.rowData.email || '-'}</p>
						</div>
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<p>ID Civitas</p>
							<p>{editRoleModalState.rowData.civitasId || '-'}</p>
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="role">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<Form.Label>Peran (Sistem)</Form.Label>
							<Select.Root type="single" bind:value={$formData.role} {...props}>
								<Select.Trigger class="w-45">
									{roleTriggerContent}
								</Select.Trigger>
								<Select.Content>
									<Select.Group>
										<Select.Label>Peran (Sistem)</Select.Label>
										{#each roleOptions as role (role.value)}
											<Select.Item value={role.value} label={role.label}>
												{role.label}
											</Select.Item>
										{/each}
									</Select.Group>
								</Select.Content>
							</Select.Root>
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="academicRole">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<Form.Label>Peran (Akademik)</Form.Label>
							<Select.Root type="single" bind:value={$formData.academicRole} {...props}>
								<Select.Trigger class="w-45">
									{academicRoleTriggerContent}
								</Select.Trigger>
								<Select.Content>
									<Select.Group>
										<Select.Label>Peran (Akademik)</Select.Label>
										{#each academicRoleOptions as academicRole (academicRole.value)}
											<Select.Item value={academicRole.value} label={academicRole.label}>
												{academicRole.label}
											</Select.Item>
										{/each}
									</Select.Group>
								</Select.Content>
							</Select.Root>
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</form>
		<Dialog.Footer>
			<Button
				type="submit"
				form="delete-course-form"
				class={buttonVariants({ variant: 'default' })}
			>
				Ubah Peran
			</Button>
			<Dialog.Close type="button" class={buttonVariants({ variant: 'outline' })}>
				Batalkan
			</Dialog.Close>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
