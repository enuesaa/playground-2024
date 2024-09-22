<script lang="ts">
	import Files from './Files.svelte'
	import GoToParentDir from './GoToParentDir.svelte'
</script>

<main>
	<div class="container mx-auto px-1 py-8 flex">
		<div class="w-40">
			<GoToParentDir path='./' />
			<Files path='./' />
		</div>
		<div class="flex-auto">
			<slot />
		</div>
	</div>
</main>

<style lang="postcss">
	main {
		@apply bg-gray min-h-screen sm:mx-3 sm:w-[calc(100%_-_24px)] sm:shadow-sm;
	}
</style>
