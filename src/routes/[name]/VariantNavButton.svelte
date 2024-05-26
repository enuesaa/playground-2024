<script lang="ts">
	export let name: string
	export let showing: string

	function handleClick() {
		showing = name
	}
</script>

<button on:click|preventDefault={handleClick} class={showing === name ? 'bg-graywhite' : 'bg-grayblack'}>
	{name}
</button>

<style lang="postcss">
	button {
		border: solid 2px;
		border-bottom: 0px;
		@apply mt-3 mx-1 p-3 border-grayblack;
		@apply rounded-t-lg;
	}
</style>
