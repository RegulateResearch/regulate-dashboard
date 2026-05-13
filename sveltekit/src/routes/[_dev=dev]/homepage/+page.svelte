<script>
	import { SRL } from '$lib/components/contents/srl';
	import { Button } from '$lib/components/ui/button/index';
	import { RegulateLogo } from '$lib/components/contents/logo/index';
</script>

<nav class="mx-auto flex h-20 w-full flex-row items-center justify-center gap-1 px-4">
	<Button variant="ghost" size="icon-lg" href="/"><RegulateLogo></RegulateLogo></Button>
	<Button variant="link" size="sm" href="/login">Masuk</Button>
	<Button variant="link" size="sm" href="###">Tentang reguLAte!</Button>
	<Button variant="link" size="sm" href="###">Tentang Self-Regulated Learning!</Button>
</nav>
<section
	class="container mx-auto flex w-4/5 flex-row-reverse flex-wrap items-center justify-center gap-10 py-20"
>
	<SRL animated class="w-4/5 max-w-100"></SRL>
	<h1 class="text-center text-2xl xl:text-left xl:text-4xl">
		Kendalikan belajarmu dan jadilah <br /><em class="font-bold">Self-Regulated Learners!</em>
	</h1>
</section>
