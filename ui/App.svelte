<script lang="ts">
	import { Router, Route } from 'svelte-routing'
	import TopPage from './components/TopPage.svelte'
	import './app.css'

	export let url = ''
</script>

<Router {url}>
	<Route path="/"><TopPage /></Route>
</Router>	
