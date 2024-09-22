<script lang="ts">
	import '../app.css'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'

	const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
	<main class="mx-auto my-4 max-w-5xl">
		<slot />
	</main>	
</QueryClientProvider>

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.gray.900);
		color: theme(colors.gray.300);
	}
</style>
