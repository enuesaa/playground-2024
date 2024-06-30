<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry;

	function handleClick() {
		registry.svgOnClick = ({ x, y }) => {
			registry.shapes = [
				...registry.shapes,
				{
					tag: 'circle',
					cx: x,
					cy: y,
					r: 20,
					fill: '#000000',
				},
			]
		}
		registry.svgOnMouseMove = () => {}
		registry.svgOnMouseLeave = () => {}
		registry.svgOnMouseUp = () => {}
	}
</script>

<button on:click|preventDefault={handleClick}>circle</button>
