<script lang="ts">
	import './app.css'
	import Footer from './Footer.svelte'
	import Header from './Header.svelte'
	import WipBanner from './WipBanner.svelte'
</script>

<Header />
<WipBanner />

<div class="container mx-auto p-1 mt-5">
	<slot />
</div>

<Footer />
