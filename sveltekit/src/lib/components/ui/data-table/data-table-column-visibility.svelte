<script lang="ts" generics="TData">
	import { buttonVariants } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { cn } from '$lib/utils';
	import SettingsIcon from '@lucide/svelte/icons/settings-2';
	import type { Table } from '@tanstack/table-core';
	import type { ClassValue } from 'clsx';

	let {
		table,
		columnsLabel,
		class: className
	}: {
		table: Table<TData>;
		columnsLabel: {
			id: string;
			label: string;
		}[];
		class?: ClassValue;
	} = $props();
</script>

<div class={cn(className)}>
	<DropdownMenu.Root>
		<DropdownMenu.Trigger
			class={buttonVariants({
				variant: 'outline',
				size: 'sm',
				class: 'ms-auto hidden h-8 lg:flex'
			})}
		>
			<SettingsIcon />
			Kolom
		</DropdownMenu.Trigger>
		<DropdownMenu.Content align="end">
			{#each table.getAllColumns().filter((col) => col.getCanHide()) as column (column.id)}
				<DropdownMenu.CheckboxItem
					class="capitalize"
					bind:checked={() => column.getIsVisible(), (v) => column.toggleVisibility(!!v)}
				>
					{columnsLabel.find((col) => col.id === column.id)?.label ?? column.id}
				</DropdownMenu.CheckboxItem>
			{/each}
		</DropdownMenu.Content>
	</DropdownMenu.Root>
</div>
