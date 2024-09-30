<script lang="ts">
	import { page } from '$app/stores'
	import type { Registry } from '$lib/registry'
	import Drawer from './Drawer.svelte'
	import { getDrawing } from '$lib/drawing'

	let id: string
	let registry: Registry

	$: {
		id = $page.params.id
		registry = getDrawing(id)
	}
</script>

<Drawer {registry} {id} />
