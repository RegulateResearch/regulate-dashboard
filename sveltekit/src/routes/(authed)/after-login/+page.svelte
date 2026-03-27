<script lang="ts">
	import { Button } from "$lib/components/ui/button";

	let data = $props();
</script>

<p>Congrats, you reached this page.</p>

<form method="POST">
  <Button type="submit">Logout</Button>
</form>