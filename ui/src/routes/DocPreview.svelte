<script lang="ts">
	import SlideMd from './SlideMd.svelte'
	import MenuPageNumber from './MenuPageNumber.svelte'

	export let slides: string[]

	let selected = 0
	let content = slides[selected]

	function handlePrev() {
		selected -= 1
		content = slides[selected]
	}

	function handleNext() {
		selected += 1
		content = slides[selected]
	}

	function handleClick(e: KeyboardEvent) {
		if (e.key === 'ArrowLeft') {
			if (selected === 0) {
				return
			}
			handlePrev()
		}
		if (e.key === 'ArrowRight') {
			if (slides.length === selected + 1) {
				return
			}
			handleNext()
		}
	}
</script>

<div class="w-[1000px] h-[700px] px-5 py-8 border-2 mx-auto relative">
	<SlideMd {content} />

	<div class="absolute bottom-5 right-5 z-10">
		<MenuPageNumber {selected} />
	</div>
</div>

<svelte:window on:keydown|preventDefault={handleClick} />
