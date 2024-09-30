<script lang="ts">
	import Markdown from 'svelte-exmarkdown'

	export let content: string
</script>

<div class="appdoc">
	<Markdown md={content} />
</div>

<!-- svelte-ignore css-unused-selector -->
<style lang="postcss">
	.appdoc {
		@apply py-1 px-2;
	}
	:global(.appdoc h1) {
		@apply text-6xl mt-40 mb-5 font-bold;
	}
	:global(.appdoc h2) {
		@apply text-5xl my-1 font-bold;
	}
	:global(.appdoc p) {
		@apply my-1;
	}
	:global(.appdoc ul) {
		@apply my-1 pl-3;
	}
	:global(.appdoc li) {
		list-style-type: "- ";
	}
</style>
