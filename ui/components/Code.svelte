<script lang="ts">
	import { HighlightAuto } from 'svelte-highlight'
	import githubDark from "svelte-highlight/styles/github-dark";
    import ashes from "svelte-highlight/styles/ashes";

	export let code: string
</script>

<svelte:head>
  {@html ashes}
</svelte:head>

<div>
	<HighlightAuto {code} />
</div>

<style lang="postcss">
	div {
		@apply text-xl font-medium px-5 py-3 bg-blackgray text-graywhite rounded-xl;
		@apply w-full h-full;
	}
</style>
