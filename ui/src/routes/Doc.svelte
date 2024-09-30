<script lang="ts">
	import MenuSaveBtn from './MenuSaveBtn.svelte'
	import MenuPrevBtn from './MenuPrevBtn.svelte'
	import MenuNextBtn from './MenuNextBtn.svelte'
	import SlideMd from './SlideMd.svelte'
	import MenuPreviewBtn from './MenuPreviewBtn.svelte'
	import { useUploadFile } from '$lib/api/file'

	export let slides: string[]

	let selected = 0
	let content = slides[selected]

	const uploadFile = useUploadFile()

	async function handleDrop(e: DragEvent) {
		const files = e?.dataTransfer?.files ?? null
		if (files === null) {
			return
		}
		const file = files[0]
		await $uploadFile.mutateAsync(file)
		content += `\n\n![${file.name}](./${file.name})\n\n`
	}
</script>

<section
	class="w-full h-full flex"
	on:drop|preventDefault={handleDrop}
	on:dragover|preventDefault
	on:dragleave|preventDefault
>
	<div class="w-[1000px] h-[800px] px-5 py-8 border-2 relative">
		<SlideMd {content} />

		<div class="absolute bottom-5 right-5 z-10">
			<MenuPreviewBtn />
			<MenuSaveBtn {slides} {selected} {content} />
		</div>
		<div class="absolute top-0 left-[0px] w-[150px] h-full">
			<MenuPrevBtn bind:slides bind:selected bind:content />
		</div>
		<div class="absolute top-0 left-[850px] w-[150px] h-full">
			<MenuNextBtn bind:slides bind:selected bind:content />
		</div>
	</div>
	<div class="flex-auto">
		<textarea bind:value={content} />
	</div>
</section>

<style lang="postcss">
	textarea {
		@apply block h-full bg-[rgba(0,0,0,0)];
		@apply font-normal w-full p-3 text-black text-lg outline-none;
	}
</style>
