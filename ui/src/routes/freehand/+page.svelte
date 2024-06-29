<script lang="ts">
    type Path = {d: string}
    let paths: Path[] = []
    let current: Path = {d: ''}

    function start(e: MouseEvent & {currentTarget: EventTarget & SVGSVGElement}) {
        const { left, top } = e.currentTarget.getBoundingClientRect()
        const x = e.clientX - left
        const y = e.clientY - top
        current.d = `M${x},${y}`
    }

    function draw(e: MouseEvent & {currentTarget: EventTarget & SVGSVGElement}) {
        const { left, top } = e.currentTarget.getBoundingClientRect()
        const x = e.clientX - left
        const y = e.clientY - top
        current.d += ` L${x},${y}`
    }

    function stop() {
        paths = [...paths, {d: current.d}]
        current.d = ''
    }

    function del(i: number) {
        paths = paths.filter((_, k) => k !== i)
    }
</script>

<svg 
    on:mousedown={start}
    on:mousemove={draw}
    on:mouseup={stop}
    on:mouseleave={stop}
>
    {#each paths as path, i}
        <path d={path.d} fill="none" stroke="#000000" stroke-width="2" on:click|preventDefault={() => del(i)} />
    {/each}
    {#if current.d !== ''}
        <path d={current.d} fill="none" stroke="#000000" stroke-width="2" />
    {/if}
</svg>

<style lang="postcss">
    svg {
        width: 100%;
        height: 100vh;
        cursor: crosshair;
    }
</style>
