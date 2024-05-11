<script lang="ts">
	import SidebarFile from '../components/SidebarFile.svelte'
	import Code from '../components/Code.svelte'
	import { fetchFiles } from '../lib/api'
</script>

<main>
	{#await fetchFiles()}
		a
	{/await}
	<div class="left">
		<SidebarFile filename="a.md" />
	</div>
	<div class="right">
		<Code code="const a ='b'" />
	</div>
</main>

<style lang="postcss">
	main {
		height: 85vh;
		@apply flex gap-5 mt-9 mx-auto p-3;
	}
	.left {
		@apply w-48;
	}
	.right {
		@apply flex-auto;
	}
</style>
