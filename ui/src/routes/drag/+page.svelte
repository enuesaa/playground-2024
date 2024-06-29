<script lang="ts">
    let rect = {
		x: 50,
		y: 50,
		width: 100,
		height: 100,
	}
    let dragging = false;
    let offsetX = 0;
    let offsetY = 0;

    function start(e: MouseEvent) {
        dragging = true;
        offsetX = e.clientX - rect.x;
        offsetY = e.clientY - rect.y;
    }

    function drag(e: MouseEvent & {currentTarget: EventTarget & SVGSVGElement}) {
        if (!dragging) {
			return
		}
        rect.x = e.clientX - offsetX;
        rect.y = e.clientY - offsetY;
    }

    function stop(e: MouseEvent & {currentTarget: EventTarget & SVGSVGElement}) {
        dragging = false;
    }
</script>

<svg 
    on:mousemove={drag}
    on:mouseup={stop}
    on:mouseleave={stop}
>
    <rect	
		x={rect.x} y={rect.y} width={rect.width} height={rect.height} fill="#000000"
        on:mousedown={start}
    />
</svg>

<style lang="postcss">
    svg {
        width: 100%;
        height: 100vh;
    }
</style>
