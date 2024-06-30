<script lang="ts">
	import AddCircleBtn from './AddCircleBtn.svelte'
	import AddPathBtn from './AddPathBtn.svelte'

	type SvgOnClick = (x: number, y: number) => void;
	let svgOnClick: SvgOnClick|undefined;

	function handleClick(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		const { left, top } = e.currentTarget.getBoundingClientRect()
		const x = e.clientX - left
		const y = e.clientY - top

		if (svgOnClick === undefined) {
			return
		}
		svgOnClick(x, y)
	}

	function handleMouseMove() {}
	function handleMouseUp() {}
	function handleMouseLeave() {}
</script>

<AddPathBtn bind:svgOnClick={svgOnClick} />
<AddCircleBtn bind:svgOnClick={svgOnClick} />

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
