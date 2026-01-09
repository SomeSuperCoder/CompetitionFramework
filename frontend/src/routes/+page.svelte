<script lang="ts">
    import { invokeJSONRPC } from "$lib/api";
	import Button from "$lib/components/ui/button/button.svelte";

	let count = $state(0);

	invokeJSONRPC("Competition.FindAllSortByName", {
      "priority": 123,
    })	
</script>

<Button onclick={() => count++}>{count}</Button>
