<script lang="ts">
	type Line = {
		x1: number;
		y1: number;
		x2: number;
		y2: number;
	}

	let lines: Line[] = []
	let dragging: number|undefined = undefined;

	function handleClick(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		const {left, top} = e.currentTarget.getBoundingClientRect()
		const x = e.clientX - left
		const y = e.clientY - top
		lines = [...lines, {
			x1: x,
			y1: y,
			x2: x + 100,
			y2: y,
		}]
	}

	function handleMouseMove(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (dragging === undefined) {
			return
		}
		const {left, top} = e.currentTarget.getBoundingClientRect();
		const x = e.clientX - left;
		const y = e.clientY - top;		

		lines = lines.map((l, i) => {
			if (i === dragging) {
				return { ...l, x2: x, y2: y }
			}
			return l
		})
	}

	function handleMouseUp() {
		dragging = undefined;
	}
</script>

<svg
	on:click={handleClick}
	on:mousemove={handleMouseMove}
	on:mouseup={handleMouseUp}
>
	{#each lines as line, i}
		<line
			x1={line.x1}
			y1={line.y1}
			x2={line.x2}
			y2={line.y2}
			stroke="black"
			stroke-width="2"
			on:mousedown={() => dragging = i} />
	{/each}
</svg>

<style lang="postcss">
	svg {
		width: 100%;
		height: 100vh;
	}
</style>

