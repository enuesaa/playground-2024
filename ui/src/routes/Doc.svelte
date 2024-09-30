<script lang="ts">
	import Slide from './Slide.svelte'
	import { useUpdateDoc } from '$lib/apidoc'
	import { createRegistry } from '$lib/drawing'
	import DrawerMenu from '$lib/components/drawer/DrawerMenu.svelte'
	import type { SlideSchema } from '$lib/apidoc'
	import { convertToSvg } from '$lib/convert'

	export let slides: SlideSchema[]
	let selected = 0
	let content = slides[selected].content
	let registry = createRegistry()

	const updateDoc = useUpdateDoc()

	async function handleSave() {
		const drawing = convertToSvg(registry)
		slides[selected] = { content, drawing }
		await $updateDoc.mutateAsync({slides})
	}

	function handleNext() {
		const drawing = convertToSvg(registry)
		slides[selected] = { content, drawing }
		selected += 1

		if (slides.length <= selected) {
			slides.push({content: '', drawing: ''})
		}
		content = slides[selected].content
		registry = createRegistry()
	}
	function handlePrev() {
		const drawing = convertToSvg(registry)
		slides[selected] = { content, drawing }
		selected -= 1
		content = slides[selected].content
		registry = createRegistry()
	}
</script>

<button on:click|preventDefault={handleSave}>save</button>
<button on:click|preventDefault={handlePrev}>prev</button>
<button on:click|preventDefault={handleNext}>next</button>

<DrawerMenu bind:registry />

<Slide bind:content bind:registry />
