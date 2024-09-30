<script lang="ts">
	import { type Rect } from '$lib/shape'
	import { type Position, type Registry } from '$lib/registry'

	export let shape: Rect
	export let registry: Registry

	function rebaseXY(e: MouseEvent & { currentTarget: EventTarget & SVGRectElement }): Position {
		const x = e.clientX - registry.svg.x
		const y = e.clientY - registry.svg.y
		return { x, y }
	}

	function handleOnMouseDown(e: MouseEvent & { currentTarget: EventTarget & SVGRectElement }) {
		if (registry.rectOnMouseDown === undefined) {
			return
		}
		registry.rectOnMouseDown(shape.id, rebaseXY(e))
	}
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<rect
	x={shape.x}
	y={shape.y}
	width={shape.width}
	height={shape.height}
	fill={shape.fill}
	stroke="#ff6633"
	stroke-width="5"
	rx="10"
	ry="10"
	on:mousedown={handleOnMouseDown}
/>
