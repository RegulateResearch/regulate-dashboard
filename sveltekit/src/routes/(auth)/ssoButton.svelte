<script lang="ts">
	import { page } from '$app/state';
	import { Button } from '$lib/components/ui/button';

	const handleClick = async () => {
		window.location.href = `https://sso.ui.ac.id/cas2/login?service=${encodeURIComponent(`${page.url.origin}/sso`)}`;
	};
</script>

<Button variant="outline" class="w-full" onclick={handleClick}>
	<enhanced:img src="$lib/assets/ui-logo.png" alt="Logo UI" height="24" />
	Masuk dengan SSO UI
</Button>
