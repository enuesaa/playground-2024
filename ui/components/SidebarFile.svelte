<script lang="ts">
	export let filename: string;
</script>

<div>
	{filename}
</div>

<style lang="postcss">
	div {
		user-select: none;
		@apply font-zenkaku text-xl font-medium px-5 py-3 bg-blackgray text-graywhite rounded-xl;
	}
</style>
