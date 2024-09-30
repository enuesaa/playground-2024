<script lang="ts">
	import MenuSaveBtn from './MenuSaveBtn.svelte'
	import MenuPrevBtn from './MenuPrevBtn.svelte'
	import MenuNextBtn from './MenuNextBtn.svelte'
	import SlideMd from './SlideMd.svelte'

	export let slides: string[]
	let selected = 0
	let content = slides[selected]
</script>

<section class="w-full h-full flex">
	<div class="w-[1000px] h-[800px] px-5 py-8 border-2 relative">
		<SlideMd {content} />

		<div class="absolute bottom-5 right-5">
			<MenuSaveBtn {slides} {selected} {content} />
			<MenuPrevBtn bind:slides bind:selected bind:content />
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
