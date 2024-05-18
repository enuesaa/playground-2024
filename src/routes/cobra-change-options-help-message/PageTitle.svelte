<script lang="ts">
	export let title: string
</script>

<div class="mx-2 py-1 mb-2">
	<h2>{title}</h2>
</div>

<style lang="postcss">
	h2 {
		@apply font-bold font-zenkaku text-3xl;
	}
</style>
