<script lang="ts">
	import { toPng } from 'html-to-image'
	import type { EventHandler } from 'svelte/elements'

	type Props = {
		canvas: HTMLDivElement
	}
	let { canvas }: Props = $props()

	let imgsrc = $state('')

	const handleClick: EventHandler = async (e) => {
		e.preventDefault()
		if (canvas === undefined) {
			return
		}

		const dataurl = await toPng(canvas)
		imgsrc = dataurl
	}
</script>

<button onclick={handleClick}>convert</button>

<img src={imgsrc} alt="" />
