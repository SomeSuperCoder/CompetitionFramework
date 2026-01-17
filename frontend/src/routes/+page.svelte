<script lang="ts">
    import { getCompetitons } from "$lib/api";
	import Button from "$lib/components/ui/button/button.svelte";
    import { onMount } from "svelte";

	let count = $state(0);

	onMount(async () => {
	  console.log(await getCompetitons())
    }) 
</script>

<Button onclick={() => count++}>{count}</Button>
