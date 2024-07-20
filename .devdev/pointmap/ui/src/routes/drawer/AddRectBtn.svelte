<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry;

	function handleClick() {
		registry.svgOnClick = ({ x, y }) => {
			registry.shapes = [
				...registry.shapes,
				{
					tag: 'rect',
					x,
					y,
					width: 20,
					height: 20,
					fill: '#000000',
				},
			]
		}
		registry.svgOnMouseMove = () => {}
		registry.svgOnMouseLeave = () => {}
		registry.svgOnMouseUp = () => {}
	}
</script>

<button on:click|preventDefault={handleClick}>rect</button>
