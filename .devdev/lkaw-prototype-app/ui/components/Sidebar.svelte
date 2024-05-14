<script lang="ts">
	import SidebarFile from './SidebarFile.svelte'
	import { filesStore } from '../lib/files'
</script>

{#each $filesStore.files as file}
	<SidebarFile filename={file.filename} />
{/each}
