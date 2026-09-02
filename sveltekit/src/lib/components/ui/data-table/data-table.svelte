<script lang="ts" generics="TData, TValue">
	import { createSvelteTable, FlexRender } from '$lib/components/ui/data-table/index.js';
	import * as Empty from '$lib/components/ui/empty';
	import * as Table from '$lib/components/ui/table/index.js';
	import { cn } from '$lib/utils';
	import SearchXIcon from '@lucide/svelte/icons/search-x';
	import {
		getCoreRowModel,
		getFacetedRowModel,
		getFacetedUniqueValues,
		getFilteredRowModel,
		getPaginationRowModel,
		getSortedRowModel,
		type ColumnDef,
		type ColumnFiltersState,
		type PaginationState,
		type RowModel,
		type RowSelectionState,
		type SortingState,
		type VisibilityState
	} from '@tanstack/table-core';
	import type { ClassValue } from 'clsx';
	import { type Component, type Snippet } from 'svelte';
	import DataTableFilterToolbar from './data-table-filter-toolbar.svelte';
	import DataTablePagination from './data-table-pagination.svelte';

	type DataTableProps<TData, TValue> = {
		columns: ColumnDef<TData, TValue>[];
		data: TData[];
	};

	let {
		data,
		columns,
		columnToSearch,
		columnsLabel,
		searchPlaceholder = 'Filter...',
		categorialFilters,
		columnVisibility = $bindable({}),
		rowSelection = $bindable({}),
		selectedRow = $bindable(),
		additionalMenu,
		class: className
	}: DataTableProps<TData, TValue> & {
		columnToSearch: string;
		columnsLabel: {
			id: string;
			label: string;
		}[];
		searchPlaceholder: string;
		categorialFilters?: {
			title: string;
			colName: string;
			options: {
				label?: string;
				value: string;
				icon?: Component;
			}[];
		}[];
		columnVisibility?: VisibilityState;
		rowSelection?: RowSelectionState;
		selectedRow?: RowModel<TData> | undefined;
		additionalMenu?: Snippet;
		class?: ClassValue;
	} = $props();

	let pagination = $state<PaginationState>({ pageIndex: 0, pageSize: 15 });
	let sorting = $state<SortingState>([]);
	let columnFilters = $state<ColumnFiltersState>([]);

	const table = createSvelteTable({
		get data() {
			return data;
		},
		get columns() {
			return columns;
		},
		getCoreRowModel: getCoreRowModel(),
		getPaginationRowModel: getPaginationRowModel(),
		getSortedRowModel: getSortedRowModel(),
		getFilteredRowModel: getFilteredRowModel(),
		getFacetedRowModel: getFacetedRowModel(),
		getFacetedUniqueValues: getFacetedUniqueValues(),
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

	$effect(() => {
		selectedRow = table.getSelectedRowModel();
	});
</script>

<div class={cn('flex flex-col gap-4', className)}>
	<div class="flex w-full flex-row gap-1">
		<DataTableFilterToolbar
			{table}
			{columnsLabel}
			{categorialFilters}
			{columnToSearch}
			{searchPlaceholder}
			class="grow"
		/>
		<div class="shrink-0">
			{#if additionalMenu}
				{@render additionalMenu()}
			{/if}
		</div>
	</div>
	<div class="flex min-h-0 grow">
		<Table.Root
			wrapperClass="rounded-lg border overflow-scroll scrollbar-thumb-yellow-400 max-h-max"
			class="table-fixed"
		>
			<Table.Header>
				{#each table.getHeaderGroups() as headerGroup (headerGroup.id)}
					<Table.Row>
						{#each headerGroup.headers as header (header.id)}
							<Table.Head
								style="width: {header.getSize()}px;"
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
							<Table.Cell style="width: {cell.column.getSize()}px;">
								<FlexRender content={cell.column.columnDef.cell} context={cell.getContext()} />
							</Table.Cell>
						{/each}
					</Table.Row>
				{:else}
					<Table.Row>
						<Table.Cell colspan={columns.length} class="h-24 text-center">
							<Empty.Root>
								<Empty.Header>
									<Empty.Media variant="icon">
										<SearchXIcon />
									</Empty.Media>
									<Empty.Title>Hasil pencarian kosong.</Empty.Title>
									<Empty.Description>
										Mohon sesuaikan kembali pencarian dan filter.
									</Empty.Description>
								</Empty.Header>
							</Empty.Root>
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>
	<DataTablePagination {table} class="shrink-0" />
</div>
