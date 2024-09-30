<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry

	function handleClick() {
		registry.svgOnClick = ({ x, y }) => {
			registry.shapes = [
				...registry.shapes,
				{
					id: crypto.randomUUID(),
					tag: 'line',
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
		registry.rectOnMouseDown = () => {}
	}
</script>

<button on:click|preventDefault={handleClick}>Add line</button>

<style lang="postcss">
	button {
		@apply text-sm border-solid border-[1px] rounded mb-1 p-1;
	}
</style>
