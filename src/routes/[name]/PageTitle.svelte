<script lang="ts">
	export let title: string
</script>

<div>
	<h2>{title}</h2>
</div>

<style lang="postcss">
	div {
		@apply mx-2 py-1 mb-6 mt-3;
	}
	h2 {
		@apply font-bold font-zenkaku text-4xl;
	}
</style>
