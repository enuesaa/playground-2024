<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry
	let shapeId: string | undefined = undefined

	function handleClick() {
		registry.svgOnMouseDown = ({ x, y }) => {
			shapeId = crypto.randomUUID()

			registry.shapes = [
				...registry.shapes,
				{
					id: shapeId,
					tag: 'path',
					d: `M${x},${y}`,
				},
			]
		}
		registry.svgOnClick = () => {}
		registry.svgOnMouseMove = ({ x, y }) => {
			if (shapeId === undefined) {
				return
			}
			for (let i = 0; i < registry.shapes.length; i++) {
				if (registry.shapes[i].tag !== 'path') {
					continue
				}
				if (registry.shapes[i].id === shapeId) {
					// @ts-ignore
					registry.shapes[i].d += ` L${x},${y}`
					break
				}
			}
		}
		registry.svgOnMouseLeave = () => {
			shapeId = undefined
		}
		registry.svgOnMouseUp = () => {
			shapeId = undefined
		}
		registry.rectOnMouseDown = () => {
			shapeId = undefined
		}
	}
</script>

<button on:click|preventDefault={handleClick}>Add path</button>

<style lang="postcss">
	button {
		@apply text-sm border-solid border-[1px] rounded mb-1 p-1;
	}
</style>
