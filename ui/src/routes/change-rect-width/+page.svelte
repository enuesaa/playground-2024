<script lang="ts">
    let rect = { x: 50, y: 50, width: 100, height: 100 };
    let resizing = false;
    let initialMouseX = 0;
    let initialWidth = 0;

    function start(e: MouseEvent) {
        resizing = true;
        initialMouseX = e.clientX;
        initialWidth = rect.width; // どんどん伸びていくから初期値を保持している
    }

    function resize(e: MouseEvent) {
        if (!resizing) {
			return
		}
        rect.width = initialWidth + e.clientX - initialMouseX;
    }

    function stop() {
        resizing = false;
    }
</script>

<svg 
    on:mousemove={resize}
    on:mouseup={stop}
    on:mouseleave={stop}
>
    <rect 
        x={rect.x} 
        y={rect.y} 
        width={rect.width} 
        height={rect.height} 
		fill="#000000"
    />
    <rect 
		fill="#ff9933"
        x={rect.x + rect.width - 5}
        y={rect.y}
        width="10"
        height={rect.height}
        on:mousedown={start}
    />
</svg>

<style lang="postcss">
    svg {
        width: 100%;
        height: 100vh;
    }
</style>
