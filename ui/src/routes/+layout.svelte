<script lang="ts">
	import './app.css'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'

	const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
	<main>
		<slot />
	</main>
</QueryClientProvider>

<style lang="postcss">
	main {
		@apply w-full max-w-[1500px] mx-auto h-screen;
	}
</style>
