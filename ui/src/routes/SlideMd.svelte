<script lang="ts">
	import Markdown from 'svelte-exmarkdown'

	export let content: string
</script>

<main>
	<Markdown md={content} />
</main>

<!-- svelte-ignore css-unused-selector -->
<style lang="postcss">
	main {
		@apply p-9;
	}
	:global(main) {
		@apply text-2xl;
		@apply overflow-hidden;
	}
	:global(main h1) {
		@apply text-6xl mt-40 mb-5 font-bold;
	}
	:global(main h2) {
		@apply text-5xl my-1 font-bold;
	}
	:global(main p) {
		@apply py-2;
	}
	:global(main ul) {
		@apply py-2 pl-6;
	}
	:global(main ul p) {
		@apply py-1;
	}
	:global(main a) {
		@apply underline;
	}
	:global(main li) {
		list-style-type: disc;
	}
	:global(main img) {
		@apply max-w-full max-h-full;
	}
</style>
