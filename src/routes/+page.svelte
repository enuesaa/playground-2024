<script lang="ts">
	import TrailerLink from './TrailerLink.svelte'
	import type { PageData } from './$types'

	export let data: PageData
</script>

{#each data.projects as project}
	<TrailerLink href={`/${project.name}`} title={project.name} />
{/each}
