<script lang="ts">
	import { page } from '$app/state';
	import { Button } from '$lib/components/ui/button';
	import { env } from '$env/dynamic/public';

	const handleClick = async () => {
		window.location.href = `${env.PUBLIC_SSO_UI_URL}/login?service=${encodeURIComponent(`${page.url.origin}/sso`)}`;
	};
</script>

<Button variant="outline" class="w-full" onclick={handleClick} type="button">
	<enhanced:img src="$lib/assets/ui-logo.png" alt="Logo UI" height="24" />
	Masuk dengan SSO UI
</Button>
