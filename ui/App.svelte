<script lang="ts">
	import { Router, Route } from 'svelte-routing'
	import Home from './routes/Home.svelte'
	import './app.css'
	import Layout from './Layout.svelte'

	export let url = ''
</script>

<Layout>
	<Router {url}>
		<Route path="/"><Home /></Route>
	</Router>	
</Layout>
