<script lang="ts">
	import './app.css'
	import Footer from './Footer.svelte'
	import Header from './Header.svelte'
</script>

<Header />

<main>
	<div class="container mx-auto px-1 py-8">
		<slot />
	</div>
</main>

<Footer />

<style lang="postcss">
	main {
		@apply bg-gray min-h-screen sm:mx-3 sm:w-[calc(100%_-_24px)] sm:shadow-sm;
	}
</style>
