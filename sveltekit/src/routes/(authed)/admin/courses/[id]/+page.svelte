<script lang="ts">
	import PageHeader from '$lib/components/page-header.svelte';
	import * as Empty from '$lib/components/ui/empty';
	import * as Form from '$lib/components/ui/form/index';
	import { Input } from '$lib/components/ui/input';
	import RoundedTabs from '$lib/components/ui/rounded-tabs/rounded-tabs.svelte';
	import * as Select from '$lib/components/ui/select';
	import BlocksIcon from '@lucide/svelte/icons/blocks';
	import SquareLibraryIcon from '@lucide/svelte/icons/square-library';
	import { untrack } from 'svelte';
	import { toast } from 'svelte-sonner';
	import { superForm } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import type { PageProps } from './$types';
	import { editCourseFormSchema } from './schema';
	import Button from '$lib/components/ui/button/button.svelte';
	import SquarePenIcon from '@lucide/svelte/icons/square-pen';
	import XIcon from '@lucide/svelte/icons/x';

	const tabList = [
		{ name: 'Pengaturan', value: 'settings' },
		{ name: 'Tujuan Pembelajaran', value: 'learningObjectives' },
		{ name: 'Peserta', value: 'members' },
		{ name: 'Materi & Aktivitas', value: 'courseItems' }
	];

	let tabValue = $state('settings');
	let enableEditCourseForm = $state(false);
	let { data }: PageProps = $props();

	const form = superForm(
		untrack(() => data.form),
		{
			validators: zod4Client(editCourseFormSchema),
			resetForm: false,
			onError({ result }) {
				if (result.type === 'error' || result.type === 'failure') {
					toast.error('Gagal memperbarui kelas. Silakan periksa kembali data yang dimasukkan.');
				}
			},
			onUpdate({ result }) {
				if (result.type === 'failure') {
					toast.error('Gagal memperbarui kelas. Silakan periksa kembali data yang dimasukkan.');
				}
			},
			onUpdated() {
				toast.success('Kelas berhasil diperbarui');
				enableEditCourseForm = false;
			}
		}
	);
	const { form: formData, enhance, reset } = form;

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
		<div class="flex flex-col gap-1">
			<div class="flex flex-row items-center gap-2">
				<SquareLibraryIcon class="fill-yellow-400 text-6xl" />
				<h1 class="text-xl">Kelola Detail Kelas</h1>
				{#if data.course}
					<p class="text-sm text-stone-300">{data.course.name} (ID: {data.course.id})</p>
				{/if}
			</div>
		</div>
	{/snippet}
	{#snippet actionBtn()}
		{#if data.course && data.course.url}
			<Button href={data.course.url} target="_blank" variant="outline">Buka Kelas di LMS</Button>
		{/if}
	{/snippet}
</PageHeader>

<RoundedTabs class="flex min-h-0 grow justify-center" {tabList} bind:tabValue>
	{#if tabValue === 'settings'}
		<div class="mb-8 flex w-full flex-row justify-between gap-2">
			<h1 class="my-2 font-semibold">Pengaturan Kelas</h1>
			<Button
				type="button"
				variant="ghost"
				class={enableEditCourseForm ? 'text-red-500' : 'text-yellow-400'}
				onclick={() => {
					enableEditCourseForm = !enableEditCourseForm;
					if (!enableEditCourseForm) {
						reset();
					}
				}}
			>
				{#if enableEditCourseForm}
					<XIcon />
				{:else}
					<SquarePenIcon />
				{/if}
				{enableEditCourseForm ? 'Batal' : 'Ubah Pengaturan'}
			</Button>
		</div>
		<form
			method="POST"
			use:enhance
			action="?/updateCourse"
			class="grid w-full gap-2"
			id="update-course-form"
		>
			<Form.Field {form} name="name">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[240px_1fr] gap-2">
							<Form.Label>Nama Kelas</Form.Label>
							<Input
								{...props}
								bind:value={$formData.name}
								placeholder="Masukkan nama kelas"
								disabled={!enableEditCourseForm}
							/>
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="year">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[240px_1fr] gap-2">
							<Form.Label>Tahun Akademik</Form.Label>
							<Select.Root
								type="single"
								bind:value={$formData.year}
								{...props}
								disabled={!enableEditCourseForm}
							>
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
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="term">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[240px_1fr] gap-2">
							<Form.Label>Semester</Form.Label>
							<Select.Root
								type="single"
								bind:value={$formData.term}
								{...props}
								disabled={!enableEditCourseForm}
							>
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
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="url">
				<Form.Control>
					{#snippet children({ props })}
						<div class="grid grid-cols-[240px_1fr] gap-2">
							<Form.Label>URL Kelas (cth. di LMS)</Form.Label>
							<Input
								{...props}
								bind:value={$formData.url}
								placeholder="Masukkan URL kelas (jika ada)"
								disabled={!enableEditCourseForm}
							/>
						</div>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Button type="submit" class="justify-self-end" disabled={!enableEditCourseForm}>
				Simpan Perubahan
			</Form.Button>
		</form>
	{:else if tabValue === 'learningObjectives'}
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
	{:else if tabValue === 'members'}
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
	{:else if tabValue === 'courseItems'}
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
	{/if}
</RoundedTabs>
