<script lang="ts" generics="TData">
	import type { Table } from '@tanstack/table-core';
	import type { ClassValue } from 'clsx';
	import { Label } from '../label';
	import * as Select from '$lib/components/ui/select';
	import { Button } from '$lib/components/ui//button';
	import { cn } from '$lib/utils';
	import ChevronLeftIcon from '@lucide/svelte/icons/chevron-left';
	import ChevronRightIcon from '@lucide/svelte/icons/chevron-right';
	import ChevronsLeftIcon from '@lucide/svelte/icons/chevrons-left';
	import ChevronsRightIcon from '@lucide/svelte/icons/chevrons-right';

	let {
		table,
		class: className
	}: {
		table: Table<TData>;
		class?: ClassValue;
	} = $props();

	const paginationSelectOptions = [
		{
			value: '15',
			label: '15'
		},
		{
			value: '25',
			label: '25'
		},
		{
			value: '50',
			label: '50'
		}
	];
	const paginationSelectTriggerContent = $derived(
		paginationSelectOptions.find((o) => o.value === table.getState().pagination.pageSize.toString())
			?.label ?? 'Pilih jumlah baris'
	);
</script>

<div class={cn('flex items-center justify-end gap-8', className)}>
	<div class="flex gap-2">
		<Label class="text-sm font-normal">Baris per halaman:</Label>
		<Select.Root
			type="single"
			value={table.getState().pagination.pageSize.toString()}
			onValueChange={(value) => {
				table.setPageSize(Number(value));
			}}
		>
			<Select.Trigger size="sm">
				{paginationSelectTriggerContent}
			</Select.Trigger>
			<Select.Content>
				<Select.Group>
					{#each paginationSelectOptions as option (option.value)}
						<Select.Item value={option.value} label={option.label}>
							{option.label}
						</Select.Item>
					{/each}
				</Select.Group>
			</Select.Content>
		</Select.Root>
	</div>
	<span class="flex items-center text-sm">
		Halaman {table.getState().pagination.pageIndex + 1} dari {table.getPageCount()}
	</span>
	<div class="flex gap-2">
		<Button
			variant="outline"
			size="icon-sm"
			onclick={() => table.firstPage()}
			disabled={!table.getCanPreviousPage()}
		>
			<ChevronsLeftIcon />
		</Button>
		<Button
			variant="outline"
			size="icon-sm"
			onclick={() => table.previousPage()}
			disabled={!table.getCanPreviousPage()}
		>
			<ChevronLeftIcon />
		</Button>
		<Button
			variant="outline"
			size="icon-sm"
			onclick={() => table.nextPage()}
			disabled={!table.getCanNextPage()}
		>
			<ChevronRightIcon />
		</Button>
		<Button
			variant="outline"
			size="icon-sm"
			onclick={() => table.lastPage()}
			disabled={!table.getCanNextPage()}
		>
			<ChevronsRightIcon />
		</Button>
	</div>
</div>
