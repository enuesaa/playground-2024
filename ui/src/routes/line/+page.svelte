<script lang="ts">
	type Line = {
		x1: number;
		y1: number;
		x2: number;
		y2: number;
	}
    let lines: Line[] = [];
    const lineLength = 100;

    function handleClick(e: MouseEvent & { currentTarget: EventTarget & SVGSVGElement }) {
        const rect = e.currentTarget.getBoundingClientRect()
        const x = e.clientX - rect.left;
        const y = e.clientY - rect.top;
		lines = [...lines, {
			x1: x,
			y1: y,
			x2: x + lineLength,
			y2: y,
		}]
    }
</script>

<svg on:click={handleClick}>
	{#each lines as {x1, y1, x2, y2}}
		<line x1={x1} y1={y1} x2={x2} y2={y2} stroke="black" stroke-width="2" />
	{/each}
</svg>

<style lang="postcss">
	svg {
		width: 100%;
		height: 100vh;
	}
</style>
