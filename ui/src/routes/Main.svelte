<script lang="ts">
	import Docs from './Docs.svelte'
</script>

<main>
	<div class="container mx-auto px-1 py-8 flex">
		<div class="w-40">
			<Docs />
		</div>
		<div class="flex-auto">
			<slot />
		</div>
	</div>
</main>

<style lang="postcss">
	main {
		@apply bg-gray min-h-screen;
	}
</style>
