<script lang="ts">
	import AddCircleBtn from './AddCircleBtn.svelte'
	import AddPathBtn from './AddPathBtn.svelte'
	import AddRectBtn from './AddRectBtn.svelte'
	import type { Registry } from '$lib/registry'

	let registry: Registry = {
		svgOnClick: undefined,
		svgOnMouseMove: undefined,
		svgOnMouseUp: undefined,
		svgOnMouseLeave: undefined,
	}

	function calcXY(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		const { left, top } = e.currentTarget.getBoundingClientRect()
		return {
			x: e.clientX - left,
			y: e.clientY - top,
		}
	}

	function handleClick(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnClick === undefined) {
			return
		}
		const {x, y} = calcXY(e)
		registry.svgOnClick(x, y)
	}

	function handleMouseMove(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseMove === undefined) {
			return
		}
		const {x, y} = calcXY(e)
		registry.svgOnMouseMove(x, y)
	}

	function handleMouseUp(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseUp === undefined) {
			return
		}
		const {x, y} = calcXY(e)
		registry.svgOnMouseUp(x, y)
	}

	function handleMouseLeave(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (registry.svgOnMouseLeave === undefined) {
			return
		}
		const {x, y} = calcXY(e)
		registry.svgOnMouseLeave(x, y)
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

</svg>

<style lang="postcss">
	svg {
		width: 100%;
		height: 100vh;
	}
</style>
