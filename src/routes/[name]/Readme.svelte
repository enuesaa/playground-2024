<script lang="ts">
	import SvelteMarkdown from 'svelte-markdown'
	export let content: string
</script>

<section>
	<SvelteMarkdown source={content} />
</section>

<style lang="postcss">
	section {
		border: solid 1px rgba(0, 0, 0, 0.2);
		box-shadow: 0 1px 1px rgba(255, 255, 255, 0.2);
		@apply p-3 my-4 bg-grayer;
	}
</style>
