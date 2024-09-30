<script lang="ts">
	import type { Position, Registry } from '$lib/registry'
	import type { Line } from '$lib/shape'

	export let shape: Line
	export let registry: Registry

	function rebaseXY(e: MouseEvent & { currentTarget: EventTarget & SVGLineElement }): Position {
		const x = e.clientX - registry.svg.x
		const y = e.clientY - registry.svg.y
		return { x, y }
	}

	function handleOnMouseDown(e: MouseEvent & { currentTarget: EventTarget & SVGLineElement }) {
		if (registry.rectOnMouseDown === undefined) {
			return
		}
		registry.rectOnMouseDown(shape.id, rebaseXY(e))
	}
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<line
	x1={shape.x}
	y1={shape.y}
	x2={shape.x + shape.width}
	y2={shape.y + shape.height}
	stroke={shape.fill}
	stroke-width="2"
	marker-end="url(#arrowhead)"
	on:mousedown={handleOnMouseDown}
/>
