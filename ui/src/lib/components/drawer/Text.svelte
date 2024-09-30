<script lang="ts">
	import { type Text } from '$lib/shape'
	import { type Position, type Registry } from '$lib/registry'

	export let shape: Text
	export let registry: Registry

	function rebaseXY(e: MouseEvent & { currentTarget: EventTarget & SVGTextElement }): Position {
		const x = e.clientX - registry.svg.x
		const y = e.clientY - registry.svg.y
		return { x, y }
	}

	function handleOnMouseDown(e: MouseEvent & { currentTarget: EventTarget & SVGTextElement }) {
		if (registry.rectOnMouseDown === undefined) {
			return
		}
		registry.rectOnMouseDown(shape.id, rebaseXY(e))
	}
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<text x={shape.x} y={shape.y} on:mousedown={handleOnMouseDown}>
	{shape.value}
</text>
