<script lang="ts">
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import * as Breadcrumb from '$lib/components/ui/breadcrumb/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import Library from '@lucide/svelte/icons/library';
	import * as Tabs from '$lib/components/ui/tabs/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
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
			<Tabs.Root value="srl-profile" class="h-full w-full">
				<Tabs.List>
					<Tabs.Trigger value="srl-profile">Profil SRL</Tabs.Trigger>
					<Tabs.Trigger value="learning-objectives">Tujuan Pembelajaran</Tabs.Trigger>
					<Tabs.Trigger value="gradebook">Gradebook</Tabs.Trigger>
					<Tabs.Trigger value="activities">Aktivitas</Tabs.Trigger>
				</Tabs.List>
				<Tabs.Content value="srl-profile">
					<Card.Root class="h-full w-full overflow-auto">
						<Card.Header>
							<Card.Title>Profil SRL</Card.Title>
							<Card.Description>Update your SRL profile here.</Card.Description>
						</Card.Header>
						<Card.Content>
							<p>Change your SRL profile here.</p>
						</Card.Content>
					</Card.Root>
				</Tabs.Content>
				<Tabs.Content value="learning-objectives">
					<Card.Root class="h-full w-full overflow-auto">
						<Card.Header>
							<Card.Title>Tujuan Pembelajaran</Card.Title>
							<Card.Description>Update your learning objectives here.</Card.Description>
						</Card.Header>
						<Card.Content>
							<p>Change your learning objectives here.</p>
						</Card.Content>
					</Card.Root>
				</Tabs.Content>
				<Tabs.Content value="gradebook">
					<Card.Root class="h-full w-full overflow-auto">
						<Card.Header>
							<Card.Title>Gradebook</Card.Title>
							<Card.Description>View your grades here.</Card.Description>
						</Card.Header>
						<Card.Content>
							<p>View your grades here.</p>
						</Card.Content>
					</Card.Root>
				</Tabs.Content>
				<Tabs.Content value="activities">
					<Card.Root class="h-full w-full overflow-auto">
						<Card.Header>
							<Card.Title>Aktivitas</Card.Title>
							<Card.Description>View your activities here.</Card.Description>
						</Card.Header>
						<Card.Content>
							<p>View your activities here.</p>
						</Card.Content>
					</Card.Root>
				</Tabs.Content>
			</Tabs.Root>
		</div>
	</Sidebar.Inset>
</Sidebar.Provider>
