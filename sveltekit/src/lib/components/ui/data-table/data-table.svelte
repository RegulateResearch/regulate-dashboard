<script lang="ts" generics="TData, TValue">
	import {
		type ColumnDef,
		type ColumnFiltersState,
		getCoreRowModel,
		getFilteredRowModel,
		getPaginationRowModel,
		getSortedRowModel,
		type PaginationState,
		type RowSelectionState,
		type SortingState,
		type VisibilityState
	} from '@tanstack/table-core';
	import { createSvelteTable, FlexRender } from '$lib/components/ui/data-table/index.js';
	import * as Table from '$lib/components/ui/table/index.js';
	import { Button } from '$lib/components/ui//button';
	import { Input } from '$lib/components/ui/input';
	import ChevronLeftIcon from '@lucide/svelte/icons/chevron-left';
	import ChevronRightIcon from '@lucide/svelte/icons/chevron-right';
	import * as Select from '$lib/components/ui/select';
	import Label from '../label/label.svelte';
	import type { ClassValue } from 'clsx';
	import { cn } from '$lib/utils';

	type DataTableProps<TData, TValue> = {
		columns: ColumnDef<TData, TValue>[];
		data: TData[];
	};

	let {
		data,
		columns,
		columnToFilter,
		filterPlaceholder = 'Filter...',
		class: className
	}: DataTableProps<TData, TValue> & {
		columnToFilter: string;
		filterPlaceholder: string;
		class?: ClassValue;
	} = $props();
	let pagination = $state<PaginationState>({ pageIndex: 0, pageSize: 15 });
	let sorting = $state<SortingState>([]);
	let columnFilters = $state<ColumnFiltersState>([]);
	let columnVisibility = $state<VisibilityState>({});
	let rowSelection = $state<RowSelectionState>({});
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

	const table = createSvelteTable({
		get data() {
			return data;
		},
		columns,
		getCoreRowModel: getCoreRowModel(),
		getPaginationRowModel: getPaginationRowModel(),
		getSortedRowModel: getSortedRowModel(),
		getFilteredRowModel: getFilteredRowModel(),
		onPaginationChange: (updater) => {
			if (typeof updater === 'function') {
				pagination = updater(pagination);
			} else {
				pagination = updater;
			}
		},
		onSortingChange: (updater) => {
			if (typeof updater === 'function') {
				sorting = updater(sorting);
			} else {
				sorting = updater;
			}
		},
		onColumnFiltersChange: (updater) => {
			if (typeof updater === 'function') {
				columnFilters = updater(columnFilters);
			} else {
				columnFilters = updater;
			}
		},
		onColumnVisibilityChange: (updater) => {
			if (typeof updater === 'function') {
				columnVisibility = updater(columnVisibility);
			} else {
				columnVisibility = updater;
			}
		},
		onRowSelectionChange: (updater) => {
			if (typeof updater === 'function') {
				rowSelection = updater(rowSelection);
			} else {
				rowSelection = updater;
			}
		},
		state: {
			get pagination() {
				return pagination;
			},
			get sorting() {
				return sorting;
			},
			get columnFilters() {
				return columnFilters;
			},
			get columnVisibility() {
				return columnVisibility;
			},
			get rowSelection() {
				return rowSelection;
			}
		}
	});
</script>

<div class={cn('container flex h-full flex-col gap-1', className)}>
	<div class="flex shrink-0 items-center py-4">
		<Input
			placeholder={filterPlaceholder}
			value={(table.getColumn(columnToFilter)?.getFilterValue() as string) ?? ''}
			onchange={(e) => {
				table.getColumn(columnToFilter)?.setFilterValue(e.currentTarget.value);
			}}
			oninput={(e) => {
				table.getColumn(columnToFilter)?.setFilterValue(e.currentTarget.value);
			}}
			class="max-w-sm"
		/>
	</div>
	<div class="flex min-h-0 grow">
		<Table.Root wrapperClass="rounded-md border overflow-scroll scrollbar-thumb-yellow-400 max-h-max">
			<Table.Header>
				{#each table.getHeaderGroups() as headerGroup (headerGroup.id)}
					<Table.Row>
						{#each headerGroup.headers as header (header.id)}
							<Table.Head
								colspan={header.colSpan}
								class="sticky top-0 z-10 bg-background after:absolute after:bottom-0 after:left-0 after:h-px after:w-full after:bg-border"
							>
								{#if !header.isPlaceholder}
									<FlexRender
										content={header.column.columnDef.header}
										context={header.getContext()}
									/>
								{/if}
							</Table.Head>
						{/each}
					</Table.Row>
				{/each}
			</Table.Header>

			<Table.Body>
				{#each table.getRowModel().rows as row (row.id)}
					<Table.Row data-state={row.getIsSelected() && 'selected'}>
						{#each row.getVisibleCells() as cell (cell.id)}
							<Table.Cell>
								<FlexRender content={cell.column.columnDef.cell} context={cell.getContext()} />
							</Table.Cell>
						{/each}
					</Table.Row>
				{:else}
					<Table.Row>
						<Table.Cell colspan={columns.length} class="h-24 text-center">No results.</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>
	<div class="flex shrink-0 items-center justify-end space-x-2 py-4">
		<div class="flex gap-2">
			<Label>Baris per halaman:</Label>
			<Select.Root
				type="single"
				value={table.getState().pagination.pageSize.toString()}
				onValueChange={(value) => {
					table.setPageSize(Number(value));
				}}
			>
				<Select.Trigger>
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
		<Button
			variant="outline"
			size="icon"
			onclick={() => table.previousPage()}
			disabled={!table.getCanPreviousPage()}
		>
			<ChevronLeftIcon />
		</Button>
		<span class="flex items-center gap-1 text-sm">
			Halaman {table.getState().pagination.pageIndex + 1} dari {table.getPageCount()}
		</span>
		<Button
			variant="outline"
			size="icon"
			onclick={() => table.nextPage()}
			disabled={!table.getCanNextPage()}
		>
			<ChevronRightIcon />
		</Button>
	</div>
</div>
