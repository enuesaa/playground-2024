<script lang="ts">
	import SlideMd from './SlideMd.svelte'

	export let slides: string[]
	export let selected: number

	if (slides.length <= selected) {
		slides.push('')
	}

	let content = slides[selected]
	$: {
		slides[selected] = content

		// if (slides.length > selected) {
		// 	slides[selected] = content
		// } else {
		// 	slides.push(content)
		// }
	}
</script>

{selected}
<section class="w-full flex h-full">
	<div class="w-[1000px] h-[800px] px-5 py-8 border-2">
		<SlideMd {content} />
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
