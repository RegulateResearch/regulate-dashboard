<script>
	import { SRLContent } from '$lib/components/interactiveContents/srlContents';
	import { Button } from '$lib/components/ui/button/index';
	import { Logo } from '$lib/components/ui/logo/index';
</script>

<nav class="mx-auto flex h-20 w-full flex-row items-center justify-center gap-1 px-4">
	<Button variant="ghost" size="icon-lg" href="/"><Logo></Logo></Button>
	<Button variant="link" size="sm" href="/login">Masuk</Button>
	<Button variant="link" size="sm" href="###">Tentang reguLAte!</Button>
	<Button variant="link" size="sm" href="###">Tentang Self-Regulated Learning!</Button>
</nav>
<section class="container mx-auto flex w-full flex-col items-center justify-center gap-10 py-20">
	<SRLContent class="w-0.8 max-h-0.8 h-100"></SRLContent>
	<h1 class="text-center text-4xl">
		Kendalikan belajarmu dan jadilah <br /><em class="font-bold">Self-Regulated Learners!</em>
	</h1>
</section>
