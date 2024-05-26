<script lang="ts">
	export let title: string
</script>

<div>
	<h2>{title}</h2>
</div>

<style lang="postcss">
	div {
		@apply mx-2 py-1 mt-3 mb-1;
	}
	h2 {
		@apply font-bold font-zenkaku text-3xl sm:text-4xl px-2;
		line-height: 1.5;
	}
</style>
