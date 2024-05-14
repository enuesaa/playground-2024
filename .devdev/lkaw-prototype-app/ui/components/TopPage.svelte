<script lang="ts">
	import Layout from './Layout.svelte'
	import Code from './Code.svelte'
	import Sidebar from './Sidebar.svelte'
</script>

<Layout>
	<svelte:fragment slot="left">
		<Sidebar />
	</svelte:fragment>
	<svelte:fragment slot="right">
		<Code />
	</svelte:fragment>
</Layout>
