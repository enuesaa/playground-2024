<script lang="ts">
	import { browser } from '$app/environment'
	import Drawer from './Drawer.svelte'
</script>

<Drawer />

{#if browser}
	<script src="https://cdn.tailwindcss.com"></script>
{/if}
