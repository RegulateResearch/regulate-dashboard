<script lang="ts">
	import { Button } from "$lib/components/ui/button";

	let { data } = $props();
</script>

<p>Congrats, you reached this page.</p>
<p>User Info: {JSON.stringify(data?.userInfo)}</p>

<form method="POST">
  <Button type="submit">Logout</Button>
</form>