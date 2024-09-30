<script lang="ts">
	import Slide from './Slide.svelte'
	import { useUpdateDoc } from '$lib/api/doc'
	import type { SlideSchema } from '$lib/api/doc'
	import { ArrowLeftIcon, ArrowRightIcon, SaveIcon } from 'svelte-feather-icons'

	export let slides: SlideSchema[]
	let selected = 0
	let content = slides[selected].content

	const updateDoc = useUpdateDoc()

	async function handleSave() {
		slides[selected] = { content }
		await $updateDoc.mutateAsync({slides})
	}

	function handleNext() {
		slides[selected] = { content }
		selected += 1
		content = slides[selected].content
	}
	function handlePrev() {
		slides[selected] = { content }
		selected -= 1
		content = slides[selected].content
	}
</script>

<div class="w-full h-full relative">
	<Slide bind:content />

	<div class="absolute bottom-0 right-0">
		<button on:click|preventDefault={handleSave}><SaveIcon /></button>
		{#if selected === 0}
			<button disabled class="bg-blackgray text-gray rounded-md"><ArrowLeftIcon /></button>
		{:else}
			<button on:click|preventDefault={handlePrev}><ArrowLeftIcon /></button>
		{/if}
		<button on:click|preventDefault={handleNext}><ArrowRightIcon /></button>		
	</div>
</div>
