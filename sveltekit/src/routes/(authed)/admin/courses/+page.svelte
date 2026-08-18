<script lang="ts">
	import { enhance as defaultEnhance } from '$app/forms';
	import PageHeader from '$lib/components/page-header.svelte';
	import { Button, buttonVariants } from '$lib/components/ui/button';
	import DataTable from '$lib/components/ui/data-table/data-table.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Empty from '$lib/components/ui/empty';
	import * as Field from '$lib/components/ui/field';
	import * as Form from '$lib/components/ui/form/index';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import RoundedTabs from '$lib/components/ui/rounded-tabs/rounded-tabs.svelte';
	import * as Select from '$lib/components/ui/select';
	import { Switch } from '$lib/components/ui/switch';
	import FolderIcon from '@lucide/svelte/icons/folder';
	import PlusIcon from '@lucide/svelte/icons/plus';
	import SquareLibraryIcon from '@lucide/svelte/icons/square-library';
	import { toast } from "svelte-sonner";
	import { superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import type { PageProps } from './$types';
	import { columns, columnsLabel } from './columns';
	import { newCourseFormSchema } from './schema';
	import {
		setCreateCourseModalState,
		setDeleteCourseModalState
	} from './shared-modal-state.svelte.ts';

	let { data }: PageProps = $props();

	const form = superForm(data.form, {
		validators: zod4Client(newCourseFormSchema),
		onResult({result}) {
			if (result.type === 'success') {
				toast.success('Kelas berhasil dibuat');
				newCourseModalState.close();
			} else if (result.type === 'error' || result.type === 'failure') {
				toast.error('Gagal membuat kelas. Silakan periksa kembali data yang dimasukkan.');
			}
		}
	});

	const { form: formData, enhance } = form;
	let newCourseModalState = setCreateCourseModalState();
	let deleteCourseModalState = setDeleteCourseModalState();

	const yearOptions = (() => {
		const currentYear = new Date().getFullYear();
		return [
			...Array.from({ length: 3 }, (_, i) => ({
				value: (currentYear - i - 1).toString(),
				label: `${currentYear - i - 1}/${currentYear - i}`
			})),
			...Array.from({ length: 7 }, (_, i) => ({
				value: (currentYear + i).toString(),
				label: `${currentYear + i}/${currentYear + i + 1}`
			}))
		].sort((a, b) => parseInt(a.value) - parseInt(b.value));
	})();
	const yearTriggerContent = $derived(
		yearOptions.find((y) => y.value === $formData.year.toString())?.label ?? 'Pilih tahun akademik'
	);

	const termOptions = [
		{
			value: 'odd',
			label: 'Ganjil'
		},
		{
			value: 'even',
			label: 'Genap'
		},
		{
			value: 'short',
			label: 'Semester Pendek'
		}
	];
	const termTriggerContent = $derived(
		termOptions.find((t) => t.value === $formData.term)?.label ?? 'Pilih semester'
	);

	$effect(() => {
		if (newCourseModalState.isOpen && newCourseModalState.rowData) {
			$formData.name = newCourseModalState.rowData.name;
			$formData.year = newCourseModalState.rowData.year;
			$formData.term = newCourseModalState.rowData.term;
			$formData.url = newCourseModalState.rowData.url ?? undefined;
		}
	});
</script>

<PageHeader class="shrink-0">
	{#snippet header()}
		<SquareLibraryIcon class="fill-yellow-400 text-6xl" />
		<h1 class="text-xl">Kelola Kelas</h1>
	{/snippet}
	{#snippet actionBtn()}
		<Button type="button" variant="default" onclick={() => newCourseModalState.open()}>
			<PlusIcon />
			Buat Kelas
		</Button>
	{/snippet}
</PageHeader>

<RoundedTabs class="min-h-0 grow">
	{#if data.courses && data.courses.length > 0}
		<DataTable
			{columns}
			{columnsLabel}
			data={data.courses}
			columnToSearch="name"
			searchPlaceholder="Cari kelas"
			categorialFilters={[
				{
					title: 'Tahun Akademik',
					colName: 'year',
					options: yearOptions
				},
				{
					title: 'Semester',
					colName: 'term',
					options: termOptions
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
				<Empty.Title>Belum ada kelas</Empty.Title>
				<Empty.Description>
					Belum ada kelas yang dibuat. Buat kelas baru untuk mulai mengelola materi dan tugas.
				</Empty.Description>
			</Empty.Header>
			<Empty.Content>
				<div class="flex gap-2">
					<Button size="sm">
						<PlusIcon />
						Buat Kelas
					</Button>
				</div>
			</Empty.Content>
		</Empty.Root>
	{/if}
</RoundedTabs>

<Dialog.Root bind:open={newCourseModalState.isOpen}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Buat Kelas Baru</Dialog.Title>
		</Dialog.Header>
		<form
			method="POST"
			use:enhance
			action="?/createCourse"
			class="grid gap-4"
			id="create-course-form"
		>
			<Form.Field {form} name="name">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Nama Kelas</Form.Label>
						<Input {...props} bind:value={$formData.name} placeholder="Masukkan nama kelas" />
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<div class="flex gap-2">
				<Form.Field {form} name="year">
					<Form.Control>
						{#snippet children({ props })}
							<Form.Label>Tahun Akademik</Form.Label>
							<Select.Root type="single" bind:value={$formData.year} {...props}>
								<Select.Trigger class="w-45">
									{yearTriggerContent}
								</Select.Trigger>
								<Select.Content>
									<Select.Group>
										<Select.Label>Tahun Akademik</Select.Label>
										{#each yearOptions as year (year.value)}
											<Select.Item value={year.value} label={year.label}>
												{year.label}
											</Select.Item>
										{/each}
									</Select.Group>
								</Select.Content>
							</Select.Root>
						{/snippet}
					</Form.Control>
					<Form.FieldErrors />
				</Form.Field>
				<Form.Field {form} name="term">
					<Form.Control>
						{#snippet children({ props })}
							<Form.Label>Semester</Form.Label>
							<Select.Root type="single" bind:value={$formData.term} {...props}>
								<Select.Trigger class="w-45">
									{termTriggerContent}
								</Select.Trigger>
								<Select.Content>
									<Select.Group>
										<Select.Label>Semester</Select.Label>
										{#each termOptions as term (term.value)}
											<Select.Item value={term.value} label={term.label}>
												{term.label}
											</Select.Item>
										{/each}
									</Select.Group>
								</Select.Content>
							</Select.Root>
						{/snippet}
					</Form.Control>
					<Form.FieldErrors />
				</Form.Field>
			</div>
			<Form.Field {form} name="url">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>URL Kelas (cth. di LMS)</Form.Label>
						<Input
							{...props}
							bind:value={$formData.url}
							placeholder="Masukkan URL kelas (jika ada)"
						/>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</form>
		<Dialog.Footer>
			<Button
				type="submit"
				form="create-course-form"
				class={buttonVariants({ variant: 'default' })}
			>
				Buat Kelas
			</Button>
			<Dialog.Close type="button" class={buttonVariants({ variant: 'outline' })}>
				Batalkan
			</Dialog.Close>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>

<Dialog.Root bind:open={deleteCourseModalState.isOpen}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Hapus Kelas</Dialog.Title>
		</Dialog.Header>
		<form
			method="POST"
			action="?/deleteCourse"
			use:defaultEnhance={() => {
				return async ({ result, update }) => {
					if (result.type === 'success') {
						console.log('Delete course result:', result);
						deleteCourseModalState.close();
						toast.success('Kelas berhasil dihapus');
					} else {
						toast.error('Gagal menghapus kelas. Silakan coba lagi.');
					}
					await update();
				};
			}}
			class="grid gap-4"
			id="delete-course-form"
		>
			<input type="hidden" name="id" value={deleteCourseModalState.rowData?.id} />
			<Field.Field>
				<Field.Label for="input-id">Apakah Anda yakin ingin menghapus kelas ini?</Field.Label>
				<div class="flex items-center gap-2">
					<Switch bind:checked={deleteCourseModalState.isConfirmed} />
					<Label for="airplane-mode">Ya saya yakin</Label>
				</div>
				<Field.Description>Menghapus kelas ini akan menghapus semua data terkait.</Field.Description
				>
			</Field.Field>
		</form>
		<Dialog.Footer>
			<Button
				type="submit"
				form="delete-course-form"
				class={buttonVariants({ variant: 'default' })}
				disabled={!deleteCourseModalState.isConfirmed}
			>
				Hapus Kelas
			</Button>
			<Dialog.Close type="button" class={buttonVariants({ variant: 'outline' })}>
				Batalkan
			</Dialog.Close>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
