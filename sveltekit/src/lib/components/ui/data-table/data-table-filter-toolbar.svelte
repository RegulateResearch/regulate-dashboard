<script lang="ts" generics="TData">
	import { Button } from '$lib/components/ui//button';
	import { Input } from '$lib/components/ui/input';
	import { cn } from '$lib/utils';
	import XIcon from '@lucide/svelte/icons/x';
	import type { Table } from '@tanstack/table-core';
	import type { ClassValue } from 'clsx';
	import type { Component } from 'svelte';
	import DataTableColumnVisibility from './data-table-column-visibility.svelte';
	import DataTableFacetedFilter from './data-table-faceted-filter.svelte';

	let {
		table,
		class: className,
		searchPlaceholder = 'Filter...',
		columnToSearch,
		categorialFilters,
		columnsLabel
	}: {
		table: Table<TData>;
		class?: ClassValue;
		searchPlaceholder: string;
		columnToSearch: string;
		categorialFilters?: {
			title: string;
			colName: string;
			options: {
				label?: string;
				value: string;
				icon?: Component;
			}[];
		}[];
		columnsLabel: {
			id: string;
			label: string;
		}[];
	} = $props();

	const isFiltered = $derived(table.getState().columnFilters.length > 0);
</script>

<div class={cn('flex items-center justify-between gap-2', className)}>
	<div class="flex min-h-0 grow items-center gap-2">
		<Input
			placeholder={searchPlaceholder}
			value={(table.getColumn(columnToSearch)?.getFilterValue() as string) ?? ''}
			onchange={(e) => {
				table.getColumn(columnToSearch)?.setFilterValue(e.currentTarget.value);
			}}
			oninput={(e) => {
				table.getColumn(columnToSearch)?.setFilterValue(e.currentTarget.value);
			}}
			class="h-8 max-w-sm rounded-lg"
		/>
		{#if categorialFilters}
			{#each categorialFilters as { title, colName, options }, i (i)}
				<DataTableFacetedFilter {title} column={table.getColumn(colName)!} {options} />
			{/each}
		{/if}
		{#if isFiltered}
			<Button
				variant="ghost"
				size="sm"
				class="text-red-500 hover:text-red-500"
				onclick={() => {
					table.resetColumnFilters();
				}}
			>
				Reset filter
				<XIcon></XIcon>
			</Button>
		{/if}
	</div>
	<DataTableColumnVisibility {table} {columnsLabel} />
</div>
