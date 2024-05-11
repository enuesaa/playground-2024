<script lang="ts">
	import { Router, Route } from 'svelte-routing'
	import Home from './routes/Home.svelte'
	import About from './routes/About.svelte'
	import './app.css'

	export let url = ''
</script>

<Router {url}>
	<Route path="/about" component={About} />
	<Route path="/"><Home /></Route>
</Router>
