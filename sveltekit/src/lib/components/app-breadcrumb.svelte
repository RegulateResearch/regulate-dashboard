<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb';
	import { page } from '$app/state';

	let breadcrumbs = $derived.by(() => {
		if (page.data.breadcrumbs && page.data.breadcrumbs.length) {
			return (page.data.breadcrumbs as { name: string; href?: string }[]).map((bc, i) => ({
				name: bc.name,
				href: !bc.href || i === page.data.breadcrumbs!.length - 1 ? null : bc.href
			}));
		}
		return [];
	});
</script>

<Breadcrumb.Root>
	<Breadcrumb.List>
		{#if breadcrumbs.length > 0}
			{#each breadcrumbs as item, i (i)}
				{#if item.href}
					<Breadcrumb.Item>
						<Breadcrumb.Link href={item.href} class="hover:text-foreground">
							{item.name}
						</Breadcrumb.Link>
					</Breadcrumb.Item>
					<Breadcrumb.Separator />
				{:else}
					<Breadcrumb.Item>
						<Breadcrumb.Page>{item.name}</Breadcrumb.Page>
					</Breadcrumb.Item>
					{#if i < breadcrumbs.length - 1}
						<Breadcrumb.Separator />
					{/if}
				{/if}
			{/each}
		{:else}
			<Breadcrumb.Item>
				<Breadcrumb.Page>Regulate Dashboard</Breadcrumb.Page>
			</Breadcrumb.Item>
		{/if}
	</Breadcrumb.List>
</Breadcrumb.Root>
