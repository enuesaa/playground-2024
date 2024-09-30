<script lang="ts">
	import type { Registry } from '$lib/registry'
	import DrawerMenu from './DrawerMenu.svelte'
	import DrawerDangerZone from './DrawerDangerZone.svelte'
	import { createRegistry } from '$lib/drawing'
	import DrawerSvg from './DrawerSvg.svelte'
	import DrawerHistories from './DrawerHistories.svelte'

	export let id: string | undefined = undefined
	export let registry: Registry = createRegistry()
</script>

<section class="flex justify-between px-6">
	<div class="w-48">
		<DrawerHistories {registry} {id} />
	</div>
	<div class="w-[1000px]">
		<DrawerMenu bind:registry {id} />
		<DrawerSvg bind:registry />
		<DrawerDangerZone {id} />
	</div>
</section>
