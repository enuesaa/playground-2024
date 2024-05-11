<script lang="ts">
	import SidebarFile from '../components/SidebarFile.svelte'
	import Code from '../components/Code.svelte'
	import { fetchFiles } from '../lib/api'
</script>

<main>
	{#await fetchFiles()}
		<div class="left" />
		<div class="right" />
	{:then files}
		<div class="left">
			{#each files as file}
				<SidebarFile filename={file.filename} />
			{/each}
		</div>
		<div class="right">
			<Code code="const a ='b'" />
		</div>
	{/await}
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
