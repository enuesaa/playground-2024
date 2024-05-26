<script lang="ts">
	export let content: string
</script>

<section>
	<pre>{content}</pre>
</section>

<style lang="postcss">
	section {
		border: solid 1px rgba(0, 0, 0, 0.2);
		@apply p-3 my-4 bg-graywhite;
	}
</style>
