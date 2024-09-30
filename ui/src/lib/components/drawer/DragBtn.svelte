<script lang="ts">
	import type { Position, Registry } from '$lib/registry'

	export let registry: Registry

	let shapeId: string | undefined = undefined
	let start: Position = { x: 0, y: 0 }

	function handleClick() {
		shapeId = undefined

		registry.svgOnClick = () => {}
		registry.svgOnMouseMove = ({ x, y }) => {
			if (shapeId === undefined) {
				return
			}
			for (let i = 0; i < registry.shapes.length; i++) {
				if (registry.shapes[i].tag === 'path') {
					continue
				}
				if (registry.shapes[i].id === shapeId) {
					// @ts-ignore
					registry.shapes[i].x = x - start.x
					// @ts-ignore
					registry.shapes[i].y = y - start.y
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
		registry.rectOnMouseDown = (id, { x, y }) => {
			for (let shape of registry.shapes) {
				if (shape.tag === 'path') {
					continue
				}
				if (shape.id === id) {
					start.x = x - shape.x
					start.y = y - shape.y
					shapeId = id
					break
				}
			}
		}
	}
</script>

<button on:click|preventDefault={handleClick}>Drag</button>

<style lang="postcss">
	button {
		@apply text-sm border-solid border-[1px] rounded mb-1 p-1;
	}
</style>
