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
		switch (page.data.userInfo?.role) {
			case 'admin':
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
			case 'lecturer':
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
			case 'student':
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
			default:
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
		}
	});
</script>

<Sidebar.Root bind:ref variant="inset" {...restProps}>
	<Sidebar.Header>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<Sidebar.MenuButton size="lg">
					{#snippet child()}
						<div class="flex h-10 w-full items-center p-4">
							<RegulateLogo withTitle class="h-auto w-3/4" />
						</div>
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
