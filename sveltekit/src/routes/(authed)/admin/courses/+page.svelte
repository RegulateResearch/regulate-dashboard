<script lang="ts">
	import DataTable from '$lib/components/ui/data-table/data-table.svelte';
	import SquareLibraryIcon from '@lucide/svelte/icons/square-library';
	import { columns } from './columns';
	import type { PageProps } from './$types';
	import * as Empty from '$lib/components/ui/empty';
	import FolderIcon from '@lucide/svelte/icons/folder';
	import PlusIcon from '@lucide/svelte/icons/plus';
	import { Button, buttonVariants } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import * as Select from '$lib/components/ui/select';
	import * as Form from '$lib/components/ui/form/index';
	import { superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import { formSchema } from './schema';
	import PageHeader from '$lib/components/page-header.svelte';
	import RoundedTabs from '$lib/components/ui/rounded-tabs/rounded-tabs.svelte';

	let { data }: PageProps = $props();

	const form = superForm(data.form, {
		validators: zod4Client(formSchema),
		onResult() {
			dialogOpen = false;
		}
	});

	const { form: formData, enhance } = form;
	let dialogOpen = $state(false);
	let courses = $derived(
		data.courses
			? data.courses.map((course) => {
					let term: 'Ganjil' | 'Genap' | 'Semester Pendek' = 'Ganjil';
					switch (course.term) {
						case 'odd':
							term = 'Ganjil';
							break;
						case 'even':
							term = 'Genap';
							break;
						case 'short':
							term = 'Semester Pendek';
							break;
					}
					return {
						...course,
						term
					};
				})
			: []
	);

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
</script>

<PageHeader class="shrink-0">
	{#snippet header()}
		<SquareLibraryIcon class="fill-yellow-400 text-6xl" />
		<h1 class="text-xl">Kelola Kelas</h1>
	{/snippet}
	{#snippet actionBtn()}
		<Dialog.Root bind:open={dialogOpen}>
			<Dialog.Trigger type="button" class={buttonVariants({ variant: 'default' })}>
				<PlusIcon />
				Buat Kelas
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>Buat Kelas Baru</Dialog.Title>
				</Dialog.Header>
				<form method="POST" use:enhance class="grid gap-4">
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
					<Form.Button>Buat Kelas</Form.Button>
				</form>
				<Dialog.Footer>
					<Dialog.Close type="button" class={buttonVariants({ variant: 'outline' })}>
						Batalkan
					</Dialog.Close>
				</Dialog.Footer>
			</Dialog.Content>
		</Dialog.Root>
	{/snippet}
</PageHeader>

<RoundedTabs class="min-h-0 grow">
	{#if data.courses && data.courses.length > 0}
		<DataTable {columns} data={courses} columnToFilter="name" filterPlaceholder="Cari kelas" />
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
