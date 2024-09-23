<script lang="ts">
	import './app.css'
	import Main from './Main.svelte'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'

	const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
	<Main>
		<slot />
	</Main>
</QueryClientProvider>
