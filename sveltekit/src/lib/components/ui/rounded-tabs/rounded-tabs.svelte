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

	let inactiveTabStyle: ClassValue =
		'relative h-8 rounded-t-xl border border-transparent bg-gray-100 px-6 py-2 text-xs text-gray-500 transition-colors hover:bg-gray-50';
	let activeTabStyle: ClassValue = `relative z-10 -mb-px h-8 rounded-t-xl border border-b-transparent bg-white px-6 py-2 text-xs text-yellow-500
    before:absolute before:-bottom-px before:-left-6 before:h-5 before:w-6 before:rounded-br-xl before:border-r before:border-b before:border-stone-200 before:shadow-[6px_6px_0_white] before:content-['']
    after:absolute after:-right-6 after:-bottom-px after:h-5 after:w-6 after:rounded-bl-xl after:border-b after:border-l after:border-stone-200 after:shadow-[-6px_6px_0_white] after:content-['']`;
</script>

<div class={cn('flex flex-col', className)}>
	{#if tabList}
		<ul class="relative flex items-end px-10 font-sans text-sm font-medium text-gray-600">
			{#each tabList as tabItem (tabItem.value)}
				<li class={tabValue === tabItem.value ? activeTabStyle : inactiveTabStyle}>
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

	<div class="min-h-0 grow rounded-xl border bg-white p-6">
		{@render children?.()}
	</div>
</div>
