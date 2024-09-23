<script lang="ts">
	import Docs from './Docs.svelte'
	import Header from './Header.svelte'
</script>

<main>
	<div class="w-40 px-5 py-8">
		<Header />
		<Docs />
	</div>
	<div class="flex-auto">
		<slot />
	</div>
</main>

<style lang="postcss">
	main {
		@apply w-full max-w-[1300px] mx-auto h-screen flex overflow-y-hidden;
	}
</style>
