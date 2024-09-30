<script lang="ts">
	import type { Registry } from '$lib/registry'

	export let registry: Registry

	let shapeId: string | undefined = undefined
	let initial = { x: 0, y: 0, width: 0, height: 0 }

	function handleClick() {
		shapeId = undefined

		registry.svgOnClick = () => {}
		registry.svgOnMouseMove = ({ x, y }) => {
			if (shapeId === undefined) {
				return
			}
			for (let i = 0; i < registry.shapes.length; i++) {
				if (registry.shapes[i].id === shapeId && registry.shapes[i].tag === 'rect') {
					// @ts-ignore
					registry.shapes[i].width = initial.width + x - initial.x
					// @ts-ignore
					registry.shapes[i].height = initial.height + y - initial.y
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
				if (shape.tag !== 'rect') {
					continue
				}
				if (shape.id === id) {
					// どんどん伸びていくから初期値を保持している
					initial = {
						x,
						y,
						width: shape.width,
						height: shape.height,
					}
					shapeId = id
					break
				}
			}
		}
	}
</script>

<button on:click|preventDefault={handleClick}>Resize</button>

<style lang="postcss">
	button {
		@apply text-sm border-solid border-[1px] rounded mb-1 p-1;
	}
</style>
