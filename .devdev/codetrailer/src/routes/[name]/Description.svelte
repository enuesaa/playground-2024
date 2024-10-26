<script lang="ts">
	export let content: string
</script>

<section>
	<pre>{content}</pre>
</section>

<style lang="postcss">
	section {
		@apply py-3 px-5 my-4 font-zenmaru text-base w-full;
	}
	section pre {
		letter-spacing: 0.1px;
		line-height: 1.8;
		overflow-wrap: anywhere;
		word-break: normal;
		line-break: strict;
		white-space: pre-wrap;
  		font-feature-settings: "palt";
		@apply font-ibmplex;
	}
</style>
