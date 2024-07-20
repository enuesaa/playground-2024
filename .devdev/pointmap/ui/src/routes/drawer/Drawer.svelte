<script lang="ts">
	import AddCircleBtn from './AddCircleBtn.svelte'
	import AddPathBtn from './AddPathBtn.svelte'
	import AddRectBtn from './AddRectBtn.svelte'
	import type { Registry, Position } from '$lib/registry'

	let registry: Registry = {
		shapes: [],
		svgOnClick: undefined,
		svgOnMouseMove: undefined,
		svgOnMouseUp: undefined,
		svgOnMouseLeave: undefined,
	}

	function calcXY(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }): Position {
		const { left, top } = e.currentTarget.getBoundingClientRect()
		const x = e.clientX - left;
		const y = e.clientY - top;
		return { x, y }
	}

	function handleClick(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnClick === undefined) {
			return
		}
		registry.svgOnClick(calcXY(e))
	}

	function handleMouseMove(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseMove === undefined) {
			return
		}
		registry.svgOnMouseMove(calcXY(e))
	}

	function handleMouseUp(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseUp === undefined) {
			return
		}
		registry.svgOnMouseUp(calcXY(e))
	}

	function handleMouseLeave(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseLeave === undefined) {
			return
		}
		registry.svgOnMouseLeave(calcXY(e))
	}
</script>

<AddPathBtn bind:registry={registry} />
<AddCircleBtn bind:registry={registry} />
<AddRectBtn bind:registry={registry} />

<svg
	on:click={handleClick}
	on:mousemove={handleMouseMove}
	on:mouseup={handleMouseUp}
	on:mouseleave={handleMouseLeave}
>
	{#each registry.shapes as shape}
		{#if shape.tag === 'rect'}
			<rect x={shape.x} y={shape.y} width={shape.width} height={shape.height} fill={shape.fill} />
		{:else if shape.tag === 'circle'}
			<circle cx={shape.cx} cy={shape.cy} r={shape.r} fill={shape.fill} />
		{:else if shape.tag === 'path'}
			<path d={shape.d} stroke={shape.stroke} fill="none" stroke-width={shape.strokeWidth} />
		{/if}
	{/each}
</svg>

<style lang="postcss">
	svg {
		width: 100%;
		height: 100vh;
	}
</style>
