<script lang="ts">
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import * as Breadcrumb from '$lib/components/ui/breadcrumb/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import Library from '@lucide/svelte/icons/library';

	let { children } = $props();
</script>

<Sidebar.Provider>
	<AppSidebar />
	<Sidebar.Inset>
		<header class="flex h-16 shrink-0 items-center gap-2">
			<div class="flex items-center gap-2 px-4">
				<Sidebar.Trigger class="-ms-1" />
				<Separator orientation="vertical" class="me-2 data-[orientation=vertical]:h-4" />
				<Breadcrumb.Root>
					<Breadcrumb.List>
						<Breadcrumb.Item class="hidden md:block">
							<Breadcrumb.Link href="##">Kelas Saya</Breadcrumb.Link>
						</Breadcrumb.Item>
						<Breadcrumb.Separator class="hidden md:block" />
						<Breadcrumb.Item>
							<Breadcrumb.Page class="text-yellow-400">Sistem Interaksi</Breadcrumb.Page>
						</Breadcrumb.Item>
					</Breadcrumb.List>
				</Breadcrumb.Root>
			</div>
		</header>
		<div class="flex flex-1 flex-col gap-4 p-4 pt-0">
			<div
				class="flex h-24 items-center rounded-xl bg-linear-to-r from-yellow-600 to-yellow-300 p-4 shadow-md"
			>
				<Library class="text-white" />
				<div class="flex flex-col items-start justify-center gap-1 px-4">
					<h1 class="font-cal-sans text-xl text-white">[Reg] Sistem Interaksi (A dan B)</h1>
				</div>
			</div>
			{@render children()}
		</div>
	</Sidebar.Inset>
</Sidebar.Provider>
