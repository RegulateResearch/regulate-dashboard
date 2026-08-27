<script lang="ts">
	import NavMain from './nav-main.svelte';
	import NavUser from './nav-user.svelte';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import { RegulateLogo } from './contents/logo/index';
	import type { ComponentProps } from 'svelte';
	import { page } from '$app/state';
	import LayoutDashboardIcon from '@lucide/svelte/icons/layout-dashboard';
	import SquareLibraryIcon from '@lucide/svelte/icons/square-library';
	import UsersIcon from '@lucide/svelte/icons/users';

	let { ref = $bindable(null), ...restProps }: ComponentProps<typeof Sidebar.Root> = $props();
	let navItem = $derived.by(() => {
		if (page.data.userInfo?.role === 'admin')
			return [
				{
					title: 'Dashboard',
					url: '/admin/dashboard',
					icon: LayoutDashboardIcon
				},
				{
					title: 'Kelola Kelas',
					url: '/admin/courses',
					icon: SquareLibraryIcon
				},
				{
					title: 'Kelola Pengguna',
					url: '/admin/users',
					icon: UsersIcon
				}
			];

		if (page.data.userInfo?.academicRole)
			return [
				{
					title: 'Dashboard',
					url: '/lecturer/dashboard',
					icon: LayoutDashboardIcon
				},
				{
					title: 'Kelas Saya',
					url: '/lecturer/courses',
					icon: SquareLibraryIcon
				}
			];

		return [
			{
				title: 'Dashboard',
				url: '/student/dashboard',
				icon: LayoutDashboardIcon
			},
			{
				title: 'Kelas Saya',
				url: '/student/courses',
				icon: SquareLibraryIcon
			}
		];
	});
</script>

<Sidebar.Root bind:ref {...restProps}>
	<Sidebar.Header>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<Sidebar.MenuButton
					tooltipContent="Home"
					class="px-0 py-6 group-data-[collapsible=icon]:px-0! hover:bg-white"
				>
					{#snippet child({ props })}
						<!-- eslint-disable-next-line svelte/no-navigation-without-resolve -->
						<a href="/" {...props}>
							<RegulateLogo class="size-8! p-1" />
							<span class="text-2xl font-extralight">regulate</span>
						</a>
					{/snippet}
				</Sidebar.MenuButton>
			</Sidebar.MenuItem>
		</Sidebar.Menu>
	</Sidebar.Header>
	<Sidebar.Content>
		<NavMain items={navItem} />
	</Sidebar.Content>
	<Sidebar.Footer>
		<NavUser />
	</Sidebar.Footer>
</Sidebar.Root>
