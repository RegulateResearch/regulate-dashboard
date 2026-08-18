<script lang="ts">
	import AppBreadcrumb from '$lib/components/app-breadcrumb.svelte';
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import { Toaster } from "$lib/components/ui/sonner/index.js";

	let { children } = $props();
</script>

<Toaster position="top-center" />

<Sidebar.Provider class="h-svh">
	<AppSidebar />
	<Sidebar.Inset class="max-h-full">
		<header class="flex h-16 shrink-0 items-center gap-2">
			<div class="flex items-center gap-2 px-4">
				<Sidebar.Trigger class="-ms-1" />
				<Separator orientation="vertical" class="me-2 data-[orientation=vertical]:h-4" />
				<AppBreadcrumb />
			</div>
		</header>
		<div class="flex min-h-0 grow flex-col gap-4 overflow-hidden p-4 pt-0">
			{@render children()}
		</div>
	</Sidebar.Inset>
</Sidebar.Provider>
