<script lang="ts">
	import type { Path, Circle, Rect, Shape  } from '$lib/shape'

	let shapes: Shape[] = []
	let current: Path = {
		tag: 'path',
		d: '',
		stroke: '#000000',
	}

	let selectedShape: 'path' | 'circle' | 'rect' | undefined = undefined
	let rectStyle = {
		fill: '#000000',
	}

	function handleClick(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (selectedShape === undefined) {
			return
		}
		const { left, top } = e.currentTarget.getBoundingClientRect()
		const x = e.clientX - left
		const y = e.clientY - top

		if (selectedShape === 'path') {
			current.d = `M${x},${y}`
			// continue drawing
		} else if (selectedShape === 'circle') {
			shapes = [
				...shapes,
				{
					tag: 'circle',
					cx: x,
					cy: y,
					r: 20,
					fill: '#000000',
				},
			]
			selectedShape = undefined
		} else if (selectedShape === 'rect') {
			shapes = [
				...shapes,
				{
					tag: 'rect',
					x,
					y,
					width: 20,
					height: 20,
					fill: rectStyle.fill,
				},
			]
			selectedShape = undefined
		}
	}

	function draw(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
		if (current.d === '') {
			return
		}

		const { left, top } = e.currentTarget.getBoundingClientRect()
		const x = e.clientX - left
		const y = e.clientY - top
		current.d += ` L${x},${y}`
	}

	function stop() {
		if (current.d === '') {
			return
		}
		shapes = [
			...shapes,
			{
				tag: 'path',
				d: current.d,
				stroke: current.stroke,
			},
		]
		current.d = ''
		selectedShape = undefined
	}
</script>

<button on:click={() => (selectedShape = 'path')}>path</button>
<button on:click={() => (selectedShape = 'circle')}>circle</button>
<button on:click={() => (selectedShape = 'rect')}>rect</button>

<svg on:click={handleClick} on:mousemove={draw} on:mouseup={stop} on:mouseleave={stop}>
	{#each shapes as shape}
		{#if shape.tag === 'rect'}
			<rect x={shape.x} y={shape.y} width={shape.width} height={shape.height} fill={shape.fill} />
		{:else if shape.tag === 'circle'}
			<circle cx={shape.cx} cy={shape.cy} r={shape.r} fill={shape.fill} />
		{:else if shape.tag === 'path'}
			<path d={shape.d} stroke={shape.stroke} fill="none" stroke-width="2" />
		{/if}
	{/each}
	{#if current.d !== ''}
		<path d={current.d} fill="none" stroke="#000000" stroke-width="2" />
	{/if}
</svg>

<style lang="postcss">
	svg {
		width: 100%;
		height: 100vh;
	}
</style>
