<script lang="ts">
	export let name: string
	export let title: string
	export let showing: string

	function handleClick() {
		showing = name
	}
</script>

<button on:click|preventDefault={handleClick} class={showing === name ? 'bg-graywhite' : 'bg-grayblack'}>
	{title}
</button>

<style lang="postcss">
	button {
		border: solid 2px;
		border-bottom: 0px;
		overflow: hidden;
		flex: 0 0 auto;
		@apply mt-3 mx-1 px-4 py-3 rounded-t-lg border-grayblack;
		@apply text-base font-semibold;
	}
</style>
