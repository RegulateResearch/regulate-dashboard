<script lang="ts">
	import { page } from '$app/state';
	import { Button } from '$lib/components/ui/button';
	import { env } from '$env/dynamic/public';
	import { UILogo } from '$lib/components/contents/logo/index';

	const handleClick = async () => {
		window.location.href = `${env.PUBLIC_SSO_UI_URL}/login?service=${encodeURIComponent(`${page.url.origin}/sso`)}`;
	};
</script>

<Button variant="outline" class="w-full" onclick={handleClick} type="button">
	<UILogo />
	Masuk dengan SSO UI
</Button>
