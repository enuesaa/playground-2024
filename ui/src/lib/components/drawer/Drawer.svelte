<script lang="ts">
	import type { Registry } from '$lib/registry'
	import DrawerMenu from './DrawerMenu.svelte'
	import { createRegistry } from '$lib/drawing'
	import DrawerSvg from './DrawerSvg.svelte'

	export let registry: Registry = createRegistry()
</script>


<DrawerMenu bind:registry />
<DrawerSvg bind:registry />
