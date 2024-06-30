<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry;

	function handleClick() {
		registry = {
			svgOnClick: () => {
				console.log('circle')
			},
			svgOnMouseMove: () => {},
			svgOnMouseLeave: () => {},
			svgOnMouseUp: () => {},
		}
	}
</script>

<button on:click|preventDefault={handleClick}>circle</button>
