<script lang="ts">
	type Circle = {
		cx: number;
		cy: number;
		r: number;
	}
    let circles: Circle[] = [];

    function handleClick(e: MouseEvent & {currentTarget: EventTarget & SVGSVGElement}) {
        const { left, top } = e.currentTarget.getBoundingClientRect()
		const circle = {
			cx: e.clientX - left,
			cy: e.clientY - top,
			r: 20,
		}
        circles = [...circles, circle];
    }
</script>

<svg on:click|preventDefault={handleClick}>
    {#each circles as { cx, cy, r }}
        <circle cx={cx} cy={cy} r={r} fill="#000000" />
    {/each}
</svg>

<style lang="postcss">
    svg {
        width: 100%;
        height: 100vh;
    }
</style>
