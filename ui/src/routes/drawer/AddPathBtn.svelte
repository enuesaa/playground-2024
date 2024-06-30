<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry;

	let current = ''

	function handleClick() {
		registry.svgOnClick = ({ x, y }) => {
			current = `M${x},${y}`
		}
		registry.svgOnMouseMove = ({ x, y }) => {
			if (current === '') {
				return
			}
			current += ` L${x},${y}`
		}
		registry.svgOnMouseLeave = () => {
			if (current === '') {
				return
			}
			registry.shapes = [
				...registry.shapes,
				{
					tag: 'path',
					d: current,
					stroke: '#000000',
					strokeWidth: 2,
				},
			]
			current = ''
		}
		registry.svgOnMouseUp = () => {
			if (current === '') {
				return
			}
			registry.shapes = [
				...registry.shapes,
				{
					tag: 'path',
					d: current,
					stroke: '#000000',
					strokeWidth: 2,
				},
			]
			current = ''
		}
	}
</script>

<button on:click|preventDefault={handleClick}>path</button>
