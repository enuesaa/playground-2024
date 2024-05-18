<script lang="ts">
	import './app.css'
	import Header from './Header.svelte'
	import WipBanner from './WipBanner.svelte'
</script>

<Header />
<WipBanner />

<div class="container mx-auto p-3 mt-5">
	<slot />
</div>
