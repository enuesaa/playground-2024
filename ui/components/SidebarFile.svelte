<script lang="ts">
	import { filesStore } from '../lib/files'
	export let filename: string

	function handleClick() {
		filesStore.update(val => {
			return {...val, open: filename}
		})
	}
</script>

{#if $filesStore.open === filename}
	<button style="color: #ff6633;">
		{filename}
	</button>
{:else}
	<button  on:click|preventDefault={handleClick}>
		{filename}
	</button>
{/if}

<style lang="postcss">
	button {
		@apply block w-full text-left;
		@apply font-zenkaku text-lg font-medium m-3 px-5 py-3;
		@apply bg-blackgray text-graywhite rounded-xl;
	}
</style>
