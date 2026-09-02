<script lang="ts">
	import PageHeader from '$lib/components/page-header.svelte';
	import { Button, buttonVariants } from '$lib/components/ui/button/index';
	import DataTable from '$lib/components/ui/data-table/data-table.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Empty from '$lib/components/ui/empty';
	import * as Form from '$lib/components/ui/form/index';
	import { Input } from '$lib/components/ui/input/index';
	import RoundedTabs from '$lib/components/ui/rounded-tabs/rounded-tabs.svelte';
	import * as Select from '$lib/components/ui/select';
	import type { UserWithId } from '$lib/schema';
	import { cn } from '$lib/utils';
	import FolderIcon from '@lucide/svelte/icons/folder';
	import SaveIcon from '@lucide/svelte/icons/save';
	import UserCogIcon from '@lucide/svelte/icons/user-cog';
	import UsersIcon from '@lucide/svelte/icons/users';
	import XIcon from '@lucide/svelte/icons/x';
	import type { RowModel } from '@tanstack/table-core';
	import { untrack } from 'svelte';
	import { toast } from 'svelte-sonner';
	import { superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import type { PageProps } from './$types';
	import { columns, columnsLabel } from './columns';
	import { setEditRoleModalState } from './modal-state.svelte.ts';
	import { editBulkUserRoleSchema, editSingleUserRoleSchema } from './schema';

	let { data }: PageProps = $props();
	let editRoleModalState = setEditRoleModalState();
	let editBulkModalOpen = $state(false);

	const editSingleRoleForm = superForm(
		untrack(() => data.editSingleRoleForm),
		{
			validators: zod4Client(editSingleUserRoleSchema),
			dataType: 'json',
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
			onUpdated({ form }) {
				if (form.valid) {
					toast.success('Peran pengguna berhasil diperbarui');
					editRoleModalState.close();
				}
			}
		}
	);
	const editBulkRoleForm = superForm(
		untrack(() => data.editBulkRoleForm),
		{
			validators: zod4Client(editBulkUserRoleSchema),
			dataType: 'json',
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
			onUpdated({ form }) {
				if (form.valid) {
					toast.success('Peran pengguna berhasil diperbarui');
					toggleBulkEditMode();
					toggleBulkModal();
				}
			}
		}
	);

	const { form: editSingleFormData, enhance: editSingleFormEnhance } = editSingleRoleForm;
	const { form: editBulkFormData, enhance: editBulkFormEnhance } = editBulkRoleForm;

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
	const singleRoleTriggerContent = $derived(
		roleOptions.find((t) => t.value === $editSingleFormData.role)?.label ?? 'Pilih peran sistem'
	);
	const bulkRoleTriggerContent = $derived(
		roleOptions.find((t) => t.value === $editBulkFormData.role)?.label ?? 'Pilih peran sistem'
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
	const singleAcademicRoleTriggerContent = $derived(
		academicRoleOptions.find((t) => t.value === $editSingleFormData.academicRole)?.label ??
			'Pilih peran akademik'
	);
	const bulkAcademicRoleTriggerContent = $derived(
		academicRoleOptions.find((t) => t.value === $editBulkFormData.academicRole)?.label ??
			'Pilih peran akademik'
	);

	let bulkEditMode = $state(false);
	let columnVisibility = $derived(bulkEditMode ? { select: true } : { select: false });
	let bulkSelectionState = $state({});
	let selectedBulk = $state<RowModel<UserWithId>>();

	const toggleBulkEditMode = () => {
		if (!bulkEditMode) bulkSelectionState = {};
		bulkEditMode = !bulkEditMode;
		$editBulkFormData.role = 'user';
		$editBulkFormData.academicRole = 'lecturer';
	};

	const toggleBulkModal = () => {
		editBulkModalOpen = !editBulkModalOpen;
	};

	$effect(() => {
		if (editRoleModalState.isOpen && editRoleModalState.rowData) {
			$editSingleFormData.id = editRoleModalState.rowData.id;
			$editSingleFormData.role = editRoleModalState.rowData.role;
			$editSingleFormData.academicRole = editRoleModalState.rowData.academicRole || 'student';
		}
	});

	$effect(() => {
		$editBulkFormData.id = selectedBulk ? selectedBulk.rows.map((row) => row.original.id) : [];
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
			bind:columnVisibility
			bind:rowSelection={bulkSelectionState}
			bind:selectedRow={selectedBulk}
			class={cn(['h-full w-full', bulkEditMode && 'bulk-edit-active'])}
		>
			{#snippet additionalMenu()}
				<div class="flex gap-2">
					{#if bulkEditMode}
						<Button type="button" size="sm" onclick={toggleBulkModal}>
							<SaveIcon />
							Simpan Perubahan Masal
						</Button>
					{/if}
					<Button
						type="button"
						size="sm"
						variant={bulkEditMode ? 'destructive' : 'default'}
						onclick={toggleBulkEditMode}
					>
						{#if bulkEditMode}
							<XIcon />
							Batalkan
						{:else}
							<UserCogIcon />
							Ubah Peran Massal
						{/if}
					</Button>
				</div>
			{/snippet}
		</DataTable>
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
			<Dialog.Title>Ubah Peran Pengguna</Dialog.Title>
		</Dialog.Header>
		<form
			method="POST"
			use:editSingleFormEnhance
			action="?/editSingleUserRole"
			class="grid gap-4"
			id="edit-user-role-form"
		>
			<Form.Field
				form={editSingleRoleForm}
				name="id"
				class="rounded-lg border border-yellow-200 bg-yellow-50 p-4"
			>
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2 text-yellow-500">
							<Form.Label>ID</Form.Label>
							<Input {...props} bind:value={$editSingleFormData.id} disabled />
						</div>
						<div class="grid grid-cols-[120px_1fr] gap-2 text-yellow-500">
							<p>Username</p>
							<p>{editRoleModalState.rowData.username}</p>
						</div>
						<div class="grid grid-cols-[120px_1fr] gap-2 text-yellow-500">
							<p>Nama Lengkap</p>
							<p>{editRoleModalState.rowData.displayName}</p>
						</div>
						<div class="grid grid-cols-[120px_1fr] gap-2 text-yellow-500">
							<p>Email</p>
							<p>{editRoleModalState.rowData.email || '-'}</p>
						</div>
						<div class="grid grid-cols-[120px_1fr] gap-2 text-yellow-500">
							<p>ID Civitas</p>
							<p>{editRoleModalState.rowData.civitasId || '-'}</p>
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field form={editSingleRoleForm} name="role">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<Form.Label>Peran (Sistem)</Form.Label>
							<Select.Root type="single" bind:value={$editSingleFormData.role} {...props} disabled>
								<Select.Trigger class="w-full">
									{singleRoleTriggerContent}
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
			<Form.Field form={editSingleRoleForm} name="academicRole">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<Form.Label>Peran (Akademik)</Form.Label>
							<Select.Root type="single" bind:value={$editSingleFormData.academicRole} {...props}>
								<Select.Trigger class="w-full">
									{singleAcademicRoleTriggerContent}
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
				form="edit-user-role-form"
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

<Dialog.Root bind:open={editBulkModalOpen}>
	<Dialog.Content>
		<Dialog.Header>
			{editBulkModalOpen}
			<Dialog.Title>Ubah Massal Peran Pengguna</Dialog.Title>
		</Dialog.Header>
		<form
			method="POST"
			use:editBulkFormEnhance
			action="?/editBulkUserRole"
			class="grid gap-4"
			id="edit-user-role-form"
		>
			<Form.Field
				form={editBulkRoleForm}
				name="id"
				class="rounded-lg border border-yellow-200 bg-yellow-50 p-4"
			>
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2 text-yellow-500">
							<Form.Label>ID</Form.Label>
							<Input {...props} bind:value={$editBulkFormData.id} disabled />
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field form={editBulkRoleForm} name="role">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<Form.Label>Peran (Sistem)</Form.Label>
							<Select.Root type="single" bind:value={$editBulkFormData.role} {...props} disabled>
								<Select.Trigger class="w-full">
									{bulkRoleTriggerContent}
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
			<Form.Field form={editBulkRoleForm} name="academicRole">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[120px_1fr] gap-2">
							<Form.Label>Peran (Akademik)</Form.Label>
							<Select.Root type="single" bind:value={$editBulkFormData.academicRole} {...props}>
								<Select.Trigger class="w-full">
									{bulkAcademicRoleTriggerContent}
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
				form="edit-user-role-form"
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

<style>
	:global(.bulk-edit-active [data-slot='table-container']) {
		border: 2px solid transparent !important;
		background-image:
			linear-gradient(var(--background, #fff), var(--background, #fff)),
			repeating-linear-gradient(90deg, #facc15 0, #facc15 8px, transparent 8px, transparent 16px),
			repeating-linear-gradient(180deg, #facc15 0, #facc15 8px, transparent 8px, transparent 16px),
			repeating-linear-gradient(90deg, #facc15 0, #facc15 8px, transparent 8px, transparent 16px),
			repeating-linear-gradient(180deg, #facc15 0, #facc15 8px, transparent 8px, transparent 16px);
		background-origin: padding-box, border-box, border-box, border-box, border-box;
		background-clip: padding-box, border-box, border-box, border-box, border-box;
		background-size:
			100% 100%,
			100% 2px,
			2px 100%,
			100% 2px,
			2px 100%;
		background-position:
			0 0,
			0 0,
			100% 0,
			0 100%,
			0 0;
		background-repeat: no-repeat, repeat-x, repeat-y, repeat-x, repeat-y;
		animation: marching-ants 0.8s linear infinite;
	}

	@keyframes marching-ants {
		0% {
			background-position:
				0 0,
				0 0,
				100% 0,
				0 100%,
				0 0;
		}
		100% {
			background-position:
				0 0,
				16px 0,
				100% 16px,
				-16px 100%,
				0 -16px;
		}
	}
</style>
