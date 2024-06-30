<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry;

	function handleClick() {
		registry.svgOnClick = () => {
			console.log('')
		}
		registry.svgOnMouseMove = () => {}
		registry.svgOnMouseLeave = () => {}
		registry.svgOnMouseUp = () => {}
	}
</script>

<button on:click|preventDefault={handleClick}>path</button>
