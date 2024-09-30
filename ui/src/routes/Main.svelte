<script lang="ts">
	import Docs from './Docs.svelte'
</script>

<main>
	<div class="w-20 px-5 py-8 border">
		<Docs />
	</div>
	<div class="flex-auto border">
		<slot />
	</div>
</main>

<style lang="postcss">
	main {
		@apply w-full max-w-[1300px] mx-auto h-screen flex overflow-y-hidden;
	}
</style>
