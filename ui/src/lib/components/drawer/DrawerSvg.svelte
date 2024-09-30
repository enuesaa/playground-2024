<script lang="ts">
	import type { Registry, Position } from '$lib/registry'
	import Rect from './Rect.svelte'
	import Text from './Text.svelte'
	import Path from './Path.svelte'
	import Line from './Line.svelte'
	import LineArrowDef from './LineArrowDef.svelte'
	import ImagePng from './ImagePng.svelte'

	export let registry: Registry

	function updateSvgPosition(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		const { left, top } = e.currentTarget.getBoundingClientRect()
		registry.svg.x = left
		registry.svg.y = top
	}

	function rebaseXY(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }): Position {
		const x = e.clientX - registry.svg.x
		const y = e.clientY - registry.svg.y
		return { x, y }
	}

	function handleClick(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		updateSvgPosition(e)
		if (registry.svgOnClick === undefined) {
			return
		}
		registry.svgOnClick(rebaseXY(e))
	}

	function handleMouseMove(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseMove === undefined) {
			return
		}
		registry.svgOnMouseMove(rebaseXY(e))
	}

	function handleMouseUp(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseUp === undefined) {
			return
		}
		registry.svgOnMouseUp(rebaseXY(e))
	}

	function handleMouseLeave(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseLeave === undefined) {
			return
		}
		registry.svgOnMouseLeave(rebaseXY(e))
	}

	function handleMouseDown(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseDown === undefined) {
			return
		}
		registry.svgOnMouseDown(rebaseXY(e))
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
<svg
	viewBox="0 0 1000 600"
	width="1000"
	height="600"
	on:click={handleClick}
	on:mousemove={handleMouseMove}
	on:mouseup={handleMouseUp}
	on:mouseleave={handleMouseLeave}
	on:mousedown={handleMouseDown}
>
	{#each registry.shapes as shape}
		{#if shape.tag === 'rect'}
			<Rect {shape} {registry} />
		{:else if shape.tag === 'text'}
			<Text {shape} {registry} />
		{:else if shape.tag === 'line'}
			<Line {shape} {registry} />
		{:else if shape.tag === 'path'}
			<Path {shape} />
		{/if}
	{/each}

	<LineArrowDef />
	<ImagePng />
</svg>

<style lang="postcss">
	svg {
		width: 100%;
		height: 100%;
	}
</style>
