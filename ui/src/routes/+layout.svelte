<script lang="ts">
	import './app.css'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'

	const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
	<div>
		<slot />
	</div>
</QueryClientProvider>

<style lang="postcss">
	div {
		@apply w-full max-w-[1500px] mx-auto h-screen;
	}
</style>
