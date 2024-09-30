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
		@apply py-1 px-2;
	}
	:global(main h1) {
		@apply text-6xl mt-40 mb-5 font-bold;
	}
	:global(main h2) {
		@apply text-5xl my-1 font-bold;
	}
	:global(main p) {
		@apply my-1;
	}
	:global(main ul) {
		@apply my-1 pl-3;
	}
	:global(main li) {
		list-style-type: "- ";
	}
</style>
