<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry;

	function handleClick() {
		registry = {
			svgOnClick: () => {
				console.log('rect')
			},
			svgOnMouseMove: () => {},
			svgOnMouseLeave: () => {},
			svgOnMouseUp: () => {},
		}
	}
</script>

<button on:click|preventDefault={handleClick}>rect</button>
