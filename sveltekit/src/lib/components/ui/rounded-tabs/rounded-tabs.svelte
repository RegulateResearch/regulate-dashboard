<script lang="ts">
	import { cn } from '$lib/utils';
	import type { HTMLAttributes } from 'svelte/elements';
	import type { ClassValue } from 'tailwind-variants';

	type TabListItem = {
		name: string;
		value: string;
	};
	type RoundedTabsProps = HTMLAttributes<HTMLDivElement> & {
		tabList?: TabListItem[];
		tabValue?: string;
	};

	let {
		class: className,
		tabList,
		tabValue = $bindable(''),
		children
	}: RoundedTabsProps = $props();

	let baseTabStyle: ClassValue = 'relative -mb-px h-8 px-6 py-2 text-xs transition-colors';
	let inactiveTabStyle: ClassValue = `${baseTabStyle} text-gray-500 hover:text-yellow-500`;
	let firstActiveTabStyle: ClassValue = `${baseTabStyle} z-10 -mb-px rounded-t-xl border border-b-transparent bg-white text-yellow-500
		after:absolute after:-right-6 after:-bottom-px after:h-5 after:w-6 after:rounded-bl-xl after:border-b after:border-l after:border-stone-200 after:shadow-[-6px_6px_0_white] after:content-['']`;
	let activeTabStyle: ClassValue = `${firstActiveTabStyle}
    before:absolute before:-bottom-px before:-left-6 before:h-5 before:w-6 before:rounded-br-xl before:border-r before:border-b before:border-stone-200 before:shadow-[6px_6px_0_white] before:content-['']`;
</script>

<div class={cn('flex flex-col', className)}>
	{#if tabList}
		<ul class="relative flex items-end font-sans text-sm font-medium text-gray-600">
			{#each tabList as tabItem (tabItem.value)}
				<li
					class={tabValue === tabItem.value
						? tabItem.value === tabList[0]?.value
							? firstActiveTabStyle
							: activeTabStyle
						: inactiveTabStyle}
				>
					<button
						type="button"
						class="h-full w-full cursor-pointer"
						onclick={() => {
							tabValue = tabItem.value;
						}}>{tabItem.name}</button
					>
				</li>
			{/each}
		</ul>
	{/if}

	<div
		class={cn([
			'min-h-0 grow border bg-white p-4',
			tabList && tabValue === tabList[0]?.value ? 'rounded-tr-xl rounded-b-xl' : 'rounded-xl'
		])}
	>
		{@render children?.()}
	</div>
</div>
