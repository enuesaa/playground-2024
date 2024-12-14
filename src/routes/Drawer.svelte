<script lang="ts">
	import { toPng, toJpeg, toBlob, toPixelData, toSvg } from 'html-to-image'

	let canvas: undefined|HTMLDivElement = undefined;
	let style = $state('')
	let classe = $state('bg-black')
	let border = $state('none')
	let imgsrc = $state('')

	$effect(() => {
		console.log(border)
		style = `border: ${border};`
	})

	async function handleClick() {
		if (canvas === undefined) {
			return
		}
		const dataurl = await toPng(canvas)
		imgsrc = dataurl
	}
</script>

border
<input type="text" bind:value={border} />

class
<input type="text" bind:value={classe} />

<div {style} class={classe} bind:this={canvas}>
	hello
</div>

<button onclick={handleClick}>toPng</button>

<img src={imgsrc} alt='' />
